// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/webauthn/credential.proto

package webauthn

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

// Credential is a message type that contains all needed information
// about a WebAuthn credential for storage.
type Credential struct {
	// id is a probabilistically-unique byte sequence identifying a public key
	// credential source and its authentication assertions.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// public_key is the public key portion of a Relying Party-specific credential
	// key pair, generated by an authenticator and returned to a Relying Party at
	// registration time.
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// attestation_type is the attestation format used (if any) by the
	// authenticator when creating the credential.
	AttestationType string `protobuf:"bytes,3,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	// transport is the transports used by the authenticator when creating the
	// credential.
	Transport []string `protobuf:"bytes,4,rep,name=transport,proto3" json:"transport,omitempty"`
	// Authenticator is the Authenticator information for a given certificate.
	Authenticator *Authenticator `protobuf:"bytes,5,opt,name=authenticator,proto3" json:"authenticator,omitempty"`
	// controller is the DID Controller of the credential.
	Controller string `protobuf:"bytes,6,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_9642ee7a70a978ce, []int{0}
}
func (m *Credential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return m.Size()
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Credential) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Credential) GetAttestationType() string {
	if m != nil {
		return m.AttestationType
	}
	return ""
}

func (m *Credential) GetTransport() []string {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *Credential) GetAuthenticator() *Authenticator {
	if m != nil {
		return m.Authenticator
	}
	return nil
}

func (m *Credential) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

