package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/unibrightio/baseledger/logger"
	txutil "github.com/unibrightio/baseledger/txutil"
	baseledgerTypes "github.com/unibrightio/baseledger/x/baseledger/types"
)

type signAndBroadcastTransactionRequest struct {
	TransactionId string `json:"transaction_id"`
	Payload       string `json:"payload"`
	OpCode        uint32 `json:"op_code"`
}

func signAndBroadcastTransactionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := parseSignAndBroadcastTransactionRequest(w, r, clientCtx)

		clientCtx, err := txutil.BuildClientCtx(clientCtx)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		accNum, accSeq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(*clientCtx, clientCtx.FromAddress)

		if err != nil {
			logger.Errorf("error while retrieving acc %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "error while retrieving acc")
			return
		}

		msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(req.TransactionId, clientCtx.GetFromAddress().String(), req.TransactionId, req.Payload, req.OpCode)
		if err := msg.ValidateBasic(); err != nil {
			logger.Errorf("msg validate basic failed %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		logger.Infof("msg with encrypted payload to be broadcasted %s\n", msg)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		txHash, err := txutil.BroadcastAndGetTxHash(*clientCtx, msg, accNum, accSeq, false)

		if err != nil {
			logger.Errorf("broadcasting failed %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		logger.Infof("broadcasted tx hash %v\n", *txHash)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(*txHash))
		w.WriteHeader(http.StatusOK)
		return
	}
}

func parseSignAndBroadcastTransactionRequest(w http.ResponseWriter, r *http.Request, clientCtx client.Context) *signAndBroadcastTransactionRequest {
	var req signAndBroadcastTransactionRequest
	if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
		return nil
	}

	return &req
}
