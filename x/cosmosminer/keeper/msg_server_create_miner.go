package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

var ErrTemplateIDNotFound = sdkerrors.Register(types.ModuleName, 0001, "provided template id not found")
var ErrNotPayable = sdkerrors.Register(types.ModuleName, 0002, "cannot pay for the miner")

func (k msgServer) CreateMiner(goCtx context.Context, msg *types.MsgCreateMiner) (*types.MsgCreateMinerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minerInfo, found := k.Keeper.GetMinerInfo(ctx)
	if !found {
		panic("miner info not found")
	}

	template, found := k.Keeper.GetStoredMinerTemplate(ctx, strconv.FormatUint(msg.MinerTemplateID, 10))
	if !found {
		return nil, ErrTemplateIDNotFound
	}

	//spending tokens from the signer's wallet in order to buy the miner
	coin := sdk.NewCoin("token", sdk.NewIntFromUint64(template.Price))
	err := k.bank.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], types.ModuleName, sdk.Coins{coin})
	if err != nil {
		return nil, sdkerrors.Wrap(err, ErrNotPayable.Error())
	}

	newIndex := strconv.FormatUint(minerInfo.MinerNextID, 10)
	storedMiner := types.StoredMiner{
		LastClaimed:   uint64(time.Now().Unix()),
		Index:         newIndex,
		TemplateIndex: msg.MinerTemplateID,
		Creator:       msg.Creator,
	}

	k.Keeper.SetStoredMiner(ctx, storedMiner)
	minerInfo.MinerNextID += 1
	k.Keeper.SetMinerInfo(ctx, minerInfo)

	return &types.MsgCreateMinerResponse{MinerIndex: newIndex}, nil
}
