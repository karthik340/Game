package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Round: &Round{
			Val: 0,
		},
		TxnCounter: &TxnCounter{
			Val: 0,
		},
		BetList:          []Bet{},
		ValidatorsWinner: nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in bet
	betIndexMap := make(map[string]struct{})

	for _, elem := range gs.BetList {
		index := string(BetKey(Round{Val: 0}, elem.Sender))
		if _, ok := betIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for bet")
		}
		betIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
