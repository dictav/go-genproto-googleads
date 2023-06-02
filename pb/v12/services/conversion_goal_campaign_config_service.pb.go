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
// source: google/ads/googleads/v12/services/conversion_goal_campaign_config_service.proto

package services

import (
	context "context"
	enums "github.com/dictav/go-genproto-googleads/pb/v12/enums"
	resources "github.com/dictav/go-genproto-googleads/pb/v12/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [ConversionGoalCampaignConfigService.MutateConversionGoalCampaignConfig][].
type MutateConversionGoalCampaignConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose custom conversion goals are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual conversion goal campaign
	// config.
	Operations []*ConversionGoalCampaignConfigOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,4,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v12.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateConversionGoalCampaignConfigsRequest) Reset() {
	*x = MutateConversionGoalCampaignConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateConversionGoalCampaignConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConversionGoalCampaignConfigsRequest) ProtoMessage() {}

func (x *MutateConversionGoalCampaignConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConversionGoalCampaignConfigsRequest.ProtoReflect.Descriptor instead.
func (*MutateConversionGoalCampaignConfigsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetOperations() []*ConversionGoalCampaignConfigOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (update) on a conversion goal campaign config.
type ConversionGoalCampaignConfigOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*ConversionGoalCampaignConfigOperation_Update
	Operation isConversionGoalCampaignConfigOperation_Operation `protobuf_oneof:"operation"`
}

func (x *ConversionGoalCampaignConfigOperation) Reset() {
	*x = ConversionGoalCampaignConfigOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionGoalCampaignConfigOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionGoalCampaignConfigOperation) ProtoMessage() {}

func (x *ConversionGoalCampaignConfigOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionGoalCampaignConfigOperation.ProtoReflect.Descriptor instead.
func (*ConversionGoalCampaignConfigOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConversionGoalCampaignConfigOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *ConversionGoalCampaignConfigOperation) GetOperation() isConversionGoalCampaignConfigOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *ConversionGoalCampaignConfigOperation) GetUpdate() *resources.ConversionGoalCampaignConfig {
	if x, ok := x.GetOperation().(*ConversionGoalCampaignConfigOperation_Update); ok {
		return x.Update
	}
	return nil
}

type isConversionGoalCampaignConfigOperation_Operation interface {
	isConversionGoalCampaignConfigOperation_Operation()
}

type ConversionGoalCampaignConfigOperation_Update struct {
	// Update operation: The conversion goal campaign config is expected to have
	// a valid resource name.
	Update *resources.ConversionGoalCampaignConfig `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

func (*ConversionGoalCampaignConfigOperation_Update) isConversionGoalCampaignConfigOperation_Operation() {
}

// Response message for a conversion goal campaign config mutate.
type MutateConversionGoalCampaignConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All results for the mutate.
	Results []*MutateConversionGoalCampaignConfigResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateConversionGoalCampaignConfigsResponse) Reset() {
	*x = MutateConversionGoalCampaignConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateConversionGoalCampaignConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConversionGoalCampaignConfigsResponse) ProtoMessage() {}

func (x *MutateConversionGoalCampaignConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConversionGoalCampaignConfigsResponse.ProtoReflect.Descriptor instead.
func (*MutateConversionGoalCampaignConfigsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateConversionGoalCampaignConfigsResponse) GetResults() []*MutateConversionGoalCampaignConfigResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the conversion goal campaign config mutate.
type MutateConversionGoalCampaignConfigResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated ConversionGoalCampaignConfig with only mutable fields after
	// mutate. The field will only be returned when response_content_type is set
	// to "MUTABLE_RESOURCE".
	ConversionGoalCampaignConfig *resources.ConversionGoalCampaignConfig `protobuf:"bytes,2,opt,name=conversion_goal_campaign_config,json=conversionGoalCampaignConfig,proto3" json:"conversion_goal_campaign_config,omitempty"`
}

func (x *MutateConversionGoalCampaignConfigResult) Reset() {
	*x = MutateConversionGoalCampaignConfigResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateConversionGoalCampaignConfigResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConversionGoalCampaignConfigResult) ProtoMessage() {}

func (x *MutateConversionGoalCampaignConfigResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConversionGoalCampaignConfigResult.ProtoReflect.Descriptor instead.
func (*MutateConversionGoalCampaignConfigResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateConversionGoalCampaignConfigResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateConversionGoalCampaignConfigResult) GetConversionGoalCampaignConfig() *resources.ConversionGoalCampaignConfig {
	if x != nil {
		return x.ConversionGoalCampaignConfig
	}
	return nil
}

var File_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x6f, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe7, 0x02, 0x0a, 0x2a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x6d, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7f, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x25, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x5a, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x2b, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f,
	0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47,
	0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0x95, 0x02, 0x0a, 0x28, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5f, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f,
	0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x87,
	0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x6f,
	0x61, 0x6c, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x1c, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x9c, 0x03, 0x0a, 0x23, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xad, 0x02, 0x0a, 0x23, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61,
	0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48, 0x3a, 0x01, 0x2a, 0x22, 0x43, 0x2f, 0x76, 0x31,
	0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x94, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x42, 0x28, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x61,
	0x6c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescData = file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDesc
)

func file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDescData
}

var file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_goTypes = []interface{}{
	(*MutateConversionGoalCampaignConfigsRequest)(nil),     // 0: google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigsRequest
	(*ConversionGoalCampaignConfigOperation)(nil),          // 1: google.ads.googleads.v12.services.ConversionGoalCampaignConfigOperation
	(*MutateConversionGoalCampaignConfigsResponse)(nil),    // 2: google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigsResponse
	(*MutateConversionGoalCampaignConfigResult)(nil),       // 3: google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v12.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.ConversionGoalCampaignConfig)(nil),         // 6: google.ads.googleads.v12.resources.ConversionGoalCampaignConfig
}
var file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigsRequest.operations:type_name -> google.ads.googleads.v12.services.ConversionGoalCampaignConfigOperation
	4, // 1: google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigsRequest.response_content_type:type_name -> google.ads.googleads.v12.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v12.services.ConversionGoalCampaignConfigOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v12.services.ConversionGoalCampaignConfigOperation.update:type_name -> google.ads.googleads.v12.resources.ConversionGoalCampaignConfig
	3, // 4: google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigsResponse.results:type_name -> google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigResult
	6, // 5: google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigResult.conversion_goal_campaign_config:type_name -> google.ads.googleads.v12.resources.ConversionGoalCampaignConfig
	0, // 6: google.ads.googleads.v12.services.ConversionGoalCampaignConfigService.MutateConversionGoalCampaignConfigs:input_type -> google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigsRequest
	2, // 7: google.ads.googleads.v12.services.ConversionGoalCampaignConfigService.MutateConversionGoalCampaignConfigs:output_type -> google.ads.googleads.v12.services.MutateConversionGoalCampaignConfigsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_init()
}
func file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_init() {
	if File_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateConversionGoalCampaignConfigsRequest); i {
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
		file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionGoalCampaignConfigOperation); i {
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
		file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateConversionGoalCampaignConfigsResponse); i {
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
		file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateConversionGoalCampaignConfigResult); i {
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
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ConversionGoalCampaignConfigOperation_Update)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto = out.File
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_rawDesc = nil
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_goTypes = nil
	file_google_ads_googleads_v12_services_conversion_goal_campaign_config_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConversionGoalCampaignConfigServiceClient is the client API for ConversionGoalCampaignConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConversionGoalCampaignConfigServiceClient interface {
	// Creates, updates or removes conversion goal campaign config. Operation
	// statuses are returned.
	MutateConversionGoalCampaignConfigs(ctx context.Context, in *MutateConversionGoalCampaignConfigsRequest, opts ...grpc.CallOption) (*MutateConversionGoalCampaignConfigsResponse, error)
}

type conversionGoalCampaignConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionGoalCampaignConfigServiceClient(cc grpc.ClientConnInterface) ConversionGoalCampaignConfigServiceClient {
	return &conversionGoalCampaignConfigServiceClient{cc}
}

func (c *conversionGoalCampaignConfigServiceClient) MutateConversionGoalCampaignConfigs(ctx context.Context, in *MutateConversionGoalCampaignConfigsRequest, opts ...grpc.CallOption) (*MutateConversionGoalCampaignConfigsResponse, error) {
	out := new(MutateConversionGoalCampaignConfigsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.ConversionGoalCampaignConfigService/MutateConversionGoalCampaignConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionGoalCampaignConfigServiceServer is the server API for ConversionGoalCampaignConfigService service.
type ConversionGoalCampaignConfigServiceServer interface {
	// Creates, updates or removes conversion goal campaign config. Operation
	// statuses are returned.
	MutateConversionGoalCampaignConfigs(context.Context, *MutateConversionGoalCampaignConfigsRequest) (*MutateConversionGoalCampaignConfigsResponse, error)
}

// UnimplementedConversionGoalCampaignConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConversionGoalCampaignConfigServiceServer struct {
}

func (*UnimplementedConversionGoalCampaignConfigServiceServer) MutateConversionGoalCampaignConfigs(context.Context, *MutateConversionGoalCampaignConfigsRequest) (*MutateConversionGoalCampaignConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionGoalCampaignConfigs not implemented")
}

func RegisterConversionGoalCampaignConfigServiceServer(s *grpc.Server, srv ConversionGoalCampaignConfigServiceServer) {
	s.RegisterService(&_ConversionGoalCampaignConfigService_serviceDesc, srv)
}

func _ConversionGoalCampaignConfigService_MutateConversionGoalCampaignConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionGoalCampaignConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionGoalCampaignConfigServiceServer).MutateConversionGoalCampaignConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.ConversionGoalCampaignConfigService/MutateConversionGoalCampaignConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionGoalCampaignConfigServiceServer).MutateConversionGoalCampaignConfigs(ctx, req.(*MutateConversionGoalCampaignConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConversionGoalCampaignConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.ConversionGoalCampaignConfigService",
	HandlerType: (*ConversionGoalCampaignConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateConversionGoalCampaignConfigs",
			Handler:    _ConversionGoalCampaignConfigService_MutateConversionGoalCampaignConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/conversion_goal_campaign_config_service.proto",
}
