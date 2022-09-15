package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

func (k msgServer) CreateStoredMiner(goCtx context.Context, msg *types.MsgCreateStoredMiner) (*types.MsgCreateStoredMinerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetStoredMiner(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var storedMiner = types.StoredMiner{
		Creator:     msg.Creator,
		Index:       msg.Index,
		LastClaimed: msg.LastClaimed,
	}

	k.SetStoredMiner(
		ctx,
		storedMiner,
	)
	return &types.MsgCreateStoredMinerResponse{}, nil
}

func (k msgServer) UpdateStoredMiner(goCtx context.Context, msg *types.MsgUpdateStoredMiner) (*types.MsgUpdateStoredMinerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStoredMiner(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var storedMiner = types.StoredMiner{
		Creator:     msg.Creator,
		Index:       msg.Index,
		LastClaimed: msg.LastClaimed,
	}

	k.SetStoredMiner(ctx, storedMiner)

	return &types.MsgUpdateStoredMinerResponse{}, nil
}

func (k msgServer) DeleteStoredMiner(goCtx context.Context, msg *types.MsgDeleteStoredMiner) (*types.MsgDeleteStoredMinerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStoredMiner(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStoredMiner(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteStoredMinerResponse{}, nil
}
