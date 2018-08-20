// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/data/core/v2alpha/health_check_event.proto

package envoy_data_core_v2alpha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import bytes "bytes"

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

type HealthCheckFailureType int32

const (
	HealthCheckFailureType_ACTIVE  HealthCheckFailureType = 0
	HealthCheckFailureType_PASSIVE HealthCheckFailureType = 1
	HealthCheckFailureType_NETWORK HealthCheckFailureType = 2
)

var HealthCheckFailureType_name = map[int32]string{
	0: "ACTIVE",
	1: "PASSIVE",
	2: "NETWORK",
}
var HealthCheckFailureType_value = map[string]int32{
	"ACTIVE":  0,
	"PASSIVE": 1,
	"NETWORK": 2,
}

func (x HealthCheckFailureType) String() string {
	return proto.EnumName(HealthCheckFailureType_name, int32(x))
}
func (HealthCheckFailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_health_check_event_e1f981a911dc488a, []int{0}
}

type HealthCheckerType int32

const (
	HealthCheckerType_HTTP  HealthCheckerType = 0
	HealthCheckerType_TCP   HealthCheckerType = 1
	HealthCheckerType_GRPC  HealthCheckerType = 2
	HealthCheckerType_REDIS HealthCheckerType = 3
)

var HealthCheckerType_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "GRPC",
	3: "REDIS",
}
var HealthCheckerType_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"GRPC":  2,
	"REDIS": 3,
}

func (x HealthCheckerType) String() string {
	return proto.EnumName(HealthCheckerType_name, int32(x))
}
func (HealthCheckerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_health_check_event_e1f981a911dc488a, []int{1}
}

type HealthCheckEvent struct {
	HealthCheckerType HealthCheckerType `protobuf:"varint,1,opt,name=health_checker_type,json=healthCheckerType,proto3,enum=envoy.data.core.v2alpha.HealthCheckerType" json:"health_checker_type,omitempty"`
	Host              *core.Address     `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	ClusterName       string            `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*HealthCheckEvent_EjectUnhealthyEvent
	//	*HealthCheckEvent_AddHealthyEvent
	Event                isHealthCheckEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HealthCheckEvent) Reset()         { *m = HealthCheckEvent{} }
func (m *HealthCheckEvent) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEvent) ProtoMessage()    {}
func (*HealthCheckEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_check_event_e1f981a911dc488a, []int{0}
}
func (m *HealthCheckEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthCheckEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthCheckEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HealthCheckEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEvent.Merge(dst, src)
}
func (m *HealthCheckEvent) XXX_Size() int {
	return m.Size()
}
func (m *HealthCheckEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEvent.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEvent proto.InternalMessageInfo

type isHealthCheckEvent_Event interface {
	isHealthCheckEvent_Event()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type HealthCheckEvent_EjectUnhealthyEvent struct {
	EjectUnhealthyEvent *HealthCheckEjectUnhealthy `protobuf:"bytes,4,opt,name=eject_unhealthy_event,json=ejectUnhealthyEvent,oneof"`
}
type HealthCheckEvent_AddHealthyEvent struct {
	AddHealthyEvent *HealthCheckAddHealthy `protobuf:"bytes,5,opt,name=add_healthy_event,json=addHealthyEvent,oneof"`
}

func (*HealthCheckEvent_EjectUnhealthyEvent) isHealthCheckEvent_Event() {}
func (*HealthCheckEvent_AddHealthyEvent) isHealthCheckEvent_Event()     {}

func (m *HealthCheckEvent) GetEvent() isHealthCheckEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *HealthCheckEvent) GetHealthCheckerType() HealthCheckerType {
	if m != nil {
		return m.HealthCheckerType
	}
	return HealthCheckerType_HTTP
}

func (m *HealthCheckEvent) GetHost() *core.Address {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *HealthCheckEvent) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *HealthCheckEvent) GetEjectUnhealthyEvent() *HealthCheckEjectUnhealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_EjectUnhealthyEvent); ok {
		return x.EjectUnhealthyEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetAddHealthyEvent() *HealthCheckAddHealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_AddHealthyEvent); ok {
		return x.AddHealthyEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HealthCheckEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HealthCheckEvent_OneofMarshaler, _HealthCheckEvent_OneofUnmarshaler, _HealthCheckEvent_OneofSizer, []interface{}{
		(*HealthCheckEvent_EjectUnhealthyEvent)(nil),
		(*HealthCheckEvent_AddHealthyEvent)(nil),
	}
}

