package businesslogic

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"

	uuid "github.com/kthomas/go.uuid"
	common "github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/eth"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/proxyutil"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/synctree"
	systemofrecord "github.com/unibrightio/proxy-api/systemofrecord"
	"github.com/unibrightio/proxy-api/types"
	proxytypes "github.com/unibrightio/proxy-api/types"
)

func ExecuteBusinessLogic(txResult proxytypes.Result) {
	systemofrecord.InitClient()
	var trustmeshEntry = txResult.Job.TrustmeshEntry

	if txResult.TxInfo.TxHeight == "" || txResult.TxInfo.TxTimestamp == "" {
		logger.Infof("Transaction %v not yet committed", trustmeshEntry.TransactionHash)
		return
	}
	if !txResult.TxInfo.TxValid {
		logger.Warnf("Transaction %v is invalid with code %v and log %v", trustmeshEntry.TransactionHash, txResult.TxInfo.TxCode, txResult.TxInfo.TxLog)

		systemofrecord.TriggerSorWebhook(
			types.UpdateObject,
			&trustmeshEntry,
			"",
			false,
			"Transaction is invalid",
			"")

		setTxStatus(txResult, common.InvalidCommitmentState)
		return
	}
	logger.Infof("Execute business logic for result %v\n", txResult)
	offchainMessage, err := proxytypes.GetOffchainMsgById(trustmeshEntry.OffchainProcessMessageId)
	if err != nil {
		logger.Error("Offchain process msg not found")
		systemofrecord.TriggerSorWebhook(
			types.UpdateObject,
			&trustmeshEntry,
			"",
			false,
			"Offchain process msg not found",
			"")
		return
	}
	switch trustmeshEntry.EntryType {
	case common.SuggestionSentTrustmeshEntryType:
		logger.Info(common.SuggestionSentTrustmeshEntryType)

		// TODO: bellow lines should be moved to send offchain message
		var natsMessage proxytypes.NatsMessage
		natsMessage.ProcessMessage = *offchainMessage
		natsMessage.TxHash = trustmeshEntry.TransactionHash

		var payload, _ = json.Marshal(natsMessage)

		proxyutil.SendOffchainMessage(payload, trustmeshEntry.WorkgroupId.String(), trustmeshEntry.ReceiverOrgId.String(), common.BaseledgerNatsSubject)

		systemofrecord.TriggerSorWebhook(
			types.UpdateObject,
			&trustmeshEntry,
			"",
			true,
			"success",
			"")

	case common.SuggestionReceivedTrustmeshEntryType:
		logger.Info(common.SuggestionReceivedTrustmeshEntryType)
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			logger.Error("Failed to get committed baseledger transaction")
			return
		}
		baseledgerTransactionPayload := proxytypes.BaseledgerTransactionPayload{}
		deprivitizedPayload := proxyutil.DeprivatizeBaseledgerTransactionPayload(baseledgerTransaction.Payload, trustmeshEntry.WorkgroupId)
		err = json.Unmarshal(([]byte)(deprivitizedPayload), &baseledgerTransactionPayload)
		if err != nil {
			logger.Error("Failed to unmarshal baseledger transaction payload")
			return
		}

		if synctree.VerifyHashMatch(baseledgerTransactionPayload.Proof, offchainMessage.BusinessObjectProof, offchainMessage.BaseledgerSyncTreeJson) {
			logger.Info("Hashes match, processing feedback")

			syncTree := &synctree.BaseledgerSyncTree{}
			err = json.Unmarshal([]byte(offchainMessage.BaseledgerSyncTreeJson), &syncTree)
			if err != nil {
				logger.Errorf("Error unmarshalling sync tree", err.Error())
				return
			}
			logger.Infof("Sync tree unmarshalled", syncTree)

			boJson := synctree.GetBusinessObjectJson(*syncTree)
			logger.Infof("Business object sync tree json", boJson)

			systemofrecord.TriggerSorWebhook(
				types.CreateObject,
				&trustmeshEntry,
				boJson,
				false,
				offchainMessage.StatusTextMessage,
				offchainMessage.SenderId.String(),
			)
			break
		}
		logger.Warnf("Hashes don't match, rejecting feedback %v %v %v", baseledgerTransactionPayload.Proof, offchainMessage.BusinessObjectProof, offchainMessage.BaseledgerSyncTreeJson)
		restutil.SendRejectFeedback(offchainMessage, trustmeshEntry.WorkgroupId.String())
	case common.FeedbackSentTrustmeshEntryType:
		logger.Info(common.FeedbackSentTrustmeshEntryType)

		// TODO: bellow lines should be moved to send offchain message
		var natsMessage proxytypes.NatsMessage
		natsMessage.ProcessMessage = *offchainMessage
		natsMessage.TxHash = trustmeshEntry.TransactionHash

		var payload, _ = json.Marshal(natsMessage)

		proxyutil.SendOffchainMessage(payload, trustmeshEntry.WorkgroupId.String(), trustmeshEntry.ReceiverOrgId.String(), common.BaseledgerNatsSubject)

		systemofrecord.TriggerSorWebhook(
			types.UpdateObject,
			&trustmeshEntry,
			"",
			true,
			"success",
			"")

	case common.FeedbackReceivedTrustmeshEntryType:
		logger.Info(common.FeedbackReceivedTrustmeshEntryType)
		baseledgerTransaction := getCommittedBaseledgerTransaction(offchainMessage.BaseledgerTransactionIdOfStoredProof)
		if baseledgerTransaction == nil {
			logger.Error("Failed to get committed baseledger transaction")
			return
		}

		syncTree := &synctree.BaseledgerSyncTree{}
		err = json.Unmarshal([]byte(offchainMessage.BaseledgerSyncTreeJson), &syncTree)
		if err != nil {
			logger.Errorf("Error unmarshalling sync tree", err.Error())
			return
		}
		logger.Infof("Sync tree unmarshalled", syncTree)

		// type? is it possible in go?
		// do we need it if we just pass this to sor?
		var bo map[string]interface{}
		boJson := synctree.GetBusinessObjectJson(*syncTree)
		err = json.Unmarshal([]byte(boJson), &bo)
		if err != nil {
			logger.Errorf("Error unmarshalling sync tree", err.Error())
			return
		}
		logger.Infof("Business object unmarshalled", bo)
		status := true
		if trustmeshEntry.BaseledgerTransactionType == common.BaseledgerTransactionTypeReject {
			status = false
		}

		logger.Infof("Sending feedback received status update %v\n", status)

		systemofrecord.TriggerSorWebhook(
			types.UpdateObject,
			&trustmeshEntry,
			"",
			status,
			offchainMessage.StatusTextMessage,
			offchainMessage.SenderId.String())

		if status == true {
			tryExitToEth(&trustmeshEntry)
		}

	default:
		logger.Errorf("unknown business process %v\n", trustmeshEntry.EntryType)
		panic(errors.New("uknown business process!"))
	}

	setTxStatus(txResult, common.CommittedCommitmentState)
}

