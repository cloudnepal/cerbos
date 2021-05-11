// Copyright 2021 Zenauth Ltd.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: request/v1/request.proto

package requestv1

import (
	v1 "github.com/cerbos/cerbos/internal/genpb/engine/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CheckResourceSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   string        `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Actions     []string      `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	Principal   *v1.Principal `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`
	Resource    *ResourceSet  `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	IncludeMeta bool          `protobuf:"varint,5,opt,name=include_meta,json=includeMeta,proto3" json:"include_meta,omitempty"`
}

func (x *CheckResourceSetRequest) Reset() {
	*x = CheckResourceSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResourceSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResourceSetRequest) ProtoMessage() {}

func (x *CheckResourceSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResourceSetRequest.ProtoReflect.Descriptor instead.
func (*CheckResourceSetRequest) Descriptor() ([]byte, []int) {
	return file_request_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *CheckResourceSetRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CheckResourceSetRequest) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *CheckResourceSetRequest) GetPrincipal() *v1.Principal {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *CheckResourceSetRequest) GetResource() *ResourceSet {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *CheckResourceSetRequest) GetIncludeMeta() bool {
	if x != nil {
		return x.IncludeMeta
	}
	return false
}

type ResourceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind          string                    `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	PolicyVersion string                    `protobuf:"bytes,2,opt,name=policy_version,json=policyVersion,proto3" json:"policy_version,omitempty"`
	Instances     map[string]*AttributesMap `protobuf:"bytes,3,rep,name=instances,proto3" json:"instances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResourceSet) Reset() {
	*x = ResourceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_v1_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSet) ProtoMessage() {}

func (x *ResourceSet) ProtoReflect() protoreflect.Message {
	mi := &file_request_v1_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSet.ProtoReflect.Descriptor instead.
func (*ResourceSet) Descriptor() ([]byte, []int) {
	return file_request_v1_request_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceSet) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ResourceSet) GetPolicyVersion() string {
	if x != nil {
		return x.PolicyVersion
	}
	return ""
}

func (x *ResourceSet) GetInstances() map[string]*AttributesMap {
	if x != nil {
		return x.Instances
	}
	return nil
}

type AttributesMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attr map[string]*structpb.Value `protobuf:"bytes,1,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AttributesMap) Reset() {
	*x = AttributesMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_v1_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributesMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributesMap) ProtoMessage() {}

func (x *AttributesMap) ProtoReflect() protoreflect.Message {
	mi := &file_request_v1_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributesMap.ProtoReflect.Descriptor instead.
func (*AttributesMap) Descriptor() ([]byte, []int) {
	return file_request_v1_request_proto_rawDescGZIP(), []int{2}
}

func (x *AttributesMap) GetAttr() map[string]*structpb.Value {
	if x != nil {
		return x.Attr
	}
	return nil
}

type CheckResourceBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string                                  `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Principal *v1.Principal                           `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	Resources []*CheckResourceBatchRequest_BatchEntry `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *CheckResourceBatchRequest) Reset() {
	*x = CheckResourceBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_v1_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResourceBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResourceBatchRequest) ProtoMessage() {}

func (x *CheckResourceBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_v1_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResourceBatchRequest.ProtoReflect.Descriptor instead.
func (*CheckResourceBatchRequest) Descriptor() ([]byte, []int) {
	return file_request_v1_request_proto_rawDescGZIP(), []int{3}
}

func (x *CheckResourceBatchRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CheckResourceBatchRequest) GetPrincipal() *v1.Principal {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *CheckResourceBatchRequest) GetResources() []*CheckResourceBatchRequest_BatchEntry {
	if x != nil {
		return x.Resources
	}
	return nil
}

type CheckResourceBatchRequest_BatchEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions  []string     `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	Resource *v1.Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CheckResourceBatchRequest_BatchEntry) Reset() {
	*x = CheckResourceBatchRequest_BatchEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_v1_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResourceBatchRequest_BatchEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResourceBatchRequest_BatchEntry) ProtoMessage() {}

