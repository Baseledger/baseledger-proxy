package types

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres

	"github.com/unibrightio/baseledger/dbutil"
)

const defaultTransactionStatus = "UNCOMMITTED"

type TrustmeshEntry struct {
	TendermintBlockId                    sql.NullString
	TendermintTransactionId              string
	TendermintTransactionTimestamp       sql.NullString
	Type                                 string
	SenderOrgId                          string
	ReceiverOrgId                        string
	WorkgroupId                          string
	WorkstepType                         string
	BaseledgerTransactionType            string
	BaseledgerTransactionId              string
	ReferencedBaseledgerTransactionId    string
	BusinessObjectType                   string
	BaseledgerBusinessObjectId           string
	ReferencedBaseledgerBusinessObjectId string
	OffchainProcessMessageId             string
	ReferencedProcessMessageId           string
	TransactionStatus                    string
	TransactionHash                      string
}

func (t *TrustmeshEntry) Create() bool {
	t.TransactionStatus = defaultTransactionStatus
	t.TendermintBlockId = sql.NullString{Valid: false}
	t.TendermintTransactionTimestamp = sql.NullString{Valid: false}
	if dbutil.Db.GetConn().NewRecord(t) {
		result := dbutil.Db.GetConn().Create(&t)
		rowsAffected := result.RowsAffected
		errors := result.GetErrors()
		if len(errors) > 0 {
			fmt.Printf("errors while creating new entry %v\n", errors)
			return false
		}
		return rowsAffected > 0
	}

	return false
}
