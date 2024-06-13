// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btcstaking/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9c98ee28fc92429, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9c98ee28fc92429, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type StakingRecordDisplay struct {
	TxId            string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	MintToAddress   string `protobuf:"bytes,2,opt,name=mint_to_address,json=mintToAddress,proto3" json:"mint_to_address,omitempty"`
	Amount          string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	BtcReceiverName string `protobuf:"bytes,4,opt,name=btc_receiver_name,json=btcReceiverName,proto3" json:"btc_receiver_name,omitempty"`
	BtcReceiverAddr string `protobuf:"bytes,5,opt,name=btc_receiver_addr,json=btcReceiverAddr,proto3" json:"btc_receiver_addr,omitempty"`
}

func (m *StakingRecordDisplay) Reset()         { *m = StakingRecordDisplay{} }
func (m *StakingRecordDisplay) String() string { return proto.CompactTextString(m) }
func (*StakingRecordDisplay) ProtoMessage()    {}
func (*StakingRecordDisplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9c98ee28fc92429, []int{2}
}
func (m *StakingRecordDisplay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakingRecordDisplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakingRecordDisplay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakingRecordDisplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingRecordDisplay.Merge(m, src)
}
func (m *StakingRecordDisplay) XXX_Size() int {
	return m.Size()
}
func (m *StakingRecordDisplay) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingRecordDisplay.DiscardUnknown(m)
}

var xxx_messageInfo_StakingRecordDisplay proto.InternalMessageInfo

func (m *StakingRecordDisplay) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *StakingRecordDisplay) GetMintToAddress() string {
	if m != nil {
		return m.MintToAddress
	}
	return ""
}

func (m *StakingRecordDisplay) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *StakingRecordDisplay) GetBtcReceiverName() string {
	if m != nil {
		return m.BtcReceiverName
	}
	return ""
}

func (m *StakingRecordDisplay) GetBtcReceiverAddr() string {
	if m != nil {
		return m.BtcReceiverAddr
	}
	return ""
}

