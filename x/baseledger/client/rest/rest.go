package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers baseledger-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/baseledger/BaseledgerTransactions/{id}", getBaseledgerTransactionHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/baseledger/BaseledgerTransactions", listBaseledgerTransactionHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/baseledger/BaseledgerTransactions", createBaseledgerTransactionHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/baseledger/BaseledgerTransactions/{id}", updateBaseledgerTransactionHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/baseledger/BaseledgerTransactions/{id}", deleteBaseledgerTransactionHandler(clientCtx)).Methods("POST")

}
