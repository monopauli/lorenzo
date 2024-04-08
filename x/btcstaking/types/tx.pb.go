// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btcstaking/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_Lorenzo_Protocol_lorenzo_types "github.com/Lorenzo-Protocol/lorenzo/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TransactionKey struct {
	Index uint32                                                        `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Hash  *github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderHashBytes `protobuf:"bytes,2,opt,name=hash,proto3,customtype=github.com/Lorenzo-Protocol/lorenzo/types.BTCHeaderHashBytes" json:"hash,omitempty"`
}

func (m *TransactionKey) Reset()         { *m = TransactionKey{} }
func (m *TransactionKey) String() string { return proto.CompactTextString(m) }
func (*TransactionKey) ProtoMessage()    {}
func (*TransactionKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{0}
}
func (m *TransactionKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionKey.Merge(m, src)
}
func (m *TransactionKey) XXX_Size() int {
	return m.Size()
}
func (m *TransactionKey) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionKey.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionKey proto.InternalMessageInfo

func (m *TransactionKey) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

// TransactionInfo is the info of a tx on Bitcoin,
// including
// - the position of the tx on BTC blockchain
// - the full tx content
// - the Merkle proof that this tx is on the above position
type TransactionInfo struct {
	// key is the position (txIdx, blockHash) of this tx on BTC blockchain
	// Although it is already a part of SubmissionKey, we store it here again
	// to make TransactionInfo self-contained.
	// For example, storing the key allows TransactionInfo to not relay on
	// the fact that TransactionInfo will be ordered in the same order as
	// TransactionKeys in SubmissionKey.
	Key *TransactionKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// transaction is the full transaction in bytes
	Transaction []byte `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// proof is the Merkle proof that this tx is included in the position in `key`
	// TODO: maybe it could use here better format as we already processed and
	// validated the proof?
	Proof []byte `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *TransactionInfo) Reset()         { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()    {}
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{1}
}
func (m *TransactionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInfo.Merge(m, src)
}
func (m *TransactionInfo) XXX_Size() int {
	return m.Size()
}
func (m *TransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInfo proto.InternalMessageInfo

func (m *TransactionInfo) GetKey() *TransactionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TransactionInfo) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *TransactionInfo) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type MsgCreateBTCStaking struct {
	Signer    string           `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	StakingTx *TransactionInfo `protobuf:"bytes,2,opt,name=staking_tx,json=stakingTx,proto3" json:"staking_tx,omitempty"`
}

func (m *MsgCreateBTCStaking) Reset()         { *m = MsgCreateBTCStaking{} }
func (m *MsgCreateBTCStaking) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBTCStaking) ProtoMessage()    {}
func (*MsgCreateBTCStaking) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{2}
}
func (m *MsgCreateBTCStaking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBTCStaking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBTCStaking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBTCStaking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBTCStaking.Merge(m, src)
}
func (m *MsgCreateBTCStaking) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBTCStaking) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBTCStaking.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBTCStaking proto.InternalMessageInfo

func (m *MsgCreateBTCStaking) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgCreateBTCStaking) GetStakingTx() *TransactionInfo {
	if m != nil {
		return m.StakingTx
	}
	return nil
}

type MsgCreateBTCStakingResponse struct {
}

func (m *MsgCreateBTCStakingResponse) Reset()         { *m = MsgCreateBTCStakingResponse{} }
func (m *MsgCreateBTCStakingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBTCStakingResponse) ProtoMessage()    {}
func (*MsgCreateBTCStakingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be51bab5db52b8e, []int{3}
}
func (m *MsgCreateBTCStakingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBTCStakingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBTCStakingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBTCStakingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBTCStakingResponse.Merge(m, src)
}
func (m *MsgCreateBTCStakingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBTCStakingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBTCStakingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBTCStakingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TransactionKey)(nil), "lorenzo.btcstaking.v1.TransactionKey")
	proto.RegisterType((*TransactionInfo)(nil), "lorenzo.btcstaking.v1.TransactionInfo")
	proto.RegisterType((*MsgCreateBTCStaking)(nil), "lorenzo.btcstaking.v1.MsgCreateBTCStaking")
	proto.RegisterType((*MsgCreateBTCStakingResponse)(nil), "lorenzo.btcstaking.v1.MsgCreateBTCStakingResponse")
}

