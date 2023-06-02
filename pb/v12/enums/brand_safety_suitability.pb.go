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
// source: google/ads/googleads/v12/enums/brand_safety_suitability.proto

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

// 3-Tier brand safety suitability control.
type BrandSafetySuitabilityEnum_BrandSafetySuitability int32

const (
	// Not specified.
	BrandSafetySuitabilityEnum_UNSPECIFIED BrandSafetySuitabilityEnum_BrandSafetySuitability = 0
	// Used for return value only. Represents value unknown in this version.
	BrandSafetySuitabilityEnum_UNKNOWN BrandSafetySuitabilityEnum_BrandSafetySuitability = 1
	// This option lets you show ads across all inventory on YouTube and video
	// partners that meet our standards for monetization. This option may be an
	// appropriate choice for brands that want maximum access to the full
	// breadth of videos eligible for ads, including, for example, videos that
	// have strong profanity in the context of comedy or a documentary, or
	// excessive violence as featured in video games.
	BrandSafetySuitabilityEnum_EXPANDED_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 2
	// This option lets you show ads across a wide range of content that's
	// appropriate for most brands, such as popular music videos, documentaries,
	// and movie trailers. The content you can show ads on is based on YouTube's
	// advertiser-friendly content guidelines that take into account, for
	// example, the strength or frequency of profanity, or the appropriateness
	// of subject matter like sensitive events. Ads won't show, for example, on
	// content with repeated strong profanity, strong sexual content, or graphic
	// violence.
	BrandSafetySuitabilityEnum_STANDARD_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 3
	// This option lets you show ads on a reduced range of content that's
	// appropriate for brands with particularly strict guidelines around
	// inappropriate language and sexual suggestiveness; above and beyond what
	// YouTube's advertiser-friendly content guidelines address. The videos
	// accessible in this sensitive category meet heightened requirements,
	// especially for inappropriate language and sexual suggestiveness. For
	// example, your ads will be excluded from showing on some of YouTube's most
	// popular music videos and other pop culture content across YouTube and
	// Google video partners.
	BrandSafetySuitabilityEnum_LIMITED_INVENTORY BrandSafetySuitabilityEnum_BrandSafetySuitability = 4
)

// Enum value maps for BrandSafetySuitabilityEnum_BrandSafetySuitability.
var (
	BrandSafetySuitabilityEnum_BrandSafetySuitability_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "EXPANDED_INVENTORY",
		3: "STANDARD_INVENTORY",
		4: "LIMITED_INVENTORY",
	}
	BrandSafetySuitabilityEnum_BrandSafetySuitability_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"EXPANDED_INVENTORY": 2,
		"STANDARD_INVENTORY": 3,
		"LIMITED_INVENTORY":  4,
	}
)

func (x BrandSafetySuitabilityEnum_BrandSafetySuitability) Enum() *BrandSafetySuitabilityEnum_BrandSafetySuitability {
	p := new(BrandSafetySuitabilityEnum_BrandSafetySuitability)
	*p = x
	return p
}

func (x BrandSafetySuitabilityEnum_BrandSafetySuitability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrandSafetySuitabilityEnum_BrandSafetySuitability) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_enumTypes[0].Descriptor()
}

func (BrandSafetySuitabilityEnum_BrandSafetySuitability) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_enumTypes[0]
}

func (x BrandSafetySuitabilityEnum_BrandSafetySuitability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrandSafetySuitabilityEnum_BrandSafetySuitability.Descriptor instead.
func (BrandSafetySuitabilityEnum_BrandSafetySuitability) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum with 3-Tier brand safety suitability control.
type BrandSafetySuitabilityEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BrandSafetySuitabilityEnum) Reset() {
	*x = BrandSafetySuitabilityEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandSafetySuitabilityEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandSafetySuitabilityEnum) ProtoMessage() {}

func (x *BrandSafetySuitabilityEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandSafetySuitabilityEnum.ProtoReflect.Descriptor instead.
func (*BrandSafetySuitabilityEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v12_enums_brand_safety_suitability_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x5f, 0x73, 0x75,
	0x69, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22,
	0x9b, 0x01, 0x0a, 0x1a, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x53,
	0x75, 0x69, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x7d,
	0x0a, 0x16, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x53, 0x75, 0x69,
	0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x58, 0x50, 0x41, 0x4e, 0x44,
	0x45, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x45, 0x4e, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x45, 0x4e,
	0x54, 0x4f, 0x52, 0x59, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45,
	0x44, 0x5f, 0x49, 0x4e, 0x56, 0x45, 0x4e, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x04, 0x42, 0xf5, 0x01,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1b, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x66, 0x65, 0x74,
	0x79, 0x53, 0x75, 0x69, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescData = file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDesc
)

func file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDescData
}

var file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_goTypes = []interface{}{
	(BrandSafetySuitabilityEnum_BrandSafetySuitability)(0), // 0: google.ads.googleads.v12.enums.BrandSafetySuitabilityEnum.BrandSafetySuitability
	(*BrandSafetySuitabilityEnum)(nil),                     // 1: google.ads.googleads.v12.enums.BrandSafetySuitabilityEnum
}
var file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_init() }
func file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_init() {
	if File_google_ads_googleads_v12_enums_brand_safety_suitability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandSafetySuitabilityEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_enums_brand_safety_suitability_proto = out.File
	file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_rawDesc = nil
	file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_goTypes = nil
	file_google_ads_googleads_v12_enums_brand_safety_suitability_proto_depIdxs = nil
}
