package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrSample              = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrBetValidationFailed = sdkerrors.Register(
		ModuleName,
		1101,
		"bet should be between  1 and 100 and fee should be above 5",
	)
	ErrInsufficientFunds = sdkerrors.Register(ModuleName, 1102, "insufficient funds")
)
