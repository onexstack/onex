// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.23.4
// source: toyblc/v1/toyblc.proto

package v1

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

type CreateBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateBlockRequest) Reset() {
	*x = CreateBlockRequest{}
	mi := &file_toyblc_v1_toyblc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlockRequest) ProtoMessage() {}

func (x *CreateBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_toyblc_v1_toyblc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlockRequest.ProtoReflect.Descriptor instead.
func (*CreateBlockRequest) Descriptor() ([]byte, []int) {
	return file_toyblc_v1_toyblc_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlockRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CreatePeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer string `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *CreatePeerRequest) Reset() {
	*x = CreatePeerRequest{}
	mi := &file_toyblc_v1_toyblc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePeerRequest) ProtoMessage() {}

func (x *CreatePeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_toyblc_v1_toyblc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePeerRequest.ProtoReflect.Descriptor instead.
func (*CreatePeerRequest) Descriptor() ([]byte, []int) {
	return file_toyblc_v1_toyblc_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePeerRequest) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

var File_toyblc_v1_toyblc_proto protoreflect.FileDescriptor

var file_toyblc_v1_toyblc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x6f, 0x79, 0x62, 0x6c, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x79, 0x62,
	0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x6a, 0x2f, 0x6f, 0x6e, 0x65, 0x78, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x6f, 0x79, 0x62, 0x6c, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_toyblc_v1_toyblc_proto_rawDescOnce sync.Once
	file_toyblc_v1_toyblc_proto_rawDescData = file_toyblc_v1_toyblc_proto_rawDesc
)

func file_toyblc_v1_toyblc_proto_rawDescGZIP() []byte {
	file_toyblc_v1_toyblc_proto_rawDescOnce.Do(func() {
		file_toyblc_v1_toyblc_proto_rawDescData = protoimpl.X.CompressGZIP(file_toyblc_v1_toyblc_proto_rawDescData)
	})
	return file_toyblc_v1_toyblc_proto_rawDescData
}

var file_toyblc_v1_toyblc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_toyblc_v1_toyblc_proto_goTypes = []any{
	(*CreateBlockRequest)(nil), // 0: usercenter.v1.CreateBlockRequest
	(*CreatePeerRequest)(nil),  // 1: usercenter.v1.CreatePeerRequest
}
var file_toyblc_v1_toyblc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_toyblc_v1_toyblc_proto_init() }
func file_toyblc_v1_toyblc_proto_init() {
	if File_toyblc_v1_toyblc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_toyblc_v1_toyblc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_toyblc_v1_toyblc_proto_goTypes,
		DependencyIndexes: file_toyblc_v1_toyblc_proto_depIdxs,
		MessageInfos:      file_toyblc_v1_toyblc_proto_msgTypes,
	}.Build()
	File_toyblc_v1_toyblc_proto = out.File
	file_toyblc_v1_toyblc_proto_rawDesc = nil
	file_toyblc_v1_toyblc_proto_goTypes = nil
	file_toyblc_v1_toyblc_proto_depIdxs = nil
}
