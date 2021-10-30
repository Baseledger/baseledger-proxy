package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

type pendingTrustmeshEntryDto struct {
	TrustmeshEntryId string `json:"trustmesh_entry_id"`
	WorkstepType     string `json:"workstep_type"`
	// BusinessObjectJsonPayload string `json:"business_object_json_payload"`
	NewObjectStatus int    `json:"new_object_status"`
	Message         string `json:"message"`
}

type relatedTrustmeshEntryDto struct {
	TrustmeshEntryId string `json:"trustmesh_entry_id"`
	WorkstepType     string `json:"workstep_type"`
	// BusinessObjectJsonPayload string `json:"business_object_json_payload"`
	NewObjectStatus int `json:"new_object_status"`
	// Origin string `json:"origin"`
	Message string `json:"message"`
}

func GetPendingTrustmeshEntriesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := types.GetPendingTrustmeshEntries()

		if err != nil {
			restutil.RenderError("error when fetching pending entries", 400, c)
			return
		}

		dtos := []pendingTrustmeshEntryDto{}

		for _, entry := range res {
			status := 0
			if entry.OffchainProcessMessage.BaseledgerTransactionType == "Approve" {
				status = 1
			}
			dto := &pendingTrustmeshEntryDto{
				TrustmeshEntryId: entry.Id.String(),
				WorkstepType:     entry.WorkstepType,
				Message:          entry.OffchainProcessMessage.StatusTextMessage,
				NewObjectStatus:  status,
			}

			dtos = append(dtos, *dto)
		}

		restutil.Render(dtos, 200, c)
	}
}

func GetRelatedTrustmesEntryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		entryId := c.Param("id")
		res, err := types.GetFirstRelatedTrustmeshEntry(entryId)

		if err != nil {
			restutil.RenderError("error when fetching related entries", 400, c)
			return
		}

		status := 0
		if res.OffchainProcessMessage.BaseledgerTransactionType == "Approve" {
			status = 1
		}
		dto := &relatedTrustmeshEntryDto{
			TrustmeshEntryId: res.Id.String(),
			WorkstepType:     res.WorkstepType,
			Message:          res.OffchainProcessMessage.StatusTextMessage,
			NewObjectStatus:  status,
		}

		restutil.Render(dto, 200, c)
	}
}
