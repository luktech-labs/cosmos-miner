package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/luktech-labs/cosmos-miner/testutil/keeper"
	"github.com/luktech-labs/cosmos-miner/testutil/nullify"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

func TestMinerInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.CosmosminerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestMinerInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMinerInfoRequest
		response *types.QueryGetMinerInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMinerInfoRequest{},
			response: &types.QueryGetMinerInfoResponse{MinerInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MinerInfo(wctx, tc.request)
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
