// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package io_radicalbit_nsdb_rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

type RPCInsert struct {
	Database  string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Metric    string `protobuf:"bytes,3,opt,name=metric,proto3" json:"metric,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*RPCInsert_DecimalValue
	//	*RPCInsert_LongValue
	Value                isRPCInsert_Value     `protobuf_oneof:"value"`
	Dimensions           map[string]*Dimension `protobuf:"bytes,7,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags                 map[string]*Tag       `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RPCInsert) Reset()         { *m = RPCInsert{} }
func (m *RPCInsert) String() string { return proto.CompactTextString(m) }
func (*RPCInsert) ProtoMessage()    {}
func (*RPCInsert) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *RPCInsert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCInsert.Unmarshal(m, b)
}
func (m *RPCInsert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCInsert.Marshal(b, m, deterministic)
}
func (m *RPCInsert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCInsert.Merge(m, src)
}
func (m *RPCInsert) XXX_Size() int {
	return xxx_messageInfo_RPCInsert.Size(m)
}
func (m *RPCInsert) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCInsert.DiscardUnknown(m)
}

var xxx_messageInfo_RPCInsert proto.InternalMessageInfo

func (m *RPCInsert) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *RPCInsert) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *RPCInsert) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *RPCInsert) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type isRPCInsert_Value interface {
	isRPCInsert_Value()
}

type RPCInsert_DecimalValue struct {
	DecimalValue float64 `protobuf:"fixed64,5,opt,name=decimalValue,proto3,oneof"`
}

type RPCInsert_LongValue struct {
	LongValue int64 `protobuf:"varint,6,opt,name=longValue,proto3,oneof"`
}

func (*RPCInsert_DecimalValue) isRPCInsert_Value() {}

func (*RPCInsert_LongValue) isRPCInsert_Value() {}

func (m *RPCInsert) GetValue() isRPCInsert_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RPCInsert) GetDecimalValue() float64 {
	if x, ok := m.GetValue().(*RPCInsert_DecimalValue); ok {
		return x.DecimalValue
	}
	return 0
}

func (m *RPCInsert) GetLongValue() int64 {
	if x, ok := m.GetValue().(*RPCInsert_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (m *RPCInsert) GetDimensions() map[string]*Dimension {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *RPCInsert) GetTags() map[string]*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RPCInsert) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RPCInsert_DecimalValue)(nil),
		(*RPCInsert_LongValue)(nil),
	}
}

func init() {
	proto.RegisterType((*RPCInsert)(nil), "io.radicalbit.nsdb.rpc.RPCInsert")
	proto.RegisterMapType((map[string]*Dimension)(nil), "io.radicalbit.nsdb.rpc.RPCInsert.DimensionsEntry")
	proto.RegisterMapType((map[string]*Tag)(nil), "io.radicalbit.nsdb.rpc.RPCInsert.TagsEntry")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x9b, 0xa6, 0x6f, 0x99, 0xf6, 0xcf, 0x5f, 0xf6, 0x50, 0x62, 0x14, 0x89, 0xe2, 0x21,
	0x20, 0xa4, 0xb4, 0x1e, 0x14, 0x2f, 0x82, 0x2f, 0x50, 0x6f, 0x1a, 0x8a, 0x67, 0x27, 0xc9, 0x1a,
	0x16, 0xb3, 0xbb, 0x31, 0xbb, 0x15, 0xfa, 0x45, 0xfc, 0xbc, 0x92, 0x17, 0xd2, 0x56, 0x5a, 0xbc,
	0xed, 0x3c, 0x33, 0xcf, 0x6f, 0x98, 0x87, 0x85, 0x7f, 0x39, 0xfd, 0x5c, 0x52, 0xa5, 0xfd, 0x2c,
	0x97, 0x5a, 0x92, 0x31, 0x93, 0x7e, 0x8e, 0x31, 0x8b, 0x30, 0x0d, 0x99, 0xf6, 0x85, 0x8a, 0x43,
	0x3f, 0xcf, 0x22, 0xe7, 0x30, 0x91, 0x32, 0x49, 0xe9, 0xa4, 0x9c, 0x0a, 0x97, 0xef, 0x13, 0x14,
	0xab, 0xca, 0xe2, 0x8c, 0x22, 0xc9, 0xb9, 0x14, 0x55, 0x75, 0xf6, 0xdd, 0x01, 0x2b, 0x78, 0xbe,
	0x7f, 0x12, 0x8a, 0xe6, 0x9a, 0x38, 0x30, 0x88, 0x51, 0x63, 0x88, 0x8a, 0xda, 0x86, 0x6b, 0x78,
	0x56, 0xd0, 0xd4, 0xe4, 0x18, 0x2c, 0x81, 0x9c, 0xaa, 0x0c, 0x23, 0x6a, 0xb7, 0xcb, 0xe6, 0x5a,
	0x20, 0x63, 0xe8, 0x71, 0xaa, 0x73, 0x16, 0xd9, 0x66, 0xd9, 0xaa, 0xab, 0xc2, 0xa5, 0x19, 0xa7,
	0x4a, 0x23, 0xcf, 0xec, 0x8e, 0x6b, 0x78, 0x66, 0xb0, 0x16, 0xc8, 0x39, 0x8c, 0x62, 0x1a, 0x31,
	0x8e, 0xe9, 0x2b, 0xa6, 0x4b, 0x6a, 0x77, 0x5d, 0xc3, 0x33, 0xe6, 0xad, 0x60, 0x4b, 0x25, 0x27,
	0x60, 0xa5, 0x52, 0x24, 0xd5, 0x48, 0xaf, 0x60, 0xcc, 0x5b, 0xc1, 0x5a, 0x22, 0x2f, 0x00, 0x31,
	0xe3, 0x54, 0x28, 0x26, 0x85, 0xb2, 0xfb, 0xae, 0xe9, 0x0d, 0x67, 0x53, 0x7f, 0x77, 0x32, 0x7e,
	0x73, 0xac, 0xff, 0xd0, 0x78, 0x1e, 0x85, 0xce, 0x57, 0xc1, 0x06, 0x84, 0xdc, 0x42, 0x47, 0x63,
	0xa2, 0xec, 0x41, 0x09, 0xbb, 0xf8, 0x1b, 0xb6, 0xc0, 0xa4, 0xc6, 0x94, 0x46, 0xe7, 0x0d, 0xfe,
	0xff, 0xe2, 0x93, 0x03, 0x30, 0x3f, 0xe8, 0xaa, 0xce, 0xb5, 0x78, 0x92, 0x2b, 0xe8, 0x7e, 0x95,
	0x47, 0x15, 0x71, 0x0e, 0x67, 0xa7, 0xfb, 0xd6, 0x34, 0xa4, 0xa0, 0x9a, 0xbf, 0x69, 0x5f, 0x1b,
	0xce, 0x02, 0xac, 0x66, 0xe9, 0x0e, 0xf6, 0x74, 0x9b, 0x7d, 0xb4, 0x8f, 0xbd, 0xc0, 0x64, 0x83,
	0x7a, 0xd7, 0xaf, 0x6d, 0x61, 0xaf, 0xfc, 0x1f, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdc,
	0xed, 0x57, 0x0c, 0x71, 0x02, 0x00, 0x00,
}
