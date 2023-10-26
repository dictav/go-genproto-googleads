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
// source: google/ads/googleads/v14/common/feed_item_set_filter_type_infos.proto

package common

import (
	enums "github.com/dictav/go-genproto-googleads/pb/v14/enums"
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

// Represents a filter on locations in a feed item set.
// Only applicable if the parent Feed of the FeedItemSet is a LOCATION feed.
type DynamicLocationSetFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If multiple labels are set, then only feeditems marked with all the labels
	// will be added to the FeedItemSet.
	Labels []string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	// Business name filter.
	BusinessNameFilter *BusinessNameFilter `protobuf:"bytes,2,opt,name=business_name_filter,json=businessNameFilter,proto3" json:"business_name_filter,omitempty"`
}

func (x *DynamicLocationSetFilter) Reset() {
	*x = DynamicLocationSetFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicLocationSetFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicLocationSetFilter) ProtoMessage() {}

func (x *DynamicLocationSetFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicLocationSetFilter.ProtoReflect.Descriptor instead.
func (*DynamicLocationSetFilter) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescGZIP(), []int{0}
}

func (x *DynamicLocationSetFilter) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *DynamicLocationSetFilter) GetBusinessNameFilter() *BusinessNameFilter {
	if x != nil {
		return x.BusinessNameFilter
	}
	return nil
}

// Represents a business name filter on locations in a FeedItemSet.
type BusinessNameFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Business name string to use for filtering.
	BusinessName string `protobuf:"bytes,1,opt,name=business_name,json=businessName,proto3" json:"business_name,omitempty"`
	// The type of string matching to use when filtering with business_name.
	FilterType enums.FeedItemSetStringFilterTypeEnum_FeedItemSetStringFilterType `protobuf:"varint,2,opt,name=filter_type,json=filterType,proto3,enum=google.ads.googleads.v14.enums.FeedItemSetStringFilterTypeEnum_FeedItemSetStringFilterType" json:"filter_type,omitempty"`
}

func (x *BusinessNameFilter) Reset() {
	*x = BusinessNameFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessNameFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessNameFilter) ProtoMessage() {}

func (x *BusinessNameFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessNameFilter.ProtoReflect.Descriptor instead.
func (*BusinessNameFilter) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescGZIP(), []int{1}
}

func (x *BusinessNameFilter) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *BusinessNameFilter) GetFilterType() enums.FeedItemSetStringFilterTypeEnum_FeedItemSetStringFilterType {
	if x != nil {
		return x.FilterType
	}
	return enums.FeedItemSetStringFilterTypeEnum_FeedItemSetStringFilterType(0)
}

// Represents a filter on affiliate locations in a FeedItemSet.
// Only applicable if the parent Feed of the FeedItemSet is an
// AFFILIATE_LOCATION feed.
type DynamicAffiliateLocationSetFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to filter affiliate locations by chain ids. Only affiliate locations
	// that belong to the specified chain(s) will be added to the FeedItemSet.
	ChainIds []int64 `protobuf:"varint,1,rep,packed,name=chain_ids,json=chainIds,proto3" json:"chain_ids,omitempty"`
}

func (x *DynamicAffiliateLocationSetFilter) Reset() {
	*x = DynamicAffiliateLocationSetFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicAffiliateLocationSetFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicAffiliateLocationSetFilter) ProtoMessage() {}

func (x *DynamicAffiliateLocationSetFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicAffiliateLocationSetFilter.ProtoReflect.Descriptor instead.
func (*DynamicAffiliateLocationSetFilter) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescGZIP(), []int{2}
}

func (x *DynamicAffiliateLocationSetFilter) GetChainIds() []int64 {
	if x != nil {
		return x.ChainIds
	}
	return nil
}

var File_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x99, 0x01, 0x0a, 0x18, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x65, 0x0a, 0x14, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x12, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x12,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x7c, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x40, 0x0a, 0x21, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x42, 0xff, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42,
	0x1f, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x34, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescData = file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDesc
)

func file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescData)
	})
	return file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDescData
}

var file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_goTypes = []interface{}{
	(*DynamicLocationSetFilter)(nil),                                       // 0: google.ads.googleads.v14.common.DynamicLocationSetFilter
	(*BusinessNameFilter)(nil),                                             // 1: google.ads.googleads.v14.common.BusinessNameFilter
	(*DynamicAffiliateLocationSetFilter)(nil),                              // 2: google.ads.googleads.v14.common.DynamicAffiliateLocationSetFilter
	(enums.FeedItemSetStringFilterTypeEnum_FeedItemSetStringFilterType)(0), // 3: google.ads.googleads.v14.enums.FeedItemSetStringFilterTypeEnum.FeedItemSetStringFilterType
}
var file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v14.common.DynamicLocationSetFilter.business_name_filter:type_name -> google.ads.googleads.v14.common.BusinessNameFilter
	3, // 1: google.ads.googleads.v14.common.BusinessNameFilter.filter_type:type_name -> google.ads.googleads.v14.enums.FeedItemSetStringFilterTypeEnum.FeedItemSetStringFilterType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_init() }
func file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_init() {
	if File_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicLocationSetFilter); i {
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
		file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessNameFilter); i {
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
		file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicAffiliateLocationSetFilter); i {
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
			RawDescriptor: file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto = out.File
	file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_rawDesc = nil
	file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_goTypes = nil
	file_google_ads_googleads_v14_common_feed_item_set_filter_type_infos_proto_depIdxs = nil
}
