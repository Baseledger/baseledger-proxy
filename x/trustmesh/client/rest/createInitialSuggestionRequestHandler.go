package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	baseledgerTypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"github.com/unibrightio/baseledger/x/trustmesh/proxy"
	"github.com/unibrightio/baseledger/x/trustmesh/types"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
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
}

func createInitialSuggestionRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := parseRequest(w, r, clientCtx)
		clientCtx, err := buildClientCtx(clientCtx, req.BaseReq.From)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		createSyncReq := newSynchronizationRequest(*req)

		payload, transactionId := proxy.SynchronizeBusinessObject(createSyncReq)

		msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(clientCtx.GetFromAddress().String(), transactionId, string(payload))
		if err := msg.ValidateBasic(); err != nil {
			fmt.Printf("msg validate basic failed %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		fmt.Printf("msg with encrypted payload to be broadcasted %s\n", msg)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		txBytes, err := signTxAndGetTxBytes(*clientCtx, msg)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		res, err := clientCtx.BroadcastTx(txBytes)
		if err != nil {
			fmt.Printf("error while broadcasting tx %v\n", err.Error())
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		fmt.Printf("TRANSACTION BROADCASTED WITH RESULT %v\n", res)

		// TODO: fix this mocked entry with proper entries, discuss unkown props with team
		trustmeshEntry := &types.TrustmeshEntry{
			// TODO: what should we use here? this is only availabe in "block" broadcast mode which is not recommended because of timeout
			TendermintBlockId: strconv.FormatUint(uint64(res.Height), 10),
			// TODO: what should we use here?
			TendermintTransactionId: res.TxHash,
			// TODO: what should we use here? timestamp not available
			TendermintTransactionTimestamp:       "",
			Sender:                               "123",
			Receiver:                             "123",
			WorkgroupId:                          req.WorkgroupId,
			WorkstepType:                         req.WorkstepType,
			BaseledgerTransactionType:            "Suggest",
			BaseledgerTransactionId:              transactionId,
			ReferencedBaseledgerTransactionId:    "123",
			BusinessObjectType:                   req.BusinessObjectType,
			BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
			ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
			OffchainProcessMessageId:             "123",
			ReferencedProcessMessageId:           "123",
		}

		if !trustmeshEntry.Create() {
			fmt.Printf("error when creating new trustmesh entry")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
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
		WorkgroupId:                          req.WorkgroupId,
		Recipient:                            req.Recipient,
		WorkstepType:                         req.WorkstepType,
		BusinessObjectType:                   req.BusinessObjectType,
		BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
		BusinessObject:                       req.BusinessObject,
		ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
	}
}

func buildClientCtx(clientCtx client.Context, from string) (*client.Context, error) {
	fromAddress, err := sdk.AccAddressFromBech32(from)

	keyring, err := newKeyringInstance()
	key, err := keyring.KeyByAddress(fromAddress)

	if err != nil {
		fmt.Printf("error getting key %v\n", err.Error())
		return nil, errors.New("")
	}

	fmt.Printf("key found %v %v\n", key, key.GetName())

	clientCtx = clientCtx.
		WithKeyring(keyring).
		WithFromAddress(fromAddress).
		WithSkipConfirmation(true).
		WithFromName(key.GetName()).
		WithBroadcastMode("sync")

	return &clientCtx, nil
}

// TODO: change test keyring with other (file?)
func newKeyringInstance() (keyring.Keyring, error) {
	kr, err := keyring.New("baseledger", "test", "~/.baseledger", nil)

	if err != nil {
		fmt.Printf("error fetching test keyring %v\n", err.Error())
		return nil, errors.New("error fetching key ring")
	}

	return kr, nil
}

func signTxAndGetTxBytes(clientCtx client.Context, msg sdk.Msg) ([]byte, error) {
	accNum, accSeq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(clientCtx, clientCtx.FromAddress)

	if err != nil {
		fmt.Printf("error while retrieving acc %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}
	fmt.Printf("retrieved account %v %v\n", accNum, accSeq)
	txFactory := tx.Factory{}.
		WithChainID(clientCtx.ChainID).
		WithGas(100000).
		WithTxConfig(clientCtx.TxConfig).
		WithAccountNumber(accNum).
		WithSequence(accSeq).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring)

	txFactory, err = tx.PrepareFactory(clientCtx, txFactory)
	if err != nil {
		fmt.Printf("prepare factory error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	transaction, err := tx.BuildUnsignedTx(txFactory, msg)
	if err != nil {
		fmt.Printf("build unsigned tx error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	err = tx.Sign(txFactory, clientCtx.GetFromName(), transaction, true)
	if err != nil {
		fmt.Printf("sign tx error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(transaction.GetTx())
	if err != nil {
		fmt.Printf("tx encoder %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	return txBytes, nil
}
