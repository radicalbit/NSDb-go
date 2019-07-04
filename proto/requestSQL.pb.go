// Code generated by protoc-gen-go. DO NOT EDIT.
// source: requestSQL.proto

package io_radicalbit_nsdb_rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SQLRequestStatement struct {
	Db                   string   `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Statement            string   `protobuf:"bytes,3,opt,name=statement,proto3" json:"statement,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SQLRequestStatement) Reset()         { *m = SQLRequestStatement{} }
func (m *SQLRequestStatement) String() string { return proto.CompactTextString(m) }
func (*SQLRequestStatement) ProtoMessage()    {}
func (*SQLRequestStatement) Descriptor() ([]byte, []int) {
	return fileDescriptor_6afc6e5ac5b5e42e, []int{0}
}

func (m *SQLRequestStatement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SQLRequestStatement.Unmarshal(m, b)
}
func (m *SQLRequestStatement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SQLRequestStatement.Marshal(b, m, deterministic)
}
func (m *SQLRequestStatement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SQLRequestStatement.Merge(m, src)
}
func (m *SQLRequestStatement) XXX_Size() int {
	return xxx_messageInfo_SQLRequestStatement.Size(m)
}
func (m *SQLRequestStatement) XXX_DiscardUnknown() {
	xxx_messageInfo_SQLRequestStatement.DiscardUnknown(m)
}

var xxx_messageInfo_SQLRequestStatement proto.InternalMessageInfo

func (m *SQLRequestStatement) GetDb() string {
	if m != nil {
		return m.Db
	}
	return ""
}

func (m *SQLRequestStatement) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SQLRequestStatement) GetStatement() string {
	if m != nil {
		return m.Statement
	}
	return ""
}

func init() {
	proto.RegisterType((*SQLRequestStatement)(nil), "io.radicalbit.nsdb.rpc.SQLRequestStatement")
}

func init() { proto.RegisterFile("requestSQL.proto", fileDescriptor_6afc6e5ac5b5e42e) }

var fileDescriptor_6afc6e5ac5b5e42e = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x09, 0x0e, 0xf4, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcb, 0xcc,
	0xd7, 0x2b, 0x4a, 0x4c, 0xc9, 0x4c, 0x4e, 0xcc, 0x49, 0xca, 0x2c, 0xd1, 0xcb, 0x2b, 0x4e, 0x49,
	0xd2, 0x2b, 0x2a, 0x48, 0x56, 0x4a, 0xe4, 0x12, 0x0e, 0x0e, 0xf4, 0x09, 0x82, 0x2a, 0x2f, 0x49,
	0x2c, 0x49, 0xcd, 0x4d, 0xcd, 0x2b, 0x11, 0xe2, 0xe3, 0x62, 0x4a, 0x49, 0x92, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x62, 0x4a, 0x49, 0x12, 0x92, 0xe1, 0xe2, 0xcc, 0x4b, 0xcc, 0x4d, 0x2d, 0x2e,
	0x48, 0x4c, 0x4e, 0x95, 0x60, 0x02, 0x0b, 0x23, 0x04, 0x40, 0xb2, 0xc5, 0x30, 0xad, 0x12, 0xcc,
	0x10, 0x59, 0xb8, 0x40, 0x12, 0x1b, 0xd8, 0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32,
	0xe6, 0x71, 0xfb, 0x95, 0x00, 0x00, 0x00,
}
