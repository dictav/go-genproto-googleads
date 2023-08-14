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
// source: google/ads/googleads/v13/enums/customer_pay_per_conversion_eligibility_failure_reason.proto

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

// Enum describing possible reasons a customer is not eligible to use
// PaymentMode.CONVERSIONS.
type CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason int32

const (
	// Not specified.
	CustomerPayPerConversionEligibilityFailureReasonEnum_UNSPECIFIED CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 0
	// Used for return value only. Represents value unknown in this version.
	CustomerPayPerConversionEligibilityFailureReasonEnum_UNKNOWN CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 1
	// Customer does not have enough conversions.
	CustomerPayPerConversionEligibilityFailureReasonEnum_NOT_ENOUGH_CONVERSIONS CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 2
	// Customer's conversion lag is too high.
	CustomerPayPerConversionEligibilityFailureReasonEnum_CONVERSION_LAG_TOO_HIGH CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 3
	// Customer uses shared budgets.
	CustomerPayPerConversionEligibilityFailureReasonEnum_HAS_CAMPAIGN_WITH_SHARED_BUDGET CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 4
	// Customer has conversions with ConversionActionType.UPLOAD_CLICKS.
	CustomerPayPerConversionEligibilityFailureReasonEnum_HAS_UPLOAD_CLICKS_CONVERSION CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 5
	// Customer's average daily spend is too high.
	CustomerPayPerConversionEligibilityFailureReasonEnum_AVERAGE_DAILY_SPEND_TOO_HIGH CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 6
	// Customer's eligibility has not yet been calculated by the Google Ads
	// backend. Check back soon.
	CustomerPayPerConversionEligibilityFailureReasonEnum_ANALYSIS_NOT_COMPLETE CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 7
	// Customer is not eligible due to other reasons.
	CustomerPayPerConversionEligibilityFailureReasonEnum_OTHER CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason = 8
)

// Enum value maps for CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason.
var (
	CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "NOT_ENOUGH_CONVERSIONS",
		3: "CONVERSION_LAG_TOO_HIGH",
		4: "HAS_CAMPAIGN_WITH_SHARED_BUDGET",
		5: "HAS_UPLOAD_CLICKS_CONVERSION",
		6: "AVERAGE_DAILY_SPEND_TOO_HIGH",
		7: "ANALYSIS_NOT_COMPLETE",
		8: "OTHER",
	}
	CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"UNKNOWN":                         1,
		"NOT_ENOUGH_CONVERSIONS":          2,
		"CONVERSION_LAG_TOO_HIGH":         3,
		"HAS_CAMPAIGN_WITH_SHARED_BUDGET": 4,
		"HAS_UPLOAD_CLICKS_CONVERSION":    5,
		"AVERAGE_DAILY_SPEND_TOO_HIGH":    6,
		"ANALYSIS_NOT_COMPLETE":           7,
		"OTHER":                           8,
	}
)

func (x CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) Enum() *CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason {
	p := new(CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason)
	*p = x
	return p
}

func (x CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_enumTypes[0].Descriptor()
}

func (CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_enumTypes[0]
}

func (x CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason.Descriptor instead.
func (CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing reasons why a customer is not eligible to use
// PaymentMode.CONVERSIONS.
type CustomerPayPerConversionEligibilityFailureReasonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CustomerPayPerConversionEligibilityFailureReasonEnum) Reset() {
	*x = CustomerPayPerConversionEligibilityFailureReasonEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerPayPerConversionEligibilityFailureReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerPayPerConversionEligibilityFailureReasonEnum) ProtoMessage() {}

func (x *CustomerPayPerConversionEligibilityFailureReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerPayPerConversionEligibilityFailureReasonEnum.ProtoReflect.Descriptor instead.
func (*CustomerPayPerConversionEligibilityFailureReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDesc = []byte{
	0x0a, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6c, 0x69,
	0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xd1, 0x02,
	0x0a, 0x34, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x79, 0x50, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x98, 0x02, 0x0a, 0x30, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x50, 0x61, 0x79, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x4f, 0x54,
	0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x53, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x41, 0x47, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x48, 0x49, 0x47, 0x48,
	0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x48, 0x41, 0x53, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x42,
	0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x48, 0x41, 0x53, 0x5f, 0x55,
	0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x53, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x56, 0x45,
	0x52, 0x41, 0x47, 0x45, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x5f, 0x53, 0x50, 0x45, 0x4e, 0x44,
	0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x41,
	0x4e, 0x41, 0x4c, 0x59, 0x53, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10,
	0x08, 0x42, 0x8f, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x35, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x33, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescData = file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDesc
)

func file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDescData
}

var file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_goTypes = []interface{}{
	(CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason)(0), // 0: google.ads.googleads.v13.enums.CustomerPayPerConversionEligibilityFailureReasonEnum.CustomerPayPerConversionEligibilityFailureReason
	(*CustomerPayPerConversionEligibilityFailureReasonEnum)(nil),                                               // 1: google.ads.googleads.v13.enums.CustomerPayPerConversionEligibilityFailureReasonEnum
}
var file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_init()
}
func file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_init() {
	if File_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerPayPerConversionEligibilityFailureReasonEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto = out.File
	file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_rawDesc = nil
	file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_goTypes = nil
	file_google_ads_googleads_v13_enums_customer_pay_per_conversion_eligibility_failure_reason_proto_depIdxs = nil
}