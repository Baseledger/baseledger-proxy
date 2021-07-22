package handler

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

type trustmeshEntryDto struct {
	TendermintBlockId                    string
	TendermintTransactionId              uuid.UUID
	TendermintTransactionTimestamp       time.Time
	EntryType                            string
	SenderOrgId                          uuid.UUID
	ReceiverOrgId                        uuid.UUID
	WorkgroupId                          uuid.UUID
	WorkstepType                         string
	BaseledgerTransactionType            string
	BaseledgerTransactionId              uuid.UUID
	ReferencedBaseledgerTransactionId    uuid.UUID
	BusinessObjectType                   string
	BaseledgerBusinessObjectId           uuid.UUID
	ReferencedBaseledgerBusinessObjectId uuid.UUID
	OffchainProcessMessageId             uuid.UUID
	ReferencedProcessMessageId           uuid.UUID
	CommitmentState                      string
	TransactionHash                      string
	TrustmeshId                          uuid.UUID
}

type trustmeshDto struct {
	Id                  uuid.UUID
	CreatedAt           time.Time
	StartTime           time.Time
	EndTime             time.Time
	Participants        string
	BusinessObjectTypes string
	Finalized           bool
	ContainsRejections  bool
	Entries             []trustmeshEntryDto
}

func GetTrustmeshesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var trustmeshes []types.Trustmesh
		db := dbutil.Db.GetConn().Order("trustmeshes.created_at ASC")
		// preload seems good enough for now, revisit if it turns out to be performance bottleneck
		dbutil.Paginate(c, db, &types.Trustmesh{}).Preload("Entries").Find(&trustmeshes)

		var trustmeshesDtos []trustmeshDto

		for i := 0; i < len(trustmeshes); i++ {
			trustmeshesDtos = append(trustmeshesDtos, *processTrustmesh(&trustmeshes[i]))
		}

		restutil.Render(trustmeshesDtos, 200, c)
	}
}

func processTrustmesh(trustmesh *types.Trustmesh) *trustmeshDto {
	if len(trustmesh.Entries) == 0 {
		return nil
	}
	trustmeshDto := &trustmeshDto{}
	var entriesDto []trustmeshEntryDto

	startTime := trustmesh.Entries[0].TendermintTransactionTimestamp
	endTime := trustmesh.Entries[0].TendermintTransactionTimestamp
	senders := ""
	sendersMap := make(map[string]int)
	receivers := ""
	receiversMap := make(map[string]int)
	businessObjectTypes := ""
	businessObjectTypesMap := make(map[string]int)
	finalized := false
	containsRejection := false

	for _, entry := range trustmesh.Entries {
		startTime = getBeforeTime(startTime, entry.TendermintTransactionTimestamp)
		endTime = getAfterTime(endTime, entry.TendermintTransactionTimestamp)

		appendDistinct(sendersMap, entry.SenderOrgId.String(), &senders)
		appendDistinct(receiversMap, entry.ReceiverOrgId.String(), &receivers)
		appendDistinct(businessObjectTypesMap, entry.BusinessObjectType, &businessObjectTypes)

		if entry.WorkstepType == "FinalWorkstep" && !finalized {
			finalized = true
		}

		if entry.BaseledgerTransactionType == "Reject" && !containsRejection {
			containsRejection = true
		}

		entryDto := processTrustmeshEntry(entry)
		entriesDto = append(entriesDto, *entryDto)
	}

	trustmeshDto.StartTime = startTime.Time
	trustmeshDto.EndTime = endTime.Time
	trustmeshDto.Participants = senders + ", " + receivers
	trustmeshDto.BusinessObjectTypes = businessObjectTypes
	trustmeshDto.Finalized = finalized
	trustmeshDto.ContainsRejections = containsRejection
	trustmeshDto.Entries = entriesDto

	return trustmeshDto
}

func processTrustmeshEntry(trustmeshEntry types.TrustmeshEntry) *trustmeshEntryDto {
	return &trustmeshEntryDto{
		TendermintBlockId:                    trustmeshEntry.TendermintBlockId.String,
		TendermintTransactionId:              trustmeshEntry.TendermintTransactionId,
		TendermintTransactionTimestamp:       trustmeshEntry.TendermintTransactionTimestamp.Time,
		EntryType:                            trustmeshEntry.EntryType,
		SenderOrgId:                          trustmeshEntry.SenderOrgId,
		ReceiverOrgId:                        trustmeshEntry.ReceiverOrgId,
		WorkgroupId:                          trustmeshEntry.WorkgroupId,
		WorkstepType:                         trustmeshEntry.WorkstepType,
		BaseledgerTransactionType:            trustmeshEntry.BaseledgerTransactionType,
		BaseledgerTransactionId:              trustmeshEntry.BaseledgerTransactionId,
		ReferencedBaseledgerTransactionId:    trustmeshEntry.ReferencedBaseledgerTransactionId,
		BusinessObjectType:                   trustmeshEntry.BusinessObjectType,
		BaseledgerBusinessObjectId:           trustmeshEntry.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: trustmeshEntry.ReferencedBaseledgerBusinessObjectId,
		OffchainProcessMessageId:             trustmeshEntry.OffchainProcessMessageId,
		ReferencedProcessMessageId:           trustmeshEntry.ReferencedProcessMessageId,
		CommitmentState:                      trustmeshEntry.CommitmentState,
		TransactionHash:                      trustmeshEntry.TransactionHash,
		TrustmeshId:                          trustmeshEntry.TrustmeshId,
	}
}

func getSeparator(str string) string {
	if str == "" {
		return ""
	} else {
		return ", "
	}
}

func getBeforeTime(first sql.NullTime, second sql.NullTime) sql.NullTime {
	if !first.Valid {
		return second
	}

	if !second.Valid {
		return first
	}

	if first.Time.Before(second.Time) {
		return first
	}

	return second
}

func getAfterTime(first sql.NullTime, second sql.NullTime) sql.NullTime {
	if !first.Valid {
		return second
	}

	if !second.Valid {
		return first
	}

	if first.Time.Before(second.Time) {
		return second
	}

	return first
}

func appendDistinct(itemsMap map[string]int, newItem string, acc *string) {
	if itemsMap[newItem] == 0 {
		itemsMap[newItem] = 1
		*acc = *acc + getSeparator(*acc) + newItem
	}
}
