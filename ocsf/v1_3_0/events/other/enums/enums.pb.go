// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: ocsf/v1_3_0/events/other/enums/enums.proto

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

type BASE_EVENT_ACTIVITY_ID int32

const (
	BASE_EVENT_ACTIVITY_ID_BASE_EVENT_ACTIVITY_ID_UNKNOWN BASE_EVENT_ACTIVITY_ID = 0  // Type: OCSF_VALUE; EnumValue: 0;
	BASE_EVENT_ACTIVITY_ID_BASE_EVENT_ACTIVITY_ID_OTHER   BASE_EVENT_ACTIVITY_ID = 99 // EnumValue: 99; Type: OCSF_VALUE;
)

// Enum value maps for BASE_EVENT_ACTIVITY_ID.
var (
	BASE_EVENT_ACTIVITY_ID_name = map[int32]string{
		0:  "BASE_EVENT_ACTIVITY_ID_UNKNOWN",
		99: "BASE_EVENT_ACTIVITY_ID_OTHER",
	}
	BASE_EVENT_ACTIVITY_ID_value = map[string]int32{
		"BASE_EVENT_ACTIVITY_ID_UNKNOWN": 0,
		"BASE_EVENT_ACTIVITY_ID_OTHER":   99,
	}
)

func (x BASE_EVENT_ACTIVITY_ID) Enum() *BASE_EVENT_ACTIVITY_ID {
	p := new(BASE_EVENT_ACTIVITY_ID)
	*p = x
	return p
}

func (x BASE_EVENT_ACTIVITY_ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BASE_EVENT_ACTIVITY_ID) Descriptor() protoreflect.EnumDescriptor {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[0].Descriptor()
}

func (BASE_EVENT_ACTIVITY_ID) Type() protoreflect.EnumType {
	return &file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[0]
}

