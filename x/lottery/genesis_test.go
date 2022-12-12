package lottery_test

import (
	"testing"

	keepertest "game/testutil/keeper"
	"game/testutil/nullify"
	"game/x/lottery"
	"game/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Round: &types.Round{
			Val: 80,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Round, got.Round)
	// this line is used by starport scaffolding # genesis/test/assert
}
