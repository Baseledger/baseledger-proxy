package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/synctree"
	"github.com/unibrightio/proxy-api/types"
)

type newWorkflowDto struct {
	WorkflowId                 string `json:"workflow_id"` // Id of the trustmesh
	WorkstepId                 string `json:"workstep_id"` // Id of the latest trustmesh entry id
	WorkstepType               string `json:"workstep_type"`
	BaseledgerBusinessObjectId string `json:"baseledger_business_object_id"`
	BusinessObjectJsonPayload  string `json:"business_object_json_payload"`
}

type latestTrustmeshEntryDto struct {
	WorkflowId                 string `json:"workflow_id"` // Id of the trustmesh
	WorkstepId                 string `json:"workstep_id"` // Id of the latest trustmesh entry id
	WorkstepType               string `json:"workstep_type"`
	BaseledgerBusinessObjectId string `json:"baseledger_business_object_id"`
	BusinessObjectJsonPayload  string `json:"business_object_json_payload"`
	Approved                   bool   `json:"approved"`
}

func GetNewWorkflowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := types.GetPendingTrustmeshEntries()

		if err != nil {
			restutil.RenderError("error when fetching pending entries", 400, c)
			return
		}

		dtos := []newWorkflowDto{}

		for _, entry := range res {
			syncTree := &synctree.BaseledgerSyncTree{}
			json.Unmarshal([]byte(entry.OffchainProcessMessage.BaseledgerSyncTreeJson), &syncTree)

			boJson := synctree.GetBusinessObjectJson(*syncTree)
			dto := &newWorkflowDto{
				WorkflowId:                 entry.TrustmeshId.String(),
				WorkstepId:                 entry.Id.String(),
				WorkstepType:               entry.WorkstepType,
				BaseledgerBusinessObjectId: entry.BaseledgerBusinessObjectId,
				BusinessObjectJsonPayload:  boJson,
			}

			dtos = append(dtos, *dto)
		}

		restutil.Render(dtos, 200, c)
	}
}

func GetLatestWorkflowStateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		baseledgerBusinessObjectId := c.Param("bo_id")

		entry, err := types.GetLatestTrustmeshEntryBasedOnBusinessObjectId(baseledgerBusinessObjectId)

		if err != nil {
			restutil.RenderError("error when fetching latest worfkflow entry", 400, c)
			return
		}

		approved := false
		if entry.OffchainProcessMessage.BaseledgerTransactionType == "Approve" {
			approved = true
		}
		syncTree := &synctree.BaseledgerSyncTree{}
		json.Unmarshal([]byte(entry.OffchainProcessMessage.BaseledgerSyncTreeJson), &syncTree)

		boJson := synctree.GetBusinessObjectJson(*syncTree)
		dto := &latestTrustmeshEntryDto{
			WorkflowId:                 entry.TrustmeshId.String(),
			WorkstepId:                 entry.Id.String(),
			WorkstepType:               entry.WorkstepType,
			BaseledgerBusinessObjectId: entry.BaseledgerBusinessObjectId,
			BusinessObjectJsonPayload:  boJson,
			Approved:                   approved,
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
