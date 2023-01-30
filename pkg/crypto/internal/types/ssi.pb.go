// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/crypto/ssi.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

// KeyType is the type of key used to sign a DID document.
type KeyType int32

const (
	// No key type specified
	KeyType_KeyType_UNSPECIFIED KeyType = 0
	// JsonWebKey2020 is a VerificationMethod type. https://w3c-ccg.github.io/lds-jws2020/
	KeyType_KeyType_JSON_WEB_KEY_2020 KeyType = 1
	// ED25519VerificationKey2018 is the Ed25519VerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-ed25519-2018/
	KeyType_KeyType_ED25519_VERIFICATION_KEY_2018 KeyType = 2
	// ECDSASECP256K1VerificationKey2019 is the EcdsaSecp256k1VerificationKey2019 verification key type as specified here: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
	KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019 KeyType = 3
	// RSAVerificationKey2018 is the RsaVerificationKey2018 verification key type as specified here: https://w3c-ccg.github.io/lds-rsa2018/
	KeyType_KeyType_RSA_VERIFICATION_KEY_2018 KeyType = 4
	// WebAuthnAuthentication2018 is the WebAuthnAuthentication2018 verification key type as specified here: https://w3c-ccg.github.io/lds-webauthn/
	KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 KeyType = 5
)

var KeyType_name = map[int32]string{
	0: "KeyType_UNSPECIFIED",
	1: "KeyType_JSON_WEB_KEY_2020",
	2: "KeyType_ED25519_VERIFICATION_KEY_2018",
	3: "KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019",
	4: "KeyType_RSA_VERIFICATION_KEY_2018",
	5: "KeyType_WEB_AUTHN_AUTHENTICATION_2018",
}

var KeyType_value = map[string]int32{
	"KeyType_UNSPECIFIED":                           0,
	"KeyType_JSON_WEB_KEY_2020":                     1,
	"KeyType_ED25519_VERIFICATION_KEY_2018":         2,
	"KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019": 3,
	"KeyType_RSA_VERIFICATION_KEY_2018":             4,
	"KeyType_WEB_AUTHN_AUTHENTICATION_2018":         5,
}

func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}

func (KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4d65a10f7585d94, []int{0}
}

// ProofType is the type of proof used to present claims over a DID document.
type ProofType int32

const (
	// No proof type specified
	ProofType_ProofType_UNSPECIFIED ProofType = 0
	// JsonWebSignature2020 is a proof type. https://w3c-ccg.github.io/lds-jws2020/
	ProofType_ProofType_JSON_WEB_SIGNATURE_2020 ProofType = 1
	// ED25519Signature2018 is the Ed25519Signature2018 proof type as specified here: https://w3c-ccg.github.io/lds-ed25519-2018/
	ProofType_ProofType_ED25519_SIGNATURE_2018 ProofType = 2
	// EcdsaSecp256k1Signature2019 is the EcdsaSecp256k1Signature2019 proof type as specified here: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
	ProofType_ProofType_ECDSA_SECP256K1_SIGNATURE_2019 ProofType = 3
	// RsaSignature2018 is the RsaSignature2018 proof type as specified here: https://w3c-ccg.github.io/lds-rsa2018/
	ProofType_ProofType_RSA_SIGNATURE_2018 ProofType = 4
)

var ProofType_name = map[int32]string{
	0: "ProofType_UNSPECIFIED",
	1: "ProofType_JSON_WEB_SIGNATURE_2020",
	2: "ProofType_ED25519_SIGNATURE_2018",
	3: "ProofType_ECDSA_SECP256K1_SIGNATURE_2019",
	4: "ProofType_RSA_SIGNATURE_2018",
}

var ProofType_value = map[string]int32{
	"ProofType_UNSPECIFIED":                    0,
	"ProofType_JSON_WEB_SIGNATURE_2020":        1,
	"ProofType_ED25519_SIGNATURE_2018":         2,
	"ProofType_ECDSA_SECP256K1_SIGNATURE_2019": 3,
	"ProofType_RSA_SIGNATURE_2018":             4,
}

