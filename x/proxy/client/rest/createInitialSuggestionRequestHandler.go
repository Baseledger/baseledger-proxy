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

type createInitialSuggestionRequest struct {
	BaseReq                              rest.BaseReq `json:"base_req"`
	WorkgroupId                          string       `json:"workgroup_id"`
	Recipient                            string       `json:"recipient"`
	WorkstepType                         string       `json:"workstep_type"`
	BusinessObjectType                   string       `json:"business_object_type"`
	BaseledgerBusinessObjectId           string       `json:"baseledger_business_object_id"`
	BusinessObjectJson                   string       `json:"business_object_json"`
	ReferencedBaseledgerBusinessObjectId string       `json:"referenced_baseledger_business_object_id"`
	ReferencedBaseledgerTransactionId    string       `json:"referenced_baseledger_transaction_id"`
}

func createInitialSuggestionRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := parseRequest(w, r, clientCtx)
		clientCtx, err := txutil.BuildClientCtx(clientCtx, req.BaseReq.From)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		accNum, accSeq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(*clientCtx, clientCtx.FromAddress)

		if err != nil {
			fmt.Printf("error while retrieving acc %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "error while retrieving acc")
			return
		}

		createSyncReq := newSynchronizationRequest(*req)

		hash := proxy.CreateHashFromBusinessObject(req.BusinessObjectJson)
		transactionId, _ := uuid.NewV4()

		offchainMsg := createSuggestOffchainMessage(*req, transactionId, hash)

		if !offchainMsg.Create() {
			fmt.Printf("error when creating new offchain msg entry")
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "error when creating new offchain msg entry")
			return
		}

		payload := proxy.CreateBaseledgerTransactionPayload(createSyncReq, &offchainMsg)

		msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(transactionId.String(), clientCtx.GetFromAddress().String(), transactionId.String(), string(payload))
		if err := msg.ValidateBasic(); err != nil {
			fmt.Printf("msg validate basic failed %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Printf("msg with encrypted payload to be broadcasted %s\n", msg)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		txHash, err := txutil.BroadcastAndGetTxHash(*clientCtx, msg, accNum, accSeq, false)

		if err != nil {
			fmt.Printf("broadcasting failed %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		trustmeshEntry := createSuggestionSentTrustmeshEntry(*req, transactionId, offchainMsg, *txHash)

		if !trustmeshEntry.Create() {
			fmt.Printf("error when creating new trustmesh entry")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
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

func parseRequest(w http.ResponseWriter, r *http.Request, clientCtx client.Context) *createInitialSuggestionRequest {
	var req createInitialSuggestionRequest
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
