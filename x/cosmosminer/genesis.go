package cosmosminer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/keeper"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	emptyMinerInfo := types.MinerInfo{}
	// Set if defined
	if genState.MinerInfo != emptyMinerInfo {
		k.SetMinerInfo(ctx, genState.MinerInfo)
	}
	// Set all the storedMinerTemplate
	for _, elem := range genState.StoredMinerTemplateList {
		k.SetStoredMinerTemplate(ctx, elem)
	}
	// Set all the storedMiner
	for _, elem := range genState.StoredMinerList {
		k.SetStoredMiner(ctx, elem)
	}
	// Set all the minerBalance
	for _, elem := range genState.MinerBalanceList {
		k.SetMinerBalance(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all minerInfo
	minerInfo, found := k.GetMinerInfo(ctx)
	if found {
		genesis.MinerInfo = minerInfo
	}
	genesis.StoredMinerTemplateList = k.GetAllStoredMinerTemplate(ctx)
	genesis.StoredMinerList = k.GetAllStoredMiner(ctx)
	genesis.MinerBalanceList = k.GetAllMinerBalance(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
