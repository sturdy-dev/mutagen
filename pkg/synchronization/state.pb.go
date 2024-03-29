// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: synchronization/state.proto

package synchronization

import (
	proto "github.com/golang/protobuf/proto"
	core "github.com/mutagen-io/mutagen/pkg/synchronization/core"
	rsync "github.com/mutagen-io/mutagen/pkg/synchronization/rsync"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	Status_Disconnected           Status = 0
	Status_HaltedOnRootEmptied    Status = 1
	Status_HaltedOnRootDeletion   Status = 2
	Status_HaltedOnRootTypeChange Status = 3
	Status_ConnectingAlpha        Status = 4
	Status_ConnectingBeta         Status = 5
	Status_Watching               Status = 6
	Status_Scanning               Status = 7
	Status_WaitingForRescan       Status = 8
	Status_Reconciling            Status = 9
	Status_StagingAlpha           Status = 10
	Status_StagingBeta            Status = 11
	Status_Transitioning          Status = 12
	Status_Saving                 Status = 13
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "Disconnected",
		1:  "HaltedOnRootEmptied",
		2:  "HaltedOnRootDeletion",
		3:  "HaltedOnRootTypeChange",
		4:  "ConnectingAlpha",
		5:  "ConnectingBeta",
		6:  "Watching",
		7:  "Scanning",
		8:  "WaitingForRescan",
		9:  "Reconciling",
		10: "StagingAlpha",
		11: "StagingBeta",
		12: "Transitioning",
		13: "Saving",
	}
	Status_value = map[string]int32{
		"Disconnected":           0,
		"HaltedOnRootEmptied":    1,
		"HaltedOnRootDeletion":   2,
		"HaltedOnRootTypeChange": 3,
		"ConnectingAlpha":        4,
		"ConnectingBeta":         5,
		"Watching":               6,
		"Scanning":               7,
		"WaitingForRescan":       8,
		"Reconciling":            9,
		"StagingAlpha":           10,
		"StagingBeta":            11,
		"Transitioning":          12,
		"Saving":                 13,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_state_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_synchronization_state_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_state_proto_rawDescGZIP(), []int{0}
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session                         *Session              `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Status                          Status                `protobuf:"varint,2,opt,name=status,proto3,enum=synchronization.Status" json:"status,omitempty"`
	AlphaConnected                  bool                  `protobuf:"varint,3,opt,name=alphaConnected,proto3" json:"alphaConnected,omitempty"`
	BetaConnected                   bool                  `protobuf:"varint,4,opt,name=betaConnected,proto3" json:"betaConnected,omitempty"`
	LastError                       string                `protobuf:"bytes,5,opt,name=lastError,proto3" json:"lastError,omitempty"`
	SuccessfulSynchronizationCycles uint64                `protobuf:"varint,6,opt,name=successfulSynchronizationCycles,proto3" json:"successfulSynchronizationCycles,omitempty"`
	StagingStatus                   *rsync.ReceiverStatus `protobuf:"bytes,7,opt,name=stagingStatus,proto3" json:"stagingStatus,omitempty"`
	Conflicts                       []*core.Conflict      `protobuf:"bytes,8,rep,name=conflicts,proto3" json:"conflicts,omitempty"`
	AlphaProblems                   []*core.Problem       `protobuf:"bytes,9,rep,name=alphaProblems,proto3" json:"alphaProblems,omitempty"`
	BetaProblems                    []*core.Problem       `protobuf:"bytes,10,rep,name=betaProblems,proto3" json:"betaProblems,omitempty"`
	TruncatedConflicts              uint64                `protobuf:"varint,11,opt,name=truncatedConflicts,proto3" json:"truncatedConflicts,omitempty"`
	TruncatedAlphaProblems          uint64                `protobuf:"varint,12,opt,name=truncatedAlphaProblems,proto3" json:"truncatedAlphaProblems,omitempty"`
	TruncatedBetaProblems           uint64                `protobuf:"varint,13,opt,name=truncatedBetaProblems,proto3" json:"truncatedBetaProblems,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_synchronization_state_proto_rawDescGZIP(), []int{0}
}

func (x *State) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *State) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Disconnected
}

func (x *State) GetAlphaConnected() bool {
	if x != nil {
		return x.AlphaConnected
	}
	return false
}

func (x *State) GetBetaConnected() bool {
	if x != nil {
		return x.BetaConnected
	}
	return false
}

