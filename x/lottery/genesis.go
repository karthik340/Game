package lottery

import (
	"game/x/lottery/keeper"
	"game/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.Round != nil {
		k.SetRound(ctx, *genState.Round)
	}
	// Set if defined
	if genState.TxnCounter != nil {
		k.SetTxnCounter(ctx, *genState.TxnCounter)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all round
	round, found := k.GetRound(ctx)
	if found {
		genesis.Round = &round
	}
	// Get all txnCounter
	txnCounter, found := k.GetTxnCounter(ctx)
	if found {
		genesis.TxnCounter = &txnCounter
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
