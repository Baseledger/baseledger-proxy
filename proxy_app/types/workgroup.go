package types

import (
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
)

type Workgroup struct {
	Id            uuid.UUID
	WorkgroupName string
	PrivatizeKey  string
}

func (t *Workgroup) Create() bool {
	result := dbutil.Db.GetConn().Create(&t)
	rowsAffected := result.RowsAffected
	errors := result.GetErrors()
	if len(errors) > 0 {
		logger.Errorf("errors while creating new entry %v\n", errors)
		return false
	}
	return rowsAffected > 0
}

func (t *Workgroup) Delete() bool {
	result := dbutil.Db.GetConn().Delete(&t)
	rowsAffected := result.RowsAffected
	errors := result.GetErrors()

	if len(errors) > 0 {
		logger.Errorf("errors while deleting new entry %v\n", errors)
		return false
	}

	return rowsAffected > 0
}
