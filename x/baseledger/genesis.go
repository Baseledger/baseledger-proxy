package baseledger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/example/baseledger/x/baseledger/keeper"
	"github.com/example/baseledger/x/baseledger/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the BaseledgerTransaction
	for _, elem := range genState.BaseledgerTransactionList {
		k.SetBaseledgerTransaction(ctx, *elem)
	}

	// Set BaseledgerTransaction count
	k.SetBaseledgerTransactionCount(ctx, uint64(len(genState.BaseledgerTransactionList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all BaseledgerTransaction
	BaseledgerTransactionList := k.GetAllBaseledgerTransaction(ctx)
	for _, elem := range BaseledgerTransactionList {
		elem := elem
		genesis.BaseledgerTransactionList = append(genesis.BaseledgerTransactionList, &elem)
	}

	return genesis
}
