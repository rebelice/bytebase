// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: store/project.proto

package store

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

// The type of target.
type ProtectionRule_Target int32

const (
	ProtectionRule_PROTECTION_TARGET_UNSPECIFIED ProtectionRule_Target = 0
	ProtectionRule_BRANCH                        ProtectionRule_Target = 1
	ProtectionRule_CHANGELIST                    ProtectionRule_Target = 2
)

// Enum value maps for ProtectionRule_Target.
var (
	ProtectionRule_Target_name = map[int32]string{
		0: "PROTECTION_TARGET_UNSPECIFIED",
		1: "BRANCH",
		2: "CHANGELIST",
	}
	ProtectionRule_Target_value = map[string]int32{
		"PROTECTION_TARGET_UNSPECIFIED": 0,
		"BRANCH":                        1,
		"CHANGELIST":                    2,
	}
)

func (x ProtectionRule_Target) Enum() *ProtectionRule_Target {
	p := new(ProtectionRule_Target)
	*p = x
	return p
}

func (x ProtectionRule_Target) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtectionRule_Target) Descriptor() protoreflect.EnumDescriptor {
	return file_store_project_proto_enumTypes[0].Descriptor()
}

func (ProtectionRule_Target) Type() protoreflect.EnumType {
	return &file_store_project_proto_enumTypes[0]
}

func (x ProtectionRule_Target) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtectionRule_Target.Descriptor instead.
func (ProtectionRule_Target) EnumDescriptor() ([]byte, []int) {
	return file_store_project_proto_rawDescGZIP(), []int{2, 0}
}

type ProtectionRule_BranchSource int32

const (
	ProtectionRule_BRANCH_SOURCE_UNSPECIFIED ProtectionRule_BranchSource = 0
	ProtectionRule_DATABASE                  ProtectionRule_BranchSource = 1
)

// Enum value maps for ProtectionRule_BranchSource.
var (
	ProtectionRule_BranchSource_name = map[int32]string{
		0: "BRANCH_SOURCE_UNSPECIFIED",
		1: "DATABASE",
	}
	ProtectionRule_BranchSource_value = map[string]int32{
		"BRANCH_SOURCE_UNSPECIFIED": 0,
		"DATABASE":                  1,
	}
)

func (x ProtectionRule_BranchSource) Enum() *ProtectionRule_BranchSource {
	p := new(ProtectionRule_BranchSource)
	*p = x
	return p
}

func (x ProtectionRule_BranchSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtectionRule_BranchSource) Descriptor() protoreflect.EnumDescriptor {
	return file_store_project_proto_enumTypes[1].Descriptor()
}

func (ProtectionRule_BranchSource) Type() protoreflect.EnumType {
	return &file_store_project_proto_enumTypes[1]
}

