package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "game/testutil/keeper"
	"game/testutil/nullify"
	"game/x/lottery/types"
)

func TestRoundQuery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestRound(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRoundRequest
		response *types.QueryGetRoundResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRoundRequest{},
			response: &types.QueryGetRoundResponse{Round: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Round(wctx, tc.request)
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
