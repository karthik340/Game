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
		TxnCounter: &types.TxnCounter{
			Val: 52,
		},
		BetList: []types.Bet{
			{
				Sender: "0",
			},
			{
				Sender: "1",
			},
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
	require.Equal(t, genesisState.TxnCounter, got.TxnCounter)
	require.ElementsMatch(t, genesisState.BetList, got.BetList)
	// this line is used by starport scaffolding # genesis/test/assert
}
