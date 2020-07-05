// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/putorders/putorders.proto

package go_micro_srv_PutOrders

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
	SessionId string `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	// 订单 id
	OrderId string `protobuf:"bytes,2,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	// 同意 or 拒绝
	Action               string   `protobuf:"bytes,3,opt,name=Action,proto3" json:"Action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_122dad0cbd606813, []int{0}
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

func (m *Request) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Request) GetAction() string {
	if m != nil {
		return m.Action
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
	return fileDescriptor_122dad0cbd606813, []int{1}
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
	proto.RegisterType((*Request)(nil), "go.micro.srv.PutOrders.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PutOrders.Response")
}

func init() { proto.RegisterFile("proto/putorders/putorders.proto", fileDescriptor_122dad0cbd606813) }

var fileDescriptor_122dad0cbd606813 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x28, 0x2d, 0xc9, 0x2f, 0x4a, 0x49, 0x2d, 0x2a, 0x46, 0xb0, 0xf4, 0xc0, 0x32,
	0x42, 0x62, 0xe9, 0xf9, 0x7a, 0xb9, 0x99, 0xc9, 0x45, 0xf9, 0x7a, 0xc5, 0x45, 0x65, 0x7a, 0x01,
	0xa5, 0x25, 0xfe, 0x60, 0x59, 0xa5, 0x48, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x19, 0x2e, 0xce, 0xe0, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0x3c, 0xcf, 0x14, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x84, 0x80, 0x90, 0x04, 0x17, 0x3b, 0x58, 0x8b, 0x67, 0x8a, 0x04, 0x13,
	0x58, 0x0e, 0xc6, 0x15, 0x12, 0xe3, 0x62, 0x73, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x06,
	0x4b, 0x40, 0x79, 0x4a, 0x16, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42,
	0x22, 0x5c, 0xac, 0xae, 0x45, 0x45, 0x7e, 0xf9, 0x50, 0x73, 0x21, 0x1c, 0x90, 0x4e, 0xd7, 0xa2,
	0x22, 0xdf, 0xe2, 0x74, 0xa8, 0x91, 0x50, 0x9e, 0x51, 0x2c, 0x17, 0x27, 0xdc, 0x85, 0x42, 0x01,
	0xc8, 0x1c, 0x79, 0x3d, 0xec, 0xfe, 0xd0, 0x83, 0x7a, 0x42, 0x4a, 0x01, 0xb7, 0x02, 0x88, 0x53,
	0x94, 0x18, 0x92, 0xd8, 0xc0, 0x41, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x53, 0x82, 0x83,
	0x3b, 0x35, 0x01, 0x00, 0x00,
}
