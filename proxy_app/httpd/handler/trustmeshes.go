package handler

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

type trustmeshEntryDto struct {
	TendermintBlockId                    string
	TendermintTransactionId              string
	TendermintTransactionTimestamp       time.Time
	EntryType                            string
	SenderOrgId                          string
	ReceiverOrgId                        string
	WorkgroupId                          string
	WorkstepType                         string
	BaseledgerTransactionType            string
	BaseledgerTransactionId              string
	ReferencedBaseledgerTransactionId    string
	BusinessObjectType                   string
	BaseledgerBusinessObjectId           string
	ReferencedBaseledgerBusinessObjectId string
	OffchainProcessMessageId             string
	ReferencedProcessMessageId           string
	CommitmentState                      string
	TransactionHash                      string
	TrustmeshId                          string
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

func GetTrustmeshHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		trustmeshIdParam := c.Param("id")

		trustmeshId, err := uuid.FromString(trustmeshIdParam)
		if err != nil {
			logger.Errorf("Trustmesh id param %v is in wrong format", trustmeshIdParam)
			restutil.RenderError("trustmesh id in wrong format", 400, c)
			return
		}

		trustmesh, err := types.GetTrustmeshById(trustmeshId)
		if err != nil {
			logger.Errorf("Trustmesh with id %v not found", trustmeshId)
			restutil.RenderError("trustmesh not found", 400, c)
			return
		}

		trustmeshDto := processTrustmesh(trustmesh)
		restutil.Render(trustmeshDto, 200, c)
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

	trustmeshDto.Id = trustmesh.Id
	trustmeshDto.CreatedAt = trustmesh.CreatedAt
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
		TendermintTransactionId:              uuidToString(trustmeshEntry.TendermintTransactionId),
		TendermintTransactionTimestamp:       trustmeshEntry.TendermintTransactionTimestamp.Time,
		EntryType:                            trustmeshEntry.EntryType,
		SenderOrgId:                          uuidToString(trustmeshEntry.SenderOrgId),
		ReceiverOrgId:                        uuidToString(trustmeshEntry.ReceiverOrgId),
		WorkgroupId:                          uuidToString(trustmeshEntry.WorkgroupId),
		WorkstepType:                         trustmeshEntry.WorkstepType,
		BaseledgerTransactionType:            trustmeshEntry.BaseledgerTransactionType,
		BaseledgerTransactionId:              uuidToString(trustmeshEntry.BaseledgerTransactionId),
		ReferencedBaseledgerTransactionId:    uuidToString(trustmeshEntry.ReferencedBaseledgerTransactionId),
		BusinessObjectType:                   trustmeshEntry.BusinessObjectType,
		BaseledgerBusinessObjectId:           trustmeshEntry.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: trustmeshEntry.ReferencedBaseledgerBusinessObjectId,
		OffchainProcessMessageId:             uuidToString(trustmeshEntry.OffchainProcessMessageId),
		ReferencedProcessMessageId:           uuidToString(trustmeshEntry.ReferencedProcessMessageId),
		CommitmentState:                      trustmeshEntry.CommitmentState,
		TransactionHash:                      trustmeshEntry.TransactionHash,
		TrustmeshId:                          uuidToString(trustmeshEntry.TrustmeshId),
	}
}

func uuidToString(id uuid.UUID) string {
	if id == uuid.Nil {
		return ""
	}
	return id.String()
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