func (x ProofType) String() string {
	return proto.EnumName(ProofType_name, int32(x))
}

func (ProofType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4d65a10f7585d94, []int{1}
}

// ServiceType is the type of service used to provide a DID document.
type ServiceType int32

const (
	// No service type specified
	ServiceType_ServiceType_UNSPECIFIED ServiceType = 0
	// DIDCommMessaging is the DIDCommMessaging service type as specified here: https://identity.foundation/didcomm-messaging/spec/
	ServiceType_ServiceType_DID_COMM_MESSAGING ServiceType = 1
	// EncryptedDataVault is the EncryptedDataVault service type as specified here: https://identity.foundation/edv-spec/
	ServiceType_ServiceType_ENCRYPTED_DATA_VAULT ServiceType = 2
	// LinkedDomains is the LinkedDomains service type as specified here: https://identity.foundation/linked-domains/
	ServiceType_ServiceType_LINKED_DOMAINS ServiceType = 3
)

var ServiceType_name = map[int32]string{
	0: "ServiceType_UNSPECIFIED",
	1: "ServiceType_DID_COMM_MESSAGING",
	2: "ServiceType_ENCRYPTED_DATA_VAULT",
	3: "ServiceType_LINKED_DOMAINS",
}

var ServiceType_value = map[string]int32{
	"ServiceType_UNSPECIFIED":          0,
	"ServiceType_DID_COMM_MESSAGING":   1,
	"ServiceType_ENCRYPTED_DATA_VAULT": 2,
	"ServiceType_LINKED_DOMAINS":       3,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}

func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4d65a10f7585d94, []int{2}
}

