// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package proto

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

type Types int32

const (
	Types_SYSTEM Types = 0
	Types_MAP    Types = 1
	Types_CHAT   Types = 2
	Types_OTHER  Types = 99
)

var Types_name = map[int32]string{
	0:  "SYSTEM",
	1:  "MAP",
	2:  "CHAT",
	99: "OTHER",
}

var Types_value = map[string]int32{
	"SYSTEM": 0,
	"MAP":    1,
	"CHAT":   2,
	"OTHER":  99,
}

func (x Types) String() string {
	return proto.EnumName(Types_name, int32(x))
}

func (Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

type SessionRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRequest) Reset()         { *m = SessionRequest{} }
func (m *SessionRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()    {}
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *SessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRequest.Unmarshal(m, b)
}
func (m *SessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRequest.Marshal(b, m, deterministic)
}
func (m *SessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRequest.Merge(m, src)
}
func (m *SessionRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRequest.Size(m)
}
func (m *SessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRequest proto.InternalMessageInfo

func (m *SessionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SessionRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SessionGrant struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionGrant) Reset()         { *m = SessionGrant{} }
func (m *SessionGrant) String() string { return proto.CompactTextString(m) }
func (*SessionGrant) ProtoMessage()    {}
func (*SessionGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *SessionGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionGrant.Unmarshal(m, b)
}
func (m *SessionGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionGrant.Marshal(b, m, deterministic)
}
func (m *SessionGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionGrant.Merge(m, src)
}
func (m *SessionGrant) XXX_Size() int {
	return xxx_messageInfo_SessionGrant.Size(m)
}
func (m *SessionGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionGrant.DiscardUnknown(m)
}

var xxx_messageInfo_SessionGrant proto.InternalMessageInfo

func (m *SessionGrant) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SessionGrant) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type LookAroundRequest struct {
	Myuuid               string   `protobuf:"bytes,1,opt,name=myuuid,proto3" json:"myuuid,omitempty"`
	Otheruuid            string   `protobuf:"bytes,2,opt,name=otheruuid,proto3" json:"otheruuid,omitempty"`
	Type                 Types    `protobuf:"varint,3,opt,name=type,proto3,enum=gameproto.v1.Types" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookAroundRequest) Reset()         { *m = LookAroundRequest{} }
func (m *LookAroundRequest) String() string { return proto.CompactTextString(m) }
func (*LookAroundRequest) ProtoMessage()    {}
func (*LookAroundRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *LookAroundRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookAroundRequest.Unmarshal(m, b)
}
func (m *LookAroundRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookAroundRequest.Marshal(b, m, deterministic)
}
func (m *LookAroundRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookAroundRequest.Merge(m, src)
}
func (m *LookAroundRequest) XXX_Size() int {
	return xxx_messageInfo_LookAroundRequest.Size(m)
}
func (m *LookAroundRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookAroundRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookAroundRequest proto.InternalMessageInfo

func (m *LookAroundRequest) GetMyuuid() string {
	if m != nil {
		return m.Myuuid
	}
	return ""
}

func (m *LookAroundRequest) GetOtheruuid() string {
	if m != nil {
		return m.Otheruuid
	}
	return ""
}

func (m *LookAroundRequest) GetType() Types {
	if m != nil {
		return m.Type
	}
	return Types_SYSTEM
}

type LookAroundAnswer struct {
	Results              []*Object `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Type                 Types     `protobuf:"varint,2,opt,name=type,proto3,enum=gameproto.v1.Types" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LookAroundAnswer) Reset()         { *m = LookAroundAnswer{} }
func (m *LookAroundAnswer) String() string { return proto.CompactTextString(m) }
func (*LookAroundAnswer) ProtoMessage()    {}
func (*LookAroundAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *LookAroundAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookAroundAnswer.Unmarshal(m, b)
}
func (m *LookAroundAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookAroundAnswer.Marshal(b, m, deterministic)
}
func (m *LookAroundAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookAroundAnswer.Merge(m, src)
}
func (m *LookAroundAnswer) XXX_Size() int {
	return xxx_messageInfo_LookAroundAnswer.Size(m)
}
func (m *LookAroundAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_LookAroundAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_LookAroundAnswer proto.InternalMessageInfo

func (m *LookAroundAnswer) GetResults() []*Object {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *LookAroundAnswer) GetType() Types {
	if m != nil {
		return m.Type
	}
	return Types_SYSTEM
}

type ObjectAttribute struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Currentvalue         string   `protobuf:"bytes,2,opt,name=currentvalue,proto3" json:"currentvalue,omitempty"`
	Targetvalue          string   `protobuf:"bytes,3,opt,name=targetvalue,proto3" json:"targetvalue,omitempty"`
	Ttl                  int32    `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectAttribute) Reset()         { *m = ObjectAttribute{} }
func (m *ObjectAttribute) String() string { return proto.CompactTextString(m) }
func (*ObjectAttribute) ProtoMessage()    {}
func (*ObjectAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4}
}

