package keeper

import (
	"game/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MinFee(ctx),
		k.MinBet(ctx),
		k.MaxBet(ctx),
		k.MinTxn(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MinFee returns the MinFee param
func (k Keeper) MinFee(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinFee, &res)
	return
}

// MinBet returns the MinBet param
func (k Keeper) MinBet(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinBet, &res)
	return
}

// MaxBet returns the MaxBet param
func (k Keeper) MaxBet(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxBet, &res)
	return
}

// MinTxn returns the MinTxn param
func (k Keeper) MinTxn(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinTxn, &res)
	return
}
