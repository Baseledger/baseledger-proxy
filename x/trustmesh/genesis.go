package trustmesh

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/baseledger/x/trustmesh/keeper"
	"github.com/example/baseledger/x/trustmesh/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the SynchronizationFeedback
	for _, elem := range genState.SynchronizationFeedbackList {
		k.SetSynchronizationFeedback(ctx, *elem)
	}

	// Set SynchronizationFeedback count
	k.SetSynchronizationFeedbackCount(ctx, uint64(len(genState.SynchronizationFeedbackList)))

	// Set all the SynchronizationRequest
	for _, elem := range genState.SynchronizationRequestList {
		k.SetSynchronizationRequest(ctx, *elem)
	}

	// Set SynchronizationRequest count
	k.SetSynchronizationRequestCount(ctx, uint64(len(genState.SynchronizationRequestList)))

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all SynchronizationFeedback
	SynchronizationFeedbackList := k.GetAllSynchronizationFeedback(ctx)
	for _, elem := range SynchronizationFeedbackList {
		elem := elem
		genesis.SynchronizationFeedbackList = append(genesis.SynchronizationFeedbackList, &elem)
	}

	// Get all SynchronizationRequest
	SynchronizationRequestList := k.GetAllSynchronizationRequest(ctx)
	for _, elem := range SynchronizationRequestList {
		elem := elem
		genesis.SynchronizationRequestList = append(genesis.SynchronizationRequestList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
