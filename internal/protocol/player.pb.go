// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/protocol/player.proto

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

type Player struct {
	AccountId            uint32          `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PersonaName          string          `protobuf:"bytes,3,opt,name=persona_name,json=personaName,proto3" json:"persona_name,omitempty"`
	AvatarUrl            string          `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	AvatarMediumUrl      string          `protobuf:"bytes,5,opt,name=avatar_medium_url,json=avatarMediumUrl,proto3" json:"avatar_medium_url,omitempty"`
	AvatarFullUrl        string          `protobuf:"bytes,6,opt,name=avatar_full_url,json=avatarFullUrl,proto3" json:"avatar_full_url,omitempty"`
	IsPro                bool            `protobuf:"varint,7,opt,name=is_pro,json=isPro,proto3" json:"is_pro,omitempty"`
	Team                 *Player_Team    `protobuf:"bytes,100,opt,name=team,proto3" json:"team,omitempty"`
	Matches              []*Player_Match `protobuf:"bytes,101,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ca890c379a1848, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetAccountId() uint32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetPersonaName() string {
	if m != nil {
		return m.PersonaName
	}
	return ""
}

func (m *Player) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *Player) GetAvatarMediumUrl() string {
	if m != nil {
		return m.AvatarMediumUrl
	}
	return ""
}

func (m *Player) GetAvatarFullUrl() string {
	if m != nil {
		return m.AvatarFullUrl
	}
	return ""
}

func (m *Player) GetIsPro() bool {
	if m != nil {
		return m.IsPro
	}
	return false
}

func (m *Player) GetTeam() *Player_Team {
	if m != nil {
		return m.Team
	}
	return nil
}

func (m *Player) GetMatches() []*Player_Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

type Player_Match struct {
	MatchId              uint64   `protobuf:"varint,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	HeroId               uint64   `protobuf:"varint,2,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player_Match) Reset()         { *m = Player_Match{} }
func (m *Player_Match) String() string { return proto.CompactTextString(m) }
func (*Player_Match) ProtoMessage()    {}
func (*Player_Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ca890c379a1848, []int{0, 0}
}

func (m *Player_Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player_Match.Unmarshal(m, b)
}
func (m *Player_Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player_Match.Marshal(b, m, deterministic)
}
func (m *Player_Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player_Match.Merge(m, src)
}
func (m *Player_Match) XXX_Size() int {
	return xxx_messageInfo_Player_Match.Size(m)
}
func (m *Player_Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Player_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Player_Match proto.InternalMessageInfo

func (m *Player_Match) GetMatchId() uint64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *Player_Match) GetHeroId() uint64 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

