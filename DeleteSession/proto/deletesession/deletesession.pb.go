// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/deletesession/deletesession.proto

package go_micro_srv_DeleteSession

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a8dbf7dd3f41f61, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type Response struct {
	ErrNo                string   `protobuf:"bytes,1,opt,name=ErrNo,proto3" json:"ErrNo,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a8dbf7dd3f41f61, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrNo() string {
	if m != nil {
		return m.ErrNo
	}
	return ""
}

func (m *Response) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.DeleteSession.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.DeleteSession.Response")
}

func init() {
	proto.RegisterFile("proto/deletesession/deletesession.proto", fileDescriptor_5a8dbf7dd3f41f61)
}

var fileDescriptor_5a8dbf7dd3f41f61 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0x2d, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x43,
	0xe5, 0xe9, 0x81, 0x55, 0x08, 0x49, 0xa5, 0xe7, 0xeb, 0xe5, 0x66, 0x26, 0x17, 0xe5, 0xeb, 0x15,
	0x17, 0x95, 0xe9, 0xb9, 0x80, 0x55, 0x04, 0x43, 0x54, 0x28, 0xa9, 0x73, 0xb1, 0x07, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x70, 0x71, 0x42, 0x45, 0x3d, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x10, 0x02, 0x4a, 0x16, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0x22, 0x5c, 0xac, 0xae, 0x45, 0x45, 0x7e, 0xf9, 0x50, 0x55, 0x10, 0x8e, 0x90, 0x18,
	0x17, 0x9b, 0x6b, 0x51, 0x91, 0x6f, 0x71, 0xba, 0x04, 0x13, 0x58, 0x18, 0xca, 0x33, 0xca, 0xe5,
	0xe2, 0x45, 0xb1, 0x53, 0x28, 0x06, 0x5d, 0x40, 0x59, 0x0f, 0xb7, 0x0b, 0xf5, 0xa0, 0xce, 0x93,
	0x52, 0xc1, 0xaf, 0x08, 0xe2, 0x34, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xa7, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x84, 0xfb, 0x6f, 0x85, 0x1f, 0x01, 0x00, 0x00,
}