func (x *CheckResourceBatchRequest_BatchEntry) ProtoReflect() protoreflect.Message {
	mi := &file_request_v1_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResourceBatchRequest_BatchEntry.ProtoReflect.Descriptor instead.
func (*CheckResourceBatchRequest_BatchEntry) Descriptor() ([]byte, []int) {
	return file_request_v1_request_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CheckResourceBatchRequest_BatchEntry) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *CheckResourceBatchRequest_BatchEntry) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_request_v1_request_proto protoreflect.FileDescriptor

var file_request_v1_request_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x96, 0x01, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x77, 0x92, 0x41, 0x74, 0x32, 0x4a, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x49, 0x44, 0x20, 0x75,
	0x73, 0x65, 0x66, 0x75, 0x6c, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x4a, 0x26, 0x22, 0x63, 0x32, 0x64, 0x62, 0x31,
	0x37, 0x62, 0x38, 0x2d, 0x34, 0x66, 0x39, 0x66, 0x2d, 0x34, 0x66, 0x62, 0x31, 0x2d, 0x61, 0x63,
	0x66, 0x64, 0x2d, 0x39, 0x31, 0x36, 0x32, 0x61, 0x30, 0x32, 0x62, 0x65, 0x34, 0x32, 0x62, 0x22,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x92, 0x01, 0x0a, 0x07,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x78, 0x92,
	0x41, 0x5f, 0x32, 0x38, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x62, 0x65, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4a, 0x1a, 0x5b, 0x22,
	0x76, 0x69, 0x65, 0x77, 0x3a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x2c, 0x20, 0x22, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5d, 0xa0, 0x01, 0x0a, 0xa8, 0x01, 0x01, 0xb0, 0x01,
	0x01, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x0f, 0x92, 0x01, 0x0c, 0x08, 0x01, 0x10, 0x0a, 0x18,
	0x01, 0x22, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x40, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42, 0x0c, 0xe2, 0x41, 0x01, 0x02, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x12, 0x41, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x42, 0x0c, 0xe2,
	0x41, 0x01, 0x02, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x40, 0x92, 0x41, 0x3d,
	0x32, 0x3b, 0x4f, 0x70, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x20, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x69, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x0b, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x3a, 0x12, 0x92, 0x41, 0x0f, 0x0a,
	0x0d, 0x32, 0x0b, 0x50, 0x44, 0x50, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9e,
	0x07, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0xc6,
	0x01, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb1, 0x01,
	0x92, 0x41, 0x62, 0x32, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6b, 0x69,
	0x6e, 0x64, 0x2e, 0x4a, 0x0e, 0x22, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x3a, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x8a, 0x01, 0x3f, 0x5e, 0x5b, 0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a,
	0x5d, 0x5d, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c,
	0x2d, 0x5d, 0x2a, 0x28, 0x5c, 0x3a, 0x5b, 0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d,
	0x5d, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c, 0x2d,
	0x5d, 0x2a, 0x29, 0x2a, 0x24, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x10, 0x01,
	0x32, 0x3f, 0x5e, 0x5b, 0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x5b, 0x5b,
	0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c, 0x2d, 0x5d, 0x2a, 0x28,
	0x5c, 0x3a, 0x5b, 0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x5b, 0x5b, 0x3a,
	0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5c, 0x40, 0x5c, 0x2e, 0x5c, 0x2d, 0x5d, 0x2a, 0x29, 0x2a,
	0x24, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0xdd, 0x01, 0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0xb5, 0x01, 0x92, 0x41, 0x99, 0x01, 0x32, 0x7c, 0x54, 0x68, 0x65, 0x20, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x75,
	0x73, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x20, 0x49, 0x66, 0x20,
	0x6e, 0x6f, 0x74, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2c, 0x20, 0x77,
	0x69, 0x6c, 0x6c, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x64, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x09, 0x22, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0x8a, 0x01, 0x0d, 0x5e, 0x5b, 0x5b, 0x3a, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5d, 0x2a, 0x24,
	0xe2, 0x41, 0x01, 0x01, 0xfa, 0x42, 0x11, 0x72, 0x0f, 0x32, 0x0d, 0x5e, 0x5b, 0x5b, 0x3a, 0x77,
	0x6f, 0x72, 0x64, 0x3a, 0x5d, 0x5d, 0x2a, 0x24, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0xea, 0x02, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0xa3, 0x02, 0x92, 0x41, 0x8f, 0x02, 0x32, 0x6d, 0x53, 0x65, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x20,
	0x45, 0x61, 0x63, 0x68, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x6d, 0x75,
	0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x61,
	0x6e, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4a, 0x97, 0x01, 0x7b, 0x22, 0x58, 0x58,
	0x31, 0x32, 0x35, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x74, 0x74, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x22, 0x3a, 0x22, 0x62, 0x75, 0x67, 0x73, 0x5f, 0x62, 0x75, 0x6e, 0x6e,
	0x79, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x3a, 0x20, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x2c, 0x20, 0x22, 0x66, 0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x22, 0x3a, 0x20,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x7d, 0x7d, 0x2c, 0x20, 0x22, 0x58, 0x58, 0x32, 0x32, 0x35, 0x22,
	0x3a, 0x7b, 0x22, 0x61, 0x74, 0x74, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x22, 0x3a, 0x22, 0x64, 0x61, 0x66, 0x66, 0x79, 0x5f, 0x64, 0x75, 0x63, 0x6b, 0x22, 0x2c, 0x20,
	0x22, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x3a, 0x20, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x20,
	0x22, 0x66, 0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0x7d, 0x7d, 0x7d, 0xc0, 0x01, 0x14, 0xc8, 0x01, 0x01, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x09,
	0x9a, 0x01, 0x06, 0x08, 0x01, 0x10, 0x14, 0x18, 0x01, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x1a, 0x57, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x4d,
	0x61, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x20, 0x92,
	0x41, 0x1d, 0x0a, 0x1b, 0x32, 0x19, 0x53, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22,
	0xc2, 0x02, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x4d, 0x61,
	0x70, 0x12, 0xaa, 0x01, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x71, 0x92, 0x41, 0x66, 0x32, 0x64, 0x4b, 0x65, 0x79, 0x2d,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x70, 0x61, 0x69, 0x72, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x61,
	0x62, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x64, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x20, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0xfa, 0x42, 0x05, 0x9a, 0x01, 0x02, 0x18, 0x01, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x1a, 0x4f,
	0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a,
	0x33, 0x92, 0x41, 0x30, 0x0a, 0x2e, 0x32, 0x2c, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x22, 0xa3, 0x06, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x96, 0x01, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x77, 0x92, 0x41, 0x74, 0x32, 0x4a, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x49, 0x44, 0x20, 0x75,
	0x73, 0x65, 0x66, 0x75, 0x6c, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x4a, 0x26, 0x22, 0x63, 0x32, 0x64, 0x62, 0x31,
	0x37, 0x62, 0x38, 0x2d, 0x34, 0x66, 0x39, 0x66, 0x2d, 0x34, 0x66, 0x62, 0x31, 0x2d, 0x61, 0x63,
	0x66, 0x64, 0x2d, 0x39, 0x31, 0x36, 0x32, 0x61, 0x30, 0x32, 0x62, 0x65, 0x34, 0x32, 0x62, 0x22,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x09, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x42, 0x0c, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0xbb, 0x02,
	0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0xea, 0x01, 0x92, 0x41, 0xd8, 0x01, 0x32, 0x1e, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4a, 0xac, 0x01, 0x5b, 0x7b, 0x22,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x76, 0x69, 0x65, 0x77, 0x22,
	0x2c, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5d, 0x2c, 0x20, 0x22, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x3a,
	0x22, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x3a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x58,
	0x58, 0x31, 0x32, 0x35, 0x22, 0x2c, 0x22, 0x61, 0x74, 0x74, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x22, 0x3a, 0x22, 0x62, 0x75, 0x67, 0x73, 0x5f, 0x62, 0x75, 0x6e, 0x6e,
	0x79, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x3a, 0x20, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x2c, 0x20, 0x22, 0x66, 0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x22, 0x3a, 0x20,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x7d, 0x7d, 0x7d, 0x5d, 0xa0, 0x01, 0x14, 0xa8, 0x01, 0x01, 0xb0,
	0x01, 0x01, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x07, 0x92, 0x01, 0x04, 0x08, 0x01, 0x10, 0x14,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0xd8, 0x01, 0x0a, 0x0a,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x8a, 0x01, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x70, 0x92, 0x41,
	0x57, 0x32, 0x30, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x20, 0x62, 0x65, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x4a, 0x1a, 0x5b, 0x22, 0x76, 0x69, 0x65, 0x77, 0x3a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5d, 0xa0,
	0x01, 0x0a, 0xa8, 0x01, 0x01, 0xb0, 0x01, 0x01, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x0f, 0x92,
	0x01, 0x0c, 0x08, 0x01, 0x10, 0x0a, 0x18, 0x01, 0x22, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0c,
	0xe2, 0x41, 0x01, 0x02, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x12, 0x92, 0x41, 0x0f, 0x0a, 0x0d, 0x32, 0x0b, 0x50,
	0x44, 0x50, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_request_v1_request_proto_rawDescOnce sync.Once
	file_request_v1_request_proto_rawDescData = file_request_v1_request_proto_rawDesc
)

