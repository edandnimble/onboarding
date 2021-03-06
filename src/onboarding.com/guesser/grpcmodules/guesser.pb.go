// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: onboarding.com/guesser/grpcmodules/guesser.proto

package grpcmodules

import (
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

type Guesser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginAt       uint32 `protobuf:"varint,1,opt,name=beginAt,proto3" json:"beginAt,omitempty"`
	IncrementBy   uint32 `protobuf:"varint,2,opt,name=incrementBy,proto3" json:"incrementBy,omitempty"`
	SleepInterval uint32 `protobuf:"varint,3,opt,name=sleepInterval,proto3" json:"sleepInterval,omitempty"`
}

func (x *Guesser) Reset() {
	*x = Guesser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guesser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guesser) ProtoMessage() {}

func (x *Guesser) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guesser.ProtoReflect.Descriptor instead.
func (*Guesser) Descriptor() ([]byte, []int) {
	return file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescGZIP(), []int{0}
}

func (x *Guesser) GetBeginAt() uint32 {
	if x != nil {
		return x.BeginAt
	}
	return 0
}

func (x *Guesser) GetIncrementBy() uint32 {
	if x != nil {
		return x.IncrementBy
	}
	return 0
}

func (x *Guesser) GetSleepInterval() uint32 {
	if x != nil {
		return x.SleepInterval
	}
	return 0
}

type GuesserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GuesserId) Reset() {
	*x = GuesserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuesserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuesserId) ProtoMessage() {}

func (x *GuesserId) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuesserId.ProtoReflect.Descriptor instead.
func (*GuesserId) Descriptor() ([]byte, []int) {
	return file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescGZIP(), []int{1}
}

func (x *GuesserId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResponseStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrCode uint32 `protobuf:"varint,2,opt,name=errCode,proto3" json:"errCode,omitempty"`
}

func (x *ResponseStatus) Reset() {
	*x = ResponseStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatus) ProtoMessage() {}

func (x *ResponseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatus.ProtoReflect.Descriptor instead.
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseStatus) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ResponseStatus) GetErrCode() uint32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

type GuessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num     uint32                 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Attempt int64                  `protobuf:"varint,2,opt,name=attempt,proto3" json:"attempt,omitempty"`
	FoundAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=foundAt,proto3" json:"foundAt,omitempty"`
}

func (x *GuessInfo) Reset() {
	*x = GuessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessInfo) ProtoMessage() {}

func (x *GuessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessInfo.ProtoReflect.Descriptor instead.
func (*GuessInfo) Descriptor() ([]byte, []int) {
	return file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescGZIP(), []int{3}
}

func (x *GuessInfo) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *GuessInfo) GetAttempt() int64 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

func (x *GuessInfo) GetFoundAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FoundAt
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        *ResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Active        bool            `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	BeginAt       uint32          `protobuf:"varint,3,opt,name=beginAt,proto3" json:"beginAt,omitempty"`
	IncrementBy   uint32          `protobuf:"varint,4,opt,name=incrementBy,proto3" json:"incrementBy,omitempty"`
	SleepInterval uint32          `protobuf:"varint,5,opt,name=sleepInterval,proto3" json:"sleepInterval,omitempty"`
	Guesses       []*GuessInfo    `protobuf:"bytes,6,rep,name=guesses,proto3" json:"guesses,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescGZIP(), []int{4}
}

func (x *QueryResponse) GetStatus() *ResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *QueryResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *QueryResponse) GetBeginAt() uint32 {
	if x != nil {
		return x.BeginAt
	}
	return 0
}

func (x *QueryResponse) GetIncrementBy() uint32 {
	if x != nil {
		return x.IncrementBy
	}
	return 0
}

func (x *QueryResponse) GetSleepInterval() uint32 {
	if x != nil {
		return x.SleepInterval
	}
	return 0
}

func (x *QueryResponse) GetGuesses() []*GuessInfo {
	if x != nil {
		return x.Guesses
	}
	return nil
}

type AddGuesserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *ResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id     uint32          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddGuesserResponse) Reset() {
	*x = AddGuesserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGuesserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGuesserResponse) ProtoMessage() {}

