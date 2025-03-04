// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pb/model/result.proto

package model

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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major string `protobuf:"bytes,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor string `protobuf:"bytes,2,opt,name=minor,proto3" json:"minor,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_pb_model_result_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *Version) GetMinor() string {
	if x != nil {
		return x.Minor
	}
	return ""
}

type SimulationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaVersion *Version `protobuf:"bytes,1,opt,name=schema_version,proto3" json:"schema_version,omitempty"`
	SimVersion    *string  `protobuf:"bytes,2,opt,name=sim_version,proto3,oneof" json:"sim_version,omitempty"`
	Modified      *bool    `protobuf:"varint,3,opt,name=modified,proto3,oneof" json:"modified,omitempty"`
	BuildDate     string   `protobuf:"bytes,4,opt,name=build_date,proto3" json:"build_date,omitempty"`
}

func (x *SimulationResult) Reset() {
	*x = SimulationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationResult) ProtoMessage() {}

func (x *SimulationResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationResult.ProtoReflect.Descriptor instead.
func (*SimulationResult) Descriptor() ([]byte, []int) {
	return file_pb_model_result_proto_rawDescGZIP(), []int{1}
}

func (x *SimulationResult) GetSchemaVersion() *Version {
	if x != nil {
		return x.SchemaVersion
	}
	return nil
}

func (x *SimulationResult) GetSimVersion() string {
	if x != nil && x.SimVersion != nil {
		return *x.SimVersion
	}
	return ""
}

func (x *SimulationResult) GetModified() bool {
	if x != nil && x.Modified != nil {
		return *x.Modified
	}
	return false
}

func (x *SimulationResult) GetBuildDate() string {
	if x != nil {
		return x.BuildDate
	}
	return ""
}

type IterationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetIdMapping map[int32]string `protobuf:"bytes,1,rep,name=target_id_mapping,proto3" json:"target_id_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IterationResult) Reset() {
	*x = IterationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IterationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IterationResult) ProtoMessage() {}

func (x *IterationResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IterationResult.ProtoReflect.Descriptor instead.
func (*IterationResult) Descriptor() ([]byte, []int) {
	return file_pb_model_result_proto_rawDescGZIP(), []int{2}
}

func (x *IterationResult) GetTargetIdMapping() map[int32]string {
	if x != nil {
		return x.TargetIdMapping
	}
	return nil
}

type SimulationStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SimulationStatistics) Reset() {
	*x = SimulationStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationStatistics) ProtoMessage() {}

func (x *SimulationStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationStatistics.ProtoReflect.Descriptor instead.
func (*SimulationStatistics) Descriptor() ([]byte, []int) {
	return file_pb_model_result_proto_rawDescGZIP(), []int{3}
}

type OverviewStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min  *float64 `protobuf:"fixed64,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max  *float64 `protobuf:"fixed64,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
	Mean *float64 `protobuf:"fixed64,3,opt,name=mean,proto3,oneof" json:"mean,omitempty"`
	SD   *float64 `protobuf:"fixed64,4,opt,name=SD,json=sd,proto3,oneof" json:"SD,omitempty"`
	Q1   *float64 `protobuf:"fixed64,5,opt,name=Q1,json=q1,proto3,oneof" json:"Q1,omitempty"`
	Q2   *float64 `protobuf:"fixed64,6,opt,name=Q2,json=q2,proto3,oneof" json:"Q2,omitempty"`
	Q3   *float64 `protobuf:"fixed64,7,opt,name=Q3,json=q3,proto3,oneof" json:"Q3,omitempty"`
	Hist []uint32 `protobuf:"varint,8,rep,packed,name=hist,json=histogram,proto3" json:"hist,omitempty"`
}

func (x *OverviewStats) Reset() {
	*x = OverviewStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverviewStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverviewStats) ProtoMessage() {}

func (x *OverviewStats) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverviewStats.ProtoReflect.Descriptor instead.
func (*OverviewStats) Descriptor() ([]byte, []int) {
	return file_pb_model_result_proto_rawDescGZIP(), []int{4}
}

func (x *OverviewStats) GetMin() float64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *OverviewStats) GetMax() float64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

func (x *OverviewStats) GetMean() float64 {
	if x != nil && x.Mean != nil {
		return *x.Mean
	}
	return 0
}

func (x *OverviewStats) GetSD() float64 {
	if x != nil && x.SD != nil {
		return *x.SD
	}
	return 0
}

func (x *OverviewStats) GetQ1() float64 {
	if x != nil && x.Q1 != nil {
		return *x.Q1
	}
	return 0
}

func (x *OverviewStats) GetQ2() float64 {
	if x != nil && x.Q2 != nil {
		return *x.Q2
	}
	return 0
}

