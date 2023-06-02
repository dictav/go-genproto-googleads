// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/ads/googleads/v12/resources/shared_set.proto

package resources

import (
	enums "github.com/dictav/go-genproto-googleads/pb/v12/enums"
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

// SharedSets are used for sharing criterion exclusions across multiple
// campaigns.
type SharedSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the shared set.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedSets/{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of this shared set. Read only.
	Id *int64 `protobuf:"varint,8,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Immutable. The type of this shared set: each shared set holds only a single kind
	// of resource. Required. Immutable.
	Type enums.SharedSetTypeEnum_SharedSetType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v12.enums.SharedSetTypeEnum_SharedSetType" json:"type,omitempty"`
	// The name of this shared set. Required.
	// Shared Sets must have names that are unique among active shared sets of
	// the same type.
	// The length of this string should be between 1 and 255 UTF-8 bytes,
	// inclusive.
	Name *string `protobuf:"bytes,9,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The status of this shared set. Read only.
	Status enums.SharedSetStatusEnum_SharedSetStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v12.enums.SharedSetStatusEnum_SharedSetStatus" json:"status,omitempty"`
	// Output only. The number of shared criteria within this shared set. Read only.
	MemberCount *int64 `protobuf:"varint,10,opt,name=member_count,json=memberCount,proto3,oneof" json:"member_count,omitempty"`
	// Output only. The number of campaigns associated with this shared set. Read only.
	ReferenceCount *int64 `protobuf:"varint,11,opt,name=reference_count,json=referenceCount,proto3,oneof" json:"reference_count,omitempty"`
}

func (x *SharedSet) Reset() {
	*x = SharedSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_resources_shared_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedSet) ProtoMessage() {}

func (x *SharedSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_resources_shared_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedSet.ProtoReflect.Descriptor instead.
func (*SharedSet) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_resources_shared_set_proto_rawDescGZIP(), []int{0}
}

func (x *SharedSet) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *SharedSet) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SharedSet) GetType() enums.SharedSetTypeEnum_SharedSetType {
	if x != nil {
		return x.Type
	}
	return enums.SharedSetTypeEnum_SharedSetType(0)
}

func (x *SharedSet) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SharedSet) GetStatus() enums.SharedSetStatusEnum_SharedSetStatus {
	if x != nil {
		return x.Status
	}
	return enums.SharedSetStatusEnum_SharedSetStatus(0)
}

func (x *SharedSet) GetMemberCount() int64 {
	if x != nil && x.MemberCount != nil {
		return *x.MemberCount
	}
	return 0
}

func (x *SharedSet) GetReferenceCount() int64 {
	if x != nil && x.ReferenceCount != nil {
		return *x.ReferenceCount
	}
	return 0
}

var File_google_ads_googleads_v12_resources_shared_set_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_resources_shared_set_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x04, 0x0a, 0x09, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x24,
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x58, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x60, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2b, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x0b,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x31,
	0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x0e,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x3a, 0x5b, 0xea, 0x41, 0x58, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x12, 0x32, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x7b,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x80, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_resources_shared_set_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_resources_shared_set_proto_rawDescData = file_google_ads_googleads_v12_resources_shared_set_proto_rawDesc
)

func file_google_ads_googleads_v12_resources_shared_set_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_resources_shared_set_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_resources_shared_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_resources_shared_set_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_resources_shared_set_proto_rawDescData
}

var file_google_ads_googleads_v12_resources_shared_set_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_resources_shared_set_proto_goTypes = []interface{}{
	(*SharedSet)(nil),                              // 0: google.ads.googleads.v12.resources.SharedSet
	(enums.SharedSetTypeEnum_SharedSetType)(0),     // 1: google.ads.googleads.v12.enums.SharedSetTypeEnum.SharedSetType
	(enums.SharedSetStatusEnum_SharedSetStatus)(0), // 2: google.ads.googleads.v12.enums.SharedSetStatusEnum.SharedSetStatus
}
var file_google_ads_googleads_v12_resources_shared_set_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v12.resources.SharedSet.type:type_name -> google.ads.googleads.v12.enums.SharedSetTypeEnum.SharedSetType
	2, // 1: google.ads.googleads.v12.resources.SharedSet.status:type_name -> google.ads.googleads.v12.enums.SharedSetStatusEnum.SharedSetStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_resources_shared_set_proto_init() }
func file_google_ads_googleads_v12_resources_shared_set_proto_init() {
	if File_google_ads_googleads_v12_resources_shared_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_resources_shared_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedSet); i {
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
	file_google_ads_googleads_v12_resources_shared_set_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v12_resources_shared_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_resources_shared_set_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_resources_shared_set_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_resources_shared_set_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_resources_shared_set_proto = out.File
	file_google_ads_googleads_v12_resources_shared_set_proto_rawDesc = nil
	file_google_ads_googleads_v12_resources_shared_set_proto_goTypes = nil
	file_google_ads_googleads_v12_resources_shared_set_proto_depIdxs = nil
}
