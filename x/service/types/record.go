// Utility functions for DID Service - https://w3c.github.io/did-core/#services
// I.e. Service Endpoints for IPFS Cluster
package types

import (
	"encoding/base64"
	"encoding/json"
	fmt "fmt"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	idtypes "github.com/sonrhq/core/x/identity/types"
)

const (
	VaultServiceType            = "EncryptedVault"
	LinkedDomainServiceType     = "LinkedDomains"
	DIDCommMessagingServiceType = "DIDCommMessaging"
)

// NewIPFSStoreService creates a new IPFS Store Service record for the given address.
// Addresses look like: /orbitdb/bafyreiepksmvjzvcbzsdqkf474hgfoqf3xj5t47olga5qnnhxggxssbcya/testKVStore
// The address is split into the CID and the DBName, and the CID is used to create the DID. Which results in:
// did:orbitdb:bafyreiepksmvjzvcbzsdqkf474hgfoqf3xj5t47olga5qnnhxggxssbcya
// The origin is the dbname and the type is "EncryptedVault"
func NewIPFSStoreService(address string, controllerDid string) *ServiceRecord {
	parts := strings.Split(address, "/")
	if len(parts) < 4 {
		return nil
	}
	host := parts[1]
	cid := parts[2]
	dbName := parts[3]
	id := fmt.Sprintf("did:%s:%s", host, cid)
	return &ServiceRecord{
		Id:         id,
		Type:       VaultServiceType,
		Origin:     dbName,
		Controller: controllerDid,
	}
}

func (s *ServiceRecord) CredentialEntity() protocol.CredentialEntity {
	return protocol.CredentialEntity{
		Name: s.Name,
	}
}

func (s *ServiceRecord) GetUserEntity(id string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:               []byte(id),
		DisplayName:      id,
		CredentialEntity: s.CredentialEntity(),
	}
}

// GetCredentialCreationOptions issues a challenge for the VerificationMethod to sign and return
func (vm *ServiceRecord) GetCredentialCreationOptions(username string, chal protocol.URLEncodedBase64, addr string, isMobile bool) (string, error) {
	params := DefaultParams()
	cco, err := params.NewWebauthnCreationOptions(vm, username, chal, addr, isMobile)
	if err != nil {
		return "", err
	}

	ccoJSON, err := json.Marshal(cco)
	if err != nil {
		return "", err
	}
	return string(ccoJSON), nil
}

// GetCredentialCreationOptions issues a challenge for the VerificationMethod to sign and return
func (vm *ServiceRecord) GetCredentialAssertionOptions(allowedCredentials []protocol.CredentialDescriptor, isMobile bool) (string, error) {
	hashString := base64.URLEncoding.EncodeToString([]byte(vm.Id))
	params := DefaultParams()
	chal := protocol.URLEncodedBase64(hashString)

	cco, err := params.NewWebauthnAssertionOptions(vm, chal, allowedCredentials, isMobile)
	if err != nil {
		return "", err
	}
	ccoJSON, err := json.Marshal(cco)
	if err != nil {
		return "", err
	}
	return string(ccoJSON), nil
}

// RelyingPartyEntity is a struct that represents a Relying Party entity.
func (s *ServiceRecord) RelyingPartyEntity() protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		ID: s.Id,
		CredentialEntity: protocol.CredentialEntity{
			Name: s.Origin,
		},
	}
}

// VerifyCreationChallenge verifies the challenge and a creation signature and returns an error if it fails to verify
func (vm *ServiceRecord) VerifyCreationChallenge(resp string, chal protocol.URLEncodedBase64) (*WebauthnCredential, error) {
	// Get Credential Creation Respons
	var ccr protocol.CredentialCreationResponse
	err := json.Unmarshal([]byte(resp), &ccr)
	if err != nil {
		return nil, err
	}
	pcc, err := ccr.Parse()
	if err != nil {
		return nil, err
	}
	return makeCredentialFromCreationData(pcc), nil
}

// VeriifyAssertionChallenge verifies the challenge and an assertion signature and returns an error if it fails to verify
func (vm *ServiceRecord) VerifyAssertionChallenge(resp string, creds ...*idtypes.VerificationMethod) error {
	var ccr protocol.CredentialAssertionResponse
	err := json.Unmarshal([]byte(resp), &ccr)
	if err != nil {
		return err
	}
	pca, err := ccr.Parse()
	if err != nil {
		return err
	}
	makeCredentialFromAssertionData(pca)
	return nil
}