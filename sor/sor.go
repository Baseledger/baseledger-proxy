package sor

import (
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/logger"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
)

// TODO: add some sor logic, discuss
func ProcessFeedback(offchainProcessMessage proxytypes.OffchainProcessMessage, workgroupId uuid.UUID, payload string) {
	logger.Infof("SOR PROCESSING FEEDBACK")
}
