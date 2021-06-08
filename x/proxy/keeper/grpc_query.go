package keeper

import (
	"github.com/unibrightio/baseledger/x/proxy/types"
)

var _ types.QueryServer = Keeper{}
