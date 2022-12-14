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

		if txnCount.Val >= k.MinTxn(ctx) { // if number of txns are greater than equal to 10 then pick winnerBet
			// if proposers winner or proposers winner bet not found the we use hash func to find winner
			proposersWinner, found := k.GetProposersWinner(ctx) // get the winner selected by proposer
			if found {
				winnerBet, found := k.GetBetInCurrentRound(ctx, proposersWinner.Winner) // get the bet of winner
				if found {
					k.PayWinner(ctx, winnerBet, bets)
					k.ModifyLotteryData(ctx, winnerBet, round, txnCount)
					return
				}
			}

			winnerTxNum := k.GetWinnerIndex(bets, txnCount.Val, round.Val)
			winnerBet, _ := k.GetBetByTxNumber(ctx, round, winnerTxNum)

			k.PayWinner(ctx, winnerBet, bets)

			k.ModifyLotteryData(ctx, winnerBet, round, txnCount)
		}
	}
}
