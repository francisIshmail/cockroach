// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/log.proto

package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import time "time"
import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type RangeLogEventType int32

const (
	// These are lower case to maintain compatibility with how they were
	// originally stored.
	// Split is the event type recorded when a range splits.
	RangeLogEventType_split RangeLogEventType = 0
	// Add is the event type recorded when a range adds a new replica.
	RangeLogEventType_add RangeLogEventType = 1
	// Remove is the event type recorded when a range removed an existing replica.
	RangeLogEventType_remove RangeLogEventType = 2
)

var RangeLogEventType_name = map[int32]string{
	0: "split",
	1: "add",
	2: "remove",
}
var RangeLogEventType_value = map[string]int32{
	"split":  0,
	"add":    1,
	"remove": 2,
}

func (x RangeLogEventType) String() string {
	return proto.EnumName(RangeLogEventType_name, int32(x))
}
func (RangeLogEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorLog, []int{0} }

type RangeLogEvent struct {
	Timestamp    time.Time                                            `protobuf:"bytes,1,opt,name=timestamp,stdtime" json:"timestamp"`
	RangeID      github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,2,opt,name=range_id,json=rangeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"range_id,omitempty"`
	StoreID      github_com_cockroachdb_cockroach_pkg_roachpb.StoreID `protobuf:"varint,3,opt,name=store_id,json=storeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.StoreID" json:"store_id,omitempty"`
	EventType    RangeLogEventType                                    `protobuf:"varint,4,opt,name=event_type,json=eventType,proto3,enum=cockroach.storage.RangeLogEventType" json:"event_type,omitempty"`
	OtherRangeID github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,5,opt,name=other_range_id,json=otherRangeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"other_range_id,omitempty"`
	Info         *RangeLogEvent_Info                                  `protobuf:"bytes,6,opt,name=info" json:"info,omitempty"`
}

func (m *RangeLogEvent) Reset()                    { *m = RangeLogEvent{} }
func (m *RangeLogEvent) String() string            { return proto.CompactTextString(m) }
func (*RangeLogEvent) ProtoMessage()               {}
func (*RangeLogEvent) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{0} }

type RangeLogEvent_Info struct {
	UpdatedDesc    *cockroach_roachpb.RangeDescriptor   `protobuf:"bytes,1,opt,name=updated_desc,json=updatedDesc" json:"UpdatedDesc"`
	NewDesc        *cockroach_roachpb.RangeDescriptor   `protobuf:"bytes,2,opt,name=new_desc,json=newDesc" json:"NewDesc"`
	AddedReplica   *cockroach_roachpb.ReplicaDescriptor `protobuf:"bytes,3,opt,name=added_replica,json=addedReplica" json:"AddReplica"`
	RemovedReplica *cockroach_roachpb.ReplicaDescriptor `protobuf:"bytes,4,opt,name=removed_replica,json=removedReplica" json:"RemovedReplica"`
	Reason         RangeLogEventReason                  `protobuf:"bytes,5,opt,name=reason,proto3,casttype=RangeLogEventReason" json:"Reason"`
	Details        string                               `protobuf:"bytes,6,opt,name=details,proto3" json:"Details"`
}

func (m *RangeLogEvent_Info) Reset()                    { *m = RangeLogEvent_Info{} }
func (m *RangeLogEvent_Info) String() string            { return proto.CompactTextString(m) }
func (*RangeLogEvent_Info) ProtoMessage()               {}
func (*RangeLogEvent_Info) Descriptor() ([]byte, []int) { return fileDescriptorLog, []int{0, 0} }

func init() {
	proto.RegisterType((*RangeLogEvent)(nil), "cockroach.storage.RangeLogEvent")
	proto.RegisterType((*RangeLogEvent_Info)(nil), "cockroach.storage.RangeLogEvent.Info")
	proto.RegisterEnum("cockroach.storage.RangeLogEventType", RangeLogEventType_name, RangeLogEventType_value)
}
func (m *RangeLogEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeLogEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLog(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.RangeID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.RangeID))
	}
	if m.StoreID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.StoreID))
	}
	if m.EventType != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.EventType))
	}
	if m.OtherRangeID != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.OtherRangeID))
	}
	if m.Info != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Info.Size()))
		n2, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *RangeLogEvent_Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeLogEvent_Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UpdatedDesc != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.UpdatedDesc.Size()))
		n3, err := m.UpdatedDesc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.NewDesc != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.NewDesc.Size()))
		n4, err := m.NewDesc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.AddedReplica != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.AddedReplica.Size()))
		n5, err := m.AddedReplica.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.RemovedReplica != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.RemovedReplica.Size()))
		n6, err := m.RemovedReplica.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if len(m.Details) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Details)))
		i += copy(dAtA[i:], m.Details)
	}
	return i, nil
}