func (x ProtectionRule_BranchSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtectionRule_BranchSource.Descriptor instead.
func (ProtectionRule_BranchSource) EnumDescriptor() ([]byte, []int) {
	return file_store_project_proto_rawDescGZIP(), []int{2, 1}
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_store_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_store_project_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Label) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Label) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtectionRules  []*ProtectionRule `protobuf:"bytes,1,rep,name=protection_rules,json=protectionRules,proto3" json:"protection_rules,omitempty"`
	IssueLabels      []*Label          `protobuf:"bytes,2,rep,name=issue_labels,json=issueLabels,proto3" json:"issue_labels,omitempty"`
	ForceIssueLabels bool              `protobuf:"varint,3,opt,name=force_issue_labels,json=forceIssueLabels,proto3" json:"force_issue_labels,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_store_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_store_project_proto_rawDescGZIP(), []int{1}
}

func (x *Project) GetProtectionRules() []*ProtectionRule {
	if x != nil {
		return x.ProtectionRules
	}
	return nil
}

func (x *Project) GetIssueLabels() []*Label {
	if x != nil {
		return x.IssueLabels
	}
	return nil
}

func (x *Project) GetForceIssueLabels() bool {
	if x != nil {
		return x.ForceIssueLabels
	}
	return false
}

type ProtectionRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier for a node in UUID format.
	Id     string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Target ProtectionRule_Target `protobuf:"varint,2,opt,name=target,proto3,enum=bytebase.store.ProtectionRule_Target" json:"target,omitempty"`
	// The name of the branch/changelist or wildcard.
	NameFilter string `protobuf:"bytes,3,opt,name=name_filter,json=nameFilter,proto3" json:"name_filter,omitempty"`
	// The roles allowed to create branches or changelists, rebase branches, delete branches.
	// Format: roles/projectOwner.
	AllowedRoles []string                    `protobuf:"bytes,4,rep,name=allowed_roles,json=allowedRoles,proto3" json:"allowed_roles,omitempty"`
	BranchSource ProtectionRule_BranchSource `protobuf:"varint,5,opt,name=branch_source,json=branchSource,proto3,enum=bytebase.store.ProtectionRule_BranchSource" json:"branch_source,omitempty"`
}

func (x *ProtectionRule) Reset() {
	*x = ProtectionRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtectionRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtectionRule) ProtoMessage() {}

func (x *ProtectionRule) ProtoReflect() protoreflect.Message {
	mi := &file_store_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtectionRule.ProtoReflect.Descriptor instead.
func (*ProtectionRule) Descriptor() ([]byte, []int) {
	return file_store_project_proto_rawDescGZIP(), []int{2}
}

func (x *ProtectionRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProtectionRule) GetTarget() ProtectionRule_Target {
	if x != nil {
		return x.Target
	}
	return ProtectionRule_PROTECTION_TARGET_UNSPECIFIED
}

func (x *ProtectionRule) GetNameFilter() string {
	if x != nil {
		return x.NameFilter
	}
	return ""
}

func (x *ProtectionRule) GetAllowedRoles() []string {
	if x != nil {
		return x.AllowedRoles
	}
	return nil
}

func (x *ProtectionRule) GetBranchSource() ProtectionRule_BranchSource {
	if x != nil {
		return x.BranchSource
	}
	return ProtectionRule_BRANCH_SOURCE_UNSPECIFIED
}

var File_store_project_proto protoreflect.FileDescriptor

var file_store_project_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x49, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0xbc, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x49, 0x0a, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22,
	0xfd, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x0d, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x47, 0x0a, 0x06, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x52, 0x4f, 0x54, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x4c, 0x49, 0x53, 0x54,
	0x10, 0x02, 0x22, 0x3b, 0x0a, 0x0c, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x10, 0x01, 0x42,
	0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_project_proto_rawDescOnce sync.Once
	file_store_project_proto_rawDescData = file_store_project_proto_rawDesc
)

func file_store_project_proto_rawDescGZIP() []byte {
	file_store_project_proto_rawDescOnce.Do(func() {
		file_store_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_project_proto_rawDescData)
	})
	return file_store_project_proto_rawDescData
}

var file_store_project_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_project_proto_goTypes = []interface{}{
	(ProtectionRule_Target)(0),       // 0: bytebase.store.ProtectionRule.Target
	(ProtectionRule_BranchSource)(0), // 1: bytebase.store.ProtectionRule.BranchSource
	(*Label)(nil),                    // 2: bytebase.store.Label
	(*Project)(nil),                  // 3: bytebase.store.Project
	(*ProtectionRule)(nil),           // 4: bytebase.store.ProtectionRule
}
var file_store_project_proto_depIdxs = []int32{
	4, // 0: bytebase.store.Project.protection_rules:type_name -> bytebase.store.ProtectionRule
	2, // 1: bytebase.store.Project.issue_labels:type_name -> bytebase.store.Label
	0, // 2: bytebase.store.ProtectionRule.target:type_name -> bytebase.store.ProtectionRule.Target
	1, // 3: bytebase.store.ProtectionRule.branch_source:type_name -> bytebase.store.ProtectionRule.BranchSource
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_store_project_proto_init() }
func file_store_project_proto_init() {
	if File_store_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_store_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_store_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtectionRule); i {
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
			RawDescriptor: file_store_project_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_project_proto_goTypes,
		DependencyIndexes: file_store_project_proto_depIdxs,
		EnumInfos:         file_store_project_proto_enumTypes,
		MessageInfos:      file_store_project_proto_msgTypes,
	}.Build()
	File_store_project_proto = out.File
	file_store_project_proto_rawDesc = nil
	file_store_project_proto_goTypes = nil
	file_store_project_proto_depIdxs = nil
}
