package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/baseledger/x/trustmesh/types"
	"strconv"
)

// GetSynchronizationFeedbackCount get the total number of SynchronizationFeedback
func (k Keeper) GetSynchronizationFeedbackCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackCountKey))
	byteKey := types.KeyPrefix(types.SynchronizationFeedbackCountKey)
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

// SetSynchronizationFeedbackCount set the total number of SynchronizationFeedback
func (k Keeper) SetSynchronizationFeedbackCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackCountKey))
	byteKey := types.KeyPrefix(types.SynchronizationFeedbackCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendSynchronizationFeedback appends a SynchronizationFeedback in the store with a new id and update the count
func (k Keeper) AppendSynchronizationFeedback(
	ctx sdk.Context,
	creator string,
	Approved string,
	BusinessObject string,
	BaseledgerBusinessObjectIDOfApprovedObject string,
	Workgroup string,
	Recipient string,
	HashOfObjectToApprove string,
	OriginalBaseledgerTransactionID string,
	OriginalOffchainProcessMessageID string,
	FeedbackMessage string,
) uint64 {
	// Create the SynchronizationFeedback
	count := k.GetSynchronizationFeedbackCount(ctx)
	var SynchronizationFeedback = types.SynchronizationFeedback{
		Creator:        creator,
		Id:             count,
		Approved:       Approved,
		BusinessObject: BusinessObject,
		BaseledgerBusinessObjectIDOfApprovedObject: BaseledgerBusinessObjectIDOfApprovedObject,
		Workgroup:                        Workgroup,
		Recipient:                        Recipient,
		HashOfObjectToApprove:            HashOfObjectToApprove,
		OriginalBaseledgerTransactionID:  OriginalBaseledgerTransactionID,
		OriginalOffchainProcessMessageID: OriginalOffchainProcessMessageID,
		FeedbackMessage:                  FeedbackMessage,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackKey))
	value := k.cdc.MustMarshalBinaryBare(&SynchronizationFeedback)
	store.Set(GetSynchronizationFeedbackIDBytes(SynchronizationFeedback.Id), value)

	// Update SynchronizationFeedback count
	k.SetSynchronizationFeedbackCount(ctx, count+1)

	return count
}

// SetSynchronizationFeedback set a specific SynchronizationFeedback in the store
func (k Keeper) SetSynchronizationFeedback(ctx sdk.Context, SynchronizationFeedback types.SynchronizationFeedback) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackKey))
	b := k.cdc.MustMarshalBinaryBare(&SynchronizationFeedback)
	store.Set(GetSynchronizationFeedbackIDBytes(SynchronizationFeedback.Id), b)
}

// GetSynchronizationFeedback returns a SynchronizationFeedback from its id
func (k Keeper) GetSynchronizationFeedback(ctx sdk.Context, id uint64) types.SynchronizationFeedback {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackKey))
	var SynchronizationFeedback types.SynchronizationFeedback
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetSynchronizationFeedbackIDBytes(id)), &SynchronizationFeedback)
	return SynchronizationFeedback
}

// HasSynchronizationFeedback checks if the SynchronizationFeedback exists in the store
func (k Keeper) HasSynchronizationFeedback(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackKey))
	return store.Has(GetSynchronizationFeedbackIDBytes(id))
}

// GetSynchronizationFeedbackOwner returns the creator of the SynchronizationFeedback
func (k Keeper) GetSynchronizationFeedbackOwner(ctx sdk.Context, id uint64) string {
	return k.GetSynchronizationFeedback(ctx, id).Creator
}

// RemoveSynchronizationFeedback removes a SynchronizationFeedback from the store
func (k Keeper) RemoveSynchronizationFeedback(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackKey))
	store.Delete(GetSynchronizationFeedbackIDBytes(id))
}

// GetAllSynchronizationFeedback returns all SynchronizationFeedback
func (k Keeper) GetAllSynchronizationFeedback(ctx sdk.Context) (list []types.SynchronizationFeedback) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SynchronizationFeedbackKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SynchronizationFeedback
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSynchronizationFeedbackIDBytes returns the byte representation of the ID
func GetSynchronizationFeedbackIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSynchronizationFeedbackIDFromBytes returns ID in uint64 format from a byte array
func GetSynchronizationFeedbackIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
