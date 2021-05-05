// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: ssl_gc_common.proto

package tracked

import (
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

// Team is either blue or yellow
type Team int32

const (
	// team not set
	Team_UNKNOWN Team = 0
	// yellow team
	Team_YELLOW Team = 1
	// blue team
	Team_BLUE Team = 2
)

// Enum value maps for Team.
var (
	Team_name = map[int32]string{
		0: "UNKNOWN",
		1: "YELLOW",
		2: "BLUE",
	}
	Team_value = map[string]int32{
		"UNKNOWN": 0,
		"YELLOW":  1,
		"BLUE":    2,
	}
)

func (x Team) Enum() *Team {
	p := new(Team)
	*p = x
	return p
}

func (x Team) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Team) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_common_proto_enumTypes[0].Descriptor()
}

func (Team) Type() protoreflect.EnumType {
	return &file_ssl_gc_common_proto_enumTypes[0]
}

func (x Team) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Team) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Team(num)
	return nil
}

// Deprecated: Use Team.Descriptor instead.
func (Team) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_common_proto_rawDescGZIP(), []int{0}
}

// Division denotes the current division, which influences some rules
type Division int32

const (
	Division_DIV_UNKNOWN Division = 0
	Division_DIV_A       Division = 1
	Division_DIV_B       Division = 2
)

// Enum value maps for Division.
var (
	Division_name = map[int32]string{
		0: "DIV_UNKNOWN",
		1: "DIV_A",
		2: "DIV_B",
	}
	Division_value = map[string]int32{
		"DIV_UNKNOWN": 0,
		"DIV_A":       1,
		"DIV_B":       2,
	}
)

func (x Division) Enum() *Division {
	p := new(Division)
	*p = x
	return p
}

func (x Division) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Division) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_common_proto_enumTypes[1].Descriptor()
}

func (Division) Type() protoreflect.EnumType {
	return &file_ssl_gc_common_proto_enumTypes[1]
}

func (x Division) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Division) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Division(num)
	return nil
}

// Deprecated: Use Division.Descriptor instead.
func (Division) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_common_proto_rawDescGZIP(), []int{1}
}

// RobotId is the combination of a team and a robot id
type RobotId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the robot number
	Id *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// the team that the robot belongs to
	Team *Team `protobuf:"varint,2,opt,name=team,enum=Team" json:"team,omitempty"`
}

func (x *RobotId) Reset() {
	*x = RobotId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotId) ProtoMessage() {}

func (x *RobotId) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotId.ProtoReflect.Descriptor instead.
func (*RobotId) Descriptor() ([]byte, []int) {
	return file_ssl_gc_common_proto_rawDescGZIP(), []int{0}
}

func (x *RobotId) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *RobotId) GetTeam() Team {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return Team_UNKNOWN
}

var File_ssl_gc_common_proto protoreflect.FileDescriptor

var file_ssl_gc_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x07, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x2a, 0x29, 0x0a, 0x04, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x2a, 0x31, 0x0a, 0x08, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x49, 0x56, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x49, 0x56, 0x5f, 0x41, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x49, 0x56, 0x5f, 0x42, 0x10, 0x02, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d,
	0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x64,
}

var (
	file_ssl_gc_common_proto_rawDescOnce sync.Once
	file_ssl_gc_common_proto_rawDescData = file_ssl_gc_common_proto_rawDesc
)

func file_ssl_gc_common_proto_rawDescGZIP() []byte {
	file_ssl_gc_common_proto_rawDescOnce.Do(func() {
		file_ssl_gc_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_common_proto_rawDescData)
	})
	return file_ssl_gc_common_proto_rawDescData
}

var file_ssl_gc_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ssl_gc_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ssl_gc_common_proto_goTypes = []interface{}{
	(Team)(0),       // 0: Team
	(Division)(0),   // 1: Division
	(*RobotId)(nil), // 2: RobotId
}
var file_ssl_gc_common_proto_depIdxs = []int32{
	0, // 0: RobotId.team:type_name -> Team
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ssl_gc_common_proto_init() }
func file_ssl_gc_common_proto_init() {
	if File_ssl_gc_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotId); i {
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
			RawDescriptor: file_ssl_gc_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_common_proto_goTypes,
		DependencyIndexes: file_ssl_gc_common_proto_depIdxs,
		EnumInfos:         file_ssl_gc_common_proto_enumTypes,
		MessageInfos:      file_ssl_gc_common_proto_msgTypes,
	}.Build()
	File_ssl_gc_common_proto = out.File
	file_ssl_gc_common_proto_rawDesc = nil
	file_ssl_gc_common_proto_goTypes = nil
	file_ssl_gc_common_proto_depIdxs = nil
}
