package keeper

import (
	"github.com/example/baseledger/x/baseledger/types"
)

var _ types.QueryServer = Keeper{}
