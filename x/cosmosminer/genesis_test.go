package cosmosminer_test

import (
	"testing"

	keepertest "github.com/luktech-labs/cosmos-miner/testutil/keeper"
	"github.com/luktech-labs/cosmos-miner/testutil/nullify"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MinerInfo: &types.MinerInfo{
			NextID: 26,
		},
		StoredMinerTemplateList: []types.StoredMinerTemplate{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		StoredMinerList: []types.StoredMiner{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		MinerBalanceList: []types.MinerBalance{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosminerKeeper(t)
	cosmosminer.InitGenesis(ctx, *k, genesisState)
	got := cosmosminer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.MinerInfo, got.MinerInfo)
	require.ElementsMatch(t, genesisState.StoredMinerTemplateList, got.StoredMinerTemplateList)
	require.ElementsMatch(t, genesisState.StoredMinerList, got.StoredMinerList)
	require.ElementsMatch(t, genesisState.MinerBalanceList, got.MinerBalanceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
