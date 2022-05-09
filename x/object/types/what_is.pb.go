// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: object/v1/what_is.proto

package types

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

type WhatIs struct {
	// DID is the DID of the object
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Object_doc is the object document
	ObjectDoc *ObjectDoc `protobuf:"bytes,2,opt,name=object_doc,json=objectDoc,proto3" json:"object_doc,omitempty"`
	// Creator is the DID of the creator
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	// Timestamp is the time of the last update of the DID Document
	Timestamp int64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// IsActive is the status of the DID Document
	IsActive bool `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (m *WhatIs) Reset()         { *m = WhatIs{} }
func (m *WhatIs) String() string { return proto.CompactTextString(m) }
func (*WhatIs) ProtoMessage()    {}
func (*WhatIs) Descriptor() ([]byte, []int) {
	return fileDescriptor_98619fb04ad61e9d, []int{0}
}
func (m *WhatIs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhatIs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhatIs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhatIs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhatIs.Merge(m, src)
}
func (m *WhatIs) XXX_Size() int {
	return m.Size()
}
func (m *WhatIs) XXX_DiscardUnknown() {
	xxx_messageInfo_WhatIs.DiscardUnknown(m)
}

var xxx_messageInfo_WhatIs proto.InternalMessageInfo

func (m *WhatIs) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *WhatIs) GetObjectDoc() *ObjectDoc {
	if m != nil {
		return m.ObjectDoc
	}
	return nil
}

func (m *WhatIs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *WhatIs) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *WhatIs) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func init() {
	proto.RegisterType((*WhatIs)(nil), "sonrio.sonr.object.WhatIs")
}

func init() { proto.RegisterFile("object/v1/what_is.proto", fileDescriptor_98619fb04ad61e9d) }

var fileDescriptor_98619fb04ad61e9d = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x4f, 0xca, 0x4a,
	0x4d, 0x2e, 0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0xcf, 0x48, 0x2c, 0x89, 0xcf, 0x2c, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0xce, 0xcf, 0x2b, 0xca, 0xcc, 0xd7, 0x03, 0x51, 0x7a, 0x10,
	0x45, 0x52, 0x62, 0x08, 0xc5, 0x10, 0x16, 0x44, 0xad, 0xd2, 0x5a, 0x46, 0x2e, 0xb6, 0xf0, 0x8c,
	0xc4, 0x12, 0xcf, 0x62, 0x21, 0x01, 0x2e, 0xe6, 0x94, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x10, 0x53, 0xc8, 0x86, 0x8b, 0x0b, 0xa2, 0x38, 0x3e, 0x25, 0x3f, 0x59, 0x82, 0x49,
	0x81, 0x51, 0x83, 0xdb, 0x48, 0x56, 0x0f, 0xd3, 0x74, 0x3d, 0x7f, 0x30, 0xe5, 0x92, 0x9f, 0x1c,
	0xc4, 0x99, 0x0f, 0x63, 0x0a, 0x49, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49,
	0x30, 0x83, 0xcd, 0x84, 0x71, 0x85, 0x64, 0xb8, 0x38, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12,
	0x73, 0x0b, 0x24, 0x58, 0x14, 0x18, 0x35, 0x98, 0x83, 0x10, 0x02, 0x42, 0xd2, 0x5c, 0x9c, 0x99,
	0xc5, 0xf1, 0x89, 0xc9, 0x25, 0x99, 0x65, 0xa9, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x1c,
	0x99, 0xc5, 0x8e, 0x60, 0xbe, 0x93, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x44, 0xa9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0xdc, 0xa6,
	0x9b, 0x99, 0x0f, 0xa6, 0xf5, 0x2b, 0xa0, 0x3e, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0x7b, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x20, 0x49, 0x6f, 0xef, 0x3f, 0x01, 0x00,
	0x00,
}

func (m *WhatIs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhatIs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhatIs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Timestamp != 0 {
		i = encodeVarintWhatIs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhatIs(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ObjectDoc != nil {
		{
			size, err := m.ObjectDoc.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWhatIs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintWhatIs(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhatIs(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhatIs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhatIs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovWhatIs(uint64(l))
	}
	if m.ObjectDoc != nil {
		l = m.ObjectDoc.Size()
		n += 1 + l + sovWhatIs(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhatIs(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovWhatIs(uint64(m.Timestamp))
	}
	if m.IsActive {
		n += 2
	}
	return n
}

func sovWhatIs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhatIs(x uint64) (n int) {
	return sovWhatIs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhatIs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhatIs
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
			return fmt.Errorf("proto: WhatIs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhatIs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
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
				return ErrInvalidLengthWhatIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhatIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectDoc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWhatIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhatIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ObjectDoc == nil {
				m.ObjectDoc = &ObjectDoc{}
			}
			if err := m.ObjectDoc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
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
				return ErrInvalidLengthWhatIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhatIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWhatIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhatIs
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
func skipWhatIs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhatIs
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
					return 0, ErrIntOverflowWhatIs
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
					return 0, ErrIntOverflowWhatIs
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
				return 0, ErrInvalidLengthWhatIs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhatIs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhatIs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhatIs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhatIs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhatIs = fmt.Errorf("proto: unexpected end of group")
)