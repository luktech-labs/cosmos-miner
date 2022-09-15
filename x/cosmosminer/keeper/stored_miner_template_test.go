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

func createNStoredMinerTemplate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredMinerTemplate {
	items := make([]types.StoredMinerTemplate, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredMinerTemplate(ctx, items[i])
	}
	return items
}

func TestStoredMinerTemplateGet(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	items := createNStoredMinerTemplate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredMinerTemplate(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredMinerTemplateRemove(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	items := createNStoredMinerTemplate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredMinerTemplate(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredMinerTemplate(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredMinerTemplateGetAll(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	items := createNStoredMinerTemplate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredMinerTemplate(ctx)),
	)
}
