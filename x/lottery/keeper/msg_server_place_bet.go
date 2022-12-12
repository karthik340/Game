package keeper

import (
	"context"

	"game/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceBet(goCtx context.Context, msg *types.MsgPlaceBet) (*types.MsgPlaceBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPlaceBetResponse{}, nil
}
