// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: store/slow_query.proto

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

// SlowQueryStatistics is the slow query statistics.
type SlowQueryStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Items is the list of slow query statistics.
	Items []*SlowQueryStatisticsItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SlowQueryStatistics) Reset() {
	*x = SlowQueryStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_slow_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlowQueryStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlowQueryStatistics) ProtoMessage() {}

func (x *SlowQueryStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_store_slow_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlowQueryStatistics.ProtoReflect.Descriptor instead.
func (*SlowQueryStatistics) Descriptor() ([]byte, []int) {
	return file_store_slow_query_proto_rawDescGZIP(), []int{0}
}

func (x *SlowQueryStatistics) GetItems() []*SlowQueryStatisticsItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// SlowQueryStatisticsItem is the item of slow query statistics.
type SlowQueryStatisticsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sql_fingerprint is the fingerprint of the slow query.
	SqlFingerprint string `protobuf:"bytes,1,opt,name=sql_fingerprint,json=sqlFingerprint,proto3" json:"sql_fingerprint,omitempty"`
	// count is the number of slow queries with the same fingerprint.
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// latest_log_time is the time of the latest slow query with the same fingerprint.
	LatestLogTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=latest_log_time,json=latestLogTime,proto3" json:"latest_log_time,omitempty"`
	// The total query time of the slow query log.
	TotalQueryTime *durationpb.Duration `protobuf:"bytes,4,opt,name=total_query_time,json=totalQueryTime,proto3" json:"total_query_time,omitempty"`
	// The maximum query time of the slow query log.
	MaximumQueryTime *durationpb.Duration `protobuf:"bytes,5,opt,name=maximum_query_time,json=maximumQueryTime,proto3" json:"maximum_query_time,omitempty"`
	// The total rows sent of the slow query log.
	TotalRowsSent int64 `protobuf:"varint,6,opt,name=total_rows_sent,json=totalRowsSent,proto3" json:"total_rows_sent,omitempty"`
	// The maximum rows sent of the slow query log.
	MaximumRowsSent int32 `protobuf:"varint,7,opt,name=maximum_rows_sent,json=maximumRowsSent,proto3" json:"maximum_rows_sent,omitempty"`
	// The total rows examined of the slow query log.
	TotalRowsExamined int64 `protobuf:"varint,8,opt,name=total_rows_examined,json=totalRowsExamined,proto3" json:"total_rows_examined,omitempty"`
	// The maximum rows examined of the slow query log.
	MaximumRowsExamined int32 `protobuf:"varint,9,opt,name=maximum_rows_examined,json=maximumRowsExamined,proto3" json:"maximum_rows_examined,omitempty"`
	// samples are the details of the sample slow queries with the same fingerprint.
	Samples []*SlowQueryDetails `protobuf:"bytes,10,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *SlowQueryStatisticsItem) Reset() {
	*x = SlowQueryStatisticsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_slow_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlowQueryStatisticsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlowQueryStatisticsItem) ProtoMessage() {}

func (x *SlowQueryStatisticsItem) ProtoReflect() protoreflect.Message {
	mi := &file_store_slow_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlowQueryStatisticsItem.ProtoReflect.Descriptor instead.
func (*SlowQueryStatisticsItem) Descriptor() ([]byte, []int) {
	return file_store_slow_query_proto_rawDescGZIP(), []int{1}
}

func (x *SlowQueryStatisticsItem) GetSqlFingerprint() string {
	if x != nil {
		return x.SqlFingerprint
	}
	return ""
}

func (x *SlowQueryStatisticsItem) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SlowQueryStatisticsItem) GetLatestLogTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LatestLogTime
	}
	return nil
}

func (x *SlowQueryStatisticsItem) GetTotalQueryTime() *durationpb.Duration {
	if x != nil {
		return x.TotalQueryTime
	}
	return nil
}

func (x *SlowQueryStatisticsItem) GetMaximumQueryTime() *durationpb.Duration {
	if x != nil {
		return x.MaximumQueryTime
	}
	return nil
}

func (x *SlowQueryStatisticsItem) GetTotalRowsSent() int64 {
	if x != nil {
		return x.TotalRowsSent
	}
	return 0
}

func (x *SlowQueryStatisticsItem) GetMaximumRowsSent() int32 {
	if x != nil {
		return x.MaximumRowsSent
	}
	return 0
}

func (x *SlowQueryStatisticsItem) GetTotalRowsExamined() int64 {
	if x != nil {
		return x.TotalRowsExamined
	}
	return 0
}

func (x *SlowQueryStatisticsItem) GetMaximumRowsExamined() int32 {
	if x != nil {
		return x.MaximumRowsExamined
	}
	return 0
}

func (x *SlowQueryStatisticsItem) GetSamples() []*SlowQueryDetails {
	if x != nil {
		return x.Samples
	}
	return nil
}

// SlowQueryDetails is the details of a slow query.
type SlowQueryDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start_time is the start time of the slow query.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// query_time is the query time of the slow query.
	QueryTime *durationpb.Duration `protobuf:"bytes,2,opt,name=query_time,json=queryTime,proto3" json:"query_time,omitempty"`
	// lock_time is the lock time of the slow query.
	LockTime *durationpb.Duration `protobuf:"bytes,3,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`
	// rows_sent is the number of rows sent by the slow query.
	RowsSent int32 `protobuf:"varint,4,opt,name=rows_sent,json=rowsSent,proto3" json:"rows_sent,omitempty"`
	// rows_examined is the number of rows examined by the slow query.
	RowsExamined int32 `protobuf:"varint,5,opt,name=rows_examined,json=rowsExamined,proto3" json:"rows_examined,omitempty"`
	// sql_text is the SQL text of the slow query.
	SqlText string `protobuf:"bytes,6,opt,name=sql_text,json=sqlText,proto3" json:"sql_text,omitempty"`
}

