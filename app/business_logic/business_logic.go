package businesslogic

import (
	"context"
	"errors"
	"fmt"

	"github.com/unibrightio/baseledger/app/types"
	"github.com/unibrightio/baseledger/sor"
	baseledgertypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"github.com/unibrightio/baseledger/x/proxy/proxy"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

// TODO: currently this is saving after result is written to output channel
// do we save inside worker or after?
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
	// TODO: should we reuse db connection here, or open new one?
	offchainMessage, err := proxytypes.GetOffchainMsgById(txResult.Job.TrustmeshEntry.OffchainProcessMessageId)
	if err != nil {
		// TODO: what do to here? is return enough?
		fmt.Println("Offchain process msg not found")
		return
	}
	switch txResult.Job.TrustmeshEntry.Type {
	case "SuggestionSent":
		fmt.Println("SuggestionSent")
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.Receiver)
	case "SuggestionReceived":
		fmt.Println("SuggestionReceived")
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			// TODO: what do to here? is return enough?
			return
		}
		proof := proxy.CreateHashFromBusinessObject(offchainMessage.BusinessObject)
		if proof != offchainMessage.Hash {
			fmt.Println("Hashes don't match")
			// TODO: what do to here? is return enough?
			return
		}
		sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
	case "FeedbackSent":
		fmt.Println("FeedbackSent")
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.Sender)
	case "FeedbackReceived":
		fmt.Println("FeedbackReceived")
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			// TODO: what do to here? is return enough?
			return
		}

		sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
	default:
		// TODO: this should not happen, probably panic is ok to use here?
		fmt.Printf("unknown business process %v\n", txResult.Job.TrustmeshEntry.Type)
		panic(errors.New("uknown business process!"))
	}
}

func getCommittedBaseledgerTransaction(transactionId string) *baseledgertypes.BaseledgerTransaction {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",
		// The SDK doesn't support any transport security mechanism.
		grpc.WithInsecure(),
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

	fmt.Printf("found baseledger transaction %v\n", res.BaseledgerTransaction)
	return res.BaseledgerTransaction
}
