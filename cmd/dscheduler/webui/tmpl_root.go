// -*- html -*-
package main

const rootTmpl = `
{{ $root := . }}
<h2>Scheduler stats</h2>
    {{if .Stats.SchedulingStats}}
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
    {{end}}

    <h2>Active leases</h2>
    <table>
      <tr>
        <th>Order</th>
        <th>Lease</th>
        <th>Created</th>
        <th>Lifetime</th>
        <th>Updated</th>
        <th>Expires</th>
        <!--  <th>Package</th> -->
        <th>File</th>
        <th>Client</th>
      </tr>
      {{range .Leases}}
        <tr>
          <td nowrap class="fixed">{{.OrderId}}</td>
          <td nowrap class="fixed"><a href="{{$root.Root}}/lease/{{.LeaseId}}">{{.LeaseId}}</a></td>
          <td nowrap>{{.CreatedMs|fmsdate "2006-01-02 15:04"}}</td>
          <td nowrap>{{.CreatedMs|fmssince}}</td>
          <td nowrap>{{.UpdatedMs|fmssince}}</td>
          <td nowrap>{{.ExpiresMs|fmsuntil}}</td>
          <!--  <td nowrap>{{.Order.Package|fileonly}}</td> -->
          <td nowrap>{{.Order.File}}</td>
          <td nowrap>{{.Address}} {{.Hostname}}</td>
        </tr>
      {{end}}
    </table>

    <h2>Batches</h2>
    <table>
      <tr>
        <th>Batch</th>
        <th>Done</th>
        <th>Total</th>
        <th>Completion</th>
      </tr>
      {{range .StatsOverall.BatchStats}}
      <tr>
        <td>{{.BatchId}}</td>
        <td>{{.Done}}</td>
        <td>{{.Total}}</td>
        <td>{{fmtpercent .Done .Total}}%</td>
      </tr>
      {{end}}
    </table>

    <h2>Finished</h2>
    <table>
      <tr>
        <th>Order</th>
        <th>Lease</th>
        <th>Created</th>
        <th>Done</th>
        <th>Time</th>
        <th>Image</th>
        <!--  <th>Package</th> -->
        <th>File</th>
        <!--  <th>Args</th> -->
        <th>Client</th>
      </tr>
      {{range .DoneLeases}}
        <tr>
          <td nowrap class="fixed">{{.OrderId}}</td>
          <td nowrap class="fixed"><a href="{{$root.Root}}/lease/{{.LeaseId}}">{{.LeaseId}}</a></td>
          <td nowrap>{{.CreatedMs|fmsdate "2006-01-02 15:04"}}</td>
          <td nowrap>{{.UpdatedMs|fmsdate "2006-01-02 15:04"}}</td>
          <td nowrap>{{.CreatedMs|fmssub .UpdatedMs}}</td>
          <td nowrap><a href="{{$root.Root}}/image/{{.LeaseId}}">Image</a></td>
          <!--  <td nowrap>{{.Order.Package|fileonly}}</td> -->
          <td nowrap>{{.Order.File}}</td>
          <td nowrap>{{.Address}} {{.Hostname}}</td>
          <!--  <td nowrap>{{.Order.Args}}</td> -->
        </tr>
      {{end}}
    </table>
`
