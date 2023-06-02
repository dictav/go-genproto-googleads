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
// source: google/ads/googleads/v12/services/billing_setup_service.proto

package services

import (
	context "context"
	resources "github.com/dictav/go-genproto-googleads/pb/v12/resources"
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

// Request message for billing setup mutate operations.
type MutateBillingSetupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Id of the customer to apply the billing setup mutate operation to.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform.
	Operation *BillingSetupOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *MutateBillingSetupRequest) Reset() {
	*x = MutateBillingSetupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateBillingSetupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBillingSetupRequest) ProtoMessage() {}

func (x *MutateBillingSetupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBillingSetupRequest.ProtoReflect.Descriptor instead.
func (*MutateBillingSetupRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateBillingSetupRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateBillingSetupRequest) GetOperation() *BillingSetupOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

// A single operation on a billing setup, which describes the cancellation of an
// existing billing setup.
type BillingSetupOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only one of these operations can be set. "Update" operations are not
	// supported.
	//
	// Types that are assignable to Operation:
	//	*BillingSetupOperation_Create
	//	*BillingSetupOperation_Remove
	Operation isBillingSetupOperation_Operation `protobuf_oneof:"operation"`
}

func (x *BillingSetupOperation) Reset() {
	*x = BillingSetupOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingSetupOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingSetupOperation) ProtoMessage() {}

