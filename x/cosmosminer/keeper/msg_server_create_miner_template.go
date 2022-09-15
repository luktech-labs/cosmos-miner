package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

func (k msgServer) CreateMinerTemplate(goCtx context.Context, msg *types.MsgCreateMinerTemplate) (*types.MsgCreateMinerTemplateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minerInfo, found := k.Keeper.GetMinerInfo(ctx)
	if !found {
		panic("miner info not found")
	}

	newIndex := strconv.FormatUint(minerInfo.TemplateNextID, 10)

	storedMinerTemplate := types.StoredMinerTemplate{
		Index:         newIndex,
		Name:          msg.Name,
		Price:         msg.Price,
		ClaimTime:     msg.MinClaimTime,
		ProdPerSecond: msg.ProdPerSecond,
	}

	k.Keeper.SetStoredMinerTemplate(ctx, storedMinerTemplate)
	minerInfo.TemplateNextID += 1
	k.Keeper.SetMinerInfo(ctx, minerInfo)

	return &types.MsgCreateMinerTemplateResponse{MinerTemplateIndex: newIndex}, nil
}
