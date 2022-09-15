package cosmosminer

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/luktech-labs/cosmos-miner/testutil/sample"
	cosmosminersimulation "github.com/luktech-labs/cosmos-miner/x/cosmosminer/simulation"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cosmosminersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateStoredMiner = "op_weight_msg_stored_miner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStoredMiner int = 100

	opWeightMsgUpdateStoredMiner = "op_weight_msg_stored_miner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStoredMiner int = 100

	opWeightMsgDeleteStoredMiner = "op_weight_msg_stored_miner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStoredMiner int = 100

	opWeightMsgCreateMinerTemplate = "op_weight_msg_create_miner_template"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMinerTemplate int = 100

	opWeightMsgCreateMiner = "op_weight_msg_create_miner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMiner int = 100

	opWeightMsgMine = "op_weight_msg_mine"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMine int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cosmosminerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		StoredMinerList: []types.StoredMiner{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cosmosminerGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateStoredMiner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateStoredMiner, &weightMsgCreateStoredMiner, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStoredMiner = defaultWeightMsgCreateStoredMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStoredMiner,
		cosmosminersimulation.SimulateMsgCreateStoredMiner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStoredMiner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateStoredMiner, &weightMsgUpdateStoredMiner, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStoredMiner = defaultWeightMsgUpdateStoredMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStoredMiner,
		cosmosminersimulation.SimulateMsgUpdateStoredMiner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteStoredMiner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteStoredMiner, &weightMsgDeleteStoredMiner, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStoredMiner = defaultWeightMsgDeleteStoredMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStoredMiner,
		cosmosminersimulation.SimulateMsgDeleteStoredMiner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMinerTemplate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMinerTemplate, &weightMsgCreateMinerTemplate, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMinerTemplate = defaultWeightMsgCreateMinerTemplate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMinerTemplate,
		cosmosminersimulation.SimulateMsgCreateMinerTemplate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMiner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMiner, &weightMsgCreateMiner, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMiner = defaultWeightMsgCreateMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMiner,
		cosmosminersimulation.SimulateMsgCreateMiner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMine int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMine, &weightMsgMine, nil,
		func(_ *rand.Rand) {
			weightMsgMine = defaultWeightMsgMine
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMine,
		cosmosminersimulation.SimulateMsgMine(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
