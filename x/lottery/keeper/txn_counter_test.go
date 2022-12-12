package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "game/testutil/keeper"
	"game/testutil/nullify"
	"game/x/lottery/keeper"
	"game/x/lottery/types"
)

func createTestTxnCounter(keeper *keeper.Keeper, ctx sdk.Context) types.TxnCounter {
	item := types.TxnCounter{}
	keeper.SetTxnCounter(ctx, item)
	return item
}

func TestTxnCounterGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestTxnCounter(keeper, ctx)
	rst, found := keeper.GetTxnCounter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTxnCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestTxnCounter(keeper, ctx)
	keeper.RemoveTxnCounter(ctx)
	_, found := keeper.GetTxnCounter(ctx)
	require.False(t, found)
}
