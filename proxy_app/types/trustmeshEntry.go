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
	CommitmentState                      string
	TransactionHash                      string
	TrustmeshId                          uuid.UUID
	SorBusinessObjectId                  string // TODO: rename to remove SOR
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
	EthExitTxHash       string
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
		logger.Errorf("error when getting trustmesh from db %v\n", res.Error)
		return nil, res.Error
	}

	return &trustmesh, nil
}

func UpdateTrustmeshEthTxHash(id uuid.UUID, ethTxHash string) error {
	db := dbutil.Db.GetConn()

	res := db.Exec("update trustmeshes set eth_exit_tx_hash = ? where id = ?", ethTxHash, id.String())

	if res.Error != nil {
		logger.Errorf("Error when setting eth tx hash %v", res.Error.Error())
		return res.Error
	}

	if res.RowsAffected == 0 {
		logger.Infof("Trustmesh %v was not updated", id.String())
		return nil
	}

	logger.Infof("Eth tx hash %v set for trustmesh %v", ethTxHash, id)
	return nil
}

func GetTrustmeshEntryById(id uuid.UUID) (*TrustmeshEntry, error) {
	db := dbutil.Db.GetConn()
	var trustmeshEntry TrustmeshEntry
	res := db.Preload("OffchainProcessMessage").
		First(&trustmeshEntry, "id = ?", id.String())

	if res.Error != nil {
		logger.Errorf("error when getting trustmesh entry from db %v\n", res.Error)
		return nil, res.Error
	}

	return &trustmeshEntry, nil
}

func GetPendingTrustmeshEntries() ([]*TrustmeshEntry, error) {
	db := dbutil.Db.GetConn()

	entries := []*TrustmeshEntry{}

	res := db.Preload("OffchainProcessMessage").Raw("select * from trustmesh_entries te1 where entry_type = 'SuggestionReceived' and workstep_type = 'INITIAL' and not exists (select * from trustmesh_entries te2 where te1.baseledger_transaction_id = te2.referenced_baseledger_transaction_id)").Find(&entries)

	if res.Error != nil {
		logger.Errorf("Error when getting pending entries %v", res.Error.Error())
		return nil, res.Error
	}

	return entries, nil
}

func GetLatestTrustmeshEntryBasedOnTrustmeshId(trustmeshId string) (*TrustmeshEntry, error) {
	db := dbutil.Db.GetConn()

	latestEntry := &TrustmeshEntry{}
	res := db.Preload("OffchainProcessMessage").Raw("select * from trustmesh_entries where trustmesh_id = ? order by created_at desc limit 1", trustmeshId).Find(&latestEntry)

	if res.Error != nil {
		logger.Errorf("error when getting latest trustmesh entry from db %v\n", res.Error)
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		logger.Errorf("no trustmesh entries found for trustmesh %s", trustmeshId)
		return nil, nil
	}

	return latestEntry, nil
}

func GetLatestTrustmeshEntryBasedOnBboid(bboid string) (*TrustmeshEntry, error) {
	db := dbutil.Db.GetConn()

	latestEntry := &TrustmeshEntry{}
	res := db.Preload("OffchainProcessMessage").
		Raw("select * from trustmesh_entries where trustmesh_id = (select trustmesh_id from trustmesh_entries where baseledger_business_object_id = ? or referenced_baseledger_business_object_id = ? limit 1) order by created_at desc limit 1", bboid, bboid).
		Find(&latestEntry)
	if res.Error != nil {
		logger.Errorf("error when getting latest trustmesh entry from db %v\n", res.Error)
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		logger.Errorf("no trustmesh entries found for bboid %s", bboid)
		return nil, nil
	}

	return latestEntry, nil
}
