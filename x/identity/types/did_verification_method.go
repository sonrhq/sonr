package types

import (
	"encoding/base64"
	"encoding/json"
	fmt "fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/pkg/crypto"
)

var (
	knownAddrPrefixes = []string{
		"snr",
		"btc",
		"0x",
		"cosmos",
		"fil",
	}
)

// VerificationMethodOption is used to define options that modify the creation of the verification method
type VerificationMethodOption func(vm *VerificationMethod) error

// WithController sets the controller of a verificationMethod
func WithController(v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		_, err := ParseDID(v)
		if err != nil {
			return err
		}
		vm.Controller = v
		return nil
	}
}

// WithIDFragmentSuffix sets the fragment of the ID on a verificationMethod
func WithIDFragmentSuffix(v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		vm.Id = fmt.Sprintf("%s#%s", vm.Id, v)
		return nil
	}
}

// WithBlockchainAccount sets the blockchain account of a verificationMethod
func WithBlockchainAccount(v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		vm.BlockchainAccountId = v
		return nil
	}
}

// WithMetadataValue sets the metadata value of a verificationMethod
func WithMetadataValue(k, v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		vm.SetMetadataValue(k, v)
		return nil
	}
}

//
// VerificationMethod Creation Functions
//

// // VerificationMethod applies the given options and builds a verification method from this Key
func NewVMFromPubKey(pk *crypto.PubKey, opts ...VerificationMethodOption) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		Id:                 pk.DID(),
		Type:               pk.KeyType,
		PublicKeyMultibase: pk.Multibase(),
	}
	for _, opt := range opts {
		if err := opt(vm); err != nil {
			return nil, err
		}
	}
	return vm, nil
}

// NewVerificationMethod is a convenience method to easily create verificationMethods based on a set of given params.
// It automatically encodes the provided public key based on the keyType.
func NewVerificationMethod(id string, keyType crypto.KeyType, controller string, key interface{}) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		Id:         id,
		Type:       keyType.PrettyString(),
		Controller: controller,
	}
	// Check for Secp256k1 key
	if keyType == crypto.Secp256k1KeyType {
		// Switch Interface to *secp256k1.PublicKey or string
		switch key.(type) {
		case *secp256k1.PubKey:
			vm.BlockchainAccountId = key.(*secp256k1.PubKey).Address().String()
		case string:
			vm.BlockchainAccountId = key.(string)
		default:
			return nil, fmt.Errorf("key is not a secp256k1.PublicKey or string")
		}
	}
	return vm, nil
}

// PubKey returns the public key of the verification method
func (v *VerificationMethod) PubKey() (*crypto.PubKey, error) {
	return crypto.PubKeyFromDID(v.Id)
}

// UnmarshalJSON implements the json.Unmarshaler interface for the VerificationMethod type.
func (v *VerificationMethod) UnmarshalJSON(bytes []byte) error {
	type Alias VerificationMethod
	tmp := Alias{}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	*v = (VerificationMethod)(tmp)
	return nil
}

// CredentialDiscriptor is a descriptor for a credential for VerificationMethod which contains WebAuthnCredential
func (vm *VerificationMethod) CredentialDescriptor() (protocol.CredentialDescriptor, error) {
	if vm.Type != crypto.WebAuthnKeyType.PrettyString() {
		return protocol.CredentialDescriptor{}, fmt.Errorf("verification method is not of type WebAuthn")
	}
	cred, err := vm.WebAuthnCredential()
	if err != nil {
		return protocol.CredentialDescriptor{}, err
	}
	stdCred := cred.ToStdCredential()
	return stdCred.Descriptor(), nil
}

// IDFragmentSuffix returns the fragment of the ID of the VerificationMethod
func (vm *VerificationMethod) IDFragmentSuffix() string {
	ptrs := strings.Split(vm.Id, "#")
	return ptrs[len(ptrs)-1]
}

// IssueChallenge issues a challenge for the VerificationMethod to sign and return
func (vm *VerificationMethod) IssueChallenge(unsignedUcanStr string) (protocol.URLEncodedBase64, error) {
	b64Ucan := base64.RawURLEncoding.EncodeToString([]byte(unsignedUcanStr))
	return protocol.URLEncodedBase64(b64Ucan), nil
}

// IsBlockchainAccount returns true if the VerificationMethod is a blockchain account
func (vm *VerificationMethod) IsBlockchainAccount() bool {
	return vm.BlockchainAccountId != ""
}

// PublicKey returns the public key of the VerificationMethod
func (vm *VerificationMethod) PublicKey() ([]byte, error) {
	switch vm.Type {
	case crypto.Ed25519KeyType.PrettyString():
		return base58.Decode(vm.PublicKeyMultibase, base58.BitcoinAlphabet)
	case crypto.Secp256k1KeyType.PrettyString():
		_, bz, err := bech32.DecodeAndConvert(vm.BlockchainAccountId)
		if err != nil {
			return nil, err
		}
		return bz, nil
	default:
		return nil, fmt.Errorf("unsupported key type")
	}
}

func (vm *VerificationMethod) SetMetadata(data map[string]string) {
	vm.Metadata = MapToKeyValueList(data)
}

// SetMetadataValue sets the metadata value for the given key
func (vm *VerificationMethod) SetMetadataValue(key, value string) {
	for i, kv := range vm.Metadata {
		if kv.Key == key {
			vm.Metadata[i].Value = value
			return
		}
	}
	vm.Metadata = append(vm.Metadata, &KeyValuePair{Key: key, Value: value})
}

// GetMetadata returns the metadata value for the given key
func (vm *VerificationMethod) GetMetadataValue(key string) (string, bool) {
	ok := vm.HasMetadataValue(key)
	if !ok {
		return "", false
	}
	for _, kv := range vm.Metadata {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return "", false
}

// HasMetadata returns true if the VerificationMethod has the given metadata key
func (vm *VerificationMethod) HasMetadataValue(key string) bool {
	for _, kv := range vm.Metadata {
		if kv.Key == key {
			return true
		}
	}
	return false
}