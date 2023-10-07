// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/policy.proto

package store

import (
	expr "google.golang.org/genproto/googleapis/type/expr"
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

type SQLReviewRuleLevel int32

const (
	SQLReviewRuleLevel_LEVEL_UNSPECIFIED SQLReviewRuleLevel = 0
	SQLReviewRuleLevel_ERROR             SQLReviewRuleLevel = 1
	SQLReviewRuleLevel_WARNING           SQLReviewRuleLevel = 2
	SQLReviewRuleLevel_DISABLED          SQLReviewRuleLevel = 3
)

// Enum value maps for SQLReviewRuleLevel.
var (
	SQLReviewRuleLevel_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "ERROR",
		2: "WARNING",
		3: "DISABLED",
	}
	SQLReviewRuleLevel_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"ERROR":             1,
		"WARNING":           2,
		"DISABLED":          3,
	}
)

func (x SQLReviewRuleLevel) Enum() *SQLReviewRuleLevel {
	p := new(SQLReviewRuleLevel)
	*p = x
	return p
}

func (x SQLReviewRuleLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SQLReviewRuleLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_store_policy_proto_enumTypes[0].Descriptor()
}

func (SQLReviewRuleLevel) Type() protoreflect.EnumType {
	return &file_store_policy_proto_enumTypes[0]
}

func (x SQLReviewRuleLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SQLReviewRuleLevel.Descriptor instead.
func (SQLReviewRuleLevel) EnumDescriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{0}
}

type MaskingExceptionPolicy_MaskingException_Action int32

const (
	MaskingExceptionPolicy_MaskingException_ACTION_UNSPECIFIED MaskingExceptionPolicy_MaskingException_Action = 0
	MaskingExceptionPolicy_MaskingException_QUERY              MaskingExceptionPolicy_MaskingException_Action = 1
	MaskingExceptionPolicy_MaskingException_EXPORT             MaskingExceptionPolicy_MaskingException_Action = 2
)

// Enum value maps for MaskingExceptionPolicy_MaskingException_Action.
var (
	MaskingExceptionPolicy_MaskingException_Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "QUERY",
		2: "EXPORT",
	}
	MaskingExceptionPolicy_MaskingException_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"QUERY":              1,
		"EXPORT":             2,
	}
)

func (x MaskingExceptionPolicy_MaskingException_Action) Enum() *MaskingExceptionPolicy_MaskingException_Action {
	p := new(MaskingExceptionPolicy_MaskingException_Action)
	*p = x
	return p
}

func (x MaskingExceptionPolicy_MaskingException_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MaskingExceptionPolicy_MaskingException_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_store_policy_proto_enumTypes[1].Descriptor()
}

func (MaskingExceptionPolicy_MaskingException_Action) Type() protoreflect.EnumType {
	return &file_store_policy_proto_enumTypes[1]
}

func (x MaskingExceptionPolicy_MaskingException_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MaskingExceptionPolicy_MaskingException_Action.Descriptor instead.
func (MaskingExceptionPolicy_MaskingException_Action) EnumDescriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{4, 0, 0}
}

type IamPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Collection of binding.
	Bindings []*Binding `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
}

func (x *IamPolicy) Reset() {
	*x = IamPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPolicy) ProtoMessage() {}

func (x *IamPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPolicy.ProtoReflect.Descriptor instead.
func (*IamPolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{0}
}

func (x *IamPolicy) GetBindings() []*Binding {
	if x != nil {
		return x.Bindings
	}
	return nil
}

// Reference: https://cloud.google.com/pubsub/docs/reference/rpc/google.iam.v1#binding
type Binding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Role that is assigned to the list of members.
	// Format: roles/{role}
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// Specifies the principals requesting access for a Bytebase resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone.
	// * `user:{emailid}`: An email address that represents a specific Bytebase account. For example, `alice@example.com`.
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// The condition that is associated with this binding.
	// If the condition evaluates to true, then this binding applies to the current request.
	// If the condition evaluates to false, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding.
	Condition *expr.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *Binding) Reset() {
	*x = Binding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Binding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binding) ProtoMessage() {}

func (x *Binding) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binding.ProtoReflect.Descriptor instead.
func (*Binding) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Binding) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Binding) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Binding) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

type MaskingPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaskData []*MaskData `protobuf:"bytes,1,rep,name=mask_data,json=maskData,proto3" json:"mask_data,omitempty"`
}

func (x *MaskingPolicy) Reset() {
	*x = MaskingPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingPolicy) ProtoMessage() {}

func (x *MaskingPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingPolicy.ProtoReflect.Descriptor instead.
func (*MaskingPolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{2}
}

func (x *MaskingPolicy) GetMaskData() []*MaskData {
	if x != nil {
		return x.MaskData
	}
	return nil
}

type MaskData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema       string       `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Table        string       `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Column       string       `protobuf:"bytes,3,opt,name=column,proto3" json:"column,omitempty"`
	MaskingLevel MaskingLevel `protobuf:"varint,4,opt,name=masking_level,json=maskingLevel,proto3,enum=bytebase.store.MaskingLevel" json:"masking_level,omitempty"`
}

func (x *MaskData) Reset() {
	*x = MaskData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskData) ProtoMessage() {}

func (x *MaskData) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskData.ProtoReflect.Descriptor instead.
func (*MaskData) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{3}
}

func (x *MaskData) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *MaskData) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *MaskData) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *MaskData) GetMaskingLevel() MaskingLevel {
	if x != nil {
		return x.MaskingLevel
	}
	return MaskingLevel_MASKING_LEVEL_UNSPECIFIED
}

// MaskingExceptionPolicy is the allowlist of users who can access sensitive data.
type MaskingExceptionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaskingExceptions []*MaskingExceptionPolicy_MaskingException `protobuf:"bytes,1,rep,name=masking_exceptions,json=maskingExceptions,proto3" json:"masking_exceptions,omitempty"`
}

func (x *MaskingExceptionPolicy) Reset() {
	*x = MaskingExceptionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingExceptionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingExceptionPolicy) ProtoMessage() {}

func (x *MaskingExceptionPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingExceptionPolicy.ProtoReflect.Descriptor instead.
func (*MaskingExceptionPolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{4}
}

func (x *MaskingExceptionPolicy) GetMaskingExceptions() []*MaskingExceptionPolicy_MaskingException {
	if x != nil {
		return x.MaskingExceptions
	}
	return nil
}

type MaskingRulePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*MaskingRulePolicy_MaskingRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *MaskingRulePolicy) Reset() {
	*x = MaskingRulePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingRulePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingRulePolicy) ProtoMessage() {}

func (x *MaskingRulePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingRulePolicy.ProtoReflect.Descriptor instead.
func (*MaskingRulePolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{5}
}

func (x *MaskingRulePolicy) GetRules() []*MaskingRulePolicy_MaskingRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type SQLReviewPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RuleList []*SQLReviewRule `protobuf:"bytes,2,rep,name=rule_list,json=ruleList,proto3" json:"rule_list,omitempty"`
}

func (x *SQLReviewPolicy) Reset() {
	*x = SQLReviewPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLReviewPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLReviewPolicy) ProtoMessage() {}

func (x *SQLReviewPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLReviewPolicy.ProtoReflect.Descriptor instead.
func (*SQLReviewPolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{6}
}

func (x *SQLReviewPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SQLReviewPolicy) GetRuleList() []*SQLReviewRule {
	if x != nil {
		return x.RuleList
	}
	return nil
}

type SQLReviewRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string             `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Level   SQLReviewRuleLevel `protobuf:"varint,2,opt,name=level,proto3,enum=bytebase.store.SQLReviewRuleLevel" json:"level,omitempty"`
	Payload string             `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Engine  Engine             `protobuf:"varint,4,opt,name=engine,proto3,enum=bytebase.store.Engine" json:"engine,omitempty"`
	Comment string             `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *SQLReviewRule) Reset() {
	*x = SQLReviewRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLReviewRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLReviewRule) ProtoMessage() {}

func (x *SQLReviewRule) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLReviewRule.ProtoReflect.Descriptor instead.
func (*SQLReviewRule) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{7}
}

func (x *SQLReviewRule) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SQLReviewRule) GetLevel() SQLReviewRuleLevel {
	if x != nil {
		return x.Level
	}
	return SQLReviewRuleLevel_LEVEL_UNSPECIFIED
}

func (x *SQLReviewRule) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *SQLReviewRule) GetEngine() Engine {
	if x != nil {
		return x.Engine
	}
	return Engine_ENGINE_UNSPECIFIED
}

