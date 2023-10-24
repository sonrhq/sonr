// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/identity/identifier.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// UserIdentifierType defines the type of the user identifier\
type UserIdentifierType int32

const (
	UserIdentifierType_NONE    UserIdentifierType = 0
	UserIdentifierType_EMAIL   UserIdentifierType = 1
	UserIdentifierType_PHONE   UserIdentifierType = 2
	UserIdentifierType_FIDO    UserIdentifierType = 3
	UserIdentifierType_PASSKEY UserIdentifierType = 4
	UserIdentifierType_GPG     UserIdentifierType = 5
)

var UserIdentifierType_name = map[int32]string{
	0: "NONE",
	1: "EMAIL",
	2: "PHONE",
	3: "FIDO",
	4: "PASSKEY",
	5: "GPG",
}

var UserIdentifierType_value = map[string]int32{
	"NONE":    0,
	"EMAIL":   1,
	"PHONE":   2,
	"FIDO":    3,
	"PASSKEY": 4,
	"GPG":     5,
}

func (x UserIdentifierType) String() string {
	return proto.EnumName(UserIdentifierType_name, int32(x))
}

func (UserIdentifierType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_279efc507d29e268, []int{0}
}

// UserIdentifier defines a user identifier
type UserIdentifier struct {
	Type       UserIdentifierType `protobuf:"varint,1,opt,name=type,proto3,enum=core.identity.UserIdentifierType" json:"type,omitempty"`
	Identifier string             `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (m *UserIdentifier) Reset()         { *m = UserIdentifier{} }
func (m *UserIdentifier) String() string { return proto.CompactTextString(m) }
func (*UserIdentifier) ProtoMessage()    {}
func (*UserIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_279efc507d29e268, []int{0}
}
func (m *UserIdentifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserIdentifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdentifier.Merge(m, src)
}
func (m *UserIdentifier) XXX_Size() int {
	return m.Size()
}
func (m *UserIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdentifier proto.InternalMessageInfo

func (m *UserIdentifier) GetType() UserIdentifierType {
	if m != nil {
		return m.Type
	}
	return UserIdentifierType_NONE
}

func (m *UserIdentifier) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

// WalletIdentifier defines a wallet identifier
type WalletIdentifier struct {
	ChainId uint64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *WalletIdentifier) Reset()         { *m = WalletIdentifier{} }
func (m *WalletIdentifier) String() string { return proto.CompactTextString(m) }
func (*WalletIdentifier) ProtoMessage()    {}
func (*WalletIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_279efc507d29e268, []int{1}
}
func (m *WalletIdentifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WalletIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WalletIdentifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WalletIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletIdentifier.Merge(m, src)
}
func (m *WalletIdentifier) XXX_Size() int {
	return m.Size()
}
func (m *WalletIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_WalletIdentifier proto.InternalMessageInfo

func (m *WalletIdentifier) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *WalletIdentifier) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterEnum("core.identity.UserIdentifierType", UserIdentifierType_name, UserIdentifierType_value)
	proto.RegisterType((*UserIdentifier)(nil), "core.identity.UserIdentifier")
	proto.RegisterType((*WalletIdentifier)(nil), "core.identity.WalletIdentifier")
}

func init() { proto.RegisterFile("core/identity/identifier.proto", fileDescriptor_279efc507d29e268) }

var fileDescriptor_279efc507d29e268 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0x4a,
	0xd5, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x84, 0x32, 0xd2, 0x32, 0x53, 0x8b, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x41, 0xf2, 0x7a, 0x30, 0x79, 0xa5, 0x74, 0x2e, 0xbe,
	0xd0, 0xe2, 0xd4, 0x22, 0x4f, 0xb8, 0x32, 0x21, 0x53, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x45, 0x3d, 0x14, 0xf5, 0x7a, 0xa8, 0x8a, 0x43, 0x2a, 0x0b,
	0x52, 0x83, 0xc0, 0xca, 0x85, 0xe4, 0xb8, 0xb8, 0x10, 0x76, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x06, 0x21, 0x89, 0x28, 0xb9, 0x73, 0x09, 0x84, 0x27, 0xe6, 0xe4, 0xa4, 0x96, 0x20, 0x59, 0x25,
	0xc9, 0xc5, 0x91, 0x9c, 0x91, 0x98, 0x99, 0x17, 0x9f, 0x99, 0x02, 0xb6, 0x8e, 0x25, 0x88, 0x1d,
	0xcc, 0xf7, 0x4c, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x86, 0x9a,
	0x05, 0xe3, 0x6a, 0x85, 0x70, 0x09, 0x61, 0x3a, 0x42, 0x88, 0x83, 0x8b, 0xc5, 0xcf, 0xdf, 0xcf,
	0x55, 0x80, 0x41, 0x88, 0x93, 0x8b, 0xd5, 0xd5, 0xd7, 0xd1, 0xd3, 0x47, 0x80, 0x11, 0xc4, 0x0c,
	0xf0, 0x00, 0x89, 0x32, 0x81, 0xe4, 0xdd, 0x3c, 0x5d, 0xfc, 0x05, 0x98, 0x85, 0xb8, 0xb9, 0xd8,
	0x03, 0x1c, 0x83, 0x83, 0xbd, 0x5d, 0x23, 0x05, 0x58, 0x84, 0xd8, 0xb9, 0x98, 0xdd, 0x03, 0xdc,
	0x05, 0x58, 0x9d, 0x9c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x23,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x38, 0x3f, 0xaf, 0x48, 0x37,
	0x33, 0x5f, 0x1f, 0x1c, 0xc6, 0x15, 0x88, 0x50, 0x06, 0x85, 0x40, 0x71, 0x12, 0x1b, 0x38, 0x84,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x85, 0xb2, 0x39, 0x83, 0x01, 0x00, 0x00,
}

func (m *UserIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserIdentifier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserIdentifier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintIdentifier(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintIdentifier(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *WalletIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WalletIdentifier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WalletIdentifier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintIdentifier(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainId != 0 {
		i = encodeVarintIdentifier(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIdentifier(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdentifier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserIdentifier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovIdentifier(uint64(m.Type))
	}
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovIdentifier(uint64(l))
	}
	return n
}

func (m *WalletIdentifier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovIdentifier(uint64(m.ChainId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovIdentifier(uint64(l))
	}
	return n
}

func sovIdentifier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdentifier(x uint64) (n int) {
	return sovIdentifier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentifier
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
			return fmt.Errorf("proto: UserIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= UserIdentifierType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifier
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
				return ErrInvalidLengthIdentifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentifier
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
func (m *WalletIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentifier
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
			return fmt.Errorf("proto: WalletIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WalletIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifier
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
				return ErrInvalidLengthIdentifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentifier
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
func skipIdentifier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentifier
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
					return 0, ErrIntOverflowIdentifier
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
					return 0, ErrIntOverflowIdentifier
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
				return 0, ErrInvalidLengthIdentifier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdentifier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdentifier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdentifier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentifier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdentifier = fmt.Errorf("proto: unexpected end of group")
)
