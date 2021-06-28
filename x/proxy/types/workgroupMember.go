package types

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/dbutil"
)

type WorkgroupMember struct {
	Id                   uuid.UUID
	WorkgroupId          string
	OrganizationId       string
	OrganizationEndpoint string // localhost:4223
	OrganizationToken    string // testToken1
}

func (t *WorkgroupMember) Create() bool {
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
