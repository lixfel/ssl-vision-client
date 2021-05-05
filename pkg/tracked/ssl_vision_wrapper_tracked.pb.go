// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: ssl_vision_wrapper_tracked.proto

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

// A wrapper packet containing meta data of the source
// Also serves for the possibility to extend the protocol later
type TrackerWrapperPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A random UUID of the source that is kept constant at the source while running
	// If multiple sources are broadcasting to the same network, this id can be used to identify individual sources
	Uuid *string `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	// The name of the source software that is producing this messages.
	SourceName *string `protobuf:"bytes,2,opt,name=source_name,json=sourceName" json:"source_name,omitempty"`
	// The tracked frame
	TrackedFrame *TrackedFrame `protobuf:"bytes,3,opt,name=tracked_frame,json=trackedFrame" json:"tracked_frame,omitempty"`
}

func (x *TrackerWrapperPacket) Reset() {
	*x = TrackerWrapperPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_vision_wrapper_tracked_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackerWrapperPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackerWrapperPacket) ProtoMessage() {}

func (x *TrackerWrapperPacket) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_vision_wrapper_tracked_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackerWrapperPacket.ProtoReflect.Descriptor instead.
func (*TrackerWrapperPacket) Descriptor() ([]byte, []int) {
	return file_ssl_vision_wrapper_tracked_proto_rawDescGZIP(), []int{0}
}

func (x *TrackerWrapperPacket) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *TrackerWrapperPacket) GetSourceName() string {
	if x != nil && x.SourceName != nil {
		return *x.SourceName
	}
	return ""
}

func (x *TrackerWrapperPacket) GetTrackedFrame() *TrackedFrame {
	if x != nil {
		return x.TrackedFrame
	}
	return nil
}

var File_ssl_vision_wrapper_tracked_proto protoreflect.FileDescriptor

var file_ssl_vision_wrapper_tracked_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53,
	0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64,
}

var (
	file_ssl_vision_wrapper_tracked_proto_rawDescOnce sync.Once
	file_ssl_vision_wrapper_tracked_proto_rawDescData = file_ssl_vision_wrapper_tracked_proto_rawDesc
)

func file_ssl_vision_wrapper_tracked_proto_rawDescGZIP() []byte {
	file_ssl_vision_wrapper_tracked_proto_rawDescOnce.Do(func() {
		file_ssl_vision_wrapper_tracked_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_vision_wrapper_tracked_proto_rawDescData)
	})
	return file_ssl_vision_wrapper_tracked_proto_rawDescData
}

var file_ssl_vision_wrapper_tracked_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ssl_vision_wrapper_tracked_proto_goTypes = []interface{}{
	(*TrackerWrapperPacket)(nil), // 0: TrackerWrapperPacket
	(*TrackedFrame)(nil),         // 1: TrackedFrame
}
var file_ssl_vision_wrapper_tracked_proto_depIdxs = []int32{
	1, // 0: TrackerWrapperPacket.tracked_frame:type_name -> TrackedFrame
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ssl_vision_wrapper_tracked_proto_init() }
func file_ssl_vision_wrapper_tracked_proto_init() {
	if File_ssl_vision_wrapper_tracked_proto != nil {
		return
	}
	file_ssl_vision_detection_tracked_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ssl_vision_wrapper_tracked_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackerWrapperPacket); i {
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
			RawDescriptor: file_ssl_vision_wrapper_tracked_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_vision_wrapper_tracked_proto_goTypes,
		DependencyIndexes: file_ssl_vision_wrapper_tracked_proto_depIdxs,
		MessageInfos:      file_ssl_vision_wrapper_tracked_proto_msgTypes,
	}.Build()
	File_ssl_vision_wrapper_tracked_proto = out.File
	file_ssl_vision_wrapper_tracked_proto_rawDesc = nil
	file_ssl_vision_wrapper_tracked_proto_goTypes = nil
	file_ssl_vision_wrapper_tracked_proto_depIdxs = nil
}
