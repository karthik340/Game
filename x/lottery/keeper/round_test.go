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

func createTestRound(keeper *keeper.Keeper, ctx sdk.Context) types.Round {
	item := types.Round{}
	keeper.SetRound(ctx, item)
	return item
}

func TestRoundGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestRound(keeper, ctx)
	rst, found := keeper.GetRound(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestRoundRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestRound(keeper, ctx)
	keeper.RemoveRound(ctx)
	_, found := keeper.GetRound(ctx)
	require.False(t, found)
}
