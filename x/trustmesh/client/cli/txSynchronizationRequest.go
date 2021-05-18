package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	//"strconv"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/simapp/params"

	// "github.com/cosmos/cosmos-sdk/client/tx"

	baseledgerTypes "github.com/example/baseledger/x/baseledger/types"
	"github.com/example/baseledger/x/trustmesh/proxy"
	"github.com/example/baseledger/x/trustmesh/types"

	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"context"

	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CmdCreateSynchronizationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-SynchronizationRequest [WorkgroupId] [Recipient] [WorkstepType] [BusinessObjectType] [BaseledgerBusinessObjectId] [BusinessObject] [ReferencedBaseledgerBusinessObjectId]",
		Short: "Creates a new SynchronizationRequest",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {

			createSyncReq := &types.SynchronizationRequest{
				WorkgroupId:                          string(args[0]),
				Recipient:                            string(args[1]),
				WorkstepType:                         string(args[2]),
				BusinessObjectType:                   string(args[3]),
				BaseledgerBusinessObjectId:           string(args[4]),
				BusinessObject:                       string(args[5]),
				ReferencedBaseledgerBusinessObjectId: string(args[6]),
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			payload, transactionId := proxy.SynchronizeBusinessObject(createSyncReq)

			msg := baseledgerTypes.NewMsgCreateBaseledgerTransaction(clientCtx.GetFromAddress().String(), transactionId, string(payload))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			fmt.Printf("msg with encrypted payload to be broadcasted %s\n", msg)

			encCfg := simapp.MakeTestEncodingConfig()
			txBuilder := encCfg.TxConfig.NewTxBuilder()

			txBuilder.SetGasLimit(100000)

			// TODO: move path to constant, and figure out how to pull mnemonic from config
			keyBytes, _ := hd.Secp256k1.Derive()("favorite recipe chef seek cigar below leg flower napkin income ten situate follow order rural swear lawn frozen involve insect erode praise oblige pupil", "", "m/44'/118'/0'/0/0")

			key := hd.Secp256k1.Generate()(keyBytes)
			address := sdk.AccAddress(key.PubKey().Address()).String()
			fmt.Printf("cosmos address %v\n", address)

			err = txBuilder.SetMsgs(msg)
			if err != nil {
				fmt.Printf("Error when setting msg %v \n", err)
			}

			// TODO: sequence always increments?
			accNum := 0
			accSeq := 3

			sigV2 := signing.SignatureV2{
				PubKey: key.PubKey(),
				Data: &signing.SingleSignatureData{
					SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
					Signature: nil,
				},
				Sequence: uint64(accSeq),
			}

			err = txBuilder.SetSignatures(sigV2)
			if err != nil {
				fmt.Printf("Error when setting signatures 1 %v \n", err)
			}

			signerData := xauthsigning.SignerData{
				ChainID:       clientCtx.ChainID,
				AccountNumber: uint64(accNum),
				Sequence:      uint64(accSeq),
			}

			sigV22, err := tx.SignWithPrivKey(
				encCfg.TxConfig.SignModeHandler().DefaultMode(),
				signerData,
				txBuilder,
				key,
				encCfg.TxConfig,
				uint64(accSeq))

			if err != nil {
				fmt.Printf("Error when sign priv key %v \n", err)
			}

			err = txBuilder.SetSignatures(sigV22)
			if err != nil {
				fmt.Printf("Error when setting signatures 2 %v \n", err)
			}

			txBytes := getTxBytes(encCfg, txBuilder)
			sendTx(txBytes)
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getTxBytes(encCfg params.EncodingConfig, txBuilder client.TxBuilder) []byte {
	// Generated Protobuf-encoded bytes.
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		fmt.Printf("tx bytes error %v \n", err)
	}

	fmt.Printf("tx bytes %v \n", txBytes)

	// Generate a JSON string.
	// TODO: something missing here to encode to json, figure out what
	txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		fmt.Printf("tx json error %v \n", err)
	}
	txJSON := string(txJSONBytes)

	fmt.Printf("tx json %v \n", txJSON)

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
