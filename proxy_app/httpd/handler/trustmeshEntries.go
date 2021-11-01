package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/synctree"
	"github.com/unibrightio/proxy-api/types"
)

type pendingTrustmeshEntryDto struct {
	TrustmeshEntryId          string `json:"trustmesh_entry_id"`
	WorkstepType              string `json:"workstep_type"`
	BusinessObjectJsonPayload string `json:"business_object_json_payload"`
}

type relatedTrustmeshEntryDto struct {
	TrustmeshEntryId          string `json:"trustmesh_entry_id"`
	WorkstepType              string `json:"workstep_type"`
	BusinessObjectJsonPayload string `json:"business_object_json_payload"`
	NewObjectStatus           int    `json:"new_object_status"`
	Origin                    string `json:"origin"`
	Message                   string `json:"message"`
}

func FetchNewSuggestions() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := types.GetPendingTrustmeshEntries()

		if err != nil {
			restutil.RenderError("error when fetching pending entries", 400, c)
			return
		}

		dtos := []pendingTrustmeshEntryDto{}

		for _, entry := range res {
			syncTree := &synctree.BaseledgerSyncTree{}
			json.Unmarshal([]byte(entry.OffchainProcessMessage.BaseledgerSyncTreeJson), &syncTree)

			boJson := synctree.GetBusinessObjectJson(*syncTree)
			dto := &pendingTrustmeshEntryDto{
				TrustmeshEntryId:          entry.Id.String(),
				WorkstepType:              entry.WorkstepType,
				BusinessObjectJsonPayload: boJson,
			}

			dtos = append(dtos, *dto)
		}

		restutil.Render(dtos, 200, c)
	}
}

func FetchTrustmeshEntryUpdates() gin.HandlerFunc {
	return func(c *gin.Context) {
		entryId := c.Param("id")
		res, err := types.GetFirstSubsequentTrustmeshEntry(entryId)

		if err != nil {
			restutil.RenderError("error when fetching related entries", 400, c)
			return
		}

		status := 0
		if res.OffchainProcessMessage.BaseledgerTransactionType == "Approve" {
			status = 1
		}
		syncTree := &synctree.BaseledgerSyncTree{}
		json.Unmarshal([]byte(res.OffchainProcessMessage.BaseledgerSyncTreeJson), &syncTree)

		boJson := synctree.GetBusinessObjectJson(*syncTree)
		dto := &relatedTrustmeshEntryDto{
			TrustmeshEntryId:          res.Id.String(),
			WorkstepType:              res.WorkstepType,
			Message:                   res.OffchainProcessMessage.StatusTextMessage,
			BusinessObjectJsonPayload: boJson,
			NewObjectStatus:           status,
			Origin:                    getEntryOrigin(res),
		}

		restutil.Render(dto, 200, c)
	}
}

func getEntryOrigin(entry *types.TrustmeshEntry) string {
	isInitiatorProxy := entry.CommitmentState == common.InvalidCommitmentState
	isEntryComingFromOtherParty := entry.EntryType == common.SuggestionReceivedTrustmeshEntryType || entry.EntryType == common.FeedbackReceivedTrustmeshEntryType

	if isInitiatorProxy || !isEntryComingFromOtherParty {
		return ""
	}

	return entry.SenderOrgId.String()
}
