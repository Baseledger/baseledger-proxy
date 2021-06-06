package types

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	"github.com/spf13/viper"
)

const defaultTransactionStatus = "UNCOMMITTED"

type TrustmeshEntry struct {
	TendermintBlockId                    sql.NullString
	TendermintTransactionId              string
	TendermintTransactionTimestamp       sql.NullString
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
	TransactionStatus                    string
}

func (t *TrustmeshEntry) Create() bool {
	dbHost, _ := viper.Get("DB_HOST").(string)
	dbPwd, _ := viper.Get("DB_UB_PWD").(string)
	sslMode, _ := viper.Get("DB_SSLMODE").(string)
	dbUser, _ := viper.Get("DB_BASELEDGER_USER").(string)
	dbName, _ := viper.Get("DB_BASELEDGER_NAME").(string)

	args := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbUser,
		dbPwd,
		dbName,
		sslMode,
	)

	db, err := gorm.Open("postgres", args)

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
