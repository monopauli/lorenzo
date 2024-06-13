// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/agent/v1/event.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// agent creation event
type EventAddAgent struct {
	// id is the unique identifier of the agent
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// agent name,required
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// btc_receiving_address is agent’s fund escrow address,required
	BtcReceivingAddress string `protobuf:"bytes,3,opt,name=btc_receiving_address,json=btcReceivingAddress,proto3" json:"btc_receiving_address,omitempty"`
	// like 0xBAb28FF7659481F1c8516f616A576339936AFB06
	EthAddr string `protobuf:"bytes,4,opt,name=eth_addr,json=ethAddr,proto3" json:"eth_addr,omitempty"`
	// description is a brief description of the agent, optional
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// url is the agent's link, used for detailed introduction, optional
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	// sender is the address of the governance account or module admin
	Sender string `protobuf:"bytes,7,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *EventAddAgent) Reset()         { *m = EventAddAgent{} }
func (m *EventAddAgent) String() string { return proto.CompactTextString(m) }
func (*EventAddAgent) ProtoMessage()    {}
func (*EventAddAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8e8f07d92c984c2, []int{0}
}
func (m *EventAddAgent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAddAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAddAgent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAddAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddAgent.Merge(m, src)
}
func (m *EventAddAgent) XXX_Size() int {
	return m.Size()
}
func (m *EventAddAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddAgent.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddAgent proto.InternalMessageInfo

func (m *EventAddAgent) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventAddAgent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventAddAgent) GetBtcReceivingAddress() string {
	if m != nil {
		return m.BtcReceivingAddress
	}
	return ""
}

func (m *EventAddAgent) GetEthAddr() string {
	if m != nil {
		return m.EthAddr
	}
	return ""
}

func (m *EventAddAgent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventAddAgent) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *EventAddAgent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// agent edit event
type EventEditAgent struct {
	// id is the unique identifier of the agent
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// agent name,required
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// description is a brief description of the agent, optional
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// url is the agent's link, used for detailed introduction, optional
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	// sender is the address of the governance account or module admin
	Sender string `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *EventEditAgent) Reset()         { *m = EventEditAgent{} }
func (m *EventEditAgent) String() string { return proto.CompactTextString(m) }
func (*EventEditAgent) ProtoMessage()    {}
func (*EventEditAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8e8f07d92c984c2, []int{1}
}
func (m *EventEditAgent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEditAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEditAgent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEditAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEditAgent.Merge(m, src)
}
func (m *EventEditAgent) XXX_Size() int {
	return m.Size()
}
func (m *EventEditAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEditAgent.DiscardUnknown(m)
}

var xxx_messageInfo_EventEditAgent proto.InternalMessageInfo

func (m *EventEditAgent) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventEditAgent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventEditAgent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventEditAgent) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *EventEditAgent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// agent remove event
type EventRemoveAgent struct {
	// id is the unique identifier of the agent
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// sender is the address of the governance account or module admin
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *EventRemoveAgent) Reset()         { *m = EventRemoveAgent{} }
func (m *EventRemoveAgent) String() string { return proto.CompactTextString(m) }
func (*EventRemoveAgent) ProtoMessage()    {}
func (*EventRemoveAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8e8f07d92c984c2, []int{2}
}
func (m *EventRemoveAgent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRemoveAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRemoveAgent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRemoveAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRemoveAgent.Merge(m, src)
}
func (m *EventRemoveAgent) XXX_Size() int {
	return m.Size()
}
func (m *EventRemoveAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRemoveAgent.DiscardUnknown(m)
}

var xxx_messageInfo_EventRemoveAgent proto.InternalMessageInfo

func (m *EventRemoveAgent) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventRemoveAgent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func init() {
	proto.RegisterType((*EventAddAgent)(nil), "lorenzo.agent.v1.EventAddAgent")
	proto.RegisterType((*EventEditAgent)(nil), "lorenzo.agent.v1.EventEditAgent")
	proto.RegisterType((*EventRemoveAgent)(nil), "lorenzo.agent.v1.EventRemoveAgent")
}

func init() { proto.RegisterFile("lorenzo/agent/v1/event.proto", fileDescriptor_f8e8f07d92c984c2) }