func (x BASE_EVENT_ACTIVITY_ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BASE_EVENT_ACTIVITY_ID.Descriptor instead.
func (BASE_EVENT_ACTIVITY_ID) EnumDescriptor() ([]byte, []int) {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescGZIP(), []int{0}
}

type BASE_EVENT_CATEGORY_UID int32

const (
	BASE_EVENT_CATEGORY_UID_BASE_EVENT_CATEGORY_UID_UNCATEGORIZED BASE_EVENT_CATEGORY_UID = 0 // Type: OCSF_VALUE; EnumValue: 0;
)

// Enum value maps for BASE_EVENT_CATEGORY_UID.
var (
	BASE_EVENT_CATEGORY_UID_name = map[int32]string{
		0: "BASE_EVENT_CATEGORY_UID_UNCATEGORIZED",
	}
	BASE_EVENT_CATEGORY_UID_value = map[string]int32{
		"BASE_EVENT_CATEGORY_UID_UNCATEGORIZED": 0,
	}
)

func (x BASE_EVENT_CATEGORY_UID) Enum() *BASE_EVENT_CATEGORY_UID {
	p := new(BASE_EVENT_CATEGORY_UID)
	*p = x
	return p
}

func (x BASE_EVENT_CATEGORY_UID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BASE_EVENT_CATEGORY_UID) Descriptor() protoreflect.EnumDescriptor {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[1].Descriptor()
}

func (BASE_EVENT_CATEGORY_UID) Type() protoreflect.EnumType {
	return &file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[1]
}

func (x BASE_EVENT_CATEGORY_UID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BASE_EVENT_CATEGORY_UID.Descriptor instead.
func (BASE_EVENT_CATEGORY_UID) EnumDescriptor() ([]byte, []int) {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescGZIP(), []int{1}
}

type BASE_EVENT_CLASS_UID int32

const (
	BASE_EVENT_CLASS_UID_BASE_EVENT_CLASS_UID_BASE_EVENT BASE_EVENT_CLASS_UID = 0 // Type: OCSF_VALUE; EnumValue: 0;
)

// Enum value maps for BASE_EVENT_CLASS_UID.
var (
	BASE_EVENT_CLASS_UID_name = map[int32]string{
		0: "BASE_EVENT_CLASS_UID_BASE_EVENT",
	}
	BASE_EVENT_CLASS_UID_value = map[string]int32{
		"BASE_EVENT_CLASS_UID_BASE_EVENT": 0,
	}
)

func (x BASE_EVENT_CLASS_UID) Enum() *BASE_EVENT_CLASS_UID {
	p := new(BASE_EVENT_CLASS_UID)
	*p = x
	return p
}

func (x BASE_EVENT_CLASS_UID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BASE_EVENT_CLASS_UID) Descriptor() protoreflect.EnumDescriptor {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[2].Descriptor()
}

func (BASE_EVENT_CLASS_UID) Type() protoreflect.EnumType {
	return &file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[2]
}

func (x BASE_EVENT_CLASS_UID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BASE_EVENT_CLASS_UID.Descriptor instead.
func (BASE_EVENT_CLASS_UID) EnumDescriptor() ([]byte, []int) {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescGZIP(), []int{2}
}

type BASE_EVENT_SEVERITY_ID int32

const (
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_UNKNOWN       BASE_EVENT_SEVERITY_ID = 0  // EnumValue: 0; Type: OCSF_VALUE;
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_INFORMATIONAL BASE_EVENT_SEVERITY_ID = 1  // Type: OCSF_VALUE; EnumValue: 1;
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_LOW           BASE_EVENT_SEVERITY_ID = 2  // Type: OCSF_VALUE; EnumValue: 2;
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_MEDIUM        BASE_EVENT_SEVERITY_ID = 3  // Type: OCSF_VALUE; EnumValue: 3;
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_HIGH          BASE_EVENT_SEVERITY_ID = 4  // Type: OCSF_VALUE; EnumValue: 4;
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_CRITICAL      BASE_EVENT_SEVERITY_ID = 5  // Type: OCSF_VALUE; EnumValue: 5;
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_FATAL         BASE_EVENT_SEVERITY_ID = 6  // EnumValue: 6; Type: OCSF_VALUE;
	BASE_EVENT_SEVERITY_ID_BASE_EVENT_SEVERITY_ID_OTHER         BASE_EVENT_SEVERITY_ID = 99 // Type: OCSF_VALUE; EnumValue: 99;
)

// Enum value maps for BASE_EVENT_SEVERITY_ID.
var (
	BASE_EVENT_SEVERITY_ID_name = map[int32]string{
		0:  "BASE_EVENT_SEVERITY_ID_UNKNOWN",
		1:  "BASE_EVENT_SEVERITY_ID_INFORMATIONAL",
		2:  "BASE_EVENT_SEVERITY_ID_LOW",
		3:  "BASE_EVENT_SEVERITY_ID_MEDIUM",
		4:  "BASE_EVENT_SEVERITY_ID_HIGH",
		5:  "BASE_EVENT_SEVERITY_ID_CRITICAL",
		6:  "BASE_EVENT_SEVERITY_ID_FATAL",
		99: "BASE_EVENT_SEVERITY_ID_OTHER",
	}
	BASE_EVENT_SEVERITY_ID_value = map[string]int32{
		"BASE_EVENT_SEVERITY_ID_UNKNOWN":       0,
		"BASE_EVENT_SEVERITY_ID_INFORMATIONAL": 1,
		"BASE_EVENT_SEVERITY_ID_LOW":           2,
		"BASE_EVENT_SEVERITY_ID_MEDIUM":        3,
		"BASE_EVENT_SEVERITY_ID_HIGH":          4,
		"BASE_EVENT_SEVERITY_ID_CRITICAL":      5,
		"BASE_EVENT_SEVERITY_ID_FATAL":         6,
		"BASE_EVENT_SEVERITY_ID_OTHER":         99,
	}
)

func (x BASE_EVENT_SEVERITY_ID) Enum() *BASE_EVENT_SEVERITY_ID {
	p := new(BASE_EVENT_SEVERITY_ID)
	*p = x
	return p
}

func (x BASE_EVENT_SEVERITY_ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BASE_EVENT_SEVERITY_ID) Descriptor() protoreflect.EnumDescriptor {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[3].Descriptor()
}

func (BASE_EVENT_SEVERITY_ID) Type() protoreflect.EnumType {
	return &file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[3]
}

func (x BASE_EVENT_SEVERITY_ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BASE_EVENT_SEVERITY_ID.Descriptor instead.
func (BASE_EVENT_SEVERITY_ID) EnumDescriptor() ([]byte, []int) {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescGZIP(), []int{3}
}

type BASE_EVENT_STATUS_ID int32

const (
	BASE_EVENT_STATUS_ID_BASE_EVENT_STATUS_ID_UNKNOWN BASE_EVENT_STATUS_ID = 0  // Type: OCSF_VALUE; EnumValue: 0;
	BASE_EVENT_STATUS_ID_BASE_EVENT_STATUS_ID_SUCCESS BASE_EVENT_STATUS_ID = 1  // Type: OCSF_VALUE; EnumValue: 1;
	BASE_EVENT_STATUS_ID_BASE_EVENT_STATUS_ID_FAILURE BASE_EVENT_STATUS_ID = 2  // Type: OCSF_VALUE; EnumValue: 2;
	BASE_EVENT_STATUS_ID_BASE_EVENT_STATUS_ID_OTHER   BASE_EVENT_STATUS_ID = 99 // Type: OCSF_VALUE; EnumValue: 99;
)

// Enum value maps for BASE_EVENT_STATUS_ID.
var (
	BASE_EVENT_STATUS_ID_name = map[int32]string{
		0:  "BASE_EVENT_STATUS_ID_UNKNOWN",
		1:  "BASE_EVENT_STATUS_ID_SUCCESS",
		2:  "BASE_EVENT_STATUS_ID_FAILURE",
		99: "BASE_EVENT_STATUS_ID_OTHER",
	}
	BASE_EVENT_STATUS_ID_value = map[string]int32{
		"BASE_EVENT_STATUS_ID_UNKNOWN": 0,
		"BASE_EVENT_STATUS_ID_SUCCESS": 1,
		"BASE_EVENT_STATUS_ID_FAILURE": 2,
		"BASE_EVENT_STATUS_ID_OTHER":   99,
	}
)

func (x BASE_EVENT_STATUS_ID) Enum() *BASE_EVENT_STATUS_ID {
	p := new(BASE_EVENT_STATUS_ID)
	*p = x
	return p
}

func (x BASE_EVENT_STATUS_ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BASE_EVENT_STATUS_ID) Descriptor() protoreflect.EnumDescriptor {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[4].Descriptor()
}

func (BASE_EVENT_STATUS_ID) Type() protoreflect.EnumType {
	return &file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes[4]
}

func (x BASE_EVENT_STATUS_ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BASE_EVENT_STATUS_ID.Descriptor instead.
func (BASE_EVENT_STATUS_ID) EnumDescriptor() ([]byte, []int) {
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescGZIP(), []int{4}
}

var File_ocsf_v1_3_0_events_other_enums_enums_proto protoreflect.FileDescriptor

var file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6f, 0x63, 0x73, 0x66, 0x2f, 0x76, 0x31, 0x5f, 0x33, 0x5f, 0x30, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6f, 0x63,
	0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x33, 0x5f, 0x30, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x5e, 0x0a, 0x16,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x44,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x44, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x63, 0x2a, 0x44, 0x0a, 0x17,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x55, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x25, 0x42, 0x41, 0x53, 0x45, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55,
	0x49, 0x44, 0x5f, 0x55, 0x4e, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44,
	0x10, 0x00, 0x2a, 0x3b, 0x0a, 0x14, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x55, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x1f, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x55,
	0x49, 0x44, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x2a,
	0xb3, 0x02, 0x0a, 0x16, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x28,
	0x0a, 0x24, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x42, 0x41, 0x53, 0x45,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x49, 0x44, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x42, 0x41, 0x53, 0x45,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x49, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x42,
	0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x44, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10,
	0x05, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x41, 0x54, 0x41,
	0x4c, 0x10, 0x06, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x5f, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x63, 0x2a, 0x9c, 0x01, 0x0a, 0x14, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x1c, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x5f, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x10, 0x63, 0x42, 0x8d, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x63, 0x73,
	0x66, 0x2e, 0x76, 0x31, 0x5f, 0x33, 0x5f, 0x30, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6c, 0x6c, 0x6c, 0x61, 0x62, 0x68, 0x2f, 0x6f,
	0x63, 0x73, 0x66, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x6f, 0x63, 0x73, 0x66, 0x2f, 0x76, 0x31, 0x5f, 0x33, 0x5f, 0x30, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0xa2, 0x02, 0x05, 0x4f, 0x56, 0x45, 0x4f, 0x45, 0xaa, 0x02, 0x1c, 0x4f, 0x63, 0x73, 0x66, 0x2e,
	0x56, 0x31, 0x33, 0x30, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1c, 0x4f, 0x63, 0x73, 0x66, 0x5c, 0x56,
	0x31, 0x33, 0x30, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xe2, 0x02, 0x28, 0x4f, 0x63, 0x73, 0x66, 0x5c, 0x56, 0x31,
	0x33, 0x30, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x5c,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x20, 0x4f, 0x63, 0x73, 0x66, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x30, 0x3a, 0x3a,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescOnce sync.Once
	file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescData = file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDesc
)

func file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescGZIP() []byte {
	file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescOnce.Do(func() {
		file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescData)
	})
	return file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDescData
}

var file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_ocsf_v1_3_0_events_other_enums_enums_proto_goTypes = []any{
	(BASE_EVENT_ACTIVITY_ID)(0),  // 0: ocsf.v1_3_0.events.other.enums.BASE_EVENT_ACTIVITY_ID
	(BASE_EVENT_CATEGORY_UID)(0), // 1: ocsf.v1_3_0.events.other.enums.BASE_EVENT_CATEGORY_UID
	(BASE_EVENT_CLASS_UID)(0),    // 2: ocsf.v1_3_0.events.other.enums.BASE_EVENT_CLASS_UID
	(BASE_EVENT_SEVERITY_ID)(0),  // 3: ocsf.v1_3_0.events.other.enums.BASE_EVENT_SEVERITY_ID
	(BASE_EVENT_STATUS_ID)(0),    // 4: ocsf.v1_3_0.events.other.enums.BASE_EVENT_STATUS_ID
}
var file_ocsf_v1_3_0_events_other_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ocsf_v1_3_0_events_other_enums_enums_proto_init() }
func file_ocsf_v1_3_0_events_other_enums_enums_proto_init() {
	if File_ocsf_v1_3_0_events_other_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocsf_v1_3_0_events_other_enums_enums_proto_goTypes,
		DependencyIndexes: file_ocsf_v1_3_0_events_other_enums_enums_proto_depIdxs,
		EnumInfos:         file_ocsf_v1_3_0_events_other_enums_enums_proto_enumTypes,
	}.Build()
	File_ocsf_v1_3_0_events_other_enums_enums_proto = out.File
	file_ocsf_v1_3_0_events_other_enums_enums_proto_rawDesc = nil
	file_ocsf_v1_3_0_events_other_enums_enums_proto_goTypes = nil
	file_ocsf_v1_3_0_events_other_enums_enums_proto_depIdxs = nil
}
