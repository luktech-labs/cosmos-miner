package types_test

import (
	"testing"

	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				MinerInfo: types.MinerInfo{
					TemplateNextID: 51,
					MinerNextID:    51,
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedMinerTemplate",
			genState: &types.GenesisState{
				StoredMinerTemplateList: []types.StoredMinerTemplate{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated storedMiner",
			genState: &types.GenesisState{
				StoredMinerList: []types.StoredMiner{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated minerBalance",
			genState: &types.GenesisState{
				MinerBalanceList: []types.MinerBalance{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