var fileDescriptor_f8e8f07d92c984c2 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0x97, 0xae, 0xdb, 0xde, 0x37, 0xe2, 0x18, 0x11, 0x25, 0x82, 0x84, 0xb1, 0xd3, 0x2e,
	0xb6, 0x4c, 0x6f, 0xde, 0x26, 0xec, 0xa4, 0x07, 0xe9, 0xd1, 0xcb, 0x58, 0x93, 0x87, 0x2d, 0xb0,
	0x26, 0x25, 0xcd, 0x8a, 0x7a, 0xf2, 0x23, 0xf8, 0xb1, 0x3c, 0xc9, 0x8e, 0x1e, 0xa5, 0xfd, 0x22,
	0xd2, 0xac, 0xc3, 0x29, 0x13, 0xbc, 0x25, 0xcf, 0xef, 0xff, 0xc0, 0xef, 0xe1, 0x8f, 0xcf, 0x96,
	0xda, 0x80, 0x7a, 0xd2, 0xe1, 0x6c, 0x0e, 0xca, 0x86, 0xf9, 0x28, 0x84, 0x1c, 0x94, 0x0d, 0x52,
	0xa3, 0xad, 0x26, 0xbd, 0x9a, 0x06, 0x8e, 0x06, 0xf9, 0x68, 0xf0, 0x86, 0xf0, 0xe1, 0xa4, 0x4a,
	0x8c, 0x85, 0x18, 0x57, 0x43, 0xd2, 0xc5, 0x9e, 0x14, 0x14, 0xf5, 0xd1, 0xd0, 0x8f, 0x3c, 0x29,
	0x08, 0xc1, 0xbe, 0x9a, 0x25, 0x40, 0xbd, 0x3e, 0x1a, 0xfe, 0x8f, 0xdc, 0x9b, 0x5c, 0xe0, 0xe3,
	0xd8, 0xf2, 0xa9, 0x01, 0x0e, 0x32, 0x97, 0x6a, 0x3e, 0x9d, 0x09, 0x61, 0x20, 0xcb, 0x68, 0xd3,
	0x85, 0x8e, 0x62, 0xcb, 0xa3, 0x2d, 0x1b, 0x6f, 0x10, 0x39, 0xc5, 0xff, 0xc0, 0x2e, 0x5c, 0x92,
	0xfa, 0x2e, 0xd6, 0x01, 0xbb, 0xa8, 0x28, 0xe9, 0xe3, 0x03, 0x01, 0x19, 0x37, 0x32, 0xb5, 0x52,
	0x2b, 0xda, 0x72, 0x74, 0x77, 0x44, 0x7a, 0xb8, 0xb9, 0x32, 0x4b, 0xda, 0x76, 0xa4, 0x7a, 0x92,
	0x13, 0xdc, 0xce, 0x40, 0x09, 0x30, 0xb4, 0xe3, 0x86, 0xf5, 0x6f, 0xf0, 0x8c, 0x70, 0xd7, 0x1d,
	0x34, 0x11, 0xd2, 0xfe, 0xfd, 0xa2, 0x1f, 0x0a, 0xcd, 0x5f, 0x15, 0xfc, 0x7d, 0x0a, 0xad, 0x6f,
	0x0a, 0x57, 0xb8, 0xe7, 0x0c, 0x22, 0x48, 0x74, 0x0e, 0xfb, 0x1d, 0xbe, 0x76, 0xbd, 0xdd, 0xdd,
	0xeb, 0x9b, 0xd7, 0x82, 0xa1, 0x75, 0xc1, 0xd0, 0x47, 0xc1, 0xd0, 0x4b, 0xc9, 0x1a, 0xeb, 0x92,
	0x35, 0xde, 0x4b, 0xd6, 0xb8, 0x1f, 0xcd, 0xa5, 0x5d, 0xac, 0xe2, 0x80, 0xeb, 0x24, 0xbc, 0xdd,
	0xd4, 0x78, 0x7e, 0x57, 0xb5, 0xca, 0xf5, 0x32, 0xdc, 0xb6, 0xfe, 0x50, 0xf7, 0x6e, 0x1f, 0x53,
	0xc8, 0xe2, 0xb6, 0x6b, 0xfd, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x44, 0x2b, 0x12, 0xdb, 0x15,
	0x02, 0x00, 0x00,
}

func (m *EventAddAgent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAddAgent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAddAgent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EthAddr) > 0 {
		i -= len(m.EthAddr)
		copy(dAtA[i:], m.EthAddr)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EthAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BtcReceivingAddress) > 0 {
		i -= len(m.BtcReceivingAddress)
		copy(dAtA[i:], m.BtcReceivingAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.BtcReceivingAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventEditAgent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEditAgent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEditAgent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRemoveAgent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRemoveAgent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRemoveAgent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventAddAgent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvent(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.BtcReceivingAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.EthAddr)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventEditAgent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvent(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventRemoveAgent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvent(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventAddAgent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventAddAgent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAddAgent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcReceivingAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcReceivingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventEditAgent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventEditAgent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEditAgent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventRemoveAgent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventRemoveAgent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRemoveAgent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
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
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)