func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RangeLogEvent) Size() (n int) {
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovLog(uint64(l))
	if m.RangeID != 0 {
		n += 1 + sovLog(uint64(m.RangeID))
	}
	if m.StoreID != 0 {
		n += 1 + sovLog(uint64(m.StoreID))
	}
	if m.EventType != 0 {
		n += 1 + sovLog(uint64(m.EventType))
	}
	if m.OtherRangeID != 0 {
		n += 1 + sovLog(uint64(m.OtherRangeID))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	return n
}

func (m *RangeLogEvent_Info) Size() (n int) {
	var l int
	_ = l
	if m.UpdatedDesc != nil {
		l = m.UpdatedDesc.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.NewDesc != nil {
		l = m.NewDesc.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.AddedReplica != nil {
		l = m.AddedReplica.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.RemovedReplica != nil {
		l = m.RemovedReplica.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	return n
}

func sovLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RangeLogEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: RangeLogEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeLogEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			m.StoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreID |= (github_com_cockroachdb_cockroach_pkg_roachpb.StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			m.EventType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventType |= (RangeLogEventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherRangeID", wireType)
			}
			m.OtherRangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherRangeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &RangeLogEvent_Info{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
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
func (m *RangeLogEvent_Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
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
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedDesc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedDesc == nil {
				m.UpdatedDesc = &cockroach_roachpb.RangeDescriptor{}
			}
			if err := m.UpdatedDesc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewDesc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NewDesc == nil {
				m.NewDesc = &cockroach_roachpb.RangeDescriptor{}
			}
			if err := m.NewDesc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AddedReplica == nil {
				m.AddedReplica = &cockroach_roachpb.ReplicaDescriptor{}
			}
			if err := m.AddedReplica.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovedReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
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
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RemovedReplica == nil {
				m.RemovedReplica = &cockroach_roachpb.ReplicaDescriptor{}
			}
			if err := m.RemovedReplica.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = RangeLogEventReason(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
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
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
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
					return 0, ErrIntOverflowLog
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
					return 0, ErrIntOverflowLog
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
				return 0, ErrInvalidLengthLog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLog
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
				next, err := skipLog(dAtA[start:])
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
	ErrInvalidLengthLog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage/log.proto", fileDescriptorLog) }

var fileDescriptorLog = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xe3, 0xdc, 0x9c, 0x4c, 0xd2, 0x34, 0x19, 0x10, 0x8a, 0xb2, 0xf0, 0xa4, 0x15, 0x95,
	0x22, 0x16, 0xb6, 0xd4, 0xb2, 0xe9, 0x12, 0x37, 0x48, 0x04, 0x21, 0x90, 0x86, 0x82, 0x04, 0x0b,
	0xa2, 0x89, 0x67, 0xe2, 0x58, 0x4d, 0x3c, 0x96, 0x3d, 0x69, 0xd5, 0xb7, 0xe8, 0x63, 0x45, 0xac,
	0xba, 0x64, 0x65, 0xc0, 0xec, 0xf2, 0x08, 0x5d, 0x21, 0xcf, 0x38, 0x37, 0x75, 0xc1, 0xa5, 0xbb,
	0x73, 0xfd, 0x7c, 0xe6, 0x3f, 0xc7, 0xa0, 0x15, 0x09, 0x1e, 0x12, 0x97, 0x59, 0x53, 0xee, 0x9a,
	0x41, 0xc8, 0x05, 0x87, 0x2d, 0x87, 0x3b, 0x17, 0x21, 0x27, 0xce, 0xc4, 0xcc, 0x92, 0x9d, 0x27,
	0xd2, 0x0d, 0x46, 0xd6, 0x8c, 0x09, 0x42, 0x89, 0x20, 0xaa, 0xb4, 0xf3, 0xd8, 0xe5, 0x2e, 0x97,
	0xa6, 0x95, 0x5a, 0x59, 0x14, 0xb9, 0x9c, 0xbb, 0x53, 0x66, 0x49, 0x6f, 0x34, 0x1f, 0x5b, 0xc2,
	0x9b, 0xb1, 0x48, 0x90, 0x59, 0xa0, 0x0a, 0x0e, 0x6f, 0x75, 0xb0, 0x87, 0x89, 0xef, 0xb2, 0x37,
	0xdc, 0x7d, 0x79, 0xc9, 0x7c, 0x01, 0x6d, 0x50, 0x5d, 0x17, 0xb5, 0xb5, 0xae, 0xd6, 0xab, 0x1d,
	0x77, 0x4c, 0x85, 0x31, 0x57, 0x18, 0xf3, 0x7c, 0x55, 0x61, 0x57, 0x16, 0x31, 0xca, 0xdd, 0x7c,
	0x47, 0x1a, 0xde, 0xb4, 0xc1, 0x2f, 0xa0, 0x12, 0xa6, 0xd0, 0xa1, 0x47, 0xdb, 0xf9, 0xae, 0xd6,
	0x2b, 0xd8, 0x67, 0x49, 0x8c, 0x74, 0xf9, 0xa1, 0x41, 0xff, 0x2e, 0x46, 0xcf, 0x5d, 0x4f, 0x4c,
	0xe6, 0x23, 0xd3, 0xe1, 0x33, 0x6b, 0xfd, 0x46, 0x3a, 0xda, 0xd8, 0x56, 0x70, 0xe1, 0x5a, 0xd9,
	0x53, 0xcd, 0xac, 0x0f, 0xeb, 0x12, 0x3a, 0xa0, 0x29, 0x3f, 0xd5, 0x43, 0xf2, 0x0b, 0x5d, 0xad,
	0x57, 0x52, 0xfc, 0xf7, 0x69, 0xec, 0x3f, 0xf8, 0x59, 0x1f, 0xd6, 0x25, 0x74, 0x40, 0xe1, 0x19,
	0x00, 0x2c, 0x15, 0x63, 0x28, 0xae, 0x03, 0xd6, 0x2e, 0x76, 0xb5, 0x5e, 0xe3, 0xf8, 0xa9, 0x79,
	0x6f, 0x19, 0xe6, 0x8e, 0x72, 0xe7, 0xd7, 0x01, 0xc3, 0x55, 0xb6, 0x32, 0xa1, 0x0f, 0x1a, 0x5c,
	0x4c, 0x58, 0x38, 0x5c, 0x4b, 0x51, 0x92, 0x52, 0xbc, 0x4a, 0x62, 0x54, 0x7f, 0x97, 0x66, 0x1e,
	0xaa, 0x47, 0x9d, 0x6f, 0x28, 0x14, 0x9e, 0x82, 0xa2, 0xe7, 0x8f, 0x79, 0xbb, 0x2c, 0x77, 0x76,
	0xf4, 0xa7, 0x71, 0xcd, 0x81, 0x3f, 0xe6, 0x58, 0xb6, 0x74, 0xbe, 0x16, 0x40, 0x31, 0x75, 0xe1,
	0x47, 0x50, 0x9f, 0x07, 0x94, 0x08, 0x46, 0x87, 0x94, 0x45, 0x4e, 0xb6, 0xff, 0xc3, 0x2d, 0xd6,
	0xce, 0x0c, 0x7d, 0x16, 0x39, 0xa1, 0x17, 0x08, 0x1e, 0xda, 0xfb, 0xcb, 0x18, 0xd5, 0x3e, 0xa8,
	0xde, 0x34, 0x8c, 0x6b, 0xf3, 0x8d, 0x03, 0x5f, 0x83, 0x8a, 0xcf, 0xae, 0x14, 0x33, 0xff, 0xd7,
	0xcc, 0xda, 0x32, 0x46, 0xfa, 0x5b, 0x76, 0x25, 0x79, 0xba, 0xaf, 0x0c, 0xf8, 0x09, 0xec, 0x11,
	0x4a, 0x19, 0x1d, 0x86, 0x2c, 0x98, 0x7a, 0x0e, 0x91, 0x17, 0x50, 0xdb, 0xd9, 0xcf, 0x1a, 0xa8,
	0x2a, 0xb6, 0x90, 0x8d, 0x65, 0x8c, 0xc0, 0x0b, 0x4a, 0xb3, 0x0c, 0xae, 0x4b, 0x54, 0xe6, 0x41,
	0x02, 0xf6, 0x43, 0x36, 0xe3, 0x97, 0x5b, 0xf0, 0xe2, 0x3f, 0xc0, 0xe1, 0x32, 0x46, 0x0d, 0xac,
	0x00, 0xab, 0x0f, 0x34, 0xc2, 0x1d, 0x1f, 0x9e, 0x82, 0x72, 0xc8, 0x48, 0xc4, 0x7d, 0x79, 0x0d,
	0x55, 0xfb, 0x60, 0x19, 0xa3, 0x32, 0x96, 0x91, 0xbb, 0x18, 0x3d, 0xda, 0x59, 0x91, 0x0a, 0xe3,
	0xac, 0x01, 0x1e, 0x01, 0x9d, 0x32, 0x41, 0xbc, 0x69, 0x24, 0x77, 0x5c, 0x55, 0xfa, 0xf4, 0x55,
	0x08, 0xaf, 0x72, 0xcf, 0x4e, 0x40, 0xeb, 0xde, 0x5d, 0xc2, 0x2a, 0x28, 0x45, 0xc1, 0xd4, 0x13,
	0xcd, 0x1c, 0xd4, 0x41, 0x81, 0x50, 0xda, 0xd4, 0x20, 0x48, 0x47, 0x49, 0x87, 0x6b, 0xe6, 0xed,
	0x83, 0xc5, 0x4f, 0x23, 0xb7, 0x48, 0x0c, 0xed, 0x36, 0x31, 0xb4, 0x6f, 0x89, 0xa1, 0xfd, 0x48,
	0x0c, 0xed, 0xe6, 0x97, 0x91, 0xfb, 0xac, 0x67, 0xd7, 0x33, 0x2a, 0xcb, 0xbf, 0xff, 0xe4, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0xfe, 0xcf, 0x8a, 0xa8, 0x04, 0x00, 0x00,
}
