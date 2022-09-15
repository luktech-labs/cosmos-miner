package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/luktech-labs/cosmos-miner/testutil/keeper"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/keeper"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosminerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
