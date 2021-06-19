package businesslogic

import (
	"context"
	"fmt"

	"github.com/unibrightio/baseledger/app/types"
	"github.com/unibrightio/baseledger/sor"
	baseledgertypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"github.com/unibrightio/baseledger/x/proxy/proxy"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func SetTxStatusToCommitted(txResult types.Result, db *gorm.DB) {
	result := db.Exec("UPDATE trustmesh_entries SET transaction_status = 'COMMITTED', tendermint_block_id = ?, tendermint_transaction_timestamp = ? WHERE tendermint_transaction_id = ?",
		txResult.TxInfo.TxHeight,
		txResult.TxInfo.TxTimestamp,
		txResult.Job.TrustmeshEntry.TendermintTransactionId)
	if result.RowsAffected == 1 {
		fmt.Printf("Tx %v committed \n", txResult.Job.TrustmeshEntry.TendermintTransactionId)
	} else {
		fmt.Printf("Error setting tx status to committed %v\n", result.Error)
	}
}

func ExecuteBusinessLogic(txResult types.Result) {
	switch txResult.Job.TrustmeshEntry.BaseledgerTransactionType {
	case "SuggestionSent":
		fmt.Println("SuggestionSent")
		// send offchain msg to all receivers
		// should we reuse db connection here, or open new one?
		offchainMessage, err := proxytypes.GetOffchainMsgById(txResult.Job.TrustmeshEntry.OffchainProcessMessageId)
		if err != nil {
			// TODO: what do to here?
			fmt.Println("Offchain process msg not found")
			return
		}
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.Receiver)
	case "SuggestionReceived":
		fmt.Println("SuggestionReceived")
	case "FeedbackSent":
		fmt.Println("FeedbackSent")
		// send offchain msg to sender of sugestion
		offchainMessage, err := proxytypes.GetOffchainMsgById(txResult.Job.TrustmeshEntry.OffchainProcessMessageId)
		if err != nil {
			// TODO: what do to here?
			fmt.Println("Offchain process msg not found")
			return
		}
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.Sender)
	case "FeedbackReceived":
		fmt.Println("FeedbackReceived")
		offchainMessage, err := proxytypes.GetOffchainMsgById(txResult.Job.TrustmeshEntry.OffchainProcessMessageId)
		if err != nil {
			// TODO: what do to here?
			fmt.Println("Offchain process msg not found")
			return
		}
		baseledgerTransaction := getCommittedBaseledgerTransactionId(txResult.Job.TrustmeshEntry.TendermintTransactionId)

		if baseledgerTransaction == nil {
			// TODO: what do to here?
			return
		}
		sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
	default:
		// TODO panic
		fmt.Println("UNKNOWN BUSINESS LOGIC")
	}
}

func getCommittedBaseledgerTransactionId(transactionId string) *baseledgertypes.BaseledgerTransaction {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // your gRPC server address.
		grpc.WithInsecure(), // The SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	if err != nil {
		// TODO: error handling
		fmt.Printf("grpc conn failed %v\n", err.Error())
		return nil
	}

	queryClient := baseledgertypes.NewQueryClient(grpcConn)

	res, err := queryClient.BaseledgerTransaction(context.Background(), &baseledgertypes.QueryGetBaseledgerTransactionRequest{Id: transactionId})

	if err != nil {
		// TODO: error handling
		fmt.Printf("grpc query failed %v\n", err.Error())
		return nil
	}

	fmt.Printf("FOUND BASELEDGER TRANSACTION %v\n", res.BaseledgerTransaction)
	return res.BaseledgerTransaction
}
