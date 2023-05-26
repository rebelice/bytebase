// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/sql_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The instance to execute the query against.
	// Format: instances/{instance}
	Instance string `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// The connection database name to execute the query against.
	// For PostgreSQL, it's required.
	// For other database engines, it's optional. Use empty string to execute against without specifying a database.
	ConnectionDatabase string `protobuf:"bytes,2,opt,name=connection_database,json=connectionDatabase,proto3" json:"connection_database,omitempty"`
	// The SQL statement to execute.
	Statement string `protobuf:"bytes,3,opt,name=statement,proto3" json:"statement,omitempty"`
	// The maximum number of rows to return.
	Limit int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *QueryRequest) GetConnectionDatabase() string {
	if x != nil {
		return x.ConnectionDatabase
	}
	return ""
}

func (x *QueryRequest) GetStatement() string {
	if x != nil {
		return x.Statement
	}
	return ""
}

func (x *QueryRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The query result.
	Result []*QueryResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	// The query advice.
	Advice []*Advice `protobuf:"bytes,2,rep,name=advice,proto3" json:"advice,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{1}
}

func (x *QueryResponse) GetResult() []*QueryResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *QueryResponse) GetAdvice() []*Advice {
	if x != nil {
		return x.Advice
	}
	return nil
}

type QueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The column names of the query result.
	ColumnNames []string `protobuf:"bytes,1,rep,name=column_names,json=columnNames,proto3" json:"column_names,omitempty"`
	// The column types of the query result.
	ColumnTypeNames []string `protobuf:"bytes,2,rep,name=column_type_names,json=columnTypeNames,proto3" json:"column_type_names,omitempty"`
	// The data of the query result.
	Data []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	// The column is masked or not.
	Masked []bool `protobuf:"varint,4,rep,packed,name=masked,proto3" json:"masked,omitempty"`
	// The error message if the query failed.
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResult) GetColumnNames() []string {
	if x != nil {
		return x.ColumnNames
	}
	return nil
}

func (x *QueryResult) GetColumnTypeNames() []string {
	if x != nil {
		return x.ColumnTypeNames
	}
	return nil
}

func (x *QueryResult) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryResult) GetMasked() []bool {
	if x != nil {
		return x.Masked
	}
	return nil
}

func (x *QueryResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Advice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The advice status.
	// Should be one of the following:
	// - SUCCESS
	// - WARN
	// - ERROR
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The advice code.
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// The advice title.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// The advice content.
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	// The advice line number in the SQL statement.
	Line int32 `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
	// The advice details.
	Details string `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Advice) Reset() {
	*x = Advice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Advice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Advice) ProtoMessage() {}

func (x *Advice) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Advice.ProtoReflect.Descriptor instead.
func (*Advice) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{3}
}

func (x *Advice) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Advice) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Advice) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Advice) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Advice) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Advice) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type PrettyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engine Engine `protobuf:"varint,1,opt,name=engine,proto3,enum=bytebase.v1.Engine" json:"engine,omitempty"`
	// The SDL format SQL schema information that was dumped from a database engine.
	// This information will be sorted to match the order of statements in the userSchema.
	CurrentSchema string `protobuf:"bytes,2,opt,name=current_schema,json=currentSchema,proto3" json:"current_schema,omitempty"`
	// The expected SDL schema. This schema will be checked for correctness and normalized.
	ExpectedSchema string `protobuf:"bytes,3,opt,name=expected_schema,json=expectedSchema,proto3" json:"expected_schema,omitempty"`
}

func (x *PrettyRequest) Reset() {
	*x = PrettyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrettyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrettyRequest) ProtoMessage() {}

func (x *PrettyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrettyRequest.ProtoReflect.Descriptor instead.
func (*PrettyRequest) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{4}
}

func (x *PrettyRequest) GetEngine() Engine {
	if x != nil {
		return x.Engine
	}
	return Engine_ENGINE_UNSPECIFIED
}

func (x *PrettyRequest) GetCurrentSchema() string {
	if x != nil {
		return x.CurrentSchema
	}
	return ""
}

func (x *PrettyRequest) GetExpectedSchema() string {
	if x != nil {
		return x.ExpectedSchema
	}
	return ""
}

type PrettyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pretty-formatted version of current schema.
	CurrentSchema string `protobuf:"bytes,1,opt,name=current_schema,json=currentSchema,proto3" json:"current_schema,omitempty"`
	// The expected SDL schema after normalizing.
	ExpectedSchema string `protobuf:"bytes,2,opt,name=expected_schema,json=expectedSchema,proto3" json:"expected_schema,omitempty"`
}

func (x *PrettyResponse) Reset() {
	*x = PrettyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sql_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrettyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrettyResponse) ProtoMessage() {}

func (x *PrettyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sql_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrettyResponse.ProtoReflect.Descriptor instead.
func (*PrettyResponse) Descriptor() ([]byte, []int) {
	return file_v1_sql_service_proto_rawDescGZIP(), []int{5}
}

func (x *PrettyResponse) GetCurrentSchema() string {
	if x != nil {
		return x.CurrentSchema
	}
	return ""
}

func (x *PrettyResponse) GetExpectedSchema() string {
	if x != nil {
		return x.ExpectedSchema
	}
	return ""
}

var File_v1_sql_service_proto protoreflect.FileDescriptor

var file_v1_sql_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x73, 0x71, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a,
	0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x6e, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x64, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x61, 0x64, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x08, 0x52, 0x06, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x06, 0x41, 0x64, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x50, 0x72,
	0x65, 0x74, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x60, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x74,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x32, 0xe6, 0x01, 0x0a, 0x0a, 0x53,
	0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x06, 0x50, 0x72, 0x65,
	0x74, 0x74, 0x79, 0x12, 0x1a, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x65, 0x74, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x65, 0x74, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x71, 0x6c,
	0x2f, 0x70, 0x72, 0x65, 0x74, 0x74, 0x79, 0x12, 0x7a, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0xda, 0x41, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x3d, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x71, 0x6c, 0x3a, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_sql_service_proto_rawDescOnce sync.Once
	file_v1_sql_service_proto_rawDescData = file_v1_sql_service_proto_rawDesc
)

func file_v1_sql_service_proto_rawDescGZIP() []byte {
	file_v1_sql_service_proto_rawDescOnce.Do(func() {
		file_v1_sql_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_sql_service_proto_rawDescData)
	})
	return file_v1_sql_service_proto_rawDescData
}

var file_v1_sql_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_sql_service_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),   // 0: bytebase.v1.QueryRequest
	(*QueryResponse)(nil),  // 1: bytebase.v1.QueryResponse
	(*QueryResult)(nil),    // 2: bytebase.v1.QueryResult
	(*Advice)(nil),         // 3: bytebase.v1.Advice
	(*PrettyRequest)(nil),  // 4: bytebase.v1.PrettyRequest
	(*PrettyResponse)(nil), // 5: bytebase.v1.PrettyResponse
	(Engine)(0),            // 6: bytebase.v1.Engine
}
var file_v1_sql_service_proto_depIdxs = []int32{
	2, // 0: bytebase.v1.QueryResponse.result:type_name -> bytebase.v1.QueryResult
	3, // 1: bytebase.v1.QueryResponse.advice:type_name -> bytebase.v1.Advice
	6, // 2: bytebase.v1.PrettyRequest.engine:type_name -> bytebase.v1.Engine
	4, // 3: bytebase.v1.SQLService.Pretty:input_type -> bytebase.v1.PrettyRequest
	0, // 4: bytebase.v1.SQLService.Query:input_type -> bytebase.v1.QueryRequest
	5, // 5: bytebase.v1.SQLService.Pretty:output_type -> bytebase.v1.PrettyResponse
	1, // 6: bytebase.v1.SQLService.Query:output_type -> bytebase.v1.QueryResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_sql_service_proto_init() }
func file_v1_sql_service_proto_init() {
	if File_v1_sql_service_proto != nil {
		return
	}
	file_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_sql_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_v1_sql_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_v1_sql_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResult); i {
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
		file_v1_sql_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Advice); i {
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
		file_v1_sql_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrettyRequest); i {
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
		file_v1_sql_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrettyResponse); i {
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
			RawDescriptor: file_v1_sql_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_sql_service_proto_goTypes,
		DependencyIndexes: file_v1_sql_service_proto_depIdxs,
		MessageInfos:      file_v1_sql_service_proto_msgTypes,
	}.Build()
	File_v1_sql_service_proto = out.File
	file_v1_sql_service_proto_rawDesc = nil
	file_v1_sql_service_proto_goTypes = nil
	file_v1_sql_service_proto_depIdxs = nil
}
