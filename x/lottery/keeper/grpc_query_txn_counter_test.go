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

func TestTxnCounterQuery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTxnCounter(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTxnCounterRequest
		response *types.QueryGetTxnCounterResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTxnCounterRequest{},
			response: &types.QueryGetTxnCounterResponse{TxnCounter: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TxnCounter(wctx, tc.request)
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