// Authenticator is a message type that contains certificate information
// about a WebAuthn authenticator.
type Authenticator struct {
	// aaguid is the AAGUID of the authenticator. An AAGUID is defined as an array
	// containing the globally unique identifier of the authenticator model being
	// sought.
	Aaguid []byte `protobuf:"bytes,1,opt,name=aaguid,proto3" json:"aaguid,omitempty"`
	// sign_count is the SignCount -Upon a new login operation, the Relying Party
	// compares the stored signature counter value with the new signCount value
	// returned in the assertion’s authenticator data.
	SignCount uint32 `protobuf:"varint,2,opt,name=sign_count,json=signCount,proto3" json:"sign_count,omitempty"`
	// attachment is a signal that the authenticator may be cloned, i.e. at
	// least two copies of the credential private key may exist and are being used
	// in parallel.
	Attachment string `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (m *Authenticator) Reset()         { *m = Authenticator{} }
func (m *Authenticator) String() string { return proto.CompactTextString(m) }
func (*Authenticator) ProtoMessage()    {}
func (*Authenticator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9642ee7a70a978ce, []int{1}
}
func (m *Authenticator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Authenticator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Authenticator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Authenticator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authenticator.Merge(m, src)
}
func (m *Authenticator) XXX_Size() int {
	return m.Size()
}
func (m *Authenticator) XXX_DiscardUnknown() {
	xxx_messageInfo_Authenticator.DiscardUnknown(m)
}

var xxx_messageInfo_Authenticator proto.InternalMessageInfo

func (m *Authenticator) GetAaguid() []byte {
	if m != nil {
		return m.Aaguid
	}
	return nil
}

func (m *Authenticator) GetSignCount() uint32 {
	if m != nil {
		return m.SignCount
	}
	return 0
}

func (m *Authenticator) GetAttachment() string {
	if m != nil {
		return m.Attachment
	}
	return ""
}

func init() {
	proto.RegisterType((*Credential)(nil), "common.webauthn.Credential")
	proto.RegisterType((*Authenticator)(nil), "common.webauthn.Authenticator")
}

func init() { proto.RegisterFile("common/webauthn/credential.proto", fileDescriptor_9642ee7a70a978ce) }

var fileDescriptor_9642ee7a70a978ce = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xeb, 0xf6, 0xde, 0x4a, 0x39, 0x50, 0x8a, 0x3c, 0x20, 0x0f, 0x60, 0x45, 0x1d, 0x50,
	0x18, 0x48, 0x24, 0x78, 0x01, 0xa0, 0x6c, 0x6c, 0x11, 0x13, 0x4b, 0xe5, 0xb8, 0xa6, 0xb5, 0x48,
	0xed, 0xc8, 0x39, 0x11, 0xca, 0x5b, 0xf0, 0x58, 0x8c, 0x1d, 0x19, 0x51, 0xbb, 0xf1, 0x14, 0xc8,
	0x69, 0x29, 0x29, 0x93, 0xe5, 0x4f, 0xff, 0x91, 0xff, 0xcf, 0x07, 0x42, 0x69, 0x17, 0x0b, 0x6b,
	0x92, 0x57, 0x95, 0x89, 0x0a, 0xe7, 0x26, 0x91, 0x4e, 0x4d, 0x95, 0x41, 0x2d, 0xf2, 0xb8, 0x70,
	0x16, 0x2d, 0x1d, 0x6e, 0x12, 0xf1, 0x4f, 0x62, 0xf4, 0x45, 0x00, 0xc6, 0xbb, 0x14, 0x3d, 0x82,
	0xae, 0x9e, 0x32, 0x12, 0x92, 0xe8, 0x30, 0xed, 0xea, 0x29, 0x3d, 0x03, 0x28, 0xaa, 0x2c, 0xd7,
	0x72, 0xf2, 0xa2, 0x6a, 0xd6, 0x6d, 0x78, 0xb0, 0x21, 0x0f, 0xaa, 0xa6, 0x17, 0x70, 0x2c, 0x10,
	0x55, 0x89, 0x02, 0xb5, 0x35, 0x13, 0xac, 0x0b, 0xc5, 0x7a, 0x21, 0x89, 0x82, 0x74, 0xd8, 0xe2,
	0x8f, 0x75, 0xa1, 0xe8, 0x29, 0x04, 0xe8, 0x84, 0x29, 0x0b, 0xeb, 0x90, 0xfd, 0x0b, 0x7b, 0x51,
	0x90, 0xfe, 0x02, 0x7a, 0x0f, 0x03, 0xdf, 0xc7, 0xb7, 0x90, 0x02, 0xad, 0x63, 0xff, 0x43, 0x12,
	0x1d, 0x5c, 0xf1, 0xf8, 0x4f, 0xdf, 0xf8, 0xb6, 0x9d, 0x4a, 0xf7, 0x87, 0x28, 0x07, 0x90, 0xd6,
	0xa0, 0xb3, 0x79, 0xae, 0x1c, 0xeb, 0x37, 0x45, 0x5a, 0x64, 0xf4, 0x0c, 0x83, 0xbd, 0x79, 0x7a,
	0x02, 0x7d, 0x21, 0x66, 0xd5, 0x4e, 0x79, 0x7b, 0xf3, 0xda, 0xa5, 0x9e, 0x99, 0x89, 0xb4, 0x95,
	0xc1, 0x46, 0x7b, 0x90, 0x06, 0x9e, 0x8c, 0x3d, 0xf0, 0xef, 0x08, 0x44, 0x21, 0xe7, 0x0b, 0x65,
	0x70, 0x2b, 0xdc, 0x22, 0x77, 0x37, 0xef, 0x2b, 0x4e, 0x96, 0x2b, 0x4e, 0x3e, 0x57, 0x9c, 0xbc,
	0xad, 0x79, 0x67, 0xb9, 0xe6, 0x9d, 0x8f, 0x35, 0xef, 0x3c, 0x9d, 0xcf, 0x34, 0xce, 0xab, 0xcc,
	0x6b, 0x25, 0xa5, 0x35, 0xee, 0x52, 0xdb, 0xe6, 0x4c, 0xfc, 0xd7, 0x95, 0xbb, 0xc5, 0x65, 0xfd,
	0x66, 0x5d, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0x68, 0x24, 0xee, 0xd2, 0x01, 0x00,
	0x00,
}

func (m *Credential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x32
	}
	if m.Authenticator != nil {
		{
			size, err := m.Authenticator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCredential(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Transport) > 0 {
		for iNdEx := len(m.Transport) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transport[iNdEx])
			copy(dAtA[i:], m.Transport[iNdEx])
			i = encodeVarintCredential(dAtA, i, uint64(len(m.Transport[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AttestationType) > 0 {
		i -= len(m.AttestationType)
		copy(dAtA[i:], m.AttestationType)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.AttestationType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Authenticator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Authenticator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Authenticator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attachment) > 0 {
		i -= len(m.Attachment)
		copy(dAtA[i:], m.Attachment)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Attachment)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SignCount != 0 {
		i = encodeVarintCredential(dAtA, i, uint64(m.SignCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Aaguid) > 0 {
		i -= len(m.Aaguid)
		copy(dAtA[i:], m.Aaguid)
		i = encodeVarintCredential(dAtA, i, uint64(len(m.Aaguid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCredential(dAtA []byte, offset int, v uint64) int {
	offset -= sovCredential(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Credential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.AttestationType)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	if len(m.Transport) > 0 {
		for _, s := range m.Transport {
			l = len(s)
			n += 1 + l + sovCredential(uint64(l))
		}
	}
	if m.Authenticator != nil {
		l = m.Authenticator.Size()
		n += 1 + l + sovCredential(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	return n
}

func (m *Authenticator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Aaguid)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	if m.SignCount != 0 {
		n += 1 + sovCredential(uint64(m.SignCount))
	}
	l = len(m.Attachment)
	if l > 0 {
		n += 1 + l + sovCredential(uint64(l))
	}
	return n
}

func sovCredential(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCredential(x uint64) (n int) {
	return sovCredential(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredential
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
			return fmt.Errorf("proto: Credential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transport = append(m.Transport, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authenticator == nil {
				m.Authenticator = &Authenticator{}
			}
			if err := m.Authenticator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredential
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
func (m *Authenticator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredential
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
			return fmt.Errorf("proto: Authenticator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Authenticator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aaguid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aaguid = append(m.Aaguid[:0], dAtA[iNdEx:postIndex]...)
			if m.Aaguid == nil {
				m.Aaguid = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignCount", wireType)
			}
			m.SignCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attachment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredential
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
				return ErrInvalidLengthCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attachment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredential
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
func skipCredential(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredential
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
					return 0, ErrIntOverflowCredential
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
					return 0, ErrIntOverflowCredential
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
				return 0, ErrInvalidLengthCredential
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCredential
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCredential
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCredential        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredential          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCredential = fmt.Errorf("proto: unexpected end of group")
)