func (x *AddGuesserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGuesserResponse.ProtoReflect.Descriptor instead.
func (*AddGuesserResponse) Descriptor() ([]byte, []int) {
	return file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescGZIP(), []int{5}
}

func (x *AddGuesserResponse) GetStatus() *ResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AddGuesserResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_onboarding_com_guesser_grpcmodules_guesser_proto protoreflect.FileDescriptor

var file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6b, 0x0a, 0x07, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6c, 0x65, 0x65,
	0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x1b,
	0x0a, 0x09, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x6d, 0x0a, 0x09, 0x47, 0x75, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x12, 0x34, 0x0a, 0x07, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x41, 0x74, 0x22, 0xf2, 0x01, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x75, 0x65, 0x73, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x41, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x6c, 0x65, 0x65, 0x70,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x75, 0x65, 0x73,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd2, 0x01, 0x0a, 0x0a, 0x47, 0x75, 0x65, 0x73,
	0x73, 0x65, 0x72, 0x52, 0x70, 0x63, 0x12, 0x40, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e,
	0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x75, 0x65,
	0x73, 0x73, 0x65, 0x72, 0x1a, 0x20, 0x2e, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x1c, 0x2e, 0x67, 0x75,
	0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x05, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x1b, 0x2e,
	0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x75, 0x65, 0x73, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescOnce sync.Once
	file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescData = file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDesc
)

func file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescGZIP() []byte {
	file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescOnce.Do(func() {
		file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescData = protoimpl.X.CompressGZIP(file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescData)
	})
	return file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDescData
}

var file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_onboarding_com_guesser_grpcmodules_guesser_proto_goTypes = []interface{}{
	(*Guesser)(nil),               // 0: guesserModel.Guesser
	(*GuesserId)(nil),             // 1: guesserModel.GuesserId
	(*ResponseStatus)(nil),        // 2: guesserModel.ResponseStatus
	(*GuessInfo)(nil),             // 3: guesserModel.GuessInfo
	(*QueryResponse)(nil),         // 4: guesserModel.QueryResponse
	(*AddGuesserResponse)(nil),    // 5: guesserModel.AddGuesserResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_onboarding_com_guesser_grpcmodules_guesser_proto_depIdxs = []int32{
	6, // 0: guesserModel.GuessInfo.foundAt:type_name -> google.protobuf.Timestamp
	2, // 1: guesserModel.QueryResponse.status:type_name -> guesserModel.ResponseStatus
	3, // 2: guesserModel.QueryResponse.guesses:type_name -> guesserModel.GuessInfo
	2, // 3: guesserModel.AddGuesserResponse.status:type_name -> guesserModel.ResponseStatus
	0, // 4: guesserModel.GuesserRpc.Add:input_type -> guesserModel.Guesser
	1, // 5: guesserModel.GuesserRpc.Remove:input_type -> guesserModel.GuesserId
	1, // 6: guesserModel.GuesserRpc.Query:input_type -> guesserModel.GuesserId
	5, // 7: guesserModel.GuesserRpc.Add:output_type -> guesserModel.AddGuesserResponse
	2, // 8: guesserModel.GuesserRpc.Remove:output_type -> guesserModel.ResponseStatus
	4, // 9: guesserModel.GuesserRpc.Query:output_type -> guesserModel.QueryResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_onboarding_com_guesser_grpcmodules_guesser_proto_init() }
func file_onboarding_com_guesser_grpcmodules_guesser_proto_init() {
	if File_onboarding_com_guesser_grpcmodules_guesser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guesser); i {
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
		file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuesserId); i {
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
		file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStatus); i {
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
		file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessInfo); i {
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
		file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGuesserResponse); i {
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
			RawDescriptor: file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_onboarding_com_guesser_grpcmodules_guesser_proto_goTypes,
		DependencyIndexes: file_onboarding_com_guesser_grpcmodules_guesser_proto_depIdxs,
		MessageInfos:      file_onboarding_com_guesser_grpcmodules_guesser_proto_msgTypes,
	}.Build()
	File_onboarding_com_guesser_grpcmodules_guesser_proto = out.File
	file_onboarding_com_guesser_grpcmodules_guesser_proto_rawDesc = nil
	file_onboarding_com_guesser_grpcmodules_guesser_proto_goTypes = nil
	file_onboarding_com_guesser_grpcmodules_guesser_proto_depIdxs = nil
}
