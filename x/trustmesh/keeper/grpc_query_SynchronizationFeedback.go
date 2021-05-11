package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/example/baseledger/x/trustmesh/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SynchronizationFeedbackAll(c context.Context, req *types.QueryAllSynchronizationFeedbackRequest) (*types.QueryAllSynchronizationFeedbackResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var SynchronizationFeedbacks []*types.SynchronizationFeedback
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	SynchronizationFeedbackStore := prefix.NewStore(store, types.KeyPrefix(types.SynchronizationFeedbackKey))

	pageRes, err := query.Paginate(SynchronizationFeedbackStore, req.Pagination, func(key []byte, value []byte) error {
		var SynchronizationFeedback types.SynchronizationFeedback
		if err := k.cdc.UnmarshalBinaryBare(value, &SynchronizationFeedback); err != nil {
			return err
		}

		SynchronizationFeedbacks = append(SynchronizationFeedbacks, &SynchronizationFeedback)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSynchronizationFeedbackResponse{SynchronizationFeedback: SynchronizationFeedbacks, Pagination: pageRes}, nil
}

func (k Keeper) SynchronizationFeedback(c context.Context, req *types.QueryGetSynchronizationFeedbackRequest) (*types.QueryGetSynchronizationFeedbackResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var SynchronizationFeedback types.SynchronizationFeedback
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasSynchronizationFeedback(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetSynchronizationFeedbackIDBytes(req.Id)), &SynchronizationFeedback)

	return &types.QueryGetSynchronizationFeedbackResponse{SynchronizationFeedback: &SynchronizationFeedback}, nil
}
