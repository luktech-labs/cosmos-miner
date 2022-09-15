package keeper

import (
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

var _ types.QueryServer = Keeper{}
