// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

/*
Package testing is a generated protocol buffer package.

It is generated from these files:
	grpc/grpc.proto

It has these top-level messages:
	SimpleRequest
	SimpleResponse
	StreamMsg
	StreamMsg2
*/
package testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SimpleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()                    { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string            { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()               {}
func (*SimpleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
func (m *SimpleRequest) Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (dst *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(dst, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

type SimpleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }
func (m *SimpleResponse) Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (dst *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(dst, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

type StreamMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMsg) Reset()                    { *m = StreamMsg{} }
func (m *StreamMsg) String() string            { return proto.CompactTextString(m) }
func (*StreamMsg) ProtoMessage()               {}
func (*StreamMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }
func (m *StreamMsg) Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg.Unmarshal(m, b)
}
func (m *StreamMsg) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg.Marshal(b, m, deterministic)
}
func (dst *StreamMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg.Merge(dst, src)
}
func (m *StreamMsg) XXX_Size() int {
	return xxx_messageInfo_StreamMsg.Size(m)
}
func (m *StreamMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg proto.InternalMessageInfo

type StreamMsg2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMsg2) Reset()                    { *m = StreamMsg2{} }
func (m *StreamMsg2) String() string            { return proto.CompactTextString(m) }
func (*StreamMsg2) ProtoMessage()               {}
func (*StreamMsg2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }
func (m *StreamMsg2) Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg2.Unmarshal(m, b)
}
func (m *StreamMsg2) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg2.Marshal(b, m, deterministic)
}
func (dst *StreamMsg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg2.Merge(dst, src)
}
func (m *StreamMsg2) XXX_Size() int {
	return xxx_messageInfo_StreamMsg2.Size(m)
}
func (m *StreamMsg2) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg2.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
	proto.RegisterType((*StreamMsg)(nil), "grpc.testing.StreamMsg")
	proto.RegisterType((*StreamMsg2)(nil), "grpc.testing.StreamMsg2")
}

func init() { proto.RegisterFile("grpc/grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2f, 0x2a, 0x48,
	0xd6, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x3c, 0x60, 0x76, 0x49, 0x6a, 0x71,
	0x49, 0x66, 0x5e, 0xba, 0x12, 0x3f, 0x17, 0x6f, 0x70, 0x66, 0x6e, 0x41, 0x4e, 0x6a, 0x50, 0x6a,
	0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x00, 0x17, 0x1f, 0x4c, 0xa0, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x55, 0x89, 0x9b, 0x8b, 0x33, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0xd7, 0xb7, 0x38, 0x5d, 0x89, 0x87,
	0x8b, 0x0b, 0xce, 0x31, 0x32, 0x9a, 0xc1, 0xc4, 0xc5, 0x12, 0x92, 0x5a, 0x5c, 0x22, 0xe4, 0xc6,
	0xc5, 0x19, 0x9a, 0x97, 0x58, 0x54, 0xe9, 0x9c, 0x98, 0x93, 0x23, 0x24, 0xad, 0x87, 0x6c, 0x85,
	0x1e, 0x8a, 0xf9, 0x52, 0x32, 0xd8, 0x25, 0x21, 0x76, 0x09, 0xb9, 0x70, 0x71, 0xb9, 0xe4, 0x97,
	0xe7, 0x15, 0x83, 0xad, 0xc0, 0x6f, 0x90, 0x38, 0x9a, 0x24, 0xcc, 0x55, 0x06, 0x8c, 0x42, 0xce,
	0x5c, 0x1c, 0xa1, 0x05, 0x50, 0x33, 0x70, 0x29, 0xc3, 0xef, 0x10, 0x0d, 0x46, 0x21, 0x5b, 0x2e,
	0x16, 0xa7, 0xcc, 0x94, 0x4c, 0xdc, 0x06, 0x48, 0xe0, 0x90, 0x30, 0xd2, 0x60, 0x34, 0x60, 0x74,
	0x72, 0x88, 0xb2, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x07, 0xc7, 0x40, 0x52, 0x69, 0x1a, 0x84, 0x91, 0xac, 0x9b, 0x9e,
	0x9a, 0xa7, 0x9b, 0x9e, 0xaf, 0x0f, 0x32, 0x22, 0x25, 0xb1, 0x24, 0x11, 0x1c, 0x4d, 0xd6, 0x50,
	0x03, 0x93, 0xd8, 0xc0, 0x8a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0xb9, 0x95, 0x42,
	0xc2, 0x01, 0x00, 0x00,
}
