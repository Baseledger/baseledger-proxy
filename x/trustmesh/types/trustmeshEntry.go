package types

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
)

type TrustmeshEntry struct {
	TendermintBlockId                    string
	TendermintTransactionId              string
	TendermintTransactionTimestamp       string
	Sender                               string
	Receiver                             string
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
}

func (t *TrustmeshEntry) Create() bool {
	args := "host=localhost user=baseledger password=ub123 dbname=baseledger sslmode=disable"

	// TODO: use kyle's package
	db, err := gorm.Open("postgres", args)

	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
		return false
	}

	if db.NewRecord(t) {
		result := db.Create(&t)
		rowsAffected := result.RowsAffected
		errors := result.GetErrors()
		if len(errors) > 0 {
			fmt.Printf("errors while creating new entry %v\n", errors)
			return false
		}
		fmt.Printf("ROWS AFFECTED %v\n", rowsAffected)
		return rowsAffected > 0
	}

	return false
}
