package keeper

import (
	"github.com/carmonasl/loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
