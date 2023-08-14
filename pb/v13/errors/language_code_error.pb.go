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
// source: google/ads/googleads/v13/errors/language_code_error.proto

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

// Enum describing language code errors.
type LanguageCodeErrorEnum_LanguageCodeError int32

const (
	// Enum unspecified.
	LanguageCodeErrorEnum_UNSPECIFIED LanguageCodeErrorEnum_LanguageCodeError = 0
	// The received error code is not known in this version.
	LanguageCodeErrorEnum_UNKNOWN LanguageCodeErrorEnum_LanguageCodeError = 1
	// The input language code is not recognized.
	LanguageCodeErrorEnum_LANGUAGE_CODE_NOT_FOUND LanguageCodeErrorEnum_LanguageCodeError = 2
	// The language code is not supported.
	LanguageCodeErrorEnum_INVALID_LANGUAGE_CODE LanguageCodeErrorEnum_LanguageCodeError = 3
)

// Enum value maps for LanguageCodeErrorEnum_LanguageCodeError.
var (
	LanguageCodeErrorEnum_LanguageCodeError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "LANGUAGE_CODE_NOT_FOUND",
		3: "INVALID_LANGUAGE_CODE",
	}
	LanguageCodeErrorEnum_LanguageCodeError_value = map[string]int32{
		"UNSPECIFIED":             0,
		"UNKNOWN":                 1,
		"LANGUAGE_CODE_NOT_FOUND": 2,
		"INVALID_LANGUAGE_CODE":   3,
	}
)

func (x LanguageCodeErrorEnum_LanguageCodeError) Enum() *LanguageCodeErrorEnum_LanguageCodeError {
	p := new(LanguageCodeErrorEnum_LanguageCodeError)
	*p = x
	return p
}

func (x LanguageCodeErrorEnum_LanguageCodeError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LanguageCodeErrorEnum_LanguageCodeError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v13_errors_language_code_error_proto_enumTypes[0].Descriptor()
}

func (LanguageCodeErrorEnum_LanguageCodeError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v13_errors_language_code_error_proto_enumTypes[0]
}

func (x LanguageCodeErrorEnum_LanguageCodeError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LanguageCodeErrorEnum_LanguageCodeError.Descriptor instead.
func (LanguageCodeErrorEnum_LanguageCodeError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing language code errors.
type LanguageCodeErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LanguageCodeErrorEnum) Reset() {
	*x = LanguageCodeErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_errors_language_code_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageCodeErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageCodeErrorEnum) ProtoMessage() {}

func (x *LanguageCodeErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_errors_language_code_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageCodeErrorEnum.ProtoReflect.Descriptor instead.
func (*LanguageCodeErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v13_errors_language_code_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_errors_language_code_error_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x82, 0x01, 0x0a,
	0x15, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x69, 0x0a, 0x11, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x41, 0x4e,
	0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10,
	0x03, 0x42, 0xf6, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x16, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x33, 0x5c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x33, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescData = file_google_ads_googleads_v13_errors_language_code_error_proto_rawDesc
)

func file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_errors_language_code_error_proto_rawDescData
}

var file_google_ads_googleads_v13_errors_language_code_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v13_errors_language_code_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_errors_language_code_error_proto_goTypes = []interface{}{
	(LanguageCodeErrorEnum_LanguageCodeError)(0), // 0: google.ads.googleads.v13.errors.LanguageCodeErrorEnum.LanguageCodeError
	(*LanguageCodeErrorEnum)(nil),                // 1: google.ads.googleads.v13.errors.LanguageCodeErrorEnum
}
var file_google_ads_googleads_v13_errors_language_code_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_errors_language_code_error_proto_init() }
func file_google_ads_googleads_v13_errors_language_code_error_proto_init() {
	if File_google_ads_googleads_v13_errors_language_code_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_errors_language_code_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageCodeErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v13_errors_language_code_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_errors_language_code_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_errors_language_code_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v13_errors_language_code_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v13_errors_language_code_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_errors_language_code_error_proto = out.File
	file_google_ads_googleads_v13_errors_language_code_error_proto_rawDesc = nil
	file_google_ads_googleads_v13_errors_language_code_error_proto_goTypes = nil
	file_google_ads_googleads_v13_errors_language_code_error_proto_depIdxs = nil
}