func (m *ObjectAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectAttribute.Unmarshal(m, b)
}
func (m *ObjectAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectAttribute.Marshal(b, m, deterministic)
}
func (m *ObjectAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectAttribute.Merge(m, src)
}
func (m *ObjectAttribute) XXX_Size() int {
	return xxx_messageInfo_ObjectAttribute.Size(m)
}
func (m *ObjectAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectAttribute proto.InternalMessageInfo

func (m *ObjectAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectAttribute) GetCurrentvalue() string {
	if m != nil {
		return m.Currentvalue
	}
	return ""
}

func (m *ObjectAttribute) GetTargetvalue() string {
	if m != nil {
		return m.Targetvalue
	}
	return ""
}

func (m *ObjectAttribute) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type Object struct {
	Uuid                 string             `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	X                    int32              `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32              `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Attributes           []*ObjectAttribute `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{5}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Object) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Object) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Object) GetAttributes() []*ObjectAttribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterEnum("gameproto.v1.Types", Types_name, Types_value)
	proto.RegisterType((*SessionRequest)(nil), "gameproto.v1.SessionRequest")
	proto.RegisterType((*SessionGrant)(nil), "gameproto.v1.SessionGrant")
	proto.RegisterType((*LookAroundRequest)(nil), "gameproto.v1.LookAroundRequest")
	proto.RegisterType((*LookAroundAnswer)(nil), "gameproto.v1.LookAroundAnswer")
	proto.RegisterType((*ObjectAttribute)(nil), "gameproto.v1.ObjectAttribute")
	proto.RegisterType((*Object)(nil), "gameproto.v1.Object")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x65, 0x13, 0xdb, 0x69, 0xa6, 0x51, 0x31, 0x03, 0xaa, 0xac, 0xa8, 0x80, 0xe5, 0x0b, 0x16,
	0x07, 0x8b, 0x86, 0x1b, 0x12, 0x07, 0x83, 0x0a, 0x39, 0x50, 0x0a, 0x4e, 0x2e, 0x70, 0x73, 0x92,
	0x51, 0x09, 0x4d, 0xbc, 0x61, 0x3f, 0x9a, 0xfa, 0xc0, 0xbf, 0xe1, 0x87, 0x22, 0xaf, 0xed, 0xc4,
	0x16, 0x41, 0x3d, 0x79, 0xde, 0xcc, 0x9b, 0x79, 0x9e, 0xd9, 0x07, 0x70, 0x9d, 0xae, 0x29, 0xda,
	0x08, 0xae, 0x38, 0x0e, 0x8a, 0xd8, 0x84, 0xd1, 0xed, 0x79, 0x30, 0x86, 0x93, 0x09, 0x49, 0xb9,
	0xe4, 0x59, 0x42, 0xbf, 0x34, 0x49, 0x85, 0x43, 0x38, 0xd2, 0x92, 0x44, 0x96, 0xae, 0xc9, 0x63,
	0x3e, 0x0b, 0xfb, 0xc9, 0x0e, 0x17, 0xb5, 0x4d, 0x2a, 0xe5, 0x96, 0x8b, 0x85, 0xd7, 0x29, 0x6b,
	0x35, 0x0e, 0xde, 0xc0, 0xa0, 0x9a, 0xf4, 0x51, 0xa4, 0x99, 0x42, 0x04, 0x4b, 0xeb, 0xe5, 0xa2,
	0x9a, 0x61, 0x62, 0x3c, 0x05, 0x47, 0x50, 0x2a, 0x79, 0x56, 0x75, 0x57, 0x28, 0x10, 0xf0, 0xe8,
	0x13, 0xe7, 0x37, 0xb1, 0xe0, 0x3a, 0x5b, 0xd4, 0x3f, 0x72, 0x0a, 0xce, 0x3a, 0x6f, 0x8c, 0xa8,
	0x10, 0x9e, 0x41, 0x9f, 0xab, 0x1f, 0x24, 0x4c, 0xa9, 0x9c, 0xb3, 0x4f, 0xe0, 0x0b, 0xb0, 0x54,
	0xbe, 0x21, 0xaf, 0xeb, 0xb3, 0xf0, 0x64, 0xf4, 0x38, 0x6a, 0x6e, 0x1b, 0x4d, 0xf3, 0x0d, 0xc9,
	0xc4, 0x10, 0x82, 0x1b, 0x70, 0xf7, 0x9a, 0x71, 0x26, 0xb7, 0x24, 0x30, 0x82, 0x9e, 0x20, 0xa9,
	0x57, 0x4a, 0x7a, 0xcc, 0xef, 0x86, 0xc7, 0xa3, 0x27, 0xed, 0xfe, 0xab, 0xd9, 0x4f, 0x9a, 0xab,
	0xa4, 0x26, 0xed, 0xc4, 0x3a, 0xf7, 0x89, 0xfd, 0x86, 0x87, 0x65, 0x6f, 0xac, 0x94, 0x58, 0xce,
	0xb4, 0xa2, 0xe2, 0x3e, 0x8d, 0x1b, 0x9b, 0x18, 0x03, 0x18, 0xcc, 0xb5, 0x10, 0x94, 0xa9, 0xdb,
	0x74, 0xa5, 0xa9, 0xda, 0xae, 0x95, 0x43, 0x1f, 0x8e, 0x55, 0x2a, 0xae, 0xa9, 0xa2, 0x74, 0x0d,
	0xa5, 0x99, 0x42, 0x17, 0xba, 0x4a, 0xad, 0x3c, 0xcb, 0x67, 0xa1, 0x9d, 0x14, 0x61, 0xb0, 0x05,
	0xa7, 0x94, 0x3f, 0xf8, 0x2a, 0x03, 0x60, 0x77, 0x46, 0xca, 0x4e, 0xd8, 0x5d, 0x81, 0x72, 0x33,
	0xd5, 0x4e, 0x58, 0x8e, 0x6f, 0x01, 0xd2, 0xfa, 0x97, 0xa5, 0x67, 0x99, 0xa3, 0x3c, 0x3d, 0x74,
	0x94, 0xdd, 0x62, 0x49, 0xa3, 0xe1, 0xe5, 0x39, 0xd8, 0xe6, 0x0c, 0x08, 0xe0, 0x4c, 0xbe, 0x4d,
	0xa6, 0x17, 0x97, 0xee, 0x03, 0xec, 0x41, 0xf7, 0x32, 0xfe, 0xe2, 0x32, 0x3c, 0x02, 0xeb, 0xfd,
	0x38, 0x9e, 0xba, 0x1d, 0xec, 0x83, 0x7d, 0x35, 0x1d, 0x5f, 0x24, 0xee, 0x7c, 0xf4, 0x87, 0x41,
	0xaf, 0x32, 0x12, 0x7e, 0x00, 0xf8, 0x4c, 0xdb, 0x1a, 0x9d, 0xb5, 0x75, 0xdb, 0xbe, 0x1d, 0x0e,
	0x0f, 0x56, 0x4b, 0x2f, 0x7e, 0x05, 0xd8, 0xbf, 0x35, 0x3e, 0x6f, 0x33, 0xff, 0x71, 0xde, 0xf0,
	0xd9, 0xff, 0x08, 0xa5, 0x4d, 0x42, 0xf6, 0x8a, 0xbd, 0xeb, 0x7d, 0xb7, 0x0d, 0x61, 0xe6, 0x98,
	0xcf, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0x6d, 0x46, 0xdd, 0x64, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionClient interface {
	NewSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionGrant, error)
	LookAround(ctx context.Context, opts ...grpc.CallOption) (Session_LookAroundClient, error)
}

