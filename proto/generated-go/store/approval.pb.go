// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: store/approval.proto

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

type IssuePayloadApproval_Approver_Status int32

const (
	IssuePayloadApproval_Approver_STATUS_UNSPECIFIED IssuePayloadApproval_Approver_Status = 0
	IssuePayloadApproval_Approver_PENDING            IssuePayloadApproval_Approver_Status = 1
	IssuePayloadApproval_Approver_APPROVED           IssuePayloadApproval_Approver_Status = 2
	IssuePayloadApproval_Approver_REJECTED           IssuePayloadApproval_Approver_Status = 3
)

// Enum value maps for IssuePayloadApproval_Approver_Status.
var (
	IssuePayloadApproval_Approver_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "APPROVED",
		3: "REJECTED",
	}
	IssuePayloadApproval_Approver_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"PENDING":            1,
		"APPROVED":           2,
		"REJECTED":           3,
	}
)

func (x IssuePayloadApproval_Approver_Status) Enum() *IssuePayloadApproval_Approver_Status {
	p := new(IssuePayloadApproval_Approver_Status)
	*p = x
	return p
}

func (x IssuePayloadApproval_Approver_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssuePayloadApproval_Approver_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_store_approval_proto_enumTypes[0].Descriptor()
}

func (IssuePayloadApproval_Approver_Status) Type() protoreflect.EnumType {
	return &file_store_approval_proto_enumTypes[0]
}

func (x IssuePayloadApproval_Approver_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssuePayloadApproval_Approver_Status.Descriptor instead.
func (IssuePayloadApproval_Approver_Status) EnumDescriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Type of the ApprovalStep
// ALL means every node must be approved to proceed.
// ANY means approving any node will proceed.
type ApprovalStep_Type int32

const (
	ApprovalStep_TYPE_UNSPECIFIED ApprovalStep_Type = 0
	ApprovalStep_ALL              ApprovalStep_Type = 1
	ApprovalStep_ANY              ApprovalStep_Type = 2
)

// Enum value maps for ApprovalStep_Type.
var (
	ApprovalStep_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "ALL",
		2: "ANY",
	}
	ApprovalStep_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"ALL":              1,
		"ANY":              2,
	}
)

func (x ApprovalStep_Type) Enum() *ApprovalStep_Type {
	p := new(ApprovalStep_Type)
	*p = x
	return p
}

func (x ApprovalStep_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApprovalStep_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_approval_proto_enumTypes[1].Descriptor()
}

func (ApprovalStep_Type) Type() protoreflect.EnumType {
	return &file_store_approval_proto_enumTypes[1]
}

func (x ApprovalStep_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApprovalStep_Type.Descriptor instead.
func (ApprovalStep_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{3, 0}
}

// Type of the ApprovalNode.
// type determines who should approve this node.
// ANY_IN_GROUP means the ApprovalNode can be approved by an user from our predefined user group.
// See GroupValue below for the predefined user groups.
type ApprovalNode_Type int32

const (
	ApprovalNode_TYPE_UNSPECIFIED ApprovalNode_Type = 0
	ApprovalNode_ANY_IN_GROUP     ApprovalNode_Type = 1
)

// Enum value maps for ApprovalNode_Type.
var (
	ApprovalNode_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "ANY_IN_GROUP",
	}
	ApprovalNode_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"ANY_IN_GROUP":     1,
	}
)

func (x ApprovalNode_Type) Enum() *ApprovalNode_Type {
	p := new(ApprovalNode_Type)
	*p = x
	return p
}

func (x ApprovalNode_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApprovalNode_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_approval_proto_enumTypes[2].Descriptor()
}

func (ApprovalNode_Type) Type() protoreflect.EnumType {
	return &file_store_approval_proto_enumTypes[2]
}

