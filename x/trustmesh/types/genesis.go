package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		SynchronizationFeedbackList: []*SynchronizationFeedback{},
		SynchronizationRequestList:  []*SynchronizationRequest{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in SynchronizationFeedback
	SynchronizationFeedbackIdMap := make(map[uint64]bool)

	for _, elem := range gs.SynchronizationFeedbackList {
		if _, ok := SynchronizationFeedbackIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for SynchronizationFeedback")
		}
		SynchronizationFeedbackIdMap[elem.Id] = true
	}
	// Check for duplicated ID in SynchronizationRequest
	SynchronizationRequestIdMap := make(map[uint64]bool)

	for _, elem := range gs.SynchronizationRequestList {
		if _, ok := SynchronizationRequestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for SynchronizationRequest")
		}
		SynchronizationRequestIdMap[elem.Id] = true
	}

	return nil
}
