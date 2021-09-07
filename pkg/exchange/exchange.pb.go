// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: api/protocols/exchange.proto

package exchange

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

// Local Type Event
type ExchangeEvent_Subject int32

const (
	ExchangeEvent_JOIN    ExchangeEvent_Subject = 0 // Peer has joined wants Info
	ExchangeEvent_STANDBY ExchangeEvent_Subject = 1 // Peer is taking break from lobby
	ExchangeEvent_RESUME  ExchangeEvent_Subject = 2 // Peer has rejoined Lobby
	ExchangeEvent_UPDATE  ExchangeEvent_Subject = 3 // Peer has changed position
	ExchangeEvent_EXIT    ExchangeEvent_Subject = 4 // Peer has Exited
)

// Enum value maps for ExchangeEvent_Subject.
var (
	ExchangeEvent_Subject_name = map[int32]string{
		0: "JOIN",
		1: "STANDBY",
		2: "RESUME",
		3: "UPDATE",
		4: "EXIT",
	}
	ExchangeEvent_Subject_value = map[string]int32{
		"JOIN":    0,
		"STANDBY": 1,
		"RESUME":  2,
		"UPDATE":  3,
		"EXIT":    4,
	}
)

func (x ExchangeEvent_Subject) Enum() *ExchangeEvent_Subject {
	p := new(ExchangeEvent_Subject)
	*p = x
	return p
}

func (x ExchangeEvent_Subject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExchangeEvent_Subject) Descriptor() protoreflect.EnumDescriptor {
	return file_api_protocols_exchange_proto_enumTypes[0].Descriptor()
}

func (ExchangeEvent_Subject) Type() protoreflect.EnumType {
	return &file_api_protocols_exchange_proto_enumTypes[0]
}

func (x ExchangeEvent_Subject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExchangeEvent_Subject.Descriptor instead.
func (ExchangeEvent_Subject) EnumDescriptor() ([]byte, []int) {
	return file_api_protocols_exchange_proto_rawDescGZIP(), []int{0, 0}
}

// Message Sent when peer messages Local Room
type ExchangeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Olc     string                `protobuf:"bytes,1,opt,name=olc,proto3" json:"olc,omitempty"`                                                             // OLC Code of Topic
	Id      string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                                               // Peer ID
	Peer    *common.Peer          `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`                                                           // User Information
	Subject ExchangeEvent_Subject `protobuf:"varint,4,opt,name=subject,proto3,enum=sonr.protocols.exchange.ExchangeEvent_Subject" json:"subject,omitempty"` // Local Event Subject
}

func (x *ExchangeEvent) Reset() {
	*x = ExchangeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocols_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeEvent) ProtoMessage() {}

func (x *ExchangeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocols_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeEvent.ProtoReflect.Descriptor instead.
func (*ExchangeEvent) Descriptor() ([]byte, []int) {
	return file_api_protocols_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeEvent) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *ExchangeEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExchangeEvent) GetPeer() *common.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *ExchangeEvent) GetSubject() ExchangeEvent_Subject {
	if x != nil {
		return x.Subject
	}
	return ExchangeEvent_JOIN
}

// JoinExchangeRequest is Message for Signing Request (Hmac Sha256)
type JoinExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value to be signed
	SName string       `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"` // SName combined with Device ID and Hashed
	Peer  *common.Peer `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *JoinExchangeRequest) Reset() {
	*x = JoinExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocols_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinExchangeRequest) ProtoMessage() {}

func (x *JoinExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocols_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinExchangeRequest.ProtoReflect.Descriptor instead.
func (*JoinExchangeRequest) Descriptor() ([]byte, []int) {
	return file_api_protocols_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *JoinExchangeRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *JoinExchangeRequest) GetPeer() *common.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

// JoinExchangeResponse is Message for Signing Response (Hmac Sha256)
type JoinExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool         `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	SName   string       `protobuf:"bytes,2,opt,name=sName,proto3" json:"sName,omitempty"` // SName combined with Device ID and Hashed
	Peer    *common.Peer `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`   // Peer Data
}

func (x *JoinExchangeResponse) Reset() {
	*x = JoinExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocols_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinExchangeResponse) ProtoMessage() {}

