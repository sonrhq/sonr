// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sonr/identity/ssi.proto

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
	return fileDescriptor_7951e27a89149433, []int{0}
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
	return fileDescriptor_7951e27a89149433, []int{1}
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
	return fileDescriptor_7951e27a89149433, []int{2}
}

func init() {
	proto.RegisterEnum("sonrhq.sonr.identity.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("sonrhq.sonr.identity.ProofType", ProofType_name, ProofType_value)
	proto.RegisterEnum("sonrhq.sonr.identity.ServiceType", ServiceType_name, ServiceType_value)
}

func init() { proto.RegisterFile("sonr/identity/ssi.proto", fileDescriptor_7951e27a89149433) }

var fileDescriptor_7951e27a89149433 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0x94, 0x40,
	0x18, 0x87, 0x97, 0xb6, 0x6a, 0x7c, 0xbd, 0x10, 0xd4, 0x6c, 0xea, 0x9f, 0x49, 0x6b, 0xdc, 0xa4,
	0x6e, 0xec, 0x52, 0x30, 0x6b, 0xec, 0x71, 0x16, 0xa6, 0xeb, 0x48, 0x19, 0x08, 0x33, 0xd4, 0xd4,
	0xcb, 0x24, 0xad, 0xe8, 0x72, 0xb0, 0x6c, 0x01, 0x8d, 0x7c, 0x0b, 0x0f, 0x7e, 0x23, 0x2f, 0x1e,
	0x7b, 0xf4, 0x68, 0x76, 0x3f, 0x81, 0xdf, 0xc0, 0x80, 0xcb, 0x82, 0x44, 0x2f, 0x0c, 0x93, 0xdf,
	0xf3, 0xbe, 0x79, 0x9f, 0xc9, 0x0b, 0xfd, 0x2c, 0xb9, 0x48, 0xf5, 0xf8, 0x6d, 0x74, 0x91, 0xc7,
	0x79, 0xa1, 0x67, 0x59, 0x3c, 0x9a, 0xa7, 0x49, 0x9e, 0x68, 0x77, 0xca, 0x60, 0x76, 0x39, 0x2a,
	0x8f, 0x51, 0x9d, 0x0f, 0x7f, 0x29, 0x70, 0xc3, 0x89, 0x0a, 0x51, 0xcc, 0x23, 0xad, 0x0f, 0xb7,
	0x57, 0xbf, 0x32, 0x64, 0xdc, 0x27, 0x16, 0x3d, 0xa2, 0xc4, 0x56, 0x7b, 0xda, 0x43, 0xd8, 0xae,
	0x83, 0x57, 0xdc, 0x63, 0xf2, 0x35, 0x99, 0x48, 0x87, 0x9c, 0x4a, 0xf3, 0xc0, 0x3c, 0x50, 0x15,
	0xed, 0x09, 0x0c, 0xea, 0x98, 0xd8, 0xe6, 0x78, 0x6c, 0x1c, 0xca, 0x13, 0x12, 0xd0, 0x23, 0x6a,
	0x61, 0x41, 0x3d, 0xb6, 0x42, 0x8d, 0x17, 0xea, 0x86, 0x66, 0xc0, 0xfe, 0x1a, 0xb5, 0x6c, 0x8e,
	0x25, 0x27, 0x96, 0x6f, 0x8e, 0x9f, 0x3b, 0xc6, 0x3f, 0x4b, 0x0e, 0xd5, 0x4d, 0x6d, 0x00, 0xbb,
	0x75, 0x49, 0xc0, 0xf1, 0x7f, 0x3a, 0x6f, 0xb5, 0x87, 0x28, 0xc7, 0xc3, 0xa1, 0x78, 0xc9, 0xaa,
	0x2f, 0x61, 0xa2, 0xc6, 0x2b, 0xf4, 0xda, 0xf0, 0x9b, 0x02, 0x37, 0xfd, 0x34, 0x49, 0xde, 0x55,
	0xd6, 0xdb, 0x70, 0x77, 0x7d, 0xe9, 0x78, 0x0f, 0x60, 0xb7, 0x89, 0xd6, 0xe6, 0x9c, 0x4e, 0x19,
	0x16, 0x61, 0x40, 0x6a, 0xff, 0xc7, 0xb0, 0xd3, 0x60, 0xf5, 0x0b, 0xb4, 0xa9, 0x4a, 0xfd, 0x29,
	0xec, 0xb5, 0xa8, 0x8e, 0xfc, 0x5f, 0x74, 0x69, 0xbd, 0x03, 0x0f, 0x1a, 0xba, 0xf4, 0xee, 0xf4,
	0xdb, 0x1a, 0x7e, 0x55, 0xe0, 0x16, 0x8f, 0xd2, 0x4f, 0xf1, 0x79, 0x54, 0x79, 0xdc, 0x87, 0x7e,
	0xeb, 0xda, 0x31, 0x79, 0x04, 0xa8, 0x1d, 0xda, 0xd4, 0x96, 0x96, 0xe7, 0xba, 0xd2, 0x25, 0x9c,
	0xe3, 0x29, 0x65, 0xd3, 0x3f, 0x1a, 0x6d, 0x86, 0x30, 0x2b, 0x38, 0xf5, 0x05, 0xb1, 0xa5, 0x8d,
	0x05, 0x96, 0x27, 0x38, 0x3c, 0x16, 0xea, 0x86, 0x86, 0xe0, 0x5e, 0x9b, 0x3a, 0xa6, 0xcc, 0x29,
	0x11, 0xcf, 0xc5, 0x94, 0x71, 0x75, 0x73, 0x32, 0xf9, 0xbe, 0x40, 0xca, 0xd5, 0x02, 0x29, 0x3f,
	0x17, 0x48, 0xf9, 0xb2, 0x44, 0xbd, 0xab, 0x25, 0xea, 0xfd, 0x58, 0xa2, 0xde, 0x9b, 0xbd, 0xf7,
	0x71, 0x3e, 0xfb, 0x78, 0x36, 0x3a, 0x4f, 0x3e, 0xe8, 0xe5, 0x12, 0xee, 0xcf, 0x2e, 0xab, 0x53,
	0xff, 0xdc, 0xac, 0x6b, 0x5e, 0xcc, 0xa3, 0xec, 0xec, 0x7a, 0xb5, 0xb1, 0xcf, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0x49, 0x47, 0x60, 0xcc, 0x02, 0x00, 0x00,
}
