// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/actuator_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// The request message for getting the theme resource.
type GetResourcePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetResourcePackageRequest) Reset() {
	*x = GetResourcePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePackageRequest) ProtoMessage() {}

func (x *GetResourcePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePackageRequest.ProtoReflect.Descriptor instead.
func (*GetResourcePackageRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{0}
}

// The theme resources.
type ResourcePackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The branding logo.
	Logo []byte `protobuf:"bytes,1,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *ResourcePackage) Reset() {
	*x = ResourcePackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcePackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePackage) ProtoMessage() {}

func (x *ResourcePackage) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePackage.ProtoReflect.Descriptor instead.
func (*ResourcePackage) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResourcePackage) GetLogo() []byte {
	if x != nil {
		return x.Logo
	}
	return nil
}

type GetActuatorInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetActuatorInfoRequest) Reset() {
	*x = GetActuatorInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActuatorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActuatorInfoRequest) ProtoMessage() {}

func (x *GetActuatorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActuatorInfoRequest.ProtoReflect.Descriptor instead.
func (*GetActuatorInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{2}
}

type UpdateActuatorInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actuator to update.
	Actuator *ActuatorInfo `protobuf:"bytes,1,opt,name=actuator,proto3" json:"actuator,omitempty"`
	// The list of fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateActuatorInfoRequest) Reset() {
	*x = UpdateActuatorInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateActuatorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActuatorInfoRequest) ProtoMessage() {}

func (x *UpdateActuatorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActuatorInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateActuatorInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateActuatorInfoRequest) GetActuator() *ActuatorInfo {
	if x != nil {
		return x.Actuator
	}
	return nil
}

func (x *UpdateActuatorInfoRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type ListDebugLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of logs to return. The service may return fewer than
	// this value.
	// If unspecified, at most 50 logs will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListDebugLog` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListDebugLog` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListDebugLogRequest) Reset() {
	*x = ListDebugLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDebugLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDebugLogRequest) ProtoMessage() {}

func (x *ListDebugLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDebugLogRequest.ProtoReflect.Descriptor instead.
func (*ListDebugLogRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListDebugLogRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDebugLogRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListDebugLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The logs from the specified request.
	Logs []*DebugLog `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListDebugLogResponse) Reset() {
	*x = ListDebugLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDebugLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDebugLogResponse) ProtoMessage() {}