func (x ApprovalNode_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApprovalNode_Type.Descriptor instead.
func (ApprovalNode_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{4, 0}
}

// The predefined user groups are:
// - WORKSPACE_OWNER
// - WORKSPACE_DBA
// - PROJECT_OWNER
// - PROJECT_MEMBER
type ApprovalNode_GroupValue int32

const (
	ApprovalNode_GROUP_VALUE_UNSPECIFILED ApprovalNode_GroupValue = 0
	ApprovalNode_WORKSPACE_OWNER          ApprovalNode_GroupValue = 1
	ApprovalNode_WORKSPACE_DBA            ApprovalNode_GroupValue = 2
	ApprovalNode_PROJECT_OWNER            ApprovalNode_GroupValue = 3
	ApprovalNode_PROJECT_MEMBER           ApprovalNode_GroupValue = 4
)

// Enum value maps for ApprovalNode_GroupValue.
var (
	ApprovalNode_GroupValue_name = map[int32]string{
		0: "GROUP_VALUE_UNSPECIFILED",
		1: "WORKSPACE_OWNER",
		2: "WORKSPACE_DBA",
		3: "PROJECT_OWNER",
		4: "PROJECT_MEMBER",
	}
	ApprovalNode_GroupValue_value = map[string]int32{
		"GROUP_VALUE_UNSPECIFILED": 0,
		"WORKSPACE_OWNER":          1,
		"WORKSPACE_DBA":            2,
		"PROJECT_OWNER":            3,
		"PROJECT_MEMBER":           4,
	}
)

func (x ApprovalNode_GroupValue) Enum() *ApprovalNode_GroupValue {
	p := new(ApprovalNode_GroupValue)
	*p = x
	return p
}

func (x ApprovalNode_GroupValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApprovalNode_GroupValue) Descriptor() protoreflect.EnumDescriptor {
	return file_store_approval_proto_enumTypes[3].Descriptor()
}

func (ApprovalNode_GroupValue) Type() protoreflect.EnumType {
	return &file_store_approval_proto_enumTypes[3]
}

func (x ApprovalNode_GroupValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApprovalNode_GroupValue.Descriptor instead.
func (ApprovalNode_GroupValue) EnumDescriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{4, 1}
}

// IssuePayloadApproval is a part of the payload of an issue.
// IssuePayloadApproval records the approval template used and the approval history.
type IssuePayloadApproval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApprovalTemplates []*ApprovalTemplate              `protobuf:"bytes,1,rep,name=approval_templates,json=approvalTemplates,proto3" json:"approval_templates,omitempty"`
	Approvers         []*IssuePayloadApproval_Approver `protobuf:"bytes,2,rep,name=approvers,proto3" json:"approvers,omitempty"`
	// If the value is `false`, it means that the backend is still finding matching approval templates.
	// If `true`, other fields are available.
	ApprovalFindingDone  bool   `protobuf:"varint,3,opt,name=approval_finding_done,json=approvalFindingDone,proto3" json:"approval_finding_done,omitempty"`
	ApprovalFindingError string `protobuf:"bytes,4,opt,name=approval_finding_error,json=approvalFindingError,proto3" json:"approval_finding_error,omitempty"`
	// Format: risks/{name}
	Risk string `protobuf:"bytes,5,opt,name=risk,proto3" json:"risk,omitempty"`
}

func (x *IssuePayloadApproval) Reset() {
	*x = IssuePayloadApproval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_approval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuePayloadApproval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuePayloadApproval) ProtoMessage() {}

func (x *IssuePayloadApproval) ProtoReflect() protoreflect.Message {
	mi := &file_store_approval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuePayloadApproval.ProtoReflect.Descriptor instead.
func (*IssuePayloadApproval) Descriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{0}
}

func (x *IssuePayloadApproval) GetApprovalTemplates() []*ApprovalTemplate {
	if x != nil {
		return x.ApprovalTemplates
	}
	return nil
}

func (x *IssuePayloadApproval) GetApprovers() []*IssuePayloadApproval_Approver {
	if x != nil {
		return x.Approvers
	}
	return nil
}

func (x *IssuePayloadApproval) GetApprovalFindingDone() bool {
	if x != nil {
		return x.ApprovalFindingDone
	}
	return false
}

func (x *IssuePayloadApproval) GetApprovalFindingError() string {
	if x != nil {
		return x.ApprovalFindingError
	}
	return ""
}

func (x *IssuePayloadApproval) GetRisk() string {
	if x != nil {
		return x.Risk
	}
	return ""
}

type ApprovalTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flow        *ApprovalFlow `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
	Title       string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   int32         `protobuf:"varint,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *ApprovalTemplate) Reset() {
	*x = ApprovalTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_approval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalTemplate) ProtoMessage() {}

func (x *ApprovalTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_store_approval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalTemplate.ProtoReflect.Descriptor instead.
func (*ApprovalTemplate) Descriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{1}
}

func (x *ApprovalTemplate) GetFlow() *ApprovalFlow {
	if x != nil {
		return x.Flow
	}
	return nil
}

func (x *ApprovalTemplate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ApprovalTemplate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApprovalTemplate) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