func (x *JoinExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocols_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinExchangeResponse.ProtoReflect.Descriptor instead.
func (*JoinExchangeResponse) Descriptor() ([]byte, []int) {
	return file_api_protocols_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *JoinExchangeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *JoinExchangeResponse) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *JoinExchangeResponse) GetPeer() *common.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

// UpdateRequest is message for Exchange Protocol
type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Olc  string       `protobuf:"bytes,1,opt,name=olc,proto3" json:"olc,omitempty"`   // Users Open Location Code
	Peer *common.Peer `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"` // Users Peer Data
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocols_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocols_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_protocols_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetOlc() string {
	if x != nil {
		return x.Olc
	}
	return ""
}

func (x *UpdateRequest) GetPeer() *common.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

var File_api_protocols_exchange_proto protoreflect.FileDescriptor

var file_api_protocols_exchange_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4,
	0x01, 0x0a, 0x0d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x6c, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x42, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x08, 0x0a, 0x04,
	0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x42,
	0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x53, 0x55, 0x4d, 0x45, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x45,
	0x58, 0x49, 0x54, 0x10, 0x04, 0x22, 0x50, 0x0a, 0x13, 0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x6b, 0x0a, 0x14, 0x4a, 0x6f, 0x69, 0x6e, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04,
	0x70, 0x65, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6f, 0x6c, 0x63, 0x12, 0x23, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d,
	0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protocols_exchange_proto_rawDescOnce sync.Once
	file_api_protocols_exchange_proto_rawDescData = file_api_protocols_exchange_proto_rawDesc
)

func file_api_protocols_exchange_proto_rawDescGZIP() []byte {
	file_api_protocols_exchange_proto_rawDescOnce.Do(func() {
		file_api_protocols_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protocols_exchange_proto_rawDescData)
	})
	return file_api_protocols_exchange_proto_rawDescData
}

var file_api_protocols_exchange_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_protocols_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_protocols_exchange_proto_goTypes = []interface{}{
	(ExchangeEvent_Subject)(0),   // 0: sonr.protocols.exchange.ExchangeEvent.Subject
	(*ExchangeEvent)(nil),        // 1: sonr.protocols.exchange.ExchangeEvent
	(*JoinExchangeRequest)(nil),  // 2: sonr.protocols.exchange.JoinExchangeRequest
	(*JoinExchangeResponse)(nil), // 3: sonr.protocols.exchange.JoinExchangeResponse
	(*UpdateRequest)(nil),        // 4: sonr.protocols.exchange.UpdateRequest
	(*common.Peer)(nil),          // 5: sonr.core.Peer
}
var file_api_protocols_exchange_proto_depIdxs = []int32{
	5, // 0: sonr.protocols.exchange.ExchangeEvent.peer:type_name -> sonr.core.Peer
	0, // 1: sonr.protocols.exchange.ExchangeEvent.subject:type_name -> sonr.protocols.exchange.ExchangeEvent.Subject
	5, // 2: sonr.protocols.exchange.JoinExchangeRequest.peer:type_name -> sonr.core.Peer
	5, // 3: sonr.protocols.exchange.JoinExchangeResponse.peer:type_name -> sonr.core.Peer
	5, // 4: sonr.protocols.exchange.UpdateRequest.peer:type_name -> sonr.core.Peer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_protocols_exchange_proto_init() }
func file_api_protocols_exchange_proto_init() {
	if File_api_protocols_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protocols_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeEvent); i {
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
		file_api_protocols_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinExchangeRequest); i {
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
		file_api_protocols_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinExchangeResponse); i {
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
		file_api_protocols_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
			RawDescriptor: file_api_protocols_exchange_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protocols_exchange_proto_goTypes,
		DependencyIndexes: file_api_protocols_exchange_proto_depIdxs,
		EnumInfos:         file_api_protocols_exchange_proto_enumTypes,
		MessageInfos:      file_api_protocols_exchange_proto_msgTypes,
	}.Build()
	File_api_protocols_exchange_proto = out.File
	file_api_protocols_exchange_proto_rawDesc = nil
	file_api_protocols_exchange_proto_goTypes = nil
	file_api_protocols_exchange_proto_depIdxs = nil
}
