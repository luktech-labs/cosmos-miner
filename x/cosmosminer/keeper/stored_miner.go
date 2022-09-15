package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// SetStoredMiner set a specific storedMiner in the store from its index
func (k Keeper) SetStoredMiner(ctx sdk.Context, storedMiner types.StoredMiner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerKeyPrefix))
	b := k.cdc.MustMarshal(&storedMiner)
	store.Set(types.StoredMinerKey(
		storedMiner.Index,
	), b)
}

// GetStoredMiner returns a storedMiner from its index
func (k Keeper) GetStoredMiner(
	ctx sdk.Context,
	index string,

) (val types.StoredMiner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerKeyPrefix))

	b := store.Get(types.StoredMinerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredMiner removes a storedMiner from the store
func (k Keeper) RemoveStoredMiner(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerKeyPrefix))
	store.Delete(types.StoredMinerKey(
		index,
	))
}

// GetAllStoredMiner returns all storedMiner
func (k Keeper) GetAllStoredMiner(ctx sdk.Context) (list []types.StoredMiner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredMiner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
