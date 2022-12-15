package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/karthik340/game/x/lottery/keeper"
	"github.com/karthik340/game/x/lottery/types"
)

func SimulateMsgSetWinner(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSetWinner{
			Winner: simAccount.Address.String(),
		}

		// TODO: Handling the SetWinner simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SetWinner simulation not implemented"), nil, nil
	}
}
