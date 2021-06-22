package workgroups

import (
	"fmt"

	"github.com/unibrightio/baseledger/dbutil"
	"github.com/unibrightio/baseledger/x/proxy/types"
	// gorm
)

type IWorkgroupClient interface {
	FindWorkgroup(workgroupId string) *types.Workgroup
	FindWorkgroupMember(workgroupId string, recipientId string) *types.WorkgroupMember
}

type PostgresWorkgroupClient struct {
}

func (client *PostgresWorkgroupClient) FindWorkgroup(workgroupId string) *types.Workgroup {
	db, err := dbutil.InitBaseledgerDBConnection()

	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
		return nil
	}

	var workgroup types.Workgroup
	dbError := db.First(&workgroup, "id = ?", workgroupId).Error

	if dbError != nil {
		fmt.Printf("error trying to fetch workgroup with id %s\n", workgroupId)
		return nil
	}

	return &workgroup
}

func (client *PostgresWorkgroupClient) FindWorkgroupMember(workgroupId string, recipientId string) *types.WorkgroupMember {
	db, err := dbutil.InitBaseledgerDBConnection()

	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
		return nil
	}

	var member types.WorkgroupMember
	dbError := db.First(&member, "workgroup_id = ? AND organization_id = ?", workgroupId, recipientId).Error

	if dbError != nil {
		fmt.Printf("error trying to fetch workgroup membership with workgroup id %s and organization with id %s\n", recipientId, recipientId)
		return nil
	}

	return &member
}
