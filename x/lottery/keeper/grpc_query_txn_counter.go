package keeper

import (
	"context"

	"game/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TxnCounter(c context.Context, req *types.QueryGetTxnCounterRequest) (*types.QueryGetTxnCounterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTxnCounter(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTxnCounterResponse{TxnCounter: val}, nil
}