type QueryStakingRecordRequest struct {
	TxHash []byte `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (m *QueryStakingRecordRequest) Reset()         { *m = QueryStakingRecordRequest{} }
func (m *QueryStakingRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStakingRecordRequest) ProtoMessage()    {}
func (*QueryStakingRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9c98ee28fc92429, []int{3}
}
func (m *QueryStakingRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStakingRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStakingRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStakingRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStakingRecordRequest.Merge(m, src)
}
func (m *QueryStakingRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStakingRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStakingRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStakingRecordRequest proto.InternalMessageInfo

func (m *QueryStakingRecordRequest) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

type QueryStakingRecordResponse struct {
	Record *BTCStakingRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (m *QueryStakingRecordResponse) Reset()         { *m = QueryStakingRecordResponse{} }
func (m *QueryStakingRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStakingRecordResponse) ProtoMessage()    {}
func (*QueryStakingRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9c98ee28fc92429, []int{4}
}
func (m *QueryStakingRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStakingRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStakingRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStakingRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStakingRecordResponse.Merge(m, src)
}
func (m *QueryStakingRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStakingRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStakingRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStakingRecordResponse proto.InternalMessageInfo

func (m *QueryStakingRecordResponse) GetRecord() *BTCStakingRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "lorenzo.btcstaking.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "lorenzo.btcstaking.v1.QueryParamsResponse")
	proto.RegisterType((*StakingRecordDisplay)(nil), "lorenzo.btcstaking.v1.StakingRecordDisplay")
	proto.RegisterType((*QueryStakingRecordRequest)(nil), "lorenzo.btcstaking.v1.QueryStakingRecordRequest")
	proto.RegisterType((*QueryStakingRecordResponse)(nil), "lorenzo.btcstaking.v1.QueryStakingRecordResponse")
}

func init() { proto.RegisterFile("lorenzo/btcstaking/v1/query.proto", fileDescriptor_d9c98ee28fc92429) }

var fileDescriptor_d9c98ee28fc92429 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x43, 0x63, 0x89, 0x85, 0xaa, 0x62, 0x1b, 0x90, 0xb1, 0xa8, 0x01, 0x0b, 0x0a, 0x44,
	0xe0, 0x25, 0xad, 0xe0, 0x8a, 0x28, 0x1c, 0x40, 0xaa, 0x50, 0x30, 0x3d, 0x21, 0x21, 0x6b, 0x6d,
	0xaf, 0x5c, 0x8b, 0x78, 0xc7, 0xf5, 0x6e, 0xa2, 0x84, 0x23, 0x07, 0xce, 0x48, 0xfc, 0x03, 0xfe,
	0x03, 0xbf, 0x80, 0x0b, 0xc7, 0x4a, 0x5c, 0x38, 0xa2, 0x84, 0x1f, 0x82, 0xbc, 0xbb, 0x7c, 0x44,
	0x71, 0x23, 0xb8, 0xad, 0x67, 0xdf, 0x7b, 0xf3, 0xe6, 0x79, 0x07, 0x5d, 0x1d, 0x42, 0xc5, 0xf8,
	0x1b, 0x20, 0xb1, 0x4c, 0x84, 0xa4, 0xaf, 0x73, 0x9e, 0x91, 0x71, 0x9f, 0x1c, 0x8d, 0x58, 0x35,
	0x0d, 0xca, 0x0a, 0x24, 0xe0, 0xf3, 0x06, 0x12, 0xfc, 0x81, 0x04, 0xe3, 0xbe, 0xdb, 0x6b, 0x66,
	0x9a, 0x63, 0x54, 0xb1, 0x04, 0xaa, 0x54, 0x4b, 0xb8, 0x7e, 0x33, 0xb6, 0xa4, 0x15, 0x2d, 0x84,
	0xc1, 0x74, 0x33, 0xc8, 0x40, 0x1d, 0x49, 0x7d, 0x32, 0xd5, 0x4b, 0x19, 0x40, 0x36, 0x64, 0x84,
	0x96, 0x39, 0xa1, 0x9c, 0x83, 0xa4, 0x32, 0x07, 0xfe, 0x8b, 0xd3, 0x4b, 0x40, 0x14, 0x20, 0x48,
	0x4c, 0x05, 0xd3, 0x9e, 0xc9, 0xb8, 0x1f, 0x33, 0x49, 0x6b, 0xed, 0x2c, 0xe7, 0x0a, 0xac, 0xb1,
	0x7e, 0x17, 0xe1, 0xe7, 0x35, 0x62, 0xa0, 0x9a, 0x86, 0xec, 0x68, 0xc4, 0x84, 0xf4, 0xf7, 0xd1,
	0xe6, 0x42, 0x55, 0x94, 0xc0, 0x05, 0xc3, 0xf7, 0x90, 0xad, 0xcd, 0x39, 0xd6, 0x15, 0xeb, 0xe6,
	0x99, 0x9d, 0xad, 0xa0, 0x31, 0x84, 0xc0, 0xd0, 0x0c, 0xd8, 0xff, 0x6c, 0xa1, 0xee, 0x0b, 0x7d,
	0x1b, 0xaa, 0xf9, 0x1f, 0xe7, 0xa2, 0x1c, 0xd2, 0x29, 0xde, 0x44, 0x1d, 0x39, 0x89, 0xf2, 0x54,
	0xc9, 0x9d, 0x0e, 0xd7, 0xe4, 0xe4, 0x69, 0x8a, 0xb7, 0xd1, 0x46, 0x91, 0x73, 0x19, 0x49, 0x88,
	0x68, 0x9a, 0x56, 0x4c, 0x08, 0xa7, 0xad, 0xae, 0xd7, 0xeb, 0xf2, 0x01, 0x3c, 0xd4, 0x45, 0x7c,
	0x01, 0xd9, 0xb4, 0x80, 0x11, 0x97, 0xce, 0x29, 0x75, 0x6d, 0xbe, 0x70, 0x0f, 0x9d, 0x8b, 0x65,
	0x52, 0x27, 0xcd, 0xf2, 0x31, 0xab, 0x22, 0x4e, 0x0b, 0xe6, 0xac, 0x29, 0xc8, 0x46, 0x2c, 0x93,
	0xd0, 0xd4, 0x9f, 0xd1, 0x82, 0x2d, 0x61, 0xeb, 0x86, 0x4e, 0x67, 0x09, 0x5b, 0xb7, 0xf4, 0x77,
	0xd1, 0x45, 0x95, 0xc9, 0xc2, 0x24, 0x26, 0xb0, 0xda, 0x8c, 0x9c, 0x3c, 0xa1, 0xe2, 0x50, 0x8d,
	0x72, 0x36, 0x34, 0x5f, 0xfe, 0x2b, 0xe4, 0x36, 0x91, 0x4c, 0x9e, 0x0f, 0x90, 0xad, 0x1f, 0x84,
	0xc9, 0xf3, 0xc6, 0x09, 0x79, 0xee, 0x1d, 0x3c, 0x5a, 0x14, 0x30, 0xb4, 0x9d, 0x4f, 0x6d, 0xd4,
	0x51, 0xfa, 0xf8, 0x9d, 0x85, 0x6c, 0x1d, 0x3b, 0xbe, 0x75, 0x82, 0xca, 0xf2, 0x7f, 0x76, 0x7b,
	0xff, 0x02, 0xd5, 0x66, 0xfd, 0xeb, 0x6f, 0xbf, 0xfe, 0xf8, 0xd0, 0xbe, 0x8c, 0xb7, 0xc8, 0xaa,
	0x67, 0x8b, 0x3f, 0x5a, 0x68, 0x7d, 0xc1, 0x2c, 0xbe, 0xbb, 0xaa, 0x49, 0x53, 0x9a, 0x6e, 0xff,
	0x3f, 0x18, 0xc6, 0xdd, 0x6d, 0xe5, 0x6e, 0x1b, 0x5f, 0x23, 0x2b, 0x17, 0x50, 0xb3, 0xf6, 0x06,
	0x5f, 0x66, 0x9e, 0x75, 0x3c, 0xf3, 0xac, 0xef, 0x33, 0xcf, 0x7a, 0x3f, 0xf7, 0x5a, 0xc7, 0x73,
	0xaf, 0xf5, 0x6d, 0xee, 0xb5, 0x5e, 0xde, 0xcf, 0x72, 0x79, 0x38, 0x8a, 0x83, 0x04, 0x0a, 0xb2,
	0xaf, 0x95, 0xee, 0x0c, 0xea, 0x4d, 0x49, 0x60, 0xf8, 0x5b, 0x7a, 0xf2, 0xb7, 0xb8, 0x9c, 0x96,
	0x4c, 0xc4, 0xb6, 0x5a, 0xa7, 0xdd, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x63, 0x38, 0x44,
	0x3a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	StakingRecord(ctx context.Context, in *QueryStakingRecordRequest, opts ...grpc.CallOption) (*QueryStakingRecordResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btcstaking.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StakingRecord(ctx context.Context, in *QueryStakingRecordRequest, opts ...grpc.CallOption) (*QueryStakingRecordResponse, error) {
	out := new(QueryStakingRecordResponse)
	err := c.cc.Invoke(ctx, "/lorenzo.btcstaking.v1.Query/StakingRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	StakingRecord(context.Context, *QueryStakingRecordRequest) (*QueryStakingRecordResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) StakingRecord(ctx context.Context, req *QueryStakingRecordRequest) (*QueryStakingRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakingRecord not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btcstaking.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StakingRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStakingRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StakingRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorenzo.btcstaking.v1.Query/StakingRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StakingRecord(ctx, req.(*QueryStakingRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorenzo.btcstaking.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "StakingRecord",
			Handler:    _Query_StakingRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorenzo/btcstaking/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StakingRecordDisplay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakingRecordDisplay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakingRecordDisplay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BtcReceiverAddr) > 0 {
		i -= len(m.BtcReceiverAddr)
		copy(dAtA[i:], m.BtcReceiverAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BtcReceiverAddr)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BtcReceiverName) > 0 {
		i -= len(m.BtcReceiverName)
		copy(dAtA[i:], m.BtcReceiverName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BtcReceiverName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MintToAddress) > 0 {
		i -= len(m.MintToAddress)
		copy(dAtA[i:], m.MintToAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MintToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStakingRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStakingRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStakingRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStakingRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStakingRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStakingRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Record != nil {
		{
			size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *StakingRecordDisplay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.MintToAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BtcReceiverName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BtcReceiverAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStakingRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStakingRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *StakingRecordDisplay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: StakingRecordDisplay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakingRecordDisplay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcReceiverName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcReceiverName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcReceiverAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcReceiverAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryStakingRecordRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryStakingRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStakingRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryStakingRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryStakingRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStakingRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Record == nil {
				m.Record = &BTCStakingRecord{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
