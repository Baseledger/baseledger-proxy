package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers trustmesh-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/trustmesh/SynchronizationFeedbacks/{id}", getSynchronizationFeedbackHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/trustmesh/SynchronizationFeedbacks", listSynchronizationFeedbackHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/trustmesh/SynchronizationRequests/{id}", getSynchronizationRequestHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/trustmesh/SynchronizationRequests", listSynchronizationRequestHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/trustmesh/SynchronizationFeedbacks", createSynchronizationFeedbackHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/trustmesh/SynchronizationFeedbacks/{id}", updateSynchronizationFeedbackHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/trustmesh/SynchronizationFeedbacks/{id}", deleteSynchronizationFeedbackHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/trustmesh/SynchronizationRequests", createSynchronizationRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/trustmesh/SynchronizationRequests/{id}", updateSynchronizationRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/trustmesh/SynchronizationRequests/{id}", deleteSynchronizationRequestHandler(clientCtx)).Methods("POST")

}
