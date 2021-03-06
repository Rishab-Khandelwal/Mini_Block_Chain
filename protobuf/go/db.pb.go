// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

/*
Package miniblockchaindb is a generated protocol buffer package.

It is generated from these files:
	db.proto

It has these top-level messages:
	GetBlockRequest
	GetRequest
	GetResponse
	GetHeightResponse
	BooleanResponse
	VerifyResponse
	Null
	Transaction
	JsonBlockString
	Block
*/
package miniblockchaindb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type VerifyResponse_Results int32

const (
	VerifyResponse_FAILED    VerifyResponse_Results = 0
	VerifyResponse_PENDING   VerifyResponse_Results = 1
	VerifyResponse_SUCCEEDED VerifyResponse_Results = 2
)

var VerifyResponse_Results_name = map[int32]string{
	0: "FAILED",
	1: "PENDING",
	2: "SUCCEEDED",
}
var VerifyResponse_Results_value = map[string]int32{
	"FAILED":    0,
	"PENDING":   1,
	"SUCCEEDED": 2,
}

func (x VerifyResponse_Results) Enum() *VerifyResponse_Results {
	p := new(VerifyResponse_Results)
	*p = x
	return p
}
func (x VerifyResponse_Results) String() string {
	return proto.EnumName(VerifyResponse_Results_name, int32(x))
}
func (x *VerifyResponse_Results) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VerifyResponse_Results_value, data, "VerifyResponse_Results")
	if err != nil {
		return err
	}
	*x = VerifyResponse_Results(value)
	return nil
}
func (VerifyResponse_Results) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Transaction_Types int32

const (
	Transaction_UNKNOWN  Transaction_Types = 0
	Transaction_TRANSFER Transaction_Types = 5
)

var Transaction_Types_name = map[int32]string{
	0: "UNKNOWN",
	5: "TRANSFER",
}
var Transaction_Types_value = map[string]int32{
	"UNKNOWN":  0,
	"TRANSFER": 5,
}

