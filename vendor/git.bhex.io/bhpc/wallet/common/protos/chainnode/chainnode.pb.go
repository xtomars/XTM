// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chainnode.proto

package chainnode

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "git.bhex.io/bhpc/wallet/common/protos/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Action int32

const (
	Action_ADD    Action = 0
	Action_REMOVE Action = 1
	Action_CREATE Action = 2
)

var Action_name = map[int32]string{
	0: "ADD",
	1: "REMOVE",
	2: "CREATE",
}
var Action_value = map[string]int32{
	"ADD":    0,
	"REMOVE": 1,
	"CREATE": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}
func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{0}
}

type QueryCmd int32

const (
	QueryCmd_BALANCE              QueryCmd = 0
	QueryCmd_NONCE                QueryCmd = 1
	QueryCmd_GAS_PRICE            QueryCmd = 2
	QueryCmd_BLOCK_HEIGHT         QueryCmd = 3
	QueryCmd_DECIMALS             QueryCmd = 4
	QueryCmd_BALANCE_OF           QueryCmd = 5
	QueryCmd_ADDRESS_EXIST        QueryCmd = 6
	QueryCmd_HEADBLOCK_ID         QueryCmd = 7
	QueryCmd_TX_ID                QueryCmd = 8
	QueryCmd_ADDRESS_CONTRACT     QueryCmd = 9
	QueryCmd_CREATE_RAW_UNSIGN_TX QueryCmd = 10
	QueryCmd_CHAIN_ID             QueryCmd = 11
	QueryCmd_CHAIN_STATUS         QueryCmd = 12
)

var QueryCmd_name = map[int32]string{
	0:  "BALANCE",
	1:  "NONCE",
	2:  "GAS_PRICE",
	3:  "BLOCK_HEIGHT",
	4:  "DECIMALS",
	5:  "BALANCE_OF",
	6:  "ADDRESS_EXIST",
	7:  "HEADBLOCK_ID",
	8:  "TX_ID",
	9:  "ADDRESS_CONTRACT",
	10: "CREATE_RAW_UNSIGN_TX",
	11: "CHAIN_ID",
	12: "CHAIN_STATUS",
}
var QueryCmd_value = map[string]int32{
	"BALANCE":              0,
	"NONCE":                1,
	"GAS_PRICE":            2,
	"BLOCK_HEIGHT":         3,
	"DECIMALS":             4,
	"BALANCE_OF":           5,
	"ADDRESS_EXIST":        6,
	"HEADBLOCK_ID":         7,
	"TX_ID":                8,
	"ADDRESS_CONTRACT":     9,
	"CREATE_RAW_UNSIGN_TX": 10,
	"CHAIN_ID":             11,
	"CHAIN_STATUS":         12,
}

func (x QueryCmd) String() string {
	return proto.EnumName(QueryCmd_name, int32(x))
}
func (QueryCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{1}
}

