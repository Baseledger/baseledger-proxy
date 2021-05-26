package keeper

import (
	"github.com/unibrightio/baseledger/x/trustmesh/types"
)

var _ types.QueryServer = Keeper{}