func (x *BillingSetupOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingSetupOperation.ProtoReflect.Descriptor instead.
func (*BillingSetupOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescGZIP(), []int{1}
}

func (m *BillingSetupOperation) GetOperation() isBillingSetupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *BillingSetupOperation) GetCreate() *resources.BillingSetup {
	if x, ok := x.GetOperation().(*BillingSetupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *BillingSetupOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*BillingSetupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isBillingSetupOperation_Operation interface {
	isBillingSetupOperation_Operation()
}

type BillingSetupOperation_Create struct {
	// Creates a billing setup. No resource name is expected for the new billing
	// setup.
	Create *resources.BillingSetup `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type BillingSetupOperation_Remove struct {
	// Resource name of the billing setup to remove. A setup cannot be
	// removed unless it is in a pending state or its scheduled start time is in
	// the future. The resource name looks like
	// `customers/{customer_id}/billingSetups/{billing_id}`.
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

func (*BillingSetupOperation_Create) isBillingSetupOperation_Operation() {}

func (*BillingSetupOperation_Remove) isBillingSetupOperation_Operation() {}

// Response message for a billing setup operation.
type MutateBillingSetupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A result that identifies the resource affected by the mutate request.
	Result *MutateBillingSetupResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MutateBillingSetupResponse) Reset() {
	*x = MutateBillingSetupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateBillingSetupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBillingSetupResponse) ProtoMessage() {}

func (x *MutateBillingSetupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBillingSetupResponse.ProtoReflect.Descriptor instead.
func (*MutateBillingSetupResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateBillingSetupResponse) GetResult() *MutateBillingSetupResult {
	if x != nil {
		return x.Result
	}
	return nil
}

// Result for a single billing setup mutate.
type MutateBillingSetupResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateBillingSetupResult) Reset() {
	*x = MutateBillingSetupResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateBillingSetupResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBillingSetupResult) ProtoMessage() {}

func (x *MutateBillingSetupResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBillingSetupResult.ProtoReflect.Descriptor instead.
func (*MutateBillingSetupResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateBillingSetupResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v12_services_billing_setup_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x65, 0x74, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01,
	0x0a, 0x19, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x5b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x74, 0x75, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb6,
	0x01, 0x0a, 0x15, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x1a, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6b, 0x0a, 0x18, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xfa,
	0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xc8, 0x02, 0x0a, 0x13, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xe9, 0x01, 0x0a, 0x12, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x56, 0xda, 0x41, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x38, 0x3a, 0x01, 0x2a, 0x22, 0x33, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x74, 0x75, 0x70, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41, 0x18,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x84, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x18, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescData = file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDesc
)

func file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDescData
}

var file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v12_services_billing_setup_service_proto_goTypes = []interface{}{
	(*MutateBillingSetupRequest)(nil),  // 0: google.ads.googleads.v12.services.MutateBillingSetupRequest
	(*BillingSetupOperation)(nil),      // 1: google.ads.googleads.v12.services.BillingSetupOperation
	(*MutateBillingSetupResponse)(nil), // 2: google.ads.googleads.v12.services.MutateBillingSetupResponse
	(*MutateBillingSetupResult)(nil),   // 3: google.ads.googleads.v12.services.MutateBillingSetupResult
	(*resources.BillingSetup)(nil),     // 4: google.ads.googleads.v12.resources.BillingSetup
}
var file_google_ads_googleads_v12_services_billing_setup_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v12.services.MutateBillingSetupRequest.operation:type_name -> google.ads.googleads.v12.services.BillingSetupOperation
	4, // 1: google.ads.googleads.v12.services.BillingSetupOperation.create:type_name -> google.ads.googleads.v12.resources.BillingSetup
	3, // 2: google.ads.googleads.v12.services.MutateBillingSetupResponse.result:type_name -> google.ads.googleads.v12.services.MutateBillingSetupResult
	0, // 3: google.ads.googleads.v12.services.BillingSetupService.MutateBillingSetup:input_type -> google.ads.googleads.v12.services.MutateBillingSetupRequest
	2, // 4: google.ads.googleads.v12.services.BillingSetupService.MutateBillingSetup:output_type -> google.ads.googleads.v12.services.MutateBillingSetupResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_services_billing_setup_service_proto_init() }
func file_google_ads_googleads_v12_services_billing_setup_service_proto_init() {
	if File_google_ads_googleads_v12_services_billing_setup_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateBillingSetupRequest); i {
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
		file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingSetupOperation); i {
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
		file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateBillingSetupResponse); i {
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
		file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateBillingSetupResult); i {
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
	file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BillingSetupOperation_Create)(nil),
		(*BillingSetupOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v12_services_billing_setup_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_services_billing_setup_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_services_billing_setup_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_services_billing_setup_service_proto = out.File
	file_google_ads_googleads_v12_services_billing_setup_service_proto_rawDesc = nil
	file_google_ads_googleads_v12_services_billing_setup_service_proto_goTypes = nil
	file_google_ads_googleads_v12_services_billing_setup_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BillingSetupServiceClient is the client API for BillingSetupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BillingSetupServiceClient interface {
	// Creates a billing setup, or cancels an existing billing setup.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BillingSetupError]()
	//   [DateError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error)
}

type billingSetupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingSetupServiceClient(cc grpc.ClientConnInterface) BillingSetupServiceClient {
	return &billingSetupServiceClient{cc}
}

func (c *billingSetupServiceClient) MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error) {
	out := new(MutateBillingSetupResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.BillingSetupService/MutateBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingSetupServiceServer is the server API for BillingSetupService service.
type BillingSetupServiceServer interface {
	// Creates a billing setup, or cancels an existing billing setup.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BillingSetupError]()
	//   [DateError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error)
}

// UnimplementedBillingSetupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBillingSetupServiceServer struct {
}

func (*UnimplementedBillingSetupServiceServer) MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBillingSetup not implemented")
}

func RegisterBillingSetupServiceServer(s *grpc.Server, srv BillingSetupServiceServer) {
	s.RegisterService(&_BillingSetupService_serviceDesc, srv)
}

func _BillingSetupService_MutateBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.BillingSetupService/MutateBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, req.(*MutateBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BillingSetupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.BillingSetupService",
	HandlerType: (*BillingSetupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBillingSetup",
			Handler:    _BillingSetupService_MutateBillingSetup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/billing_setup_service.proto",
}
