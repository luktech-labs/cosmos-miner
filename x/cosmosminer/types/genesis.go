package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MinerInfo:               MinerInfo{TemplateNextID: DefaultIndex, MinerNextID: DefaultIndex},
		StoredMinerTemplateList: []StoredMinerTemplate{},
		StoredMinerList:         []StoredMiner{},
		MinerBalanceList:        []MinerBalance{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedMinerTemplate
	storedMinerTemplateIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredMinerTemplateList {
		index := string(StoredMinerTemplateKey(elem.Index))
		if _, ok := storedMinerTemplateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedMinerTemplate")
		}
		storedMinerTemplateIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in storedMiner
	storedMinerIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredMinerList {
		index := string(StoredMinerKey(elem.Index))
		if _, ok := storedMinerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedMiner")
		}
		storedMinerIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in minerBalance
	minerBalanceIndexMap := make(map[string]struct{})

	for _, elem := range gs.MinerBalanceList {
		index := string(MinerBalanceKey(elem.Index))
		if _, ok := minerBalanceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for minerBalance")
		}
		minerBalanceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
