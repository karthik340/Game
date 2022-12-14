package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetValidatorsWinner set validatorsWinner in the store
func (k Keeper) SetValidatorsWinner(ctx sdk.Context, validatorsWinner types.ValidatorsWinner) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsWinnerKey))
	b := k.cdc.MustMarshal(&validatorsWinner)
	store.Set([]byte{0}, b)
}

// GetValidatorsWinner returns validatorsWinner
func (k Keeper) GetValidatorsWinner(ctx sdk.Context) (val types.ValidatorsWinner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsWinnerKey))

	b := store.Get([]byte{0})
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorsWinner removes validatorsWinner from the store
func (k Keeper) RemoveValidatorsWinner(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorsWinnerKey))
	store.Delete([]byte{0})
}
