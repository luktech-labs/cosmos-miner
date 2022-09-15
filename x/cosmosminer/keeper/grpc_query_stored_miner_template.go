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

func (k Keeper) StoredMinerTemplateAll(c context.Context, req *types.QueryAllStoredMinerTemplateRequest) (*types.QueryAllStoredMinerTemplateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedMinerTemplates []types.StoredMinerTemplate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedMinerTemplateStore := prefix.NewStore(store, types.KeyPrefix(types.StoredMinerTemplateKeyPrefix))

	pageRes, err := query.Paginate(storedMinerTemplateStore, req.Pagination, func(key []byte, value []byte) error {
		var storedMinerTemplate types.StoredMinerTemplate
		if err := k.cdc.Unmarshal(value, &storedMinerTemplate); err != nil {
			return err
		}

		storedMinerTemplates = append(storedMinerTemplates, storedMinerTemplate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredMinerTemplateResponse{StoredMinerTemplate: storedMinerTemplates, Pagination: pageRes}, nil
}

func (k Keeper) StoredMinerTemplate(c context.Context, req *types.QueryGetStoredMinerTemplateRequest) (*types.QueryGetStoredMinerTemplateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredMinerTemplate(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredMinerTemplateResponse{StoredMinerTemplate: val}, nil
}
