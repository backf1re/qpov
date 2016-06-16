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
	UserSeconds   int64 `protobuf:"varint,2,opt,name=user_seconds,json=userSeconds" json:"user_seconds,omitempty"`
	SystemSeconds int64 `protobuf:"varint,3,opt,name=system_seconds,json=systemSeconds" json:"system_seconds,omitempty"`
	// Standardized compute core-seconds.
	// Baseline CPU: One core from a "Core(TM)2 Quad CPU Q6600 @ 2.40GHz".
	ComputeSeconds int64 `protobuf:"varint,4,opt,name=compute_seconds,json=computeSeconds" json:"compute_seconds,omitempty"`
}

func (m *StatsCPUTime) Reset()                    { *m = StatsCPUTime{} }
func (m *StatsCPUTime) String() string            { return proto.CompactTextString(m) }
func (*StatsCPUTime) ProtoMessage()               {}
func (*StatsCPUTime) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type MachineStats struct {
	// E.g. "Amazon/c4.8xlarge 36 x Intel(R) Xeon(R) CPU E5-2666 v3 @ 2.90GHz".
	ArchSummary string        `protobuf:"bytes,1,opt,name=arch_summary,json=archSummary" json:"arch_summary,omitempty"`
	Cpu         string        `protobuf:"bytes,2,opt,name=cpu" json:"cpu,omitempty"`
	NumCpu      int32         `protobuf:"varint,3,opt,name=num_cpu,json=numCpu" json:"num_cpu,omitempty"`
	Cloud       *Cloud        `protobuf:"bytes,4,opt,name=cloud" json:"cloud,omitempty"`
	CpuTime     *StatsCPUTime `protobuf:"bytes,5,opt,name=cpu_time,json=cpuTime" json:"cpu_time,omitempty"`
	Jobs        int64         `protobuf:"varint,6,opt,name=jobs" json:"jobs,omitempty"`
}

func (m *MachineStats) Reset()                    { *m = MachineStats{} }
func (m *MachineStats) String() string            { return proto.CompactTextString(m) }
func (*MachineStats) ProtoMessage()               {}
func (*MachineStats) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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
	BatchId string        `protobuf:"bytes,1,opt,name=batch_id,json=batchId" json:"batch_id,omitempty"`
	Total   int64         `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	Done    int64         `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	Comment string        `protobuf:"bytes,4,opt,name=comment" json:"comment,omitempty"`
	Ctime   int64         `protobuf:"varint,5,opt,name=ctime" json:"ctime,omitempty"`
	CpuTime *StatsCPUTime `protobuf:"bytes,6,opt,name=cpu_time,json=cpuTime" json:"cpu_time,omitempty"`
}

func (m *BatchStats) Reset()                    { *m = BatchStats{} }
func (m *BatchStats) String() string            { return proto.CompactTextString(m) }
func (*BatchStats) ProtoMessage()               {}
func (*BatchStats) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *BatchStats) GetCpuTime() *StatsCPUTime {
	if m != nil {
		return m.CpuTime
	}
	return nil
}

type StatsOverall struct {
	// When these stats were calculated.
	StatsTimestamp int64 `protobuf:"varint,1,opt,name=stats_timestamp,json=statsTimestamp" json:"stats_timestamp,omitempty"`
	// User and system time.
	CpuTime *StatsCPUTime `protobuf:"bytes,2,opt,name=cpu_time,json=cpuTime" json:"cpu_time,omitempty"`
	// (user_seconds+system_seconds) / CPU cores.
	MachineTime *StatsCPUTime `protobuf:"bytes,3,opt,name=machine_time,json=machineTime" json:"machine_time,omitempty"`
	// Total time completed leases have been outstanding.
	LeaseSeconds int64 `protobuf:"varint,4,opt,name=lease_seconds,json=leaseSeconds" json:"lease_seconds,omitempty"`
	// Split out by machine type.
	MachineStats []*MachineStats `protobuf:"bytes,5,rep,name=machine_stats,json=machineStats" json:"machine_stats,omitempty"`
	BatchStats   []*BatchStats   `protobuf:"bytes,6,rep,name=batch_stats,json=batchStats" json:"batch_stats,omitempty"`
}

func (m *StatsOverall) Reset()                    { *m = StatsOverall{} }
func (m *StatsOverall) String() string            { return proto.CompactTextString(m) }
func (*StatsOverall) ProtoMessage()               {}
func (*StatsOverall) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

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

func init() {
	proto.RegisterType((*StatsCPUTime)(nil), "qpov.StatsCPUTime")
	proto.RegisterType((*MachineStats)(nil), "qpov.MachineStats")
	proto.RegisterType((*BatchStats)(nil), "qpov.BatchStats")
	proto.RegisterType((*StatsOverall)(nil), "qpov.StatsOverall")
}

