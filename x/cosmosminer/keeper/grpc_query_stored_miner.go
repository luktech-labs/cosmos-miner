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

func (k Keeper) StoredMinerAll(c context.Context, req *types.QueryAllStoredMinerRequest) (*types.QueryAllStoredMinerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedMiners []types.StoredMiner
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedMinerStore := prefix.NewStore(store, types.KeyPrefix(types.StoredMinerKeyPrefix))

	pageRes, err := query.Paginate(storedMinerStore, req.Pagination, func(key []byte, value []byte) error {
		var storedMiner types.StoredMiner
		if err := k.cdc.Unmarshal(value, &storedMiner); err != nil {
			return err
		}

		storedMiners = append(storedMiners, storedMiner)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredMinerResponse{StoredMiner: storedMiners, Pagination: pageRes}, nil
}

func (k Keeper) StoredMiner(c context.Context, req *types.QueryGetStoredMinerRequest) (*types.QueryGetStoredMinerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredMiner(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredMinerResponse{StoredMiner: val}, nil
}
