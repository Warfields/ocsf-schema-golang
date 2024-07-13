// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: ocsf/v1_2_0/events/other/other.proto

package other

import (
	enums "github.com/valllabh/ocsf-schema-golang/ocsf/v1_2_0/events/other/enums"
	objects "github.com/valllabh/ocsf-schema-golang/ocsf/v1_2_0/objects"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Event: other
// Event UID: 0
// URL: https://schema.ocsf.io/1.2.0/classes/base_event
type BaseEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId     enums.BASE_EVENT_ACTIVITY_ID  `protobuf:"varint,1,opt,name=activity_id,json=activityId,proto3,enum=ocsf.v1_2_0.events.other.enums.BASE_EVENT_ACTIVITY_ID" json:"activity_id,omitempty"`     // Caption: Activity ID;
	ActivityName   string                        `protobuf:"bytes,2,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`                                                           // Caption: Activity;
	Api            *objects.Api                  `protobuf:"bytes,3,opt,name=api,proto3" json:"api,omitempty"`                                                                                                 // Caption: API Details; Profile: cloud;
	CategoryName   string                        `protobuf:"bytes,4,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`                                                           // Caption: Category;
	CategoryUid    enums.BASE_EVENT_CATEGORY_UID `protobuf:"varint,5,opt,name=category_uid,json=categoryUid,proto3,enum=ocsf.v1_2_0.events.other.enums.BASE_EVENT_CATEGORY_UID" json:"category_uid,omitempty"` // Caption: Category ID;
	ClassName      string                        `protobuf:"bytes,6,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`                                                                    // Caption: Class;
	ClassUid       enums.BASE_EVENT_CLASS_UID    `protobuf:"varint,7,opt,name=class_uid,json=classUid,proto3,enum=ocsf.v1_2_0.events.other.enums.BASE_EVENT_CLASS_UID" json:"class_uid,omitempty"`             // Caption: Class ID;
	Cloud          *objects.Cloud                `protobuf:"bytes,8,opt,name=cloud,proto3" json:"cloud,omitempty"`                                                                                             // Caption: Cloud; Profile: cloud;
	Count          int32                         `protobuf:"varint,9,opt,name=count,proto3" json:"count,omitempty"`                                                                                            // Caption: Count;
	Duration       int32                         `protobuf:"varint,10,opt,name=duration,proto3" json:"duration,omitempty"`                                                                                     // Caption: Duration;
	EndTime        int64                         `protobuf:"varint,11,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`                                                                        // Caption: End Time;
	Enrichments    []*objects.Enrichment         `protobuf:"bytes,12,rep,name=enrichments,proto3" json:"enrichments,omitempty"`                                                                                // Caption: Enrichments;
	Message        string                        `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`                                                                                        // Caption: Message;
	Metadata       *objects.Metadata             `protobuf:"bytes,14,opt,name=metadata,proto3" json:"metadata,omitempty"`                                                                                      // Caption: Metadata;
	Observables    []*objects.Observable         `protobuf:"bytes,15,rep,name=observables,proto3" json:"observables,omitempty"`                                                                                // Caption: Observables;
	RawData        string                        `protobuf:"bytes,16,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`                                                                         // Caption: Raw Data;
	Severity       string                        `protobuf:"bytes,17,opt,name=severity,proto3" json:"severity,omitempty"`                                                                                      // Caption: Severity;
	SeverityId     enums.BASE_EVENT_SEVERITY_ID  `protobuf:"varint,18,opt,name=severity_id,json=severityId,proto3,enum=ocsf.v1_2_0.events.other.enums.BASE_EVENT_SEVERITY_ID" json:"severity_id,omitempty"`    // Caption: Severity ID;
	StartTime      int64                         `protobuf:"varint,19,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`                                                                  // Caption: Start Time;
	Status         string                        `protobuf:"bytes,20,opt,name=status,proto3" json:"status,omitempty"`                                                                                          // Caption: Status;
	StatusCode     string                        `protobuf:"bytes,21,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`                                                                // Caption: Status Code;
	StatusDetail   string                        `protobuf:"bytes,22,opt,name=status_detail,json=statusDetail,proto3" json:"status_detail,omitempty"`                                                          // Caption: Status Details;
	StatusId       enums.BASE_EVENT_STATUS_ID    `protobuf:"varint,23,opt,name=status_id,json=statusId,proto3,enum=ocsf.v1_2_0.events.other.enums.BASE_EVENT_STATUS_ID" json:"status_id,omitempty"`            // Caption: Status ID;
	Time           int64                         `protobuf:"varint,24,opt,name=time,proto3" json:"time,omitempty"`                                                                                             // Caption: Event Time;
	TimezoneOffset int32                         `protobuf:"varint,25,opt,name=timezone_offset,json=timezoneOffset,proto3" json:"timezone_offset,omitempty"`                                                   // Caption: Timezone Offset;
	TypeName       string                        `protobuf:"bytes,26,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`                                                                      // Caption: Type Name;
	TypeUid        int64                         `protobuf:"varint,27,opt,name=type_uid,json=typeUid,proto3" json:"type_uid,omitempty"`                                                                        // Caption: Type ID;
	Unmapped       *structpb.Struct              `protobuf:"bytes,28,opt,name=unmapped,proto3" json:"unmapped,omitempty"`                                                                                      // Caption: Unmapped Data;
}