func init() { proto.RegisterFile("lorenzo/btcstaking/v1/tx.proto", fileDescriptor_6be51bab5db52b8e) }

var fileDescriptor_6be51bab5db52b8e = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x3b, 0xd6, 0x5d, 0xd8, 0x89, 0xff, 0x18, 0x57, 0xad, 0x11, 0x63, 0x09, 0x28, 0x4b,
	0xc1, 0x84, 0x8d, 0xa0, 0x20, 0x1e, 0x24, 0x45, 0x58, 0xd1, 0x85, 0x25, 0xe6, 0xe4, 0x65, 0x49,
	0xb3, 0xd3, 0x49, 0x68, 0x33, 0x13, 0x66, 0xc6, 0x92, 0x88, 0x87, 0xda, 0x4f, 0xe0, 0x47, 0xe9,
	0xc7, 0xf0, 0xd8, 0xa3, 0x78, 0x10, 0x69, 0x0f, 0xfd, 0x1a, 0x92, 0x99, 0x94, 0x56, 0xac, 0x6c,
	0x6f, 0x79, 0xde, 0xe7, 0xcd, 0xfb, 0xfc, 0xde, 0xe1, 0x85, 0xd6, 0x90, 0x71, 0x4c, 0x3f, 0x33,
	0xb7, 0x27, 0x63, 0x21, 0xa3, 0x41, 0x4a, 0x89, 0x3b, 0x3a, 0x76, 0x65, 0xe1, 0xe4, 0x9c, 0x49,
	0x86, 0xee, 0xd4, 0xbe, 0xb3, 0xf6, 0x9d, 0xd1, 0xb1, 0x79, 0x2f, 0x66, 0x22, 0x63, 0xc2, 0xcd,
	0x84, 0x6a, 0xcf, 0x04, 0xd1, 0xfd, 0xe6, 0x7d, 0x6d, 0x9c, 0x2b, 0xe5, 0x6a, 0x51, 0x5b, 0x87,
	0x84, 0x11, 0xa6, 0xeb, 0xd5, 0x97, 0xae, 0xda, 0x5f, 0xe0, 0x8d, 0x90, 0x47, 0x54, 0x44, 0xb1,
	0x4c, 0x19, 0x7d, 0x87, 0x4b, 0x74, 0x08, 0xf7, 0x52, 0x7a, 0x81, 0x8b, 0x16, 0x68, 0x83, 0xa3,
	0xeb, 0x81, 0x16, 0x28, 0x84, 0x57, 0x93, 0x48, 0x24, 0xad, 0x2b, 0x6d, 0x70, 0x74, 0xcd, 0x7f,
	0xfd, 0xf3, 0xd7, 0xa3, 0x57, 0x24, 0x95, 0xc9, 0xa7, 0x9e, 0x13, 0xb3, 0xcc, 0x7d, 0xaf, 0x29,
	0x9f, 0x9e, 0x55, 0x33, 0x63, 0x36, 0x74, 0x57, 0x6b, 0xc9, 0x32, 0xc7, 0xc2, 0xf1, 0xc3, 0xee,
	0x09, 0x8e, 0x2e, 0x30, 0x3f, 0x89, 0x44, 0xe2, 0x97, 0x12, 0x8b, 0x40, 0x4d, 0xb3, 0x27, 0x00,
	0xde, 0xdc, 0x88, 0x7f, 0x4b, 0xfb, 0x0c, 0xbd, 0x80, 0xcd, 0x01, 0x2e, 0x55, 0xba, 0xe1, 0x3d,
	0x76, 0xb6, 0x3e, 0x80, 0xf3, 0x37, 0x73, 0x50, 0xfd, 0x81, 0xda, 0xd0, 0x90, 0xeb, 0xb2, 0x26,
	0x0d, 0x36, 0x4b, 0xd5, 0x6a, 0x39, 0x67, 0xac, 0xdf, 0x6a, 0x2a, 0x4f, 0x0b, 0xfb, 0x2b, 0x80,
	0xb7, 0x4f, 0x05, 0xe9, 0x72, 0x1c, 0x49, 0xec, 0x87, 0xdd, 0x0f, 0x3a, 0x06, 0xdd, 0x85, 0xfb,
	0x22, 0x25, 0x14, 0x73, 0xc5, 0x72, 0x10, 0xd4, 0x0a, 0xbd, 0x81, 0xb0, 0x26, 0x39, 0x97, 0x85,
	0x8a, 0x31, 0xbc, 0x27, 0x97, 0x73, 0x56, 0xcb, 0x05, 0x07, 0xb5, 0x17, 0x16, 0x2f, 0x8d, 0xc9,
	0x72, 0xda, 0xa9, 0x67, 0xda, 0x0f, 0xe1, 0x83, 0x2d, 0x08, 0x01, 0x16, 0x39, 0xa3, 0x02, 0x7b,
	0x63, 0x00, 0x9b, 0xa7, 0x82, 0x20, 0x0e, 0x6f, 0xfd, 0x83, 0xd9, 0xf9, 0x4f, 0xf4, 0x96, 0x79,
	0xa6, 0xb7, 0x7b, 0xef, 0x2a, 0xdb, 0xdc, 0x1b, 0x2f, 0xa7, 0x1d, 0xe0, 0x9f, 0x7d, 0x9f, 0x5b,
	0x60, 0x36, 0xb7, 0xc0, 0xef, 0xb9, 0x05, 0xbe, 0x2d, 0xac, 0xc6, 0x6c, 0x61, 0x35, 0x7e, 0x2c,
	0xac, 0xc6, 0xc7, 0xe7, 0xbb, 0x1c, 0x42, 0xb1, 0x79, 0xe1, 0xea, 0x2a, 0x7a, 0xfb, 0xea, 0x02,
	0x9f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x11, 0x4e, 0x1f, 0xbf, 0x04, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateBTCDelegation creates a new BTC delegation
	CreateBTCStaking(ctx context.Context, in *MsgCreateBTCStaking, opts ...grpc.CallOption) (*MsgCreateBTCStakingResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBTCStaking(ctx context.Context, in *MsgCreateBTCStaking, opts ...grpc.CallOption) (*MsgCreateBTCStakingResponse, error) {
	out := new(MsgCreateBTCStakingResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btcstaking.v1.Msg/CreateBTCStaking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateBTCDelegation creates a new BTC delegation
	CreateBTCStaking(context.Context, *MsgCreateBTCStaking) (*MsgCreateBTCStakingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBTCStaking(ctx context.Context, req *MsgCreateBTCStaking) (*MsgCreateBTCStakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBTCStaking not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBTCStaking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBTCStaking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBTCStaking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btcstaking.v1.Msg/CreateBTCStaking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBTCStaking(ctx, req.(*MsgCreateBTCStaking))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorenzo.btcstaking.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBTCStaking",
			Handler:    _Msg_CreateBTCStaking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorenzo/btcstaking/v1/tx.proto",
}

func (m *TransactionKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Hash != nil {
		{
			size := m.Hash.Size()
			i -= size
			if _, err := m.Hash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TransactionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Transaction) > 0 {
		i -= len(m.Transaction)
		copy(dAtA[i:], m.Transaction)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Transaction)))
		i--
		dAtA[i] = 0x12
	}
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateBTCStaking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBTCStaking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBTCStaking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StakingTx != nil {
		{
			size, err := m.StakingTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateBTCStakingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBTCStakingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBTCStakingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovTx(uint64(m.Index))
	}
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *TransactionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Transaction)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateBTCStaking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StakingTx != nil {
		l = m.StakingTx.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateBTCStakingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: TransactionKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderHashBytes
			m.Hash = &v
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *TransactionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: TransactionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &TransactionKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transaction = append(m.Transaction[:0], dAtA[iNdEx:postIndex]...)
			if m.Transaction == nil {
				m.Transaction = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateBTCStaking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateBTCStaking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBTCStaking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StakingTx == nil {
				m.StakingTx = &TransactionInfo{}
			}
			if err := m.StakingTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateBTCStakingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateBTCStakingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBTCStakingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)