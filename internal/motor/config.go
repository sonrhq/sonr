package motor

import (
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
)

func setupDefault() (*MotorNode, error) {
	// Create Client instance
	c := client.NewClient(true)

	// Generate wallet
	w, err := crypto.Generate()
	if err != nil {
		return nil, err
	}
	bechAddr, err := w.Address()
	if err != nil {
		return nil, err
	}
	err = c.RequestFaucet(bechAddr)
	if err != nil {
		return nil, err
	}

	pk, err := w.PublicKeyProto()
	if err != nil {
		return nil, err
	}

	doc, err := w.DIDDocument()
	if err != nil {
		return nil, err
	}

	acc, err := c.QueryAccount(bechAddr)
	if err != nil {
		return nil, err
	}

	return &MotorNode{
		Cosmos:  c,
		Wallet:  w,
		Address: bechAddr,
		PubKey:  pk,
		DIDDoc:  doc,
		Account: acc,
	}, nil
}