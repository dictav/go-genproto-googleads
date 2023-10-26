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
// source: google/ads/googleads/v15/services/payments_account_service.proto

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

// Request message for fetching all accessible payments accounts.
type ListPaymentsAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer to apply the PaymentsAccount list
	// operation to.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ListPaymentsAccountsRequest) Reset() {
	*x = ListPaymentsAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentsAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsAccountsRequest) ProtoMessage() {}

func (x *ListPaymentsAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsAccountsRequest.ProtoReflect.Descriptor instead.
func (*ListPaymentsAccountsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListPaymentsAccountsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

// Response message for
// [PaymentsAccountService.ListPaymentsAccounts][google.ads.googleads.v15.services.PaymentsAccountService.ListPaymentsAccounts].
type ListPaymentsAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of accessible payments accounts.
	PaymentsAccounts []*resources.PaymentsAccount `protobuf:"bytes,1,rep,name=payments_accounts,json=paymentsAccounts,proto3" json:"payments_accounts,omitempty"`
}

func (x *ListPaymentsAccountsResponse) Reset() {
	*x = ListPaymentsAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentsAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsAccountsResponse) ProtoMessage() {}

func (x *ListPaymentsAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsAccountsResponse.ProtoReflect.Descriptor instead.
func (*ListPaymentsAccountsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPaymentsAccountsResponse) GetPaymentsAccounts() []*resources.PaymentsAccount {
	if x != nil {
		return x.PaymentsAccounts
	}
	return nil
}

var File_google_ads_googleads_v15_services_payments_account_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_services_payments_account_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x80, 0x01,
	0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60,
	0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x10,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x32, 0xc0, 0x02, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xde, 0x01, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0xda, 0x41, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f, 0x76, 0x31,
	0x35, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x45, 0xca, 0x41,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x87, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1b, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x35, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescData = file_google_ads_googleads_v15_services_payments_account_service_proto_rawDesc
)

func file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_services_payments_account_service_proto_rawDescData
}

var file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v15_services_payments_account_service_proto_goTypes = []interface{}{
	(*ListPaymentsAccountsRequest)(nil),  // 0: google.ads.googleads.v15.services.ListPaymentsAccountsRequest
	(*ListPaymentsAccountsResponse)(nil), // 1: google.ads.googleads.v15.services.ListPaymentsAccountsResponse
	(*resources.PaymentsAccount)(nil),    // 2: google.ads.googleads.v15.resources.PaymentsAccount
}
var file_google_ads_googleads_v15_services_payments_account_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v15.services.ListPaymentsAccountsResponse.payments_accounts:type_name -> google.ads.googleads.v15.resources.PaymentsAccount
	0, // 1: google.ads.googleads.v15.services.PaymentsAccountService.ListPaymentsAccounts:input_type -> google.ads.googleads.v15.services.ListPaymentsAccountsRequest
	1, // 2: google.ads.googleads.v15.services.PaymentsAccountService.ListPaymentsAccounts:output_type -> google.ads.googleads.v15.services.ListPaymentsAccountsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_services_payments_account_service_proto_init() }
func file_google_ads_googleads_v15_services_payments_account_service_proto_init() {
	if File_google_ads_googleads_v15_services_payments_account_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentsAccountsRequest); i {
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
		file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentsAccountsResponse); i {
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
			RawDescriptor: file_google_ads_googleads_v15_services_payments_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v15_services_payments_account_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_services_payments_account_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v15_services_payments_account_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_services_payments_account_service_proto = out.File
	file_google_ads_googleads_v15_services_payments_account_service_proto_rawDesc = nil
	file_google_ads_googleads_v15_services_payments_account_service_proto_goTypes = nil
	file_google_ads_googleads_v15_services_payments_account_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentsAccountServiceClient is the client API for PaymentsAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentsAccountServiceClient interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PaymentsAccountError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error)
}

type paymentsAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsAccountServiceClient(cc grpc.ClientConnInterface) PaymentsAccountServiceClient {
	return &paymentsAccountServiceClient{cc}
}

func (c *paymentsAccountServiceClient) ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error) {
	out := new(ListPaymentsAccountsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v15.services.PaymentsAccountService/ListPaymentsAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsAccountServiceServer is the server API for PaymentsAccountService service.
type PaymentsAccountServiceServer interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [PaymentsAccountError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error)
}

// UnimplementedPaymentsAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentsAccountServiceServer struct {
}

func (*UnimplementedPaymentsAccountServiceServer) ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPaymentsAccounts not implemented")
}

func RegisterPaymentsAccountServiceServer(s *grpc.Server, srv PaymentsAccountServiceServer) {
	s.RegisterService(&_PaymentsAccountService_serviceDesc, srv)
}

func _PaymentsAccountService_ListPaymentsAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v15.services.PaymentsAccountService/ListPaymentsAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, req.(*ListPaymentsAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentsAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.PaymentsAccountService",
	HandlerType: (*PaymentsAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPaymentsAccounts",
			Handler:    _PaymentsAccountService_ListPaymentsAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/payments_account_service.proto",
}
