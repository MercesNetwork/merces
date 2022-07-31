package keeper

import (
	"github.com/MercesNetwork/merces/x/planet/types"
)

var _ types.QueryServer = Keeper{}
