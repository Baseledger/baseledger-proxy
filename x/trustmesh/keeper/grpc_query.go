package keeper

import (
	"github.com/example/baseledger/x/trustmesh/types"
)

var _ types.QueryServer = Keeper{}
