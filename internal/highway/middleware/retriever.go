package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/sonr/config"
	"github.com/sonr-io/sonr/internal/local"
	domaintypes "github.com/sonr-io/sonr/x/domain/types"
	identitytypes "github.com/sonr-io/sonr/x/identity/types"
	servicetypes "github.com/sonr-io/sonr/x/service/types"
)

var baseAPIUrl = fmt.Sprintf("http://%s", config.NodeAPIHostAddress())

// GetEmailRecordCreator returns the creator of a given email
func GetEmailRecordCreator(email string) (string, error) {
	endpoint := fmt.Sprintf("%s/core/domain/username/%s", baseAPIUrl, domaintypes.EmailIndex(email))
	resp := new(domaintypes.QueryGetUsernameRecordsResponse)
	bz, err := local.GetJSON(endpoint)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return "", err
	}
	return resp.UsernameRecords.Address, nil
}

// GetControllerAccount returns the DIDDocument of a given DID or Alias
func GetControllerAccount(address string) (*identitytypes.ControllerAccount, error) {
	endpoint := fmt.Sprintf("%s/core/id/controllers/%s", baseAPIUrl, address)
	resp := new(identitytypes.QueryGetControllerAccountResponse)
	bz, err := local.GetJSON(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return &resp.ControllerAccount, nil
}

// GetDID returns the DIDDocument of a given DID or Alias
func GetDID(alias string) (*identitytypes.DIDDocument, error) {
	endpoint := fmt.Sprintf("%s/core/id/alias/%s", baseAPIUrl, alias)
	resp := new(identitytypes.QueryDidByAlsoKnownAsResponse)
	bz, err := local.GetJSON(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetServiceRecord returns the ServiceRecord of a given origin
func GetServiceRecord(origin string) (*servicetypes.ServiceRecord, error) {
	endpoint := fmt.Sprintf("%s/core/service/%s", baseAPIUrl, origin)
	resp := new(servicetypes.QueryGetServiceRecordResponse)
	bz, err := local.GetJSON(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return &resp.ServiceRecord, nil
}
