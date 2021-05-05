package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/baseledger/x/baseledger/types"
	"strconv"
)

// GetBaseledgerTransactionCount get the total number of BaseledgerTransaction
func (k Keeper) GetBaseledgerTransactionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionCountKey))
	byteKey := types.KeyPrefix(types.BaseledgerTransactionCountKey)
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

// SetBaseledgerTransactionCount set the total number of BaseledgerTransaction
func (k Keeper) SetBaseledgerTransactionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionCountKey))
	byteKey := types.KeyPrefix(types.BaseledgerTransactionCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendBaseledgerTransaction appends a BaseledgerTransaction in the store with a new id and update the count
func (k Keeper) AppendBaseledgerTransaction(
	ctx sdk.Context,
	creator string,
	baseid string,
	payload string,
) uint64 {
	// Create the BaseledgerTransaction
	count := k.GetBaseledgerTransactionCount(ctx)
	var BaseledgerTransaction = types.BaseledgerTransaction{
		Creator: creator,
		Id:      count,
		Baseid:  baseid,
		Payload: payload,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionKey))
	value := k.cdc.MustMarshalBinaryBare(&BaseledgerTransaction)
	store.Set(GetBaseledgerTransactionIDBytes(BaseledgerTransaction.Id), value)

	// Update BaseledgerTransaction count
	k.SetBaseledgerTransactionCount(ctx, count+1)

	return count
}

// SetBaseledgerTransaction set a specific BaseledgerTransaction in the store
func (k Keeper) SetBaseledgerTransaction(ctx sdk.Context, BaseledgerTransaction types.BaseledgerTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionKey))
	b := k.cdc.MustMarshalBinaryBare(&BaseledgerTransaction)
	store.Set(GetBaseledgerTransactionIDBytes(BaseledgerTransaction.Id), b)
}

// GetBaseledgerTransaction returns a BaseledgerTransaction from its id
func (k Keeper) GetBaseledgerTransaction(ctx sdk.Context, id uint64) types.BaseledgerTransaction {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionKey))
	var BaseledgerTransaction types.BaseledgerTransaction
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBaseledgerTransactionIDBytes(id)), &BaseledgerTransaction)
	return BaseledgerTransaction
}

// HasBaseledgerTransaction checks if the BaseledgerTransaction exists in the store
func (k Keeper) HasBaseledgerTransaction(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionKey))
	return store.Has(GetBaseledgerTransactionIDBytes(id))
}

// GetBaseledgerTransactionOwner returns the creator of the BaseledgerTransaction
func (k Keeper) GetBaseledgerTransactionOwner(ctx sdk.Context, id uint64) string {
	return k.GetBaseledgerTransaction(ctx, id).Creator
}

// RemoveBaseledgerTransaction removes a BaseledgerTransaction from the store
func (k Keeper) RemoveBaseledgerTransaction(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionKey))
	store.Delete(GetBaseledgerTransactionIDBytes(id))
}

// GetAllBaseledgerTransaction returns all BaseledgerTransaction
func (k Keeper) GetAllBaseledgerTransaction(ctx sdk.Context) (list []types.BaseledgerTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseledgerTransactionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BaseledgerTransaction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBaseledgerTransactionIDBytes returns the byte representation of the ID
func GetBaseledgerTransactionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBaseledgerTransactionIDFromBytes returns ID in uint64 format from a byte array
func GetBaseledgerTransactionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
