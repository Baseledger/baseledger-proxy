package businesslogic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/app/types"
	"github.com/unibrightio/baseledger/dbutil"
	"github.com/unibrightio/baseledger/sor"
	baseledgertypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"github.com/unibrightio/baseledger/x/proxy/proxy"
	restutil "github.com/unibrightio/baseledger/x/proxy/restutil"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"

	"google.golang.org/grpc"
)

func ExecuteBusinessLogic(txResult types.Result) {
	// TODO: it looks like we should do this first?
	setTxStatusToCommitted(txResult)
	// TODO: should we reuse db connection here, or open new one?
	offchainMessage, err := proxytypes.GetOffchainMsgById(txResult.Job.TrustmeshEntry.OffchainProcessMessageId)
	if err != nil {
		// TODO: logging
		fmt.Println("Offchain process msg not found")
		return
	}
	switch txResult.Job.TrustmeshEntry.EntryType {
	case "SuggestionSent":
		fmt.Println("SuggestionSent")
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.ReceiverOrgId.String(), txResult.Job.TrustmeshEntry.TransactionHash)
	case "SuggestionReceived":
		fmt.Println("SuggestionReceived")
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			// TODO: logging
			return
		}
		baseledgerTransactionPayload := proxytypes.BaseledgerTransactionPayload{}
		deprivitizedPayload := proxy.DeprivatizeBaseledgerTransactionPayload(baseledgerTransaction.Payload, txResult.Job.TrustmeshEntry.WorkgroupId)

		err = json.Unmarshal(([]byte)(deprivitizedPayload), &baseledgerTransactionPayload)
		if err != nil {
			fmt.Println("Failed to unmarshal baseledger transaction payload")
			return
		}

		offchainMessageBusinessObjectProof := proxy.CreateHashFromBusinessObject(offchainMessage.BaseledgerSyncTreeJson)
		if baseledgerTransactionPayload.Proof == offchainMessage.BusinessObjectProof && baseledgerTransactionPayload.Proof == offchainMessageBusinessObjectProof {
			fmt.Println("Hashes match, processing feedback")
			sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
			return
		}
		fmt.Println("Hashes don't match")
		restutil.RejectFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId.String())
	case "FeedbackSent":
		fmt.Println("FeedbackSent")
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.SenderOrgId.String(), txResult.Job.TrustmeshEntry.TransactionHash)
	case "FeedbackReceived":
		fmt.Println("FeedbackReceived")
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			// TODO: logging
			return
		}

		sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
	default:
		// TODO: this should not happen, probably panic is ok to use here?
		fmt.Printf("unknown business process %v\n", txResult.Job.TrustmeshEntry.EntryType)
		panic(errors.New("uknown business process!"))
	}
}

func setTxStatusToCommitted(txResult types.Result) {
	if txResult.TxInfo.TxHeight == "" || txResult.TxInfo.TxTimestamp == "" {
		return
	}
	// TODO: should we reuse db connection here, or open new one?
	db, err := dbutil.InitBaseledgerDBConnection()
	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
	}
	result := db.Exec("UPDATE trustmesh_entries SET commitment_state = 'COMMITTED', tendermint_block_id = ?, tendermint_transaction_timestamp = ? WHERE tendermint_transaction_id = ?",
		txResult.TxInfo.TxHeight,
		txResult.TxInfo.TxTimestamp,
		txResult.Job.TrustmeshEntry.TendermintTransactionId)
	if result.RowsAffected == 1 {
		fmt.Printf("Tx %v committed \n", txResult.Job.TrustmeshEntry.TendermintTransactionId)
	} else {
		fmt.Printf("Error setting tx status to committed %v\n", result.Error)
	}
}

func getCommittedBaseledgerTransaction(transactionId uuid.UUID) *baseledgertypes.BaseledgerTransaction {
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

	res, err := queryClient.BaseledgerTransaction(context.Background(), &baseledgertypes.QueryGetBaseledgerTransactionRequest{Id: transactionId.String()})

	if err != nil {
		// TODO: error handling
		fmt.Printf("grpc query failed %v\n", err.Error())
		return nil
	}

	fmt.Printf("found baseledger transaction %v\n", res.BaseledgerTransaction)
	return res.BaseledgerTransaction
}