func _HealthCheckEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HealthCheckEvent)
	// event
	switch x := m.Event.(type) {
	case *HealthCheckEvent_EjectUnhealthyEvent:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EjectUnhealthyEvent); err != nil {
			return err
		}
	case *HealthCheckEvent_AddHealthyEvent:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AddHealthyEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HealthCheckEvent.Event has unexpected type %T", x)
	}
	return nil
}

func _HealthCheckEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HealthCheckEvent)
	switch tag {
	case 4: // event.eject_unhealthy_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheckEjectUnhealthy)
		err := b.DecodeMessage(msg)
		m.Event = &HealthCheckEvent_EjectUnhealthyEvent{msg}
		return true, err
	case 5: // event.add_healthy_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HealthCheckAddHealthy)
		err := b.DecodeMessage(msg)
		m.Event = &HealthCheckEvent_AddHealthyEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HealthCheckEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HealthCheckEvent)
	// event
	switch x := m.Event.(type) {
	case *HealthCheckEvent_EjectUnhealthyEvent:
		s := proto.Size(x.EjectUnhealthyEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HealthCheckEvent_AddHealthyEvent:
		s := proto.Size(x.AddHealthyEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HealthCheckEjectUnhealthy struct {
	// The type of failure that caused this ejection.
	FailureType          HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v2alpha.HealthCheckFailureType" json:"failure_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheckEjectUnhealthy) Reset()         { *m = HealthCheckEjectUnhealthy{} }
func (m *HealthCheckEjectUnhealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEjectUnhealthy) ProtoMessage()    {}
func (*HealthCheckEjectUnhealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_check_event_e1f981a911dc488a, []int{1}
}
func (m *HealthCheckEjectUnhealthy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthCheckEjectUnhealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthCheckEjectUnhealthy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HealthCheckEjectUnhealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEjectUnhealthy.Merge(dst, src)
}
func (m *HealthCheckEjectUnhealthy) XXX_Size() int {
	return m.Size()
}
func (m *HealthCheckEjectUnhealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEjectUnhealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEjectUnhealthy proto.InternalMessageInfo

func (m *HealthCheckEjectUnhealthy) GetFailureType() HealthCheckFailureType {
	if m != nil {
		return m.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

type HealthCheckAddHealthy struct {
	// Whether this addition is the result of the first ever health check on a host, in which case
	// the configured :ref:`healthy threshold <envoy_api_field_core.HealthCheck.healthy_threshold>`
	// is bypassed and the host is immediately added.
	FirstCheck           bool     `protobuf:"varint,1,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckAddHealthy) Reset()         { *m = HealthCheckAddHealthy{} }
func (m *HealthCheckAddHealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckAddHealthy) ProtoMessage()    {}
func (*HealthCheckAddHealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_check_event_e1f981a911dc488a, []int{2}
}
func (m *HealthCheckAddHealthy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthCheckAddHealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthCheckAddHealthy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HealthCheckAddHealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckAddHealthy.Merge(dst, src)
}
func (m *HealthCheckAddHealthy) XXX_Size() int {
	return m.Size()
}
func (m *HealthCheckAddHealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckAddHealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckAddHealthy proto.InternalMessageInfo

func (m *HealthCheckAddHealthy) GetFirstCheck() bool {
	if m != nil {
		return m.FirstCheck
	}
	return false
}

func init() {
	proto.RegisterType((*HealthCheckEvent)(nil), "envoy.data.core.v2alpha.HealthCheckEvent")
	proto.RegisterType((*HealthCheckEjectUnhealthy)(nil), "envoy.data.core.v2alpha.HealthCheckEjectUnhealthy")
	proto.RegisterType((*HealthCheckAddHealthy)(nil), "envoy.data.core.v2alpha.HealthCheckAddHealthy")
	proto.RegisterEnum("envoy.data.core.v2alpha.HealthCheckFailureType", HealthCheckFailureType_name, HealthCheckFailureType_value)
	proto.RegisterEnum("envoy.data.core.v2alpha.HealthCheckerType", HealthCheckerType_name, HealthCheckerType_value)
}
func (this *HealthCheckEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheckEvent)
	if !ok {
		that2, ok := that.(HealthCheckEvent)
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
	if this.HealthCheckerType != that1.HealthCheckerType {
		return false
	}
	if !this.Host.Equal(that1.Host) {
		return false
	}
	if this.ClusterName != that1.ClusterName {
		return false
	}
	if that1.Event == nil {
		if this.Event != nil {
			return false
		}
	} else if this.Event == nil {
		return false
	} else if !this.Event.Equal(that1.Event) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HealthCheckEvent_EjectUnhealthyEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheckEvent_EjectUnhealthyEvent)
	if !ok {
		that2, ok := that.(HealthCheckEvent_EjectUnhealthyEvent)
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
	if !this.EjectUnhealthyEvent.Equal(that1.EjectUnhealthyEvent) {
		return false
	}
	return true
}
func (this *HealthCheckEvent_AddHealthyEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheckEvent_AddHealthyEvent)
	if !ok {
		that2, ok := that.(HealthCheckEvent_AddHealthyEvent)
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
	if !this.AddHealthyEvent.Equal(that1.AddHealthyEvent) {
		return false
	}
	return true
}
func (this *HealthCheckEjectUnhealthy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheckEjectUnhealthy)
	if !ok {
		that2, ok := that.(HealthCheckEjectUnhealthy)
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
	if this.FailureType != that1.FailureType {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HealthCheckAddHealthy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheckAddHealthy)
	if !ok {
		that2, ok := that.(HealthCheckAddHealthy)
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
	if this.FirstCheck != that1.FirstCheck {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *HealthCheckEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HealthCheckerType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHealthCheckEvent(dAtA, i, uint64(m.HealthCheckerType))
	}
	if m.Host != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHealthCheckEvent(dAtA, i, uint64(m.Host.Size()))
		n1, err := m.Host.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.ClusterName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHealthCheckEvent(dAtA, i, uint64(len(m.ClusterName)))
		i += copy(dAtA[i:], m.ClusterName)
	}
	if m.Event != nil {
		nn2, err := m.Event.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HealthCheckEvent_EjectUnhealthyEvent) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.EjectUnhealthyEvent != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHealthCheckEvent(dAtA, i, uint64(m.EjectUnhealthyEvent.Size()))
		n3, err := m.EjectUnhealthyEvent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *HealthCheckEvent_AddHealthyEvent) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.AddHealthyEvent != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHealthCheckEvent(dAtA, i, uint64(m.AddHealthyEvent.Size()))
		n4, err := m.AddHealthyEvent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *HealthCheckEjectUnhealthy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckEjectUnhealthy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FailureType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHealthCheckEvent(dAtA, i, uint64(m.FailureType))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HealthCheckAddHealthy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheckAddHealthy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FirstCheck {
		dAtA[i] = 0x8
		i++
		if m.FirstCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHealthCheckEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HealthCheckEvent) Size() (n int) {
	var l int
	_ = l
	if m.HealthCheckerType != 0 {
		n += 1 + sovHealthCheckEvent(uint64(m.HealthCheckerType))
	}
	if m.Host != nil {
		l = m.Host.Size()
		n += 1 + l + sovHealthCheckEvent(uint64(l))
	}
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovHealthCheckEvent(uint64(l))
	}
	if m.Event != nil {
		n += m.Event.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HealthCheckEvent_EjectUnhealthyEvent) Size() (n int) {
	var l int
	_ = l
	if m.EjectUnhealthyEvent != nil {
		l = m.EjectUnhealthyEvent.Size()
		n += 1 + l + sovHealthCheckEvent(uint64(l))
	}
	return n
}
func (m *HealthCheckEvent_AddHealthyEvent) Size() (n int) {
	var l int
	_ = l
	if m.AddHealthyEvent != nil {
		l = m.AddHealthyEvent.Size()
		n += 1 + l + sovHealthCheckEvent(uint64(l))
	}
	return n
}
func (m *HealthCheckEjectUnhealthy) Size() (n int) {
	var l int
	_ = l
	if m.FailureType != 0 {
		n += 1 + sovHealthCheckEvent(uint64(m.FailureType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HealthCheckAddHealthy) Size() (n int) {
	var l int
	_ = l
	if m.FirstCheck {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHealthCheckEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealthCheckEvent(x uint64) (n int) {
	return sovHealthCheckEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthCheckEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheckEvent
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
			return fmt.Errorf("proto: HealthCheckEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HealthCheckerType", wireType)
			}
			m.HealthCheckerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheckEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HealthCheckerType |= (HealthCheckerType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheckEvent
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
				return ErrInvalidLengthHealthCheckEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Host == nil {
				m.Host = &core.Address{}
			}
			if err := m.Host.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheckEvent
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
				return ErrInvalidLengthHealthCheckEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EjectUnhealthyEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheckEvent
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
				return ErrInvalidLengthHealthCheckEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HealthCheckEjectUnhealthy{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &HealthCheckEvent_EjectUnhealthyEvent{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddHealthyEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheckEvent
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
				return ErrInvalidLengthHealthCheckEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HealthCheckAddHealthy{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &HealthCheckEvent_AddHealthyEvent{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheckEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheckEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HealthCheckEjectUnhealthy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheckEvent
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
			return fmt.Errorf("proto: HealthCheckEjectUnhealthy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckEjectUnhealthy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureType", wireType)
			}
			m.FailureType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheckEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FailureType |= (HealthCheckFailureType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheckEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheckEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HealthCheckAddHealthy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheckEvent
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
			return fmt.Errorf("proto: HealthCheckAddHealthy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckAddHealthy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheckEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FirstCheck = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheckEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheckEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHealthCheckEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealthCheckEvent
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
					return 0, ErrIntOverflowHealthCheckEvent
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
					return 0, ErrIntOverflowHealthCheckEvent
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
				return 0, ErrInvalidLengthHealthCheckEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealthCheckEvent
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
				next, err := skipHealthCheckEvent(dAtA[start:])
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
	ErrInvalidLengthHealthCheckEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealthCheckEvent   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/data/core/v2alpha/health_check_event.proto", fileDescriptor_health_check_event_e1f981a911dc488a)
}

var fileDescriptor_health_check_event_e1f981a911dc488a = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0xee, 0xb0, 0xd0, 0x96, 0x47, 0x53, 0x97, 0xa9, 0xb5, 0x48, 0x0c, 0x25, 0x9c, 0x08, 0x31,
	0xbb, 0x66, 0xbd, 0x98, 0x98, 0x98, 0x00, 0xa2, 0x34, 0x26, 0x95, 0x2c, 0xab, 0x5e, 0x8c, 0x9b,
	0x81, 0x7d, 0xb0, 0xab, 0x5b, 0x76, 0x33, 0xbb, 0xac, 0x21, 0xde, 0xfc, 0x35, 0x1e, 0x3d, 0x1a,
	0x4f, 0x3d, 0x7a, 0xf4, 0x27, 0x18, 0x6e, 0xfd, 0x17, 0x66, 0x67, 0xa0, 0x52, 0x29, 0x09, 0xb7,
	0x79, 0xef, 0x9b, 0xef, 0xfb, 0xe6, 0x9b, 0xf7, 0xe0, 0x11, 0x4e, 0x92, 0x60, 0xa6, 0x3b, 0x2c,
	0x66, 0xfa, 0x30, 0xe0, 0xa8, 0x27, 0x06, 0xf3, 0x43, 0x97, 0xe9, 0x2e, 0x32, 0x3f, 0x76, 0xed,
	0xa1, 0x8b, 0xc3, 0x4f, 0x36, 0x26, 0x38, 0x89, 0xb5, 0x90, 0x07, 0x71, 0x40, 0x4f, 0x04, 0x43,
	0x4b, 0x19, 0x5a, 0xca, 0xd0, 0x16, 0x8c, 0xf2, 0xa9, 0x94, 0x62, 0xa1, 0xa7, 0x27, 0x86, 0x14,
	0x63, 0x8e, 0xc3, 0x31, 0x8a, 0x24, 0xb3, 0xfc, 0x60, 0xfd, 0xc2, 0x80, 0x45, 0xb8, 0x40, 0x2b,
	0xe3, 0x20, 0x18, 0xfb, 0xa8, 0x8b, 0x6a, 0x30, 0x1d, 0xe9, 0xce, 0x94, 0xb3, 0xd8, 0x0b, 0x26,
	0x9b, 0xf0, 0xcf, 0x9c, 0x85, 0x21, 0xf2, 0xa5, 0xfa, 0x49, 0xc2, 0x7c, 0xcf, 0x61, 0x31, 0xea,
	0xcb, 0xc3, 0x02, 0xb8, 0x3b, 0x0e, 0xc6, 0x81, 0x38, 0xea, 0xe9, 0x49, 0x76, 0x6b, 0xdf, 0x15,
	0x50, 0xbb, 0x22, 0x63, 0x3b, 0x8d, 0xd8, 0x49, 0x13, 0xd2, 0x11, 0x1c, 0xad, 0xe6, 0x46, 0x6e,
	0xc7, 0xb3, 0x10, 0x4b, 0xa4, 0x4a, 0xea, 0x87, 0x46, 0x43, 0xdb, 0x90, 0x5c, 0x5b, 0xd1, 0x41,
	0x6e, 0xcd, 0x42, 0x6c, 0xc1, 0xcf, 0xab, 0x4b, 0x25, 0xf7, 0x95, 0x64, 0x54, 0x62, 0x16, 0xdd,
	0xff, 0x61, 0xaa, 0x41, 0xd6, 0x0d, 0xa2, 0xb8, 0x94, 0xa9, 0x92, 0x7a, 0xc1, 0x28, 0x2f, 0x84,
	0x59, 0xe8, 0x69, 0x89, 0x21, 0xa5, 0x9b, 0xf2, 0xe7, 0x4c, 0x71, 0x8f, 0x3e, 0x84, 0x83, 0xa1,
	0x3f, 0x8d, 0x62, 0xe4, 0xf6, 0x84, 0x5d, 0x60, 0x49, 0xa9, 0x92, 0x7a, 0xbe, 0x95, 0x4f, 0x4d,
	0xb2, 0x3c, 0x53, 0x25, 0x66, 0x61, 0x01, 0x9f, 0xb3, 0x0b, 0xa4, 0x2e, 0x1c, 0xe3, 0x47, 0x1c,
	0xc6, 0xf6, 0x74, 0x22, 0xad, 0x67, 0x72, 0x80, 0xa5, 0xac, 0xb0, 0x33, 0xb6, 0xc9, 0xd1, 0x49,
	0x05, 0xde, 0x2c, 0xf9, 0xdd, 0x1d, 0xf3, 0x08, 0x6f, 0x74, 0xe4, 0x7f, 0xbd, 0x87, 0x22, 0x73,
	0x1c, 0xfb, 0xa6, 0x4b, 0x4e, 0xb8, 0x68, 0xdb, 0xb8, 0x34, 0x1d, 0xa7, 0x7b, 0xed, 0x70, 0x87,
	0x5d, 0x57, 0x42, 0xbd, 0x75, 0x08, 0x39, 0xa1, 0x48, 0x73, 0x3f, 0xae, 0x2e, 0x15, 0x52, 0xfb,
	0x02, 0xf7, 0x37, 0xbe, 0x90, 0x7e, 0x80, 0x83, 0x11, 0xf3, 0xfc, 0x29, 0xc7, 0xd5, 0x99, 0xe9,
	0xdb, 0xbc, 0xe2, 0x85, 0xe4, 0xad, 0x0d, 0xae, 0x30, 0xfa, 0x07, 0xd4, 0x9e, 0xc0, 0xf1, 0xad,
	0x0f, 0xa7, 0xa7, 0x50, 0x18, 0x79, 0x3c, 0x8a, 0xe5, 0xca, 0x08, 0xdf, 0x7d, 0x13, 0x44, 0x4b,
	0x5c, 0x6d, 0x3c, 0x83, 0x7b, 0xb7, 0x9b, 0x51, 0x80, 0xdd, 0x66, 0xdb, 0x3a, 0x7b, 0xdb, 0x51,
	0x77, 0x68, 0x01, 0xf6, 0x7a, 0xcd, 0x7e, 0x3f, 0x2d, 0x48, 0x5a, 0x9c, 0x77, 0xac, 0x77, 0xaf,
	0xcd, 0x57, 0x6a, 0xa6, 0xf1, 0x14, 0x8a, 0x6b, 0x0b, 0x46, 0xf7, 0x21, 0xdb, 0xb5, 0xac, 0x9e,
	0xba, 0x43, 0xf7, 0x40, 0xb1, 0xda, 0x3d, 0x95, 0xa4, 0xad, 0x97, 0x66, 0xaf, 0xad, 0x66, 0x68,
	0x1e, 0x72, 0x66, 0xe7, 0xf9, 0x59, 0x5f, 0x55, 0x5a, 0xea, 0xb7, 0x79, 0x85, 0xfc, 0x9a, 0x57,
	0xc8, 0xef, 0x79, 0x85, 0xfc, 0x99, 0x57, 0xc8, 0x60, 0x57, 0xec, 0xff, 0xe3, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x82, 0xd8, 0x12, 0xca, 0xfa, 0x03, 0x00, 0x00,
}