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
// source: google/ads/googleads/v13/errors/feed_item_target_error.proto

package errors

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

// Enum describing possible feed item target errors.
type FeedItemTargetErrorEnum_FeedItemTargetError int32

const (
	// Enum unspecified.
	FeedItemTargetErrorEnum_UNSPECIFIED FeedItemTargetErrorEnum_FeedItemTargetError = 0
	// The received error code is not known in this version.
	FeedItemTargetErrorEnum_UNKNOWN FeedItemTargetErrorEnum_FeedItemTargetError = 1
	// On CREATE, the FeedItemTarget must have a populated field in the oneof
	// target.
	FeedItemTargetErrorEnum_MUST_SET_TARGET_ONEOF_ON_CREATE FeedItemTargetErrorEnum_FeedItemTargetError = 2
	// The specified feed item target already exists, so it cannot be added.
	FeedItemTargetErrorEnum_FEED_ITEM_TARGET_ALREADY_EXISTS FeedItemTargetErrorEnum_FeedItemTargetError = 3
	// The schedules for a given feed item cannot overlap.
	FeedItemTargetErrorEnum_FEED_ITEM_SCHEDULES_CANNOT_OVERLAP FeedItemTargetErrorEnum_FeedItemTargetError = 4
	// Too many targets of a given type were added for a single feed item.
	FeedItemTargetErrorEnum_TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE FeedItemTargetErrorEnum_FeedItemTargetError = 5
	// Too many AdSchedules are enabled for the feed item for the given day.
	FeedItemTargetErrorEnum_TOO_MANY_SCHEDULES_PER_DAY FeedItemTargetErrorEnum_FeedItemTargetError = 6
	// A feed item may either have an enabled campaign target or an enabled ad
	// group target.
	FeedItemTargetErrorEnum_CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS FeedItemTargetErrorEnum_FeedItemTargetError = 7
	// Duplicate ad schedules aren't allowed.
	FeedItemTargetErrorEnum_DUPLICATE_AD_SCHEDULE FeedItemTargetErrorEnum_FeedItemTargetError = 8
	// Duplicate keywords aren't allowed.
	FeedItemTargetErrorEnum_DUPLICATE_KEYWORD FeedItemTargetErrorEnum_FeedItemTargetError = 9
)

// Enum value maps for FeedItemTargetErrorEnum_FeedItemTargetError.
var (
	FeedItemTargetErrorEnum_FeedItemTargetError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "MUST_SET_TARGET_ONEOF_ON_CREATE",
		3: "FEED_ITEM_TARGET_ALREADY_EXISTS",
		4: "FEED_ITEM_SCHEDULES_CANNOT_OVERLAP",
		5: "TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE",
		6: "TOO_MANY_SCHEDULES_PER_DAY",
		7: "CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS",
		8: "DUPLICATE_AD_SCHEDULE",
		9: "DUPLICATE_KEYWORD",
	}
	FeedItemTargetErrorEnum_FeedItemTargetError_value = map[string]int32{
		"UNSPECIFIED":                                               0,
		"UNKNOWN":                                                   1,
		"MUST_SET_TARGET_ONEOF_ON_CREATE":                           2,
		"FEED_ITEM_TARGET_ALREADY_EXISTS":                           3,
		"FEED_ITEM_SCHEDULES_CANNOT_OVERLAP":                        4,
		"TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE":                      5,
		"TOO_MANY_SCHEDULES_PER_DAY":                                6,
		"CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS": 7,
		"DUPLICATE_AD_SCHEDULE":                                     8,
		"DUPLICATE_KEYWORD":                                         9,
	}
)

func (x FeedItemTargetErrorEnum_FeedItemTargetError) Enum() *FeedItemTargetErrorEnum_FeedItemTargetError {
	p := new(FeedItemTargetErrorEnum_FeedItemTargetError)
	*p = x
	return p
}

func (x FeedItemTargetErrorEnum_FeedItemTargetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedItemTargetErrorEnum_FeedItemTargetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v13_errors_feed_item_target_error_proto_enumTypes[0].Descriptor()
}

func (FeedItemTargetErrorEnum_FeedItemTargetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v13_errors_feed_item_target_error_proto_enumTypes[0]
}

func (x FeedItemTargetErrorEnum_FeedItemTargetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedItemTargetErrorEnum_FeedItemTargetError.Descriptor instead.
func (FeedItemTargetErrorEnum_FeedItemTargetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible feed item target errors.
type FeedItemTargetErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedItemTargetErrorEnum) Reset() {
	*x = FeedItemTargetErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_errors_feed_item_target_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemTargetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemTargetErrorEnum) ProtoMessage() {}

func (x *FeedItemTargetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_errors_feed_item_target_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemTargetErrorEnum.ProtoReflect.Descriptor instead.
func (*FeedItemTargetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v13_errors_feed_item_target_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22,
	0xfc, 0x02, 0x0a, 0x17, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xe0, 0x02, 0x0a, 0x13,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x41,
	0x52, 0x47, 0x45, 0x54, 0x5f, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x49,
	0x54, 0x45, 0x4d, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x46,
	0x45, 0x45, 0x44, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c,
	0x45, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41,
	0x50, 0x10, 0x04, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x47, 0x49, 0x56, 0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x05, 0x12, 0x1e, 0x0a,
	0x1a, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55,
	0x4c, 0x45, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x59, 0x10, 0x06, 0x12, 0x3d, 0x0a,
	0x39, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x48, 0x41, 0x56, 0x45, 0x5f, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x53, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15,
	0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x44, 0x5f, 0x53, 0x43, 0x48,
	0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x55, 0x50, 0x4c, 0x49,
	0x43, 0x41, 0x54, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x09, 0x42, 0xf8,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x18, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x33, 0x5c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x33, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescData = file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDesc
)

func file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDescData
}

var file_google_ads_googleads_v13_errors_feed_item_target_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v13_errors_feed_item_target_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_errors_feed_item_target_error_proto_goTypes = []interface{}{
	(FeedItemTargetErrorEnum_FeedItemTargetError)(0), // 0: google.ads.googleads.v13.errors.FeedItemTargetErrorEnum.FeedItemTargetError
	(*FeedItemTargetErrorEnum)(nil),                  // 1: google.ads.googleads.v13.errors.FeedItemTargetErrorEnum
}
var file_google_ads_googleads_v13_errors_feed_item_target_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_errors_feed_item_target_error_proto_init() }
func file_google_ads_googleads_v13_errors_feed_item_target_error_proto_init() {
	if File_google_ads_googleads_v13_errors_feed_item_target_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_errors_feed_item_target_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemTargetErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_errors_feed_item_target_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_errors_feed_item_target_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v13_errors_feed_item_target_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v13_errors_feed_item_target_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_errors_feed_item_target_error_proto = out.File
	file_google_ads_googleads_v13_errors_feed_item_target_error_proto_rawDesc = nil
	file_google_ads_googleads_v13_errors_feed_item_target_error_proto_goTypes = nil
	file_google_ads_googleads_v13_errors_feed_item_target_error_proto_depIdxs = nil
}
