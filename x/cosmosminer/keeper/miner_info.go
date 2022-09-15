package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// SetMinerInfo set minerInfo in the store
func (k Keeper) SetMinerInfo(ctx sdk.Context, minerInfo types.MinerInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerInfoKey))
	b := k.cdc.MustMarshal(&minerInfo)
	store.Set([]byte{0}, b)
}

// GetMinerInfo returns minerInfo
func (k Keeper) GetMinerInfo(ctx sdk.Context) (val types.MinerInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMinerInfo removes minerInfo from the store
func (k Keeper) RemoveMinerInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MinerInfoKey))
	store.Delete([]byte{0})
}
