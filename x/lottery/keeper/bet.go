package keeper

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
)

// SetBetInCurrentRound store bet in the store in the current round
func (k Keeper) SetBetInCurrentRound(ctx sdk.Context, bet types.Bet) {

	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&bet)

	round, _ := k.GetRound(ctx)

	key := types.BetKey(
		round,
		bet.Sender,
	)

	store.Set(key, b)
}

// GetBetInCurrentRound returns a bet by given sender in current round
func (k Keeper) GetBetInCurrentRound(ctx sdk.Context, sender string) (val types.Bet, found bool) {

	var (
		round, _ = k.GetRound(ctx)
		store    = ctx.KVStore(k.storeKey)
		key      = types.BetKey(round, sender)
	)
	rawBytes := store.Get(key)
	if rawBytes == nil {
		return
	}

	k.cdc.MustUnmarshal(rawBytes, &val)
	return val, true
}

// CheckForProposerInBets checks if proposer has participated in lottery
func (k Keeper) CheckForProposerInBets(ctx sdk.Context) bool {
	var addr sdk.AccAddress

	err := addr.Unmarshal(ctx.BlockHeader().ProposerAddress)
	if err != nil {
		panic(err)
	}

	_, found := k.GetBetInCurrentRound(ctx, addr.String())

	return found
}

// GetAllBet returns all bets in all rounds
func (k Keeper) GetAllBet(ctx sdk.Context) (list []types.Bet) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BetKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBetsByRound return bets in given round
func (k Keeper) GetBetsByRound(ctx sdk.Context, round types.Round) Bets {
	var (
		key      = types.GetBetKeyWithRound(round)
		store    = ctx.KVStore(k.storeKey)
		iterator = sdk.KVStorePrefixIterator(store, key)
	)

	defer iterator.Close()

	list := make([]types.Bet, 0)

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}

// GetBetByTxNumber returns bet of n th transaction in the given round
func (k Keeper) GetBetByTxNumber(ctx sdk.Context, round types.Round, txNum uint64) (types.Bet, bool) {
	fmt.Println("tx num", txNum)
	var (
		key      = types.GetBetKeyWithRound(round)
		store    = ctx.KVStore(k.storeKey)
		iterator = sdk.KVStorePrefixIterator(store, key)
	)

	defer iterator.Close()

	fmt.Println(k.GetBetsByRound(ctx, types.Round{Val: 0}))

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bet
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		fmt.Println("see ", val)
		if val.TxNum == txNum {
			return val, true
		}
	}

	return types.Bet{}, false
}

// SetWinner sets the winner in given round
func (k Keeper) SetWinner(ctx sdk.Context, round types.Round, winner string) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&types.Winner{
		Winner: winner,
	})

	key := types.WinnerKey(
		round,
	)

	store.Set(key, b)
}

// GetWinner gets the winner in given round
func (k Keeper) GetWinner(ctx sdk.Context, round types.Round) (val types.Winner, found bool) {
	var (
		store    = ctx.KVStore(k.storeKey)
		key      = types.WinnerKey(round)
		rawBytes = store.Get(key)
	)
	if rawBytes == nil {
		return
	}

	k.cdc.MustUnmarshal(rawBytes, &val)
	return val, true
}

// MarshalBets serializes bets and return bytes
func (k Keeper) MarshalBets(bets []types.Bet) []byte {
	var bytes []byte
	for _, bet := range bets {
		bytes = append(bytes, k.cdc.MustMarshal(&bet)...)
	}
	return bytes
}

// PayWinner pays the winner based on winner's bet size
func (k Keeper) PayWinner(ctx sdk.Context, winner types.Bet, bets Bets) {
	highestBet, lowestBet := bets.GetHighAndLowBets()
	totalBet, totalFee := bets.GetTotalBetSizeAndFee(winner.Bet.Denom)

	winnerAddr, err := sdk.AccAddressFromBech32(winner.Sender)
	if err != nil {
		panic(err)
	}

	if winner.Bet.Amount.Equal(highestBet) { // if highest bet, pay totalBet and totalFee collected in the round
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx, types.ModuleName,
			winnerAddr,
			sdk.NewCoins(totalBet.Add(totalFee)))
		if err != nil {
			panic(err)
		}
		return
	}

	if winner.Bet.Amount.Equal(lowestBet) { // if lowest bet, don't pay anything
		return
	} // if lowest bet, don't pay

	// In other case pay only bet size
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winnerAddr, sdk.NewCoins(totalBet))
	if err != nil {
		panic(err)
	}
}

// GetWinnerIndex hashes the transactions, does modulo on last 16 bits in hash by transaction count
func (k Keeper) GetWinnerIndex(bets []types.Bet, txnCount uint64, round uint64) uint64 {
	rawData := k.MarshalBets(bets)
	rawData = append(rawData, sdk.Uint64ToBigEndian(round)...) // append round so that there is randomness,
	hash := crypto.Keccak256Hash(rawData)                      // in case there are same transactions as previous round,
	num := new(big.Int).SetBytes(hash[16:]).Uint64()
	fmt.Println("num txcount", num, txnCount)

	return num % txnCount
}
