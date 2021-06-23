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
	Id                                   uuid.UUID
	SenderId                             uuid.UUID
	ReceiverId                           uuid.UUID
	Topic                                string
	ReferencedOffchainProcessMessageId   string
	BaseledgerSyncTreeJson               string
	WorkstepType                         string
	BusinessObjectProof                  string
	BusinessObjectType                   string
	TendermintTransactionIdOfStoredProof string
	BaseledgerTransactionIdOfStoredProof string
	BaseledgerBusinessObjectId           string
	ReferencedBaseledgerBusinessObjectId string
	StatusTextMessage                    string
	BaseledgerTransactionType            string
	ReferencedBaseledgerTransactionId    string
	EntryType                            string
}

// TODO rename after clean up
type SynchronizationRequest struct {
	WorkgroupId                          uuid.UUID
	Recipient                            string
	WorkstepType                         string
	BusinessObjectType                   string
	BaseledgerBusinessObjectId           string
	BusinessObjectJson                   string
	ReferencedBaseledgerBusinessObjectId string
	ReferencedBaseledgerTransactionId    string
}

type SynchronizationFeedback struct {
	WorkgroupId                                uuid.UUID
	BaseledgerProvenBusinessObjectJson         string
	BusinessObjectType                         string
	Recipient                                  string
	Approved                                   bool
	BaseledgerBusinessObjectIdOfApprovedObject string
	HashOfObjectToApprove                      string
	OriginalBaseledgerTransactionId            string
	OriginalOffchainProcessMessageId           string
	FeedbackMessage                            string
}

type BaseledgerTransactionPayload struct {
	SenderId                             string
	TransactionType                      string
	OffchainMessageId                    string
	ReferencedOffchainMessageId          string
	ReferencedBaseledgerTransactionId    string
	BaseledgerTransactionID              string
	Proof                                string
	BaseledgerBusinessObjectId           string
	ReferencedBaseledgerBusinessObjectId string
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

func GetOffchainMsgById(id uuid.UUID) (msg *OffchainProcessMessage, err error) {
	db, err := dbutil.InitBaseledgerDBConnection()

	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
		return nil, err
	}

	var offchainMsg OffchainProcessMessage
	res := db.First(&offchainMsg, "id = ?", id.String())

	if res.Error != nil {
		fmt.Printf("error when getting offchain msg from db %v\n", err)
		return nil, res.Error
	}

	return &offchainMsg, nil
}

// all other types for hasing, privacy, off-chain messaging
