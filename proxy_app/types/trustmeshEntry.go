package types

import (
	"database/sql"
	"time"

	uuid "github.com/kthomas/go.uuid"
	common "github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
)

type TrustmeshEntry struct {
	Id                                   uuid.UUID
	TendermintBlockId                    sql.NullString
	TendermintTransactionId              uuid.UUID
	TendermintTransactionTimestamp       sql.NullTime
	EntryType                            string
	SenderOrgId                          uuid.UUID
	ReceiverOrgId                        uuid.UUID
	SenderOrg                            Organization
	ReceiverOrg                          Organization
	WorkgroupId                          uuid.UUID
	Workgroup                            Workgroup
	WorkstepType                         string
	BaseledgerTransactionType            string
	BaseledgerTransactionId              uuid.UUID
	ReferencedBaseledgerTransactionId    uuid.UUID
	BusinessObjectType                   string
	BaseledgerBusinessObjectId           string
	ReferencedBaseledgerBusinessObjectId string
	OffchainProcessMessageId             uuid.UUID
	OffchainProcessMessage               OffchainProcessMessage
	ReferencedProcessMessageId           uuid.UUID
	CommitmentState                      string
	TransactionHash                      string
	TrustmeshId                          uuid.UUID
	SorBusinessObjectId                  string
}

type Trustmesh struct {
	Id                  uuid.UUID
	CreatedAt           time.Time
	StartTime           time.Time
	EndTime             time.Time
	Participants        string
	BusinessObjectTypes string
	Finalized           bool
	ContainsRejections  bool
	Entries             []TrustmeshEntry
}

func (t *TrustmeshEntry) Create() bool {
	t.CommitmentState = common.UncommittedCommitmentState
	t.TendermintBlockId = sql.NullString{Valid: false}
	t.TendermintTransactionTimestamp = sql.NullTime{Valid: false}
	if dbutil.Db.GetConn().NewRecord(t) {
		result := dbutil.Db.GetConn().Create(&t)
		rowsAffected := result.RowsAffected
		errors := result.GetErrors()
		if len(errors) > 0 {
			logger.Errorf("errors while creating new entry %v\n", errors)
			return false
		}
		return rowsAffected > 0
	}

	return false
}

func GetTrustmeshById(id uuid.UUID) (*Trustmesh, error) {
	db := dbutil.Db.GetConn()
	var trustmesh Trustmesh
	res := db.Preload("Entries").
		Preload("Entries.SenderOrg").
		Preload("Entries.ReceiverOrg").
		Preload("Entries.Workgroup").
		Preload("Entries.OffchainProcessMessage").
		First(&trustmesh, "id = ?", id.String())

	if res.Error != nil {
		logger.Errorf("error when getting offchain msg from db %v\n", res.Error)
		return nil, res.Error
	}

	return &trustmesh, nil
}

func GetPendingTrustmeshEntries() ([]*TrustmeshEntry, error) {
	db := dbutil.Db.GetConn()

	entries := []*TrustmeshEntry{}

	res := db.Preload("OffchainProcessMessage").Raw("select * from trustmesh_entries te1 where not exists (select * from trustmesh_entries te2 where te1.baseledger_transaction_id = te2.referenced_baseledger_transaction_id)").Find(&entries)

	if res.Error != nil {
		logger.Errorf("Error when getting pending entries %v", res.Error.Error())
		return nil, res.Error
	}

	return entries, nil
}

func GetFirstRelatedTrustmeshEntry(id string) (*TrustmeshEntry, error) {
	db := dbutil.Db.GetConn()

	entry := &TrustmeshEntry{}
	res := db.Preload("OffchainProcessMessage").First(&entry, "id = ?", id)
	if res.Error != nil {
		logger.Errorf("error when getting trustmesh entry from db %v\n", res.Error)
		return nil, res.Error
	}

	relatedEntry := &TrustmeshEntry{}
	res = db.Preload("OffchainProcessMessage").Raw("select * from trustmesh_entries where referenced_baseledger_transaction_id = ? order by created_at desc limit 1", entry.BaseledgerTransactionId).Find(&relatedEntry)

	if res.Error != nil {
		logger.Errorf("error when getting related trustmesh entry from db %v\n", res.Error)
		return entry, nil
	}

	if res.RowsAffected == 0 {
		return entry, nil
	}

	return relatedEntry, nil
}
