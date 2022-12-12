package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/keeper"
)

// EndBlocker picks winner from lottery at the end of block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	found := k.CheckForProposerInBets(ctx)
	if !found { // continue only if proposer has not participated in lottery
		round, _ := k.GetRound(ctx)
		bets := k.GetBetsByRound(ctx, round)
		txnCount, _ := k.GetTxnCounter(ctx)

		if txnCount.Val >= 10 { // if number of txns are greater than equal to 10 then pick winner
			winnerIndex := k.GetWinnerIndex(bets, txnCount.Val, round.Val)
			winner, _ := k.GetBetByTxNumberInCurrentRound(ctx, round, winnerIndex)

			k.PayWinner(ctx, winner, bets)

			winner.Status = true // change the status of winner status to true
			round.Val += 1       // increment round
			txnCount.Val = 0     // make txn txnCount zero

			k.SetBetInCurrentRound(ctx, winner)
			k.SetRound(ctx, round)
			k.SetTxnCounter(ctx, txnCount)
		}
	}
}
