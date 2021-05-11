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

func (k Keeper) SynchronizationRequestAll(c context.Context, req *types.QueryAllSynchronizationRequestRequest) (*types.QueryAllSynchronizationRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var SynchronizationRequests []*types.SynchronizationRequest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	SynchronizationRequestStore := prefix.NewStore(store, types.KeyPrefix(types.SynchronizationRequestKey))

	pageRes, err := query.Paginate(SynchronizationRequestStore, req.Pagination, func(key []byte, value []byte) error {
		var SynchronizationRequest types.SynchronizationRequest
		if err := k.cdc.UnmarshalBinaryBare(value, &SynchronizationRequest); err != nil {
			return err
		}

		SynchronizationRequests = append(SynchronizationRequests, &SynchronizationRequest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSynchronizationRequestResponse{SynchronizationRequest: SynchronizationRequests, Pagination: pageRes}, nil
}

func (k Keeper) SynchronizationRequest(c context.Context, req *types.QueryGetSynchronizationRequestRequest) (*types.QueryGetSynchronizationRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var SynchronizationRequest types.SynchronizationRequest
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasSynchronizationRequest(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetSynchronizationRequestIDBytes(req.Id)), &SynchronizationRequest)

	return &types.QueryGetSynchronizationRequestResponse{SynchronizationRequest: &SynchronizationRequest}, nil
}