type ApprovalFlow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steps []*ApprovalStep `protobuf:"bytes,1,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *ApprovalFlow) Reset() {
	*x = ApprovalFlow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_approval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalFlow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalFlow) ProtoMessage() {}

func (x *ApprovalFlow) ProtoReflect() protoreflect.Message {
	mi := &file_store_approval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalFlow.ProtoReflect.Descriptor instead.
func (*ApprovalFlow) Descriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{2}
}

func (x *ApprovalFlow) GetSteps() []*ApprovalStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

type ApprovalStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  ApprovalStep_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bytebase.store.ApprovalStep_Type" json:"type,omitempty"`
	Nodes []*ApprovalNode   `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ApprovalStep) Reset() {
	*x = ApprovalStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_approval_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalStep) ProtoMessage() {}

func (x *ApprovalStep) ProtoReflect() protoreflect.Message {
	mi := &file_store_approval_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalStep.ProtoReflect.Descriptor instead.
func (*ApprovalStep) Descriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{3}
}

func (x *ApprovalStep) GetType() ApprovalStep_Type {
	if x != nil {
		return x.Type
	}
	return ApprovalStep_TYPE_UNSPECIFIED
}

func (x *ApprovalStep) GetNodes() []*ApprovalNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type ApprovalNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ApprovalNode_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bytebase.store.ApprovalNode_Type" json:"type,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*ApprovalNode_GroupValue_
	//	*ApprovalNode_Role
	//	*ApprovalNode_ExternalNodeId
	Payload isApprovalNode_Payload `protobuf_oneof:"payload"`
}

func (x *ApprovalNode) Reset() {
	*x = ApprovalNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_approval_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalNode) ProtoMessage() {}

func (x *ApprovalNode) ProtoReflect() protoreflect.Message {
	mi := &file_store_approval_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalNode.ProtoReflect.Descriptor instead.
func (*ApprovalNode) Descriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{4}
}

func (x *ApprovalNode) GetType() ApprovalNode_Type {
	if x != nil {
		return x.Type
	}
	return ApprovalNode_TYPE_UNSPECIFIED
}

func (m *ApprovalNode) GetPayload() isApprovalNode_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ApprovalNode) GetGroupValue() ApprovalNode_GroupValue {
	if x, ok := x.GetPayload().(*ApprovalNode_GroupValue_); ok {
		return x.GroupValue
	}
	return ApprovalNode_GROUP_VALUE_UNSPECIFILED
}

func (x *ApprovalNode) GetRole() string {
	if x, ok := x.GetPayload().(*ApprovalNode_Role); ok {
		return x.Role
	}
	return ""
}

func (x *ApprovalNode) GetExternalNodeId() string {
	if x, ok := x.GetPayload().(*ApprovalNode_ExternalNodeId); ok {
		return x.ExternalNodeId
	}
	return ""
}

type isApprovalNode_Payload interface {
	isApprovalNode_Payload()
}

type ApprovalNode_GroupValue_ struct {
	GroupValue ApprovalNode_GroupValue `protobuf:"varint,2,opt,name=group_value,json=groupValue,proto3,enum=bytebase.store.ApprovalNode_GroupValue,oneof"`
}

type ApprovalNode_Role struct {
	// Format: roles/{role}
	Role string `protobuf:"bytes,3,opt,name=role,proto3,oneof"`
}

type ApprovalNode_ExternalNodeId struct {
	ExternalNodeId string `protobuf:"bytes,4,opt,name=external_node_id,json=externalNodeId,proto3,oneof"`
}

func (*ApprovalNode_GroupValue_) isApprovalNode_Payload() {}

func (*ApprovalNode_Role) isApprovalNode_Payload() {}

func (*ApprovalNode_ExternalNodeId) isApprovalNode_Payload() {}

type IssuePayloadApproval_Approver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new status.
	Status IssuePayloadApproval_Approver_Status `protobuf:"varint,1,opt,name=status,proto3,enum=bytebase.store.IssuePayloadApproval_Approver_Status" json:"status,omitempty"`
	// The principal id of the approver.
	PrincipalId int32 `protobuf:"varint,2,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
}

func (x *IssuePayloadApproval_Approver) Reset() {
	*x = IssuePayloadApproval_Approver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_approval_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuePayloadApproval_Approver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuePayloadApproval_Approver) ProtoMessage() {}

