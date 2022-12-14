package keeper_test

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/karthik340/game/testutil/keeper"
	testutil "github.com/karthik340/game/testutil/keeper"
	"github.com/karthik340/game/testutil/sample"
	abci "github.com/karthik340/game/x/lottery"
	"github.com/karthik340/game/x/lottery/keeper"
	"github.com/karthik340/game/x/lottery/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"strconv"
	"testing"
)

// Prevent strconv unused error
var _ = strconv.IntSize

// Below are useful helpers to write test code easily.

func fundAddr(t *testing.T, ctx sdk.Context, bk types.BankKeeper, addr sdk.AccAddress, amt sdk.Coin) {
	t.Helper()

	amt1 := sdk.NewCoins(amt)
	err := bk.MintCoins(ctx, types.ModuleName, amt1)
	require.NoError(t, err)

	err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, amt1)
	require.NoError(t, err)
}

func GetModuleAccountBalance(
	ctx sdk.Context,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	moduleName string,
	denom string,
) sdk.Coin {
	address := ak.GetModuleAddress(moduleName)
	return sdk.NewCoin(denom, bk.GetBalance(ctx, address, denom).Amount)
}

func getRandomAddressesWithFunds(
	t *testing.T,
	ctx sdk.Context,
	bk types.BankKeeper,
	count int,
) ([]string, []sdk.AccAddress) {
	t.Helper()

	var err error
	senders := make([]string, count)
	addresses := make([]sdk.AccAddress, count)

	for i := 0; i < count; i++ {
		senders[i] = sample.AccAddress()
		addresses[i], err = sdk.AccAddressFromBech32(senders[i])
		require.NoError(t, err)

		fundAddr(t, ctx, bk, addresses[i], sdk.NewCoin("token", sdk.NewInt(500)))
	}

	return senders, addresses
}

func getBet(creator string, fee, bet int64) *types.MsgPlaceBet {
	return &types.MsgPlaceBet{
		Sender: creator,
		Fee:    sdk.NewCoin("token", sdk.NewInt(fee)),
		Bet:    sdk.NewCoin("token", sdk.NewInt(bet)),
	}
}

func placeBets(
	t *testing.T,
	msgServer types.MsgServer,
	wctx context.Context,
	bets []*types.MsgPlaceBet,
) {
	for _, bet := range bets {
		_, err := msgServer.PlaceBet(wctx, bet) // 20 clients pace bets
		require.NoError(t, err)
	}
}

func getBalances(ctx sdk.Context, bk types.BankKeeper, addresses []sdk.AccAddress, count int) []sdk.Coin {
	balance := make([]sdk.Coin, count)

	for i := 0; i < count; i++ {
		balance[i] = bk.GetBalance(ctx, addresses[i], "token")
	}

	return balance
}

// verifyBalancesAfterBet checks if balances are collected by lottery
func verifyBalancesAfterBet(
	t *testing.T,
	balancesBeforeBet []sdk.Coin,
	balancesAfterBet []sdk.Coin,
	bets []*types.MsgPlaceBet,
) {
	for i, bet := range bets {
		require.Equal(t, balancesBeforeBet[i].Sub(bet.Bet.Add(bet.Fee)), balancesAfterBet[i])
	}
}

func verifyModuleBalanceAfterBet(
	t *testing.T,
	k *keeper.Keeper,
	ctx sdk.Context,
	moduleBalanceBeforeBet sdk.Coin,
	moduleBalanceAfterBet sdk.Coin,
	round types.Round,
) {
	bets := k.GetBetsByRound(ctx, round)
	totalBet, totalFee := bets.GetTotalBetSizeAndFee("token")
	require.Equal(t, moduleBalanceBeforeBet.Add(totalBet.Add(totalFee)), moduleBalanceAfterBet)
}

func checkIfBalancesAreSame(
	t *testing.T,
	beforeBalances []sdk.Coin,
	afterBalances []sdk.Coin,
) {
	for i, beforeBal := range beforeBalances {
		require.Equal(t, beforeBal, afterBalances[i])
	}
}

