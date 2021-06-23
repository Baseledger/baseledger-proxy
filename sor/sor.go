package sor

import (
	"fmt"

	uuid "github.com/kthomas/go.uuid"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
)

// TODO: add some sor logic, discuss
func ProcessFeedback(offchainProcessMessage proxytypes.OffchainProcessMessage, workgroupId uuid.UUID, payload string) {
	fmt.Println("SOR PROCESSING FEEDBACK")
}
