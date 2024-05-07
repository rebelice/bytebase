// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: store/instance.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// InstanceOptions is the option for instances.
type InstanceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How often the instance is synced.
	SyncInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	// The maximum number of connections.
	// The default is 10 if the value is unset or zero.
	MaximumConnections int32 `protobuf:"varint,3,opt,name=maximum_connections,json=maximumConnections,proto3" json:"maximum_connections,omitempty"`
}

func (x *InstanceOptions) Reset() {
	*x = InstanceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceOptions) ProtoMessage() {}

func (x *InstanceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_store_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceOptions.ProtoReflect.Descriptor instead.
func (*InstanceOptions) Descriptor() ([]byte, []int) {
	return file_store_instance_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceOptions) GetSyncInterval() *durationpb.Duration {
	if x != nil {
		return x.SyncInterval
	}
	return nil
}

func (x *InstanceOptions) GetMaximumConnections() int32 {
	if x != nil {
		return x.MaximumConnections
	}
	return 0
}

// InstanceMetadata is the metadata for instances.
type InstanceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The lower_case_table_names config for MySQL instances.
	// It is used to determine whether the table names and database names are case sensitive.
	MysqlLowerCaseTableNames int32                  `protobuf:"varint,1,opt,name=mysql_lower_case_table_names,json=mysqlLowerCaseTableNames,proto3" json:"mysql_lower_case_table_names,omitempty"`
	LastSyncTime             *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
}

func (x *InstanceMetadata) Reset() {
	*x = InstanceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_instance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceMetadata) ProtoMessage() {}

func (x *InstanceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_store_instance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceMetadata.ProtoReflect.Descriptor instead.
func (*InstanceMetadata) Descriptor() ([]byte, []int) {
	return file_store_instance_proto_rawDescGZIP(), []int{1}
}

func (x *InstanceMetadata) GetMysqlLowerCaseTableNames() int32 {
	if x != nil {
		return x.MysqlLowerCaseTableNames
	}
	return 0
}

func (x *InstanceMetadata) GetLastSyncTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSyncTime
	}
	return nil
}

var File_store_instance_proto protoreflect.FileDescriptor

var file_store_instance_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73,
	0x79, 0x6e, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x13, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x22, 0x94, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x1c, 0x6d, 0x79, 0x73, 0x71, 0x6c,
	0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6d,
	0x79, 0x73, 0x71, 0x6c, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x61, 0x73, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_instance_proto_rawDescOnce sync.Once
	file_store_instance_proto_rawDescData = file_store_instance_proto_rawDesc
)

func file_store_instance_proto_rawDescGZIP() []byte {
	file_store_instance_proto_rawDescOnce.Do(func() {
		file_store_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_instance_proto_rawDescData)
	})
	return file_store_instance_proto_rawDescData
}

var file_store_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_instance_proto_goTypes = []interface{}{
	(*InstanceOptions)(nil),       // 0: bytebase.store.InstanceOptions
	(*InstanceMetadata)(nil),      // 1: bytebase.store.InstanceMetadata
	(*durationpb.Duration)(nil),   // 2: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_store_instance_proto_depIdxs = []int32{
	2, // 0: bytebase.store.InstanceOptions.sync_interval:type_name -> google.protobuf.Duration
	3, // 1: bytebase.store.InstanceMetadata.last_sync_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_store_instance_proto_init() }
func file_store_instance_proto_init() {
	if File_store_instance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceOptions); i {
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
		file_store_instance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceMetadata); i {
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
			RawDescriptor: file_store_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_instance_proto_goTypes,
		DependencyIndexes: file_store_instance_proto_depIdxs,
		MessageInfos:      file_store_instance_proto_msgTypes,
	}.Build()
	File_store_instance_proto = out.File
	file_store_instance_proto_rawDesc = nil
	file_store_instance_proto_goTypes = nil
	file_store_instance_proto_depIdxs = nil
}