func init() { proto.RegisterFile("stats.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x9a, 0x26, 0xd9, 0x9d, 0xa4, 0x65, 0x65, 0x21, 0x11, 0x38, 0xb1, 0x41, 0x08, 0x2e,
	0x54, 0x02, 0x84, 0xb8, 0xd3, 0x13, 0x07, 0x04, 0x72, 0x97, 0x73, 0xe4, 0x24, 0x96, 0xb6, 0x28,
	0x8e, 0x43, 0xec, 0xac, 0xb4, 0x47, 0x7e, 0x08, 0xf1, 0x0d, 0x7c, 0x19, 0xf6, 0xd8, 0x49, 0x5b,
	0xa4, 0x55, 0x6f, 0x33, 0x2f, 0xef, 0x8d, 0x9f, 0xdf, 0x38, 0x90, 0x2a, 0xcd, 0xb4, 0xda, 0xf4,
	0x83, 0xd4, 0x92, 0x2c, 0x7f, 0xf6, 0xf2, 0xee, 0xd9, 0x5a, 0x70, 0xcd, 0x1a, 0xa6, 0x99, 0x43,
	0x8b, 0x5f, 0x01, 0x64, 0x3b, 0xcb, 0xda, 0x7e, 0xfb, 0x7e, 0xb3, 0x17, 0x9c, 0x5c, 0x43, 0x36,
	0x2a, 0x3e, 0x94, 0x8a, 0xd7, 0xb2, 0x6b, 0x54, 0xbe, 0x78, 0x1e, 0xbc, 0x0e, 0x69, 0x6a, 0xb1,
	0x9d, 0x83, 0xc8, 0x4b, 0x58, 0xab, 0x7b, 0xa5, 0xb9, 0x98, 0x49, 0x21, 0x92, 0x56, 0x0e, 0x9d,
	0x68, 0xaf, 0xe0, 0x51, 0x2d, 0x45, 0x3f, 0x6a, 0x3e, 0xf3, 0x96, 0xc8, 0x5b, 0x7b, 0xd8, 0x13,
	0x8b, 0xbf, 0xc6, 0xc3, 0x17, 0x56, 0xdf, 0xee, 0x3b, 0x8e, 0x56, 0xac, 0x07, 0x36, 0xd4, 0xb7,
	0xa5, 0x1a, 0x85, 0x60, 0xc3, 0x7d, 0x1e, 0x18, 0xd9, 0x25, 0x4d, 0x2d, 0xb6, 0x73, 0x10, 0xb9,
	0x82, 0xb0, 0xee, 0x47, 0x74, 0x77, 0x49, 0x6d, 0x49, 0x9e, 0x40, 0xd2, 0x8d, 0xa2, 0xb4, 0xa8,
	0xb5, 0x13, 0xd1, 0xd8, 0xb4, 0x5b, 0xf3, 0xe1, 0x1a, 0xa2, 0xba, 0x95, 0x63, 0x83, 0xa7, 0xa7,
	0xef, 0xd2, 0x8d, 0x0d, 0x62, 0xb3, 0xb5, 0x10, 0x75, 0x5f, 0xc8, 0x1b, 0xb8, 0x30, 0xba, 0x52,
	0x9b, 0x00, 0xf2, 0x08, 0x59, 0xc4, 0xb1, 0x8e, 0xa3, 0xa1, 0x89, 0xe1, 0x60, 0x46, 0x04, 0x96,
	0x3f, 0x64, 0xa5, 0xf2, 0x18, 0xaf, 0x83, 0x75, 0xf1, 0x27, 0x00, 0xf8, 0xc4, 0xb4, 0x71, 0x88,
	0x57, 0x78, 0x0a, 0x17, 0x95, 0xed, 0xca, 0x7d, 0xe3, 0xed, 0x27, 0xd8, 0x7f, 0x6e, 0xc8, 0x63,
	0x88, 0xb4, 0xd4, 0xac, 0xf5, 0xd1, 0xba, 0xc6, 0xce, 0x6c, 0x64, 0xc7, 0x7d, 0x94, 0x58, 0x93,
	0x1c, 0x12, 0x13, 0x95, 0xe0, 0x9d, 0x46, 0xef, 0x66, 0x86, 0x6f, 0xed, 0x8c, 0x7a, 0x76, 0x6b,
	0x66, 0x60, 0x73, 0x72, 0x8d, 0xf8, 0xec, 0x35, 0x8a, 0xdf, 0x0b, 0xbf, 0xfb, 0xaf, 0x77, 0x7c,
	0x60, 0x6d, 0x6b, 0x37, 0x86, 0x2f, 0x06, 0x27, 0x98, 0x4a, 0xf4, 0xe8, 0xdd, 0x6c, 0x0c, 0xe1,
	0x9b, 0x09, 0x3d, 0x39, 0x68, 0x71, 0x3e, 0xaf, 0x0f, 0x90, 0x09, 0xb7, 0x5f, 0x27, 0x09, 0x1f,
	0x94, 0xa4, 0x9e, 0x87, 0xb2, 0x17, 0xb0, 0x6a, 0x39, 0x53, 0xff, 0x3f, 0x9f, 0x0c, 0xc1, 0xe9,
	0x95, 0x7d, 0x84, 0xd5, 0x34, 0x1b, 0x4d, 0x9a, 0x44, 0xc2, 0xc3, 0xf0, 0xe3, 0x67, 0x45, 0x27,
	0x13, 0x6e, 0x43, 0x6f, 0x21, 0x75, 0x1b, 0x72, 0xb2, 0x18, 0x65, 0x57, 0x4e, 0x76, 0x58, 0x24,
	0x85, 0x6a, 0xae, 0xab, 0x18, 0xff, 0x99, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0xe6,
	0xb7, 0x46, 0x58, 0x03, 0x00, 0x00,
}
