// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: google/ads/googleads/v7/services/customer_user_access_service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/dictav/go-genproto-googleads/pb/v7/resources"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for
// [CustomerUserAccessService.GetCustomerUserAccess][google.ads.googleads.v7.services.CustomerUserAccessService.GetCustomerUserAccess].
type GetCustomerUserAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the customer user access.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetCustomerUserAccessRequest) Reset() {
	*x = GetCustomerUserAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerUserAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerUserAccessRequest) ProtoMessage() {}

func (x *GetCustomerUserAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerUserAccessRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerUserAccessRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerUserAccessRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Mutate Request for
// [CustomerUserAccessService.MutateCustomerUserAccess][google.ads.googleads.v7.services.CustomerUserAccessService.MutateCustomerUserAccess].
type MutateCustomerUserAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform on the customer
	Operation *CustomerUserAccessOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *MutateCustomerUserAccessRequest) Reset() {
	*x = MutateCustomerUserAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessRequest) ProtoMessage() {}

func (x *MutateCustomerUserAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerUserAccessRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerUserAccessRequest) GetOperation() *CustomerUserAccessOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

// A single operation (update, remove) on customer user access.
type CustomerUserAccessOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*CustomerUserAccessOperation_Update
	//	*CustomerUserAccessOperation_Remove
	Operation isCustomerUserAccessOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CustomerUserAccessOperation) Reset() {
	*x = CustomerUserAccessOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerUserAccessOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUserAccessOperation) ProtoMessage() {}

func (x *CustomerUserAccessOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUserAccessOperation.ProtoReflect.Descriptor instead.
func (*CustomerUserAccessOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerUserAccessOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *CustomerUserAccessOperation) GetOperation() isCustomerUserAccessOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CustomerUserAccessOperation) GetUpdate() *resources.CustomerUserAccess {
	if x, ok := x.GetOperation().(*CustomerUserAccessOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *CustomerUserAccessOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*CustomerUserAccessOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isCustomerUserAccessOperation_Operation interface {
	isCustomerUserAccessOperation_Operation()
}

type CustomerUserAccessOperation_Update struct {
	// Update operation: The customer user access is expected to have a valid
	// resource name.
	Update *resources.CustomerUserAccess `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type CustomerUserAccessOperation_Remove struct {
	// Remove operation: A resource name for the removed access is
	// expected, in this format:
	//
	// `customers/{customer_id}/customerUserAccesses/{CustomerUserAccess.user_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerUserAccessOperation_Update) isCustomerUserAccessOperation_Operation() {}

func (*CustomerUserAccessOperation_Remove) isCustomerUserAccessOperation_Operation() {}

// Response message for customer user access mutate.
type MutateCustomerUserAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result for the mutate.
	Result *MutateCustomerUserAccessResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MutateCustomerUserAccessResponse) Reset() {
	*x = MutateCustomerUserAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessResponse) ProtoMessage() {}

func (x *MutateCustomerUserAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerUserAccessResponse) GetResult() *MutateCustomerUserAccessResult {
	if x != nil {
		return x.Result
	}
	return nil
}

// The result for the customer user access mutate.
type MutateCustomerUserAccessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCustomerUserAccessResult) Reset() {
	*x = MutateCustomerUserAccessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessResult) ProtoMessage() {}

func (x *MutateCustomerUserAccessResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateCustomerUserAccessResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v7_services_customer_user_access_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x1f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x60,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xd2, 0x01, 0x0a, 0x1b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x4f, 0x0a,
	0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x20, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x45, 0x0a, 0x1e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xc5, 0x04, 0x0a, 0x19, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xde, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x38, 0x12, 0x36, 0x2f, 0x76, 0x37, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x2a, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xff, 0x01, 0x0a, 0x18, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x22, 0x39, 0x2f, 0x76, 0x37, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0xda, 0x41, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x45, 0xca, 0x41, 0x18,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x85, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x20,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x37, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x37, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x37,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescData = file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDesc
)

func file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDescData
}

var file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v7_services_customer_user_access_service_proto_goTypes = []interface{}{
	(*GetCustomerUserAccessRequest)(nil),     // 0: google.ads.googleads.v7.services.GetCustomerUserAccessRequest
	(*MutateCustomerUserAccessRequest)(nil),  // 1: google.ads.googleads.v7.services.MutateCustomerUserAccessRequest
	(*CustomerUserAccessOperation)(nil),      // 2: google.ads.googleads.v7.services.CustomerUserAccessOperation
	(*MutateCustomerUserAccessResponse)(nil), // 3: google.ads.googleads.v7.services.MutateCustomerUserAccessResponse
	(*MutateCustomerUserAccessResult)(nil),   // 4: google.ads.googleads.v7.services.MutateCustomerUserAccessResult
	(*fieldmaskpb.FieldMask)(nil),            // 5: google.protobuf.FieldMask
	(*resources.CustomerUserAccess)(nil),     // 6: google.ads.googleads.v7.resources.CustomerUserAccess
}
var file_google_ads_googleads_v7_services_customer_user_access_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v7.services.MutateCustomerUserAccessRequest.operation:type_name -> google.ads.googleads.v7.services.CustomerUserAccessOperation
	5, // 1: google.ads.googleads.v7.services.CustomerUserAccessOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.ads.googleads.v7.services.CustomerUserAccessOperation.update:type_name -> google.ads.googleads.v7.resources.CustomerUserAccess
	4, // 3: google.ads.googleads.v7.services.MutateCustomerUserAccessResponse.result:type_name -> google.ads.googleads.v7.services.MutateCustomerUserAccessResult
	0, // 4: google.ads.googleads.v7.services.CustomerUserAccessService.GetCustomerUserAccess:input_type -> google.ads.googleads.v7.services.GetCustomerUserAccessRequest
	1, // 5: google.ads.googleads.v7.services.CustomerUserAccessService.MutateCustomerUserAccess:input_type -> google.ads.googleads.v7.services.MutateCustomerUserAccessRequest
	6, // 6: google.ads.googleads.v7.services.CustomerUserAccessService.GetCustomerUserAccess:output_type -> google.ads.googleads.v7.resources.CustomerUserAccess
	3, // 7: google.ads.googleads.v7.services.CustomerUserAccessService.MutateCustomerUserAccess:output_type -> google.ads.googleads.v7.services.MutateCustomerUserAccessResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v7_services_customer_user_access_service_proto_init() }
