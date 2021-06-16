package businessprocess

import (
	"fmt"

	"github.com/unibrightio/baseledger/app/types"

	"github.com/jinzhu/gorm"
	uuid "github.com/kthomas/go.uuid"

	proxy "github.com/unibrightio/baseledger/x/proxy/proxy"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
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
	case "Suggest":
		fmt.Println("SUGGEST BUSINESS LOGIC")
		offchainMessage := createSuggestOffchainMessage(txResult)
		proxy.SendOffchainProcessMessage(offchainMessage, txResult.Job.TrustmeshEntry.Receiver)

		// JUST FOR TESTING
		proxy.OffchainProcessMessageReceived(offchainMessage)
	case "Feedback":
		fmt.Println("FEEDBACK BUSSINESS LOGIC")
	default:
		// TODO panic
		fmt.Println("UNKNOWN BUSINESS LOGIC")
	}
}

func createSuggestOffchainMessage(txResult types.Result) proxytypes.OffchainProcessMessage {
	offchainMsgId, _ := uuid.NewV4()
	// var offchainProcessMessage = new OffchainProcessMessage(workstepType, String.Empty, businessObject, hashOfBusinessObject, baseledgerBusinessObjectID, referencedBaseledgerBusinessObjectID, workstepType.ToString() + " suggested");
	offchainMessage := proxytypes.OffchainProcessMessage{
		OffchainProcessMessageId:             offchainMsgId.String(),
		WorkstepType:                         txResult.Job.TrustmeshEntry.WorkstepType,
		ReferencedOffchainProcessMessage:     "",
		BusinessObject:                       "we need to store json in db?",
		Hash:                                 "we need to calculate hash here using proxy based on json above?",
		BaseledgerBusinessObjectId:           txResult.Job.TrustmeshEntry.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: txResult.Job.TrustmeshEntry.ReferencedBaseledgerBusinessObjectId,
		StatusTextMessage:                    txResult.Job.TrustmeshEntry.WorkstepType + " suggested",
		BaseledgerTransactionIdOfStoredProof: txResult.Job.TrustmeshEntry.BaseledgerTransactionId,
	}

	return offchainMessage
}
