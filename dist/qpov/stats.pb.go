// Code generated by protoc-gen-go.
// source: stats.proto
// DO NOT EDIT!

package qpov

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StatsCPUTime struct {
	UserSeconds   int64 `protobuf:"varint,2,opt,name=user_seconds" json:"user_seconds,omitempty"`
	SystemSeconds int64 `protobuf:"varint,3,opt,name=system_seconds" json:"system_seconds,omitempty"`
}

func (m *StatsCPUTime) Reset()         { *m = StatsCPUTime{} }
func (m *StatsCPUTime) String() string { return proto.CompactTextString(m) }
func (*StatsCPUTime) ProtoMessage()    {}

type MachineStats struct {
	// E.g. "Amazon/c4.8xlarge 36 x Intel(R) Xeon(R) CPU E5-2666 v3 @ 2.90GHz".
	ArchSummary string        `protobuf:"bytes,1,opt,name=arch_summary" json:"arch_summary,omitempty"`
	Cpu         string        `protobuf:"bytes,2,opt,name=cpu" json:"cpu,omitempty"`
	NumCpu      int32         `protobuf:"varint,3,opt,name=num_cpu" json:"num_cpu,omitempty"`
	Cloud       *Cloud        `protobuf:"bytes,4,opt,name=cloud" json:"cloud,omitempty"`
	CpuTime     *StatsCPUTime `protobuf:"bytes,5,opt,name=cpu_time" json:"cpu_time,omitempty"`
	Jobs        int64         `protobuf:"varint,6,opt,name=jobs" json:"jobs,omitempty"`
}

func (m *MachineStats) Reset()         { *m = MachineStats{} }
func (m *MachineStats) String() string { return proto.CompactTextString(m) }
func (*MachineStats) ProtoMessage()    {}

func (m *MachineStats) GetCloud() *Cloud {
	if m != nil {
		return m.Cloud
	}
	return nil
}

func (m *MachineStats) GetCpuTime() *StatsCPUTime {
	if m != nil {
		return m.CpuTime
	}
	return nil
}

type BatchStats struct {
	BatchId string `protobuf:"bytes,1,opt,name=batch_id" json:"batch_id,omitempty"`
	Total   int64  `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	Done    int64  `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	Comment string `protobuf:"bytes,4,opt,name=comment" json:"comment,omitempty"`
	Ctime   int64  `protobuf:"varint,5,opt,name=ctime" json:"ctime,omitempty"`
}

func (m *BatchStats) Reset()         { *m = BatchStats{} }
func (m *BatchStats) String() string { return proto.CompactTextString(m) }
func (*BatchStats) ProtoMessage()    {}

type StatsOverall struct {
	// When these stats were calculated.
	StatsTimestamp int64 `protobuf:"varint,1,opt,name=stats_timestamp" json:"stats_timestamp,omitempty"`
	// User and system time.
	CpuTime *StatsCPUTime `protobuf:"bytes,2,opt,name=cpu_time" json:"cpu_time,omitempty"`
	// (user_seconds+system_seconds) / CPU cores.
	MachineTime *StatsCPUTime `protobuf:"bytes,3,opt,name=machine_time" json:"machine_time,omitempty"`
	// Total time completed leases have been outstanding.
	LeaseSeconds int64 `protobuf:"varint,4,opt,name=lease_seconds" json:"lease_seconds,omitempty"`
	// Split out by machine type.
	MachineStats []*MachineStats `protobuf:"bytes,5,rep,name=machine_stats" json:"machine_stats,omitempty"`
	BatchStats   []*BatchStats   `protobuf:"bytes,6,rep,name=batch_stats" json:"batch_stats,omitempty"`
}

func (m *StatsOverall) Reset()         { *m = StatsOverall{} }
func (m *StatsOverall) String() string { return proto.CompactTextString(m) }
func (*StatsOverall) ProtoMessage()    {}

func (m *StatsOverall) GetCpuTime() *StatsCPUTime {
	if m != nil {
		return m.CpuTime
	}
	return nil
}

func (m *StatsOverall) GetMachineTime() *StatsCPUTime {
	if m != nil {
		return m.MachineTime
	}
	return nil
}

func (m *StatsOverall) GetMachineStats() []*MachineStats {
	if m != nil {
		return m.MachineStats
	}
	return nil
}

func (m *StatsOverall) GetBatchStats() []*BatchStats {
	if m != nil {
		return m.BatchStats
	}
	return nil
}
