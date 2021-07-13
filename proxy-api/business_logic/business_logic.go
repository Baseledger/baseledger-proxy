package businesslogic

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	uuid "github.com/kthomas/go.uuid"
	common "github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
	proxytypes "github.com/unibrightio/proxy-api/types"
)

func ExecuteBusinessLogic(txResult proxytypes.Result) {
	if txResult.TxInfo.TxHeight == "" || txResult.TxInfo.TxTimestamp == "" {
		return
	}
	logger.Infof("Execute business logic for result %v\n", txResult)
	offchainMessage, err := proxytypes.GetOffchainMsgById(txResult.Job.TrustmeshEntry.OffchainProcessMessageId)
	if err != nil {
		// TODO: logging
		logger.Error("Offchain process msg not found")
		return
	}
	switch txResult.Job.TrustmeshEntry.EntryType {
	case common.SuggestionSentTrustmeshEntryType:
		logger.Info(common.SuggestionSentTrustmeshEntryType)
		// TODO: Ognjen, use msg client
		// proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.ReceiverOrgId.String(), txResult.Job.TrustmeshEntry.TransactionHash)
	case common.SuggestionReceivedTrustmeshEntryType:
		logger.Info(common.SuggestionReceivedTrustmeshEntryType)
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			return
		}
		baseledgerTransactionPayload := proxytypes.BaseledgerTransactionPayload{}
		// deprivitizedPayload := proxy.DeprivatizeBaseledgerTransactionPayload(baseledgerTransaction.Payload, txResult.Job.TrustmeshEntry.WorkgroupId)
		deprivitizedPayload := ""
		err = json.Unmarshal(([]byte)(deprivitizedPayload), &baseledgerTransactionPayload)
		if err != nil {
			logger.Error("Failed to unmarshal baseledger transaction payload")
			return
		}

		// offchainMessageBusinessObjectProof := proxy.CreateHashFromBusinessObject(offchainMessage.BaseledgerSyncTreeJson)
		offchainMessageBusinessObjectProof := ""
		if baseledgerTransactionPayload.Proof == offchainMessage.BusinessObjectProof && baseledgerTransactionPayload.Proof == offchainMessageBusinessObjectProof {
			logger.Info("Hashes match, processing feedback")
			// sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
			return
		}
		logger.Warn("Hashes don't match")
		// restutil.RejectFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId.String())
	case common.FeedbackSentTrustmeshEntryType:
		logger.Info(common.FeedbackSentTrustmeshEntryType)
		// TODO: Ognjen, use msg client
		// proxy.SendOffchainProcessMessage(*offchainMessage, txResult.Job.TrustmeshEntry.SenderOrgId.String(), txResult.Job.TrustmeshEntry.TransactionHash)
	case common.FeedbackReceivedTrustmeshEntryType:
		logger.Info(common.FeedbackReceivedTrustmeshEntryType)
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			return
		}

		// sor.ProcessFeedback(*offchainMessage, txResult.Job.TrustmeshEntry.WorkgroupId, baseledgerTransaction.Payload)
	default:
		logger.Errorf("unknown business process %v\n", txResult.Job.TrustmeshEntry.EntryType)
		panic(errors.New("uknown business process!"))
	}
	setTxStatusToCommitted(txResult)
}

func setTxStatusToCommitted(txResult proxytypes.Result) {
	result := dbutil.Db.GetConn().Exec("UPDATE trustmesh_entries SET commitment_state = ?, tendermint_block_id = ?, tendermint_transaction_timestamp = ? WHERE tendermint_transaction_id = ?",
		common.CommittedCommitmentState,
		txResult.TxInfo.TxHeight,
		txResult.TxInfo.TxTimestamp,
		txResult.Job.TrustmeshEntry.TendermintTransactionId)
	if result.RowsAffected == 1 {
		logger.Infof("Tx %v committed \n", txResult.Job.TrustmeshEntry.TendermintTransactionId)
	} else {
		logger.Errorf("Error setting tx status to committed %v\n", result.Error)
	}
}

// TODO: this needs to be tested, just building ok for now
func getCommittedBaseledgerTransaction(id uuid.UUID) *proxytypes.BaseledgerTransaction {
	resp, err := http.Get("http://localhost:1317/proxy/committedTx/" + id.String())

	if err != nil {
		logger.Errorf("error while fetching committed baseledger transaction %v\n", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	var transaction proxytypes.BaseledgerTransaction
	err = json.Unmarshal(body, &transaction)

	if err != nil {
		logger.Errorf("error while unmarshalling fetched committed baseledger transaction %v\n", err.Error())
		return nil
	}

	logger.Infof("get committed baseleger transaction", transaction)
	return &transaction
}
