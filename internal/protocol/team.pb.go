// Code generated by protoc-gen-go. DO NOT EDIT.
// source: team.proto

package protocol

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

type Team struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	LogoUrl              string   `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b4e9e93d7b2c6bb, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Team) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Team)(nil), "protocol.Team")
}

func init() { proto.RegisterFile("team.proto", fileDescriptor_8b4e9e93d7b2c6bb) }

var fileDescriptor_8b4e9e93d7b2c6bb = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x4d, 0xcc,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0xe1, 0x5c,
	0x2c, 0x21, 0xa9, 0x89, 0xb9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x49, 0x62, 0xba, 0x04, 0x33, 0x58,
	0x08, 0xc4, 0x14, 0x92, 0xe4, 0xe2, 0xc8, 0xc9, 0x4f, 0xcf, 0x8f, 0x2f, 0x2d, 0xca, 0x91, 0x60,
	0x01, 0x0b, 0xb3, 0x83, 0xf8, 0xa1, 0x45, 0x39, 0x49, 0x6c, 0x60, 0x2b, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0xae, 0x7b, 0xda, 0x77, 0x00, 0x00, 0x00,
}