type AddressRequest struct {
	TokenId              string   `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Action               Action   `protobuf:"varint,3,opt,name=action,proto3,enum=chainnode.Action" json:"action,omitempty"`
	Expiration           uint64   `protobuf:"varint,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressRequest) Reset()         { *m = AddressRequest{} }
func (m *AddressRequest) String() string { return proto.CompactTextString(m) }
func (*AddressRequest) ProtoMessage()    {}
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{0}
}
func (m *AddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressRequest.Unmarshal(m, b)
}
func (m *AddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressRequest.Marshal(b, m, deterministic)
}
func (dst *AddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressRequest.Merge(dst, src)
}
func (m *AddressRequest) XXX_Size() int {
	return xxx_messageInfo_AddressRequest.Size(m)
}
func (m *AddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressRequest proto.InternalMessageInfo

func (m *AddressRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *AddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_ADD
}

func (m *AddressRequest) GetExpiration() uint64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type AddressReply struct {
	Code                 common.ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=common.ReturnCode" json:"code,omitempty"`
	Msg                  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Address              string            `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddressReply) Reset()         { *m = AddressReply{} }
func (m *AddressReply) String() string { return proto.CompactTextString(m) }
func (*AddressReply) ProtoMessage()    {}
func (*AddressReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{1}
}
func (m *AddressReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressReply.Unmarshal(m, b)
}
func (m *AddressReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressReply.Marshal(b, m, deterministic)
}
func (dst *AddressReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressReply.Merge(dst, src)
}
func (m *AddressReply) XXX_Size() int {
	return xxx_messageInfo_AddressReply.Size(m)
}
func (m *AddressReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddressReply proto.InternalMessageInfo

func (m *AddressReply) GetCode() common.ReturnCode {
	if m != nil {
		return m.Code
	}
	return common.ReturnCode_SUCCESS
}

func (m *AddressReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AddressReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type SendTransactionRequest struct {
	TokenId              string            `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	BusinessNumber       string            `protobuf:"bytes,2,opt,name=businessNumber,proto3" json:"businessNumber,omitempty"`
	TxData               []byte            `protobuf:"bytes,3,opt,name=txData,proto3" json:"txData,omitempty"`
	Extension            map[string]string `protobuf:"bytes,4,rep,name=extension,proto3" json:"extension,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SendTransactionRequest) Reset()         { *m = SendTransactionRequest{} }
func (m *SendTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SendTransactionRequest) ProtoMessage()    {}
func (*SendTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{2}
}
func (m *SendTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTransactionRequest.Unmarshal(m, b)
}
func (m *SendTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTransactionRequest.Marshal(b, m, deterministic)
}
func (dst *SendTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTransactionRequest.Merge(dst, src)
}
func (m *SendTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SendTransactionRequest.Size(m)
}
func (m *SendTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendTransactionRequest proto.InternalMessageInfo

func (m *SendTransactionRequest) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *SendTransactionRequest) GetBusinessNumber() string {
	if m != nil {
		return m.BusinessNumber
	}
	return ""
}

func (m *SendTransactionRequest) GetTxData() []byte {
	if m != nil {
		return m.TxData
	}
	return nil
}

func (m *SendTransactionRequest) GetExtension() map[string]string {
	if m != nil {
		return m.Extension
	}
	return nil
}

type SendTransactionReply struct {
	Code                 common.ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=common.ReturnCode" json:"code,omitempty"`
	Msg                  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	TxHash               string            `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Extension            map[string]string `protobuf:"bytes,4,rep,name=extension,proto3" json:"extension,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SendTransactionReply) Reset()         { *m = SendTransactionReply{} }
func (m *SendTransactionReply) String() string { return proto.CompactTextString(m) }
func (*SendTransactionReply) ProtoMessage()    {}
func (*SendTransactionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{3}
}
func (m *SendTransactionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTransactionReply.Unmarshal(m, b)
}
func (m *SendTransactionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTransactionReply.Marshal(b, m, deterministic)
}
func (dst *SendTransactionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTransactionReply.Merge(dst, src)
}
func (m *SendTransactionReply) XXX_Size() int {
	return xxx_messageInfo_SendTransactionReply.Size(m)
}
func (m *SendTransactionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTransactionReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendTransactionReply proto.InternalMessageInfo

func (m *SendTransactionReply) GetCode() common.ReturnCode {
	if m != nil {
		return m.Code
	}
	return common.ReturnCode_SUCCESS
}

func (m *SendTransactionReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SendTransactionReply) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *SendTransactionReply) GetExtension() map[string]string {
	if m != nil {
		return m.Extension
	}
	return nil
}

type QueryRequest struct {
	Cmd                  QueryCmd          `protobuf:"varint,1,opt,name=cmd,proto3,enum=chainnode.QueryCmd" json:"cmd,omitempty"`
	Paras                []string          `protobuf:"bytes,2,rep,name=paras,proto3" json:"paras,omitempty"`
	Extension            map[string]string `protobuf:"bytes,3,rep,name=extension,proto3" json:"extension,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{4}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetCmd() QueryCmd {
	if m != nil {
		return m.Cmd
	}
	return QueryCmd_BALANCE
}

func (m *QueryRequest) GetParas() []string {
	if m != nil {
		return m.Paras
	}
	return nil
}

func (m *QueryRequest) GetExtension() map[string]string {
	if m != nil {
		return m.Extension
	}
	return nil
}

type QueryReply struct {
	Code                 common.ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=common.ReturnCode" json:"code,omitempty"`
	Msg                  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Result               uint64            `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	Extension            map[string]string `protobuf:"bytes,4,rep,name=extension,proto3" json:"extension,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainnode_1078e547442ad984, []int{5}
}
func (m *QueryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReply.Unmarshal(m, b)
}
func (m *QueryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReply.Marshal(b, m, deterministic)
}
func (dst *QueryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReply.Merge(dst, src)
}
func (m *QueryReply) XXX_Size() int {
	return xxx_messageInfo_QueryReply.Size(m)
}
func (m *QueryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReply proto.InternalMessageInfo

func (m *QueryReply) GetCode() common.ReturnCode {
	if m != nil {
		return m.Code
	}
	return common.ReturnCode_SUCCESS
}

func (m *QueryReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *QueryReply) GetResult() uint64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *QueryReply) GetExtension() map[string]string {
	if m != nil {
		return m.Extension
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressRequest)(nil), "chainnode.AddressRequest")
	proto.RegisterType((*AddressReply)(nil), "chainnode.AddressReply")
	proto.RegisterType((*SendTransactionRequest)(nil), "chainnode.SendTransactionRequest")
	proto.RegisterMapType((map[string]string)(nil), "chainnode.SendTransactionRequest.ExtensionEntry")
	proto.RegisterType((*SendTransactionReply)(nil), "chainnode.SendTransactionReply")
	proto.RegisterMapType((map[string]string)(nil), "chainnode.SendTransactionReply.ExtensionEntry")
	proto.RegisterType((*QueryRequest)(nil), "chainnode.QueryRequest")
	proto.RegisterMapType((map[string]string)(nil), "chainnode.QueryRequest.ExtensionEntry")
	proto.RegisterType((*QueryReply)(nil), "chainnode.QueryReply")
	proto.RegisterMapType((map[string]string)(nil), "chainnode.QueryReply.ExtensionEntry")
	proto.RegisterEnum("chainnode.Action", Action_name, Action_value)
	proto.RegisterEnum("chainnode.QueryCmd", QueryCmd_name, QueryCmd_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChainnodeClient is the client API for Chainnode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainnodeClient interface {
	// Wallet调用通知地址的增删信息
	Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressReply, error)
	// Wallet调用发送签名交易到区块链网络
	SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionReply, error)
	SendOnlineWalletTransaction(ctx context.Context, in *common.SendOnlineWalletTransactionRequest, opts ...grpc.CallOption) (*common.SendOnlineWalletTransactionReply, error)
	// Wallet查询链上信息
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
}

type chainnodeClient struct {
	cc *grpc.ClientConn
}

func NewChainnodeClient(cc *grpc.ClientConn) ChainnodeClient {
	return &chainnodeClient{cc}
}

func (c *chainnodeClient) Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressReply, error) {
	out := new(AddressReply)
	err := c.cc.Invoke(ctx, "/chainnode.Chainnode/Address", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainnodeClient) SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionReply, error) {
	out := new(SendTransactionReply)
	err := c.cc.Invoke(ctx, "/chainnode.Chainnode/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainnodeClient) SendOnlineWalletTransaction(ctx context.Context, in *common.SendOnlineWalletTransactionRequest, opts ...grpc.CallOption) (*common.SendOnlineWalletTransactionReply, error) {
	out := new(common.SendOnlineWalletTransactionReply)
	err := c.cc.Invoke(ctx, "/chainnode.Chainnode/SendOnlineWalletTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainnodeClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/chainnode.Chainnode/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainnodeServer is the server API for Chainnode service.
type ChainnodeServer interface {
	// Wallet调用通知地址的增删信息
	Address(context.Context, *AddressRequest) (*AddressReply, error)
	// Wallet调用发送签名交易到区块链网络
	SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionReply, error)
	SendOnlineWalletTransaction(context.Context, *common.SendOnlineWalletTransactionRequest) (*common.SendOnlineWalletTransactionReply, error)
	// Wallet查询链上信息
	Query(context.Context, *QueryRequest) (*QueryReply, error)
}

func RegisterChainnodeServer(s *grpc.Server, srv ChainnodeServer) {
	s.RegisterService(&_Chainnode_serviceDesc, srv)
}

func _Chainnode_Address_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainnodeServer).Address(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainnode.Chainnode/Address",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainnodeServer).Address(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chainnode_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainnodeServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainnode.Chainnode/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainnodeServer).SendTransaction(ctx, req.(*SendTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chainnode_SendOnlineWalletTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.SendOnlineWalletTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainnodeServer).SendOnlineWalletTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainnode.Chainnode/SendOnlineWalletTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainnodeServer).SendOnlineWalletTransaction(ctx, req.(*common.SendOnlineWalletTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chainnode_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainnodeServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chainnode.Chainnode/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainnodeServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chainnode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainnode.Chainnode",
	HandlerType: (*ChainnodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Address",
			Handler:    _Chainnode_Address_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _Chainnode_SendTransaction_Handler,
		},
		{
			MethodName: "SendOnlineWalletTransaction",
			Handler:    _Chainnode_SendOnlineWalletTransaction_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Chainnode_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chainnode.proto",
}

func init() { proto.RegisterFile("chainnode.proto", fileDescriptor_chainnode_1078e547442ad984) }

var fileDescriptor_chainnode_1078e547442ad984 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdb, 0x4e, 0xdb, 0x4c,
	0x10, 0xc6, 0x71, 0x8e, 0x43, 0x08, 0xcb, 0xfe, 0xf9, 0x21, 0xa4, 0x52, 0x4b, 0xa3, 0x16, 0x05,
	0x2e, 0x92, 0x2a, 0xbd, 0xe8, 0x41, 0xed, 0x85, 0xb1, 0x5d, 0x62, 0x35, 0x38, 0xed, 0xda, 0x14,
	0xee, 0x2c, 0x27, 0x5e, 0x11, 0x8b, 0xc4, 0x4e, 0x6d, 0xa7, 0x4d, 0x2e, 0xfb, 0x06, 0x7d, 0x33,
	0x5e, 0xa1, 0x7d, 0x87, 0x3e, 0x40, 0xb5, 0x3e, 0x40, 0x20, 0x14, 0x90, 0x10, 0x57, 0xde, 0xd9,
	0x99, 0x9d, 0xef, 0xfb, 0x66, 0x76, 0xc7, 0xb0, 0xda, 0x1f, 0x98, 0xb6, 0xe3, 0xb8, 0x16, 0x6d,
	0x8c, 0x3d, 0x37, 0x70, 0x71, 0xe1, 0x7c, 0xa3, 0xfa, 0xe6, 0xc4, 0x0e, 0x1a, 0xbd, 0x01, 0x9d,
	0x36, 0x6c, 0xb7, 0xd9, 0x1b, 0x8c, 0xfb, 0xcd, 0xef, 0xe6, 0x70, 0x48, 0x83, 0x66, 0xdf, 0x1d,
	0x8d, 0x5c, 0xa7, 0x19, 0x1e, 0xf0, 0x13, 0x2b, 0xfa, 0x44, 0x59, 0x6a, 0x3f, 0x39, 0x28, 0x09,
	0x96, 0xe5, 0x51, 0xdf, 0x27, 0xf4, 0xeb, 0x84, 0xfa, 0x01, 0xde, 0x84, 0x7c, 0xe0, 0x9e, 0x52,
	0xc7, 0xb0, 0xad, 0x0a, 0xb7, 0xc5, 0xd5, 0x0b, 0x24, 0x17, 0xda, 0x8a, 0x85, 0x2b, 0x90, 0x33,
	0xa3, 0xe0, 0x4a, 0x2a, 0xf2, 0xc4, 0x26, 0xde, 0x81, 0xac, 0xd9, 0x0f, 0x6c, 0xd7, 0xa9, 0xf0,
	0x5b, 0x5c, 0xbd, 0xd4, 0x5a, 0x6b, 0x5c, 0xf0, 0x15, 0x42, 0x07, 0x89, 0x03, 0xf0, 0x63, 0x00,
	0x3a, 0x1d, 0xdb, 0x9e, 0x19, 0x86, 0xa7, 0xb7, 0xb8, 0x7a, 0x9a, 0xcc, 0xed, 0xd4, 0x7a, 0x50,
	0x3c, 0x67, 0x34, 0x1e, 0xce, 0xf0, 0x36, 0xa4, 0xfb, 0xae, 0x45, 0x43, 0x2e, 0xa5, 0x16, 0x6e,
	0xc4, 0xfc, 0x09, 0x0d, 0x26, 0x9e, 0x23, 0xba, 0x16, 0x25, 0xa1, 0x1f, 0x23, 0xe0, 0x47, 0xfe,
	0x49, 0x4c, 0x8c, 0x2d, 0xe7, 0xe9, 0xf2, 0x97, 0xe8, 0xd6, 0x7e, 0xa4, 0x60, 0x5d, 0xa3, 0x8e,
	0xa5, 0x7b, 0xa6, 0xe3, 0x47, 0xbc, 0xee, 0x20, 0x7f, 0x1b, 0x4a, 0xbd, 0x89, 0x6f, 0x3b, 0xd4,
	0xf7, 0xd5, 0xc9, 0xa8, 0x47, 0xbd, 0x18, 0xec, 0xca, 0x2e, 0x5e, 0x87, 0x6c, 0x30, 0x95, 0xcc,
	0xc0, 0x0c, 0x61, 0x8b, 0x24, 0xb6, 0xb0, 0x0a, 0x05, 0x3a, 0x0d, 0xa8, 0xe3, 0x47, 0xc2, 0xf9,
	0xfa, 0x72, 0xeb, 0xc5, 0x5c, 0x9d, 0xae, 0x27, 0xd4, 0x90, 0x93, 0x23, 0xb2, 0x13, 0x78, 0x33,
	0x72, 0x91, 0xa2, 0xfa, 0x0e, 0x4a, 0x97, 0x9d, 0xac, 0x06, 0xa7, 0x74, 0x16, 0xf3, 0x66, 0x4b,
	0x5c, 0x86, 0xcc, 0x37, 0x73, 0x38, 0xa1, 0x31, 0xd5, 0xc8, 0x78, 0x9b, 0x7a, 0xcd, 0xd5, 0xfe,
	0x70, 0x50, 0x5e, 0x80, 0xbc, 0x5f, 0xc1, 0x43, 0xe1, 0x6d, 0xd3, 0x1f, 0xc4, 0xf5, 0x8e, 0x2d,
	0xdc, 0x59, 0x14, 0xde, 0xb8, 0x49, 0xf8, 0x78, 0x38, 0x7b, 0x30, 0xd9, 0x67, 0x1c, 0x14, 0x3f,
	0x4f, 0xa8, 0x37, 0x4b, 0x1a, 0xfe, 0x1c, 0xf8, 0xfe, 0xc8, 0x8a, 0xd5, 0xfe, 0x37, 0x47, 0x2b,
	0x8c, 0x12, 0x47, 0x16, 0x61, 0x7e, 0x96, 0x71, 0x6c, 0x7a, 0x26, 0xbb, 0xf9, 0x3c, 0xcb, 0x18,
	0x1a, 0x58, 0x9a, 0x57, 0xc6, 0x87, 0xca, 0xb6, 0xaf, 0xa6, 0x78, 0xe8, 0x46, 0xfe, 0xe2, 0x00,
	0x62, 0xa0, 0x7b, 0xb7, 0xcf, 0xa3, 0xfe, 0x64, 0x18, 0x84, 0xed, 0x4b, 0x93, 0xd8, 0xc2, 0x7b,
	0x8b, 0xed, 0x7b, 0xb6, 0x28, 0xf2, 0x01, 0x9b, 0xb6, 0xbb, 0x03, 0xd9, 0x68, 0x8a, 0xe0, 0x1c,
	0xf0, 0x82, 0x24, 0xa1, 0x25, 0x0c, 0x90, 0x25, 0xf2, 0x41, 0xf7, 0x8b, 0x8c, 0x38, 0xb6, 0x16,
	0x89, 0x2c, 0xe8, 0x32, 0x4a, 0xed, 0xfe, 0xe6, 0x20, 0x9f, 0x74, 0x0e, 0x2f, 0x43, 0x6e, 0x4f,
	0xe8, 0x08, 0xaa, 0x28, 0xa3, 0x25, 0x5c, 0x80, 0x8c, 0xda, 0x65, 0x4b, 0x0e, 0xaf, 0x40, 0x61,
	0x5f, 0xd0, 0x8c, 0x4f, 0x44, 0x11, 0x65, 0x94, 0xc2, 0x08, 0x8a, 0x7b, 0x9d, 0xae, 0xf8, 0xd1,
	0x68, 0xcb, 0xca, 0x7e, 0x5b, 0x47, 0x3c, 0x2e, 0x42, 0x5e, 0x92, 0x45, 0xe5, 0x40, 0xe8, 0x68,
	0x28, 0x8d, 0x4b, 0x00, 0x71, 0x1a, 0xa3, 0xfb, 0x01, 0x65, 0xf0, 0x1a, 0xac, 0x08, 0x92, 0x44,
	0x64, 0x4d, 0x33, 0xe4, 0x63, 0x45, 0xd3, 0x51, 0x96, 0xa5, 0x68, 0xcb, 0x82, 0x14, 0xa5, 0x51,
	0x24, 0x94, 0x63, 0x70, 0xfa, 0x31, 0x5b, 0xe6, 0x71, 0x19, 0x50, 0x12, 0x2f, 0x76, 0x55, 0x9d,
	0x08, 0xa2, 0x8e, 0x0a, 0xb8, 0x02, 0xe5, 0x88, 0xb5, 0x41, 0x84, 0x23, 0xe3, 0x50, 0xd5, 0x94,
	0x7d, 0xd5, 0xd0, 0x8f, 0x11, 0x30, 0x74, 0xb1, 0x2d, 0x28, 0x2a, 0x3b, 0xbd, 0xcc, 0x52, 0x47,
	0x96, 0xa6, 0x0b, 0xfa, 0xa1, 0x86, 0x8a, 0xad, 0xb3, 0x14, 0x14, 0xc4, 0xa4, 0xfe, 0xf8, 0x3d,
	0xe4, 0xe2, 0x81, 0x89, 0x37, 0xe7, 0xc7, 0xee, 0xa5, 0xb1, 0x5e, 0xdd, 0xb8, 0xce, 0xc5, 0xee,
	0xcb, 0x21, 0xac, 0x5e, 0x79, 0x80, 0xf8, 0xe9, 0xad, 0x53, 0xa9, 0xfa, 0xe4, 0x96, 0xf7, 0x8b,
	0x7d, 0x78, 0xc4, 0xf6, 0xbb, 0xce, 0xd0, 0x76, 0xe8, 0x51, 0xf8, 0x3f, 0x9a, 0x87, 0xd8, 0x4d,
	0xee, 0xe5, 0x0d, 0x41, 0x09, 0x56, 0xfd, 0x4e, 0xb1, 0x0c, 0xf4, 0x15, 0x64, 0xc2, 0xde, 0xe3,
	0x8d, 0x7f, 0x3c, 0xc2, 0xea, 0xff, 0xd7, 0x5e, 0xdc, 0x5e, 0x36, 0xfc, 0x1d, 0xbe, 0xfc, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x38, 0x7d, 0x78, 0xcd, 0x67, 0x07, 0x00, 0x00,
}