package rest

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	baseledgerTypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"github.com/unibrightio/baseledger/x/proxy/proxy"
	"github.com/unibrightio/baseledger/x/proxy/types"

	uuid "github.com/kthomas/go.uuid"
	txutil "github.com/unibrightio/baseledger/txutil"
)

const (
	errCodeMismatch = 32
)

var (
	// errors are of the form:
	// "account sequence mismatch, expected 25, got 27: incorrect account sequence"
	recoverRegexp = regexp.MustCompile(`^account sequence mismatch, expected (\d+), got (\d+):`)
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

		txBytes, err := txutil.SignTxAndGetTxBytes(*clientCtx, msg, accNum, accSeq)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, err := clientCtx.BroadcastTx(txBytes)
		if err != nil {
			fmt.Printf("error while broadcasting tx %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Printf("TRANSACTION BROADCASTED WITH RESULT %v\n", res)

		trustmeshEntry := &types.TrustmeshEntry{
			TendermintTransactionId:  transactionId,
			OffchainProcessMessageId: offchainMsg.Id,
			// TODO: define proxy identifier
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
			TransactionHash:                      res.TxHash,
			EntryType:                            "SuggestionSent",
		}

		// if broadcast was successful, save new trustmesh entry
		if res.Code == 0 {
			if !trustmeshEntry.Create() {
				fmt.Printf("error when creating new trustmesh entry")
				rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New(res.RawLog).Error())
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}

		// if code is not 0 and is not missmatch, don't handle it and return
		if res.Code != errCodeMismatch {
			fmt.Printf("error while broadcasting tx 1 %v\n", res)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New(res.RawLog).Error())
			return
		}

		// if code is missmatch, parse log and try again
		fmt.Printf("ACCOUNT SEQUENCE MISSMATCH %v\n", res.RawLog)

		nextSequence, ok := parseNextSequence(accSeq, res.RawLog)

		if !ok {
			fmt.Printf("error while broadcasting tx 2 %v\n", res.Code)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New(res.RawLog).Error())
			return
		}

		fmt.Printf("RETRYING WITH SEQUENCE %v\n", nextSequence)

		txBytes, err = txutil.SignTxAndGetTxBytes(*clientCtx, msg, accNum, nextSequence)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, err = clientCtx.BroadcastTx(txBytes)
		if err != nil {
			fmt.Printf("error while broadcasting tx 3 %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// do not try again if another missmatch (or any other error) occurs
		if res.Code != 0 {
			fmt.Printf("error while broadcasting tx 4 %v\n", res.Code)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, errors.New(res.RawLog).Error())
			return
		}

		fmt.Printf("TRANSACTION REBROADCASTED WITH RESULT %v\n", res)

		// make sure to overwrite transaction hash with new one
		trustmeshEntry.TransactionHash = res.TxHash
		if !trustmeshEntry.Create() {
			fmt.Printf("error when creating new trustmesh entry")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func parseNextSequence(current uint64, message string) (uint64, bool) {
	// "account sequence mismatch, expected 25, got 27: incorrect account sequence"
	matches := recoverRegexp.FindStringSubmatch(message)

	if len(matches) != 3 {
		return 0, false
	}

	if len(matches[1]) == 0 || len(matches[2]) == 0 {
		return 0, false
	}

	expected, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil || expected == 0 {
		return 0, false
	}

	received, err := strconv.ParseUint(matches[2], 10, 64)
	if err != nil || received == 0 {
		return 0, false
	}

	if received != current {
		return expected, true
	}

	return expected, true
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
		EntryType:                            "SugggestionSent",
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
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		Recipient:                            req.Recipient,
		WorkstepType:                         req.WorkstepType,
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
		BusinessObjectJson:                   req.BusinessObjectJson,
		ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
	}
}
