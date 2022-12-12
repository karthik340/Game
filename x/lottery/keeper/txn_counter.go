package keeper

import (
	"game/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTxnCounter set txnCounter in the store
func (k Keeper) SetTxnCounter(ctx sdk.Context, txnCounter types.TxnCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnCounterKey))
	b := k.cdc.MustMarshal(&txnCounter)
	store.Set([]byte{0}, b)
}

// GetTxnCounter returns txnCounter
func (k Keeper) GetTxnCounter(ctx sdk.Context) (val types.TxnCounter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnCounterKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTxnCounter removes txnCounter from the store
func (k Keeper) RemoveTxnCounter(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnCounterKey))
	store.Delete([]byte{0})
}
