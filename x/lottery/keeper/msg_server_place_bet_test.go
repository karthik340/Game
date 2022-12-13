package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/karthik340/game/testutil/keeper"
	"github.com/karthik340/game/testutil/sample"
	abci "github.com/karthik340/game/x/lottery"
	"github.com/karthik340/game/x/lottery/keeper"
	"github.com/karthik340/game/x/lottery/types"
	"github.com/stretchr/testify/require"
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

func getAddresses(
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

func TestUserMsgServerCreate(t *testing.T) {
	k, bk, ak, ctx := keepertest.LotteryKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	count := 20
	senders, addresses := getAddresses(t, ctx, bk, count)

	bets := make([]*types.MsgPlaceBet, count)
	for i := 0; i < count; i++ {
		bets[i] = getBet(senders[i], 5, int64(i+1))
	}

	for i := 0; i < count; i++ {
		_, err := srv.PlaceBet(wctx, bets[i])
		require.NoError(t, err)
	}

	for i := 0; i < count; i++ {
		bal1 := bk.GetBalance(ctx, addresses[i], "token")
		fmt.Println("bal ", i, bal1)
	}

	modBal := GetModuleAccountBalance(ctx, ak, bk, types.ModuleName, "token")
	fmt.Println("module bal : ", modBal)

	abci.EndBlocker(ctx, *k)

	for i := 0; i < count; i++ {
		bal1 := bk.GetBalance(ctx, addresses[i], "token")
		fmt.Println("bal ", i, bal1)
	}

	modBal = GetModuleAccountBalance(ctx, ak, bk, types.ModuleName, "token")
	fmt.Println("module bal : ", modBal)
}
