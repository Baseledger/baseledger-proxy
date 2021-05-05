package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/example/baseledger/x/baseledger/types"
)

func (k msgServer) CreateBaseledgerTransaction(goCtx context.Context, msg *types.MsgCreateBaseledgerTransaction) (*types.MsgCreateBaseledgerTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendBaseledgerTransaction(
		ctx,
		msg.Creator,
		msg.Baseid,
		msg.Payload,
	)

	return &types.MsgCreateBaseledgerTransactionResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBaseledgerTransaction(goCtx context.Context, msg *types.MsgUpdateBaseledgerTransaction) (*types.MsgUpdateBaseledgerTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var BaseledgerTransaction = types.BaseledgerTransaction{
		Creator: msg.Creator,
		Id:      msg.Id,
		Baseid:  msg.Baseid,
		Payload: msg.Payload,
	}

	// Checks that the element exists
	if !k.HasBaseledgerTransaction(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBaseledgerTransactionOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBaseledgerTransaction(ctx, BaseledgerTransaction)

	return &types.MsgUpdateBaseledgerTransactionResponse{}, nil
}

func (k msgServer) DeleteBaseledgerTransaction(goCtx context.Context, msg *types.MsgDeleteBaseledgerTransaction) (*types.MsgDeleteBaseledgerTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBaseledgerTransaction(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBaseledgerTransactionOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBaseledgerTransaction(ctx, msg.Id)

	return &types.MsgDeleteBaseledgerTransactionResponse{}, nil
}
