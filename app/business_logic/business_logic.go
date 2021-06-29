package businesslogic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/baseledger/app/types"
	common "github.com/unibrightio/baseledger/common"
	"github.com/unibrightio/baseledger/dbutil"
	"github.com/unibrightio/baseledger/sor"
	baseledgertypes "github.com/unibrightio/baseledger/x/baseledger/types"
	"github.com/unibrightio/baseledger/x/proxy/proxy"
	restutil "github.com/unibrightio/baseledger/x/proxy/restutil"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
	"google.golang.org/grpc"
)

func ExecuteBusinessLogic(txResult types.Result) {
	if txResult.TxInfo.TxHeight == "" || txResult.TxInfo.TxTimestamp == "" {
		return
	}
	fmt.Printf("Execute business logic for result %v\n", txResult)
	offchainMessage, err := proxytypes.GetOffchainMsgById(txResult.Job.TrustmeshEntry.OffchainProcessMessageId)
	if err != nil {
		// TODO: logging
		fmt.Println("Offchain process msg not found")
		return
	}
	switch txResult.Job.TrustmeshEntry.EntryType {
	case common.SuggestionSentTrustmeshEntryType:
		fmt.Println(common.SuggestionSentTrustmeshEntryType)
		// TODO: Ognjen, use msg client
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.ReceiverOrgId.String(), txResult.Job.TrustmeshEntry.TransactionHash)
	case common.SuggestionReceivedTrustmeshEntryType:
		fmt.Println(common.SuggestionReceivedTrustmeshEntryType)
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
	case common.FeedbackSentTrustmeshEntryType:
		fmt.Println(common.FeedbackSentTrustmeshEntryType)
		// TODO: Ognjen, use msg client
		proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.SenderOrgId.String(), txResult.Job.TrustmeshEntry.TransactionHash)
	case common.FeedbackReceivedTrustmeshEntryType:
		fmt.Println(common.FeedbackReceivedTrustmeshEntryType)
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			// TODO: logging
			return
		}

		sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
	default:
		fmt.Printf("unknown business process %v\n", txResult.Job.TrustmeshEntry.EntryType)
		panic(errors.New("uknown business process!"))
	}
	setTxStatusToCommitted(txResult)
}

func setTxStatusToCommitted(txResult types.Result) {
	result := dbutil.Db.GetConn().Exec("UPDATE trustmesh_entries SET commitment_state = ?, tendermint_block_id = ?, tendermint_transaction_timestamp = ? WHERE tendermint_transaction_id = ?",
		common.CommittedCommitmentState,
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
	// TODO: BAS-33
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
