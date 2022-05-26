package jwt

import (
	"crypto/ed25519"
	cryptrand "crypto/rand"
	"fmt"
	"strings"
	"testing"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func Test_JWT(t *testing.T) {
	root := "123456"

	didController, err := did.ParseDID(fmt.Sprintf("did:snr:%s#test", root))
	if err != nil {
		t.Errorf("Error while creating test controller %s", err)
	}

	didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", root))
	if err != nil {
		t.Errorf("Error while creating did url %s", err)
	}

	priv := secp256k1.GenPrivKey()
	pub := priv.PubKey()
	vm, _ := did.NewVerificationMethod(
		*didUrl,
		ssi.ECDSASECP256K1VerificationKey2019,
		*didController,
		pub,
	)

	thing := func() *did.Document {
		doc, _ := CreateMockDidDocument(root)
		doc.AddController(*didController)
		doc.AddAuthenticationMethod(vm)
		return &doc
	}

	t.Run("JWT creation should contain options", func(t *testing.T) {
		jwt := DefaultNew()
		assert.NotNil(t, jwt)
		assert.NotNil(t, jwt.options)
	})

	t.Run("Should generate JWT from did uri", func(t *testing.T) {
		doc := thing()
		jwt := DefaultNew()
		token, err := jwt.Generate(*doc)
		if err != nil {
			t.Errorf("Error while generating token %s", err)
		}

		tokenObj, error := jwt.Parse(token)

		if error != nil {
			t.Errorf("Error while generating token %s", err)
		}

		assert.NotNil(t, tokenObj)
		assert.Greater(t, len(token), 0)
	})

	t.Run("Should create token from string and token should be non nil", func(t *testing.T) {
		doc := thing()
		jwt := DefaultNew()
		token, err := jwt.Generate(*doc)
		if err != nil {
			t.Errorf("Error while generating token %s", err)
		}

		tokenObj, error := jwt.Parse(token)
		if error != nil {
			t.Errorf("Error while generating token %s", err)
		}

		// add did check when figure out claims obj ???
		// weird map that isnt a map in the struct
		assert.NotNil(t, tokenObj)
	})

	t.Run("Should create token from string and token should be valid", func(t *testing.T) {
		doc := thing()
		jwt := DefaultNew()
		token, err := jwt.Generate(*doc)
		if err != nil {
			t.Errorf("Error while generating token %s", err)
		}

		tokenObj, error := jwt.Parse(token)
		if error != nil {
			t.Errorf("Error while generating token %s", err)
		}

		// add did check when figure out claims obj ???
		// weird map that isnt a map in the struct
		assert.Equal(t, tokenObj.Valid, true)
	})

	t.Run("Parse claims from token should include did", func(t *testing.T) {
		doc := thing()
		jwt := DefaultNew()
		token, err := jwt.Generate(*doc)
		if err != nil {
			t.Errorf("Error while generating token %s", err)
		}

		tokenObj, error := jwt.Parse(token)
		if error != nil {
			t.Errorf("Error while generating token %s", err)
		}
		claims, ok := GetClaims(tokenObj)
		if !ok {
			t.Error("Unable to parse claims from token")
		}

		assert.NotNil(t, claims.Issuer)
		assert.NotNil(t, claims.ExpiresAt)
		assert.NotNil(t, claims.IssuedAt)
	})
}

// CreateMockDidDocument creates a mock did document for testing
func CreateMockDidDocument(acct string) (did.Document, error) {
	rawCreator := acct

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		rawCreator = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(rawCreator, "cosmos") {
		rawCreator = strings.TrimLeft(rawCreator, "cosmos")
	}

	// UnmarshalJSON from DID document
	doc, err := did.NewDocument(fmt.Sprintf("did:snr:%s", rawCreator))
	if err != nil {
		return nil, err
	}

	//webauthncred := CreateMockCredential()
	pubKey, _, err := ed25519.GenerateKey(cryptrand.Reader)
	if err != nil {
		return nil, err
	}

	didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", rawCreator))
	if err != nil {
		return nil, err
	}
	didController, err := did.ParseDID(fmt.Sprintf("did:snr:%s#test", rawCreator))
	if err != nil {
		return nil, err
	}

	vm, err := did.NewVerificationMethod(*didUrl, ssi.JsonWebKey2020, *didController, pubKey)
	if err != nil {
		return nil, err
	}
	doc.AddAuthenticationMethod(vm)
	return doc, nil
}