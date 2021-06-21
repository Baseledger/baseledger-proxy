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
	txutil "github.com/unibrightio/baseledger/x/proxy/txutil"
)

type createInitialSuggestionRequest struct {
	BaseReq                              rest.BaseReq `json:"base_req"`
	WorkgroupId                          string       `json:"workgroup_id"`
	Recipient                            string       `json:"recipient"`
	WorkstepType                         string       `json:"workstep_type"`
	BusinessObjectType                   string       `json:"business_object_type"`
	BaseledgerBusinessObjectId           string       `json:"baseledger_business_object_id"`
	BusinessObject                       string       `json:"business_object"`
	ReferencedBaseledgerBusinessObjectId string       `json:"referenced_baseledger_business_object_id"`
	ReferencedBaseledgerTransactionId    string       `json:"referenced_baseledger_transaction_id"`
}

func createInitialSuggestionRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := parseRequest(w, r, clientCtx)
		clientCtx, err := txutil.BuildClientCtx(clientCtx, req.BaseReq.From)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		createSyncReq := newSynchronizationRequest(*req)

		hash := proxy.CreateHashFromBusinessObject(req.BusinessObject)
		transactionId, _ := uuid.NewV4()

		offchainMsg := createSuggestOffchainMessage(*req, transactionId.String(), hash)

		if !offchainMsg.Create() {
			fmt.Printf("error when creating new offchain msg entry")
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "error when creating new offchain msg entry")
		}

		payload := proxy.CreateBaseledgerTransactionPayload(createSyncReq, &offchainMsg)

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

		trustmeshEntry := &types.TrustmeshEntry{
			TendermintTransactionId:  transactionId.String(),
			OffchainProcessMessageId: offchainMsg.Id,
			// TODO: define proxy identifier
			Sender:                               "123",
			Receiver:                             req.Recipient,
			WorkgroupId:                          req.WorkgroupId,
			WorkstepType:                         offchainMsg.WorkstepType,
			BaseledgerTransactionType:            "Suggest",
			BaseledgerTransactionId:              transactionId.String(),
			ReferencedBaseledgerTransactionId:    req.ReferencedBaseledgerTransactionId,
			BusinessObjectType:                   req.BusinessObjectType,
			BaseledgerBusinessObjectId:           offchainMsg.BaseledgerBusinessObjectId,
			ReferencedBaseledgerBusinessObjectId: offchainMsg.ReferencedBaseledgerBusinessObjectId,
			ReferencedProcessMessageId:           offchainMsg.ReferencedOffchainProcessMessageId,
			TransactionHash:                      res.TxHash,
			Type:                                 "SuggestionSent",
		}

		trustmeshEntry.OffchainProcessMessageId = offchainMsg.Id
		if !trustmeshEntry.Create() {
			fmt.Printf("error when creating new trustmesh entry")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func createSuggestOffchainMessage(req createInitialSuggestionRequest, transactionId string, hash string) types.OffchainProcessMessage {
	offchainMessage := types.OffchainProcessMessage{
		WorkstepType:                         req.WorkstepType,
		ReferencedOffchainProcessMessageId:   "",
		BusinessObject:                       req.BusinessObject,
		Hash:                                 hash,
		BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
		StatusTextMessage:                    req.WorkstepType + " suggested",
		BaseledgerTransactionIdOfStoredProof: transactionId,
	}

	return offchainMessage
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
		WorkgroupId:                          req.WorkgroupId,
		Recipient:                            req.Recipient,
		WorkstepType:                         req.WorkstepType,
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
		BusinessObject:                       req.BusinessObject,
		ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
	}
}
