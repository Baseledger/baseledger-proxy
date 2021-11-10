package workgroups

import (
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/types"
	// gorm
)

type IWorkgroupClient interface {
	FindWorkgroup(workgroupId string) *types.Workgroup
	FindWorkgroupMember(workgroupId string, recipientId string) *types.WorkgroupMember
}

type PostgresWorkgroupClient struct {
}

func (client *PostgresWorkgroupClient) FindWorkgroup(workgroupId string) *types.Workgroup {
	var workgroup types.Workgroup
	dbError := dbutil.Db.GetConn().First(&workgroup, "id = ?", workgroupId).Error

	if dbError != nil {
		logger.Errorf("error trying to fetch workgroup with id %s\n", workgroupId)
		return nil
	}

	return &workgroup
}

func (client *PostgresWorkgroupClient) FindWorkgroupMember(workgroupId string, recipientId string) *types.WorkgroupMember {
	var member types.WorkgroupMember
	dbError := dbutil.Db.GetConn().First(&member, "workgroup_id = ? AND organization_id = ?", workgroupId, recipientId).Error

	if dbError != nil {
		logger.Errorf("error %s trying to fetch workgroup membership with workgroup id %s and organization with id %s\n", dbError, workgroupId, recipientId)
		return nil
	}

	return &member
}

// valid only until we go with the assumtion 1 recipient == 1 workgroup
func (client *PostgresWorkgroupClient) GetRecipientWorkgroupMember(recipientId string) *types.WorkgroupMember {
	var member types.WorkgroupMember
	dbError := dbutil.Db.GetConn().First(&member, "organization_id = ?", recipientId).Error

	if dbError != nil {
		logger.Errorf("error %s trying to fetch workgroup membership for organization with id %s\n", dbError, recipientId)
		return nil
	}

	return &member
}
