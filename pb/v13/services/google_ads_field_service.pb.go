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
// source: google/ads/googleads/v13/services/google_ads_field_service.proto

package services

import (
	context "context"
	resources "github.com/dictav/go-genproto-googleads/pb/v13/resources"
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

// Request message for
// [GoogleAdsFieldService.GetGoogleAdsField][google.ads.googleads.v13.services.GoogleAdsFieldService.GetGoogleAdsField].
type GetGoogleAdsFieldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the field to get.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetGoogleAdsFieldRequest) Reset() {
	*x = GetGoogleAdsFieldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoogleAdsFieldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoogleAdsFieldRequest) ProtoMessage() {}

func (x *GetGoogleAdsFieldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoogleAdsFieldRequest.ProtoReflect.Descriptor instead.
func (*GetGoogleAdsFieldRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetGoogleAdsFieldRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [GoogleAdsFieldService.SearchGoogleAdsFields][google.ads.googleads.v13.services.GoogleAdsFieldService.SearchGoogleAdsFields].
type SearchGoogleAdsFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The query string.
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Token of the page to retrieve. If not specified, the first page of
	// results will be returned. Use the value obtained from `next_page_token`
	// in the previous response in order to request the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Number of elements to retrieve in a single page.
	// When too large a page is requested, the server may decide to further
	// limit the number of returned resources.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *SearchGoogleAdsFieldsRequest) Reset() {
	*x = SearchGoogleAdsFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchGoogleAdsFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGoogleAdsFieldsRequest) ProtoMessage() {}

func (x *SearchGoogleAdsFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGoogleAdsFieldsRequest.ProtoReflect.Descriptor instead.
func (*SearchGoogleAdsFieldsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchGoogleAdsFieldsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchGoogleAdsFieldsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *SearchGoogleAdsFieldsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Response message for
// [GoogleAdsFieldService.SearchGoogleAdsFields][google.ads.googleads.v13.services.GoogleAdsFieldService.SearchGoogleAdsFields].
type SearchGoogleAdsFieldsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of fields that matched the query.
	Results []*resources.GoogleAdsField `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// Pagination token used to retrieve the next page of results. Pass the
	// content of this string as the `page_token` attribute of the next request.
	// `next_page_token` is not returned for the last page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total number of results that match the query ignoring the LIMIT clause.
	TotalResultsCount int64 `protobuf:"varint,3,opt,name=total_results_count,json=totalResultsCount,proto3" json:"total_results_count,omitempty"`
}

func (x *SearchGoogleAdsFieldsResponse) Reset() {
	*x = SearchGoogleAdsFieldsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchGoogleAdsFieldsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGoogleAdsFieldsResponse) ProtoMessage() {}

func (x *SearchGoogleAdsFieldsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGoogleAdsFieldsResponse.ProtoReflect.Descriptor instead.
func (*SearchGoogleAdsFieldsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescGZIP(), []int{2}
}

func (x *SearchGoogleAdsFieldsResponse) GetResults() []*resources.GoogleAdsField {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SearchGoogleAdsFieldsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *SearchGoogleAdsFieldsResponse) GetTotalResultsCount() int64 {
	if x != nil {
		return x.TotalResultsCount
	}
	return 0
}

var File_google_ads_googleads_v13_services_google_ads_field_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x75, 0x0a, 0x1c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xc5, 0x01, 0x0a,
	0x1d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xf2, 0x03, 0x0a, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc4,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x3e, 0xda, 0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f,
	0x76, 0x31, 0x33, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xca, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2e, 0xda, 0x41, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x86, 0x02, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x42, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x33, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescData = file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDesc
)

func file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDescData
}

var file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v13_services_google_ads_field_service_proto_goTypes = []interface{}{
	(*GetGoogleAdsFieldRequest)(nil),      // 0: google.ads.googleads.v13.services.GetGoogleAdsFieldRequest
	(*SearchGoogleAdsFieldsRequest)(nil),  // 1: google.ads.googleads.v13.services.SearchGoogleAdsFieldsRequest
	(*SearchGoogleAdsFieldsResponse)(nil), // 2: google.ads.googleads.v13.services.SearchGoogleAdsFieldsResponse
	(*resources.GoogleAdsField)(nil),      // 3: google.ads.googleads.v13.resources.GoogleAdsField
}
var file_google_ads_googleads_v13_services_google_ads_field_service_proto_depIdxs = []int32{
	3, // 0: google.ads.googleads.v13.services.SearchGoogleAdsFieldsResponse.results:type_name -> google.ads.googleads.v13.resources.GoogleAdsField
	0, // 1: google.ads.googleads.v13.services.GoogleAdsFieldService.GetGoogleAdsField:input_type -> google.ads.googleads.v13.services.GetGoogleAdsFieldRequest
	1, // 2: google.ads.googleads.v13.services.GoogleAdsFieldService.SearchGoogleAdsFields:input_type -> google.ads.googleads.v13.services.SearchGoogleAdsFieldsRequest
	3, // 3: google.ads.googleads.v13.services.GoogleAdsFieldService.GetGoogleAdsField:output_type -> google.ads.googleads.v13.resources.GoogleAdsField
	2, // 4: google.ads.googleads.v13.services.GoogleAdsFieldService.SearchGoogleAdsFields:output_type -> google.ads.googleads.v13.services.SearchGoogleAdsFieldsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_services_google_ads_field_service_proto_init() }
func file_google_ads_googleads_v13_services_google_ads_field_service_proto_init() {
	if File_google_ads_googleads_v13_services_google_ads_field_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoogleAdsFieldRequest); i {
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
		file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchGoogleAdsFieldsRequest); i {
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
		file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchGoogleAdsFieldsResponse); i {
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
			RawDescriptor: file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v13_services_google_ads_field_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_services_google_ads_field_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v13_services_google_ads_field_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_services_google_ads_field_service_proto = out.File
	file_google_ads_googleads_v13_services_google_ads_field_service_proto_rawDesc = nil
	file_google_ads_googleads_v13_services_google_ads_field_service_proto_goTypes = nil
	file_google_ads_googleads_v13_services_google_ads_field_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GoogleAdsFieldServiceClient is the client API for GoogleAdsFieldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoogleAdsFieldServiceClient interface {
	// Returns just the requested field.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetGoogleAdsField(ctx context.Context, in *GetGoogleAdsFieldRequest, opts ...grpc.CallOption) (*resources.GoogleAdsField, error)
	// Returns all fields that match the search query.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QueryError]()
	//   [QuotaError]()
	//   [RequestError]()
	SearchGoogleAdsFields(ctx context.Context, in *SearchGoogleAdsFieldsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsFieldsResponse, error)
}

type googleAdsFieldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoogleAdsFieldServiceClient(cc grpc.ClientConnInterface) GoogleAdsFieldServiceClient {
	return &googleAdsFieldServiceClient{cc}
}

func (c *googleAdsFieldServiceClient) GetGoogleAdsField(ctx context.Context, in *GetGoogleAdsFieldRequest, opts ...grpc.CallOption) (*resources.GoogleAdsField, error) {
	out := new(resources.GoogleAdsField)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.GoogleAdsFieldService/GetGoogleAdsField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleAdsFieldServiceClient) SearchGoogleAdsFields(ctx context.Context, in *SearchGoogleAdsFieldsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsFieldsResponse, error) {
	out := new(SearchGoogleAdsFieldsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.GoogleAdsFieldService/SearchGoogleAdsFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoogleAdsFieldServiceServer is the server API for GoogleAdsFieldService service.
type GoogleAdsFieldServiceServer interface {
	// Returns just the requested field.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetGoogleAdsField(context.Context, *GetGoogleAdsFieldRequest) (*resources.GoogleAdsField, error)
	// Returns all fields that match the search query.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QueryError]()
	//   [QuotaError]()
	//   [RequestError]()
	SearchGoogleAdsFields(context.Context, *SearchGoogleAdsFieldsRequest) (*SearchGoogleAdsFieldsResponse, error)
}

// UnimplementedGoogleAdsFieldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGoogleAdsFieldServiceServer struct {
}

func (*UnimplementedGoogleAdsFieldServiceServer) GetGoogleAdsField(context.Context, *GetGoogleAdsFieldRequest) (*resources.GoogleAdsField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoogleAdsField not implemented")
}
func (*UnimplementedGoogleAdsFieldServiceServer) SearchGoogleAdsFields(context.Context, *SearchGoogleAdsFieldsRequest) (*SearchGoogleAdsFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGoogleAdsFields not implemented")
}

func RegisterGoogleAdsFieldServiceServer(s *grpc.Server, srv GoogleAdsFieldServiceServer) {
	s.RegisterService(&_GoogleAdsFieldService_serviceDesc, srv)
}

func _GoogleAdsFieldService_GetGoogleAdsField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoogleAdsFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsFieldServiceServer).GetGoogleAdsField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.GoogleAdsFieldService/GetGoogleAdsField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsFieldServiceServer).GetGoogleAdsField(ctx, req.(*GetGoogleAdsFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleAdsFieldService_SearchGoogleAdsFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGoogleAdsFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsFieldServiceServer).SearchGoogleAdsFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.GoogleAdsFieldService/SearchGoogleAdsFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsFieldServiceServer).SearchGoogleAdsFields(ctx, req.(*SearchGoogleAdsFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoogleAdsFieldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.GoogleAdsFieldService",
	HandlerType: (*GoogleAdsFieldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoogleAdsField",
			Handler:    _GoogleAdsFieldService_GetGoogleAdsField_Handler,
		},
		{
			MethodName: "SearchGoogleAdsFields",
			Handler:    _GoogleAdsFieldService_SearchGoogleAdsFields_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/google_ads_field_service.proto",
}