func (x *IssuePayloadApproval_Approver) ProtoReflect() protoreflect.Message {
	mi := &file_store_approval_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuePayloadApproval_Approver.ProtoReflect.Descriptor instead.
func (*IssuePayloadApproval_Approver) Descriptor() ([]byte, []int) {
	return file_store_approval_proto_rawDescGZIP(), []int{0, 0}
}

func (x *IssuePayloadApproval_Approver) GetStatus() IssuePayloadApproval_Approver_Status {
	if x != nil {
		return x.Status
	}
	return IssuePayloadApproval_Approver_STATUS_UNSPECIFIED
}

func (x *IssuePayloadApproval_Approver) GetPrincipalId() int32 {
	if x != nil {
		return x.PrincipalId
	}
	return 0
}

var File_store_approval_proto protoreflect.FileDescriptor

var file_store_approval_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xfb, 0x03, 0x0a, 0x14, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12,
	0x4f, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x11, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x4b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x72, 0x52, 0x09, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x6f, 0x6e,
	0x65, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x66, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x46, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x69, 0x73, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x69, 0x73, 0x6b, 0x1a, 0xc6, 0x01, 0x0a, 0x08,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52,
	0x4f, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x22, 0x9b, 0x01, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x66, 0x6c, 0x6f,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x46, 0x6c,
	0x6f, 0x77, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53,
	0x74, 0x65, 0x70, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59,
	0x10, 0x02, 0x22, 0x89, 0x03, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4e, 0x59, 0x5f, 0x49, 0x4e,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x01, 0x22, 0x79, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x4f, 0x52,
	0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x44, 0x42, 0x41, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x04, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_approval_proto_rawDescOnce sync.Once
	file_store_approval_proto_rawDescData = file_store_approval_proto_rawDesc
)

func file_store_approval_proto_rawDescGZIP() []byte {
	file_store_approval_proto_rawDescOnce.Do(func() {
		file_store_approval_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_approval_proto_rawDescData)
	})
	return file_store_approval_proto_rawDescData
}

var file_store_approval_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_store_approval_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_store_approval_proto_goTypes = []interface{}{
	(IssuePayloadApproval_Approver_Status)(0), // 0: bytebase.store.IssuePayloadApproval.Approver.Status
	(ApprovalStep_Type)(0),                    // 1: bytebase.store.ApprovalStep.Type
	(ApprovalNode_Type)(0),                    // 2: bytebase.store.ApprovalNode.Type
	(ApprovalNode_GroupValue)(0),              // 3: bytebase.store.ApprovalNode.GroupValue
	(*IssuePayloadApproval)(nil),              // 4: bytebase.store.IssuePayloadApproval
	(*ApprovalTemplate)(nil),                  // 5: bytebase.store.ApprovalTemplate
	(*ApprovalFlow)(nil),                      // 6: bytebase.store.ApprovalFlow
	(*ApprovalStep)(nil),                      // 7: bytebase.store.ApprovalStep
	(*ApprovalNode)(nil),                      // 8: bytebase.store.ApprovalNode
	(*IssuePayloadApproval_Approver)(nil),     // 9: bytebase.store.IssuePayloadApproval.Approver
}
var file_store_approval_proto_depIdxs = []int32{
	5, // 0: bytebase.store.IssuePayloadApproval.approval_templates:type_name -> bytebase.store.ApprovalTemplate
	9, // 1: bytebase.store.IssuePayloadApproval.approvers:type_name -> bytebase.store.IssuePayloadApproval.Approver
	6, // 2: bytebase.store.ApprovalTemplate.flow:type_name -> bytebase.store.ApprovalFlow
	7, // 3: bytebase.store.ApprovalFlow.steps:type_name -> bytebase.store.ApprovalStep
	1, // 4: bytebase.store.ApprovalStep.type:type_name -> bytebase.store.ApprovalStep.Type
	8, // 5: bytebase.store.ApprovalStep.nodes:type_name -> bytebase.store.ApprovalNode
	2, // 6: bytebase.store.ApprovalNode.type:type_name -> bytebase.store.ApprovalNode.Type
	3, // 7: bytebase.store.ApprovalNode.group_value:type_name -> bytebase.store.ApprovalNode.GroupValue
	0, // 8: bytebase.store.IssuePayloadApproval.Approver.status:type_name -> bytebase.store.IssuePayloadApproval.Approver.Status
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_store_approval_proto_init() }
func file_store_approval_proto_init() {
	if File_store_approval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_approval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuePayloadApproval); i {
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
		file_store_approval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalTemplate); i {
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
		file_store_approval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalFlow); i {
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
		file_store_approval_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalStep); i {
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
		file_store_approval_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalNode); i {
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
		file_store_approval_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuePayloadApproval_Approver); i {
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
	file_store_approval_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ApprovalNode_GroupValue_)(nil),
		(*ApprovalNode_Role)(nil),
		(*ApprovalNode_ExternalNodeId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_approval_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_approval_proto_goTypes,
		DependencyIndexes: file_store_approval_proto_depIdxs,
		EnumInfos:         file_store_approval_proto_enumTypes,
		MessageInfos:      file_store_approval_proto_msgTypes,
	}.Build()
	File_store_approval_proto = out.File
	file_store_approval_proto_rawDesc = nil
	file_store_approval_proto_goTypes = nil
	file_store_approval_proto_depIdxs = nil
}
