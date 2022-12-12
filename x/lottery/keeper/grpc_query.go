package keeper

import (
	"game/x/lottery/types"
)

var _ types.QueryServer = Keeper{}
