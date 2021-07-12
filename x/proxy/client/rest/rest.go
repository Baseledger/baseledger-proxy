package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerTxHandlers(clientCtx, r)
	registerQueryRoutes(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/proxy/trustmesh_entries", listTrustmeshEntriesHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/proxy/trustmeshes", listTrustmeshesHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/proxy/send_offchain_message", sendOffchainMessageHandler(clientCtx)).Methods("POST")
}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/proxy/suggestion", createInitialSuggestionRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/proxy/feedback", createSynchronizationFeedbackHandler(clientCtx)).Methods("POST")
}
