package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/karthik340/game/testutil/keeper"
	"github.com/karthik340/game/testutil/nullify"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/keeper"
	"github.com/karthik340/game/x/lottery/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Bet {
	items := make([]types.Bet, n)
	for i := range items {
		items[i].Sender = strconv.Itoa(i)

		keeper.SetBetInCurrentRound(ctx, items[i])
	}
	return items
}

func TestBetGet(t *testing.T) {
	k, _, _, ctx := keepertest.LotteryKeeper(t, nil)
	items := createNBet(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetBetInCurrentRound(ctx,
			item.Sender,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestBetGetAll(t *testing.T) {
	k, _, _, ctx := keepertest.LotteryKeeper(t, nil)
	items := createNBet(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllBet(ctx)),
	)
}