func (x *BaseEvent) Reset() {
	*x = BaseEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocsf_v1_2_0_events_other_other_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseEvent) ProtoMessage() {}

func (x *BaseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ocsf_v1_2_0_events_other_other_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseEvent.ProtoReflect.Descriptor instead.
func (*BaseEvent) Descriptor() ([]byte, []int) {
	return file_ocsf_v1_2_0_events_other_other_proto_rawDescGZIP(), []int{0}
}

func (x *BaseEvent) GetActivityId() enums.BASE_EVENT_ACTIVITY_ID {
	if x != nil {
		return x.ActivityId
	}
	return enums.BASE_EVENT_ACTIVITY_ID(0)
}

func (x *BaseEvent) GetActivityName() string {
	if x != nil {
		return x.ActivityName
	}
	return ""
}

func (x *BaseEvent) GetApi() *objects.Api {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *BaseEvent) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *BaseEvent) GetCategoryUid() enums.BASE_EVENT_CATEGORY_UID {
	if x != nil {
		return x.CategoryUid
	}
	return enums.BASE_EVENT_CATEGORY_UID(0)
}

func (x *BaseEvent) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *BaseEvent) GetClassUid() enums.BASE_EVENT_CLASS_UID {
	if x != nil {
		return x.ClassUid
	}
	return enums.BASE_EVENT_CLASS_UID(0)
}

func (x *BaseEvent) GetCloud() *objects.Cloud {
	if x != nil {
		return x.Cloud
	}
	return nil
}

func (x *BaseEvent) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *BaseEvent) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *BaseEvent) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *BaseEvent) GetEnrichments() []*objects.Enrichment {
	if x != nil {
		return x.Enrichments
	}
	return nil
}

func (x *BaseEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaseEvent) GetMetadata() *objects.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BaseEvent) GetObservables() []*objects.Observable {
	if x != nil {
		return x.Observables
	}
	return nil
}

func (x *BaseEvent) GetRawData() string {
	if x != nil {
		return x.RawData
	}
	return ""
}

func (x *BaseEvent) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *BaseEvent) GetSeverityId() enums.BASE_EVENT_SEVERITY_ID {
	if x != nil {
		return x.SeverityId
	}
	return enums.BASE_EVENT_SEVERITY_ID(0)
}

func (x *BaseEvent) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *BaseEvent) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BaseEvent) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *BaseEvent) GetStatusDetail() string {
	if x != nil {
		return x.StatusDetail
	}
	return ""
}

