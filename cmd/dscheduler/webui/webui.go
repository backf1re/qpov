package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"sync"
	"time"

	"github.com/ThomasHabets/go-uuid/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/ThomasHabets/qpov/dist"
	pb "github.com/ThomasHabets/qpov/dist/qpov"
)

const (
	userAgent = "dscheduler-webui"

	rootTmpl = `
<style>
td {
  text-align: right;
}
.fixed {
  font-family: monospace;
}
</style>
<h1>QPov</h1>

{{if .Errors}}
  <h2>Errors while rendering this page:</h2>
  <ul>
    {{range .Errors}}
      {{.}}
    {{end}}
  </ul>
{{end}}

<h2>Scheduler stats</h2>
<table>
<tr><th colspan="2">Orders</th></tr>
<tr><th>Total</th><td>{{.Stats.SchedulingStats.Orders}}</td></tr>
<tr><th>Active</th><td>{{.Stats.SchedulingStats.ActiveOrders}}</td></tr>
<tr><th>Done</th><td>{{.Stats.SchedulingStats.DoneOrders}}</td></tr>
<tr><th>Unstarted</th><td>{{.UnstartedOrders}}</td></tr>
<tr><th colspan="2">Leases</th></tr>
<tr><th>Total</th><td>{{.Stats.SchedulingStats.Leases}}</td></tr>
<tr><th>Active</th><td>{{.Stats.SchedulingStats.ActiveLeases}}</td></tr>
<tr><th>Done</th><td>{{.Stats.SchedulingStats.DoneLeases}}</td></tr>
</table>

<h2>Active leases</h2>
<table>
<tr>
  <th>Order</th>
  <th>Created</th>
  <th>Lifetime</th>
  <th>Updated</th>
  <th>Expires</th>
</tr>
{{range .Leases}}
<tr>
  <td class="fixed">{{.OrderId}}</td>
  <td>{{.CreatedMs|fmsdate "2006-01-02 15:04"}}</td>
  <td>{{.CreatedMs|fmssince}}</td>
  <td>{{.UpdatedMs|fmssince}}</td>
  <td>{{.ExpiresMs|fmsuntil}}</td>
</tr>
{{end}}
</table>

<h2>Finished</h2>
<table>
<tr>
  <th>Order</th>
  <th>Created</th>
  <th>Done</th>
  <th>Time</th>
</tr>
{{range .DoneLeases}}
<tr>
  <td class="fixed">{{.OrderId}}</td>
  <td>{{.CreatedMs|fmsdate "2006-01-02 15:04"}}</td>
  <td>{{.UpdatedMs|fmsdate "2006-01-02 15:04"}}</td>
  <td>{{.CreatedMs|fmssub .UpdatedMs}}</td>
</tr>
{{end}}
</table>

<hr>
Page server time: {{.PageTime}}
`
)

var (
	pageDeadline = flag.Duration("page_deadline", time.Second, "Page timeout.")
	socketPath   = flag.String("socket", "", "Unix socket to listen to.")
	caFile       = flag.String("ca_file", "", "Server CA file.")
	certFile     = flag.String("cert_file", "", "Client cert file.")
	keyFile      = flag.String("key_file", "", "Client key file.")
	schedAddr    = flag.String("scheduler", "", "Scheduler address.")

	sched    pb.SchedulerClient
	tmplRoot template.Template
)

func httpContext(r *http.Request) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "source", "http")
	ctx = context.WithValue(ctx, "id", uuid.New())
	ctx = context.WithValue(ctx, "http:remote_addr", r.RemoteAddr)
	return ctx
}

func fmsdate(s string, ms int64) string {
	return time.Unix(ms/1000, 0).Format(s)
}

func fmssub(a, b int64) string {
	return time.Unix(a/1000, 0).Sub(time.Unix(b/1000, 0)).String()
}

func fmsuntil(ms int64) string {
	now := time.Now().UnixNano() / 1000000000
	return time.Unix(ms/1000, 0).Sub(time.Unix(now, 0)).String()
}

func fmssince(ms int64) string {
	now := time.Now().UnixNano() / 1000000000
	return time.Unix(now, 0).Sub(time.Unix(ms/1000, 0)).String()
}

func getLeases(ctx context.Context, done bool) ([]*pb.Lease, error) {
	stream, err := sched.Leases(ctx, &pb.LeasesRequest{Done: done})
	if err != nil {
		return nil, err
	}
	var leases []*pb.Lease
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		leases = append(leases, r.Lease)
	}
	return leases, nil
}

