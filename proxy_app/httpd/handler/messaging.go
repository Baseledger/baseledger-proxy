package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/messaging"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/workgroups"
)

type sendOffchainMessageRequest struct {
	WorkgroupId string `json:"workgroup_id"`
	RecipientId string `json:"recipient_id"`
	Payload     string `json:"payload"`
}

func SendOffchainMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Infof("sendOffchainMessageHandler initiated\n")
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		req := &sendOffchainMessageRequest{}
		err = json.Unmarshal(buf, &req)
		if err != nil {
			restutil.RenderError(err.Error(), 422, c)
			return
		}

		workgroupClient := &workgroups.PostgresWorkgroupClient{}
		logger.Infof("trying to find workgroup member\n")
		workgroupMembership := workgroupClient.FindWorkgroupMember(req.WorkgroupId, req.RecipientId)

		if workgroupMembership == nil {
			restutil.RenderError("failed to find a workgroup member", 404, c)
			return
		}

		logger.Infof("trying to message on url: %s with token: %s\n", workgroupMembership.OrganizationEndpoint, workgroupMembership.OrganizationToken)
		messagingClient := &messaging.NatsMessagingClient{}
		messagingClient.SendMessage([]byte(req.Payload), workgroupMembership.OrganizationEndpoint, workgroupMembership.OrganizationToken)
		restutil.Render(nil, 200, c)
	}
}