func (x *BaseEvent) GetStatusId() enums.BASE_EVENT_STATUS_ID {
	if x != nil {
		return x.StatusId
	}
	return enums.BASE_EVENT_STATUS_ID(0)
}

func (x *BaseEvent) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *BaseEvent) GetTimezoneOffset() int32 {
	if x != nil {
		return x.TimezoneOffset
	}
	return 0
}

func (x *BaseEvent) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *BaseEvent) GetTypeUid() int64 {
	if x != nil {
		return x.TypeUid
	}
	return 0
}

func (x *BaseEvent) GetUnmapped() *structpb.Struct {
	if x != nil {
		return x.Unmapped
	}
	return nil
}

var File_ocsf_v1_2_0_events_other_other_proto protoreflect.FileDescriptor

var file_ocsf_v1_2_0_events_other_other_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x63, 0x73, 0x66, 0x2f, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f,
	0x32, 0x5f, 0x30, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x1a, 0x2a, 0x6f, 0x63, 0x73, 0x66, 0x2f, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6f, 0x63,
	0x73, 0x66, 0x2f, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x0a,
	0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x36, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x70, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31,
	0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x37, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x49, 0x44, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e,
	0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x5f, 0x55, 0x49, 0x44, 0x52, 0x08,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x55, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76,
	0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f,
	0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31,
	0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x41, 0x0a, 0x0b, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x32,
	0x5f, 0x30, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0b, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x57, 0x0a, 0x0b, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36,
	0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x49, 0x44, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x51, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x34, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x74, 0x79, 0x70, 0x65, 0x55, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x75, 0x6e, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x64, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x08, 0x75, 0x6e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x42, 0xe7, 0x01, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x63, 0x73, 0x66, 0x2e, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x0a, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6c, 0x6c, 0x6c, 0x61, 0x62, 0x68,
	0x2f, 0x6f, 0x63, 0x73, 0x66, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x63, 0x73, 0x66, 0x2f, 0x76, 0x31, 0x5f, 0x32, 0x5f, 0x30, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0xa2, 0x02, 0x04, 0x4f,
	0x56, 0x45, 0x4f, 0xaa, 0x02, 0x16, 0x4f, 0x63, 0x73, 0x66, 0x2e, 0x56, 0x31, 0x32, 0x30, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0xca, 0x02, 0x16, 0x4f,
	0x63, 0x73, 0x66, 0x5c, 0x56, 0x31, 0x32, 0x30, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c,
	0x4f, 0x74, 0x68, 0x65, 0x72, 0xe2, 0x02, 0x22, 0x4f, 0x63, 0x73, 0x66, 0x5c, 0x56, 0x31, 0x32,
	0x30, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x4f, 0x63, 0x73,
	0x66, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x30, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a,
	0x3a, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocsf_v1_2_0_events_other_other_proto_rawDescOnce sync.Once
	file_ocsf_v1_2_0_events_other_other_proto_rawDescData = file_ocsf_v1_2_0_events_other_other_proto_rawDesc
)

func file_ocsf_v1_2_0_events_other_other_proto_rawDescGZIP() []byte {
	file_ocsf_v1_2_0_events_other_other_proto_rawDescOnce.Do(func() {
		file_ocsf_v1_2_0_events_other_other_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocsf_v1_2_0_events_other_other_proto_rawDescData)
	})
	return file_ocsf_v1_2_0_events_other_other_proto_rawDescData
}

