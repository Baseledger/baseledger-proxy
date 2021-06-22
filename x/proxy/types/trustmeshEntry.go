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
	db, err := dbutil.InitBaseledgerDBConnection()

	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
		return false
	}

	t.TransactionStatus = defaultTransactionStatus
	t.TendermintBlockId = sql.NullString{Valid: false}
	t.TendermintTransactionTimestamp = sql.NullString{Valid: false}
	if db.NewRecord(t) {
		result := db.Create(&t)
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