func TestValidate(t *testing.T) {
	k, bk, _, ctx := keepertest.LotteryKeeper(t, nil)
	msgServer := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	count := 1
	senders, _ := getRandomAddressesWithFunds(t, ctx, bk, count)

	testCases := []struct {
		name          string
		msgPlacebet   *types.MsgPlaceBet
		expectedError error
	}{
		{
			name:          "invalid fee",
			msgPlacebet:   getBet(senders[0], 4, int64(100)),
			expectedError: types.ErrBetValidationFailed,
		},
		{
			name:          "invalid max bet",
			msgPlacebet:   getBet(senders[0], 5, int64(101)),
			expectedError: types.ErrBetValidationFailed,
		},
		{
			name:          "invalid min bet",
			msgPlacebet:   getBet(senders[0], 5, int64(0)),
			expectedError: types.ErrBetValidationFailed,
		},
		{
			name:        "valid fee",
			msgPlacebet: getBet(senders[0], 5, int64(55)),
		},
		{
			name:        "valid max bet",
			msgPlacebet: getBet(senders[0], 5, int64(100)),
		},
		{
			name:        "valid min bet",
			msgPlacebet: getBet(senders[0], 5, int64(1)),
		},
		{
			name:        "valid in between bet",
			msgPlacebet: getBet(senders[0], 5, int64(33)),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := msgServer.PlaceBet(wctx, test.msgPlacebet)
			if test.expectedError != nil {
				require.Equal(t, test.expectedError, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestPlaceBets(t *testing.T) {
	k, bk, ak, ctx := keepertest.LotteryKeeper(t, nil)
	msgServer := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	count := 1
	senders, addresses := getRandomAddressesWithFunds(t, ctx, bk, count)

	msgPlacebet := getBet(senders[0], 5, int64(100))

	_, err := msgServer.PlaceBet(wctx, msgPlacebet) // first bet from user-1
	require.NoError(t, err)

	balancesAfterBet := bk.GetBalance(ctx, addresses[0], "token")
	moduleBalanceAfterBet := GetModuleAccountBalance(ctx, ak, bk, types.ModuleName, "token")

	require.Equal(t, moduleBalanceAfterBet.Amount, sdk.NewInt(105))
	require.Equal(t, balancesAfterBet.Amount, sdk.NewInt(395))

	msgPlacebet.Bet.Amount = sdk.NewInt(50)

	_, err = msgServer.PlaceBet(wctx, msgPlacebet) // second bet from user-1
	require.NoError(t, err)

	balanceAfterSecondBet := bk.GetBalance(ctx, addresses[0], "token")
	moduleBalanceAfterSecondBet := GetModuleAccountBalance(ctx, ak, bk, types.ModuleName, "token")

	require.Equal(t, sdk.NewInt(445), balanceAfterSecondBet.Amount) // check if previous bet returned to user
	require.Equal(t, sdk.NewInt(55), moduleBalanceAfterSecondBet.Amount)

	bet, found := k.GetBetInCurrentRound(ctx, senders[0]) // check if bet stored in kv store
	require.True(t, found)

	require.Equal(t, msgPlacebet.Bet, bet.Bet)
	require.Equal(t, msgPlacebet.Fee, bet.Fee)
	require.Equal(t, msgPlacebet.Sender, bet.Sender)
	require.Equal(t, false, bet.Status)
	require.Equal(t, uint64(0), bet.TxNum)
}

// End to End Test tests both place bet and end blocker
func TestEndToEnd(t *testing.T) {
	k, bk, ak, ctx := keepertest.LotteryKeeper(t, nil)
	msgServer := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	round := types.Round{Val: 0}

	// get random 20 addresses with balance of 500 tokens in each
	count := 20
	senders, addresses := getRandomAddressesWithFunds(t, ctx, bk, count)

	msgPlacebets := make([]*types.MsgPlaceBet, count)
	for i := 0; i < count; i++ {
		// get msgPlacebets of bet sizes 1 2 3  ... 20 with fee 5
		msgPlacebets[i] = getBet(senders[i], 5, int64(i+1))
	}

	balancesBeforeBet := getBalances(ctx, bk, addresses, count)
	moduleBalanceBeforeBet := GetModuleAccountBalance(ctx, ak, bk, types.ModuleName, "token")

	placeBets(t, msgServer, wctx, msgPlacebets) // All 20 clients place msgPlacebets

	balancesAfterBet := getBalances(ctx, bk, addresses, count)
	moduleBalanceAfterBet := GetModuleAccountBalance(ctx, ak, bk, types.ModuleName, "token")

	verifyBalancesAfterBet(t, balancesBeforeBet, balancesAfterBet, msgPlacebets) // verify senders balances
	// verify modules balances
	verifyModuleBalanceAfterBet(t, k, ctx, moduleBalanceBeforeBet, moduleBalanceAfterBet, round)

	bets := k.GetBetsByRound(ctx, round)
	winnerTxNum := k.GetWinnerIndex(bets, 20, 0)
	expectedWinner, _ := k.GetBetByTxNumber(ctx, round, winnerTxNum) // calculate the expected winner

	abci.EndBlocker(ctx, *k) // call end block to pick winner

	balancesAfterRound := getBalances(ctx, bk, addresses, count)
	moduleBalanceAfterRound := GetModuleAccountBalance(ctx, ak, bk, types.ModuleName, "token")

	actualWinner, found := k.GetWinner(ctx, types.Round{
		Val: uint64(0),
	})
	require.True(t, found)

	require.Equal(t, expectedWinner.Sender, actualWinner.Winner) // check if winner matches
	require.Equal(t, winnerTxNum, expectedWinner.TxNum)          // check if txNum matches

	totalBet, totalFee := bets.GetTotalBetSizeAndFee("token")

	bet, found := k.GetBetByTxNumber(ctx, round, expectedWinner.TxNum) // check if winner status is set
	require.True(t, found)
	require.Equal(t, bet.Status, true)

	// make sure balances of non-winner clients didn't get credited tokens
	// check if module balances are properly debited
	if expectedWinner.Bet.Amount.Equal(sdk.NewInt(1)) { // if lowest bet, then shouldn't receive tokens
		checkIfBalancesAreSame(t, balancesAfterBet[:], balancesAfterRound[:])
		require.Equal(t, moduleBalanceAfterBet, moduleBalanceAfterRound)
	} else if expectedWinner.Bet.Amount.Equal(sdk.NewInt(20)) {
		// if highest bet, then should receive tokens (total bets,total fee)
		require.Equal(t, balancesAfterBet[19].Add(totalBet.Add(totalFee)), balancesAfterRound[19])
		checkIfBalancesAreSame(t, balancesAfterBet[:19], balancesAfterRound[:19])
		require.Equal(t, sdk.NewInt(0), moduleBalanceAfterRound.Amount)
	} else {
		// otherwise receives only total bet but not fee
		i := expectedWinner.Bet.Amount.Int64() - 1 // users index = users bet - 1
		require.Equal(t, balancesAfterBet[i].Add(totalBet), balancesAfterRound[i])

		checkIfBalancesAreSame(t, balancesAfterBet[:i], balancesAfterRound[:i])
		checkIfBalancesAreSame(t, balancesAfterBet[i+1:], balancesAfterRound[i+1:])

		require.Equal(t, moduleBalanceAfterBet.Sub(totalBet), moduleBalanceAfterRound)
	}

	actualRound, found := k.GetRound(ctx)
	require.True(t, found)
	require.Equal(t, uint64(1), actualRound.Val)

	actualTxnCounter, found := k.GetTxnCounter(ctx)
	require.True(t, found)
	require.Equal(t, uint64(0), actualTxnCounter.Val)
}

func TestEndBlocker_CheckProposerBet(t *testing.T) {
	proposer := sample.AccAddress()
	proposerAddr, err := sdk.AccAddressFromBech32(proposer)
	require.NoError(t, err)

	k, bk, _, ctx := keepertest.LotteryKeeper(t, &tmproto.Header{
		Time:            testutil.ExampleTimestamp,
		Height:          testutil.ExampleHeight,
		ProposerAddress: proposerAddr, // set proposer address so that we can make a bet
	})

	k.SetRound(ctx, types.Round{
		Val: 0, // initialize with round as 0
	})

	msgServer := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	// get random 20 addresses with balance of 500 tokens in each
	count := 20
	senders, _ := getRandomAddressesWithFunds(t, ctx, bk, count)

	fundAddr(t, ctx, bk, proposerAddr, sdk.NewCoin("token", sdk.NewInt(500)))
	senders[8] = proposer // replace one of the participant as proposer

	msgPlacebets := make([]*types.MsgPlaceBet, count)
	for i := 0; i < count; i++ {
		// get msgPlacebets of bet sizes 1 2 3  ... 20 with fee 5
		msgPlacebets[i] = getBet(senders[i], 5, int64(i+1))
	}

	placeBets(t, msgServer, wctx, msgPlacebets) // All 20 clients place msgPlacebets

	abci.EndBlocker(ctx, *k) // call end block to pick winner

	// make sure transaction not executed as proposer address present in  bets
	actualRound, found := k.GetRound(ctx)
	require.True(t, found)
	require.Equal(t, uint64(0), actualRound.Val)
}

func TestEndBlocker_CheckTxnCounter(t *testing.T) {
	k, bk, _, ctx := keepertest.LotteryKeeper(t, nil)

	k.SetRound(ctx, types.Round{
		Val: 0, // initialize with round as 0
	})

	msgServer := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	// get random 9 addresses with balance of 500 tokens in each
	count := 9
	senders, _ := getRandomAddressesWithFunds(t, ctx, bk, count)

	msgPlacebets := make([]*types.MsgPlaceBet, count)
	for i := 0; i < count; i++ {
		// get msgPlacebets of bet sizes 1 2 3  ... 9 with fee 5
		msgPlacebets[i] = getBet(senders[i], 5, int64(i+1))
	}

	placeBets(t, msgServer, wctx, msgPlacebets) // All 9 clients place msgPlacebets

	abci.EndBlocker(ctx, *k) // call end block to pick winner

	// make sure transaction not executed as number of transactions less than 10
	actualRound, found := k.GetRound(ctx)
	require.True(t, found)
	require.Equal(t, uint64(0), actualRound.Val)
}

func TestEndBlocker_CheckProposerWinner(t *testing.T) {
	proposer := sample.AccAddress()
	proposerAddr, err := sdk.AccAddressFromBech32(proposer)
	require.NoError(t, err)

	k, bk, _, ctx := keepertest.LotteryKeeper(t, &tmproto.Header{
		Time:            testutil.ExampleTimestamp,
		Height:          testutil.ExampleHeight,
		ProposerAddress: proposerAddr, // set proposer address so that we can make a bet
	})

	k.SetRound(ctx, types.Round{
		Val: 0, // initialize with round as 0
	})

	msgServer := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	// get random 20 addresses with balance of 500 tokens in each
	count := 20
	senders, _ := getRandomAddressesWithFunds(t, ctx, bk, count)

	k.SetValidatorsWinner(ctx, proposer, senders[18]) // proposer setting client18 as winner

	msgPlacebets := make([]*types.MsgPlaceBet, count)
	for i := 0; i < count; i++ {
		// get msgPlacebets of bet sizes 1 2 3  ... 20 with fee 5
		msgPlacebets[i] = getBet(senders[i], 5, int64(i+1))
	}

	placeBets(t, msgServer, wctx, msgPlacebets) // All 20 clients place msgPlacebets

	abci.EndBlocker(ctx, *k) // call end block to pick winner

	actualWinner, found := k.GetWinner(ctx, types.Round{
		Val: uint64(0),
	})
	require.True(t, found)

	require.Equal(t, senders[18], actualWinner.Winner) // check if proposer's winner is the one who actual won
}
