package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatorsWinner(c context.Context, req *types.QueryGetValidatorsWinnerRequest) (*types.QueryGetValidatorsWinnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidatorsWinner(ctx)
	if !found {
	    return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetValidatorsWinnerResponse{ValidatorsWinner: val}, nil
}