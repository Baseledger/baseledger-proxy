package workgroups

import (
	uuid "github.com/kthomas/go.uuid"
	// gorm
)

type IWorkgroupClient interface {
	FindWorkgroup(workgroupId string) *workgroupMock
	FindRecipientMessagingEndpoint(recipientId string) string
	FindRecipientMessagingToken(recipientId string) string
}

type PostgresWorkgroupClient struct {
}

type workgroupMock struct {
	BaselineWorkgroupID string
	Description         string
	PrivatizeKey        string
}

func (client *PostgresWorkgroupClient) FindWorkgroup(workgroupId string) *workgroupMock {
	newUuid, _ := uuid.NewV4()
	return &workgroupMock{
		BaselineWorkgroupID: newUuid.String(),
		Description:         "Mocked workgroup",
		PrivatizeKey:        "0c2e08bc9249fb42568e5a478e9af87a208471c46211a08f3ad9f0c5dbf57314",
	}
}

func (client *PostgresWorkgroupClient) FindRecipientMessagingEndpoint(recipientId string) string {
	return "nil"
}

func (client *PostgresWorkgroupClient) FindRecipientMessagingToken(recipientId string) string {
	return "nil"
}
