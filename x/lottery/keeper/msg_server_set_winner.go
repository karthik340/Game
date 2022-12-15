package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
)

func (k msgServer) SetWinner(goCtx context.Context, msg *types.MsgSetWinner) (*types.MsgSetWinnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	k.SetValidatorsWinner(
		ctx,
		msg.Validator,
		msg.Winner,
	)

	return &types.MsgSetWinnerResponse{}, nil
}
