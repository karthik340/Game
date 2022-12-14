package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/keeper"
	"github.com/karthik340/game/x/lottery/types"
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
	// Set all the bet
	for _, elem := range genState.BetList {
		k.SetBetInCurrentRound(ctx, elem)
	}
	// Set if defined
if genState.ValidatorsWinner != nil {
	k.SetValidatorsWinner(ctx, *genState.ValidatorsWinner)
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
	genesis.BetList = k.GetAllBet(ctx)
	// Get all validatorsWinner
validatorsWinner, found := k.GetValidatorsWinner(ctx)
if found {
	genesis.ValidatorsWinner = &validatorsWinner
}
// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
