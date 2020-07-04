// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/posthousesimage/posthousesimage.proto

package go_micro_srv_PostHousesImage

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
	// 房屋 Id
	HouseId string `protobuf:"bytes,1,opt,name=HouseId,proto3" json:"HouseId,omitempty"`
	// 图片名字
	FileName string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`
	// 图片
	Image                []byte   `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba15f6ee9286afde, []int{0}
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

func (m *Request) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

func (m *Request) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Request) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type Response struct {
	ErrNo                string   `protobuf:"bytes,1,opt,name=ErrNo,proto3" json:"ErrNo,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=Url,proto3" json:"Url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba15f6ee9286afde, []int{1}
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

func (m *Response) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostHousesImage.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostHousesImage.Response")
}

func init() {
	proto.RegisterFile("proto/posthousesimage/posthousesimage.proto", fileDescriptor_ba15f6ee9286afde)
}

var fileDescriptor_ba15f6ee9286afde = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0x2f, 0x2e, 0xc9, 0xc8, 0x2f, 0x2d, 0x4e, 0x2d, 0xce, 0xcc, 0x4d, 0x4c,
	0x4f, 0x45, 0xe7, 0xeb, 0x81, 0x55, 0x09, 0xc9, 0xa4, 0xe7, 0xeb, 0xe5, 0x66, 0x26, 0x17, 0xe5,
	0xeb, 0x15, 0x17, 0x95, 0xe9, 0x05, 0xe4, 0x17, 0x97, 0x78, 0x80, 0xd5, 0x78, 0x82, 0xd4, 0x28,
	0x85, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0x83, 0x65,
	0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x29, 0x2e, 0x0e, 0xb7,
	0xcc, 0x9c, 0x54, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x14, 0x9c, 0x2f, 0x24, 0xc2, 0xc5,
	0x0a, 0x36, 0x49, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc2, 0x51, 0xf2, 0xe2, 0xe2, 0x08,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x06, 0xab, 0x70, 0x2d, 0x2a, 0xf2, 0xcb, 0x87, 0x9a, 0x0a,
	0xe1, 0x08, 0x89, 0x71, 0xb1, 0xb9, 0x16, 0x15, 0xf9, 0x16, 0xa7, 0x43, 0x4d, 0x84, 0xf2, 0x84,
	0x04, 0xb8, 0x98, 0x43, 0x8b, 0x72, 0xc0, 0xa6, 0x71, 0x06, 0x81, 0x98, 0x46, 0xa5, 0x5c, 0xfc,
	0x68, 0xae, 0x16, 0x4a, 0xc2, 0x14, 0x52, 0xd5, 0xc3, 0xe7, 0x4f, 0x3d, 0xa8, 0x27, 0xa5, 0xd4,
	0x08, 0x29, 0x83, 0x38, 0x5a, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x7c, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x74, 0xb9, 0xe1, 0x4b, 0x6d, 0x01, 0x00, 0x00,
}