func (x Transaction_Types) Enum() *Transaction_Types {
	p := new(Transaction_Types)
	*p = x
	return p
}
func (x Transaction_Types) String() string {
	return proto.EnumName(Transaction_Types_name, int32(x))
}
func (x *Transaction_Types) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Transaction_Types_value, data, "Transaction_Types")
	if err != nil {
		return err
	}
	*x = Transaction_Types(value)
	return nil
}
func (Transaction_Types) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type GetBlockRequest struct {
	BlockHash        *string `protobuf:"bytes,1,req,name=BlockHash" json:"BlockHash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetBlockRequest) Reset()                    { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()               {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetBlockRequest) GetBlockHash() string {
	if m != nil && m.BlockHash != nil {
		return *m.BlockHash
	}
	return ""
}

type GetRequest struct {
	UserID           *string `protobuf:"bytes,1,req,name=UserID" json:"UserID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRequest) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

type GetResponse struct {
	Value            *int32 `protobuf:"varint,1,req,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type GetHeightResponse struct {
	Height           *int32  `protobuf:"varint,1,req,name=Height" json:"Height,omitempty"`
	LeafHash         *string `protobuf:"bytes,2,req,name=LeafHash" json:"LeafHash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetHeightResponse) Reset()                    { *m = GetHeightResponse{} }
func (m *GetHeightResponse) String() string            { return proto.CompactTextString(m) }
func (*GetHeightResponse) ProtoMessage()               {}
func (*GetHeightResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetHeightResponse) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *GetHeightResponse) GetLeafHash() string {
	if m != nil && m.LeafHash != nil {
		return *m.LeafHash
	}
	return ""
}

type BooleanResponse struct {
	Success          *bool  `protobuf:"varint,1,req,name=Success" json:"Success,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BooleanResponse) Reset()                    { *m = BooleanResponse{} }
func (m *BooleanResponse) String() string            { return proto.CompactTextString(m) }
func (*BooleanResponse) ProtoMessage()               {}
func (*BooleanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BooleanResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

type VerifyResponse struct {
	Result           *VerifyResponse_Results `protobuf:"varint,1,req,name=Result,enum=miniblockchaindb.VerifyResponse_Results" json:"Result,omitempty"`
	BlockHash        *string                 `protobuf:"bytes,2,req,name=BlockHash" json:"BlockHash,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *VerifyResponse) Reset()                    { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string            { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()               {}
func (*VerifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VerifyResponse) GetResult() VerifyResponse_Results {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return VerifyResponse_FAILED
}

func (m *VerifyResponse) GetBlockHash() string {
	if m != nil && m.BlockHash != nil {
		return *m.BlockHash
	}
	return ""
}

type Null struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Null) Reset()                    { *m = Null{} }
func (m *Null) String() string            { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()               {}
func (*Null) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Transaction struct {
	Type *Transaction_Types `protobuf:"varint,1,req,name=Type,enum=miniblockchaindb.Transaction_Types" json:"Type,omitempty"`
	// reserved 2;
	FromID           *string `protobuf:"bytes,3,req,name=FromID" json:"FromID,omitempty"`
	ToID             *string `protobuf:"bytes,4,req,name=ToID" json:"ToID,omitempty"`
	Value            *int32  `protobuf:"varint,5,req,name=Value" json:"Value,omitempty"`
	MiningFee        *int32  `protobuf:"varint,6,req,name=MiningFee" json:"MiningFee,omitempty"`
	UUID             *string `protobuf:"bytes,7,req,name=UUID" json:"UUID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Transaction) GetType() Transaction_Types {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Transaction_UNKNOWN
}

func (m *Transaction) GetFromID() string {
	if m != nil && m.FromID != nil {
		return *m.FromID
	}
	return ""
}

func (m *Transaction) GetToID() string {
	if m != nil && m.ToID != nil {
		return *m.ToID
	}
	return ""
}

func (m *Transaction) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Transaction) GetMiningFee() int32 {
	if m != nil && m.MiningFee != nil {
		return *m.MiningFee
	}
	return 0
}

func (m *Transaction) GetUUID() string {
	if m != nil && m.UUID != nil {
		return *m.UUID
	}
	return ""
}

type JsonBlockString struct {
	Json             *string `protobuf:"bytes,1,req,name=Json" json:"Json,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JsonBlockString) Reset()                    { *m = JsonBlockString{} }
func (m *JsonBlockString) String() string            { return proto.CompactTextString(m) }
func (*JsonBlockString) ProtoMessage()               {}
func (*JsonBlockString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *JsonBlockString) GetJson() string {
	if m != nil && m.Json != nil {
		return *m.Json
	}
	return ""
}

type Block struct {
	BlockID          *int32         `protobuf:"varint,1,req,name=BlockID" json:"BlockID,omitempty"`
	PrevHash         *string        `protobuf:"bytes,2,req,name=PrevHash" json:"PrevHash,omitempty"`
	Transactions     []*Transaction `protobuf:"bytes,3,rep,name=Transactions" json:"Transactions,omitempty"`
	MinerID          *string        `protobuf:"bytes,4,req,name=MinerID" json:"MinerID,omitempty"`
	Nonce            *string        `protobuf:"bytes,5,req,name=Nonce" json:"Nonce,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Block) GetBlockID() int32 {
	if m != nil && m.BlockID != nil {
		return *m.BlockID
	}
	return 0
}

func (m *Block) GetPrevHash() string {
	if m != nil && m.PrevHash != nil {
		return *m.PrevHash
	}
	return ""
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetMinerID() string {
	if m != nil && m.MinerID != nil {
		return *m.MinerID
	}
	return ""
}

func (m *Block) GetNonce() string {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBlockRequest)(nil), "miniblockchaindb.GetBlockRequest")
	proto.RegisterType((*GetRequest)(nil), "miniblockchaindb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "miniblockchaindb.GetResponse")
	proto.RegisterType((*GetHeightResponse)(nil), "miniblockchaindb.GetHeightResponse")
	proto.RegisterType((*BooleanResponse)(nil), "miniblockchaindb.BooleanResponse")
	proto.RegisterType((*VerifyResponse)(nil), "miniblockchaindb.VerifyResponse")
	proto.RegisterType((*Null)(nil), "miniblockchaindb.Null")
	proto.RegisterType((*Transaction)(nil), "miniblockchaindb.Transaction")
	proto.RegisterType((*JsonBlockString)(nil), "miniblockchaindb.JsonBlockString")
	proto.RegisterType((*Block)(nil), "miniblockchaindb.Block")
	proto.RegisterEnum("miniblockchaindb.VerifyResponse_Results", VerifyResponse_Results_name, VerifyResponse_Results_value)
	proto.RegisterEnum("miniblockchaindb.Transaction_Types", Transaction_Types_name, Transaction_Types_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlockChainMiner service

type BlockChainMinerClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Transfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*BooleanResponse, error)
	Verify(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*VerifyResponse, error)
	GetHeight(ctx context.Context, in *Null, opts ...grpc.CallOption) (*GetHeightResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*JsonBlockString, error)
	PushBlock(ctx context.Context, in *JsonBlockString, opts ...grpc.CallOption) (*Null, error)
	PushTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Null, error)
}

type blockChainMinerClient struct {
	cc *grpc.ClientConn
}

func NewBlockChainMinerClient(cc *grpc.ClientConn) BlockChainMinerClient {
	return &blockChainMinerClient{cc}
}

func (c *blockChainMinerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/miniblockchaindb.BlockChainMiner/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainMinerClient) Transfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := grpc.Invoke(ctx, "/miniblockchaindb.BlockChainMiner/Transfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainMinerClient) Verify(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := grpc.Invoke(ctx, "/miniblockchaindb.BlockChainMiner/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainMinerClient) GetHeight(ctx context.Context, in *Null, opts ...grpc.CallOption) (*GetHeightResponse, error) {
	out := new(GetHeightResponse)
	err := grpc.Invoke(ctx, "/miniblockchaindb.BlockChainMiner/GetHeight", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainMinerClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*JsonBlockString, error) {
	out := new(JsonBlockString)
	err := grpc.Invoke(ctx, "/miniblockchaindb.BlockChainMiner/GetBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainMinerClient) PushBlock(ctx context.Context, in *JsonBlockString, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := grpc.Invoke(ctx, "/miniblockchaindb.BlockChainMiner/PushBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainMinerClient) PushTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := grpc.Invoke(ctx, "/miniblockchaindb.BlockChainMiner/PushTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlockChainMiner service

type BlockChainMinerServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Transfer(context.Context, *Transaction) (*BooleanResponse, error)
	Verify(context.Context, *Transaction) (*VerifyResponse, error)
	GetHeight(context.Context, *Null) (*GetHeightResponse, error)
	GetBlock(context.Context, *GetBlockRequest) (*JsonBlockString, error)
	PushBlock(context.Context, *JsonBlockString) (*Null, error)
	PushTransaction(context.Context, *Transaction) (*Null, error)
}

func RegisterBlockChainMinerServer(s *grpc.Server, srv BlockChainMinerServer) {
	s.RegisterService(&_BlockChainMiner_serviceDesc, srv)
}

func _BlockChainMiner_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainMinerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniblockchaindb.BlockChainMiner/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainMinerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChainMiner_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainMinerServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniblockchaindb.BlockChainMiner/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainMinerServer).Transfer(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChainMiner_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainMinerServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniblockchaindb.BlockChainMiner/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainMinerServer).Verify(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChainMiner_GetHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainMinerServer).GetHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniblockchaindb.BlockChainMiner/GetHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainMinerServer).GetHeight(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChainMiner_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainMinerServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniblockchaindb.BlockChainMiner/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainMinerServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChainMiner_PushBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonBlockString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainMinerServer).PushBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniblockchaindb.BlockChainMiner/PushBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainMinerServer).PushBlock(ctx, req.(*JsonBlockString))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChainMiner_PushTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainMinerServer).PushTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miniblockchaindb.BlockChainMiner/PushTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainMinerServer).PushTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockChainMiner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "miniblockchaindb.BlockChainMiner",
	HandlerType: (*BlockChainMinerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BlockChainMiner_Get_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _BlockChainMiner_Transfer_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _BlockChainMiner_Verify_Handler,
		},
		{
			MethodName: "GetHeight",
			Handler:    _BlockChainMiner_GetHeight_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BlockChainMiner_GetBlock_Handler,
		},
		{
			MethodName: "PushBlock",
			Handler:    _BlockChainMiner_PushBlock_Handler,
		},
		{
			MethodName: "PushTransaction",
			Handler:    _BlockChainMiner_PushTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db.proto",
}

func init() { proto.RegisterFile("db.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0x35, 0x01, 0x1b, 0xfb, 0x42, 0xb0, 0x33, 0x2b, 0x07, 0x25, 0x22, 0x1a, 0x65, 0x91, 0xcd,
	0xc7, 0xa7, 0xa2, 0x76, 0xd7, 0x45, 0x03, 0x36, 0xc4, 0x69, 0xe2, 0x22, 0x7e, 0xd2, 0xb5, 0x71,
	0x06, 0xb0, 0xea, 0x7a, 0xa8, 0xc7, 0xae, 0x14, 0xf5, 0x15, 0xfa, 0x2c, 0x7d, 0xb8, 0x3e, 0x41,
	0x35, 0x63, 0x30, 0x98, 0x10, 0x76, 0xf6, 0xdc, 0x33, 0x67, 0xee, 0x39, 0xe7, 0x5e, 0x50, 0x9f,
	0x67, 0xed, 0x55, 0x4c, 0x13, 0x8a, 0xea, 0xb3, 0x90, 0xfa, 0xdf, 0xfc, 0xa5, 0x17, 0x44, 0xcf,
	0x33, 0x7c, 0x0d, 0xfa, 0x80, 0x24, 0x5d, 0x7e, 0x34, 0x22, 0x3f, 0x52, 0xc2, 0x12, 0x74, 0x06,
	0x9a, 0xf8, 0xbf, 0xf3, 0xd8, 0xd2, 0x2c, 0x5d, 0x9d, 0xdc, 0x68, 0xf8, 0x02, 0x60, 0x40, 0x92,
	0x0d, 0xa0, 0x01, 0xca, 0x94, 0x91, 0xd8, 0xb1, 0xf2, 0x6a, 0x4d, 0x54, 0xd9, 0x8a, 0x46, 0x8c,
	0xa0, 0x53, 0x90, 0x9f, 0xbc, 0x30, 0x25, 0xa2, 0x2a, 0xe3, 0x0f, 0x70, 0x36, 0x20, 0xc9, 0x1d,
	0x09, 0x16, 0xcb, 0x2d, 0xa6, 0x01, 0x4a, 0x76, 0x92, 0x81, 0x90, 0x01, 0xea, 0x03, 0xf1, 0xe6,
	0xe2, 0xc9, 0x13, 0x41, 0x8a, 0x41, 0xef, 0x52, 0x1a, 0x12, 0x2f, 0xca, 0x2f, 0xe9, 0x50, 0x1d,
	0xa7, 0xbe, 0x4f, 0x18, 0x13, 0xb7, 0x54, 0xfc, 0xbb, 0x04, 0x8d, 0x27, 0x12, 0x07, 0xf3, 0x97,
	0x1c, 0xf3, 0x1e, 0x94, 0x11, 0x61, 0x69, 0x98, 0x11, 0x37, 0x3a, 0xd7, 0xed, 0x5d, 0xb9, 0xed,
	0x22, 0xba, 0x9d, 0x41, 0x59, 0x51, 0x72, 0xf6, 0xfe, 0x3b, 0xa8, 0x6e, 0xaa, 0x00, 0x4a, 0xff,
	0xd6, 0x79, 0xb0, 0x2d, 0x43, 0x42, 0x35, 0xa8, 0x0e, 0x6d, 0xd7, 0x72, 0xdc, 0x81, 0x51, 0x42,
	0xa7, 0xa0, 0x8d, 0xa7, 0xbd, 0x9e, 0x6d, 0x5b, 0xb6, 0x65, 0x9c, 0x60, 0x05, 0x2a, 0x6e, 0x1a,
	0x86, 0xf8, 0x4f, 0x09, 0x6a, 0x93, 0xd8, 0x8b, 0x98, 0xe7, 0x27, 0x01, 0x8d, 0xd0, 0x7f, 0x50,
	0x99, 0xbc, 0xac, 0xc8, 0xba, 0xa3, 0x56, 0xb1, 0xa3, 0x1d, 0x60, 0x9b, 0xa3, 0x18, 0xf7, 0xa6,
	0x1f, 0xd3, 0xef, 0x8e, 0x65, 0x96, 0x79, 0x27, 0xa8, 0x0e, 0x95, 0x09, 0x75, 0x2c, 0xb3, 0x22,
	0xfe, 0x72, 0x77, 0x65, 0x61, 0xdc, 0x19, 0x68, 0x8f, 0x41, 0x14, 0x44, 0x8b, 0x3e, 0x21, 0xa6,
	0x22, 0x8e, 0xea, 0x50, 0x99, 0x4e, 0x1d, 0xcb, 0xac, 0xae, 0x7d, 0x94, 0x33, 0xda, 0x1a, 0x54,
	0xa7, 0xee, 0x67, 0xf7, 0xcb, 0x57, 0xd7, 0x90, 0x50, 0x1d, 0xd4, 0xc9, 0xe8, 0xd6, 0x1d, 0xf7,
	0xed, 0x91, 0x21, 0xe3, 0x16, 0xe8, 0xf7, 0x8c, 0x46, 0xc2, 0x82, 0x71, 0x12, 0x07, 0xd1, 0x82,
	0x93, 0xf0, 0xa3, 0x75, 0xc2, 0xbf, 0x40, 0x16, 0x45, 0x1e, 0x81, 0xf8, 0x58, 0x67, 0x2f, 0x82,
	0x1b, 0xc6, 0xe4, 0xe7, 0xd6, 0x38, 0xf4, 0x3f, 0xd4, 0x77, 0x34, 0x31, 0xb3, 0x7c, 0x55, 0xbe,
	0xa9, 0x75, 0xce, 0xdf, 0x54, 0xcd, 0x39, 0x1f, 0x83, 0x48, 0xcc, 0x53, 0x2e, 0xd1, 0xa5, 0x91,
	0x9f, 0x49, 0xd4, 0x3a, 0x7f, 0xcb, 0xa0, 0x8b, 0x47, 0x7b, 0xfc, 0xb2, 0x80, 0xa2, 0x8f, 0x50,
	0x1e, 0x90, 0x04, 0x99, 0x45, 0xd6, 0xed, 0x8c, 0x36, 0xcf, 0x0f, 0x54, 0xb2, 0xd0, 0xb1, 0x84,
	0xfa, 0xa0, 0x8a, 0x06, 0xe6, 0x24, 0x46, 0x6f, 0x37, 0xd6, 0xbc, 0x2c, 0x96, 0xf6, 0xc6, 0x11,
	0x4b, 0xa8, 0x07, 0x4a, 0x36, 0x50, 0xc7, 0x58, 0x2e, 0x8e, 0x4d, 0x20, 0x96, 0x50, 0x17, 0xb4,
	0x7c, 0x3f, 0x10, 0x2a, 0x82, 0xf9, 0x38, 0x35, 0x5b, 0xaf, 0xa4, 0x14, 0x97, 0x09, 0x4b, 0xe8,
	0x1e, 0xd4, 0xcd, 0x16, 0xa3, 0xcb, 0x57, 0xf0, 0xdd, 0xed, 0xde, 0x17, 0xb5, 0x97, 0x3b, 0x96,
	0xd0, 0x27, 0xd0, 0x86, 0x29, 0x5b, 0x1e, 0x24, 0xdb, 0x43, 0x37, 0x0f, 0xb4, 0x2b, 0x14, 0xe9,
	0x9c, 0x61, 0x37, 0xe3, 0x23, 0xfe, 0x1c, 0xe4, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xa0,
	0x0a, 0x20, 0xb0, 0x04, 0x00, 0x00,
}
