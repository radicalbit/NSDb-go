// Code generated by protoc-gen-go. DO NOT EDIT.
// source: migration.proto

package io_radicalbit_nsdb_rpc

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

type MigrateRequest struct {
	SourcePath           string   `protobuf:"bytes,1,opt,name=sourcePath,proto3" json:"sourcePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MigrateRequest) Reset()         { *m = MigrateRequest{} }
func (m *MigrateRequest) String() string { return proto.CompactTextString(m) }
func (*MigrateRequest) ProtoMessage()    {}
func (*MigrateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e5742f4f907425, []int{0}
}

func (m *MigrateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrateRequest.Unmarshal(m, b)
}
func (m *MigrateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrateRequest.Marshal(b, m, deterministic)
}
func (m *MigrateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrateRequest.Merge(m, src)
}
func (m *MigrateRequest) XXX_Size() int {
	return xxx_messageInfo_MigrateRequest.Size(m)
}
func (m *MigrateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MigrateRequest proto.InternalMessageInfo

func (m *MigrateRequest) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

type MigrateResponse struct {
	StartedSuccessfully  bool     `protobuf:"varint,1,opt,name=startedSuccessfully,proto3" json:"startedSuccessfully,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	SourcePath           string   `protobuf:"bytes,3,opt,name=sourcePath,proto3" json:"sourcePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MigrateResponse) Reset()         { *m = MigrateResponse{} }
func (m *MigrateResponse) String() string { return proto.CompactTextString(m) }
func (*MigrateResponse) ProtoMessage()    {}
func (*MigrateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e5742f4f907425, []int{1}
}

func (m *MigrateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrateResponse.Unmarshal(m, b)
}
func (m *MigrateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrateResponse.Marshal(b, m, deterministic)
}
func (m *MigrateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrateResponse.Merge(m, src)
}
func (m *MigrateResponse) XXX_Size() int {
	return xxx_messageInfo_MigrateResponse.Size(m)
}
func (m *MigrateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MigrateResponse proto.InternalMessageInfo

func (m *MigrateResponse) GetStartedSuccessfully() bool {
	if m != nil {
		return m.StartedSuccessfully
	}
	return false
}

func (m *MigrateResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func (m *MigrateResponse) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func init() {
	proto.RegisterType((*MigrateRequest)(nil), "io.radicalbit.nsdb.rpc.MigrateRequest")
	proto.RegisterType((*MigrateResponse)(nil), "io.radicalbit.nsdb.rpc.MigrateResponse")
}

func init() { proto.RegisterFile("migration.proto", fileDescriptor_30e5742f4f907425) }

var fileDescriptor_30e5742f4f907425 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcd, 0x4c, 0x2f,
	0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcb, 0xcc, 0xd7,
	0x2b, 0x4a, 0x4c, 0xc9, 0x4c, 0x4e, 0xcc, 0x49, 0xca, 0x2c, 0xd1, 0xcb, 0x2b, 0x4e, 0x49, 0xd2,
	0x2b, 0x2a, 0x48, 0x56, 0x32, 0xe0, 0xe2, 0xf3, 0x05, 0x2b, 0x4d, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2, 0x2a, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x0d, 0x48, 0x2c, 0xc9,
	0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x12, 0x51, 0xaa, 0xe7, 0xe2, 0x87, 0xeb, 0x28,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x32, 0xe0, 0x12, 0x2e, 0x2e, 0x49, 0x2c, 0x2a, 0x49, 0x4d,
	0x09, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0x4e, 0x2b, 0xcd, 0xc9, 0xa9, 0x04, 0xeb, 0xe5, 0x08,
	0xc2, 0x26, 0x25, 0x24, 0xc5, 0xc5, 0x91, 0x5a, 0x54, 0x94, 0x5f, 0xe4, 0x5b, 0x9c, 0x2e, 0xc1,
	0x04, 0xb6, 0x02, 0xce, 0x47, 0x73, 0x00, 0x33, 0xba, 0x03, 0x8c, 0xd2, 0xb9, 0x38, 0x7d, 0x61,
	0xbe, 0x13, 0x8a, 0xe2, 0x62, 0x87, 0xba, 0x46, 0x48, 0x4d, 0x0f, 0xbb, 0x1f, 0xf5, 0x50, 0x3d,
	0x28, 0xa5, 0x4e, 0x50, 0x1d, 0xc4, 0x5b, 0x49, 0x6c, 0xe0, 0xa0, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0xe4, 0x82, 0x57, 0x4d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MigrationClient is the client API for Migration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MigrationClient interface {
	Migrate(ctx context.Context, in *MigrateRequest, opts ...grpc.CallOption) (*MigrateResponse, error)
}

type migrationClient struct {
	cc *grpc.ClientConn
}

func NewMigrationClient(cc *grpc.ClientConn) MigrationClient {
	return &migrationClient{cc}
}

func (c *migrationClient) Migrate(ctx context.Context, in *MigrateRequest, opts ...grpc.CallOption) (*MigrateResponse, error) {
	out := new(MigrateResponse)
	err := c.cc.Invoke(ctx, "/io.radicalbit.nsdb.rpc.Migration/Migrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MigrationServer is the server API for Migration service.
type MigrationServer interface {
	Migrate(context.Context, *MigrateRequest) (*MigrateResponse, error)
}

// UnimplementedMigrationServer can be embedded to have forward compatible implementations.
type UnimplementedMigrationServer struct {
}

func (*UnimplementedMigrationServer) Migrate(ctx context.Context, req *MigrateRequest) (*MigrateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Migrate not implemented")
}

func RegisterMigrationServer(s *grpc.Server, srv MigrationServer) {
	s.RegisterService(&_Migration_serviceDesc, srv)
}

func _Migration_Migrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).Migrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.radicalbit.nsdb.rpc.Migration/Migrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).Migrate(ctx, req.(*MigrateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Migration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.radicalbit.nsdb.rpc.Migration",
	HandlerType: (*MigrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Migrate",
			Handler:    _Migration_Migrate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "migration.proto",
}