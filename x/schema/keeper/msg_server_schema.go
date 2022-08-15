package keeper

import (
	"context"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/x/schema/types"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if len(msg.Definition.Fields) < 1 {
		k.Logger(ctx).Info("Create schema request empty, must have defined fields. aborting operation")
		return &types.MsgCreateSchemaResponse{
			Code:    http.StatusBadRequest,
			Message: "Schema Fields must non nil",
			WhatIs:  nil,
		}, nil
	}

	k.Logger(ctx).Info("msg validation successful")

	accts := msg.GetSigners()
	if len(accts) < 1 {
		return nil, sdkerrors.ErrNotFound
	}

	creator_did := msg.GetCreatorDid()
	k.Logger(ctx).Info(fmt.Sprintf("Creating schema for creator did %s", creator_did))

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	b, err := msg.Definition.Marshal()

	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Error while pinning schema definition to storage")
	}

	cid_str, err := k.PinContent(b)
	k.Logger(ctx).Info(fmt.Sprintf("Schema persisted with cid %s", cid_str))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "Error while persisting schema fields")
	}

	what_is_did, err := did.ParseDID(fmt.Sprintf("did:snr:%s", cid_str))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "error while creating did from cid")
	}
	k.Logger(ctx).Info(fmt.Sprintf("Creating schema with did %s", what_is_did))

	var schema = types.SchemaReference{
		Label: msg.Definition.Label,
		Did:   what_is_did.String(),
		Cid:   cid_str,
	}

	var whatIs = types.WhatIs{
		Creator:   creator_did,
		Did:       what_is_did.String(),
		Schema:    &schema,
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}

	k.SetWhatIs(ctx, whatIs)

	resp := types.MsgCreateSchemaResponse{
		Code:    http.StatusAccepted,
		Message: "Schema Registered Sucessfully",
		WhatIs:  &whatIs,
	}

	return &resp, nil
}

func (k msgServer) DeprecateSchema(goCtx context.Context, msg *types.MsgDeprecateSchema) (*types.MsgDeprecateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	schemas, found := k.GetWhatIsFromCreator(ctx, msg.GetCreator())
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No Schemas found under same creator as message creator.")
	}

	var what_is *types.WhatIs
	var foundSchemaWI bool
	for _, a := range schemas {
		if a.GetDid() == msg.GetDid() {
			what_is = a
			foundSchemaWI = true
			break
		}
	}

	if !foundSchemaWI {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No Schema with same creator as message creator found.")
	}

	//If already deactivated, do nothing.
	//Responsibility of caller to check if isActive beforehand
	if what_is.GetIsActive() {
		what_is.IsActive = false
		k.SetWhatIs(ctx, *what_is)
	}

	return &types.MsgDeprecateSchemaResponse{
		Code:    200,
		Message: "Schema deprecated successfully.",
	}, nil
}
