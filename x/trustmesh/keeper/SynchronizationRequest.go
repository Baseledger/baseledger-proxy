package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/baseledger/x/trustmesh/types"
	"strconv"
)

// GetSynchronizationRequestCount get the total number of SynchronizationRequest
func (k Keeper) GetSynchronizationRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestCountKey))
	byteKey := types.KeyPrefix(types.SynchronizationRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetSynchronizationRequestCount set the total number of SynchronizationRequest
func (k Keeper) SetSynchronizationRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestCountKey))
	byteKey := types.KeyPrefix(types.SynchronizationRequestCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendSynchronizationRequest appends a SynchronizationRequest in the store with a new id and update the count
func (k Keeper) AppendSynchronizationRequest(
	ctx sdk.Context,
	creator string,
	WorkgroupID string,
	Recipient string,
	WorkstepType string,
	BusinessObjectType string,
	BaseledgerBusinessObjectID string,
	BusinessObject string,
	ReferencedBaseledgerBusinessObjectID string,
) uint64 {
	// Create the SynchronizationRequest
	count := k.GetSynchronizationRequestCount(ctx)
	var SynchronizationRequest = types.SynchronizationRequest{
		Creator:                              creator,
		Id:                                   count,
		WorkgroupID:                          WorkgroupID,
		Recipient:                            Recipient,
		WorkstepType:                         WorkstepType,
		BusinessObjectType:                   BusinessObjectType,
		BaseledgerBusinessObjectID:           BaseledgerBusinessObjectID,
		BusinessObject:                       BusinessObject,
		ReferencedBaseledgerBusinessObjectID: ReferencedBaseledgerBusinessObjectID,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestKey))
	value := k.cdc.MustMarshalBinaryBare(&SynchronizationRequest)
	store.Set(GetSynchronizationRequestIDBytes(SynchronizationRequest.Id), value)

	// Update SynchronizationRequest count
	k.SetSynchronizationRequestCount(ctx, count+1)

	return count
}

// SetSynchronizationRequest set a specific SynchronizationRequest in the store
func (k Keeper) SetSynchronizationRequest(ctx sdk.Context, SynchronizationRequest types.SynchronizationRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestKey))
	b := k.cdc.MustMarshalBinaryBare(&SynchronizationRequest)
	store.Set(GetSynchronizationRequestIDBytes(SynchronizationRequest.Id), b)
}

// GetSynchronizationRequest returns a SynchronizationRequest from its id
func (k Keeper) GetSynchronizationRequest(ctx sdk.Context, id uint64) types.SynchronizationRequest {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestKey))
	var SynchronizationRequest types.SynchronizationRequest
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetSynchronizationRequestIDBytes(id)), &SynchronizationRequest)
	return SynchronizationRequest
}

// HasSynchronizationRequest checks if the SynchronizationRequest exists in the store
func (k Keeper) HasSynchronizationRequest(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestKey))
	return store.Has(GetSynchronizationRequestIDBytes(id))
}

// GetSynchronizationRequestOwner returns the creator of the SynchronizationRequest
func (k Keeper) GetSynchronizationRequestOwner(ctx sdk.Context, id uint64) string {
	return k.GetSynchronizationRequest(ctx, id).Creator
}

// RemoveSynchronizationRequest removes a SynchronizationRequest from the store
func (k Keeper) RemoveSynchronizationRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestKey))
	store.Delete(GetSynchronizationRequestIDBytes(id))
}

// GetAllSynchronizationRequest returns all SynchronizationRequest
func (k Keeper) GetAllSynchronizationRequest(ctx sdk.Context) (list []types.SynchronizationRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SynchronizationRequest
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSynchronizationRequestIDBytes returns the byte representation of the ID
func GetSynchronizationRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSynchronizationRequestIDFromBytes returns ID in uint64 format from a byte array
func GetSynchronizationRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
