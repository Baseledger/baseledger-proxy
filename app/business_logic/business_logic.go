package businesslogic

import (
	"fmt"

	"github.com/unibrightio/baseledger/app/types"

	"github.com/jinzhu/gorm"
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
		// proxy.SendOffchainProcessMessage(offchainMessage, txResult.Job.TrustmeshEntry.Receiver)
	case "SuggestionReceived":
		fmt.Println("SuggestionReceived")
	case "FeedbackSent":
		fmt.Println("FeedbackSent")
		// send offchain msg to sender of sugestion
	case "FeedbackReceived":
		fmt.Println("FeedbackReceived")
	default:
		// TODO panic
		fmt.Println("UNKNOWN BUSINESS LOGIC")
	}
}