func file_google_ads_googleads_v7_services_customer_user_access_service_proto_init() {
	if File_google_ads_googleads_v7_services_customer_user_access_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerUserAccessRequest); i {
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
		file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessRequest); i {
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
		file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerUserAccessOperation); i {
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
		file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessResponse); i {
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
		file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessResult); i {
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
	file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CustomerUserAccessOperation_Update)(nil),
		(*CustomerUserAccessOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v7_services_customer_user_access_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v7_services_customer_user_access_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v7_services_customer_user_access_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v7_services_customer_user_access_service_proto = out.File
	file_google_ads_googleads_v7_services_customer_user_access_service_proto_rawDesc = nil
	file_google_ads_googleads_v7_services_customer_user_access_service_proto_goTypes = nil
	file_google_ads_googleads_v7_services_customer_user_access_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerUserAccessServiceClient is the client API for CustomerUserAccessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerUserAccessServiceClient interface {
	// Returns the CustomerUserAccess in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerUserAccess(ctx context.Context, in *GetCustomerUserAccessRequest, opts ...grpc.CallOption) (*resources.CustomerUserAccess, error)
	// Updates, removes permission of a user on a given customer. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CustomerUserAccessError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerUserAccess(ctx context.Context, in *MutateCustomerUserAccessRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessResponse, error)
}

type customerUserAccessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerUserAccessServiceClient(cc grpc.ClientConnInterface) CustomerUserAccessServiceClient {
	return &customerUserAccessServiceClient{cc}
}

func (c *customerUserAccessServiceClient) GetCustomerUserAccess(ctx context.Context, in *GetCustomerUserAccessRequest, opts ...grpc.CallOption) (*resources.CustomerUserAccess, error) {
	out := new(resources.CustomerUserAccess)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.CustomerUserAccessService/GetCustomerUserAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerUserAccessServiceClient) MutateCustomerUserAccess(ctx context.Context, in *MutateCustomerUserAccessRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessResponse, error) {
	out := new(MutateCustomerUserAccessResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.CustomerUserAccessService/MutateCustomerUserAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerUserAccessServiceServer is the server API for CustomerUserAccessService service.
type CustomerUserAccessServiceServer interface {
	// Returns the CustomerUserAccess in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerUserAccess(context.Context, *GetCustomerUserAccessRequest) (*resources.CustomerUserAccess, error)
	// Updates, removes permission of a user on a given customer. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CustomerUserAccessError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerUserAccess(context.Context, *MutateCustomerUserAccessRequest) (*MutateCustomerUserAccessResponse, error)
}

// UnimplementedCustomerUserAccessServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerUserAccessServiceServer struct {
}

func (*UnimplementedCustomerUserAccessServiceServer) GetCustomerUserAccess(context.Context, *GetCustomerUserAccessRequest) (*resources.CustomerUserAccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerUserAccess not implemented")
}
func (*UnimplementedCustomerUserAccessServiceServer) MutateCustomerUserAccess(context.Context, *MutateCustomerUserAccessRequest) (*MutateCustomerUserAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerUserAccess not implemented")
}

func RegisterCustomerUserAccessServiceServer(s *grpc.Server, srv CustomerUserAccessServiceServer) {
	s.RegisterService(&_CustomerUserAccessService_serviceDesc, srv)
}

func _CustomerUserAccessService_GetCustomerUserAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerUserAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessServiceServer).GetCustomerUserAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.CustomerUserAccessService/GetCustomerUserAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessServiceServer).GetCustomerUserAccess(ctx, req.(*GetCustomerUserAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerUserAccessService_MutateCustomerUserAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerUserAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessServiceServer).MutateCustomerUserAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.CustomerUserAccessService/MutateCustomerUserAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessServiceServer).MutateCustomerUserAccess(ctx, req.(*MutateCustomerUserAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerUserAccessService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.CustomerUserAccessService",
	HandlerType: (*CustomerUserAccessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerUserAccess",
			Handler:    _CustomerUserAccessService_GetCustomerUserAccess_Handler,
		},
		{
			MethodName: "MutateCustomerUserAccess",
			Handler:    _CustomerUserAccessService_MutateCustomerUserAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/customer_user_access_service.proto",
}