var file_ocsf_v1_2_0_events_other_other_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ocsf_v1_2_0_events_other_other_proto_goTypes = []any{
	(*BaseEvent)(nil),                  // 0: ocsf.v1_2_0.events.other.BaseEvent
	(enums.BASE_EVENT_ACTIVITY_ID)(0),  // 1: ocsf.v1_2_0.events.other.enums.BASE_EVENT_ACTIVITY_ID
	(*objects.Api)(nil),                // 2: ocsf.v1_2_0.objects.Api
	(enums.BASE_EVENT_CATEGORY_UID)(0), // 3: ocsf.v1_2_0.events.other.enums.BASE_EVENT_CATEGORY_UID
	(enums.BASE_EVENT_CLASS_UID)(0),    // 4: ocsf.v1_2_0.events.other.enums.BASE_EVENT_CLASS_UID
	(*objects.Cloud)(nil),              // 5: ocsf.v1_2_0.objects.Cloud
	(*objects.Enrichment)(nil),         // 6: ocsf.v1_2_0.objects.Enrichment
	(*objects.Metadata)(nil),           // 7: ocsf.v1_2_0.objects.Metadata
	(*objects.Observable)(nil),         // 8: ocsf.v1_2_0.objects.Observable
	(enums.BASE_EVENT_SEVERITY_ID)(0),  // 9: ocsf.v1_2_0.events.other.enums.BASE_EVENT_SEVERITY_ID
	(enums.BASE_EVENT_STATUS_ID)(0),    // 10: ocsf.v1_2_0.events.other.enums.BASE_EVENT_STATUS_ID
	(*structpb.Struct)(nil),            // 11: google.protobuf.Struct
}
var file_ocsf_v1_2_0_events_other_other_proto_depIdxs = []int32{
	1,  // 0: ocsf.v1_2_0.events.other.BaseEvent.activity_id:type_name -> ocsf.v1_2_0.events.other.enums.BASE_EVENT_ACTIVITY_ID
	2,  // 1: ocsf.v1_2_0.events.other.BaseEvent.api:type_name -> ocsf.v1_2_0.objects.Api
	3,  // 2: ocsf.v1_2_0.events.other.BaseEvent.category_uid:type_name -> ocsf.v1_2_0.events.other.enums.BASE_EVENT_CATEGORY_UID
	4,  // 3: ocsf.v1_2_0.events.other.BaseEvent.class_uid:type_name -> ocsf.v1_2_0.events.other.enums.BASE_EVENT_CLASS_UID
	5,  // 4: ocsf.v1_2_0.events.other.BaseEvent.cloud:type_name -> ocsf.v1_2_0.objects.Cloud
	6,  // 5: ocsf.v1_2_0.events.other.BaseEvent.enrichments:type_name -> ocsf.v1_2_0.objects.Enrichment
	7,  // 6: ocsf.v1_2_0.events.other.BaseEvent.metadata:type_name -> ocsf.v1_2_0.objects.Metadata
	8,  // 7: ocsf.v1_2_0.events.other.BaseEvent.observables:type_name -> ocsf.v1_2_0.objects.Observable
	9,  // 8: ocsf.v1_2_0.events.other.BaseEvent.severity_id:type_name -> ocsf.v1_2_0.events.other.enums.BASE_EVENT_SEVERITY_ID
	10, // 9: ocsf.v1_2_0.events.other.BaseEvent.status_id:type_name -> ocsf.v1_2_0.events.other.enums.BASE_EVENT_STATUS_ID
	11, // 10: ocsf.v1_2_0.events.other.BaseEvent.unmapped:type_name -> google.protobuf.Struct
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_ocsf_v1_2_0_events_other_other_proto_init() }
func file_ocsf_v1_2_0_events_other_other_proto_init() {
	if File_ocsf_v1_2_0_events_other_other_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocsf_v1_2_0_events_other_other_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BaseEvent); i {
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
			RawDescriptor: file_ocsf_v1_2_0_events_other_other_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocsf_v1_2_0_events_other_other_proto_goTypes,
		DependencyIndexes: file_ocsf_v1_2_0_events_other_other_proto_depIdxs,
		MessageInfos:      file_ocsf_v1_2_0_events_other_other_proto_msgTypes,
	}.Build()
	File_ocsf_v1_2_0_events_other_other_proto = out.File
	file_ocsf_v1_2_0_events_other_other_proto_rawDesc = nil
	file_ocsf_v1_2_0_events_other_other_proto_goTypes = nil
	file_ocsf_v1_2_0_events_other_other_proto_depIdxs = nil
}