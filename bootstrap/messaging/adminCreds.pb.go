// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adminCreds.proto

package messaging

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_adminCreds_fa5a5cf8c48a22ae, []int{0}
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminCreds_fa5a5cf8c48a22ae, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (dst *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(dst, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=messaging.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_adminCreds_fa5a5cf8c48a22ae, []int{1}
}
func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (dst *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(dst, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterType((*Chunk)(nil), "messaging.Chunk")
	proto.RegisterType((*UploadStatus)(nil), "messaging.UploadStatus")
	proto.RegisterEnum("messaging.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
}

func init() { proto.RegisterFile("adminCreds.proto", fileDescriptor_adminCreds_fa5a5cf8c48a22ae) }

var fileDescriptor_adminCreds_fa5a5cf8c48a22ae = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4c, 0xc9, 0xcd,
	0xcc, 0x73, 0x2e, 0x4a, 0x4d, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0xcf, 0xcc, 0x4b, 0x57, 0x52, 0xe4, 0x62, 0x75, 0xce, 0x28, 0xcd, 0xcb,
	0x16, 0x92, 0xe0, 0x62, 0x77, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x09, 0x82, 0x71, 0x95, 0x22, 0xb9, 0x78, 0x42, 0x0b, 0x72, 0xf2, 0x13, 0x53, 0x82, 0x4b,
	0x12, 0x4b, 0x4a, 0x8b, 0x41, 0x2a, 0x7d, 0xc1, 0xfa, 0x53, 0xc1, 0x2a, 0x39, 0x83, 0x60, 0x5c,
	0x21, 0x7d, 0x2e, 0x16, 0xe7, 0xfc, 0x94, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x69,
	0x3d, 0xb8, 0x35, 0x7a, 0xc8, 0x06, 0x80, 0x94, 0x04, 0x81, 0x15, 0x6a, 0x19, 0x73, 0x09, 0xa0,
	0xcb, 0x08, 0x71, 0x73, 0xb1, 0x87, 0xe6, 0x65, 0xe7, 0xe5, 0x97, 0xe7, 0x09, 0x30, 0x08, 0xb1,
	0x71, 0x31, 0xf9, 0x67, 0x0b, 0x30, 0x0a, 0x71, 0x71, 0xb1, 0xb9, 0x25, 0x66, 0xe6, 0xa4, 0xa6,
	0x08, 0x30, 0x19, 0x79, 0x72, 0xf1, 0xb9, 0x97, 0x42, 0x74, 0xa5, 0x16, 0x95, 0x65, 0x26, 0xa7,
	0x0a, 0x99, 0x73, 0xb1, 0x41, 0x8c, 0x11, 0x12, 0x40, 0xb2, 0x13, 0xec, 0x2f, 0x29, 0x71, 0x1c,
	0xae, 0x50, 0x62, 0xd0, 0x60, 0x4c, 0x62, 0x03, 0x87, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0x04, 0xd4, 0x79, 0x23, 0x01, 0x00, 0x00,
}
