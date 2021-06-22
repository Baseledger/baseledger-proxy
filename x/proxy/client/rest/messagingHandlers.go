package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	"github.com/unibrightio/baseledger/x/proxy/messaging"
	"github.com/unibrightio/baseledger/x/proxy/workgroups"
)

type sendOffchainMessageRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	WorkgroupId string       `json:"workgroup_id"`
	RecipientId string       `json:"recipient_id"`
	Payload     string       `json:"payload"`
}

func sendOffchainMessageHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("sendOffchainMessageHandler initiated\n")

		fmt.Printf("trying to parse request\n")
		req := parseMessageRequest(w, r, clientCtx)
		fmt.Printf("Request parsed succesfully %s %s\n", req.WorkgroupId, req.RecipientId)

		workgroupClient := &workgroups.PostgresWorkgroupClient{}
		fmt.Printf("trying to find workgroup member\n")
		workgroupMembership := workgroupClient.FindWorkgroupMember(req.WorkgroupId, req.RecipientId)

		if workgroupMembership == nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to find a workgroup member")
			return
		}

		fmt.Printf("trying to message on url: %s with token: %s\n", workgroupMembership.OrganizationEndpoint, workgroupMembership.OrganizationToken)
		messagingClient := &messaging.NatsMessagingClient{}
		messagingClient.SendMessage(req.Payload, workgroupMembership.OrganizationEndpoint, workgroupMembership.OrganizationToken)

		w.WriteHeader(http.StatusOK)
	}
}

// TODO: common function for all rest handlers
func parseMessageRequest(w http.ResponseWriter, r *http.Request, clientCtx client.Context) *sendOffchainMessageRequest {
	var req sendOffchainMessageRequest
	if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
		return nil
	}

	baseReq := req.BaseReq.Sanitize()
	if !baseReq.ValidateBasic(w) {
		rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
		return nil
	}

	return &req
}
