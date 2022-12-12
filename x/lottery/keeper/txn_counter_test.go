package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/karthik340/game/testutil/keeper"
	"github.com/karthik340/game/testutil/nullify"

	"github.com/karthik340/game/x/lottery/keeper"
	"github.com/karthik340/game/x/lottery/types"
)

func createTestTxnCounter(keeper *keeper.Keeper, ctx sdk.Context) types.TxnCounter {
	item := types.TxnCounter{}
	keeper.SetTxnCounter(ctx, item)
	return item
}

func TestTxnCounterGet(t *testing.T) {
	keeper, _, _, ctx := keepertest.LotteryKeeper(t)
	item := createTestTxnCounter(keeper, ctx)
	rst, found := keeper.GetTxnCounter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