func (x *SQLReviewRule) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type MaskingExceptionPolicy_MaskingException struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// action is the action that the user can access sensitive data.
	Action MaskingExceptionPolicy_MaskingException_Action `protobuf:"varint,1,opt,name=action,proto3,enum=bytebase.store.MaskingExceptionPolicy_MaskingException_Action" json:"action,omitempty"`
	// Level is the masking level that the user can access sensitive data.
	MaskingLevel MaskingLevel `protobuf:"varint,2,opt,name=masking_level,json=maskingLevel,proto3,enum=bytebase.store.MaskingLevel" json:"masking_level,omitempty"`
	// Member is the principal who bind to this exception policy instance.
	//
	// * `user:{emailid}`: An email address that represents a specific Bytebase account. For example, `alice@example.com`.
	Member string `protobuf:"bytes,4,opt,name=member,proto3" json:"member,omitempty"`
	// The condition that is associated with this exception policy instance.
	Condition *expr.Expr `protobuf:"bytes,5,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *MaskingExceptionPolicy_MaskingException) Reset() {
	*x = MaskingExceptionPolicy_MaskingException{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingExceptionPolicy_MaskingException) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingExceptionPolicy_MaskingException) ProtoMessage() {}

func (x *MaskingExceptionPolicy_MaskingException) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingExceptionPolicy_MaskingException.ProtoReflect.Descriptor instead.
func (*MaskingExceptionPolicy_MaskingException) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{4, 0}
}

func (x *MaskingExceptionPolicy_MaskingException) GetAction() MaskingExceptionPolicy_MaskingException_Action {
	if x != nil {
		return x.Action
	}
	return MaskingExceptionPolicy_MaskingException_ACTION_UNSPECIFIED
}

func (x *MaskingExceptionPolicy_MaskingException) GetMaskingLevel() MaskingLevel {
	if x != nil {
		return x.MaskingLevel
	}
	return MaskingLevel_MASKING_LEVEL_UNSPECIFIED
}

func (x *MaskingExceptionPolicy_MaskingException) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *MaskingExceptionPolicy_MaskingException) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

type MaskingRulePolicy_MaskingRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier for a node in UUID format.
	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Condition    *expr.Expr   `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty"`
	MaskingLevel MaskingLevel `protobuf:"varint,3,opt,name=masking_level,json=maskingLevel,proto3,enum=bytebase.store.MaskingLevel" json:"masking_level,omitempty"`
}

func (x *MaskingRulePolicy_MaskingRule) Reset() {
	*x = MaskingRulePolicy_MaskingRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingRulePolicy_MaskingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingRulePolicy_MaskingRule) ProtoMessage() {}

func (x *MaskingRulePolicy_MaskingRule) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingRulePolicy_MaskingRule.ProtoReflect.Descriptor instead.
func (*MaskingRulePolicy_MaskingRule) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{5, 0}
}

func (x *MaskingRulePolicy_MaskingRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaskingRulePolicy_MaskingRule) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *MaskingRulePolicy_MaskingRule) GetMaskingLevel() MaskingLevel {
	if x != nil {
		return x.MaskingLevel
	}
	return MaskingLevel_MASKING_LEVEL_UNSPECIFIED
}

var File_store_policy_proto protoreflect.FileDescriptor

var file_store_policy_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x09, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x33, 0x0a,
	0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x68, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x0d,
	0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x35, 0x0a,
	0x09, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x93, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x6d, 0x61, 0x73, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x6d, 0x61,
	0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xb2, 0x03, 0x0a, 0x16, 0x4d,
	0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x66, 0x0a, 0x12, 0x6d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e,
	0x67, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x61, 0x73, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xaf, 0x02,
	0x0a, 0x10, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69,
	0x6e, 0x67, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x6d, 0x61,
	0x73, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x0c, 0x6d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x52,
	0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x22,
	0xec, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x43, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x91, 0x01, 0x0a, 0x0b, 0x4d,
	0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x6d,
	0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x0c, 0x6d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x61,
	0x0a, 0x0f, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x51, 0x0a, 0x12, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49,
	0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_policy_proto_rawDescOnce sync.Once
	file_store_policy_proto_rawDescData = file_store_policy_proto_rawDesc
)

func file_store_policy_proto_rawDescGZIP() []byte {
	file_store_policy_proto_rawDescOnce.Do(func() {
		file_store_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_policy_proto_rawDescData)
	})
	return file_store_policy_proto_rawDescData
}

