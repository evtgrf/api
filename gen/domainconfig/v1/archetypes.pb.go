// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: domainconfig/v1/archetypes.proto

package domainconfigv1

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

type QueryTerm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentEntity int64                 `protobuf:"varint,1,opt,name=component_entity,json=componentEntity,proto3" json:"component_entity,omitempty"`
	FromEntity      *QueryTerm_Connection `protobuf:"bytes,2,opt,name=from_entity,json=fromEntity,proto3" json:"from_entity,omitempty"`
	ToEntity        *QueryTerm_Connection `protobuf:"bytes,3,opt,name=to_entity,json=toEntity,proto3" json:"to_entity,omitempty"`
}

func (x *QueryTerm) Reset() {
	*x = QueryTerm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domainconfig_v1_archetypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTerm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTerm) ProtoMessage() {}

func (x *QueryTerm) ProtoReflect() protoreflect.Message {
	mi := &file_domainconfig_v1_archetypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTerm.ProtoReflect.Descriptor instead.
func (*QueryTerm) Descriptor() ([]byte, []int) {
	return file_domainconfig_v1_archetypes_proto_rawDescGZIP(), []int{0}
}

func (x *QueryTerm) GetComponentEntity() int64 {
	if x != nil {
		return x.ComponentEntity
	}
	return 0
}

func (x *QueryTerm) GetFromEntity() *QueryTerm_Connection {
	if x != nil {
		return x.FromEntity
	}
	return nil
}

func (x *QueryTerm) GetToEntity() *QueryTerm_Connection {
	if x != nil {
		return x.ToEntity
	}
	return nil
}

type ArchetypeDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata   *Metadata    `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	QueryTerms []*QueryTerm `protobuf:"bytes,2,rep,name=query_terms,json=queryTerms,proto3" json:"query_terms,omitempty"`
}

func (x *ArchetypeDefinition) Reset() {
	*x = ArchetypeDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domainconfig_v1_archetypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchetypeDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchetypeDefinition) ProtoMessage() {}

func (x *ArchetypeDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_domainconfig_v1_archetypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchetypeDefinition.ProtoReflect.Descriptor instead.
func (*ArchetypeDefinition) Descriptor() ([]byte, []int) {
	return file_domainconfig_v1_archetypes_proto_rawDescGZIP(), []int{1}
}

func (x *ArchetypeDefinition) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ArchetypeDefinition) GetQueryTerms() []*QueryTerm {
	if x != nil {
		return x.QueryTerms
	}
	return nil
}

type QueryTerm_Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mode:
	//
	//	*QueryTerm_Connection_Empty
	//	*QueryTerm_Connection_Instance
	//	*QueryTerm_Connection_InitVariable
	//	*QueryTerm_Connection_Variable
	Mode isQueryTerm_Connection_Mode `protobuf_oneof:"mode"`
}

func (x *QueryTerm_Connection) Reset() {
	*x = QueryTerm_Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domainconfig_v1_archetypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTerm_Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTerm_Connection) ProtoMessage() {}

func (x *QueryTerm_Connection) ProtoReflect() protoreflect.Message {
	mi := &file_domainconfig_v1_archetypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTerm_Connection.ProtoReflect.Descriptor instead.
func (*QueryTerm_Connection) Descriptor() ([]byte, []int) {
	return file_domainconfig_v1_archetypes_proto_rawDescGZIP(), []int{0, 0}
}

func (m *QueryTerm_Connection) GetMode() isQueryTerm_Connection_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (x *QueryTerm_Connection) GetEmpty() bool {
	if x, ok := x.GetMode().(*QueryTerm_Connection_Empty); ok {
		return x.Empty
	}
	return false
}

func (x *QueryTerm_Connection) GetInstance() int64 {
	if x, ok := x.GetMode().(*QueryTerm_Connection_Instance); ok {
		return x.Instance
	}
	return 0
}

func (x *QueryTerm_Connection) GetInitVariable() string {
	if x, ok := x.GetMode().(*QueryTerm_Connection_InitVariable); ok {
		return x.InitVariable
	}
	return ""
}

func (x *QueryTerm_Connection) GetVariable() string {
	if x, ok := x.GetMode().(*QueryTerm_Connection_Variable); ok {
		return x.Variable
	}
	return ""
}

type isQueryTerm_Connection_Mode interface {
	isQueryTerm_Connection_Mode()
}

type QueryTerm_Connection_Empty struct {
	Empty bool `protobuf:"varint,1,opt,name=empty,proto3,oneof"`
}

type QueryTerm_Connection_Instance struct {
	Instance int64 `protobuf:"varint,2,opt,name=instance,proto3,oneof"`
}

type QueryTerm_Connection_InitVariable struct {
	InitVariable string `protobuf:"bytes,3,opt,name=init_variable,json=initVariable,proto3,oneof"`
}

type QueryTerm_Connection_Variable struct {
	Variable string `protobuf:"bytes,4,opt,name=variable,proto3,oneof"`
}

func (*QueryTerm_Connection_Empty) isQueryTerm_Connection_Mode() {}

func (*QueryTerm_Connection_Instance) isQueryTerm_Connection_Mode() {}

func (*QueryTerm_Connection_InitVariable) isQueryTerm_Connection_Mode() {}

func (*QueryTerm_Connection_Variable) isQueryTerm_Connection_Mode() {}

var File_domainconfig_v1_archetypes_proto protoreflect.FileDescriptor

var file_domainconfig_v1_archetypes_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x42, 0x0a, 0x09, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x72,
	0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x74, 0x6f,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x8f, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1c, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x69,
	0x6e, 0x69, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x06, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x41, 0x72, 0x63,
	0x68, 0x65, 0x74, 0x79, 0x70, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54,
	0x65, 0x72, 0x6d, 0x73, 0x42, 0xc1, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x41, 0x72,
	0x63, 0x68, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x44, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domainconfig_v1_archetypes_proto_rawDescOnce sync.Once
	file_domainconfig_v1_archetypes_proto_rawDescData = file_domainconfig_v1_archetypes_proto_rawDesc
)

func file_domainconfig_v1_archetypes_proto_rawDescGZIP() []byte {
	file_domainconfig_v1_archetypes_proto_rawDescOnce.Do(func() {
		file_domainconfig_v1_archetypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_domainconfig_v1_archetypes_proto_rawDescData)
	})
	return file_domainconfig_v1_archetypes_proto_rawDescData
}

var file_domainconfig_v1_archetypes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_domainconfig_v1_archetypes_proto_goTypes = []interface{}{
	(*QueryTerm)(nil),            // 0: domainconfig.v1.QueryTerm
	(*ArchetypeDefinition)(nil),  // 1: domainconfig.v1.ArchetypeDefinition
	(*QueryTerm_Connection)(nil), // 2: domainconfig.v1.QueryTerm.Connection
	(*Metadata)(nil),             // 3: domainconfig.v1.Metadata
}
var file_domainconfig_v1_archetypes_proto_depIdxs = []int32{
	2, // 0: domainconfig.v1.QueryTerm.from_entity:type_name -> domainconfig.v1.QueryTerm.Connection
	2, // 1: domainconfig.v1.QueryTerm.to_entity:type_name -> domainconfig.v1.QueryTerm.Connection
	3, // 2: domainconfig.v1.ArchetypeDefinition.metadata:type_name -> domainconfig.v1.Metadata
	0, // 3: domainconfig.v1.ArchetypeDefinition.query_terms:type_name -> domainconfig.v1.QueryTerm
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_domainconfig_v1_archetypes_proto_init() }
func file_domainconfig_v1_archetypes_proto_init() {
	if File_domainconfig_v1_archetypes_proto != nil {
		return
	}
	file_domainconfig_v1_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_domainconfig_v1_archetypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTerm); i {
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
		file_domainconfig_v1_archetypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchetypeDefinition); i {
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
		file_domainconfig_v1_archetypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTerm_Connection); i {
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
	file_domainconfig_v1_archetypes_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*QueryTerm_Connection_Empty)(nil),
		(*QueryTerm_Connection_Instance)(nil),
		(*QueryTerm_Connection_InitVariable)(nil),
		(*QueryTerm_Connection_Variable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domainconfig_v1_archetypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domainconfig_v1_archetypes_proto_goTypes,
		DependencyIndexes: file_domainconfig_v1_archetypes_proto_depIdxs,
		MessageInfos:      file_domainconfig_v1_archetypes_proto_msgTypes,
	}.Build()
	File_domainconfig_v1_archetypes_proto = out.File
	file_domainconfig_v1_archetypes_proto_rawDesc = nil
	file_domainconfig_v1_archetypes_proto_goTypes = nil
	file_domainconfig_v1_archetypes_proto_depIdxs = nil
}
