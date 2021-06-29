package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	baseledgerTypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"github.com/unibrightio/baseledger/x/proxy/proxy"
	"github.com/unibrightio/baseledger/x/proxy/types"

	uuid "github.com/kthomas/go.uuid"
	common "github.com/unibrightio/baseledger/common"
	txutil "github.com/unibrightio/baseledger/txutil"
)

// keep in sync with restuitl struct
type createSynchronizationFeedbackRequest struct {
	BaseReq                                    rest.BaseReq `json:"base_req"`
	WorkgroupId                                string       `json:"workgroup_id"`
	BusinessObjectType                         string       `json:"business_object_type"`
	Recipient                                  string       `json:"recipient"`
	Approved                                   bool         `json:"approved"`
	BaseledgerBusinessObjectIdOfApprovedObject string       `json:"baseledger_business_object_id_of_approved_object"`
	HashOfObjectToApprove                      string       `json:"hash_of_object_to_approve"`
	OriginalBaseledgerTransactionId            string       `json:"original_baseledger_transaction_id"`
	OriginalOffchainProcessMessageId           string       `json:"original_offchain_process_message_id"`
	FeedbackMessage                            string       `json:"feedback_message"`
	BaseledgerProvenBusinessObjectJson         string       `json:"baseledger_proven_business_object_json"`
}

func createSynchronizationFeedbackHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := parseFeedbackRequest(w, r, clientCtx)
		clientCtx, err := txutil.BuildClientCtx(clientCtx, req.BaseReq.From)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		createFeedbackReq := newFeedbackRequest(*req)

		transactionId, _ := uuid.NewV4()
		feedbackMsg := "Approve"
		if !req.Approved {
			feedbackMsg = "Reject"
		}

		offchainMsg := createFeedbackOffchainMessage(*req, transactionId, feedbackMsg)

		if !offchainMsg.Create() {
			fmt.Printf("error when creating new offchain msg entry")
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "error when creating new offchain msg entry")
		}

		payload := proxy.CreateBaseledgerTransactionFeedbackPayload(createFeedbackReq, &offchainMsg)

		msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(transactionId.String(), clientCtx.GetFromAddress().String(), transactionId.String(), string(payload))
		if err := msg.ValidateBasic(); err != nil {
			fmt.Printf("msg validate basic failed %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		fmt.Printf("msg with encrypted payload to be broadcasted %s\n", msg)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		txBytes, err := txutil.SignTxAndGetTxBytes(*clientCtx, msg)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		res, err := clientCtx.BroadcastTx(txBytes)
		if err != nil {
			fmt.Printf("error while broadcasting tx %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		fmt.Printf("TRANSACTION BROADCASTED WITH RESULT %v\n", res)
		trustmeshEntry := createFeedbackSentTrustmeshEntry(*req, transactionId, offchainMsg, feedbackMsg, res.TxHash)

		trustmeshEntry.OffchainProcessMessageId = offchainMsg.Id
		if !trustmeshEntry.Create() {
			fmt.Printf("error when creating new trustmesh entry")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
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

func parseFeedbackRequest(w http.ResponseWriter, r *http.Request, clientCtx client.Context) *createSynchronizationFeedbackRequest {
	var req createSynchronizationFeedbackRequest
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
