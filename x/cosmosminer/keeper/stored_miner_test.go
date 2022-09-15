package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/luktech-labs/cosmos-miner/testutil/keeper"
	"github.com/luktech-labs/cosmos-miner/testutil/nullify"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/keeper"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredMiner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredMiner {
	items := make([]types.StoredMiner, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredMiner(ctx, items[i])
	}
	return items
}

func TestStoredMinerGet(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	items := createNStoredMiner(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredMiner(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredMinerRemove(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	items := createNStoredMiner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredMiner(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredMiner(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredMinerGetAll(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	items := createNStoredMiner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredMiner(ctx)),
	)
}
