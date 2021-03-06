// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/shorten.proto

package shorten

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateUrlReq struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	KeepAlive            int64    `protobuf:"varint,2,opt,name=keep_alive,json=keepAlive,proto3" json:"keep_alive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUrlReq) Reset()         { *m = CreateUrlReq{} }
func (m *CreateUrlReq) String() string { return proto.CompactTextString(m) }
func (*CreateUrlReq) ProtoMessage()    {}
func (*CreateUrlReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd06f1be78c5ee63, []int{0}
}

func (m *CreateUrlReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUrlReq.Unmarshal(m, b)
}
func (m *CreateUrlReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUrlReq.Marshal(b, m, deterministic)
}
func (m *CreateUrlReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUrlReq.Merge(m, src)
}
func (m *CreateUrlReq) XXX_Size() int {
	return xxx_messageInfo_CreateUrlReq.Size(m)
}
func (m *CreateUrlReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUrlReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUrlReq proto.InternalMessageInfo

func (m *CreateUrlReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CreateUrlReq) GetKeepAlive() int64 {
	if m != nil {
		return m.KeepAlive
	}
	return 0
}

type CreateUrlRes struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUrlRes) Reset()         { *m = CreateUrlRes{} }
func (m *CreateUrlRes) String() string { return proto.CompactTextString(m) }
func (*CreateUrlRes) ProtoMessage()    {}
func (*CreateUrlRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd06f1be78c5ee63, []int{1}
}

func (m *CreateUrlRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUrlRes.Unmarshal(m, b)
}
func (m *CreateUrlRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUrlRes.Marshal(b, m, deterministic)
}
func (m *CreateUrlRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUrlRes.Merge(m, src)
}
func (m *CreateUrlRes) XXX_Size() int {
	return xxx_messageInfo_CreateUrlRes.Size(m)
}
func (m *CreateUrlRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUrlRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUrlRes proto.InternalMessageInfo

func (m *CreateUrlRes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUrlReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUrlReq) Reset()         { *m = GetUrlReq{} }
func (m *GetUrlReq) String() string { return proto.CompactTextString(m) }
func (*GetUrlReq) ProtoMessage()    {}
func (*GetUrlReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd06f1be78c5ee63, []int{2}
}

func (m *GetUrlReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUrlReq.Unmarshal(m, b)
}
func (m *GetUrlReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUrlReq.Marshal(b, m, deterministic)
}
func (m *GetUrlReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUrlReq.Merge(m, src)
}
func (m *GetUrlReq) XXX_Size() int {
	return xxx_messageInfo_GetUrlReq.Size(m)
}
func (m *GetUrlReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUrlReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUrlReq proto.InternalMessageInfo

func (m *GetUrlReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Url struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	KeepAlive            int64    `protobuf:"varint,3,opt,name=keep_alive,json=keepAlive,proto3" json:"keep_alive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Url) Reset()         { *m = Url{} }
func (m *Url) String() string { return proto.CompactTextString(m) }
func (*Url) ProtoMessage()    {}
func (*Url) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd06f1be78c5ee63, []int{3}
}

func (m *Url) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Url.Unmarshal(m, b)
}
func (m *Url) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Url.Marshal(b, m, deterministic)
}
func (m *Url) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Url.Merge(m, src)
}
func (m *Url) XXX_Size() int {
	return xxx_messageInfo_Url.Size(m)
}
func (m *Url) XXX_DiscardUnknown() {
	xxx_messageInfo_Url.DiscardUnknown(m)
}

var xxx_messageInfo_Url proto.InternalMessageInfo

func (m *Url) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Url) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Url) GetKeepAlive() int64 {
	if m != nil {
		return m.KeepAlive
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateUrlReq)(nil), "CreateUrlReq")
	proto.RegisterType((*CreateUrlRes)(nil), "CreateUrlRes")
	proto.RegisterType((*GetUrlReq)(nil), "GetUrlReq")
	proto.RegisterType((*Url)(nil), "Url")
}

func init() { proto.RegisterFile("protos/shorten.proto", fileDescriptor_cd06f1be78c5ee63) }

