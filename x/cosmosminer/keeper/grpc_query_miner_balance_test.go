package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/luktech-labs/cosmos-miner/testutil/keeper"
	"github.com/luktech-labs/cosmos-miner/testutil/nullify"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMinerBalanceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMinerBalance(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMinerBalanceRequest
		response *types.QueryGetMinerBalanceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMinerBalanceRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetMinerBalanceResponse{MinerBalance: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMinerBalanceRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetMinerBalanceResponse{MinerBalance: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMinerBalanceRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MinerBalance(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestMinerBalanceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMinerBalance(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMinerBalanceRequest {
		return &types.QueryAllMinerBalanceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MinerBalanceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MinerBalance), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MinerBalance),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MinerBalanceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MinerBalance), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MinerBalance),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MinerBalanceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MinerBalance),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MinerBalanceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
