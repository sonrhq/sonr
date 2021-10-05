/// This file contains service for the Node RPC Server

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/node/client.proto

package node

import (
	common "github.com/sonr-io/core/internal/common"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_node_client_proto_rawDescGZIP(), []int{0}
}

var File_proto_node_client_proto protoreflect.FileDescriptor

var file_proto_node_client_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xa5, 0x07,
	0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x06, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x05, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x05, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x12, 0x19, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04,
	0x53, 0x74, 0x61, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x4f, 0x6e, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x10, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x10, 0x4f, 0x6e, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x12, 0x4f,
	0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x12, 0x10, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x44, 0x0a, 0x12, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44,
	0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x10, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x10, 0x4f, 0x6e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x12, 0x4f, 0x6e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x18, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x44, 0x0a, 0x12, 0x4f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_node_client_proto_rawDescOnce sync.Once
	file_proto_node_client_proto_rawDescData = file_proto_node_client_proto_rawDesc
)

func file_proto_node_client_proto_rawDescGZIP() []byte {
	file_proto_node_client_proto_rawDescOnce.Do(func() {
		file_proto_node_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_node_client_proto_rawDescData)
	})
	return file_proto_node_client_proto_rawDescData
}

var file_proto_node_client_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_node_client_proto_goTypes = []interface{}{
	(*Empty)(nil),                // 0: sonr.node.Empty
	(*SupplyRequest)(nil),        // 1: sonr.node.SupplyRequest
	(*EditRequest)(nil),          // 2: sonr.node.EditRequest
	(*FetchRequest)(nil),         // 3: sonr.node.FetchRequest
	(*ShareRequest)(nil),         // 4: sonr.node.ShareRequest
	(*RespondRequest)(nil),       // 5: sonr.node.RespondRequest
	(*SearchRequest)(nil),        // 6: sonr.node.SearchRequest
	(*StatRequest)(nil),          // 7: sonr.node.StatRequest
	(*SupplyResponse)(nil),       // 8: sonr.node.SupplyResponse
	(*EditResponse)(nil),         // 9: sonr.node.EditResponse
	(*FetchResponse)(nil),        // 10: sonr.node.FetchResponse
	(*ShareResponse)(nil),        // 11: sonr.node.ShareResponse
	(*RespondResponse)(nil),      // 12: sonr.node.RespondResponse
	(*SearchResponse)(nil),       // 13: sonr.node.SearchResponse
	(*StatResponse)(nil),         // 14: sonr.node.StatResponse
	(*common.RefreshEvent)(nil),  // 15: sonr.core.RefreshEvent
	(*common.MailboxEvent)(nil),  // 16: sonr.core.MailboxEvent
	(*common.DecisionEvent)(nil), // 17: sonr.core.DecisionEvent
	(*common.InviteEvent)(nil),   // 18: sonr.core.InviteEvent
	(*common.ProgressEvent)(nil), // 19: sonr.core.ProgressEvent
	(*common.CompleteEvent)(nil), // 20: sonr.core.CompleteEvent
}
var file_proto_node_client_proto_depIdxs = []int32{
	1,  // 0: sonr.node.ClientService.Supply:input_type -> sonr.node.SupplyRequest
	2,  // 1: sonr.node.ClientService.Edit:input_type -> sonr.node.EditRequest
	3,  // 2: sonr.node.ClientService.Fetch:input_type -> sonr.node.FetchRequest
	4,  // 3: sonr.node.ClientService.Share:input_type -> sonr.node.ShareRequest
	5,  // 4: sonr.node.ClientService.Respond:input_type -> sonr.node.RespondRequest
	6,  // 5: sonr.node.ClientService.Search:input_type -> sonr.node.SearchRequest
	7,  // 6: sonr.node.ClientService.Stat:input_type -> sonr.node.StatRequest
	0,  // 7: sonr.node.ClientService.OnLobbyRefresh:input_type -> sonr.node.Empty
	0,  // 8: sonr.node.ClientService.OnMailboxMessage:input_type -> sonr.node.Empty
	0,  // 9: sonr.node.ClientService.OnTransferAccepted:input_type -> sonr.node.Empty
	0,  // 10: sonr.node.ClientService.OnTransferDeclined:input_type -> sonr.node.Empty
	0,  // 11: sonr.node.ClientService.OnTransferInvite:input_type -> sonr.node.Empty
	0,  // 12: sonr.node.ClientService.OnTransferProgress:input_type -> sonr.node.Empty
	0,  // 13: sonr.node.ClientService.OnTransferComplete:input_type -> sonr.node.Empty
	8,  // 14: sonr.node.ClientService.Supply:output_type -> sonr.node.SupplyResponse
	9,  // 15: sonr.node.ClientService.Edit:output_type -> sonr.node.EditResponse
	10, // 16: sonr.node.ClientService.Fetch:output_type -> sonr.node.FetchResponse
	11, // 17: sonr.node.ClientService.Share:output_type -> sonr.node.ShareResponse
	12, // 18: sonr.node.ClientService.Respond:output_type -> sonr.node.RespondResponse
	13, // 19: sonr.node.ClientService.Search:output_type -> sonr.node.SearchResponse
	14, // 20: sonr.node.ClientService.Stat:output_type -> sonr.node.StatResponse
	15, // 21: sonr.node.ClientService.OnLobbyRefresh:output_type -> sonr.core.RefreshEvent
	16, // 22: sonr.node.ClientService.OnMailboxMessage:output_type -> sonr.core.MailboxEvent
	17, // 23: sonr.node.ClientService.OnTransferAccepted:output_type -> sonr.core.DecisionEvent
	17, // 24: sonr.node.ClientService.OnTransferDeclined:output_type -> sonr.core.DecisionEvent
	18, // 25: sonr.node.ClientService.OnTransferInvite:output_type -> sonr.core.InviteEvent
	19, // 26: sonr.node.ClientService.OnTransferProgress:output_type -> sonr.core.ProgressEvent
	20, // 27: sonr.node.ClientService.OnTransferComplete:output_type -> sonr.core.CompleteEvent
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_node_client_proto_init() }
func file_proto_node_client_proto_init() {
	if File_proto_node_client_proto != nil {
		return
	}
	file_proto_node_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_node_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_node_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_node_client_proto_goTypes,
		DependencyIndexes: file_proto_node_client_proto_depIdxs,
		MessageInfos:      file_proto_node_client_proto_msgTypes,
	}.Build()
	File_proto_node_client_proto = out.File
	file_proto_node_client_proto_rawDesc = nil
	file_proto_node_client_proto_goTypes = nil
	file_proto_node_client_proto_depIdxs = nil
}
