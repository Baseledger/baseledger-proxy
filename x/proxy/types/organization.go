package types

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/dbutil"
)

type Organization struct {
	Id               uuid.UUID
	OrganizationName string
}

func (t *Organization) Create() bool {
	db, err := dbutil.InitBaseledgerDBConnection()

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
		return rowsAffected > 0
	}

	return false
}
