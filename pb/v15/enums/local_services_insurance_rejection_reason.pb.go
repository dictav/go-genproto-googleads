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
// source: google/ads/googleads/v15/enums/local_services_insurance_rejection_reason.proto

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

// Enums describing possible rejection reasons of a local services insurance
// verification artifact.
type LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason int32

const (
	// Not specified.
	LocalServicesInsuranceRejectionReasonEnum_UNSPECIFIED LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 0
	// Used for return value only. Represents value unknown in this version.
	LocalServicesInsuranceRejectionReasonEnum_UNKNOWN LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 1
	// Business name doesn't match business name for the Local Services Ad.
	LocalServicesInsuranceRejectionReasonEnum_BUSINESS_NAME_MISMATCH LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 2
	// Insurance amount doesn't meet requirement listed in the legal summaries
	// documentation for that geographic + category ID combination.
	LocalServicesInsuranceRejectionReasonEnum_INSURANCE_AMOUNT_INSUFFICIENT LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 3
	// Insurance document is expired.
	LocalServicesInsuranceRejectionReasonEnum_EXPIRED LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 4
	// Insurance document is missing a signature.
	LocalServicesInsuranceRejectionReasonEnum_NO_SIGNATURE LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 5
	// Insurance document is missing a policy number.
	LocalServicesInsuranceRejectionReasonEnum_NO_POLICY_NUMBER LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 6
	// Commercial General Liability(CGL) box is not marked in the insurance
	// document.
	LocalServicesInsuranceRejectionReasonEnum_NO_COMMERCIAL_GENERAL_LIABILITY LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 7
	// Insurance document is in an editable format.
	LocalServicesInsuranceRejectionReasonEnum_EDITABLE_FORMAT LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 8
	// Insurance document does not cover insurance for a particular category.
	LocalServicesInsuranceRejectionReasonEnum_CATEGORY_MISMATCH LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 9
	// Insurance document is missing an expiration date.
	LocalServicesInsuranceRejectionReasonEnum_MISSING_EXPIRATION_DATE LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 10
	// Insurance document is poor quality - blurry images, illegible, etc...
	LocalServicesInsuranceRejectionReasonEnum_POOR_QUALITY LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 11
	// Insurance document is suspected of being edited.
	LocalServicesInsuranceRejectionReasonEnum_POTENTIALLY_EDITED LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 12
	// Insurance document not accepted. For example, documents of insurance
	// proposals, but missing coverages are not accepted.
	LocalServicesInsuranceRejectionReasonEnum_WRONG_DOCUMENT_TYPE LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 13
	// Insurance document is not final.
	LocalServicesInsuranceRejectionReasonEnum_NON_FINAL LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 14
	// Insurance document has another flaw not listed explicitly.
	LocalServicesInsuranceRejectionReasonEnum_OTHER LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason = 15
)

// Enum value maps for LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason.
var (
	LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "BUSINESS_NAME_MISMATCH",
		3:  "INSURANCE_AMOUNT_INSUFFICIENT",
		4:  "EXPIRED",
		5:  "NO_SIGNATURE",
		6:  "NO_POLICY_NUMBER",
		7:  "NO_COMMERCIAL_GENERAL_LIABILITY",
		8:  "EDITABLE_FORMAT",
		9:  "CATEGORY_MISMATCH",
		10: "MISSING_EXPIRATION_DATE",
		11: "POOR_QUALITY",
		12: "POTENTIALLY_EDITED",
		13: "WRONG_DOCUMENT_TYPE",
		14: "NON_FINAL",
		15: "OTHER",
	}
	LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"UNKNOWN":                         1,
		"BUSINESS_NAME_MISMATCH":          2,
		"INSURANCE_AMOUNT_INSUFFICIENT":   3,
		"EXPIRED":                         4,
		"NO_SIGNATURE":                    5,
		"NO_POLICY_NUMBER":                6,
		"NO_COMMERCIAL_GENERAL_LIABILITY": 7,
		"EDITABLE_FORMAT":                 8,
		"CATEGORY_MISMATCH":               9,
		"MISSING_EXPIRATION_DATE":         10,
		"POOR_QUALITY":                    11,
		"POTENTIALLY_EDITED":              12,
		"WRONG_DOCUMENT_TYPE":             13,
		"NON_FINAL":                       14,
		"OTHER":                           15,
	}
)

func (x LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason) Enum() *LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason {
	p := new(LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason)
	*p = x
	return p
}

func (x LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_enumTypes[0].Descriptor()
}

func (LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_enumTypes[0]
}

func (x LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason.Descriptor instead.
func (LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the rejection reason of a local services
// insurance verification artifact.
type LocalServicesInsuranceRejectionReasonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalServicesInsuranceRejectionReasonEnum) Reset() {
	*x = LocalServicesInsuranceRejectionReasonEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalServicesInsuranceRejectionReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalServicesInsuranceRejectionReasonEnum) ProtoMessage() {}

func (x *LocalServicesInsuranceRejectionReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalServicesInsuranceRejectionReasonEnum.ProtoReflect.Descriptor instead.
func (*LocalServicesInsuranceRejectionReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f,
	0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x22, 0xb2, 0x03, 0x0a, 0x29, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x84,
	0x03, 0x0a, 0x25, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45,
	0x53, 0x53, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48,
	0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x4f, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43,
	0x59, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x06, 0x12, 0x23, 0x0a, 0x1f, 0x4e, 0x4f,
	0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x52, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x47, 0x45, 0x4e, 0x45,
	0x52, 0x41, 0x4c, 0x5f, 0x4c, 0x49, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x07, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x44, 0x49, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f, 0x4f, 0x52,
	0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4f,
	0x54, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x45, 0x44,
	0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x44, 0x4f, 0x43, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x4e,
	0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x0e, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x0f, 0x42, 0x84, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x2a, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x35, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescData = file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDesc
)

func file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDescData
}

var file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_goTypes = []interface{}{
	(LocalServicesInsuranceRejectionReasonEnum_LocalServicesInsuranceRejectionReason)(0), // 0: google.ads.googleads.v15.enums.LocalServicesInsuranceRejectionReasonEnum.LocalServicesInsuranceRejectionReason
	(*LocalServicesInsuranceRejectionReasonEnum)(nil),                                    // 1: google.ads.googleads.v15.enums.LocalServicesInsuranceRejectionReasonEnum
}
var file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_init()
}
func file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_init() {
	if File_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalServicesInsuranceRejectionReasonEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto = out.File
	file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_rawDesc = nil
	file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_goTypes = nil
	file_google_ads_googleads_v15_enums_local_services_insurance_rejection_reason_proto_depIdxs = nil
}
