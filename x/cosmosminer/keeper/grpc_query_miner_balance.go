package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MinerBalanceAll(c context.Context, req *types.QueryAllMinerBalanceRequest) (*types.QueryAllMinerBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var minerBalances []types.MinerBalance
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	minerBalanceStore := prefix.NewStore(store, types.KeyPrefix(types.MinerBalanceKeyPrefix))

	pageRes, err := query.Paginate(minerBalanceStore, req.Pagination, func(key []byte, value []byte) error {
		var minerBalance types.MinerBalance
		if err := k.cdc.Unmarshal(value, &minerBalance); err != nil {
			return err
		}

		minerBalances = append(minerBalances, minerBalance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMinerBalanceResponse{MinerBalance: minerBalances, Pagination: pageRes}, nil
}

func (k Keeper) MinerBalance(c context.Context, req *types.QueryGetMinerBalanceRequest) (*types.QueryGetMinerBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMinerBalance(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMinerBalanceResponse{MinerBalance: val}, nil
}
