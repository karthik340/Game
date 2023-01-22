package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetWinnerByRound(goCtx context.Context, req *types.QueryGetWinnerByRoundRequest) (*types.QueryGetWinnerByRoundResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	winner, found := k.GetWinner(ctx, types.Round{
		Val: req.GetRound(),
	})

	if !found {
		return &types.QueryGetWinnerByRoundResponse{}, nil
	}

	return &types.QueryGetWinnerByRoundResponse{
		Winner: winner.Winner,
	}, nil
}
