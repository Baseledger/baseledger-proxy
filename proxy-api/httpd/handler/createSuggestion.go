package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/proxyutil"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

type createInitialSuggestionRequest struct {
	BaseReq                              restutil.BaseReq `json:"base_req"`
	WorkgroupId                          string           `json:"workgroup_id"`
	Recipient                            string           `json:"recipient"`
	WorkstepType                         string           `json:"workstep_type"`
	BusinessObjectType                   string           `json:"business_object_type"`
	BaseledgerBusinessObjectId           string           `json:"baseledger_business_object_id"`
	BusinessObjectJson                   string           `json:"business_object_json"`
	ReferencedBaseledgerBusinessObjectId string           `json:"referenced_baseledger_business_object_id"`
	ReferencedBaseledgerTransactionId    string           `json:"referenced_baseledger_transaction_id"`
}

func CreateInitialSuggestionRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		req := &createInitialSuggestionRequest{}
		err = json.Unmarshal(buf, &req)
		if err != nil {
			restutil.RenderError(err.Error(), 422, c)
			return
		}

		syncReq := newSynchronizationRequest(*req)

		hash := proxyutil.CreateHashFromBusinessObject(req.BusinessObjectJson)
		transactionId := uuid.NewV4()

		offchainMsg := createSuggestOffchainMessage(*req, transactionId, hash)

		if !offchainMsg.Create() {
			logger.Errorf("error when creating new offchain msg entry")
			restutil.RenderError("error when creating new offchain msg entry", 500, c)
			return
		}

		payload := proxyutil.CreateBaseledgerTransactionPayload(syncReq, &offchainMsg)

		signAndBroadcastPayload := restutil.SignAndBroadcastPayload{
			BaseReq:       req.BaseReq,
			TransactionId: transactionId.String(),
			Payload:       payload,
		}

		jsonValue, err := json.Marshal(signAndBroadcastPayload)

		if err != nil {
			logger.Error("Error marshaling sign and broadcast json")
			restutil.RenderError("Error marshaling sign and broadcast json", 500, c)
			return
		}

		resp, err := http.Post("http://localhost:1317/proxy/signAndBroadcast", "application/json", bytes.NewBuffer(jsonValue))

		if err != nil {
			logger.Errorf("error while sending feedback request %v\n", err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("error while reading sign and broadcast transaction response %v\n", err.Error())
			restutil.RenderError("error while reading sign and broadcast transaction response", 500, c)
			return
		}

		transactionHash := string(body)

		trustmeshEntry := createSuggestionSentTrustmeshEntry(*req, transactionId, offchainMsg, transactionHash)

		if !trustmeshEntry.Create() {
			logger.Errorf("error when creating new trustmesh entry")
			restutil.RenderError("error when creating new trustmesh entry", 500, c)
			return
		}

		restutil.Render(transactionHash, 200, c)
	}
}

func newSynchronizationRequest(req createInitialSuggestionRequest) *types.SynchronizationRequest {
	return &types.SynchronizationRequest{
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		Recipient:                            req.Recipient,
		WorkstepType:                         req.WorkstepType,
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
		BusinessObjectJson:                   req.BusinessObjectJson,
		ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
	}
}

func createSuggestionSentTrustmeshEntry(req createInitialSuggestionRequest, transactionId uuid.UUID, offchainMsg types.OffchainProcessMessage, txHash string) *types.TrustmeshEntry {
	return &types.TrustmeshEntry{
		TendermintTransactionId:  transactionId,
		OffchainProcessMessageId: offchainMsg.Id,
		// TODO: define proxy identifier, BAS-33
		SenderOrgId:                          uuid.FromStringOrNil("5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8"),
		ReceiverOrgId:                        uuid.FromStringOrNil(req.Recipient),
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		WorkstepType:                         offchainMsg.WorkstepType,
		BaseledgerTransactionType:            "Suggest",
		BaseledgerTransactionId:              transactionId,
		ReferencedBaseledgerTransactionId:    uuid.FromStringOrNil(req.ReferencedBaseledgerTransactionId),
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerBusinessObjectId:           offchainMsg.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: offchainMsg.ReferencedBaseledgerBusinessObjectId,
		ReferencedProcessMessageId:           offchainMsg.ReferencedOffchainProcessMessageId,
		TransactionHash:                      txHash,
		EntryType:                            common.SuggestionSentTrustmeshEntryType,
	}
}

func createSuggestOffchainMessage(req createInitialSuggestionRequest, transactionId uuid.UUID, hash string) types.OffchainProcessMessage {
	offchainMessage := types.OffchainProcessMessage{
		// TODO: define proxy identifier
		SenderId:                             uuid.FromStringOrNil("5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8"),
		ReceiverId:                           uuid.FromStringOrNil(req.Recipient),
		Topic:                                req.WorkgroupId,
		WorkstepType:                         req.WorkstepType,
		ReferencedOffchainProcessMessageId:   uuid.FromStringOrNil(""),
		BaseledgerSyncTreeJson:               req.BusinessObjectJson,
		BusinessObjectProof:                  hash,
		BaseledgerBusinessObjectId:           uuid.FromStringOrNil(req.BaseledgerBusinessObjectId),
		ReferencedBaseledgerBusinessObjectId: uuid.FromStringOrNil(req.ReferencedBaseledgerBusinessObjectId),
		StatusTextMessage:                    req.WorkstepType + " suggested",
		BaseledgerTransactionIdOfStoredProof: transactionId,
		TendermintTransactionIdOfStoredProof: transactionId,
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerTransactionType:            "Suggest",
		ReferencedBaseledgerTransactionId:    uuid.FromStringOrNil(req.ReferencedBaseledgerTransactionId),
		EntryType:                            common.SuggestionSentTrustmeshEntryType,
	}

	return offchainMessage
}
