// Copyright 2021-2024 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2-devel
// 	protoc        (unknown)
// source: cerbos/audit/v1/audit.proto

package auditv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	v11 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccessLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId     string                 `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
	Timestamp  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Peer       *Peer                  `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`
	Metadata   map[string]*MetaValues `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Method     string                 `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	StatusCode uint32                 `protobuf:"varint,6,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
}

func (x *AccessLogEntry) Reset() {
	*x = AccessLogEntry{}
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogEntry) ProtoMessage() {}

func (x *AccessLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogEntry.ProtoReflect.Descriptor instead.
func (*AccessLogEntry) Descriptor() ([]byte, []int) {
	return file_cerbos_audit_v1_audit_proto_rawDescGZIP(), []int{0}
}

func (x *AccessLogEntry) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *AccessLogEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AccessLogEntry) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *AccessLogEntry) GetMetadata() map[string]*MetaValues {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AccessLogEntry) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AccessLogEntry) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type DecisionLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId    string                 `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Peer      *Peer                  `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`
	// Deprecated. Use method.check_resources.inputs instead.
	//
	// Deprecated: Marked as deprecated in cerbos/audit/v1/audit.proto.
	Inputs []*v1.CheckInput `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// Deprecated. Use method.check_resources.outputs instead.
	//
	// Deprecated: Marked as deprecated in cerbos/audit/v1/audit.proto.
	Outputs []*v1.CheckOutput `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// Deprecated. Use method.check_resources.error instead.
	//
	// Deprecated: Marked as deprecated in cerbos/audit/v1/audit.proto.
	Error string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are assignable to Method:
	//
	//	*DecisionLogEntry_CheckResources_
	//	*DecisionLogEntry_PlanResources_
	Method     isDecisionLogEntry_Method `protobuf_oneof:"method"`
	Metadata   map[string]*MetaValues    `protobuf:"bytes,15,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AuditTrail *AuditTrail               `protobuf:"bytes,16,opt,name=audit_trail,json=auditTrail,proto3" json:"audit_trail,omitempty"`
}

func (x *DecisionLogEntry) Reset() {
	*x = DecisionLogEntry{}
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecisionLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionLogEntry) ProtoMessage() {}

func (x *DecisionLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionLogEntry.ProtoReflect.Descriptor instead.
func (*DecisionLogEntry) Descriptor() ([]byte, []int) {
	return file_cerbos_audit_v1_audit_proto_rawDescGZIP(), []int{1}
}

func (x *DecisionLogEntry) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *DecisionLogEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DecisionLogEntry) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

// Deprecated: Marked as deprecated in cerbos/audit/v1/audit.proto.
func (x *DecisionLogEntry) GetInputs() []*v1.CheckInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

// Deprecated: Marked as deprecated in cerbos/audit/v1/audit.proto.
func (x *DecisionLogEntry) GetOutputs() []*v1.CheckOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

// Deprecated: Marked as deprecated in cerbos/audit/v1/audit.proto.
func (x *DecisionLogEntry) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (m *DecisionLogEntry) GetMethod() isDecisionLogEntry_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (x *DecisionLogEntry) GetCheckResources() *DecisionLogEntry_CheckResources {
	if x, ok := x.GetMethod().(*DecisionLogEntry_CheckResources_); ok {
		return x.CheckResources
	}
	return nil
}

func (x *DecisionLogEntry) GetPlanResources() *DecisionLogEntry_PlanResources {
	if x, ok := x.GetMethod().(*DecisionLogEntry_PlanResources_); ok {
		return x.PlanResources
	}
	return nil
}

func (x *DecisionLogEntry) GetMetadata() map[string]*MetaValues {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DecisionLogEntry) GetAuditTrail() *AuditTrail {
	if x != nil {
		return x.AuditTrail
	}
	return nil
}

type isDecisionLogEntry_Method interface {
	isDecisionLogEntry_Method()
}

type DecisionLogEntry_CheckResources_ struct {
	CheckResources *DecisionLogEntry_CheckResources `protobuf:"bytes,7,opt,name=check_resources,json=checkResources,proto3,oneof"`
}

type DecisionLogEntry_PlanResources_ struct {
	PlanResources *DecisionLogEntry_PlanResources `protobuf:"bytes,8,opt,name=plan_resources,json=planResources,proto3,oneof"`
}

func (*DecisionLogEntry_CheckResources_) isDecisionLogEntry_Method() {}

func (*DecisionLogEntry_PlanResources_) isDecisionLogEntry_Method() {}

type MetaValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MetaValues) Reset() {
	*x = MetaValues{}
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetaValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaValues) ProtoMessage() {}

func (x *MetaValues) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaValues.ProtoReflect.Descriptor instead.
func (*MetaValues) Descriptor() ([]byte, []int) {
	return file_cerbos_audit_v1_audit_proto_rawDescGZIP(), []int{2}
}