func (x *ListDebugLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDebugLogResponse.ProtoReflect.Descriptor instead.
func (*ListDebugLogResponse) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListDebugLogResponse) GetLogs() []*DebugLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *ListDebugLogResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type DebugLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordTime  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=record_time,json=recordTime,proto3" json:"record_time,omitempty"`
	RequestPath string                 `protobuf:"bytes,2,opt,name=request_path,json=requestPath,proto3" json:"request_path,omitempty"`
	User        string                 `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Error       string                 `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	StackTrace  string                 `protobuf:"bytes,5,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
}

func (x *DebugLog) Reset() {
	*x = DebugLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugLog) ProtoMessage() {}

func (x *DebugLog) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugLog.ProtoReflect.Descriptor instead.
func (*DebugLog) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{6}
}

func (x *DebugLog) GetRecordTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordTime
	}
	return nil
}

func (x *DebugLog) GetRequestPath() string {
	if x != nil {
		return x.RequestPath
	}
	return ""
}

func (x *DebugLog) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *DebugLog) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *DebugLog) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

type DeleteCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCacheRequest) Reset() {
	*x = DeleteCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCacheRequest) ProtoMessage() {}

func (x *DeleteCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCacheRequest.ProtoReflect.Descriptor instead.
func (*DeleteCacheRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{7}
}

// ServerInfo is the API message for server info.
// Actuator concept is similar to the Spring Boot Actuator.
type ActuatorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version is the bytebase's server version
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// git_commit is the git commit hash of the build
	GitCommit string `protobuf:"bytes,2,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	// readonly flag means if the Bytebase is running in readonly mode.
	Readonly bool `protobuf:"varint,3,opt,name=readonly,proto3" json:"readonly,omitempty"`
	// saas flag means if the Bytebase is running in SaaS mode, some features are not allowed to edit by users.
	Saas bool `protobuf:"varint,4,opt,name=saas,proto3" json:"saas,omitempty"`
	// demo_name specifies the demo name, empty string means no demo.
	DemoName string `protobuf:"bytes,5,opt,name=demo_name,json=demoName,proto3" json:"demo_name,omitempty"`
	// host is the Bytebase instance host.
	Host string `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	// port is the Bytebase instance port.
	Port string `protobuf:"bytes,7,opt,name=port,proto3" json:"port,omitempty"`
	// external_url is the URL where user or webhook callback visits Bytebase.
	ExternalUrl string `protobuf:"bytes,8,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	// need_admin_setup flag means the Bytebase instance doesn't have any end users.
	NeedAdminSetup bool `protobuf:"varint,9,opt,name=need_admin_setup,json=needAdminSetup,proto3" json:"need_admin_setup,omitempty"`
	// disallow_signup is the flag to disable self-service signup.
	DisallowSignup bool `protobuf:"varint,10,opt,name=disallow_signup,json=disallowSignup,proto3" json:"disallow_signup,omitempty"`
	// last_active_time is the service last active time in UTC Time Format, any API calls will refresh this value.
	LastActiveTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=last_active_time,json=lastActiveTime,proto3" json:"last_active_time,omitempty"`
	// require_2fa is the flag to require 2FA for all users.
	Require_2Fa bool `protobuf:"varint,12,opt,name=require_2fa,json=require2fa,proto3" json:"require_2fa,omitempty"`
	// workspace_id is the identifier for the workspace.
	WorkspaceId string `protobuf:"bytes,13,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// gitops_webhook_url is the webhook URL for GitOps.
	GitopsWebhookUrl string `protobuf:"bytes,14,opt,name=gitops_webhook_url,json=gitopsWebhookUrl,proto3" json:"gitops_webhook_url,omitempty"`
	// debug flag means if the debug mode is enabled.
	Debug bool `protobuf:"varint,15,opt,name=debug,proto3" json:"debug,omitempty"`
	// lsp is the enablement of lsp in SQL Editor.
	Lsp bool `protobuf:"varint,16,opt,name=lsp,proto3" json:"lsp,omitempty"`
	// lsp is the enablement of data backup prior to data update.
	PreUpdateBackup bool `protobuf:"varint,17,opt,name=pre_update_backup,json=preUpdateBackup,proto3" json:"pre_update_backup,omitempty"`
	// iam_guard is the enablement of IAM checks.
	IamGuard           bool     `protobuf:"varint,18,opt,name=iam_guard,json=iamGuard,proto3" json:"iam_guard,omitempty"`
	UnlicensedFeatures []string `protobuf:"bytes,19,rep,name=unlicensed_features,json=unlicensedFeatures,proto3" json:"unlicensed_features,omitempty"`
}

func (x *ActuatorInfo) Reset() {
	*x = ActuatorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_actuator_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActuatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActuatorInfo) ProtoMessage() {}

func (x *ActuatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActuatorInfo.ProtoReflect.Descriptor instead.
func (*ActuatorInfo) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{8}
}

func (x *ActuatorInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ActuatorInfo) GetGitCommit() string {
	if x != nil {
		return x.GitCommit
	}
	return ""
}

func (x *ActuatorInfo) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

func (x *ActuatorInfo) GetSaas() bool {
	if x != nil {
		return x.Saas
	}
	return false
}

func (x *ActuatorInfo) GetDemoName() string {
	if x != nil {
		return x.DemoName
	}
	return ""
}

func (x *ActuatorInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ActuatorInfo) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ActuatorInfo) GetExternalUrl() string {
	if x != nil {
		return x.ExternalUrl
	}
	return ""
}

func (x *ActuatorInfo) GetNeedAdminSetup() bool {
	if x != nil {
		return x.NeedAdminSetup
	}
	return false
}

func (x *ActuatorInfo) GetDisallowSignup() bool {
	if x != nil {
		return x.DisallowSignup
	}
	return false
}

func (x *ActuatorInfo) GetLastActiveTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActiveTime
	}
	return nil
}

func (x *ActuatorInfo) GetRequire_2Fa() bool {
	if x != nil {
		return x.Require_2Fa
	}
	return false
}

func (x *ActuatorInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *ActuatorInfo) GetGitopsWebhookUrl() string {
	if x != nil {
		return x.GitopsWebhookUrl
	}
	return ""
}

func (x *ActuatorInfo) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *ActuatorInfo) GetLsp() bool {
	if x != nil {
		return x.Lsp
	}
	return false
}

func (x *ActuatorInfo) GetPreUpdateBackup() bool {
	if x != nil {
		return x.PreUpdateBackup
	}
	return false
}

func (x *ActuatorInfo) GetIamGuard() bool {
	if x != nil {
		return x.IamGuard
	}
	return false
}

func (x *ActuatorInfo) GetUnlicensedFeatures() []string {
	if x != nil {
		return x.UnlicensedFeatures
	}
	return nil
}

var File_v1_actuator_service_proto protoreflect.FileDescriptor

var file_v1_actuator_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9b,
	0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x08,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x51, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x69, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb5, 0x01, 0x0a, 0x08, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe0, 0x05, 0x0a, 0x0c, 0x41, 0x63, 0x74,
	0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0a, 0x67, 0x69, 0x74,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x03, 0x52, 0x09, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x20,
	0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x04, 0x73, 0x61, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x73, 0x61, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x64, 0x65,
	0x6d, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x03, 0x52, 0x08, 0x64, 0x65, 0x6d, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x03, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x27, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x10, 0x6e, 0x65,
	0x65, 0x64, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0e, 0x6e, 0x65, 0x65, 0x64,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x0f, 0x64, 0x69,
	0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x5f, 0x32, 0x66, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x32, 0x66, 0x61, 0x12, 0x27, 0x0a, 0x0c,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x12, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x5f,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x10, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x73, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6c, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x72,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x61, 0x6d, 0x5f, 0x67, 0x75, 0x61, 0x72, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x61, 0x6d, 0x47, 0x75, 0x61, 0x72, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x75, 0x6e,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x64, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x32, 0xef, 0x04, 0x0a, 0x0f,
	0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x23, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x1c, 0xda, 0x41, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x93, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x75, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x75,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3a, 0xda, 0x41, 0x14, 0x61,
	0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x08, 0x61, 0x63, 0x74, 0x75, 0x61,
	0x74, 0x6f, 0x72, 0x32, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x62, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x72, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x12, 0x20, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0xda, 0x41, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x7d,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x21, 0xda, 0x41, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x75, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x11, 0x5a,
	0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_actuator_service_proto_rawDescOnce sync.Once
	file_v1_actuator_service_proto_rawDescData = file_v1_actuator_service_proto_rawDesc
)

func file_v1_actuator_service_proto_rawDescGZIP() []byte {
	file_v1_actuator_service_proto_rawDescOnce.Do(func() {
		file_v1_actuator_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_actuator_service_proto_rawDescData)
	})
	return file_v1_actuator_service_proto_rawDescData
}

var file_v1_actuator_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_actuator_service_proto_goTypes = []any{
	(*GetResourcePackageRequest)(nil), // 0: bytebase.v1.GetResourcePackageRequest
	(*ResourcePackage)(nil),           // 1: bytebase.v1.ResourcePackage
	(*GetActuatorInfoRequest)(nil),    // 2: bytebase.v1.GetActuatorInfoRequest
	(*UpdateActuatorInfoRequest)(nil), // 3: bytebase.v1.UpdateActuatorInfoRequest
	(*ListDebugLogRequest)(nil),       // 4: bytebase.v1.ListDebugLogRequest
	(*ListDebugLogResponse)(nil),      // 5: bytebase.v1.ListDebugLogResponse
	(*DebugLog)(nil),                  // 6: bytebase.v1.DebugLog
	(*DeleteCacheRequest)(nil),        // 7: bytebase.v1.DeleteCacheRequest
	(*ActuatorInfo)(nil),              // 8: bytebase.v1.ActuatorInfo
	(*fieldmaskpb.FieldMask)(nil),     // 9: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 11: google.protobuf.Empty
}
var file_v1_actuator_service_proto_depIdxs = []int32{
	8,  // 0: bytebase.v1.UpdateActuatorInfoRequest.actuator:type_name -> bytebase.v1.ActuatorInfo
	9,  // 1: bytebase.v1.UpdateActuatorInfoRequest.update_mask:type_name -> google.protobuf.FieldMask
	6,  // 2: bytebase.v1.ListDebugLogResponse.logs:type_name -> bytebase.v1.DebugLog
	10, // 3: bytebase.v1.DebugLog.record_time:type_name -> google.protobuf.Timestamp
	10, // 4: bytebase.v1.ActuatorInfo.last_active_time:type_name -> google.protobuf.Timestamp
	2,  // 5: bytebase.v1.ActuatorService.GetActuatorInfo:input_type -> bytebase.v1.GetActuatorInfoRequest
	3,  // 6: bytebase.v1.ActuatorService.UpdateActuatorInfo:input_type -> bytebase.v1.UpdateActuatorInfoRequest
	7,  // 7: bytebase.v1.ActuatorService.DeleteCache:input_type -> bytebase.v1.DeleteCacheRequest
	4,  // 8: bytebase.v1.ActuatorService.ListDebugLog:input_type -> bytebase.v1.ListDebugLogRequest
	0,  // 9: bytebase.v1.ActuatorService.GetResourcePackage:input_type -> bytebase.v1.GetResourcePackageRequest
	8,  // 10: bytebase.v1.ActuatorService.GetActuatorInfo:output_type -> bytebase.v1.ActuatorInfo
	8,  // 11: bytebase.v1.ActuatorService.UpdateActuatorInfo:output_type -> bytebase.v1.ActuatorInfo
	11, // 12: bytebase.v1.ActuatorService.DeleteCache:output_type -> google.protobuf.Empty
	5,  // 13: bytebase.v1.ActuatorService.ListDebugLog:output_type -> bytebase.v1.ListDebugLogResponse
	1,  // 14: bytebase.v1.ActuatorService.GetResourcePackage:output_type -> bytebase.v1.ResourcePackage
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_v1_actuator_service_proto_init() }
func file_v1_actuator_service_proto_init() {
	if File_v1_actuator_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_actuator_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetResourcePackageRequest); i {
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
		file_v1_actuator_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResourcePackage); i {
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
		file_v1_actuator_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetActuatorInfoRequest); i {
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
		file_v1_actuator_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateActuatorInfoRequest); i {
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
		file_v1_actuator_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListDebugLogRequest); i {
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
		file_v1_actuator_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListDebugLogResponse); i {
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
		file_v1_actuator_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DebugLog); i {
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
		file_v1_actuator_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCacheRequest); i {
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
		file_v1_actuator_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ActuatorInfo); i {
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
			RawDescriptor: file_v1_actuator_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_actuator_service_proto_goTypes,
		DependencyIndexes: file_v1_actuator_service_proto_depIdxs,
		MessageInfos:      file_v1_actuator_service_proto_msgTypes,
	}.Build()
	File_v1_actuator_service_proto = out.File
	file_v1_actuator_service_proto_rawDesc = nil
	file_v1_actuator_service_proto_goTypes = nil
	file_v1_actuator_service_proto_depIdxs = nil
}
