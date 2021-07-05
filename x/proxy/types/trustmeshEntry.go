package types

import (
	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres

	uuid "github.com/kthomas/go.uuid"
	common "github.com/unibrightio/baseledger/common"
	"github.com/unibrightio/baseledger/dbutil"
	"github.com/unibrightio/baseledger/logger"
)

type TrustmeshEntry struct {
	TendermintBlockId                    sql.NullString
	TendermintTransactionId              uuid.UUID
	TendermintTransactionTimestamp       sql.NullString
	EntryType                            string
	SenderOrgId                          uuid.UUID
	ReceiverOrgId                        uuid.UUID
	WorkgroupId                          uuid.UUID
	WorkstepType                         string
	BaseledgerTransactionType            string
	BaseledgerTransactionId              uuid.UUID
	ReferencedBaseledgerTransactionId    uuid.UUID
	BusinessObjectType                   string
	BaseledgerBusinessObjectId           uuid.UUID
	ReferencedBaseledgerBusinessObjectId uuid.UUID
	OffchainProcessMessageId             uuid.UUID
	ReferencedProcessMessageId           uuid.UUID
	CommitmentState                      string
	TransactionHash                      string
}

func (t *TrustmeshEntry) Create() bool {
	t.CommitmentState = common.UncommittedCommitmentState
	t.TendermintBlockId = sql.NullString{Valid: false}
	t.TendermintTransactionTimestamp = sql.NullString{Valid: false}
	if dbutil.Db.GetConn().NewRecord(t) {
		result := dbutil.Db.GetConn().Create(&t)
		rowsAffected := result.RowsAffected
		errors := result.GetErrors()
		if len(errors) > 0 {
			logger.Errorf("errors while creating new entry %v\n", errors)
			return false
		}
		return rowsAffected > 0
	}

	return false
}
