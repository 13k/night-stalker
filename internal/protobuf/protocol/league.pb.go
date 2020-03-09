// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol/league.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type League struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tier                 LeagueTier           `protobuf:"varint,3,opt,name=tier,proto3,enum=protocol.LeagueTier" json:"tier,omitempty"`
	Region               LeagueRegion         `protobuf:"varint,4,opt,name=region,proto3,enum=protocol.LeagueRegion" json:"region,omitempty"`
	Status               LeagueStatus         `protobuf:"varint,5,opt,name=status,proto3,enum=protocol.LeagueStatus" json:"status,omitempty"`
	TotalPrizePool       uint32               `protobuf:"varint,6,opt,name=total_prize_pool,json=totalPrizePool,proto3" json:"total_prize_pool,omitempty"`
	LastActivityAt       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=last_activity_at,json=lastActivityAt,proto3" json:"last_activity_at,omitempty"`
	StartAt              *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	FinishAt             *timestamp.Timestamp `protobuf:"bytes,9,opt,name=finish_at,json=finishAt,proto3" json:"finish_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *League) Reset()         { *m = League{} }
func (m *League) String() string { return proto.CompactTextString(m) }
func (*League) ProtoMessage()    {}
func (*League) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd5d5b2730e85bf3, []int{0}
}

func (m *League) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_League.Unmarshal(m, b)
}
func (m *League) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_League.Marshal(b, m, deterministic)
}
func (m *League) XXX_Merge(src proto.Message) {
	xxx_messageInfo_League.Merge(m, src)
}
func (m *League) XXX_Size() int {
	return xxx_messageInfo_League.Size(m)
}
func (m *League) XXX_DiscardUnknown() {
	xxx_messageInfo_League.DiscardUnknown(m)
}

var xxx_messageInfo_League proto.InternalMessageInfo

func (m *League) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *League) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *League) GetTier() LeagueTier {
	if m != nil {
		return m.Tier
	}
	return LeagueTier_LEAGUE_TIER_UNSET
}

func (m *League) GetRegion() LeagueRegion {
	if m != nil {
		return m.Region
	}
	return LeagueRegion_LEAGUE_REGION_UNSET
}

func (m *League) GetStatus() LeagueStatus {
	if m != nil {
		return m.Status
	}
	return LeagueStatus_LEAGUE_STATUS_UNSET
}

func (m *League) GetTotalPrizePool() uint32 {
	if m != nil {
		return m.TotalPrizePool
	}
	return 0
}

func (m *League) GetLastActivityAt() *timestamp.Timestamp {
	if m != nil {
		return m.LastActivityAt
	}
	return nil
}

func (m *League) GetStartAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartAt
	}
	return nil
}

func (m *League) GetFinishAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishAt
	}
	return nil
}

func init() {
	proto.RegisterType((*League)(nil), "protocol.League")
}

func init() { proto.RegisterFile("protocol/league.proto", fileDescriptor_cd5d5b2730e85bf3) }

var fileDescriptor_cd5d5b2730e85bf3 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xe9, 0xd6, 0xb7, 0xeb, 0xf2, 0x62, 0x19, 0x61, 0x4a, 0xd9, 0xc5, 0xe2, 0xa9, 0x17,
	0x53, 0xdc, 0x90, 0xe1, 0xb1, 0xe2, 0xd1, 0xc3, 0x88, 0x3b, 0x79, 0x29, 0xd9, 0x96, 0x75, 0x61,
	0x69, 0x52, 0x92, 0x7f, 0x05, 0xfd, 0x20, 0x7e, 0x5e, 0x69, 0xba, 0x2a, 0x88, 0xb0, 0x5b, 0xfa,
	0xfc, 0x7f, 0xbf, 0xe7, 0xf0, 0x14, 0x5d, 0xd6, 0x46, 0x83, 0xde, 0x6a, 0x99, 0x49, 0xce, 0xca,
	0x86, 0x13, 0xf7, 0x8d, 0xc3, 0x3e, 0x9e, 0x5d, 0x97, 0x5a, 0x97, 0x92, 0x67, 0x2e, 0xd8, 0x34,
	0xfb, 0x0c, 0x44, 0xc5, 0x2d, 0xb0, 0xaa, 0xee, 0xd0, 0xd9, 0xf4, 0xbb, 0x81, 0xab, 0xa6, 0xb2,
	0x5d, 0x7a, 0xf3, 0x39, 0x44, 0xc1, 0xb3, 0x6b, 0xc4, 0x11, 0x1a, 0x88, 0x5d, 0xec, 0x25, 0x5e,
	0xea, 0xd3, 0x81, 0xd8, 0x61, 0x8c, 0x7c, 0xc5, 0x2a, 0x1e, 0x0f, 0x12, 0x2f, 0x1d, 0x53, 0xf7,
	0xc6, 0x29, 0xf2, 0x41, 0x70, 0x13, 0x0f, 0x13, 0x2f, 0x8d, 0xe6, 0x53, 0xd2, 0x77, 0x92, 0xae,
	0x63, 0x2d, 0xb8, 0xa1, 0x8e, 0xc0, 0x04, 0x05, 0x86, 0x97, 0x42, 0xab, 0xd8, 0x77, 0xec, 0xd5,
	0x6f, 0x96, 0xba, 0x2b, 0x3d, 0x51, 0x2d, 0x6f, 0x81, 0x41, 0x63, 0xe3, 0x7f, 0x7f, 0xf3, 0x2f,
	0xee, 0x4a, 0x4f, 0x14, 0x4e, 0xd1, 0x04, 0x34, 0x30, 0x59, 0xd4, 0x46, 0x7c, 0xf0, 0xa2, 0xd6,
	0x5a, 0xc6, 0x41, 0xe2, 0xa5, 0x17, 0x34, 0x72, 0xf9, 0xaa, 0x8d, 0x57, 0x5a, 0x4b, 0xfc, 0x84,
	0x26, 0x92, 0x59, 0x28, 0xd8, 0x16, 0xc4, 0x9b, 0x80, 0xf7, 0x82, 0x41, 0x3c, 0x4a, 0xbc, 0xf4,
	0xff, 0x7c, 0x46, 0xba, 0xd1, 0x48, 0x3f, 0x1a, 0x59, 0xf7, 0xa3, 0xd1, 0xa8, 0x75, 0xf2, 0x93,
	0x92, 0x03, 0xbe, 0x47, 0xa1, 0x05, 0x66, 0xa0, 0xb5, 0xc3, 0xb3, 0xf6, 0xc8, 0xb1, 0x39, 0xe0,
	0x25, 0x1a, 0xef, 0x85, 0x12, 0xf6, 0xd0, 0x7a, 0xe3, 0xb3, 0x5e, 0xd8, 0xc1, 0x39, 0x3c, 0x3e,
	0xbc, 0x2e, 0x4b, 0x01, 0x87, 0x66, 0x43, 0xb6, 0xba, 0xca, 0xee, 0x16, 0xc7, 0x4c, 0x89, 0xf2,
	0x00, 0xb7, 0x16, 0x98, 0x3c, 0x72, 0x93, 0x09, 0x05, 0xdc, 0x28, 0x26, 0x7f, 0x7e, 0x78, 0xbf,
	0xd7, 0x26, 0x70, 0xaf, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x4f, 0x76, 0x3f, 0x34,
	0x02, 0x00, 0x00,
}