func (x *SlowQueryDetails) Reset() {
	*x = SlowQueryDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_slow_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlowQueryDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlowQueryDetails) ProtoMessage() {}

func (x *SlowQueryDetails) ProtoReflect() protoreflect.Message {
	mi := &file_store_slow_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlowQueryDetails.ProtoReflect.Descriptor instead.
func (*SlowQueryDetails) Descriptor() ([]byte, []int) {
	return file_store_slow_query_proto_rawDescGZIP(), []int{2}
}

func (x *SlowQueryDetails) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *SlowQueryDetails) GetQueryTime() *durationpb.Duration {
	if x != nil {
		return x.QueryTime
	}
	return nil
}

func (x *SlowQueryDetails) GetLockTime() *durationpb.Duration {
	if x != nil {
		return x.LockTime
	}
	return nil
}

func (x *SlowQueryDetails) GetRowsSent() int32 {
	if x != nil {
		return x.RowsSent
	}
	return 0
}

func (x *SlowQueryDetails) GetRowsExamined() int32 {
	if x != nil {
		return x.RowsExamined
	}
	return 0
}

func (x *SlowQueryDetails) GetSqlText() string {
	if x != nil {
		return x.SqlText
	}
	return ""
}

var File_store_slow_query_proto protoreflect.FileDescriptor

var file_store_slow_query_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x6c, 0x6f, 0x77, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x13, 0x53, 0x6c, 0x6f,
	0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x6c, 0x6f, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x9e, 0x04, 0x0a, 0x17, 0x53, 0x6c, 0x6f, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x71, 0x6c, 0x5f, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x71, 0x6c, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43,
	0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73,
	0x53, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f,
	0x72, 0x6f, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x53, 0x65, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x65,
	0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73,
	0x5f, 0x65, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x13, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x45, 0x78, 0x61, 0x6d,
	0x69, 0x6e, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x6c, 0x6f, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x22, 0x9c, 0x02, 0x0a, 0x10, 0x53, 0x6c, 0x6f, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x38, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x45, 0x78, 0x61, 0x6d,
	0x69, 0x6e, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x71, 0x6c, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x71, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x42,
	0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_slow_query_proto_rawDescOnce sync.Once
	file_store_slow_query_proto_rawDescData = file_store_slow_query_proto_rawDesc
)

func file_store_slow_query_proto_rawDescGZIP() []byte {
	file_store_slow_query_proto_rawDescOnce.Do(func() {
		file_store_slow_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_slow_query_proto_rawDescData)
	})
	return file_store_slow_query_proto_rawDescData
}

var file_store_slow_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_slow_query_proto_goTypes = []interface{}{
	(*SlowQueryStatistics)(nil),     // 0: bytebase.store.SlowQueryStatistics
	(*SlowQueryStatisticsItem)(nil), // 1: bytebase.store.SlowQueryStatisticsItem
	(*SlowQueryDetails)(nil),        // 2: bytebase.store.SlowQueryDetails
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),     // 4: google.protobuf.Duration
}
var file_store_slow_query_proto_depIdxs = []int32{
	1, // 0: bytebase.store.SlowQueryStatistics.items:type_name -> bytebase.store.SlowQueryStatisticsItem
	3, // 1: bytebase.store.SlowQueryStatisticsItem.latest_log_time:type_name -> google.protobuf.Timestamp
	4, // 2: bytebase.store.SlowQueryStatisticsItem.total_query_time:type_name -> google.protobuf.Duration
	4, // 3: bytebase.store.SlowQueryStatisticsItem.maximum_query_time:type_name -> google.protobuf.Duration
	2, // 4: bytebase.store.SlowQueryStatisticsItem.samples:type_name -> bytebase.store.SlowQueryDetails
	3, // 5: bytebase.store.SlowQueryDetails.start_time:type_name -> google.protobuf.Timestamp
	4, // 6: bytebase.store.SlowQueryDetails.query_time:type_name -> google.protobuf.Duration
	4, // 7: bytebase.store.SlowQueryDetails.lock_time:type_name -> google.protobuf.Duration
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_store_slow_query_proto_init() }
func file_store_slow_query_proto_init() {
	if File_store_slow_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_slow_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlowQueryStatistics); i {
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
		file_store_slow_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlowQueryStatisticsItem); i {
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
		file_store_slow_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlowQueryDetails); i {
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
			RawDescriptor: file_store_slow_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_slow_query_proto_goTypes,
		DependencyIndexes: file_store_slow_query_proto_depIdxs,
		MessageInfos:      file_store_slow_query_proto_msgTypes,
	}.Build()
	File_store_slow_query_proto = out.File
	file_store_slow_query_proto_rawDesc = nil
	file_store_slow_query_proto_goTypes = nil
	file_store_slow_query_proto_depIdxs = nil
}
