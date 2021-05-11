package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/example/baseledger/x/trustmesh/types"
)

func (k msgServer) CreateSynchronizationRequest(goCtx context.Context, msg *types.MsgCreateSynchronizationRequest) (*types.MsgCreateSynchronizationRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendSynchronizationRequest(
		ctx,
		msg.Creator,
		msg.WorkgroupID,
		msg.Recipient,
		msg.WorkstepType,
		msg.BusinessObjectType,
		msg.BaseledgerBusinessObjectID,
		msg.BusinessObject,
		msg.ReferencedBaseledgerBusinessObjectID,
	)

	return &types.MsgCreateSynchronizationRequestResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSynchronizationRequest(goCtx context.Context, msg *types.MsgUpdateSynchronizationRequest) (*types.MsgUpdateSynchronizationRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var SynchronizationRequest = types.SynchronizationRequest{
		Creator:                              msg.Creator,
		Id:                                   msg.Id,
		WorkgroupID:                          msg.WorkgroupID,
		Recipient:                            msg.Recipient,
		WorkstepType:                         msg.WorkstepType,
		BusinessObjectType:                   msg.BusinessObjectType,
		BaseledgerBusinessObjectID:           msg.BaseledgerBusinessObjectID,
		BusinessObject:                       msg.BusinessObject,
		ReferencedBaseledgerBusinessObjectID: msg.ReferencedBaseledgerBusinessObjectID,
	}

	// Checks that the element exists
	if !k.HasSynchronizationRequest(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetSynchronizationRequestOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetSynchronizationRequest(ctx, SynchronizationRequest)

	return &types.MsgUpdateSynchronizationRequestResponse{}, nil
}

func (k msgServer) DeleteSynchronizationRequest(goCtx context.Context, msg *types.MsgDeleteSynchronizationRequest) (*types.MsgDeleteSynchronizationRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasSynchronizationRequest(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetSynchronizationRequestOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSynchronizationRequest(ctx, msg.Id)

	return &types.MsgDeleteSynchronizationRequestResponse{}, nil
}
