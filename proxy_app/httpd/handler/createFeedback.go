package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/proxyutil"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

type createSynchronizationFeedbackRequest struct {
	WorkgroupId                                string `json:"workgroup_id"`
	BusinessObjectType                         string `json:"business_object_type"`
	Recipient                                  string `json:"recipient"`
	Approved                                   bool   `json:"approved"`
	BaseledgerBusinessObjectIdOfApprovedObject string `json:"baseledger_business_object_id_of_approved_object"`
	HashOfObjectToApprove                      string `json:"hash_of_object_to_approve"`
	OriginalBaseledgerTransactionId            string `json:"original_baseledger_transaction_id"`
	OriginalOffchainProcessMessageId           string `json:"original_offchain_process_message_id"`
	FeedbackMessage                            string `json:"feedback_message"`
	BaseledgerProvenBusinessObjectJson         string `json:"baseledger_proven_business_object_json"`
}

func CreateSynchronizationFeedbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		req := &createSynchronizationFeedbackRequest{}
		err = json.Unmarshal(buf, &req)
		if err != nil {
			restutil.RenderError(err.Error(), 422, c)
			return
		}

		createFeedbackReq := newFeedbackRequest(*req)

		transactionId := uuid.NewV4()
		feedbackMsg := "Approve"
		if !req.Approved {
			feedbackMsg = "Reject"
		}

		offchainMsg := createFeedbackOffchainMessage(*req, transactionId, feedbackMsg)

		if !offchainMsg.Create() {
			logger.Errorf("error when creating new offchain msg entry")
			restutil.RenderError("error when creating new offchain msg entry", 500, c)
			return
		}

		payload := proxyutil.CreateBaseledgerTransactionFeedbackPayload(createFeedbackReq, &offchainMsg)

		signAndBroadcastPayload := restutil.SignAndBroadcastPayload{
			TransactionId: transactionId.String(),
			Payload:       payload,
		}

		transactionHash := restutil.SignAndBroadcast(signAndBroadcastPayload, c)

		if transactionHash == nil {
			restutil.RenderError("sign and broadcast transaction error", 500, c)
			return
		}

		trustmeshEntry := createFeedbackSentTrustmeshEntry(*req, transactionId, offchainMsg, feedbackMsg, *transactionHash)

		if !trustmeshEntry.Create() {
			logger.Errorf("error when creating new trustmesh entry")
			restutil.RenderError("error when creating new trustmesh entry", 500, c)
			return
		}

		restutil.Render(transactionHash, 200, c)
	}
}

func createFeedbackOffchainMessage(req createSynchronizationFeedbackRequest, transactionId uuid.UUID, baseledgerTransactionType string) types.OffchainProcessMessage {
	offchainMessage := types.OffchainProcessMessage{
		SenderId:                             uuid.FromStringOrNil("5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8"),
		ReceiverId:                           uuid.FromStringOrNil(req.Recipient),
		Topic:                                req.WorkgroupId,
		WorkstepType:                         "Feedback",
		ReferencedOffchainProcessMessageId:   uuid.FromStringOrNil(req.OriginalOffchainProcessMessageId),
		BaseledgerSyncTreeJson:               req.BaseledgerProvenBusinessObjectJson,
		BusinessObjectProof:                  req.HashOfObjectToApprove,
		BaseledgerBusinessObjectId:           uuid.FromStringOrNil(""),
		ReferencedBaseledgerBusinessObjectId: uuid.FromStringOrNil(req.BaseledgerBusinessObjectIdOfApprovedObject),
		StatusTextMessage:                    req.FeedbackMessage,
		BaseledgerTransactionIdOfStoredProof: transactionId,
		TendermintTransactionIdOfStoredProof: transactionId,
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerTransactionType:            baseledgerTransactionType,
		ReferencedBaseledgerTransactionId:    uuid.FromStringOrNil(req.OriginalBaseledgerTransactionId),
		EntryType:                            common.FeedbackSentTrustmeshEntryType,
	}

	return offchainMessage
}

func createFeedbackSentTrustmeshEntry(req createSynchronizationFeedbackRequest, transactionId uuid.UUID, offchainMsg types.OffchainProcessMessage, feedbackMsg string, txHash string) *types.TrustmeshEntry {
	trustmeshEntry := &types.TrustmeshEntry{
		TendermintTransactionId:  transactionId,
		OffchainProcessMessageId: offchainMsg.Id,
		// TODO: define proxy identifier, BAS-33
		SenderOrgId:                          uuid.FromStringOrNil("5d187a23-c4f6-4780-b8bf-aeeaeafcb1e8"),
		ReceiverOrgId:                        uuid.FromStringOrNil(req.Recipient),
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		WorkstepType:                         offchainMsg.WorkstepType,
		BaseledgerTransactionType:            feedbackMsg,
		BaseledgerTransactionId:              transactionId,
		ReferencedBaseledgerTransactionId:    uuid.FromStringOrNil(req.OriginalBaseledgerTransactionId),
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerBusinessObjectId:           offchainMsg.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: offchainMsg.ReferencedBaseledgerBusinessObjectId,
		ReferencedProcessMessageId:           offchainMsg.ReferencedOffchainProcessMessageId,
		TransactionHash:                      txHash,
		EntryType:                            common.FeedbackSentTrustmeshEntryType,
	}

	return trustmeshEntry
}

func newFeedbackRequest(req createSynchronizationFeedbackRequest) *types.SynchronizationFeedback {
	return &types.SynchronizationFeedback{
		WorkgroupId:                        uuid.FromStringOrNil(req.WorkgroupId),
		BaseledgerProvenBusinessObjectJson: req.BaseledgerProvenBusinessObjectJson,
		Recipient:                          req.Recipient,
		Approved:                           req.Approved,
		BaseledgerBusinessObjectIdOfApprovedObject: req.BaseledgerBusinessObjectIdOfApprovedObject,
		HashOfObjectToApprove:                      req.HashOfObjectToApprove,
		OriginalBaseledgerTransactionId:            req.OriginalBaseledgerTransactionId,
		OriginalOffchainProcessMessageId:           req.OriginalOffchainProcessMessageId,
		FeedbackMessage:                            req.FeedbackMessage,
		BusinessObjectType:                         req.BusinessObjectType,
	}
}
