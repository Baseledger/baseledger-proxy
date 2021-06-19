package types

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres

	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/dbutil"
)

//Put here our Types needed for the proxy elements?
type OffchainProcessMessageReferenceType string

type OffchainProcessMessage struct {
	Id                                 uuid.UUID
	SenderId                           string
	ReceiverId                         string
	Topic                              string
	ReferencedOffchainProcessMessageId string
	// todo replace string with proper type?
	BusinessObject                       string
	WorkstepType                         string
	Hash                                 string
	TendermintTransactionIdOfStoredProof string
	BaseledgerTransactionIdOfStoredProof string
	BaseledgerBusinessObjectId           string
	ReferencedBaseledgerBusinessObjectId string
	StatusTextMessage                    string
}

// TODO rename after clean up
type SynchronizationRequest struct {
	WorkgroupId                          string
	Recipient                            string
	WorkstepType                         string
	BusinessObjectType                   string
	BaseledgerBusinessObjectId           string
	BusinessObject                       string
	ReferencedBaseledgerBusinessObjectId string
	ReferencedBaseledgerTransactionId    string
}

type BaseledgerTransactionPayload struct {
	PhonebookIdentifier                  string `json:"phonebookIdentifier"`
	TransactionType                      string `json:"baseledgerTransactionType"`
	OffchainMessageId                    string `json:"offchainMessageId"`
	ReferencedOffchainMessageId          string `json:"referencedOffchainMessageId"`
	ReferencedBaseledgerTransactionId    string `json:"referencedBaseledgerTransactionId"`
	BaseledgerTransactionID              string `json:"baseledgerTransactionID"`
	Proof                                string `json:"proof"`
	BaseledgerBusinessObjectId           string `json:"baseledgerBusinessObjectID"`
	ReferencedBaseledgerBusinessObjectId string `json:"referencedBaseledgerBusinessObjectID"`
}

func (o *OffchainProcessMessage) Create() bool {
	db, err := dbutil.InitBaseledgerDBConnection()

	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
		return false
	}

	if db.NewRecord(o) {
		result := db.Create(&o)
		rowsAffected := result.RowsAffected
		errors := result.GetErrors()
		if len(errors) > 0 {
			fmt.Printf("errors while creating new offchain process msg entry %v\n", errors)
			return false
		}
		return rowsAffected > 0
	}

	return false
}

// all other types for hasing, privacy, off-chain messaging
