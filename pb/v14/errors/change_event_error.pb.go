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
// source: google/ads/googleads/v14/errors/change_event_error.proto

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

// Enum describing possible change event errors.
type ChangeEventErrorEnum_ChangeEventError int32

const (
	// Enum unspecified.
	ChangeEventErrorEnum_UNSPECIFIED ChangeEventErrorEnum_ChangeEventError = 0
	// The received error code is not known in this version.
	ChangeEventErrorEnum_UNKNOWN ChangeEventErrorEnum_ChangeEventError = 1
	// The requested start date is too old. It cannot be older than 30 days.
	ChangeEventErrorEnum_START_DATE_TOO_OLD ChangeEventErrorEnum_ChangeEventError = 2
	// The change_event search request must specify a finite range filter
	// on change_date_time.
	ChangeEventErrorEnum_CHANGE_DATE_RANGE_INFINITE ChangeEventErrorEnum_ChangeEventError = 3
	// The change event search request has specified invalid date time filters
	// that can never logically produce any valid results (for example, start
	// time after end time).
	ChangeEventErrorEnum_CHANGE_DATE_RANGE_NEGATIVE ChangeEventErrorEnum_ChangeEventError = 4
	// The change_event search request must specify a LIMIT.
	ChangeEventErrorEnum_LIMIT_NOT_SPECIFIED ChangeEventErrorEnum_ChangeEventError = 5
	// The LIMIT specified by change_event request should be less than or equal
	// to 10K.
	ChangeEventErrorEnum_INVALID_LIMIT_CLAUSE ChangeEventErrorEnum_ChangeEventError = 6
)

// Enum value maps for ChangeEventErrorEnum_ChangeEventError.
var (
	ChangeEventErrorEnum_ChangeEventError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "START_DATE_TOO_OLD",
		3: "CHANGE_DATE_RANGE_INFINITE",
		4: "CHANGE_DATE_RANGE_NEGATIVE",
		5: "LIMIT_NOT_SPECIFIED",
		6: "INVALID_LIMIT_CLAUSE",
	}
	ChangeEventErrorEnum_ChangeEventError_value = map[string]int32{
		"UNSPECIFIED":                0,
		"UNKNOWN":                    1,
		"START_DATE_TOO_OLD":         2,
		"CHANGE_DATE_RANGE_INFINITE": 3,
		"CHANGE_DATE_RANGE_NEGATIVE": 4,
		"LIMIT_NOT_SPECIFIED":        5,
		"INVALID_LIMIT_CLAUSE":       6,
	}
)

func (x ChangeEventErrorEnum_ChangeEventError) Enum() *ChangeEventErrorEnum_ChangeEventError {
	p := new(ChangeEventErrorEnum_ChangeEventError)
	*p = x
	return p
}

func (x ChangeEventErrorEnum_ChangeEventError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeEventErrorEnum_ChangeEventError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v14_errors_change_event_error_proto_enumTypes[0].Descriptor()
}

func (ChangeEventErrorEnum_ChangeEventError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v14_errors_change_event_error_proto_enumTypes[0]
}

func (x ChangeEventErrorEnum_ChangeEventError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeEventErrorEnum_ChangeEventError.Descriptor instead.
func (ChangeEventErrorEnum_ChangeEventError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible change event errors.
type ChangeEventErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeEventErrorEnum) Reset() {
	*x = ChangeEventErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_errors_change_event_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeEventErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeEventErrorEnum) ProtoMessage() {}

func (x *ChangeEventErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_errors_change_event_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeEventErrorEnum.ProtoReflect.Descriptor instead.
func (*ChangeEventErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v14_errors_change_event_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v14_errors_change_event_error_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x14,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0xbb, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4f, 0x4c, 0x44, 0x10, 0x02, 0x12,
	0x1e, 0x0a, 0x1a, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x45, 0x10, 0x03, 0x12,
	0x1e, 0x0a, 0x1a, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12,
	0x17, 0x0a, 0x13, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x43, 0x4c, 0x41, 0x55, 0x53, 0x45,
	0x10, 0x06, 0x42, 0xf5, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x34, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x15, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x34, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescData = file_google_ads_googleads_v14_errors_change_event_error_proto_rawDesc
)

func file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v14_errors_change_event_error_proto_rawDescData
}

var file_google_ads_googleads_v14_errors_change_event_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v14_errors_change_event_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v14_errors_change_event_error_proto_goTypes = []interface{}{
	(ChangeEventErrorEnum_ChangeEventError)(0), // 0: google.ads.googleads.v14.errors.ChangeEventErrorEnum.ChangeEventError
	(*ChangeEventErrorEnum)(nil),               // 1: google.ads.googleads.v14.errors.ChangeEventErrorEnum
}
var file_google_ads_googleads_v14_errors_change_event_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v14_errors_change_event_error_proto_init() }
func file_google_ads_googleads_v14_errors_change_event_error_proto_init() {
	if File_google_ads_googleads_v14_errors_change_event_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v14_errors_change_event_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeEventErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v14_errors_change_event_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v14_errors_change_event_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v14_errors_change_event_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v14_errors_change_event_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v14_errors_change_event_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v14_errors_change_event_error_proto = out.File
	file_google_ads_googleads_v14_errors_change_event_error_proto_rawDesc = nil
	file_google_ads_googleads_v14_errors_change_event_error_proto_goTypes = nil
	file_google_ads_googleads_v14_errors_change_event_error_proto_depIdxs = nil
}