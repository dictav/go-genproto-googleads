// Copyright 2023 Google LLC
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
// source: google/ads/googleads/v15/services/customer_sk_ad_network_conversion_value_schema_service.proto

package services

import (
	context "context"
	resources "github.com/dictav/go-genproto-googleads/pb/v15/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// A single update operation for a CustomerSkAdNetworkConversionValueSchema.
type CustomerSkAdNetworkConversionValueSchemaOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Update operation: The schema is expected to have a valid resource name.
	Update *resources.CustomerSkAdNetworkConversionValueSchema `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) Reset() {
	*x = CustomerSkAdNetworkConversionValueSchemaOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerSkAdNetworkConversionValueSchemaOperation) ProtoMessage() {}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerSkAdNetworkConversionValueSchemaOperation.ProtoReflect.Descriptor instead.
func (*CustomerSkAdNetworkConversionValueSchemaOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) GetUpdate() *resources.CustomerSkAdNetworkConversionValueSchema {
	if x != nil {
		return x.Update
	}
	return nil
}

// Request message for
// [CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema][google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema].
type MutateCustomerSkAdNetworkConversionValueSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the customer whose shared sets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform.
	Operation *CustomerSkAdNetworkConversionValueSchemaOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaRequest) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetOperation() *CustomerSkAdNetworkConversionValueSchemaOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// The result for the CustomerSkAdNetworkConversionValueSchema mutate.
type MutateCustomerSkAdNetworkConversionValueSchemaResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the customer that was modified.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// App ID of the SkanConversionValue modified.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaResult) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

// Response message for MutateCustomerSkAdNetworkConversionValueSchema.
type MutateCustomerSkAdNetworkConversionValueSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All results for the mutate.
	Result *MutateCustomerSkAdNetworkConversionValueSchemaResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaResponse) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) GetResult() *MutateCustomerSkAdNetworkConversionValueSchemaResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x6b, 0x5f,
	0x61, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x57, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x73, 0x6b, 0x5f, 0x61, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99,
	0x01, 0x0a, 0x31, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x64, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x35, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41,
	0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x72, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xba,
	0x01, 0x0a, 0x34, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x6b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x46,
	0xfa, 0x41, 0x43, 0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x36,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b,
	0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x57, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xbc, 0x03, 0x0a, 0x2f, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc1, 0x02, 0x0a, 0x2e,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b,
	0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x58,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x59, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x54, 0x3a, 0x01, 0x2a, 0x22, 0x4f,
	0x2f, 0x76, 0x31, 0x35, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x1a,
	0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61,
	0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0xa0, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x42, 0x34, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x35, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData = file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc
)

func file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData
}

var file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes = []interface{}{
	(*CustomerSkAdNetworkConversionValueSchemaOperation)(nil),      // 0: google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaOperation
	(*MutateCustomerSkAdNetworkConversionValueSchemaRequest)(nil),  // 1: google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest
	(*MutateCustomerSkAdNetworkConversionValueSchemaResult)(nil),   // 2: google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaResult
	(*MutateCustomerSkAdNetworkConversionValueSchemaResponse)(nil), // 3: google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse
	(*resources.CustomerSkAdNetworkConversionValueSchema)(nil),     // 4: google.ads.googleads.v15.resources.CustomerSkAdNetworkConversionValueSchema
}
var file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs = []int32{
	4, // 0: google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaOperation.update:type_name -> google.ads.googleads.v15.resources.CustomerSkAdNetworkConversionValueSchema
	0, // 1: google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest.operation:type_name -> google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaOperation
	2, // 2: google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse.result:type_name -> google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaResult
	1, // 3: google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema:input_type -> google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest
	3, // 4: google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema:output_type -> google.ads.googleads.v15.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_init()
}
func file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_init() {
	if File_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerSkAdNetworkConversionValueSchemaOperation); i {
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
		file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerSkAdNetworkConversionValueSchemaRequest); i {
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
		file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerSkAdNetworkConversionValueSchemaResult); i {
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
		file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerSkAdNetworkConversionValueSchemaResponse); i {
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
			RawDescriptor: file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto = out.File
	file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc = nil
	file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes = nil
	file_google_ads_googleads_v15_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerSkAdNetworkConversionValueSchemaServiceClient is the client API for CustomerSkAdNetworkConversionValueSchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerSkAdNetworkConversionValueSchemaServiceClient interface {
	// Creates or updates the CustomerSkAdNetworkConversionValueSchema.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [InternalError]()
	//   [MutateError]()
	MutateCustomerSkAdNetworkConversionValueSchema(ctx context.Context, in *MutateCustomerSkAdNetworkConversionValueSchemaRequest, opts ...grpc.CallOption) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error)
}

type customerSkAdNetworkConversionValueSchemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerSkAdNetworkConversionValueSchemaServiceClient(cc grpc.ClientConnInterface) CustomerSkAdNetworkConversionValueSchemaServiceClient {
	return &customerSkAdNetworkConversionValueSchemaServiceClient{cc}
}

func (c *customerSkAdNetworkConversionValueSchemaServiceClient) MutateCustomerSkAdNetworkConversionValueSchema(ctx context.Context, in *MutateCustomerSkAdNetworkConversionValueSchemaRequest, opts ...grpc.CallOption) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error) {
	out := new(MutateCustomerSkAdNetworkConversionValueSchemaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaService/MutateCustomerSkAdNetworkConversionValueSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerSkAdNetworkConversionValueSchemaServiceServer is the server API for CustomerSkAdNetworkConversionValueSchemaService service.
type CustomerSkAdNetworkConversionValueSchemaServiceServer interface {
	// Creates or updates the CustomerSkAdNetworkConversionValueSchema.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [InternalError]()
	//   [MutateError]()
	MutateCustomerSkAdNetworkConversionValueSchema(context.Context, *MutateCustomerSkAdNetworkConversionValueSchemaRequest) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error)
}

// UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer struct {
}

func (*UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer) MutateCustomerSkAdNetworkConversionValueSchema(context.Context, *MutateCustomerSkAdNetworkConversionValueSchemaRequest) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerSkAdNetworkConversionValueSchema not implemented")
}

func RegisterCustomerSkAdNetworkConversionValueSchemaServiceServer(s *grpc.Server, srv CustomerSkAdNetworkConversionValueSchemaServiceServer) {
	s.RegisterService(&_CustomerSkAdNetworkConversionValueSchemaService_serviceDesc, srv)
}

func _CustomerSkAdNetworkConversionValueSchemaService_MutateCustomerSkAdNetworkConversionValueSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerSkAdNetworkConversionValueSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerSkAdNetworkConversionValueSchemaServiceServer).MutateCustomerSkAdNetworkConversionValueSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaService/MutateCustomerSkAdNetworkConversionValueSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerSkAdNetworkConversionValueSchemaServiceServer).MutateCustomerSkAdNetworkConversionValueSchema(ctx, req.(*MutateCustomerSkAdNetworkConversionValueSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerSkAdNetworkConversionValueSchemaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.CustomerSkAdNetworkConversionValueSchemaService",
	HandlerType: (*CustomerSkAdNetworkConversionValueSchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerSkAdNetworkConversionValueSchema",
			Handler:    _CustomerSkAdNetworkConversionValueSchemaService_MutateCustomerSkAdNetworkConversionValueSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/customer_sk_ad_network_conversion_value_schema_service.proto",
}
