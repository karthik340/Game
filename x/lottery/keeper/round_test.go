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

func createTestRound(keeper *keeper.Keeper, ctx sdk.Context) types.Round {
	item := types.Round{}
	keeper.SetRound(ctx, item)
	return item
}

func TestRoundGet(t *testing.T) {
	keeper, _, _, ctx := keepertest.LotteryKeeper(t, nil)
	item := createTestRound(keeper, ctx)
	rst, found := keeper.GetRound(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