type sessionClient struct {
	cc *grpc.ClientConn
}

func NewSessionClient(cc *grpc.ClientConn) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) NewSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*SessionGrant, error) {
	out := new(SessionGrant)
	err := c.cc.Invoke(ctx, "/gameproto.v1.Session/NewSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) LookAround(ctx context.Context, opts ...grpc.CallOption) (Session_LookAroundClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Session_serviceDesc.Streams[0], "/gameproto.v1.Session/LookAround", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionLookAroundClient{stream}
	return x, nil
}

type Session_LookAroundClient interface {
	Send(*LookAroundRequest) error
	Recv() (*LookAroundAnswer, error)
	grpc.ClientStream
}

type sessionLookAroundClient struct {
	grpc.ClientStream
}

func (x *sessionLookAroundClient) Send(m *LookAroundRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionLookAroundClient) Recv() (*LookAroundAnswer, error) {
	m := new(LookAroundAnswer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SessionServer is the server API for Session service.
type SessionServer interface {
	NewSession(context.Context, *SessionRequest) (*SessionGrant, error)
	LookAround(Session_LookAroundServer) error
}

// UnimplementedSessionServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (*UnimplementedSessionServer) NewSession(ctx context.Context, req *SessionRequest) (*SessionGrant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSession not implemented")
}
func (*UnimplementedSessionServer) LookAround(srv Session_LookAroundServer) error {
	return status.Errorf(codes.Unimplemented, "method LookAround not implemented")
}

func RegisterSessionServer(s *grpc.Server, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_NewSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).NewSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameproto.v1.Session/NewSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).NewSession(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_LookAround_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionServer).LookAround(&sessionLookAroundServer{stream})
}

type Session_LookAroundServer interface {
	Send(*LookAroundAnswer) error
	Recv() (*LookAroundRequest, error)
	grpc.ServerStream
}

type sessionLookAroundServer struct {
	grpc.ServerStream
}

func (x *sessionLookAroundServer) Send(m *LookAroundAnswer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionLookAroundServer) Recv() (*LookAroundRequest, error) {
	m := new(LookAroundRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gameproto.v1.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSession",
			Handler:    _Session_NewSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LookAround",
			Handler:       _Session_LookAround_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "game.proto",
}
