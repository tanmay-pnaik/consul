// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/computed_explicit_destinations.proto

package meshv2beta1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

type ComputedExplicitDestinations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// destinations is the list of explicit destinations to define for the selected workloads.
	Destinations []*Destination `protobuf:"bytes,1,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *ComputedExplicitDestinations) Reset() {
	*x = ComputedExplicitDestinations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_computed_explicit_destinations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedExplicitDestinations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedExplicitDestinations) ProtoMessage() {}

func (x *ComputedExplicitDestinations) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_computed_explicit_destinations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedExplicitDestinations.ProtoReflect.Descriptor instead.
func (*ComputedExplicitDestinations) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescGZIP(), []int{0}
}

func (x *ComputedExplicitDestinations) GetDestinations() []*Destination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

var File_pbmesh_v2beta1_computed_explicit_destinations_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x21, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x64, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x42, 0xa2, 0x02,
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x21, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x45, 0x78, 0x70,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x48, 0x43, 0x4d, 0xaa, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xe2, 0x02, 0x29, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescData = file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDesc
)

func file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDescData
}

var file_pbmesh_v2beta1_computed_explicit_destinations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbmesh_v2beta1_computed_explicit_destinations_proto_goTypes = []interface{}{
	(*ComputedExplicitDestinations)(nil), // 0: hashicorp.consul.mesh.v2beta1.ComputedExplicitDestinations
	(*Destination)(nil),                  // 1: hashicorp.consul.mesh.v2beta1.Destination
}
var file_pbmesh_v2beta1_computed_explicit_destinations_proto_depIdxs = []int32{
	1, // 0: hashicorp.consul.mesh.v2beta1.ComputedExplicitDestinations.destinations:type_name -> hashicorp.consul.mesh.v2beta1.Destination
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_computed_explicit_destinations_proto_init() }
func file_pbmesh_v2beta1_computed_explicit_destinations_proto_init() {
	if File_pbmesh_v2beta1_computed_explicit_destinations_proto != nil {
		return
	}
	file_pbmesh_v2beta1_destinations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v2beta1_computed_explicit_destinations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedExplicitDestinations); i {
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
			RawDescriptor: file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_computed_explicit_destinations_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_computed_explicit_destinations_proto_depIdxs,
		MessageInfos:      file_pbmesh_v2beta1_computed_explicit_destinations_proto_msgTypes,
	}.Build()
	File_pbmesh_v2beta1_computed_explicit_destinations_proto = out.File
	file_pbmesh_v2beta1_computed_explicit_destinations_proto_rawDesc = nil
	file_pbmesh_v2beta1_computed_explicit_destinations_proto_goTypes = nil
	file_pbmesh_v2beta1_computed_explicit_destinations_proto_depIdxs = nil
}
