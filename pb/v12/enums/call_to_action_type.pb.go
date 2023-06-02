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
// source: google/ads/googleads/v12/enums/call_to_action_type.proto

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

// Enum describing possible types of call to action.
type CallToActionTypeEnum_CallToActionType int32

const (
	// Not specified.
	CallToActionTypeEnum_UNSPECIFIED CallToActionTypeEnum_CallToActionType = 0
	// Used for return value only. Represents value unknown in this version.
	CallToActionTypeEnum_UNKNOWN CallToActionTypeEnum_CallToActionType = 1
	// The call to action type is learn more.
	CallToActionTypeEnum_LEARN_MORE CallToActionTypeEnum_CallToActionType = 2
	// The call to action type is get quote.
	CallToActionTypeEnum_GET_QUOTE CallToActionTypeEnum_CallToActionType = 3
	// The call to action type is apply now.
	CallToActionTypeEnum_APPLY_NOW CallToActionTypeEnum_CallToActionType = 4
	// The call to action type is sign up.
	CallToActionTypeEnum_SIGN_UP CallToActionTypeEnum_CallToActionType = 5
	// The call to action type is contact us.
	CallToActionTypeEnum_CONTACT_US CallToActionTypeEnum_CallToActionType = 6
	// The call to action type is subscribe.
	CallToActionTypeEnum_SUBSCRIBE CallToActionTypeEnum_CallToActionType = 7
	// The call to action type is download.
	CallToActionTypeEnum_DOWNLOAD CallToActionTypeEnum_CallToActionType = 8
	// The call to action type is book now.
	CallToActionTypeEnum_BOOK_NOW CallToActionTypeEnum_CallToActionType = 9
	// The call to action type is shop now.
	CallToActionTypeEnum_SHOP_NOW CallToActionTypeEnum_CallToActionType = 10
)

// Enum value maps for CallToActionTypeEnum_CallToActionType.
var (
	CallToActionTypeEnum_CallToActionType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "LEARN_MORE",
		3:  "GET_QUOTE",
		4:  "APPLY_NOW",
		5:  "SIGN_UP",
		6:  "CONTACT_US",
		7:  "SUBSCRIBE",
		8:  "DOWNLOAD",
		9:  "BOOK_NOW",
		10: "SHOP_NOW",
	}
	CallToActionTypeEnum_CallToActionType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"LEARN_MORE":  2,
		"GET_QUOTE":   3,
		"APPLY_NOW":   4,
		"SIGN_UP":     5,
		"CONTACT_US":  6,
		"SUBSCRIBE":   7,
		"DOWNLOAD":    8,
		"BOOK_NOW":    9,
		"SHOP_NOW":    10,
	}
)

func (x CallToActionTypeEnum_CallToActionType) Enum() *CallToActionTypeEnum_CallToActionType {
	p := new(CallToActionTypeEnum_CallToActionType)
	*p = x
	return p
}

func (x CallToActionTypeEnum_CallToActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallToActionTypeEnum_CallToActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v12_enums_call_to_action_type_proto_enumTypes[0].Descriptor()
}

func (CallToActionTypeEnum_CallToActionType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v12_enums_call_to_action_type_proto_enumTypes[0]
}

func (x CallToActionTypeEnum_CallToActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallToActionTypeEnum_CallToActionType.Descriptor instead.
func (CallToActionTypeEnum_CallToActionType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the call to action types.
type CallToActionTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallToActionTypeEnum) Reset() {
	*x = CallToActionTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_enums_call_to_action_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallToActionTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallToActionTypeEnum) ProtoMessage() {}

func (x *CallToActionTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_enums_call_to_action_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallToActionTypeEnum.ProtoReflect.Descriptor instead.
func (*CallToActionTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v12_enums_call_to_action_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x14, 0x43,
	0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0xb4, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x5f,
	0x4d, 0x4f, 0x52, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x54, 0x5f, 0x51, 0x55,
	0x4f, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x4e,
	0x4f, 0x57, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x55, 0x50, 0x10,
	0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x55, 0x53, 0x10,
	0x06, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x07,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x08, 0x12, 0x0c,
	0x0a, 0x08, 0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x4e, 0x4f, 0x57, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x48, 0x4f, 0x50, 0x5f, 0x4e, 0x4f, 0x57, 0x10, 0x0a, 0x42, 0xef, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x15, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescData = file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDesc
)

func file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDescData
}

var file_google_ads_googleads_v12_enums_call_to_action_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v12_enums_call_to_action_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_enums_call_to_action_type_proto_goTypes = []interface{}{
	(CallToActionTypeEnum_CallToActionType)(0), // 0: google.ads.googleads.v12.enums.CallToActionTypeEnum.CallToActionType
	(*CallToActionTypeEnum)(nil),               // 1: google.ads.googleads.v12.enums.CallToActionTypeEnum
}
var file_google_ads_googleads_v12_enums_call_to_action_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_enums_call_to_action_type_proto_init() }
func file_google_ads_googleads_v12_enums_call_to_action_type_proto_init() {
	if File_google_ads_googleads_v12_enums_call_to_action_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_enums_call_to_action_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallToActionTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_enums_call_to_action_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_enums_call_to_action_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v12_enums_call_to_action_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v12_enums_call_to_action_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_enums_call_to_action_type_proto = out.File
	file_google_ads_googleads_v12_enums_call_to_action_type_proto_rawDesc = nil
	file_google_ads_googleads_v12_enums_call_to_action_type_proto_goTypes = nil
	file_google_ads_googleads_v12_enums_call_to_action_type_proto_depIdxs = nil
}
