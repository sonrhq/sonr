package types

import (
	"encoding/base64"
	fmt "fmt"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/shengdoushi/base58"
	"github.com/sonr-io/sonr/pkg/crypto"
)

// PublicKeyMultibase returns the public key in multibase format
func (c *WebauthnCredential) PubKey() *crypto.PubKey {
	return crypto.NewEd25519PubKey(c.PublicKey)
}

// CredentialFromDIDString converts a DID string into a WebauthnCredential
func CredentialFromDIDString(did string) (*WebauthnCredential, error) {
	parts := strings.Split(did, "#")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid DID string format")
	}

	multibaseKey := parts[0][8:]
	credIdBz, err := base58.Decode(parts[1], base58.BitcoinAlphabet)
	if err != nil {
		return nil, fmt.Errorf("failed to decode device label: %v", err)
	}

	if !strings.HasPrefix(multibaseKey, "z") {
		return nil, fmt.Errorf("invalid multibase prefix")
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(multibaseKey[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %v", err)
	}
	return &WebauthnCredential{PublicKey: pubKeyBytes, Id: credIdBz}, nil
}

// PublicKeyCredentialRequestOptions is a struct that contains the options for a PublicKeyCredentialRequest
// This is a modified version of the struct from the webauthn package to allow for the Attestation field
type PublicKeyCredentialRequestOptions struct {
	Challenge          protocol.URLEncodedBase64            `json:"challenge"`
	Timeout            int                                  `json:"timeout,omitempty"`
	RelyingPartyID     string                               `json:"rpId,omitempty"`
	AllowedCredentials []protocol.CredentialDescriptor      `json:"allowCredentials,omitempty"`
	UserVerification   protocol.UserVerificationRequirement `json:"userVerification,omitempty"`
	Extensions         protocol.AuthenticationExtensions    `json:"extensions,omitempty"`
	Attestion          string                               `json:"attestation,omitempty"`
	AttestionFormats   []string                             `json:"attestationFormats,omitempty"`
}
