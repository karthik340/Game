package simulation

import (
	"math/rand"

	"game/x/lottery/keeper"
	"game/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgPlaceBet(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPlaceBet{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the PlaceBet simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "PlaceBet simulation not implemented"), nil, nil
	}
}
