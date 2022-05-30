// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/prometheus/v1/prometheus.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// The request message containing the user's name.
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	St    int64  `protobuf:"varint,2,opt,name=st,proto3" json:"st,omitempty"`
	Et    int64  `protobuf:"varint,3,opt,name=et,proto3" json:"et,omitempty"`
	Step  string `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[0]
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
	return file_api_prometheus_v1_prometheus_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryRequest) GetSt() int64 {
	if x != nil {
		return x.St
	}
	return 0
}

func (x *QueryRequest) GetEt() int64 {
	if x != nil {
		return x.Et
	}
	return 0
}

func (x *QueryRequest) GetStep() string {
	if x != nil {
		return x.Step
	}
	return ""
}

// The response message containing the greetings
type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *MetricsData `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[1]
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
	return file_api_prometheus_v1_prometheus_proto_rawDescGZIP(), []int{1}
}

func (x *QueryResponse) GetResult() *MetricsData {
	if x != nil {
		return x.Result
	}
	return nil
}

// The request message containing the user's name.
type BatchQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	St    int64  `protobuf:"varint,2,opt,name=st,proto3" json:"st,omitempty"`
	Et    int64  `protobuf:"varint,3,opt,name=et,proto3" json:"et,omitempty"`
	Step  string `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *BatchQueryRequest) Reset() {
	*x = BatchQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchQueryRequest) ProtoMessage() {}

func (x *BatchQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchQueryRequest.ProtoReflect.Descriptor instead.
func (*BatchQueryRequest) Descriptor() ([]byte, []int) {
	return file_api_prometheus_v1_prometheus_proto_rawDescGZIP(), []int{2}
}

func (x *BatchQueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *BatchQueryRequest) GetSt() int64 {
	if x != nil {
		return x.St
	}
	return 0
}

func (x *BatchQueryRequest) GetEt() int64 {
	if x != nil {
		return x.Et
	}
	return 0
}

func (x *BatchQueryRequest) GetStep() string {
	if x != nil {
		return x.Step
	}
	return ""
}

// The response message containing the greetings
type BatchQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*MetricsData `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *BatchQueryResponse) Reset() {
	*x = BatchQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchQueryResponse) ProtoMessage() {}

func (x *BatchQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchQueryResponse.ProtoReflect.Descriptor instead.
func (*BatchQueryResponse) Descriptor() ([]byte, []int) {
	return file_api_prometheus_v1_prometheus_proto_rawDescGZIP(), []int{3}
}

func (x *BatchQueryResponse) GetResults() []*MetricsData {
	if x != nil {
		return x.Results
	}
	return nil
}

type MetricsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultType string         `protobuf:"bytes,1,opt,name=result_type,json=resultType,proto3" json:"result_type,omitempty"`
	Query      string         `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Result     []*MetricValue `protobuf:"bytes,3,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *MetricsData) Reset() {
	*x = MetricsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsData) ProtoMessage() {}

func (x *MetricsData) ProtoReflect() protoreflect.Message {
	mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsData.ProtoReflect.Descriptor instead.
func (*MetricsData) Descriptor() ([]byte, []int) {
	return file_api_prometheus_v1_prometheus_proto_rawDescGZIP(), []int{4}
}

func (x *MetricsData) GetResultType() string {
	if x != nil {
		return x.ResultType
	}
	return ""
}

func (x *MetricsData) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *MetricsData) GetResult() []*MetricValue {
	if x != nil {
		return x.Result
	}
	return nil
}

type MetricValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric map[string]string `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value  *Point            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Values []*Point          `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MetricValue) Reset() {
	*x = MetricValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricValue) ProtoMessage() {}

