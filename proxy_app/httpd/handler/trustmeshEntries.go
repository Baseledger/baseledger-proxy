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
	NewObjectStatus string `json:"new_object_status"`
	Message         string `json:"message"`
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
			dto := &pendingTrustmeshEntryDto{
				TrustmeshEntryId: entry.Id.String(),
				WorkstepType:     entry.WorkstepType,
				Message:          entry.OffchainProcessMessage.StatusTextMessage,
				NewObjectStatus:  entry.OffchainProcessMessage.BaseledgerTransactionType,
			}

			dtos = append(dtos, *dto)
		}

		restutil.Render(dtos, 200, c)
	}
}