func (x *MetaValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AuthInfo     string `protobuf:"bytes,2,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty"`
	UserAgent    string `protobuf:"bytes,3,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	ForwardedFor string `protobuf:"bytes,4,opt,name=forwarded_for,json=forwardedFor,proto3" json:"forwarded_for,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_cerbos_audit_v1_audit_proto_rawDescGZIP(), []int{3}
}

func (x *Peer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Peer) GetAuthInfo() string {
	if x != nil {
		return x.AuthInfo
	}
	return ""
}

func (x *Peer) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Peer) GetForwardedFor() string {
	if x != nil {
		return x.ForwardedFor
	}
	return ""
}

type AuditTrail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EffectivePolicies map[string]*v11.SourceAttributes `protobuf:"bytes,1,rep,name=effective_policies,json=effectivePolicies,proto3" json:"effective_policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AuditTrail) Reset() {
	*x = AuditTrail{}
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditTrail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditTrail) ProtoMessage() {}

func (x *AuditTrail) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditTrail.ProtoReflect.Descriptor instead.
func (*AuditTrail) Descriptor() ([]byte, []int) {
	return file_cerbos_audit_v1_audit_proto_rawDescGZIP(), []int{4}
}

func (x *AuditTrail) GetEffectivePolicies() map[string]*v11.SourceAttributes {
	if x != nil {
		return x.EffectivePolicies
	}
	return nil
}

type DecisionLogEntry_CheckResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs  []*v1.CheckInput  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*v1.CheckOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Error   string            `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DecisionLogEntry_CheckResources) Reset() {
	*x = DecisionLogEntry_CheckResources{}
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecisionLogEntry_CheckResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionLogEntry_CheckResources) ProtoMessage() {}