func tryExitToEth(trustmeshEntry *types.TrustmeshEntry) {
	trustmesh, err := types.GetTrustmeshById(trustmeshEntry.TrustmeshId)

	if err != nil {
		logger.Errorf("Could not find trustmesh for exiting %v", trustmeshEntry.TrustmeshId)
		return
	}

	if trustmesh.EthExitTxHash != "" {
		logger.Warnf("Trustmesh %v already exited", trustmesh.Id)
		return
	}

	// looking into trustmesh entries, we are trying to find suggestion sent for final workstep
	// that has baseledger business object id that feedback is approving
	containsFinal := false
	for _, entry := range trustmesh.Entries {
		if entry.WorkstepType == common.WorkstepTypeFinal &&
			entry.EntryType == common.SuggestionSentTrustmeshEntryType &&
			entry.BaseledgerBusinessObjectId == trustmeshEntry.ReferencedBaseledgerBusinessObjectId {
			containsFinal = true
			break
		}
	}

	if !containsFinal {
		logger.Infof("Current trustmesh %v does not contain final workstep, skipping exit", trustmeshEntry.TrustmeshId)
		return
	}

	logger.Infof("Current trustmesh %v contains final workstep, exiting", trustmeshEntry.TrustmeshId)

	trustmeshSyncTree := synctree.CreateFromTrustmesh(*trustmesh)
	transactionId := uuid.NewV4()

	payload := proxyutil.CreateExitBaseledgerTransactionPayload(trustmeshEntry.WorkgroupId, transactionId, trustmeshSyncTree.RootProof)
	signAndBroadcastPayload := &restutil.SignAndBroadcastPayload{
		OpCode:        0,
		TransactionId: transactionId.String(),
		Payload:       payload,
	}

	txHash := restutil.SignAndBroadcast(*signAndBroadcastPayload)
	if txHash == nil {
		logger.Error("Sign and broadcast exiting baseledger transaction failed")
		return
	}

	eth.StoreExitProofInTrustmeshAndInformCounterparty(transactionId.String(), trustmeshSyncTree.RootProof, trustmeshEntry)
}

func setTxStatus(txResult proxytypes.Result, commitmentState string) {
	result := dbutil.Db.GetConn().Exec("UPDATE trustmesh_entries SET commitment_state = ?, tendermint_block_id = ?, tendermint_transaction_timestamp = ? WHERE tendermint_transaction_id = ?",
		commitmentState,
		txResult.TxInfo.TxHeight,
		txResult.TxInfo.TxTimestamp,
		txResult.Job.TrustmeshEntry.TendermintTransactionId)
	if result.RowsAffected == 1 {
		logger.Infof("Tx %v committed \n", txResult.Job.TrustmeshEntry.TendermintTransactionId)
	} else {
		logger.Errorf("Error setting tx status to committed %v\n", result.Error)
	}
}

func getCommittedBaseledgerTransaction(id uuid.UUID) *proxytypes.BaseledgerTransactionDto {
	resp, err := http.Get("http://" + viper.Get("BLOCKCHAIN_APP_URL").(string) + "/unibrightio/baseledger/baseledger/BaseledgerTransaction/" + id.String())

	if err != nil {
		logger.Errorf("error while fetching committed baseledger transaction %v\n", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("error while reading committed baseledger transaction response %v\n", err.Error())
		return nil
	}

	logger.Infof("trying to unmarshal body: %v to CommittedBaseledgerTransactionResponse", string(body))

	var transactionResponse proxytypes.CommittedBaseledgerTransactionResponse
	err = json.Unmarshal(body, &transactionResponse)

	if err != nil {
		logger.Errorf("error while unmarshalling fetched committed baseledger transaction %v\n", err.Error())
		return nil
	}

	logger.Infof("commited baseledger transaction ", transactionResponse)
	return &transactionResponse.BaseledgerTransaction
}
