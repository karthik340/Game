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

func createTestValidatorsWinner(keeper *keeper.Keeper, ctx sdk.Context) types.ValidatorsWinner {
	item := types.ValidatorsWinner{}
	keeper.SetValidatorsWinner(ctx, item)
	return item
}

func TestValidatorsWinnerGet(t *testing.T) {
	keeper, _, _, ctx := keepertest.LotteryKeeper(t, nil)
	item := createTestValidatorsWinner(keeper, ctx)
	rst, found := keeper.GetValidatorsWinner(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestValidatorsWinnerRemove(t *testing.T) {
	keeper, _, _, ctx := keepertest.LotteryKeeper(t, nil)
	createTestValidatorsWinner(keeper, ctx)
	keeper.RemoveValidatorsWinner(ctx)
	_, found := keeper.GetValidatorsWinner(ctx)
	require.False(t, found)
}