func (x *OverviewStats) GetQ3() float64 {
	if x != nil && x.Q3 != nil {
		return *x.Q3
	}
	return 0
}

func (x *OverviewStats) GetHist() []uint32 {
	if x != nil {
		return x.Hist
	}
	return nil
}

type DescriptiveStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min  *float64 `protobuf:"fixed64,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max  *float64 `protobuf:"fixed64,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
	Mean *float64 `protobuf:"fixed64,3,opt,name=mean,proto3,oneof" json:"mean,omitempty"`
	SD   *float64 `protobuf:"fixed64,4,opt,name=SD,json=sd,proto3,oneof" json:"SD,omitempty"`
}

func (x *DescriptiveStats) Reset() {
	*x = DescriptiveStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_model_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptiveStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptiveStats) ProtoMessage() {}

func (x *DescriptiveStats) ProtoReflect() protoreflect.Message {
	mi := &file_pb_model_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptiveStats.ProtoReflect.Descriptor instead.
func (*DescriptiveStats) Descriptor() ([]byte, []int) {
	return file_pb_model_result_proto_rawDescGZIP(), []int{5}
}

func (x *DescriptiveStats) GetMin() float64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *DescriptiveStats) GetMax() float64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

func (x *DescriptiveStats) GetMean() float64 {
	if x != nil && x.Mean != nil {
		return *x.Mean
	}
	return 0
}

func (x *DescriptiveStats) GetSD() float64 {
	if x != nil && x.SD != nil {
		return *x.SD
	}
	return 0
}

var File_pb_model_result_proto protoreflect.FileDescriptor

var file_pb_model_result_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x35,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x0e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x69, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x69, 0x6d, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73,
	0x69, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x49, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x59, 0x0a, 0x11, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49,
	0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x42, 0x0a, 0x14, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x49, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x22, 0xf8, 0x01, 0x0a, 0x0d, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d,
	0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x02, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x53,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x03, 0x52, 0x02, 0x73, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x13, 0x0a, 0x02, 0x51, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x04, 0x52, 0x02,
	0x71, 0x31, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x51, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x05, 0x52, 0x02, 0x71, 0x32, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x51, 0x33,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x48, 0x06, 0x52, 0x02, 0x71, 0x33, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x68, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6d, 0x65, 0x61,
	0x6e, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x53, 0x44, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x51, 0x31, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x51, 0x32, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x51, 0x33, 0x22, 0x8e, 0x01,
	0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x02,
	0x52, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x53, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x03, 0x52, 0x02, 0x73, 0x64, 0x88, 0x01, 0x01, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6d, 0x65, 0x61, 0x6e, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x53, 0x44, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d,
	0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x2f, 0x73, 0x72, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_model_result_proto_rawDescOnce sync.Once
	file_pb_model_result_proto_rawDescData = file_pb_model_result_proto_rawDesc
)

func file_pb_model_result_proto_rawDescGZIP() []byte {
	file_pb_model_result_proto_rawDescOnce.Do(func() {
		file_pb_model_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_model_result_proto_rawDescData)
	})
	return file_pb_model_result_proto_rawDescData
}

var file_pb_model_result_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_model_result_proto_goTypes = []interface{}{
	(*Version)(nil),              // 0: model.Version
	(*SimulationResult)(nil),     // 1: model.SimulationResult
	(*IterationResult)(nil),      // 2: model.IterationResult
	(*SimulationStatistics)(nil), // 3: model.SimulationStatistics
	(*OverviewStats)(nil),        // 4: model.OverviewStats
	(*DescriptiveStats)(nil),     // 5: model.DescriptiveStats
	nil,                          // 6: model.IterationResult.TargetIdMappingEntry
}
var file_pb_model_result_proto_depIdxs = []int32{
	0, // 0: model.SimulationResult.schema_version:type_name -> model.Version
	6, // 1: model.IterationResult.target_id_mapping:type_name -> model.IterationResult.TargetIdMappingEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_model_result_proto_init() }
func file_pb_model_result_proto_init() {
	if File_pb_model_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_model_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_pb_model_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationResult); i {
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
		file_pb_model_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IterationResult); i {
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
		file_pb_model_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationStatistics); i {
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
		file_pb_model_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverviewStats); i {
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
		file_pb_model_result_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptiveStats); i {
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
	file_pb_model_result_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pb_model_result_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_pb_model_result_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_model_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_model_result_proto_goTypes,
		DependencyIndexes: file_pb_model_result_proto_depIdxs,
		MessageInfos:      file_pb_model_result_proto_msgTypes,
	}.Build()
	File_pb_model_result_proto = out.File
	file_pb_model_result_proto_rawDesc = nil
	file_pb_model_result_proto_goTypes = nil
	file_pb_model_result_proto_depIdxs = nil
}
