package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/luktech-labs/cosmos-miner/testutil/keeper"
	"github.com/luktech-labs/cosmos-miner/testutil/nullify"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/keeper"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

func createTestMinerInfo(keeper *keeper.Keeper, ctx sdk.Context) types.MinerInfo {
	item := types.MinerInfo{}
	keeper.SetMinerInfo(ctx, item)
	return item
}

func TestMinerInfoGet(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	item := createTestMinerInfo(keeper, ctx)
	rst, found := keeper.GetMinerInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMinerInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	createTestMinerInfo(keeper, ctx)
	keeper.RemoveMinerInfo(ctx)
	_, found := keeper.GetMinerInfo(ctx)
	require.False(t, found)
}
