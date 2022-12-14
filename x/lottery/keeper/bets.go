package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
)

type Bets []types.Bet

// GetHighAndLowBets returns highest and lowest bet
func (b Bets) GetHighAndLowBets() (sdk.Int, sdk.Int) {
	high := sdk.NewInt(0)
	low := sdk.NewInt(100)

	for _, bet := range b {
		high = sdk.MaxInt(high, bet.Bet.Amount)
		low = sdk.MinInt(low, bet.Bet.Amount)
	}

	return high, low
}

// GetTotalBetSizeAndFee returns sum of all bets and fee
func (b Bets) GetTotalBetSizeAndFee(denom string) (sdk.Coin, sdk.Coin) {
	fee := sdk.NewInt(0)
	totalBet := sdk.NewInt(0)

	for _, bet := range b {
		fee = fee.Add(sdk.NewInt(bet.Fee.Amount.Int64()))
		totalBet = totalBet.Add(sdk.NewInt(bet.Bet.Amount.Int64()))
	}

	return sdk.NewCoin(denom, totalBet), sdk.NewCoin(denom, fee)
}
