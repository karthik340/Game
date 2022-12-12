package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinFee = []byte("MinFee")
	// TODO: Determine the default value
	DefaultMinFee uint64 = 5
)

var (
	KeyMinBet = []byte("MinBet")
	// TODO: Determine the default value
	DefaultMinBet uint64 = 1
)

var (
	KeyMaxBet = []byte("MaxBet")
	// TODO: Determine the default value
	DefaultMaxBet uint64 = 100
)

var (
	KeyMinTxn = []byte("MinTxn")
	// TODO: Determine the default value
	DefaultMinTxn uint64 = 10
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minFee uint64,
	minBet uint64,
	maxBet uint64,
	minTxn uint64,
) Params {
	return Params{
		MinFee: minFee,
		MinBet: minBet,
		MaxBet: maxBet,
		MinTxn: minTxn,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinFee,
		DefaultMinBet,
		DefaultMaxBet,
		DefaultMinTxn,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinFee, &p.MinFee, validateMinFee),
		paramtypes.NewParamSetPair(KeyMinBet, &p.MinBet, validateMinBet),
		paramtypes.NewParamSetPair(KeyMaxBet, &p.MaxBet, validateMaxBet),
		paramtypes.NewParamSetPair(KeyMinTxn, &p.MinTxn, validateMinTxn),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinFee(p.MinFee); err != nil {
		return err
	}

	if err := validateMinBet(p.MinBet); err != nil {
		return err
	}

	if err := validateMaxBet(p.MaxBet); err != nil {
		return err
	}

	if err := validateMinTxn(p.MinTxn); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMinFee validates the MinFee param
func validateMinFee(v interface{}) error {
	minFee, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minFee

	return nil
}

// validateMinBet validates the MinBet param
func validateMinBet(v interface{}) error {
	minBet, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minBet

	return nil
}

// validateMaxBet validates the MaxBet param
func validateMaxBet(v interface{}) error {
	maxBet, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxBet

	return nil
}

// validateMinTxn validates the MinTxn param
func validateMinTxn(v interface{}) error {
	minTxn, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minTxn

	return nil
}
