package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// SetMinerBalance set a specific minerBalance in the store from its index
func (k Keeper) SetMinerBalance(ctx sdk.Context, minerBalance types.MinerBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerBalanceKeyPrefix))
	b := k.cdc.MustMarshal(&minerBalance)
	store.Set(types.MinerBalanceKey(
		minerBalance.Index,
	), b)
}

// GetMinerBalance returns a minerBalance from its index
func (k Keeper) GetMinerBalance(
	ctx sdk.Context,
	index string,

) (val types.MinerBalance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerBalanceKeyPrefix))

	b := store.Get(types.MinerBalanceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMinerBalance removes a minerBalance from the store
func (k Keeper) RemoveMinerBalance(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerBalanceKeyPrefix))
	store.Delete(types.MinerBalanceKey(
		index,
	))
}

// GetAllMinerBalance returns all minerBalance
func (k Keeper) GetAllMinerBalance(ctx sdk.Context) (list []types.MinerBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerBalanceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MinerBalance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
