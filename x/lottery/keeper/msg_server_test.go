package keeper_test

import (
	"context"
	"testing"

	keepertest "game/testutil/keeper"
	"game/x/lottery/keeper"
	"game/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LotteryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
