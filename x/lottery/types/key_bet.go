package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BetKeyPrefix is the prefix to retrieve all Bet
	BetKeyPrefix = "Bet/value/"
)

// BetKey returns the store key to retrieve a Bet from the index fields
func BetKey(
	sender string,
) []byte {
	var key []byte

	senderBytes := []byte(sender)
	key = append(key, senderBytes...)
	key = append(key, []byte("/")...)

	return key
}
