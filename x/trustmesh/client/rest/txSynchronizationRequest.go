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

type createSynchronizationRequestRequest struct {
	BaseReq                              rest.BaseReq `json:"base_req"`
	Creator                              string       `json:"creator"`
	WorkgroupID                          string       `json:"WorkgroupID"`
	Recipient                            string       `json:"Recipient"`
	WorkstepType                         string       `json:"WorkstepType"`
	BusinessObjectType                   string       `json:"BusinessObjectType"`
	BaseledgerBusinessObjectID           string       `json:"BaseledgerBusinessObjectID"`
	BusinessObject                       string       `json:"BusinessObject"`
	ReferencedBaseledgerBusinessObjectID string       `json:"ReferencedBaseledgerBusinessObjectID"`
}

func createSynchronizationRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSynchronizationRequestRequest
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

		parsedWorkgroupID := req.WorkgroupID

		parsedRecipient := req.Recipient

		parsedWorkstepType := req.WorkstepType

		parsedBusinessObjectType := req.BusinessObjectType

		parsedBaseledgerBusinessObjectID := req.BaseledgerBusinessObjectID

		parsedBusinessObject := req.BusinessObject

		parsedReferencedBaseledgerBusinessObjectID := req.ReferencedBaseledgerBusinessObjectID

		msg := types.NewMsgCreateSynchronizationRequest(
			req.Creator,
			parsedWorkgroupID,
			parsedRecipient,
			parsedWorkstepType,
			parsedBusinessObjectType,
			parsedBaseledgerBusinessObjectID,
			parsedBusinessObject,
			parsedReferencedBaseledgerBusinessObjectID,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

/* type updateSynchronizationRequestRequest struct {
	BaseReq                              rest.BaseReq `json:"base_req"`
	Creator                              string       `json:"creator"`
	WorkgroupID                          string       `json:"WorkgroupID"`
	Recipient                            string       `json:"Recipient"`
	WorkstepType                         string       `json:"WorkstepType"`
	BusinessObjectType                   string       `json:"BusinessObjectType"`
	BaseledgerBusinessObjectID           string       `json:"BaseledgerBusinessObjectID"`
	BusinessObject                       string       `json:"BusinessObject"`
	ReferencedBaseledgerBusinessObjectID string       `json:"ReferencedBaseledgerBusinessObjectID"`
} */

/* func updateSynchronizationRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateSynchronizationRequestRequest
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

		parsedWorkgroupID := req.WorkgroupID

		parsedRecipient := req.Recipient

		parsedWorkstepType := req.WorkstepType

		parsedBusinessObjectType := req.BusinessObjectType

		parsedBaseledgerBusinessObjectID := req.BaseledgerBusinessObjectID

		parsedBusinessObject := req.BusinessObject

		parsedReferencedBaseledgerBusinessObjectID := req.ReferencedBaseledgerBusinessObjectID

		msg := types.NewMsgUpdateSynchronizationRequest(
			req.Creator,
			id,
			parsedWorkgroupID,
			parsedRecipient,
			parsedWorkstepType,
			parsedBusinessObjectType,
			parsedBaseledgerBusinessObjectID,
			parsedBusinessObject,
			parsedReferencedBaseledgerBusinessObjectID,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
} */

/* type deleteSynchronizationRequestRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
} */

/* func deleteSynchronizationRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteSynchronizationRequestRequest
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

		msg := types.NewMsgDeleteSynchronizationRequest(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
} */
