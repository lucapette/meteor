// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: odpf/entities/resources/job.proto

package resources

import (
	common "github.com/odpf/meteor/proto/odpf/entities/common"
	facets "github.com/odpf/meteor/proto/odpf/entities/facets"
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

// Job is a resource that represents a job.
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique identifier of the job.
	// For example, "job_1".
	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	// Optional. The name of the job.
	// For example, "Job 1".
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The source of the job.
	// For example, "airflow".
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. The type of the job.
	// For example, "firehose", "stream".
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. The description of the job.
	// For example, "This job is used to process data from a stream."
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. The ownership of the topic.
	// For an example check out ownership.
	Ownership *facets.Ownership `protobuf:"bytes,31,opt,name=ownership,proto3" json:"ownership,omitempty"`
	// Optional. The lineage of the topic.
	// For an example check out lineage schema.
	Lineage *facets.Lineage `protobuf:"bytes,32,opt,name=lineage,proto3" json:"lineage,omitempty"`
	// Optional. List of the user's custom properties.
	// Properties facet can be used to set custom properties, tags and labels for a user.
	// Properties are defined in the properties schema.
	Properties *facets.Properties `protobuf:"bytes,33,opt,name=properties,proto3" json:"properties,omitempty"`
	// Optional. The timestamp of the user's creation.
	// Timstamp facet can be used to set the creation and updation timestamp of a user.
	// Timestamps are defined in the timestamp schema.
	Timestamps *common.Timestamp `protobuf:"bytes,34,opt,name=timestamps,proto3" json:"timestamps,omitempty"`
	// Optional. The timestamp of the generated event.
	// Event schemas is defined in the common event schema.
	Event *common.Event `protobuf:"bytes,100,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_entities_resources_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_entities_resources_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_odpf_entities_resources_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Job) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Job) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Job) GetOwnership() *facets.Ownership {
	if x != nil {
		return x.Ownership
	}
	return nil
}

func (x *Job) GetLineage() *facets.Lineage {
	if x != nil {
		return x.Lineage
	}
	return nil
}

func (x *Job) GetProperties() *facets.Properties {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Job) GetTimestamps() *common.Timestamp {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

func (x *Job) GetEvent() *common.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_odpf_entities_resources_job_proto protoreflect.FileDescriptor

var file_odpf_entities_resources_job_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x24, 0x6f, 0x64,
	0x70, 0x66, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65,
	0x74, 0x73, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6f,
	0x64, 0x70, 0x66, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x03, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x37, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6f, 0x64, 0x70, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x66, 0x61,
	0x63, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x12, 0x31, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x64,
	0x70, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x51, 0x0a, 0x1a, 0x69, 0x6f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x08, 0x4a,
	0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_odpf_entities_resources_job_proto_rawDescOnce sync.Once
	file_odpf_entities_resources_job_proto_rawDescData = file_odpf_entities_resources_job_proto_rawDesc
)

func file_odpf_entities_resources_job_proto_rawDescGZIP() []byte {
	file_odpf_entities_resources_job_proto_rawDescOnce.Do(func() {
		file_odpf_entities_resources_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_entities_resources_job_proto_rawDescData)
	})
	return file_odpf_entities_resources_job_proto_rawDescData
}

var file_odpf_entities_resources_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_odpf_entities_resources_job_proto_goTypes = []interface{}{
	(*Job)(nil),               // 0: odpf.entities.resources.Job
	(*facets.Ownership)(nil),  // 1: odpf.entities.facets.Ownership
	(*facets.Lineage)(nil),    // 2: odpf.entities.facets.Lineage
	(*facets.Properties)(nil), // 3: odpf.entities.facets.Properties
	(*common.Timestamp)(nil),  // 4: odpf.entities.common.Timestamp
	(*common.Event)(nil),      // 5: odpf.entities.common.Event
}
var file_odpf_entities_resources_job_proto_depIdxs = []int32{
	1, // 0: odpf.entities.resources.Job.ownership:type_name -> odpf.entities.facets.Ownership
	2, // 1: odpf.entities.resources.Job.lineage:type_name -> odpf.entities.facets.Lineage
	3, // 2: odpf.entities.resources.Job.properties:type_name -> odpf.entities.facets.Properties
	4, // 3: odpf.entities.resources.Job.timestamps:type_name -> odpf.entities.common.Timestamp
	5, // 4: odpf.entities.resources.Job.event:type_name -> odpf.entities.common.Event
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_odpf_entities_resources_job_proto_init() }
func file_odpf_entities_resources_job_proto_init() {
	if File_odpf_entities_resources_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_odpf_entities_resources_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
			RawDescriptor: file_odpf_entities_resources_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_entities_resources_job_proto_goTypes,
		DependencyIndexes: file_odpf_entities_resources_job_proto_depIdxs,
		MessageInfos:      file_odpf_entities_resources_job_proto_msgTypes,
	}.Build()
	File_odpf_entities_resources_job_proto = out.File
	file_odpf_entities_resources_job_proto_rawDesc = nil
	file_odpf_entities_resources_job_proto_goTypes = nil
	file_odpf_entities_resources_job_proto_depIdxs = nil
}
