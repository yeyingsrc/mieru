// Copyright (C) 2024  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: misc.proto

package appctlpb

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

type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON dump of metrics.
	Json *string `protobuf:"bytes,1,opt,name=json,proto3,oneof" json:"json,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{0}
}

func (x *Metrics) GetJson() string {
	if x != nil && x.Json != nil {
		return *x.Json
	}
	return ""
}

type ProfileSavePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Location to save profile results.
	FilePath *string `protobuf:"bytes,1,opt,name=filePath,proto3,oneof" json:"filePath,omitempty"`
}

func (x *ProfileSavePath) Reset() {
	*x = ProfileSavePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileSavePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileSavePath) ProtoMessage() {}

func (x *ProfileSavePath) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileSavePath.ProtoReflect.Descriptor instead.
func (*ProfileSavePath) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileSavePath) GetFilePath() string {
	if x != nil && x.FilePath != nil {
		return *x.FilePath
	}
	return ""
}

type SessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Protocol     *string                `protobuf:"bytes,2,opt,name=protocol,proto3,oneof" json:"protocol,omitempty"`
	LocalAddr    *string                `protobuf:"bytes,3,opt,name=localAddr,proto3,oneof" json:"localAddr,omitempty"`
	RemoteAddr   *string                `protobuf:"bytes,4,opt,name=remoteAddr,proto3,oneof" json:"remoteAddr,omitempty"`
	State        *string                `protobuf:"bytes,5,opt,name=state,proto3,oneof" json:"state,omitempty"`
	RecvQ        *uint32                `protobuf:"varint,6,opt,name=recvQ,proto3,oneof" json:"recvQ,omitempty"`
	RecvBuf      *uint32                `protobuf:"varint,7,opt,name=recvBuf,proto3,oneof" json:"recvBuf,omitempty"`
	SendQ        *uint32                `protobuf:"varint,8,opt,name=sendQ,proto3,oneof" json:"sendQ,omitempty"`
	SendBuf      *uint32                `protobuf:"varint,9,opt,name=sendBuf,proto3,oneof" json:"sendBuf,omitempty"`
	LastRecvTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=lastRecvTime,proto3,oneof" json:"lastRecvTime,omitempty"`
	LastRecvSeq  *uint32                `protobuf:"varint,11,opt,name=lastRecvSeq,proto3,oneof" json:"lastRecvSeq,omitempty"`
	LastSendTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=lastSendTime,proto3,oneof" json:"lastSendTime,omitempty"`
	LastSendSeq  *uint32                `protobuf:"varint,13,opt,name=lastSendSeq,proto3,oneof" json:"lastSendSeq,omitempty"`
}

func (x *SessionInfo) Reset() {
	*x = SessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInfo) ProtoMessage() {}

func (x *SessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInfo.ProtoReflect.Descriptor instead.
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{2}
}

func (x *SessionInfo) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *SessionInfo) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

func (x *SessionInfo) GetLocalAddr() string {
	if x != nil && x.LocalAddr != nil {
		return *x.LocalAddr
	}
	return ""
}

func (x *SessionInfo) GetRemoteAddr() string {
	if x != nil && x.RemoteAddr != nil {
		return *x.RemoteAddr
	}
	return ""
}

func (x *SessionInfo) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

func (x *SessionInfo) GetRecvQ() uint32 {
	if x != nil && x.RecvQ != nil {
		return *x.RecvQ
	}
	return 0
}

func (x *SessionInfo) GetRecvBuf() uint32 {
	if x != nil && x.RecvBuf != nil {
		return *x.RecvBuf
	}
	return 0
}

func (x *SessionInfo) GetSendQ() uint32 {
	if x != nil && x.SendQ != nil {
		return *x.SendQ
	}
	return 0
}

func (x *SessionInfo) GetSendBuf() uint32 {
	if x != nil && x.SendBuf != nil {
		return *x.SendBuf
	}
	return 0
}

func (x *SessionInfo) GetLastRecvTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRecvTime
	}
	return nil
}

func (x *SessionInfo) GetLastRecvSeq() uint32 {
	if x != nil && x.LastRecvSeq != nil {
		return *x.LastRecvSeq
	}
	return 0
}

func (x *SessionInfo) GetLastSendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSendTime
	}
	return nil
}

func (x *SessionInfo) GetLastSendSeq() uint32 {
	if x != nil && x.LastSendSeq != nil {
		return *x.LastSendSeq
	}
	return 0
}

type SessionInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SessionInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SessionInfoList) Reset() {
	*x = SessionInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInfoList) ProtoMessage() {}

func (x *SessionInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInfoList.ProtoReflect.Descriptor instead.
func (*SessionInfoList) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{3}
}

func (x *SessionInfoList) GetItems() []*SessionInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type ThreadDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full thread dump of the application.
	ThreadDump *string `protobuf:"bytes,1,opt,name=threadDump,proto3,oneof" json:"threadDump,omitempty"`
}

func (x *ThreadDump) Reset() {
	*x = ThreadDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadDump) ProtoMessage() {}

func (x *ThreadDump) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadDump.ProtoReflect.Descriptor instead.
func (*ThreadDump) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{4}
}

func (x *ThreadDump) GetThreadDump() string {
	if x != nil && x.ThreadDump != nil {
		return *x.ThreadDump
	}
	return ""
}

type MemoryStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeapBytes       *uint64 `protobuf:"varint,2,opt,name=heapBytes,proto3,oneof" json:"heapBytes,omitempty"`
	HeapObjects     *uint64 `protobuf:"varint,3,opt,name=heapObjects,proto3,oneof" json:"heapObjects,omitempty"`
	MaxHeapBytes    *uint64 `protobuf:"varint,4,opt,name=maxHeapBytes,proto3,oneof" json:"maxHeapBytes,omitempty"`
	TargetHeapBytes *uint64 `protobuf:"varint,5,opt,name=targetHeapBytes,proto3,oneof" json:"targetHeapBytes,omitempty"`
	StackBytes      *uint64 `protobuf:"varint,6,opt,name=stackBytes,proto3,oneof" json:"stackBytes,omitempty"`
}

func (x *MemoryStatistics) Reset() {
	*x = MemoryStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_misc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryStatistics) ProtoMessage() {}

func (x *MemoryStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_misc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryStatistics.ProtoReflect.Descriptor instead.
func (*MemoryStatistics) Descriptor() ([]byte, []int) {
	return file_misc_proto_rawDescGZIP(), []int{5}
}

func (x *MemoryStatistics) GetHeapBytes() uint64 {
	if x != nil && x.HeapBytes != nil {
		return *x.HeapBytes
	}
	return 0
}

func (x *MemoryStatistics) GetHeapObjects() uint64 {
	if x != nil && x.HeapObjects != nil {
		return *x.HeapObjects
	}
	return 0
}

func (x *MemoryStatistics) GetMaxHeapBytes() uint64 {
	if x != nil && x.MaxHeapBytes != nil {
		return *x.MaxHeapBytes
	}
	return 0
}

func (x *MemoryStatistics) GetTargetHeapBytes() uint64 {
	if x != nil && x.TargetHeapBytes != nil {
		return *x.TargetHeapBytes
	}
	return 0
}

func (x *MemoryStatistics) GetStackBytes() uint64 {
	if x != nil && x.StackBytes != nil {
		return *x.StackBytes
	}
	return 0
}

var File_misc_proto protoreflect.FileDescriptor

var file_misc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70,
	0x70, 0x63, 0x74, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x17, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6a, 0x73,
	0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x9b, 0x05, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x72, 0x65, 0x63, 0x76, 0x51, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x05, 0x72,
	0x65, 0x63, 0x76, 0x51, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x76, 0x42,
	0x75, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x06, 0x52, 0x07, 0x72, 0x65, 0x63, 0x76,
	0x42, 0x75, 0x66, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x65, 0x6e, 0x64, 0x51, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x05, 0x73, 0x65, 0x6e, 0x64, 0x51, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x66, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x08, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x66, 0x88, 0x01, 0x01,
	0x12, 0x43, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x76, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x76, 0x54, 0x69,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63,
	0x76, 0x53, 0x65, 0x71, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x0a, 0x52, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x76, 0x53, 0x65, 0x71, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x0b,
	0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x71,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x0c, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x65, 0x71, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x51, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x42, 0x75, 0x66, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x51, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x66, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x76, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x76, 0x53, 0x65, 0x71,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65,
	0x71, 0x22, 0x3c, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x40, 0x0a, 0x0a, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x23, 0x0a,
	0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d, 0x70, 0x88,
	0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x44, 0x75, 0x6d,
	0x70, 0x22, 0xb1, 0x02, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x70, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x68, 0x65, 0x61,
	0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x68, 0x65, 0x61,
	0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01,
	0x52, 0x0b, 0x68, 0x65, 0x61, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x27, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61,
	0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x48, 0x65, 0x61, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x03, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x65, 0x61, 0x70,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x04, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x68, 0x65, 0x61, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x68, 0x65, 0x61, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x12, 0x0a, 0x10,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x65, 0x61, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x66, 0x65, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x65, 0x72, 0x75,
	0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x63, 0x74, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_misc_proto_rawDescOnce sync.Once
	file_misc_proto_rawDescData = file_misc_proto_rawDesc
)

func file_misc_proto_rawDescGZIP() []byte {
	file_misc_proto_rawDescOnce.Do(func() {
		file_misc_proto_rawDescData = protoimpl.X.CompressGZIP(file_misc_proto_rawDescData)
	})
	return file_misc_proto_rawDescData
}

var file_misc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_misc_proto_goTypes = []interface{}{
	(*Metrics)(nil),               // 0: appctl.Metrics
	(*ProfileSavePath)(nil),       // 1: appctl.ProfileSavePath
	(*SessionInfo)(nil),           // 2: appctl.SessionInfo
	(*SessionInfoList)(nil),       // 3: appctl.SessionInfoList
	(*ThreadDump)(nil),            // 4: appctl.ThreadDump
	(*MemoryStatistics)(nil),      // 5: appctl.MemoryStatistics
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_misc_proto_depIdxs = []int32{
	6, // 0: appctl.SessionInfo.lastRecvTime:type_name -> google.protobuf.Timestamp
	6, // 1: appctl.SessionInfo.lastSendTime:type_name -> google.protobuf.Timestamp
	2, // 2: appctl.SessionInfoList.items:type_name -> appctl.SessionInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_misc_proto_init() }
func file_misc_proto_init() {
	if File_misc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_misc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
		file_misc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileSavePath); i {
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
		file_misc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionInfo); i {
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
		file_misc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionInfoList); i {
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
		file_misc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadDump); i {
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
		file_misc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryStatistics); i {
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
	file_misc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_misc_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_misc_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_misc_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_misc_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_misc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_misc_proto_goTypes,
		DependencyIndexes: file_misc_proto_depIdxs,
		MessageInfos:      file_misc_proto_msgTypes,
	}.Build()
	File_misc_proto = out.File
	file_misc_proto_rawDesc = nil
	file_misc_proto_goTypes = nil
	file_misc_proto_depIdxs = nil
}