var file_store_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_store_policy_proto_goTypes = []interface{}{
	(SQLReviewRuleLevel)(0),                             // 0: bytebase.store.SQLReviewRuleLevel
	(MaskingExceptionPolicy_MaskingException_Action)(0), // 1: bytebase.store.MaskingExceptionPolicy.MaskingException.Action
	(*IamPolicy)(nil),                                   // 2: bytebase.store.IamPolicy
	(*Binding)(nil),                                     // 3: bytebase.store.Binding
	(*MaskingPolicy)(nil),                               // 4: bytebase.store.MaskingPolicy
	(*MaskData)(nil),                                    // 5: bytebase.store.MaskData
	(*MaskingExceptionPolicy)(nil),                      // 6: bytebase.store.MaskingExceptionPolicy
	(*MaskingRulePolicy)(nil),                           // 7: bytebase.store.MaskingRulePolicy
	(*SQLReviewPolicy)(nil),                             // 8: bytebase.store.SQLReviewPolicy
	(*SQLReviewRule)(nil),                               // 9: bytebase.store.SQLReviewRule
	(*MaskingExceptionPolicy_MaskingException)(nil),     // 10: bytebase.store.MaskingExceptionPolicy.MaskingException
	(*MaskingRulePolicy_MaskingRule)(nil),               // 11: bytebase.store.MaskingRulePolicy.MaskingRule
	(*expr.Expr)(nil),                                   // 12: google.type.Expr
	(MaskingLevel)(0),                                   // 13: bytebase.store.MaskingLevel
	(Engine)(0),                                         // 14: bytebase.store.Engine
}
var file_store_policy_proto_depIdxs = []int32{
	3,  // 0: bytebase.store.IamPolicy.bindings:type_name -> bytebase.store.Binding
	12, // 1: bytebase.store.Binding.condition:type_name -> google.type.Expr
	5,  // 2: bytebase.store.MaskingPolicy.mask_data:type_name -> bytebase.store.MaskData
	13, // 3: bytebase.store.MaskData.masking_level:type_name -> bytebase.store.MaskingLevel
	10, // 4: bytebase.store.MaskingExceptionPolicy.masking_exceptions:type_name -> bytebase.store.MaskingExceptionPolicy.MaskingException
	11, // 5: bytebase.store.MaskingRulePolicy.rules:type_name -> bytebase.store.MaskingRulePolicy.MaskingRule
	9,  // 6: bytebase.store.SQLReviewPolicy.rule_list:type_name -> bytebase.store.SQLReviewRule
	0,  // 7: bytebase.store.SQLReviewRule.level:type_name -> bytebase.store.SQLReviewRuleLevel
	14, // 8: bytebase.store.SQLReviewRule.engine:type_name -> bytebase.store.Engine
	1,  // 9: bytebase.store.MaskingExceptionPolicy.MaskingException.action:type_name -> bytebase.store.MaskingExceptionPolicy.MaskingException.Action
	13, // 10: bytebase.store.MaskingExceptionPolicy.MaskingException.masking_level:type_name -> bytebase.store.MaskingLevel
	12, // 11: bytebase.store.MaskingExceptionPolicy.MaskingException.condition:type_name -> google.type.Expr
	12, // 12: bytebase.store.MaskingRulePolicy.MaskingRule.condition:type_name -> google.type.Expr
	13, // 13: bytebase.store.MaskingRulePolicy.MaskingRule.masking_level:type_name -> bytebase.store.MaskingLevel
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_store_policy_proto_init() }
func file_store_policy_proto_init() {
	if File_store_policy_proto != nil {
		return
	}
	file_store_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_store_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPolicy); i {
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
		file_store_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Binding); i {
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
		file_store_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingPolicy); i {
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
		file_store_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskData); i {
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
		file_store_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingExceptionPolicy); i {
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
		file_store_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingRulePolicy); i {
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
		file_store_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLReviewPolicy); i {
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
		file_store_policy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLReviewRule); i {
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
		file_store_policy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingExceptionPolicy_MaskingException); i {
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
		file_store_policy_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingRulePolicy_MaskingRule); i {
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
			RawDescriptor: file_store_policy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_policy_proto_goTypes,
		DependencyIndexes: file_store_policy_proto_depIdxs,
		EnumInfos:         file_store_policy_proto_enumTypes,
		MessageInfos:      file_store_policy_proto_msgTypes,
	}.Build()
	File_store_policy_proto = out.File
	file_store_policy_proto_rawDesc = nil
	file_store_policy_proto_goTypes = nil
	file_store_policy_proto_depIdxs = nil
}
