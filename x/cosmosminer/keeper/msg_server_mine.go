package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

var ErrMinerNotOwned = sdkerrors.Register(types.ModuleName, 0003, "specified miner is not owned by the signer")
var ErrMinerToReadyToClaim = sdkerrors.Register(types.ModuleName, 0004, "miner not ready to claim")

func (k msgServer) Mine(goCtx context.Context, msg *types.MsgMine) (*types.MsgMineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	miner, found := k.Keeper.GetStoredMiner(ctx, strconv.FormatUint(msg.MinerID, 10))
	if !found {
		return nil, ErrTemplateIDNotFound
	}

	// check if the miner is owned by the transaction signer
	if miner.Creator != msg.GetSigners()[0].String() {
		return nil, ErrMinerNotOwned
	}

	template, found := k.Keeper.GetStoredMinerTemplate(ctx, strconv.FormatUint(miner.TemplateIndex, 10))
	if !found {
		return nil, ErrTemplateIDNotFound
	}

	now := uint64(time.Now().Unix())

	// check if the miner is ready to claim
	secSinceLastClaim := now - miner.LastClaimed
	if secSinceLastClaim < template.ClaimTime {
		return nil, ErrMinerToReadyToClaim
	}

	produced := template.ProdPerSecond * secSinceLastClaim

	balance, _ := k.Keeper.GetMinerBalance(ctx, miner.Creator)
	balance.Amount += produced

	miner.LastClaimed = now

	k.Keeper.SetMinerBalance(ctx, balance)
	k.Keeper.SetStoredMiner(ctx, miner)

	return &types.MsgMineResponse{AmountMined: strconv.FormatUint(produced, 10)}, nil
}
