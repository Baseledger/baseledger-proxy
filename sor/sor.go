package sor

import (
	"fmt"

	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
)

// TODO: add some sor logic, discuss
func ProcessFeedback(offchainProcessMessage proxytypes.OffchainProcessMessage, workgroupId string, payload string) {
	fmt.Println("SOR PROCESSING FEEDBACK")
}
