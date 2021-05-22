package rest

import (
	"fmt"
	"net/http"

	//"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	baseledgerTypes "github.com/example/baseledger/x/baseledger/types"
	"github.com/example/baseledger/x/trustmesh/proxy"
	"github.com/example/baseledger/x/trustmesh/types"

	//"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"context"

	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/grpc"
)

type createSynchronizationRequest struct {
	BaseReq                              rest.BaseReq `json:"base_req"`
	Creator                              string       `json:"creator"`
	WorkgroupId                          string       `json:"workgroup_id"`
	Recipient                            string       `json:"recipient"`
	WorkstepType                         string       `json:"workstep_type"`
	BusinessObjectType                   string       `json:"business_object_type"`
	BaseledgerBusinessObjectId           string       `json:"baseledger_business_object_id"`
	BusinessObject                       string       `json:"business_object"`
	ReferencedBaseledgerBusinessObjectId string       `json:"referenced_baseledger_business_object_id"`
}

func createSynchronizationRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSynchronizationRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		lol, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		createSyncReq := &types.SynchronizationRequest{
			WorkgroupId:                          req.WorkgroupId,
			Recipient:                            req.Recipient,
			WorkstepType:                         req.WorkstepType,
			BusinessObjectType:                   req.BusinessObjectType,
			BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
			BusinessObject:                       req.BusinessObject,
			ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
		}

		payload, transactionId := proxy.SynchronizeBusinessObject(createSyncReq)

		msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(clientCtx.GetFromAddress().String(), transactionId, string(payload))
		if err := msg.ValidateBasic(); err != nil {
			fmt.Printf("msg validate basic failed %v\n", err.Error())
		}

		fmt.Printf("msg with encrypted payload to be broadcasted %s\n", msg)

		txBuilder := clientCtx.TxConfig.NewTxBuilder()
		txBuilder.SetGasLimit(100000)

		// TODO: move path to constant, and figure out how to pull mnemonic from config
		keyBytes, _ := hd.Secp256k1.Derive()("prize fly elder purity kiss pluck risk voice armed become elegant odor ugly prepare merge quiz nature memory setup armor evoke hello faculty advance", "", "m/44'/118'/0'/0/0")

		key := hd.Secp256k1.Generate()(keyBytes)
		address := sdk.AccAddress(key.PubKey().Address()).String()
		fmt.Printf("cosmos address %v\n", address)

		err = txBuilder.SetMsgs(msg)
		if err != nil {
			fmt.Printf("Error when setting msg %v \n", err)
		}

		// TODO: sequence always increments?
		accNum := 0
		accSeq := 1

		sigV2 := signing.SignatureV2{
			PubKey: key.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  clientCtx.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: uint64(accSeq),
		}

		err = txBuilder.SetSignatures(sigV2)
		if err != nil {
			fmt.Printf("Error when setting signatures 1 %v \n", err)
		}

		if err != nil {
			fmt.Printf("error in new acc %v\n", err.Error())
		}

		signerData := xauthsigning.SignerData{
			ChainID:       clientCtx.ChainID,
			AccountNumber: uint64(accNum),
			Sequence:      uint64(accSeq),
		}

		sigV22, err := tx.SignWithPrivKey(
			clientCtx.TxConfig.SignModeHandler().DefaultMode(),
			signerData,
			txBuilder,
			key,
			clientCtx.TxConfig,
			uint64(accSeq))

		if err != nil {
			fmt.Printf("Error when sign priv key %v \n", err)
		}

		err = txBuilder.SetSignatures(sigV22)
		if err != nil {
			fmt.Printf("Error when setting signatures 2 %v \n", err)
		}

		clientCtx.SkipConfirm = true

		txBytes := getTxBytes(clientCtx.TxConfig, txBuilder)

		clientCtx = clientCtx.WithKeyringDir("/root/.baseledger")

		clientCtx = clientCtx.WithFrom(key.PubKey().Address().String()).WithFromName("alice").WithFromAddress(lol).WithBroadcastMode("sync")

		fmt.Printf("key ring %v \n", clientCtx.Keyring)

		sendTx(txBytes)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func getTxBytes(txConfig client.TxConfig, txBuilder client.TxBuilder) []byte {
	// Generated Protobuf-encoded bytes.
	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		fmt.Printf("tx bytes error %v \n", err)
	}

	fmt.Printf("tx bytes %v \n", txBytes)
	return txBytes
}

// TODO: try to broadcast without using gRPC
func sendTx(txBytes []byte) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		fmt.Printf("dial rpc error %v\n", err)
	}
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := txTypes.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txTypes.BroadcastTxRequest{
			Mode:    txTypes.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		fmt.Printf("broadcast error %v\n", err)
	}

	fmt.Println(grpcRes.TxResponse, grpcRes.TxResponse.Code)
}
