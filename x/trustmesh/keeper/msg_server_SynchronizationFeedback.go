package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/example/baseledger/x/trustmesh/types"
)

func (k msgServer) CreateSynchronizationFeedback(goCtx context.Context, msg *types.MsgCreateSynchronizationFeedback) (*types.MsgCreateSynchronizationFeedbackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendSynchronizationFeedback(
		ctx,
		msg.Creator,
		msg.Approved,
		msg.BusinessObject,
		msg.BaseledgerBusinessObjectIDOfApprovedObject,
		msg.Workgroup,
		msg.Recipient,
		msg.HashOfObjectToApprove,
		msg.OriginalBaseledgerTransactionID,
		msg.OriginalOffchainProcessMessageID,
		msg.FeedbackMessage,
	)

	return &types.MsgCreateSynchronizationFeedbackResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSynchronizationFeedback(goCtx context.Context, msg *types.MsgUpdateSynchronizationFeedback) (*types.MsgUpdateSynchronizationFeedbackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var SynchronizationFeedback = types.SynchronizationFeedback{
		Creator:        msg.Creator,
		Id:             msg.Id,
		Approved:       msg.Approved,
		BusinessObject: msg.BusinessObject,
		BaseledgerBusinessObjectIDOfApprovedObject: msg.BaseledgerBusinessObjectIDOfApprovedObject,
		Workgroup:                        msg.Workgroup,
		Recipient:                        msg.Recipient,
		HashOfObjectToApprove:            msg.HashOfObjectToApprove,
		OriginalBaseledgerTransactionID:  msg.OriginalBaseledgerTransactionID,
		OriginalOffchainProcessMessageID: msg.OriginalOffchainProcessMessageID,
		FeedbackMessage:                  msg.FeedbackMessage,
	}

	// Checks that the element exists
	if !k.HasSynchronizationFeedback(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetSynchronizationFeedbackOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetSynchronizationFeedback(ctx, SynchronizationFeedback)

	return &types.MsgUpdateSynchronizationFeedbackResponse{}, nil
}

func (k msgServer) DeleteSynchronizationFeedback(goCtx context.Context, msg *types.MsgDeleteSynchronizationFeedback) (*types.MsgDeleteSynchronizationFeedbackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasSynchronizationFeedback(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetSynchronizationFeedbackOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSynchronizationFeedback(ctx, msg.Id)

	return &types.MsgDeleteSynchronizationFeedbackResponse{}, nil
}