func (x *MetricValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricValue.ProtoReflect.Descriptor instead.
func (*MetricValue) Descriptor() ([]byte, []int) {
	return file_api_prometheus_v1_prometheus_proto_rawDescGZIP(), []int{5}
}

func (x *MetricValue) GetMetric() map[string]string {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *MetricValue) GetValue() *Point {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *MetricValue) GetValues() []*Point {
	if x != nil {
		return x.Values
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp float64 `protobuf:"fixed64,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_api_prometheus_v1_prometheus_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_api_prometheus_v1_prometheus_proto_rawDescGZIP(), []int{6}
}

func (x *Point) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Point) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_api_prometheus_v1_prometheus_proto protoreflect.FileDescriptor

var file_api_prometheus_v1_prometheus_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x13, 0x92, 0x41, 0x0d, 0x32, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x66,
	0x69, 0x6c, 0x65, 0x64, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22,
	0x0a, 0x02, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x12, 0x92, 0x41, 0x0c, 0x32,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xe0, 0x41, 0x01, 0x52, 0x02,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x02, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x10,
	0x92, 0x41, 0x0a, 0x32, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xe0, 0x41, 0x01,
	0x52, 0x02, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x74,
	0x65, 0x70, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x58, 0x0a, 0x0d, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x44, 0x61, 0x74, 0x61, 0x42, 0x13, 0x92, 0x41, 0x10, 0x32, 0x0e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xc9, 0x01, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0x92, 0x41, 0x2b, 0x32,
	0x29, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x2c, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x65, 0x64,
	0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x27, 0x7c, 0x27, 0x20, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x02, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x12, 0x92, 0x41, 0x0c, 0x32, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0xe0, 0x41, 0x01, 0x52, 0x02, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x02, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x10, 0x92, 0x41, 0x0a, 0x32, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0xe0, 0x41, 0x01, 0x52, 0x02, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x20, 0x73, 0x74, 0x65, 0x70, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x22, 0x5f, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x13, 0x92, 0x41, 0x10, 0x32, 0x0e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x20, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x22, 0x78, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xe2, 0x01, 0x0a,
	0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x2a, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x3b, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xd1,
	0x02, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x12, 0x92, 0x01,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x92, 0x41, 0x32,
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x12, 0x10, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x68, 0x74, 0x65, 0x75, 0x73, 0x2a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02,
	0x4f, 0x4b, 0x12, 0xad, 0x01, 0x0a, 0x0a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x92, 0x41, 0x38, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x12, 0x16, 0x62, 0x61, 0x74, 0x63, 0x68, 0x20, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x20, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x68, 0x74, 0x65, 0x75, 0x73, 0x2a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02,
	0x4f, 0x4b, 0x42, 0x6d, 0x0a, 0x21, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x74,
	0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_prometheus_v1_prometheus_proto_rawDescOnce sync.Once
	file_api_prometheus_v1_prometheus_proto_rawDescData = file_api_prometheus_v1_prometheus_proto_rawDesc
)

func file_api_prometheus_v1_prometheus_proto_rawDescGZIP() []byte {
	file_api_prometheus_v1_prometheus_proto_rawDescOnce.Do(func() {
		file_api_prometheus_v1_prometheus_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_prometheus_v1_prometheus_proto_rawDescData)
	})
	return file_api_prometheus_v1_prometheus_proto_rawDescData
}

var file_api_prometheus_v1_prometheus_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_prometheus_v1_prometheus_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),       // 0: prometheus.v1.QueryRequest
	(*QueryResponse)(nil),      // 1: prometheus.v1.QueryResponse
	(*BatchQueryRequest)(nil),  // 2: prometheus.v1.BatchQueryRequest
	(*BatchQueryResponse)(nil), // 3: prometheus.v1.BatchQueryResponse
	(*MetricsData)(nil),        // 4: prometheus.v1.MetricsData
	(*MetricValue)(nil),        // 5: prometheus.v1.MetricValue
	(*Point)(nil),              // 6: prometheus.v1.Point
	nil,                        // 7: prometheus.v1.MetricValue.MetricEntry
}
var file_api_prometheus_v1_prometheus_proto_depIdxs = []int32{
	4, // 0: prometheus.v1.QueryResponse.result:type_name -> prometheus.v1.MetricsData
	4, // 1: prometheus.v1.BatchQueryResponse.results:type_name -> prometheus.v1.MetricsData
	5, // 2: prometheus.v1.MetricsData.result:type_name -> prometheus.v1.MetricValue
	7, // 3: prometheus.v1.MetricValue.metric:type_name -> prometheus.v1.MetricValue.MetricEntry
	6, // 4: prometheus.v1.MetricValue.value:type_name -> prometheus.v1.Point
	6, // 5: prometheus.v1.MetricValue.values:type_name -> prometheus.v1.Point
	0, // 6: prometheus.v1.Prometheus.Query:input_type -> prometheus.v1.QueryRequest
	2, // 7: prometheus.v1.Prometheus.BatchQuery:input_type -> prometheus.v1.BatchQueryRequest
	1, // 8: prometheus.v1.Prometheus.Query:output_type -> prometheus.v1.QueryResponse
	3, // 9: prometheus.v1.Prometheus.BatchQuery:output_type -> prometheus.v1.BatchQueryResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_prometheus_v1_prometheus_proto_init() }
func file_api_prometheus_v1_prometheus_proto_init() {
	if File_api_prometheus_v1_prometheus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_prometheus_v1_prometheus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_prometheus_v1_prometheus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_prometheus_v1_prometheus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchQueryRequest); i {
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
		file_api_prometheus_v1_prometheus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchQueryResponse); i {
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
		file_api_prometheus_v1_prometheus_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsData); i {
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
		file_api_prometheus_v1_prometheus_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricValue); i {
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
		file_api_prometheus_v1_prometheus_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
			RawDescriptor: file_api_prometheus_v1_prometheus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_prometheus_v1_prometheus_proto_goTypes,
		DependencyIndexes: file_api_prometheus_v1_prometheus_proto_depIdxs,
		MessageInfos:      file_api_prometheus_v1_prometheus_proto_msgTypes,
	}.Build()
	File_api_prometheus_v1_prometheus_proto = out.File
	file_api_prometheus_v1_prometheus_proto_rawDesc = nil
	file_api_prometheus_v1_prometheus_proto_goTypes = nil
	file_api_prometheus_v1_prometheus_proto_depIdxs = nil
}