func (x *State) GetLastError() string {
	if x != nil {
		return x.LastError
	}
	return ""
}

func (x *State) GetSuccessfulSynchronizationCycles() uint64 {
	if x != nil {
		return x.SuccessfulSynchronizationCycles
	}
	return 0
}

func (x *State) GetStagingStatus() *rsync.ReceiverStatus {
	if x != nil {
		return x.StagingStatus
	}
	return nil
}

func (x *State) GetConflicts() []*core.Conflict {
	if x != nil {
		return x.Conflicts
	}
	return nil
}

func (x *State) GetAlphaProblems() []*core.Problem {
	if x != nil {
		return x.AlphaProblems
	}
	return nil
}

func (x *State) GetBetaProblems() []*core.Problem {
	if x != nil {
		return x.BetaProblems
	}
	return nil
}

func (x *State) GetTruncatedConflicts() uint64 {
	if x != nil {
		return x.TruncatedConflicts
	}
	return 0
}

func (x *State) GetTruncatedAlphaProblems() uint64 {
	if x != nil {
		return x.TruncatedAlphaProblems
	}
	return 0
}

func (x *State) GetTruncatedBetaProblems() uint64 {
	if x != nil {
		return x.TruncatedBetaProblems
	}
	return 0
}

var File_synchronization_state_proto protoreflect.FileDescriptor

var file_synchronization_state_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73,
	0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x23,
	0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x05, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x62, 0x65, 0x74, 0x61, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x1f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x1f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x68,
	0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x3b, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d,
	0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x52, 0x0d, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73,
	0x12, 0x31, 0x0a, 0x0c, 0x62, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x0c, 0x62, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x12, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69,
	0x63, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x16, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x74,
	0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x42, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x74, 0x72, 0x75, 0x6e,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x42, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x73, 0x2a, 0x97, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x48, 0x61, 0x6c, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x45, 0x6d,
	0x70, 0x74, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x61, 0x6c, 0x74, 0x65,
	0x64, 0x4f, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x02, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x61, 0x6c, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x52, 0x6f, 0x6f,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67,
	0x42, 0x65, 0x74, 0x61, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x63, 0x61, 0x6e, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x63, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x65, 0x74, 0x61, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x0c, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x10, 0x0d, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65,
	0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_state_proto_rawDescOnce sync.Once
	file_synchronization_state_proto_rawDescData = file_synchronization_state_proto_rawDesc
)

func file_synchronization_state_proto_rawDescGZIP() []byte {
	file_synchronization_state_proto_rawDescOnce.Do(func() {
		file_synchronization_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_state_proto_rawDescData)
	})
	return file_synchronization_state_proto_rawDescData
}

var file_synchronization_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_synchronization_state_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: synchronization.Status
	(*State)(nil),                // 1: synchronization.State
	(*Session)(nil),              // 2: synchronization.Session
	(*rsync.ReceiverStatus)(nil), // 3: rsync.ReceiverStatus
	(*core.Conflict)(nil),        // 4: core.Conflict
	(*core.Problem)(nil),         // 5: core.Problem
}
var file_synchronization_state_proto_depIdxs = []int32{
	2, // 0: synchronization.State.session:type_name -> synchronization.Session
	0, // 1: synchronization.State.status:type_name -> synchronization.Status
	3, // 2: synchronization.State.stagingStatus:type_name -> rsync.ReceiverStatus
	4, // 3: synchronization.State.conflicts:type_name -> core.Conflict
	5, // 4: synchronization.State.alphaProblems:type_name -> core.Problem
	5, // 5: synchronization.State.betaProblems:type_name -> core.Problem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_synchronization_state_proto_init() }
func file_synchronization_state_proto_init() {
	if File_synchronization_state_proto != nil {
		return
	}
	file_synchronization_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_synchronization_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
			RawDescriptor: file_synchronization_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_state_proto_goTypes,
		DependencyIndexes: file_synchronization_state_proto_depIdxs,
		EnumInfos:         file_synchronization_state_proto_enumTypes,
		MessageInfos:      file_synchronization_state_proto_msgTypes,
	}.Build()
	File_synchronization_state_proto = out.File
	file_synchronization_state_proto_rawDesc = nil
	file_synchronization_state_proto_goTypes = nil
	file_synchronization_state_proto_depIdxs = nil
}
