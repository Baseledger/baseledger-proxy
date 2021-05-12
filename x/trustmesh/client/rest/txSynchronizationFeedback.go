package rest

import (
	"net/http"
	//"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/example/baseledger/x/trustmesh/types"
	//"github.com/gorilla/mux"
)

type createSynchronizationFeedbackRequest struct {
	BaseReq                                    rest.BaseReq `json:"base_req"`
	Creator                                    string       `json:"creator"`
	Approved                                   string       `json:"Approved"`
	BusinessObject                             string       `json:"BusinessObject"`
	BaseledgerBusinessObjectIDOfApprovedObject string       `json:"BaseledgerBusinessObjectIDOfApprovedObject"`
	Workgroup                                  string       `json:"Workgroup"`
	Recipient                                  string       `json:"Recipient"`
	HashOfObjectToApprove                      string       `json:"HashOfObjectToApprove"`
	OriginalBaseledgerTransactionID            string       `json:"OriginalBaseledgerTransactionID"`
	OriginalOffchainProcessMessageID           string       `json:"OriginalOffchainProcessMessageID"`
	FeedbackMessage                            string       `json:"FeedbackMessage"`
}

func createSynchronizationFeedbackHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSynchronizationFeedbackRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedApproved := req.Approved

		parsedBusinessObject := req.BusinessObject

		parsedBaseledgerBusinessObjectIDOfApprovedObject := req.BaseledgerBusinessObjectIDOfApprovedObject

		parsedWorkgroup := req.Workgroup

		parsedRecipient := req.Recipient

		parsedHashOfObjectToApprove := req.HashOfObjectToApprove

		parsedOriginalBaseledgerTransactionID := req.OriginalBaseledgerTransactionID

		parsedOriginalOffchainProcessMessageID := req.OriginalOffchainProcessMessageID

		parsedFeedbackMessage := req.FeedbackMessage

		msg := types.NewMsgCreateSynchronizationFeedback(
			req.Creator,
			parsedApproved,
			parsedBusinessObject,
			parsedBaseledgerBusinessObjectIDOfApprovedObject,
			parsedWorkgroup,
			parsedRecipient,
			parsedHashOfObjectToApprove,
			parsedOriginalBaseledgerTransactionID,
			parsedOriginalOffchainProcessMessageID,
			parsedFeedbackMessage,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

/*
type updateSynchronizationFeedbackRequest struct {
	BaseReq                                    rest.BaseReq `json:"base_req"`
	Creator                                    string       `json:"creator"`
	Approved                                   string       `json:"Approved"`
	BusinessObject                             string       `json:"BusinessObject"`
	BaseledgerBusinessObjectIDOfApprovedObject string       `json:"BaseledgerBusinessObjectIDOfApprovedObject"`
	Workgroup                                  string       `json:"Workgroup"`
	Recipient                                  string       `json:"Recipient"`
	HashOfObjectToApprove                      string       `json:"HashOfObjectToApprove"`
	OriginalBaseledgerTransactionID            string       `json:"OriginalBaseledgerTransactionID"`
	OriginalOffchainProcessMessageID           string       `json:"OriginalOffchainProcessMessageID"`
	FeedbackMessage                            string       `json:"FeedbackMessage"`
}

func updateSynchronizationFeedbackHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateSynchronizationFeedbackRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedApproved := req.Approved

		parsedBusinessObject := req.BusinessObject

		parsedBaseledgerBusinessObjectIDOfApprovedObject := req.BaseledgerBusinessObjectIDOfApprovedObject

		parsedWorkgroup := req.Workgroup

		parsedRecipient := req.Recipient

		parsedHashOfObjectToApprove := req.HashOfObjectToApprove

		parsedOriginalBaseledgerTransactionID := req.OriginalBaseledgerTransactionID

		parsedOriginalOffchainProcessMessageID := req.OriginalOffchainProcessMessageID

		parsedFeedbackMessage := req.FeedbackMessage

		msg := types.NewMsgUpdateSynchronizationFeedback(
			req.Creator,
			id,
			parsedApproved,
			parsedBusinessObject,
			parsedBaseledgerBusinessObjectIDOfApprovedObject,
			parsedWorkgroup,
			parsedRecipient,
			parsedHashOfObjectToApprove,
			parsedOriginalBaseledgerTransactionID,
			parsedOriginalOffchainProcessMessageID,
			parsedFeedbackMessage,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteSynchronizationFeedbackRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteSynchronizationFeedbackHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteSynchronizationFeedbackRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeleteSynchronizationFeedback(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
*/
