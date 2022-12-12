package keeper_test

import (
	"testing"

	testkeeper "game/testutil/keeper"
	"game/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LotteryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.MinFee, k.MinFee(ctx))
	require.EqualValues(t, params.MinBet, k.MinBet(ctx))
	require.EqualValues(t, params.MaxBet, k.MaxBet(ctx))
	require.EqualValues(t, params.MinTxn, k.MinTxn(ctx))
}
