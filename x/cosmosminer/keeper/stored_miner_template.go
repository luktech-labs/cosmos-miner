package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// SetStoredMinerTemplate set a specific storedMinerTemplate in the store from its index
func (k Keeper) SetStoredMinerTemplate(ctx sdk.Context, storedMinerTemplate types.StoredMinerTemplate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerTemplateKeyPrefix))
	b := k.cdc.MustMarshal(&storedMinerTemplate)
	store.Set(types.StoredMinerTemplateKey(
		storedMinerTemplate.Index,
	), b)
}

// GetStoredMinerTemplate returns a storedMinerTemplate from its index
func (k Keeper) GetStoredMinerTemplate(
	ctx sdk.Context,
	index string,

) (val types.StoredMinerTemplate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerTemplateKeyPrefix))

	b := store.Get(types.StoredMinerTemplateKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredMinerTemplate removes a storedMinerTemplate from the store
func (k Keeper) RemoveStoredMinerTemplate(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerTemplateKeyPrefix))
	store.Delete(types.StoredMinerTemplateKey(
		index,
	))
}

// GetAllStoredMinerTemplate returns all storedMinerTemplate
func (k Keeper) GetAllStoredMinerTemplate(ctx sdk.Context) (list []types.StoredMinerTemplate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMinerTemplateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredMinerTemplate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
