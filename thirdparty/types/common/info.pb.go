// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/v1/info.proto

// Package common defines commonly used types agnostic to the node role on the Sonr network.

package common

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Direction is the direction of a stream.
type Direction int32

const (
	// Unspecified is the default value.
	Direction_DIRECTION_UNSPECIFIED Direction = 0
	// Incoming is the direction of an incoming stream.
	Direction_DIRECTION_INCOMING Direction = 1
	// Outgoing is the direction of an outgoing stream.
	Direction_DIRECTION_OUTGOING Direction = 2
)

var Direction_name = map[int32]string{
	0: "DIRECTION_UNSPECIFIED",
	1: "DIRECTION_INCOMING",
	2: "DIRECTION_OUTGOING",
}

var Direction_value = map[string]int32{
	"DIRECTION_UNSPECIFIED": 0,
	"DIRECTION_INCOMING":    1,
	"DIRECTION_OUTGOING":    2,
}

func (x Direction) String() string {
	return proto.EnumName(Direction_name, int32(x))
}

func (Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0}
}

type EntityKind int32

const (
	EntityKind_ADDRESS EntityKind = 0
	EntityKind_DID     EntityKind = 1
	EntityKind_CID     EntityKind = 2
)

var EntityKind_name = map[int32]string{
	0: "ADDRESS",
	1: "DID",
	2: "CID",
}

var EntityKind_value = map[string]int32{
	"ADDRESS": 0,
	"DID":     1,
	"CID":     2,
}

func (x EntityKind) String() string {
	return proto.EnumName(EntityKind_name, int32(x))
}

func (EntityKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{1}
}

type BlockchainModule int32

const (
	// Query x/registry module
	BlockchainModule_REGISTRY BlockchainModule = 0
	// Query x/schema module
	BlockchainModule_SCHEMA BlockchainModule = 1
	// Query x/bucket module
	BlockchainModule_BUCKET BlockchainModule = 2
)

var BlockchainModule_name = map[int32]string{
	0: "REGISTRY",
	1: "SCHEMA",
	2: "BUCKET",
}

var BlockchainModule_value = map[string]int32{
	"REGISTRY": 0,
	"SCHEMA":   1,
	"BUCKET":   2,
}

func (x BlockchainModule) String() string {
	return proto.EnumName(BlockchainModule_name, int32(x))
}

func (BlockchainModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{2}
}

// File Content Type
type MIME_Type int32

const (
	// Other File Type - If cannot derive from Subtype
	MIME_TYPE_UNSPECIFIED MIME_Type = 0
	// Sound, Audio Files
	MIME_TYPE_AUDIO MIME_Type = 1
	// Document Files - PDF, Word, Excel, etc.
	MIME_TYPE_DOCUMENT MIME_Type = 2
	// Image Files
	MIME_TYPE_IMAGE MIME_Type = 3
	// Text Based Files
	MIME_TYPE_TEXT MIME_Type = 4
	// Video Files
	MIME_TYPE_VIDEO MIME_Type = 5
	// URL Links
	MIME_TYPE_URL MIME_Type = 6
	// Crypto Files
	MIME_TYPE_CRYPTO MIME_Type = 7
)

var MIME_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_AUDIO",
	2: "TYPE_DOCUMENT",
	3: "TYPE_IMAGE",
	4: "TYPE_TEXT",
	5: "TYPE_VIDEO",
	6: "TYPE_URL",
	7: "TYPE_CRYPTO",
}

var MIME_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"TYPE_AUDIO":       1,
	"TYPE_DOCUMENT":    2,
	"TYPE_IMAGE":       3,
	"TYPE_TEXT":        4,
	"TYPE_VIDEO":       5,
	"TYPE_URL":         6,
	"TYPE_CRYPTO":      7,
}

func (x MIME_Type) String() string {
	return proto.EnumName(MIME_Type_name, int32(x))
}

func (MIME_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0, 0}
}

// Peers Active Status
type Peer_Status int32

const (
	Peer_STATUS_UNSPECIFIED Peer_Status = 0
	Peer_STATUS_ONLINE      Peer_Status = 1
	Peer_STATUS_AWAY        Peer_Status = 2
	Peer_STATUS_BUSY        Peer_Status = 3
)

var Peer_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ONLINE",
	2: "STATUS_AWAY",
	3: "STATUS_BUSY",
}

var Peer_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ONLINE":      1,
	"STATUS_AWAY":        2,
	"STATUS_BUSY":        3,
}

func (x Peer_Status) String() string {
	return proto.EnumName(Peer_Status_name, int32(x))
}

func (Peer_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{1, 0}
}

