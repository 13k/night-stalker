// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.6.1
// source: protocol/league.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type League struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tier           LeagueTier           `protobuf:"varint,3,opt,name=tier,proto3,enum=ns.protocol.LeagueTier" json:"tier,omitempty"`
	Region         LeagueRegion         `protobuf:"varint,4,opt,name=region,proto3,enum=ns.protocol.LeagueRegion" json:"region,omitempty"`
	Status         LeagueStatus         `protobuf:"varint,5,opt,name=status,proto3,enum=ns.protocol.LeagueStatus" json:"status,omitempty"`
	TotalPrizePool uint32               `protobuf:"varint,6,opt,name=total_prize_pool,json=totalPrizePool,proto3" json:"total_prize_pool,omitempty"`
	LastActivityAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=last_activity_at,json=lastActivityAt,proto3" json:"last_activity_at,omitempty"`
	StartAt        *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	FinishAt       *timestamp.Timestamp `protobuf:"bytes,9,opt,name=finish_at,json=finishAt,proto3" json:"finish_at,omitempty"`
}

func (x *League) Reset() {
	*x = League{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_league_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *League) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*League) ProtoMessage() {}

func (x *League) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_league_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use League.ProtoReflect.Descriptor instead.
func (*League) Descriptor() ([]byte, []int) {
	return file_protocol_league_proto_rawDescGZIP(), []int{0}
}

func (x *League) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *League) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *League) GetTier() LeagueTier {
	if x != nil {
		return x.Tier
	}
	return LeagueTier_LEAGUE_TIER_UNSET
}

func (x *League) GetRegion() LeagueRegion {
	if x != nil {
		return x.Region
	}
	return LeagueRegion_LEAGUE_REGION_UNSET
}

func (x *League) GetStatus() LeagueStatus {
	if x != nil {
		return x.Status
	}
	return LeagueStatus_LEAGUE_STATUS_UNSET
}

func (x *League) GetTotalPrizePool() uint32 {
	if x != nil {
		return x.TotalPrizePool
	}
	return 0
}

func (x *League) GetLastActivityAt() *timestamp.Timestamp {
	if x != nil {
		return x.LastActivityAt
	}
	return nil
}

func (x *League) GetStartAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *League) GetFinishAt() *timestamp.Timestamp {
	if x != nil {
		return x.FinishAt
	}
	return nil
}

type Leagues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leagues []*League `protobuf:"bytes,100,rep,name=leagues,proto3" json:"leagues,omitempty"`
}

func (x *Leagues) Reset() {
	*x = Leagues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_league_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Leagues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leagues) ProtoMessage() {}

func (x *Leagues) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_league_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leagues.ProtoReflect.Descriptor instead.
func (*Leagues) Descriptor() ([]byte, []int) {
	return file_protocol_league_proto_rawDescGZIP(), []int{1}
}

func (x *Leagues) GetLeagues() []*League {
	if x != nil {
		return x.Leagues
	}
	return nil
}

var File_protocol_league_proto protoreflect.FileDescriptor

var file_protocol_league_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x03, 0x0a, 0x06,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x69,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x54, 0x69, 0x65,
	0x72, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a,
	0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x7a, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41, 0x74, 0x12, 0x35, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x41, 0x74, 0x22, 0x38, 0x0a,
	0x07, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x52, 0x07,
	0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x2d,
	0x73, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_league_proto_rawDescOnce sync.Once
	file_protocol_league_proto_rawDescData = file_protocol_league_proto_rawDesc
)

func file_protocol_league_proto_rawDescGZIP() []byte {
	file_protocol_league_proto_rawDescOnce.Do(func() {
		file_protocol_league_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_league_proto_rawDescData)
	})
	return file_protocol_league_proto_rawDescData
}

var file_protocol_league_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_league_proto_goTypes = []interface{}{
	(*League)(nil),              // 0: ns.protocol.League
	(*Leagues)(nil),             // 1: ns.protocol.Leagues
	(LeagueTier)(0),             // 2: ns.protocol.LeagueTier
	(LeagueRegion)(0),           // 3: ns.protocol.LeagueRegion
	(LeagueStatus)(0),           // 4: ns.protocol.LeagueStatus
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_protocol_league_proto_depIdxs = []int32{
	2, // 0: ns.protocol.League.tier:type_name -> ns.protocol.LeagueTier
	3, // 1: ns.protocol.League.region:type_name -> ns.protocol.LeagueRegion
	4, // 2: ns.protocol.League.status:type_name -> ns.protocol.LeagueStatus
	5, // 3: ns.protocol.League.last_activity_at:type_name -> google.protobuf.Timestamp
	5, // 4: ns.protocol.League.start_at:type_name -> google.protobuf.Timestamp
	5, // 5: ns.protocol.League.finish_at:type_name -> google.protobuf.Timestamp
	0, // 6: ns.protocol.Leagues.leagues:type_name -> ns.protocol.League
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protocol_league_proto_init() }
func file_protocol_league_proto_init() {
	if File_protocol_league_proto != nil {
		return
	}
	file_protocol_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protocol_league_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*League); i {
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
		file_protocol_league_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Leagues); i {
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
			RawDescriptor: file_protocol_league_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_league_proto_goTypes,
		DependencyIndexes: file_protocol_league_proto_depIdxs,
		MessageInfos:      file_protocol_league_proto_msgTypes,
	}.Build()
	File_protocol_league_proto = out.File
	file_protocol_league_proto_rawDesc = nil
	file_protocol_league_proto_goTypes = nil
	file_protocol_league_proto_depIdxs = nil
}
