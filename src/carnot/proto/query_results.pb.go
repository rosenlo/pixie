// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/carnot/proto/query_results.proto

package carnotpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "pixielabs.ai/pixielabs/src/carnot/schema/proto"
import _ "pixielabs.ai/pixielabs/src/shared/types/proto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type QueryTimingInfo struct {
	ExecutionTimeNs   int64 `protobuf:"varint,1,opt,name=execution_time_ns,json=executionTimeNs,proto3" json:"execution_time_ns,omitempty"`
	CompilationTimeNs int64 `protobuf:"varint,2,opt,name=compilation_time_ns,json=compilationTimeNs,proto3" json:"compilation_time_ns,omitempty"`
}

func (m *QueryTimingInfo) Reset()      { *m = QueryTimingInfo{} }
func (*QueryTimingInfo) ProtoMessage() {}
func (*QueryTimingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_results_9fb97f59eab7c5ae, []int{0}
}
func (m *QueryTimingInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTimingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTimingInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *QueryTimingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTimingInfo.Merge(dst, src)
}
func (m *QueryTimingInfo) XXX_Size() int {
	return m.Size()
}
func (m *QueryTimingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTimingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTimingInfo proto.InternalMessageInfo

func (m *QueryTimingInfo) GetExecutionTimeNs() int64 {
	if m != nil {
		return m.ExecutionTimeNs
	}
	return 0
}

func (m *QueryTimingInfo) GetCompilationTimeNs() int64 {
	if m != nil {
		return m.CompilationTimeNs
	}
	return 0
}

type QueryExecutionStats struct {
	Timing           *QueryTimingInfo `protobuf:"bytes,1,opt,name=timing,proto3" json:"timing,omitempty"`
	BytesProcessed   int64            `protobuf:"varint,2,opt,name=bytes_processed,json=bytesProcessed,proto3" json:"bytes_processed,omitempty"`
	RecordsProcessed int64            `protobuf:"varint,3,opt,name=records_processed,json=recordsProcessed,proto3" json:"records_processed,omitempty"`
}

func (m *QueryExecutionStats) Reset()      { *m = QueryExecutionStats{} }
func (*QueryExecutionStats) ProtoMessage() {}
func (*QueryExecutionStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_results_9fb97f59eab7c5ae, []int{1}
}
func (m *QueryExecutionStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryExecutionStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryExecutionStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *QueryExecutionStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryExecutionStats.Merge(dst, src)
}
func (m *QueryExecutionStats) XXX_Size() int {
	return m.Size()
}
func (m *QueryExecutionStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryExecutionStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueryExecutionStats proto.InternalMessageInfo

func (m *QueryExecutionStats) GetTiming() *QueryTimingInfo {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *QueryExecutionStats) GetBytesProcessed() int64 {
	if m != nil {
		return m.BytesProcessed
	}
	return 0
}

func (m *QueryExecutionStats) GetRecordsProcessed() int64 {
	if m != nil {
		return m.RecordsProcessed
	}
	return 0
}

type QueryResult struct {
	Tables         []*proto1.Table      `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
	TimingInfo     *QueryTimingInfo     `protobuf:"bytes,2,opt,name=timing_info,json=timingInfo,proto3" json:"timing_info,omitempty"`
	ExecutionStats *QueryExecutionStats `protobuf:"bytes,3,opt,name=execution_stats,json=executionStats,proto3" json:"execution_stats,omitempty"`
}

func (m *QueryResult) Reset()      { *m = QueryResult{} }
func (*QueryResult) ProtoMessage() {}
func (*QueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_results_9fb97f59eab7c5ae, []int{2}
}
func (m *QueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *QueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResult.Merge(dst, src)
}
func (m *QueryResult) XXX_Size() int {
	return m.Size()
}
func (m *QueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResult proto.InternalMessageInfo

func (m *QueryResult) GetTables() []*proto1.Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

func (m *QueryResult) GetTimingInfo() *QueryTimingInfo {
	if m != nil {
		return m.TimingInfo
	}
	return nil
}

func (m *QueryResult) GetExecutionStats() *QueryExecutionStats {
	if m != nil {
		return m.ExecutionStats
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryTimingInfo)(nil), "pl.carnot.carnotpb.QueryTimingInfo")
	proto.RegisterType((*QueryExecutionStats)(nil), "pl.carnot.carnotpb.QueryExecutionStats")
	proto.RegisterType((*QueryResult)(nil), "pl.carnot.carnotpb.QueryResult")
}
func (this *QueryTimingInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryTimingInfo)
	if !ok {
		that2, ok := that.(QueryTimingInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ExecutionTimeNs != that1.ExecutionTimeNs {
		return false
	}
	if this.CompilationTimeNs != that1.CompilationTimeNs {
		return false
	}
	return true
}
func (this *QueryExecutionStats) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryExecutionStats)
	if !ok {
		that2, ok := that.(QueryExecutionStats)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Timing.Equal(that1.Timing) {
		return false
	}
	if this.BytesProcessed != that1.BytesProcessed {
		return false
	}
	if this.RecordsProcessed != that1.RecordsProcessed {
		return false
	}
	return true
}
func (this *QueryResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryResult)
	if !ok {
		that2, ok := that.(QueryResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Tables) != len(that1.Tables) {
		return false
	}
	for i := range this.Tables {
		if !this.Tables[i].Equal(that1.Tables[i]) {
			return false
		}
	}
	if !this.TimingInfo.Equal(that1.TimingInfo) {
		return false
	}
	if !this.ExecutionStats.Equal(that1.ExecutionStats) {
		return false
	}
	return true
}
func (this *QueryTimingInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&carnotpb.QueryTimingInfo{")
	s = append(s, "ExecutionTimeNs: "+fmt.Sprintf("%#v", this.ExecutionTimeNs)+",\n")
	s = append(s, "CompilationTimeNs: "+fmt.Sprintf("%#v", this.CompilationTimeNs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryExecutionStats) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&carnotpb.QueryExecutionStats{")
	if this.Timing != nil {
		s = append(s, "Timing: "+fmt.Sprintf("%#v", this.Timing)+",\n")
	}
	s = append(s, "BytesProcessed: "+fmt.Sprintf("%#v", this.BytesProcessed)+",\n")
	s = append(s, "RecordsProcessed: "+fmt.Sprintf("%#v", this.RecordsProcessed)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&carnotpb.QueryResult{")
	if this.Tables != nil {
		s = append(s, "Tables: "+fmt.Sprintf("%#v", this.Tables)+",\n")
	}
	if this.TimingInfo != nil {
		s = append(s, "TimingInfo: "+fmt.Sprintf("%#v", this.TimingInfo)+",\n")
	}
	if this.ExecutionStats != nil {
		s = append(s, "ExecutionStats: "+fmt.Sprintf("%#v", this.ExecutionStats)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQueryResults(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *QueryTimingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTimingInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ExecutionTimeNs != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintQueryResults(dAtA, i, uint64(m.ExecutionTimeNs))
	}
	if m.CompilationTimeNs != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintQueryResults(dAtA, i, uint64(m.CompilationTimeNs))
	}
	return i, nil
}

func (m *QueryExecutionStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryExecutionStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timing != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQueryResults(dAtA, i, uint64(m.Timing.Size()))
		n1, err := m.Timing.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.BytesProcessed != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintQueryResults(dAtA, i, uint64(m.BytesProcessed))
	}
	if m.RecordsProcessed != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintQueryResults(dAtA, i, uint64(m.RecordsProcessed))
	}
	return i, nil
}

func (m *QueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tables) > 0 {
		for _, msg := range m.Tables {
			dAtA[i] = 0xa
			i++
			i = encodeVarintQueryResults(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.TimingInfo != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintQueryResults(dAtA, i, uint64(m.TimingInfo.Size()))
		n2, err := m.TimingInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ExecutionStats != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintQueryResults(dAtA, i, uint64(m.ExecutionStats.Size()))
		n3, err := m.ExecutionStats.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintQueryResults(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *QueryTimingInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecutionTimeNs != 0 {
		n += 1 + sovQueryResults(uint64(m.ExecutionTimeNs))
	}
	if m.CompilationTimeNs != 0 {
		n += 1 + sovQueryResults(uint64(m.CompilationTimeNs))
	}
	return n
}

func (m *QueryExecutionStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timing != nil {
		l = m.Timing.Size()
		n += 1 + l + sovQueryResults(uint64(l))
	}
	if m.BytesProcessed != 0 {
		n += 1 + sovQueryResults(uint64(m.BytesProcessed))
	}
	if m.RecordsProcessed != 0 {
		n += 1 + sovQueryResults(uint64(m.RecordsProcessed))
	}
	return n
}

func (m *QueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tables) > 0 {
		for _, e := range m.Tables {
			l = e.Size()
			n += 1 + l + sovQueryResults(uint64(l))
		}
	}
	if m.TimingInfo != nil {
		l = m.TimingInfo.Size()
		n += 1 + l + sovQueryResults(uint64(l))
	}
	if m.ExecutionStats != nil {
		l = m.ExecutionStats.Size()
		n += 1 + l + sovQueryResults(uint64(l))
	}
	return n
}

func sovQueryResults(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQueryResults(x uint64) (n int) {
	return sovQueryResults(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QueryTimingInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryTimingInfo{`,
		`ExecutionTimeNs:` + fmt.Sprintf("%v", this.ExecutionTimeNs) + `,`,
		`CompilationTimeNs:` + fmt.Sprintf("%v", this.CompilationTimeNs) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryExecutionStats) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryExecutionStats{`,
		`Timing:` + strings.Replace(fmt.Sprintf("%v", this.Timing), "QueryTimingInfo", "QueryTimingInfo", 1) + `,`,
		`BytesProcessed:` + fmt.Sprintf("%v", this.BytesProcessed) + `,`,
		`RecordsProcessed:` + fmt.Sprintf("%v", this.RecordsProcessed) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryResult{`,
		`Tables:` + strings.Replace(fmt.Sprintf("%v", this.Tables), "Table", "proto1.Table", 1) + `,`,
		`TimingInfo:` + strings.Replace(fmt.Sprintf("%v", this.TimingInfo), "QueryTimingInfo", "QueryTimingInfo", 1) + `,`,
		`ExecutionStats:` + strings.Replace(fmt.Sprintf("%v", this.ExecutionStats), "QueryExecutionStats", "QueryExecutionStats", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQueryResults(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QueryTimingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryResults
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTimingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTimingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionTimeNs", wireType)
			}
			m.ExecutionTimeNs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecutionTimeNs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompilationTimeNs", wireType)
			}
			m.CompilationTimeNs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompilationTimeNs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryResults(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryResults
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryExecutionStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryResults
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryExecutionStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryExecutionStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryResults
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timing == nil {
				m.Timing = &QueryTimingInfo{}
			}
			if err := m.Timing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesProcessed", wireType)
			}
			m.BytesProcessed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BytesProcessed |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordsProcessed", wireType)
			}
			m.RecordsProcessed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecordsProcessed |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryResults(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryResults
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryResults
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryResults
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tables = append(m.Tables, &proto1.Table{})
			if err := m.Tables[len(m.Tables)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimingInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryResults
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TimingInfo == nil {
				m.TimingInfo = &QueryTimingInfo{}
			}
			if err := m.TimingInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryResults
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExecutionStats == nil {
				m.ExecutionStats = &QueryExecutionStats{}
			}
			if err := m.ExecutionStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryResults(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryResults
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQueryResults(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryResults
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryResults
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthQueryResults
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQueryResults
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipQueryResults(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthQueryResults = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryResults   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("src/carnot/proto/query_results.proto", fileDescriptor_query_results_9fb97f59eab7c5ae)
}

var fileDescriptor_query_results_9fb97f59eab7c5ae = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x4f, 0xfa, 0x40,
	0x18, 0xc6, 0x7b, 0x7f, 0x12, 0xf2, 0xcf, 0x35, 0x01, 0x39, 0x16, 0x65, 0xb8, 0x98, 0x6a, 0x02,
	0xd1, 0xa4, 0x8d, 0x38, 0x9a, 0x38, 0x18, 0x1d, 0x5c, 0x0c, 0x56, 0x26, 0x97, 0xa6, 0x2d, 0x2f,
	0xd2, 0xa4, 0xed, 0xd5, 0xbb, 0x23, 0x91, 0xcd, 0x8f, 0xe0, 0xc7, 0xe0, 0xa3, 0x38, 0x32, 0x32,
	0x4a, 0x59, 0x1c, 0xf9, 0x08, 0xa6, 0xd7, 0x02, 0x45, 0x63, 0xe2, 0x04, 0xf7, 0xbe, 0xbf, 0xf7,
	0x7d, 0x9e, 0x7b, 0xae, 0xf8, 0x58, 0x70, 0xdf, 0xf2, 0x5d, 0x1e, 0x33, 0x69, 0x25, 0x9c, 0x49,
	0x66, 0x3d, 0x8f, 0x81, 0x4f, 0x1c, 0x0e, 0x62, 0x1c, 0x4a, 0x61, 0xaa, 0x1a, 0x21, 0x49, 0x68,
	0xe6, 0x50, 0xf1, 0x93, 0x78, 0x2d, 0x23, 0x9b, 0x14, 0x23, 0x97, 0xc3, 0xc0, 0x92, 0x93, 0x04,
	0x44, 0x31, 0xaf, 0xfe, 0xe7, 0x73, 0xad, 0xf2, 0x76, 0xe1, 0x8f, 0x20, 0x72, 0x0b, 0x28, 0x3f,
	0xe4, 0x94, 0x11, 0xe1, 0xfa, 0x7d, 0x26, 0xda, 0x0f, 0xa2, 0x20, 0x7e, 0xba, 0x8d, 0x87, 0x8c,
	0x9c, 0xe0, 0x06, 0xbc, 0x80, 0x3f, 0x96, 0x01, 0x8b, 0x1d, 0x19, 0x44, 0xe0, 0xc4, 0x62, 0x1f,
	0x1d, 0xa2, 0x4e, 0xc5, 0xae, 0x6f, 0x1a, 0xfd, 0x20, 0x82, 0x3b, 0x41, 0x4c, 0xdc, 0xf4, 0x59,
	0x94, 0x04, 0xa1, 0xbb, 0x43, 0xff, 0x53, 0x74, 0xa3, 0xd4, 0xca, 0x79, 0x63, 0x8a, 0x70, 0x53,
	0xe9, 0xdd, 0xac, 0x17, 0x3d, 0x48, 0x57, 0x0a, 0x72, 0x81, 0xab, 0x52, 0x39, 0x50, 0x42, 0x7a,
	0xf7, 0xc8, 0xfc, 0x79, 0x6b, 0xf3, 0x9b, 0x51, 0xbb, 0x18, 0x21, 0x6d, 0x5c, 0xf7, 0x26, 0x12,
	0x84, 0x93, 0x70, 0xe6, 0x83, 0x10, 0x30, 0x28, 0x0c, 0xd4, 0x54, 0xb9, 0xb7, 0xae, 0x92, 0x53,
	0xdc, 0xe0, 0xe0, 0x33, 0x3e, 0x28, 0xa3, 0x15, 0x85, 0xee, 0x15, 0x8d, 0x0d, 0x6c, 0xcc, 0x11,
	0xd6, 0x95, 0xa2, 0xad, 0x9e, 0x83, 0x9c, 0xe1, 0xaa, 0x74, 0xbd, 0x10, 0xb2, 0x2c, 0x2a, 0x1d,
	0xbd, 0x7b, 0x50, 0xb2, 0x98, 0x47, 0x9a, 0x78, 0x66, 0x3f, 0x23, 0xec, 0x02, 0x24, 0xd7, 0x58,
	0xcf, 0x2d, 0x3a, 0x41, 0x3c, 0x64, 0xca, 0xd4, 0x1f, 0xaf, 0x86, 0xe5, 0xf6, 0x3d, 0x7a, 0x78,
	0x1b, 0xbb, 0x23, 0xb2, 0xb8, 0x94, 0x67, 0xbd, 0xdb, 0xfe, 0x75, 0xd3, 0x6e, 0xba, 0x76, 0x0d,
	0x76, 0xce, 0x57, 0x97, 0xb3, 0x05, 0xd5, 0xe6, 0x0b, 0xaa, 0xad, 0x16, 0x14, 0xbd, 0xa6, 0x14,
	0x4d, 0x53, 0x8a, 0xde, 0x53, 0x8a, 0x66, 0x29, 0x45, 0x1f, 0x29, 0x45, 0x9f, 0x29, 0xd5, 0x56,
	0x29, 0x45, 0x6f, 0x4b, 0xaa, 0xcd, 0x96, 0x54, 0x9b, 0x2f, 0xa9, 0xf6, 0xf8, 0x7f, 0xad, 0xe1,
	0x55, 0xd5, 0xb7, 0x73, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x12, 0x67, 0x77, 0xc1, 0x02,
	0x00, 0x00,
}
