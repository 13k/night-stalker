// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.6.1
// source: protocol/hero.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LocalizedName      string        `protobuf:"bytes,3,opt,name=localized_name,json=localizedName,proto3" json:"localized_name,omitempty"`
	Slug               string        `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Aliases            []string      `protobuf:"bytes,5,rep,name=aliases,proto3" json:"aliases,omitempty"`
	Roles              []HeroRole    `protobuf:"varint,6,rep,packed,name=roles,proto3,enum=protocol.HeroRole" json:"roles,omitempty"`
	RoleLevels         []int64       `protobuf:"varint,7,rep,packed,name=role_levels,json=roleLevels,proto3" json:"role_levels,omitempty"`
	Complexity         int64         `protobuf:"varint,8,opt,name=complexity,proto3" json:"complexity,omitempty"`
	Legs               int64         `protobuf:"varint,9,opt,name=legs,proto3" json:"legs,omitempty"`
	AttributePrimary   DotaAttribute `protobuf:"varint,10,opt,name=attribute_primary,json=attributePrimary,proto3,enum=protocol.DotaAttribute" json:"attribute_primary,omitempty"`
	AttackCapabilities DotaUnitCap   `protobuf:"varint,11,opt,name=attack_capabilities,json=attackCapabilities,proto3,enum=protocol.DotaUnitCap" json:"attack_capabilities,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_hero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_hero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_protocol_hero_proto_rawDescGZIP(), []int{0}
}

func (x *Hero) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Hero) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hero) GetLocalizedName() string {
	if x != nil {
		return x.LocalizedName
	}
	return ""
}

func (x *Hero) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Hero) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *Hero) GetRoles() []HeroRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Hero) GetRoleLevels() []int64 {
	if x != nil {
		return x.RoleLevels
	}
	return nil
}

func (x *Hero) GetComplexity() int64 {
	if x != nil {
		return x.Complexity
	}
	return 0
}

func (x *Hero) GetLegs() int64 {
	if x != nil {
		return x.Legs
	}
	return 0
}

func (x *Hero) GetAttributePrimary() DotaAttribute {
	if x != nil {
		return x.AttributePrimary
	}
	return DotaAttribute_DOTA_ATTRIBUTE_UNSPECIFIED
}

func (x *Hero) GetAttackCapabilities() DotaUnitCap {
	if x != nil {
		return x.AttackCapabilities
	}
	return DotaUnitCap_DOTA_UNIT_CAP_NO_ATTACK
}

type HeroMatches struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hero         *Hero     `protobuf:"bytes,100,opt,name=hero,proto3" json:"hero,omitempty"`
	Matches      []*Match  `protobuf:"bytes,101,rep,name=matches,proto3" json:"matches,omitempty"`
	KnownPlayers []*Player `protobuf:"bytes,102,rep,name=known_players,json=knownPlayers,proto3" json:"known_players,omitempty"`
}

func (x *HeroMatches) Reset() {
	*x = HeroMatches{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_hero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeroMatches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroMatches) ProtoMessage() {}

func (x *HeroMatches) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_hero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroMatches.ProtoReflect.Descriptor instead.
func (*HeroMatches) Descriptor() ([]byte, []int) {
	return file_protocol_hero_proto_rawDescGZIP(), []int{1}
}

func (x *HeroMatches) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

func (x *HeroMatches) GetMatches() []*Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *HeroMatches) GetKnownPlayers() []*Player {
	if x != nil {
		return x.KnownPlayers
	}
	return nil
}

var File_protocol_hero_proto protoreflect.FileDescriptor

var file_protocol_hero_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a,
	0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48,
	0x65, 0x72, 0x6f, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x65, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c,
	0x65, 0x67, 0x73, 0x12, 0x44, 0x0a, 0x11, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x6f, 0x74, 0x61, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x10, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x46, 0x0a, 0x13, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x44, 0x6f, 0x74, 0x61, 0x55, 0x6e, 0x69, 0x74, 0x43, 0x61, 0x70, 0x52, 0x12, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x72, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52,
	0x04, 0x68, 0x65, 0x72, 0x6f, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x12, 0x35, 0x0a, 0x0d, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x18, 0x66, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x2d,
	0x73, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_hero_proto_rawDescOnce sync.Once
	file_protocol_hero_proto_rawDescData = file_protocol_hero_proto_rawDesc
)

func file_protocol_hero_proto_rawDescGZIP() []byte {
	file_protocol_hero_proto_rawDescOnce.Do(func() {
		file_protocol_hero_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_hero_proto_rawDescData)
	})
	return file_protocol_hero_proto_rawDescData
}

var file_protocol_hero_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_hero_proto_goTypes = []interface{}{
	(*Hero)(nil),        // 0: protocol.Hero
	(*HeroMatches)(nil), // 1: protocol.HeroMatches
	(HeroRole)(0),       // 2: protocol.HeroRole
	(DotaAttribute)(0),  // 3: protocol.DotaAttribute
	(DotaUnitCap)(0),    // 4: protocol.DotaUnitCap
	(*Match)(nil),       // 5: protocol.Match
	(*Player)(nil),      // 6: protocol.Player
}
var file_protocol_hero_proto_depIdxs = []int32{
	2, // 0: protocol.Hero.roles:type_name -> protocol.HeroRole
	3, // 1: protocol.Hero.attribute_primary:type_name -> protocol.DotaAttribute
	4, // 2: protocol.Hero.attack_capabilities:type_name -> protocol.DotaUnitCap
	0, // 3: protocol.HeroMatches.hero:type_name -> protocol.Hero
	5, // 4: protocol.HeroMatches.matches:type_name -> protocol.Match
	6, // 5: protocol.HeroMatches.known_players:type_name -> protocol.Player
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protocol_hero_proto_init() }
func file_protocol_hero_proto_init() {
	if File_protocol_hero_proto != nil {
		return
	}
	file_protocol_enums_proto_init()
	file_protocol_match_proto_init()
	file_protocol_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protocol_hero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protocol_hero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeroMatches); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocol_hero_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_hero_proto_goTypes,
		DependencyIndexes: file_protocol_hero_proto_depIdxs,
		MessageInfos:      file_protocol_hero_proto_msgTypes,
	}.Build()
	File_protocol_hero_proto = out.File
	file_protocol_hero_proto_rawDesc = nil
	file_protocol_hero_proto_goTypes = nil
	file_protocol_hero_proto_depIdxs = nil
}
