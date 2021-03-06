// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/simple.proto

package proto

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

type ParseMessageRequest struct {
	TenantId             string   `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	CloudfsId            string   `protobuf:"bytes,2,opt,name=cloudfs_id,json=cloudfsId,proto3" json:"cloudfs_id,omitempty"`
	NodeId               string   `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseMessageRequest) Reset()         { *m = ParseMessageRequest{} }
func (m *ParseMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ParseMessageRequest) ProtoMessage()    {}
func (*ParseMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4285f8f75e6ba5cf, []int{0}
}

func (m *ParseMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseMessageRequest.Unmarshal(m, b)
}
func (m *ParseMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseMessageRequest.Marshal(b, m, deterministic)
}
func (m *ParseMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseMessageRequest.Merge(m, src)
}
func (m *ParseMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ParseMessageRequest.Size(m)
}
func (m *ParseMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseMessageRequest proto.InternalMessageInfo

func (m *ParseMessageRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *ParseMessageRequest) GetCloudfsId() string {
	if m != nil {
		return m.CloudfsId
	}
	return ""
}

func (m *ParseMessageRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ParseMessageRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ParseMessageResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseMessageResponse) Reset()         { *m = ParseMessageResponse{} }
func (m *ParseMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ParseMessageResponse) ProtoMessage()    {}
func (*ParseMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4285f8f75e6ba5cf, []int{1}
}

func (m *ParseMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseMessageResponse.Unmarshal(m, b)
}
func (m *ParseMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseMessageResponse.Marshal(b, m, deterministic)
}
func (m *ParseMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseMessageResponse.Merge(m, src)
}
func (m *ParseMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ParseMessageResponse.Size(m)
}
func (m *ParseMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParseMessageResponse proto.InternalMessageInfo

func (m *ParseMessageResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ParseMessageResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*ParseMessageRequest)(nil), "grpc_simple.ParseMessageRequest")
	proto.RegisterType((*ParseMessageResponse)(nil), "grpc_simple.ParseMessageResponse")
}

func init() { proto.RegisterFile("proto/simple.proto", fileDescriptor_4285f8f75e6ba5cf) }

var fileDescriptor_4285f8f75e6ba5cf = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x7e, 0xf3, 0x5a, 0x63, 0x33, 0xde, 0x46, 0xd1, 0x52, 0x11, 0xda, 0x9c, 0xbc, 0x98, 0x80,
	0xfe, 0x02, 0x7b, 0x91, 0x08, 0x42, 0x49, 0x3d, 0x09, 0x52, 0xd6, 0xec, 0x18, 0x02, 0x35, 0xbb,
	0xee, 0x6c, 0xc0, 0x93, 0xbf, 0x5d, 0x32, 0x59, 0x25, 0x82, 0x78, 0xda, 0x79, 0xbe, 0xe6, 0x61,
	0x16, 0xd0, 0x3a, 0xe3, 0x4d, 0xce, 0xcd, 0xab, 0xdd, 0x51, 0x26, 0x00, 0x0f, 0x6b, 0x67, 0xab,
	0xed, 0x40, 0xa5, 0x1f, 0x70, 0xb4, 0x56, 0x8e, 0xe9, 0x9e, 0x98, 0x55, 0x4d, 0x25, 0xbd, 0x75,
	0xc4, 0x1e, 0xcf, 0x20, 0xf1, 0xd4, 0xaa, 0xd6, 0x6f, 0x1b, 0x3d, 0x8b, 0x16, 0xd1, 0x45, 0x52,
	0x4e, 0x07, 0xa2, 0xd0, 0x78, 0x0e, 0x50, 0xed, 0x4c, 0xa7, 0x5f, 0xb8, 0x57, 0xff, 0x8b, 0x9a,
	0x04, 0xa6, 0xd0, 0x78, 0x0a, 0x07, 0xad, 0xd1, 0xd4, 0x6b, 0x7b, 0xa2, 0xc5, 0x3d, 0x2c, 0x34,
	0x22, 0x4c, 0xb4, 0xf2, 0x6a, 0x36, 0x11, 0x56, 0xe6, 0xf4, 0x0e, 0x8e, 0x7f, 0xf6, 0xb3, 0x35,
	0x2d, 0x13, 0x9e, 0x40, 0xcc, 0x5e, 0xf9, 0x8e, 0xa5, 0x7d, 0xbf, 0x0c, 0x08, 0xe7, 0x30, 0x75,
	0xc1, 0x13, 0x9a, 0xbf, 0xf1, 0xd5, 0x13, 0xc4, 0xb2, 0xcb, 0xe1, 0x06, 0x40, 0xa6, 0x9b, 0x4e,
	0x37, 0x1e, 0x17, 0xd9, 0xe8, 0xe2, 0xec, 0x97, 0x73, 0xe7, 0xcb, 0x3f, 0x1c, 0xc3, 0xf2, 0xf4,
	0xdf, 0xea, 0x01, 0x96, 0xb5, 0x72, 0x95, 0x6a, 0x33, 0x7a, 0x57, 0xbd, 0x91, 0x25, 0x75, 0x39,
	0xfe, 0xdc, 0x15, 0xdc, 0x3a, 0x5b, 0x6d, 0x84, 0x59, 0x47, 0x8f, 0x21, 0x90, 0x7f, 0x05, 0xf2,
	0x51, 0x20, 0x97, 0xc0, 0x73, 0x2c, 0xcf, 0xf5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x1d,
	0x0d, 0x96, 0xaa, 0x01, 0x00, 0x00,
}