func (x *DecisionLogEntry_CheckResources) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionLogEntry_CheckResources.ProtoReflect.Descriptor instead.
func (*DecisionLogEntry_CheckResources) Descriptor() ([]byte, []int) {
	return file_cerbos_audit_v1_audit_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DecisionLogEntry_CheckResources) GetInputs() []*v1.CheckInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *DecisionLogEntry_CheckResources) GetOutputs() []*v1.CheckOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *DecisionLogEntry_CheckResources) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DecisionLogEntry_PlanResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  *v1.PlanResourcesInput  `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output *v1.PlanResourcesOutput `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Error  string                  `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DecisionLogEntry_PlanResources) Reset() {
	*x = DecisionLogEntry_PlanResources{}
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecisionLogEntry_PlanResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionLogEntry_PlanResources) ProtoMessage() {}

func (x *DecisionLogEntry_PlanResources) ProtoReflect() protoreflect.Message {
	mi := &file_cerbos_audit_v1_audit_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionLogEntry_PlanResources.ProtoReflect.Descriptor instead.
func (*DecisionLogEntry_PlanResources) Descriptor() ([]byte, []int) {
	return file_cerbos_audit_v1_audit_proto_rawDescGZIP(), []int{1, 1}
}

func (x *DecisionLogEntry_PlanResources) GetInput() *v1.PlanResourcesInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *DecisionLogEntry_PlanResources) GetOutput() *v1.PlanResourcesOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *DecisionLogEntry_PlanResources) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_cerbos_audit_v1_audit_proto protoreflect.FileDescriptor

var file_cerbos_audit_v1_audit_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1d,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02,
	0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x49,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x1a, 0x58, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x08, 0x0a,
	0x10, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x5b, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x58, 0x0a,
	0x0e, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61,
	0x69, 0x6c, 0x1a, 0x95, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xa0, 0x01, 0x0a, 0x0d, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x58, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x22, 0x24, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x22, 0xd9, 0x01, 0x0a, 0x0a,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x12, 0x61, 0x0a, 0x12, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x1a, 0x68, 0x0a,
	0x16, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x6b, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x76, 0x31, 0xaa, 0x02,
	0x13, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cerbos_audit_v1_audit_proto_rawDescOnce sync.Once
	file_cerbos_audit_v1_audit_proto_rawDescData = file_cerbos_audit_v1_audit_proto_rawDesc
)

func file_cerbos_audit_v1_audit_proto_rawDescGZIP() []byte {
	file_cerbos_audit_v1_audit_proto_rawDescOnce.Do(func() {
		file_cerbos_audit_v1_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_cerbos_audit_v1_audit_proto_rawDescData)
	})
	return file_cerbos_audit_v1_audit_proto_rawDescData
}

var file_cerbos_audit_v1_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cerbos_audit_v1_audit_proto_goTypes = []any{
	(*AccessLogEntry)(nil),                  // 0: cerbos.audit.v1.AccessLogEntry
	(*DecisionLogEntry)(nil),                // 1: cerbos.audit.v1.DecisionLogEntry
	(*MetaValues)(nil),                      // 2: cerbos.audit.v1.MetaValues
	(*Peer)(nil),                            // 3: cerbos.audit.v1.Peer
	(*AuditTrail)(nil),                      // 4: cerbos.audit.v1.AuditTrail
	nil,                                     // 5: cerbos.audit.v1.AccessLogEntry.MetadataEntry
	(*DecisionLogEntry_CheckResources)(nil), // 6: cerbos.audit.v1.DecisionLogEntry.CheckResources
	(*DecisionLogEntry_PlanResources)(nil),  // 7: cerbos.audit.v1.DecisionLogEntry.PlanResources
	nil,                                     // 8: cerbos.audit.v1.DecisionLogEntry.MetadataEntry
	nil,                                     // 9: cerbos.audit.v1.AuditTrail.EffectivePoliciesEntry
	(*timestamppb.Timestamp)(nil),           // 10: google.protobuf.Timestamp
	(*v1.CheckInput)(nil),                   // 11: cerbos.engine.v1.CheckInput
	(*v1.CheckOutput)(nil),                  // 12: cerbos.engine.v1.CheckOutput
	(*v1.PlanResourcesInput)(nil),           // 13: cerbos.engine.v1.PlanResourcesInput
	(*v1.PlanResourcesOutput)(nil),          // 14: cerbos.engine.v1.PlanResourcesOutput
	(*v11.SourceAttributes)(nil),            // 15: cerbos.policy.v1.SourceAttributes
}
var file_cerbos_audit_v1_audit_proto_depIdxs = []int32{
	10, // 0: cerbos.audit.v1.AccessLogEntry.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 1: cerbos.audit.v1.AccessLogEntry.peer:type_name -> cerbos.audit.v1.Peer
	5,  // 2: cerbos.audit.v1.AccessLogEntry.metadata:type_name -> cerbos.audit.v1.AccessLogEntry.MetadataEntry
	10, // 3: cerbos.audit.v1.DecisionLogEntry.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 4: cerbos.audit.v1.DecisionLogEntry.peer:type_name -> cerbos.audit.v1.Peer
	11, // 5: cerbos.audit.v1.DecisionLogEntry.inputs:type_name -> cerbos.engine.v1.CheckInput
	12, // 6: cerbos.audit.v1.DecisionLogEntry.outputs:type_name -> cerbos.engine.v1.CheckOutput
	6,  // 7: cerbos.audit.v1.DecisionLogEntry.check_resources:type_name -> cerbos.audit.v1.DecisionLogEntry.CheckResources
	7,  // 8: cerbos.audit.v1.DecisionLogEntry.plan_resources:type_name -> cerbos.audit.v1.DecisionLogEntry.PlanResources
	8,  // 9: cerbos.audit.v1.DecisionLogEntry.metadata:type_name -> cerbos.audit.v1.DecisionLogEntry.MetadataEntry
	4,  // 10: cerbos.audit.v1.DecisionLogEntry.audit_trail:type_name -> cerbos.audit.v1.AuditTrail
	9,  // 11: cerbos.audit.v1.AuditTrail.effective_policies:type_name -> cerbos.audit.v1.AuditTrail.EffectivePoliciesEntry
	2,  // 12: cerbos.audit.v1.AccessLogEntry.MetadataEntry.value:type_name -> cerbos.audit.v1.MetaValues
	11, // 13: cerbos.audit.v1.DecisionLogEntry.CheckResources.inputs:type_name -> cerbos.engine.v1.CheckInput
	12, // 14: cerbos.audit.v1.DecisionLogEntry.CheckResources.outputs:type_name -> cerbos.engine.v1.CheckOutput
	13, // 15: cerbos.audit.v1.DecisionLogEntry.PlanResources.input:type_name -> cerbos.engine.v1.PlanResourcesInput
	14, // 16: cerbos.audit.v1.DecisionLogEntry.PlanResources.output:type_name -> cerbos.engine.v1.PlanResourcesOutput
	2,  // 17: cerbos.audit.v1.DecisionLogEntry.MetadataEntry.value:type_name -> cerbos.audit.v1.MetaValues
	15, // 18: cerbos.audit.v1.AuditTrail.EffectivePoliciesEntry.value:type_name -> cerbos.policy.v1.SourceAttributes
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_cerbos_audit_v1_audit_proto_init() }
func file_cerbos_audit_v1_audit_proto_init() {
	if File_cerbos_audit_v1_audit_proto != nil {
		return
	}
	file_cerbos_audit_v1_audit_proto_msgTypes[1].OneofWrappers = []any{
		(*DecisionLogEntry_CheckResources_)(nil),
		(*DecisionLogEntry_PlanResources_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cerbos_audit_v1_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cerbos_audit_v1_audit_proto_goTypes,
		DependencyIndexes: file_cerbos_audit_v1_audit_proto_depIdxs,
		MessageInfos:      file_cerbos_audit_v1_audit_proto_msgTypes,
	}.Build()
	File_cerbos_audit_v1_audit_proto = out.File
	file_cerbos_audit_v1_audit_proto_rawDesc = nil
	file_cerbos_audit_v1_audit_proto_goTypes = nil
	file_cerbos_audit_v1_audit_proto_depIdxs = nil
}
