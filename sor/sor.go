package sor

import (
	"fmt"

	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
)

func ProcessFeedback(offchainProcessMessage proxytypes.OffchainProcessMessage, workgroupId string, payload string) {
	fmt.Println("SOR PROCESSING FEEDBACK")
}
