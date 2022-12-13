package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/karthik340/game/testutil/keeper"
	"github.com/karthik340/game/x/lottery/keeper"
	"github.com/karthik340/game/x/lottery/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, _, _, ctx := keepertest.LotteryKeeper(t, nil)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
