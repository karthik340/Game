package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
)

// SetRound sets round in the store
func (k Keeper) SetRound(ctx sdk.Context, round types.Round) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoundKey))
	b := k.cdc.MustMarshal(&round)
	store.Set([]byte{0}, b)
}

// GetRound returns round
func (k Keeper) GetRound(ctx sdk.Context) (val types.Round, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoundKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
