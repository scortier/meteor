// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: odpf/assets/table.proto

package assets

import (
	common "github.com/odpf/meteor/models/odpf/assets/common"
	facets "github.com/odpf/meteor/models/odpf/assets/facets"
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

// Table is a table in a database.
// It can be a file, a table, a view, a materialized view, a temporary table, or a virtual table.
type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Representation of the resource
	Resource *common.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The metrics about the table.
	// For example check the profile schem.
	Profile *TableProfile `protobuf:"bytes,21,opt,name=profile,proto3" json:"profile,omitempty"`
	// The columns of the table.
	// Example: 'id', `name`, `age'.
	Schema *facets.Columns `protobuf:"bytes,22,opt,name=schema,proto3" json:"schema,omitempty"`
	// Previews of the table.
	// For an example check out preview facet.
	Preview *facets.Preview `protobuf:"bytes,23,opt,name=preview,proto3" json:"preview,omitempty"`
	// The ownership of the topic.
	// For an example check out ownership.
	Ownership *facets.Ownership `protobuf:"bytes,31,opt,name=ownership,proto3" json:"ownership,omitempty"`
	// The lineage of the topic.
	// For an example check out lineage.
	Lineage *facets.Lineage `protobuf:"bytes,32,opt,name=lineage,proto3" json:"lineage,omitempty"`
	// List of the user's custom properties.
	// Properties facet can be used to set custom properties, tags and labels for a user.
	Properties *facets.Properties `protobuf:"bytes,33,opt,name=properties,proto3" json:"properties,omitempty"`
	// The timestamp of the user's creation.
	// Timstamp facet can be used to set the creation and updation timestamp of a user.
	Timestamps *common.Timestamp `protobuf:"bytes,34,opt,name=timestamps,proto3" json:"timestamps,omitempty"`
	// The timestamp of the generated event.
	// Event schemas is defined in the common event schema.
	Event *common.Event `protobuf:"bytes,100,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_odpf_assets_table_proto_rawDescGZIP(), []int{0}
}

func (x *Table) GetResource() *common.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Table) GetProfile() *TableProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *Table) GetSchema() *facets.Columns {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *Table) GetPreview() *facets.Preview {
	if x != nil {
		return x.Preview
	}
	return nil
}

func (x *Table) GetOwnership() *facets.Ownership {
	if x != nil {
		return x.Ownership
	}
	return nil
}

func (x *Table) GetLineage() *facets.Lineage {
	if x != nil {
		return x.Lineage
	}
	return nil
}

func (x *Table) GetProperties() *facets.Properties {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Table) GetTimestamps() *common.Timestamp {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

func (x *Table) GetEvent() *common.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// TableProfile is the metrics about the table.
type TableProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of rows in the table.
	// Example: `100`.
	TotalRows int64 `protobuf:"varint,1,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	// The number of rows in the table that are not deleted.
	// Example: `event_timestamp`.
	PartitionKey   string   `protobuf:"bytes,2,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	PartitionValue string   `protobuf:"bytes,3,opt,name=partition_value,json=partitionValue,proto3" json:"partition_value,omitempty"`
	UsageCount     int64    `protobuf:"varint,4,opt,name=usage_count,json=usageCount,proto3" json:"usage_count,omitempty"`
	Joins          []*Join  `protobuf:"bytes,5,rep,name=joins,proto3" json:"joins,omitempty"`
	Filters        []string `protobuf:"bytes,6,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *TableProfile) Reset() {
	*x = TableProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableProfile) ProtoMessage() {}

func (x *TableProfile) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableProfile.ProtoReflect.Descriptor instead.
func (*TableProfile) Descriptor() ([]byte, []int) {
	return file_odpf_assets_table_proto_rawDescGZIP(), []int{1}
}

func (x *TableProfile) GetTotalRows() int64 {
	if x != nil {
		return x.TotalRows
	}
	return 0
}

func (x *TableProfile) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

func (x *TableProfile) GetPartitionValue() string {
	if x != nil {
		return x.PartitionValue
	}
	return ""
}

func (x *TableProfile) GetUsageCount() int64 {
	if x != nil {
		return x.UsageCount
	}
	return 0
}

func (x *TableProfile) GetJoins() []*Join {
	if x != nil {
		return x.Joins
	}
	return nil
}

func (x *TableProfile) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

