package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

var (
	// BetKeyPrefix is the prefix to retrieve all Bet
	BetKeyPrefix  = []byte{0x01}
	UserKeyPrefix = []byte{0x02}
)

// BetKey returns the store key to retrieve a Bet from the index fields
func BetKey(
	round Round,
	sender string,
) []byte { // (BetKeyPrefix + round + sender) -> Bet
	return append(append(BetKeyPrefix, sdk.Uint64ToBigEndian(round.Val)...), sender...)
}

func GetBetKeyWithRound(round Round) []byte {
	return append(BetKeyPrefix, sdk.Uint64ToBigEndian(round.Val)...)
}

func UserKey(
	sender string,
	round Round,
	txnCounter TxnCounter,
) []byte { // (UserKeyPrefix + sender + round + txnNumber) -> Bet
	return append(
		append(
			append(
				UserKeyPrefix, sender...,
			),
			sdk.Uint64ToBigEndian(round.Val)...,
		),
		sdk.Uint64ToBigEndian(txnCounter.Val)...,
	)
}