type Player_Team struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	LogoUrl              string   `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player_Team) Reset()         { *m = Player_Team{} }
func (m *Player_Team) String() string { return proto.CompactTextString(m) }
func (*Player_Team) ProtoMessage()    {}
func (*Player_Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_79ca890c379a1848, []int{0, 1}
}

func (m *Player_Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player_Team.Unmarshal(m, b)
}
func (m *Player_Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player_Team.Marshal(b, m, deterministic)
}
func (m *Player_Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player_Team.Merge(m, src)
}
func (m *Player_Team) XXX_Size() int {
	return xxx_messageInfo_Player_Team.Size(m)
}
func (m *Player_Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Player_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Player_Team proto.InternalMessageInfo

func (m *Player_Team) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player_Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player_Team) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Player_Team) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Player)(nil), "protocol.Player")
	proto.RegisterType((*Player_Match)(nil), "protocol.Player.Match")
	proto.RegisterType((*Player_Team)(nil), "protocol.Player.Team")
}

func init() { proto.RegisterFile("internal/protocol/player.proto", fileDescriptor_79ca890c379a1848) }

var fileDescriptor_79ca890c379a1848 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x86, 0xc3, 0x47, 0xa1, 0x9d, 0x5a, 0x3f, 0x36, 0xa9, 0x62, 0x13, 0x0d, 0x7a, 0x30, 0xe8,
	0x81, 0x9a, 0x7a, 0xf4, 0x6e, 0xc2, 0xa1, 0xa6, 0x21, 0x36, 0x1e, 0xc9, 0x0a, 0x6b, 0x4b, 0xb2,
	0xb0, 0x64, 0x59, 0x4c, 0xfc, 0xad, 0xfe, 0x19, 0xb3, 0xc3, 0x62, 0x4c, 0xf4, 0x36, 0xf3, 0xbc,
	0x0f, 0x0c, 0xbc, 0x70, 0x59, 0xd6, 0x8a, 0xc9, 0x9a, 0xf2, 0x65, 0x23, 0x85, 0x12, 0xb9, 0xe0,
	0xcb, 0x86, 0xd3, 0x4f, 0x26, 0x63, 0xdc, 0xc9, 0x78, 0xc0, 0xd7, 0x5f, 0x0e, 0x78, 0x1b, 0x8c,
	0xc8, 0x05, 0x00, 0xcd, 0x73, 0xd1, 0xd5, 0x2a, 0x2b, 0x8b, 0xc0, 0x0a, 0xad, 0x68, 0x96, 0x4e,
	0x0c, 0x49, 0x0a, 0x42, 0xc0, 0xad, 0x69, 0xc5, 0x02, 0x3b, 0xb4, 0xa2, 0x49, 0x8a, 0x33, 0xb9,
	0x82, 0x83, 0x86, 0xc9, 0x56, 0xd4, 0x34, 0xc3, 0xcc, 0xc1, 0x6c, 0x6a, 0xd8, 0xb3, 0x56, 0xf4,
	0x5b, 0x3f, 0xa8, 0xa2, 0x32, 0xeb, 0x24, 0x0f, 0x5c, 0x14, 0x26, 0x3d, 0xd9, 0x4a, 0x4e, 0xee,
	0xe0, 0xc4, 0xc4, 0x15, 0x2b, 0xca, 0xae, 0x42, 0x6b, 0x84, 0xd6, 0x51, 0x1f, 0xac, 0x91, 0x6b,
	0xf7, 0x06, 0x0c, 0xca, 0xde, 0x3b, 0xce, 0xd1, 0xf4, 0xd0, 0x9c, 0xf5, 0xf8, 0xa9, 0xe3, 0x5c,
	0x7b, 0x73, 0xf0, 0xca, 0x36, 0x6b, 0xa4, 0x08, 0xfc, 0xd0, 0x8a, 0xc6, 0xe9, 0xa8, 0x6c, 0x37,
	0x52, 0x90, 0x5b, 0x70, 0x15, 0xa3, 0x55, 0x50, 0x84, 0x56, 0x34, 0x5d, 0xcd, 0xe3, 0xa1, 0x83,
	0xb8, 0xff, 0xff, 0xf8, 0x85, 0xd1, 0x2a, 0x45, 0x85, 0xdc, 0x83, 0x5f, 0x51, 0x95, 0xef, 0x59,
	0x1b, 0xb0, 0xd0, 0x89, 0xa6, 0xab, 0xd3, 0x3f, 0xf6, 0x5a, 0xe7, 0xe9, 0xa0, 0x2d, 0x1e, 0x61,
	0x84, 0x84, 0x9c, 0xc3, 0x18, 0xd9, 0xd0, 0xa1, 0x6b, 0x9c, 0xa4, 0x20, 0x67, 0xe0, 0xef, 0x99,
	0x14, 0x3a, 0xb1, 0x31, 0xf1, 0xf4, 0x9a, 0x14, 0x8b, 0x57, 0x70, 0xf5, 0x71, 0x72, 0x08, 0xf6,
	0xcf, 0x53, 0x76, 0xf9, 0x7f, 0xe5, 0xc7, 0xe0, 0x28, 0xba, 0x33, 0x4d, 0xeb, 0x51, 0x5f, 0xe4,
	0x62, 0x27, 0x7e, 0xf5, 0xeb, 0xeb, 0x7d, 0x2b, 0xf9, 0x9b, 0x87, 0x5f, 0xfd, 0xf0, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0x98, 0x6b, 0x2c, 0x10, 0x02, 0x00, 0x00,
}