// Join is the metric of which are other tables that are joined with this table
type Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn        string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Count      int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Conditions []string `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *Join) Reset() {
	*x = Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join) ProtoMessage() {}

func (x *Join) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join.ProtoReflect.Descriptor instead.
func (*Join) Descriptor() ([]byte, []int) {
	return file_odpf_assets_table_proto_rawDescGZIP(), []int{2}
}

func (x *Join) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Join) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Join) GetConditions() []string {
	if x != nil {
		return x.Conditions
	}
	return nil
}

var File_odpf_assets_table_proto protoreflect.FileDescriptor

var file_odpf_assets_table_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x64, 0x70, 0x66, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x1a, 0x1f, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6f, 0x64, 0x70, 0x66, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6f,
	0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74,
	0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x63,
	0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x64, 0x70,
	0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x05,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x64, 0x70,
	0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x3b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x35, 0x0a,
	0x07, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x66, 0x61, 0x63,
	0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6c, 0x69, 0x6e,
	0x65, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x73, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x6f, 0x77, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x6a, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x6a, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x4e, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x3b, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x6f, 0x64, 0x70,
	0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x42, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_odpf_assets_table_proto_rawDescOnce sync.Once
	file_odpf_assets_table_proto_rawDescData = file_odpf_assets_table_proto_rawDesc
)

func file_odpf_assets_table_proto_rawDescGZIP() []byte {
	file_odpf_assets_table_proto_rawDescOnce.Do(func() {
		file_odpf_assets_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_assets_table_proto_rawDescData)
	})
	return file_odpf_assets_table_proto_rawDescData
}

var file_odpf_assets_table_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_odpf_assets_table_proto_goTypes = []interface{}{
	(*Table)(nil),             // 0: odpf.assets.Table
	(*TableProfile)(nil),      // 1: odpf.assets.TableProfile
	(*Join)(nil),              // 2: odpf.assets.Join
	(*common.Resource)(nil),   // 3: odpf.assets.common.Resource
	(*facets.Columns)(nil),    // 4: odpf.assets.facets.Columns
	(*facets.Preview)(nil),    // 5: odpf.assets.facets.Preview
	(*facets.Ownership)(nil),  // 6: odpf.assets.facets.Ownership
	(*facets.Lineage)(nil),    // 7: odpf.assets.facets.Lineage
	(*facets.Properties)(nil), // 8: odpf.assets.facets.Properties
	(*common.Timestamp)(nil),  // 9: odpf.assets.common.Timestamp
	(*common.Event)(nil),      // 10: odpf.assets.common.Event
}
var file_odpf_assets_table_proto_depIdxs = []int32{
	3,  // 0: odpf.assets.Table.resource:type_name -> odpf.assets.common.Resource
	1,  // 1: odpf.assets.Table.profile:type_name -> odpf.assets.TableProfile
	4,  // 2: odpf.assets.Table.schema:type_name -> odpf.assets.facets.Columns
	5,  // 3: odpf.assets.Table.preview:type_name -> odpf.assets.facets.Preview
	6,  // 4: odpf.assets.Table.ownership:type_name -> odpf.assets.facets.Ownership
	7,  // 5: odpf.assets.Table.lineage:type_name -> odpf.assets.facets.Lineage
	8,  // 6: odpf.assets.Table.properties:type_name -> odpf.assets.facets.Properties
	9,  // 7: odpf.assets.Table.timestamps:type_name -> odpf.assets.common.Timestamp
	10, // 8: odpf.assets.Table.event:type_name -> odpf.assets.common.Event
	2,  // 9: odpf.assets.TableProfile.joins:type_name -> odpf.assets.Join
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_odpf_assets_table_proto_init() }
func file_odpf_assets_table_proto_init() {
	if File_odpf_assets_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_odpf_assets_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_odpf_assets_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableProfile); i {
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
		file_odpf_assets_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join); i {
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
			RawDescriptor: file_odpf_assets_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_assets_table_proto_goTypes,
		DependencyIndexes: file_odpf_assets_table_proto_depIdxs,
		MessageInfos:      file_odpf_assets_table_proto_msgTypes,
	}.Build()
	File_odpf_assets_table_proto = out.File
	file_odpf_assets_table_proto_rawDesc = nil
	file_odpf_assets_table_proto_goTypes = nil
	file_odpf_assets_table_proto_depIdxs = nil
}