func file_request_v1_request_proto_rawDescGZIP() []byte {
	file_request_v1_request_proto_rawDescOnce.Do(func() {
		file_request_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_v1_request_proto_rawDescData)
	})
	return file_request_v1_request_proto_rawDescData
}

var file_request_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_request_v1_request_proto_goTypes = []interface{}{
	(*CheckResourceSetRequest)(nil),   // 0: request.v1.CheckResourceSetRequest
	(*ResourceSet)(nil),               // 1: request.v1.ResourceSet
	(*AttributesMap)(nil),             // 2: request.v1.AttributesMap
	(*CheckResourceBatchRequest)(nil), // 3: request.v1.CheckResourceBatchRequest
	nil,                               // 4: request.v1.ResourceSet.InstancesEntry
	nil,                               // 5: request.v1.AttributesMap.AttrEntry
	(*CheckResourceBatchRequest_BatchEntry)(nil), // 6: request.v1.CheckResourceBatchRequest.BatchEntry
	(*v1.Principal)(nil),                         // 7: engine.v1.Principal
	(*structpb.Value)(nil),                       // 8: google.protobuf.Value
	(*v1.Resource)(nil),                          // 9: engine.v1.Resource
}
var file_request_v1_request_proto_depIdxs = []int32{
	7, // 0: request.v1.CheckResourceSetRequest.principal:type_name -> engine.v1.Principal
	1, // 1: request.v1.CheckResourceSetRequest.resource:type_name -> request.v1.ResourceSet
	4, // 2: request.v1.ResourceSet.instances:type_name -> request.v1.ResourceSet.InstancesEntry
	5, // 3: request.v1.AttributesMap.attr:type_name -> request.v1.AttributesMap.AttrEntry
	7, // 4: request.v1.CheckResourceBatchRequest.principal:type_name -> engine.v1.Principal
	6, // 5: request.v1.CheckResourceBatchRequest.resources:type_name -> request.v1.CheckResourceBatchRequest.BatchEntry
	2, // 6: request.v1.ResourceSet.InstancesEntry.value:type_name -> request.v1.AttributesMap
	8, // 7: request.v1.AttributesMap.AttrEntry.value:type_name -> google.protobuf.Value
	9, // 8: request.v1.CheckResourceBatchRequest.BatchEntry.resource:type_name -> engine.v1.Resource
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_request_v1_request_proto_init() }
func file_request_v1_request_proto_init() {
	if File_request_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResourceSetRequest); i {
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
		file_request_v1_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceSet); i {
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
		file_request_v1_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributesMap); i {
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
		file_request_v1_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResourceBatchRequest); i {
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
		file_request_v1_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResourceBatchRequest_BatchEntry); i {
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
			RawDescriptor: file_request_v1_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_v1_request_proto_goTypes,
		DependencyIndexes: file_request_v1_request_proto_depIdxs,
		MessageInfos:      file_request_v1_request_proto_msgTypes,
	}.Build()
	File_request_v1_request_proto = out.File
	file_request_v1_request_proto_rawDesc = nil
	file_request_v1_request_proto_goTypes = nil
	file_request_v1_request_proto_depIdxs = nil
}
