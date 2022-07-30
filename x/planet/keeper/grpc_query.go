package keeper

import (
	"github.com/MercesToken/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