func init() {
	proto.RegisterEnum("sonrhq.common.crypto.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("sonrhq.common.crypto.ProofType", ProofType_name, ProofType_value)
	proto.RegisterEnum("sonrhq.common.crypto.ServiceType", ServiceType_name, ServiceType_value)
}

func init() { proto.RegisterFile("common/crypto/ssi.proto", fileDescriptor_b4d65a10f7585d94) }

var fileDescriptor_b4d65a10f7585d94 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0x6d, 0x80, 0x78, 0xb9, 0x44, 0x01, 0x54, 0x8d, 0x3f, 0xd6, 0x86, 0xa8, 0x80,
	0x0a, 0x9a, 0xb5, 0xa8, 0x88, 0x1d, 0xbd, 0xc4, 0x2b, 0xa6, 0x8b, 0x13, 0xc5, 0xce, 0xd0, 0xb8,
	0x58, 0x5a, 0x14, 0xb6, 0x0a, 0x6d, 0x0e, 0x49, 0x40, 0xca, 0xb7, 0xe0, 0xc0, 0x37, 0xe2, 0xc2,
	0x71, 0x47, 0x8e, 0xa8, 0xfd, 0x04, 0x7c, 0x03, 0x94, 0xa4, 0x69, 0x43, 0x04, 0x97, 0x28, 0xd6,
	0xf3, 0xf3, 0xab, 0xf7, 0x67, 0x3d, 0xd0, 0x0d, 0xd5, 0xc5, 0x85, 0xba, 0x34, 0xc3, 0x24, 0x8f,
	0x33, 0x65, 0xa6, 0xe9, 0x6c, 0x10, 0x27, 0x2a, 0x53, 0xc6, 0x9d, 0x54, 0x5d, 0x26, 0xe7, 0x9f,
	0x06, 0x55, 0x3e, 0xa8, 0xf2, 0xfe, 0x6f, 0x0d, 0x6e, 0x4c, 0xa3, 0x5c, 0xe4, 0x71, 0x64, 0x74,
	0xe1, 0xf6, 0xf2, 0x57, 0x06, 0x8c, 0x7b, 0xc4, 0xa2, 0x87, 0x94, 0xd8, 0x7a, 0xc7, 0x78, 0x08,
	0xdb, 0x75, 0xf0, 0x96, 0xbb, 0x4c, 0xbe, 0x23, 0x07, 0x72, 0x4a, 0x4e, 0xe4, 0x68, 0x6f, 0xb4,
	0xa7, 0x6b, 0xc6, 0x33, 0xe8, 0xd5, 0x31, 0xb1, 0x47, 0xe3, 0xf1, 0x70, 0x5f, 0x1e, 0x13, 0x9f,
	0x1e, 0x52, 0x0b, 0x0b, 0xea, 0xb2, 0x25, 0x3a, 0x7c, 0xad, 0x6f, 0x18, 0x43, 0x78, 0xb1, 0x42,
	0x2d, 0x9b, 0x63, 0xc9, 0x89, 0xe5, 0x8d, 0xc6, 0xaf, 0xa6, 0xc3, 0x7f, 0x5e, 0xd9, 0xd7, 0x37,
	0x8d, 0x1e, 0xec, 0xd6, 0x57, 0x7c, 0x8e, 0xff, 0x33, 0x79, 0xab, 0xb9, 0x44, 0xb1, 0x1e, 0x0e,
	0xc4, 0x1b, 0x56, 0x7e, 0x09, 0x13, 0x35, 0x5e, 0xa2, 0xd7, 0xfa, 0xdf, 0x35, 0xb8, 0xe9, 0x25,
	0x4a, 0x7d, 0x28, 0xad, 0xb7, 0xe1, 0xee, 0xea, 0xd0, 0xf2, 0xee, 0xc1, 0xee, 0x3a, 0x5a, 0x99,
	0x73, 0x3a, 0x61, 0x58, 0x04, 0x3e, 0xa9, 0xfd, 0x1f, 0xc3, 0xce, 0x1a, 0xab, 0x5f, 0xa0, 0x49,
	0x95, 0xea, 0xcf, 0xe1, 0x69, 0x83, 0x6a, 0xc9, 0xff, 0x45, 0x17, 0xd6, 0x3b, 0xf0, 0x60, 0x4d,
	0x17, 0xde, 0xad, 0x79, 0x5b, 0xfd, 0x6f, 0x1a, 0xdc, 0xe2, 0x51, 0xf2, 0x65, 0x16, 0x46, 0xa5,
	0xc7, 0x7d, 0xe8, 0x36, 0x8e, 0x2d, 0x93, 0x47, 0x80, 0x9a, 0xa1, 0x4d, 0x6d, 0x69, 0xb9, 0x8e,
	0x23, 0x1d, 0xc2, 0x39, 0x9e, 0x50, 0x36, 0xa9, 0x34, 0x9a, 0x0c, 0x61, 0x96, 0x7f, 0xe2, 0x09,
	0x62, 0x4b, 0x1b, 0x0b, 0x2c, 0x8f, 0x71, 0x70, 0x24, 0xf4, 0x0d, 0x03, 0xc1, 0xbd, 0x26, 0x75,
	0x44, 0xd9, 0xb4, 0x40, 0x5c, 0x07, 0x53, 0xc6, 0xf5, 0xcd, 0x03, 0xfc, 0x63, 0x8e, 0xb4, 0xab,
	0x39, 0xd2, 0x7e, 0xcd, 0x91, 0xf6, 0x75, 0x81, 0x3a, 0x57, 0x0b, 0xd4, 0xf9, 0xb9, 0x40, 0x9d,
	0xf7, 0x4f, 0xce, 0x66, 0xd9, 0xf9, 0xe7, 0xd3, 0xa2, 0x80, 0x66, 0xd5, 0x45, 0x33, 0x54, 0x49,
	0x64, 0xc6, 0x1f, 0xcf, 0xea, 0xb6, 0x66, 0x79, 0x1c, 0xa5, 0xa7, 0xd7, 0xcb, 0xc2, 0xbe, 0xfc,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x3a, 0x16, 0x20, 0xcb, 0x02, 0x00, 0x00,
}