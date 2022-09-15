package keeper_test

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/luktech-labs/cosmos-miner/testutil/keeper"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/keeper"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	alice = "cosmos1wnth30849m6dgeyjmaa9y65wveyhcsvs2zs65e"
	bob   = "cosmos1jawrxymkxe4tfnza4a089yu4hsdvwhks7uew4e"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateMinerTemplate(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CosmosminerKeeper(t)
	cosmosminer.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateMinerTemplate(t *testing.T) {
	msgServer, _, ctx := setupMsgServerCreateMinerTemplate(t)
	response, err := msgServer.CreateMinerTemplate(ctx, &types.MsgCreateMinerTemplate{
		Creator:       alice,
		Name:          "3xRtx3080",
		Price:         10,
		ProdPerSecond: 10,
		MinClaimTime:  60,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateMinerTemplateResponse{MinerTemplateIndex: "1"}, *response)
}

func TestCreateMinerTemplateWasSaved(t *testing.T) {
	msgServer, k, ctx := setupMsgServerCreateMinerTemplate(t)
	newMinerTemplate := types.MsgCreateMinerTemplate{
		Creator:       alice,
		Name:          "3xRtx3080",
		Price:         10,
		ProdPerSecond: 10,
		MinClaimTime:  60,
	}
	response, err := msgServer.CreateMinerTemplate(ctx, &newMinerTemplate)
	require.Nil(t, err)
	unwrappedCtx := sdk.UnwrapSDKContext(ctx)

	minerInfo, found := k.GetMinerInfo(unwrappedCtx)
	require.True(t, found)
	require.EqualValues(t, types.MinerInfo{
		TemplateNextID: 2,
		MinerNextID:    types.DefaultGenesis().MinerInfo.MinerNextID,
	}, minerInfo)

	minerTemplate, found := k.GetStoredMinerTemplate(unwrappedCtx, response.MinerTemplateIndex)
	require.True(t, found)
	require.EqualValues(t, types.StoredMinerTemplate{
		Index:         response.MinerTemplateIndex,
		Name:          newMinerTemplate.Name,
		Price:         newMinerTemplate.Price,
		ClaimTime:     newMinerTemplate.MinClaimTime,
		ProdPerSecond: newMinerTemplate.ProdPerSecond,
	}, minerTemplate)
}