var fileDescriptor_cd06f1be78c5ee63 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0xce, 0xc8, 0x2f, 0x2a, 0x49, 0xcd, 0xd3, 0x03, 0x73, 0x95, 0xec, 0xb9,
	0x78, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x43, 0x8b, 0x72, 0x82, 0x52, 0x0b, 0x85, 0x04, 0xb8,
	0x98, 0x4b, 0x8b, 0x72, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x59, 0x2e,
	0xae, 0xec, 0xd4, 0xd4, 0x82, 0xf8, 0xc4, 0x9c, 0xcc, 0xb2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0xe6, 0x20, 0x4e, 0x90, 0x88, 0x23, 0x48, 0x40, 0x49, 0x0e, 0xc5, 0x80, 0x62, 0x21, 0x3e, 0x2e,
	0xa6, 0xcc, 0x14, 0xa8, 0x7e, 0xa6, 0xcc, 0x14, 0x25, 0x69, 0x2e, 0x4e, 0xf7, 0xd4, 0x12, 0xa8,
	0xe9, 0xe8, 0x92, 0x6e, 0x5c, 0xcc, 0xa1, 0x45, 0x39, 0xe8, 0xc2, 0x30, 0x47, 0x30, 0xe1, 0x72,
	0x04, 0x33, 0x9a, 0x23, 0x8c, 0xc2, 0xb9, 0xf8, 0x82, 0x21, 0xde, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x15, 0xd2, 0xe4, 0xe2, 0x84, 0x3b, 0x4b, 0x88, 0x57, 0x0f, 0xd9, 0x8f, 0x52, 0x28,
	0xdc, 0x62, 0x21, 0x29, 0x2e, 0x36, 0x88, 0x0b, 0x85, 0xb8, 0xf4, 0xe0, 0x4e, 0x95, 0x62, 0xd1,
	0x0b, 0x2d, 0xca, 0x49, 0x62, 0x03, 0x87, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x80, 0x71,
	0x4e, 0x5f, 0x3d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShortenServiceClient is the client API for ShortenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShortenServiceClient interface {
	// Create a shortened url.
	CreateUrl(ctx context.Context, in *CreateUrlReq, opts ...grpc.CallOption) (*CreateUrlRes, error)
	// Get a shortened url.
	GetUrl(ctx context.Context, in *GetUrlReq, opts ...grpc.CallOption) (*Url, error)
}

type shortenServiceClient struct {
	cc *grpc.ClientConn
}

func NewShortenServiceClient(cc *grpc.ClientConn) ShortenServiceClient {
	return &shortenServiceClient{cc}
}

func (c *shortenServiceClient) CreateUrl(ctx context.Context, in *CreateUrlReq, opts ...grpc.CallOption) (*CreateUrlRes, error) {
	out := new(CreateUrlRes)
	err := c.cc.Invoke(ctx, "/ShortenService/CreateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenServiceClient) GetUrl(ctx context.Context, in *GetUrlReq, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/ShortenService/GetUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenServiceServer is the server API for ShortenService service.
type ShortenServiceServer interface {
	// Create a shortened url.
	CreateUrl(context.Context, *CreateUrlReq) (*CreateUrlRes, error)
	// Get a shortened url.
	GetUrl(context.Context, *GetUrlReq) (*Url, error)
}

// UnimplementedShortenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShortenServiceServer struct {
}

func (*UnimplementedShortenServiceServer) CreateUrl(ctx context.Context, req *CreateUrlReq) (*CreateUrlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUrl not implemented")
}
func (*UnimplementedShortenServiceServer) GetUrl(ctx context.Context, req *GetUrlReq) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}

func RegisterShortenServiceServer(s *grpc.Server, srv ShortenServiceServer) {
	s.RegisterService(&_ShortenService_serviceDesc, srv)
}

func _ShortenService_CreateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenServiceServer).CreateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShortenService/CreateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenServiceServer).CreateUrl(ctx, req.(*CreateUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenService_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenServiceServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShortenService/GetUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenServiceServer).GetUrl(ctx, req.(*GetUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShortenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShortenService",
	HandlerType: (*ShortenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUrl",
			Handler:    _ShortenService_CreateUrl_Handler,
		},
		{
			MethodName: "GetUrl",
			Handler:    _ShortenService_GetUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shorten.proto",
}