// Standard MIME with Additional Extensions
type MIME struct {
	// Type of File
	Type MIME_Type `protobuf:"varint,1,opt,name=type,proto3,enum=sonrio.common.v1.MIME_Type" json:"type,omitempty"`
	// Extension of File
	Subtype string `protobuf:"bytes,2,opt,name=subtype,proto3" json:"subtype,omitempty"`
	// Type/Subtype i.e. (image/jpeg)
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MIME) Reset()         { *m = MIME{} }
func (m *MIME) String() string { return proto.CompactTextString(m) }
func (*MIME) ProtoMessage()    {}
func (*MIME) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0}
}
func (m *MIME) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MIME) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MIME.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MIME) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MIME.Merge(m, src)
}
func (m *MIME) XXX_Size() int {
	return m.Size()
}
func (m *MIME) XXX_DiscardUnknown() {
	xxx_messageInfo_MIME.DiscardUnknown(m)
}

var xxx_messageInfo_MIME proto.InternalMessageInfo

func (m *MIME) GetType() MIME_Type {
	if m != nil {
		return m.Type
	}
	return MIME_TYPE_UNSPECIFIED
}

func (m *MIME) GetSubtype() string {
	if m != nil {
		return m.Subtype
	}
	return ""
}

func (m *MIME) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Basic Info Sent to Peers to Establish Connections
type Peer struct {
	PeerId string      `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Did    string      `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Status Peer_Status `protobuf:"varint,3,opt,name=status,proto3,enum=sonrio.common.v1.Peer_Status" json:"status,omitempty"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{1}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *Peer) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *Peer) GetStatus() Peer_Status {
	if m != nil {
		return m.Status
	}
	return Peer_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("sonrio.common.v1.Direction", Direction_name, Direction_value)
	proto.RegisterEnum("sonrio.common.v1.EntityKind", EntityKind_name, EntityKind_value)
	proto.RegisterEnum("sonrio.common.v1.BlockchainModule", BlockchainModule_name, BlockchainModule_value)
	proto.RegisterEnum("sonrio.common.v1.MIME_Type", MIME_Type_name, MIME_Type_value)
	proto.RegisterEnum("sonrio.common.v1.Peer_Status", Peer_Status_name, Peer_Status_value)
	proto.RegisterType((*MIME)(nil), "sonrio.common.v1.MIME")
	proto.RegisterType((*Peer)(nil), "sonrio.common.v1.Peer")
}

func init() { proto.RegisterFile("common/v1/info.proto", fileDescriptor_a6ffb5b3e6a498f4) }

var fileDescriptor_a6ffb5b3e6a498f4 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0xe3, 0xb4, 0x4b, 0xe9, 0x7f, 0x6c, 0x78, 0xd6, 0x18, 0x43, 0x88, 0x68, 0xea, 0x69,
	0x1a, 0x22, 0xd1, 0x86, 0x90, 0xb8, 0xa6, 0x89, 0x29, 0xd6, 0x96, 0xa4, 0x4a, 0x9c, 0x41, 0xb9,
	0x4c, 0x5d, 0x13, 0x98, 0xc5, 0x96, 0x54, 0x69, 0x5a, 0xa9, 0x1f, 0x02, 0x89, 0x6f, 0xc4, 0x95,
	0xe3, 0x8e, 0x1c, 0x51, 0xfb, 0x21, 0xb8, 0x22, 0x3b, 0x65, 0x40, 0x39, 0xe5, 0xff, 0x7e, 0xff,
	0x17, 0x27, 0xef, 0xc9, 0xb0, 0x3b, 0x2a, 0x6e, 0x6e, 0x8a, 0xdc, 0x9e, 0x1d, 0xdb, 0x22, 0xff,
	0x50, 0x58, 0xe3, 0xb2, 0xa8, 0x0a, 0x82, 0x27, 0x45, 0x5e, 0x8a, 0xc2, 0xaa, 0x97, 0xd6, 0xec,
	0xb8, 0xf3, 0x13, 0x41, 0xd3, 0x67, 0x3e, 0x25, 0x36, 0x34, 0xab, 0xf9, 0x38, 0xdb, 0x47, 0x07,
	0xe8, 0x70, 0xfb, 0xe4, 0x89, 0xb5, 0xee, 0xb4, 0xa4, 0xcb, 0xe2, 0xf3, 0x71, 0x16, 0x29, 0x23,
	0xd9, 0x87, 0xd6, 0x64, 0x7a, 0xa9, 0xde, 0xd1, 0x0f, 0xd0, 0x61, 0x3b, 0xfa, 0x2d, 0xc9, 0x2e,
	0x6c, 0xcc, 0x86, 0xd7, 0xd3, 0x6c, 0xbf, 0xa1, 0x78, 0x2d, 0x3a, 0x9f, 0x11, 0x34, 0x79, 0xbd,
	0xc6, 0x7c, 0xd0, 0xa7, 0x17, 0x49, 0x10, 0xf7, 0xa9, 0xcb, 0x5e, 0x33, 0xea, 0x61, 0x8d, 0x6c,
	0x03, 0x28, 0xea, 0x24, 0x1e, 0x0b, 0x31, 0x22, 0x3b, 0xb0, 0xa5, 0xb4, 0x17, 0xba, 0x89, 0x4f,
	0x03, 0x8e, 0xf5, 0x3b, 0x0b, 0xf3, 0x9d, 0x1e, 0xc5, 0x0d, 0xb2, 0x05, 0x6d, 0xa5, 0x39, 0x7d,
	0xc7, 0x71, 0xf3, 0x6e, 0x7d, 0xce, 0x3c, 0x1a, 0xe2, 0x0d, 0x72, 0x1f, 0xee, 0xd5, 0xdf, 0x89,
	0xce, 0xb0, 0x41, 0x1e, 0xc0, 0xa6, 0x52, 0x6e, 0x34, 0xe8, 0xf3, 0x10, 0xb7, 0x3a, 0x5f, 0x11,
	0x34, 0xfb, 0x59, 0x56, 0x92, 0x47, 0xd0, 0x1a, 0x67, 0x59, 0x79, 0x21, 0x52, 0x15, 0xbe, 0x1d,
	0x19, 0x52, 0xb2, 0x94, 0x60, 0x68, 0xa4, 0x22, 0x5d, 0xa5, 0x93, 0x23, 0x79, 0x09, 0xc6, 0xa4,
	0x1a, 0x56, 0xd3, 0x89, 0x8a, 0xb6, 0x7d, 0xf2, 0xf4, 0xff, 0x9a, 0xe4, 0x91, 0x56, 0xac, 0x4c,
	0xd1, 0xca, 0xdc, 0x49, 0xc0, 0xa8, 0x09, 0xd9, 0x03, 0x12, 0x73, 0x87, 0x27, 0xf1, 0x5a, 0xfa,
	0x1d, 0xd8, 0x5a, 0xf1, 0x30, 0x38, 0x63, 0x01, 0xc5, 0x48, 0xfe, 0xf0, 0x0a, 0x39, 0x6f, 0x9d,
	0x01, 0xd6, 0xff, 0x02, 0xdd, 0x24, 0x1e, 0xe0, 0xc6, 0xd1, 0x39, 0xb4, 0x3d, 0x51, 0x66, 0xa3,
	0x4a, 0x14, 0x39, 0x79, 0x0c, 0x0f, 0x3d, 0x16, 0x51, 0x97, 0xb3, 0x30, 0x58, 0x3b, 0x7c, 0x0f,
	0xc8, 0x9f, 0x15, 0x0b, 0xdc, 0xd0, 0x67, 0x41, 0x0f, 0xa3, 0x7f, 0x79, 0x98, 0xf0, 0x5e, 0x28,
	0xb9, 0x7e, 0xf4, 0x0c, 0x80, 0xe6, 0x95, 0xa8, 0xe6, 0xa7, 0x22, 0x4f, 0xc9, 0x26, 0xb4, 0x1c,
	0xcf, 0x8b, 0x68, 0x1c, 0x63, 0x8d, 0xb4, 0xa0, 0xe1, 0x31, 0x0f, 0x23, 0x39, 0xb8, 0xcc, 0xc3,
	0xfa, 0xd1, 0x2b, 0xc0, 0xdd, 0xeb, 0x62, 0xf4, 0x69, 0x74, 0x35, 0x14, 0xb9, 0x5f, 0xa4, 0xd3,
	0xeb, 0x4c, 0x36, 0x1f, 0xd1, 0x1e, 0x8b, 0x79, 0x34, 0xc0, 0x1a, 0x01, 0x30, 0x62, 0xf7, 0x0d,
	0xf5, 0x1d, 0x8c, 0xe4, 0xdc, 0x4d, 0xdc, 0x53, 0xca, 0xb1, 0xde, 0x65, 0xdf, 0x16, 0x26, 0xba,
	0x5d, 0x98, 0xe8, 0xc7, 0xc2, 0x44, 0x5f, 0x96, 0xa6, 0x76, 0xbb, 0x34, 0xb5, 0xef, 0x4b, 0x53,
	0x7b, 0x6f, 0x7f, 0x14, 0xd5, 0xd5, 0xf4, 0x52, 0xb6, 0x6a, 0xcb, 0x82, 0x9f, 0x8b, 0x42, 0x3d,
	0xed, 0xea, 0x4a, 0x94, 0xe9, 0x78, 0x58, 0x56, 0x73, 0x5b, 0xde, 0xb4, 0x89, 0x5d, 0xf7, 0x7e,
	0x69, 0xa8, 0xeb, 0xfd, 0xe2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x3c, 0x14, 0xf0, 0xf6,
	0x02, 0x00, 0x00,
}

func (m *MIME) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MIME) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MIME) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Subtype) > 0 {
		i -= len(m.Subtype)
		copy(dAtA[i:], m.Subtype)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Subtype)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PeerId) > 0 {
		i -= len(m.PeerId)
		copy(dAtA[i:], m.PeerId)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.PeerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MIME) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovInfo(uint64(m.Type))
	}
	l = len(m.Subtype)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	return n
}

func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovInfo(uint64(m.Status))
	}
	return n
}

func sovInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MIME) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MIME: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MIME: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= MIME_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subtype", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subtype = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Peer_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInfo = fmt.Errorf("proto: unexpected end of group")
)
