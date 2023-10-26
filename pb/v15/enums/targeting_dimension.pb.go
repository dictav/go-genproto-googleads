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
// source: google/ads/googleads/v15/enums/targeting_dimension.proto

package enums

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

// Enum describing possible targeting dimensions.
type TargetingDimensionEnum_TargetingDimension int32

const (
	// Not specified.
	TargetingDimensionEnum_UNSPECIFIED TargetingDimensionEnum_TargetingDimension = 0
	// Used for return value only. Represents value unknown in this version.
	TargetingDimensionEnum_UNKNOWN TargetingDimensionEnum_TargetingDimension = 1
	// Keyword criteria, for example, 'mars cruise'. KEYWORD may be used as a
	// custom bid dimension. Keywords are always a targeting dimension, so may
	// not be set as a target "ALL" dimension with TargetRestriction.
	TargetingDimensionEnum_KEYWORD TargetingDimensionEnum_TargetingDimension = 2
	// Audience criteria, which include user list, user interest, custom
	// affinity,  and custom in market.
	TargetingDimensionEnum_AUDIENCE TargetingDimensionEnum_TargetingDimension = 3
	// Topic criteria for targeting categories of content, for example,
	// 'category::Animals>Pets' Used for Display and Video targeting.
	TargetingDimensionEnum_TOPIC TargetingDimensionEnum_TargetingDimension = 4
	// Criteria for targeting gender.
	TargetingDimensionEnum_GENDER TargetingDimensionEnum_TargetingDimension = 5
	// Criteria for targeting age ranges.
	TargetingDimensionEnum_AGE_RANGE TargetingDimensionEnum_TargetingDimension = 6
	// Placement criteria, which include websites like 'www.flowers4sale.com',
	// as well as mobile applications, mobile app categories, YouTube videos,
	// and YouTube channels.
	TargetingDimensionEnum_PLACEMENT TargetingDimensionEnum_TargetingDimension = 7
	// Criteria for parental status targeting.
	TargetingDimensionEnum_PARENTAL_STATUS TargetingDimensionEnum_TargetingDimension = 8
	// Criteria for income range targeting.
	TargetingDimensionEnum_INCOME_RANGE TargetingDimensionEnum_TargetingDimension = 9
)

// Enum value maps for TargetingDimensionEnum_TargetingDimension.
var (
	TargetingDimensionEnum_TargetingDimension_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "KEYWORD",
		3: "AUDIENCE",
		4: "TOPIC",
		5: "GENDER",
		6: "AGE_RANGE",
		7: "PLACEMENT",
		8: "PARENTAL_STATUS",
		9: "INCOME_RANGE",
	}
	TargetingDimensionEnum_TargetingDimension_value = map[string]int32{
		"UNSPECIFIED":     0,
		"UNKNOWN":         1,
		"KEYWORD":         2,
		"AUDIENCE":        3,
		"TOPIC":           4,
		"GENDER":          5,
		"AGE_RANGE":       6,
		"PLACEMENT":       7,
		"PARENTAL_STATUS": 8,
		"INCOME_RANGE":    9,
	}
)

func (x TargetingDimensionEnum_TargetingDimension) Enum() *TargetingDimensionEnum_TargetingDimension {
	p := new(TargetingDimensionEnum_TargetingDimension)
	*p = x
	return p
}

func (x TargetingDimensionEnum_TargetingDimension) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetingDimensionEnum_TargetingDimension) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v15_enums_targeting_dimension_proto_enumTypes[0].Descriptor()
}

func (TargetingDimensionEnum_TargetingDimension) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v15_enums_targeting_dimension_proto_enumTypes[0]
}

func (x TargetingDimensionEnum_TargetingDimension) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetingDimensionEnum_TargetingDimension.Descriptor instead.
func (TargetingDimensionEnum_TargetingDimension) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescGZIP(), []int{0, 0}
}

// The dimensions that can be targeted.
type TargetingDimensionEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TargetingDimensionEnum) Reset() {
	*x = TargetingDimensionEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_enums_targeting_dimension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetingDimensionEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetingDimensionEnum) ProtoMessage() {}

func (x *TargetingDimensionEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_enums_targeting_dimension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetingDimensionEnum.ProtoReflect.Descriptor instead.
func (*TargetingDimensionEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v15_enums_targeting_dimension_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x16, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x45,
	0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x55, 0x44, 0x49, 0x45,
	0x4e, 0x43, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x10, 0x04,
	0x12, 0x0a, 0x0a, 0x06, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x47, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41,
	0x52, 0x45, 0x4e, 0x54, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x08, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x09, 0x42, 0xf1, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x17, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescData = file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDesc
)

func file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDescData
}

var file_google_ads_googleads_v15_enums_targeting_dimension_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v15_enums_targeting_dimension_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v15_enums_targeting_dimension_proto_goTypes = []interface{}{
	(TargetingDimensionEnum_TargetingDimension)(0), // 0: google.ads.googleads.v15.enums.TargetingDimensionEnum.TargetingDimension
	(*TargetingDimensionEnum)(nil),                 // 1: google.ads.googleads.v15.enums.TargetingDimensionEnum
}
var file_google_ads_googleads_v15_enums_targeting_dimension_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_enums_targeting_dimension_proto_init() }
func file_google_ads_googleads_v15_enums_targeting_dimension_proto_init() {
	if File_google_ads_googleads_v15_enums_targeting_dimension_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_enums_targeting_dimension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetingDimensionEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_enums_targeting_dimension_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_enums_targeting_dimension_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v15_enums_targeting_dimension_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v15_enums_targeting_dimension_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_enums_targeting_dimension_proto = out.File
	file_google_ads_googleads_v15_enums_targeting_dimension_proto_rawDesc = nil
	file_google_ads_googleads_v15_enums_targeting_dimension_proto_goTypes = nil
	file_google_ads_googleads_v15_enums_targeting_dimension_proto_depIdxs = nil
}