func handleRoot(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	startTime := time.Now()
	var errs []error
	var m sync.Mutex
	var wg sync.WaitGroup

	// Get Stats.
	var st *pb.StatsReply
	wg.Add(1)
	go func() {
		var err error
		defer wg.Done()
		st, err = sched.Stats(ctx, &pb.StatsRequest{
			SchedulingStats: true,
		})
		if err != nil {
			m.Lock()
			defer m.Unlock()
			errs = append(errs, err)
			return
		}
	}()

	// Get active Leases.
	var leases []*pb.Lease
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		leases, err = getLeases(ctx, false)
		if err != nil {
			m.Lock()
			defer m.Unlock()
			errs = append(errs, err)
			return
		}
	}()

	// Get done Leases.
	var doneLeases []*pb.Lease
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		doneLeases, err = getLeases(ctx, true)
		if err != nil {
			m.Lock()
			defer m.Unlock()
			errs = append(errs, err)
			return
		}
	}()
	wg.Wait()
	if len(errs) > 0 {
		log.Printf("Errors: %v", errs)
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	return &struct {
		Stats           *pb.StatsReply
		Leases          []*pb.Lease
		DoneLeases      []*pb.Lease
		UnstartedOrders int64
		Errors          []error
		PageTime        time.Duration
	}{
		Stats:           st,
		Leases:          leases,
		DoneLeases:      doneLeases,
		Errors:          errs,
		UnstartedOrders: st.SchedulingStats.Orders - st.SchedulingStats.DoneOrders,
		PageTime:        time.Since(startTime),
	}, nil
}

type handleFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error)
type handler struct {
	tmpl *template.Template
	f    handleFunc
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Strict-Transport-Security", "max-age=2592000")
	ctx, cancel := context.WithTimeout(httpContext(r), *pageDeadline)
	defer cancel()
	data, err := h.f(ctx, w, r)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		// TODO: propagate correct error code.
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Some random error occured.")
		return
	}
	var buf bytes.Buffer
	if err := h.tmpl.Execute(&buf, data); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Template rendering failed: %v", err)
		fmt.Fprintf(w, "Internal error: Failed to render page")
		return
	}
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("Failed to write page to network: %v", err)
		return
	}
}

func wrap(f handleFunc, t string) *handler {
	tmpl := template.New("blah")
	tmpl.Funcs(template.FuncMap{
		"fmsdate":  fmsdate,
		"fmsuntil": fmsuntil,
		"fmssince": fmssince,
		"fmssub":   fmssub,
	})
	template.Must(tmpl.Parse(t))
	return &handler{f: f, tmpl: tmpl}
}

func connectScheduler(addr string) error {
	caStr := dist.CacertClass1
	if *caFile != "" {
		b, err := ioutil.ReadFile(*caFile)
		if err != nil {
			return fmt.Errorf("reading %q: %v", *caFile, err)
		}
		caStr = string(b)
	}

	// Root CA.
	cp := x509.NewCertPool()
	if ok := cp.AppendCertsFromPEM([]byte(caStr)); !ok {
		return fmt.Errorf("failed to add root CAs")
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("failed to split host/port out of %q", addr)
	}

	tlsConfig := tls.Config{
		ServerName: host,
		RootCAs:    cp,
	}

	// Client Cert.
	if *certFile != "" {
		cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
		if err != nil {
			return fmt.Errorf("failed to load client keypair %q/%q: %v", *certFile, *keyFile, err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	cr := credentials.NewTLS(&tlsConfig)
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(cr),
		grpc.WithUserAgent(userAgent),
	)
	if err != nil {
		return fmt.Errorf("dialing scheduler %q: %v", addr, err)
	}
	sched = pb.NewSchedulerClient(conn)
	return nil
}

func main() {
	flag.Parse()

	if err := connectScheduler(*schedAddr); err != nil {
		log.Fatalf("Failed to connect to scheduler: %v", err)
	}

	sock, err := net.Listen("unix", *socketPath)
	if err != nil {
		log.Fatalf("Unable to listen to socket: %v", err)
	}
	if err = os.Chmod(*socketPath, 0666); err != nil {
		log.Fatalf("Unable to chmod socket: %v", err)
	}

	r := mux.NewRouter()
	r.Handle("/", wrap(handleRoot, rootTmpl)).Methods("GET")
	log.Printf("Running dscheduler webui...")
	if err := fcgi.Serve(sock, r); err != nil {
		log.Fatal("Failed to start serving fcgi: ", err)
	}
}