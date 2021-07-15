package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/logger"
	baseledgertypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"google.golang.org/grpc"
)

func getCommittedTransactionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		committedBaseledgerTransaction, err := getCommittedBaseledgerTransaction(uuid.FromStringOrNil(vars["txId"]))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, err := json.Marshal(committedBaseledgerTransaction)
		if err != nil {
			logger.Errorf("error when marshaling committed transaction %v\n", err)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func getCommittedBaseledgerTransaction(transactionId uuid.UUID) (*baseledgertypes.BaseledgerTransaction, error) {
	// TODO: BAS-33
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",
		// The SDK doesn't support any transport security mechanism.
		grpc.WithInsecure(),
	)
	defer grpcConn.Close()

	if err != nil {
		// TODO: error handling
		logger.Errorf("grpc conn failed %v\n", err.Error())
		return nil, err
	}

	queryClient := baseledgertypes.NewQueryClient(grpcConn)

	res, err := queryClient.BaseledgerTransaction(context.Background(), &baseledgertypes.QueryGetBaseledgerTransactionRequest{Id: transactionId.String()})

	if err != nil {
		// TODO: error handling
		logger.Errorf("grpc query failed %v\n", err.Error())
		return nil, err
	}

	logger.Infof("found baseledger transaction %v\n", res.BaseledgerTransaction)
	return res.BaseledgerTransaction, nil
}
