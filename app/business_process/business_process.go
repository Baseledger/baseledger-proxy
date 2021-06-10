package businessprocess

import (
	"fmt"

	"github.com/unibrightio/baseledger/app/types"

	"github.com/jinzhu/gorm"
)

func SetTxStatusToCommitted(txResult types.Result, db *gorm.DB) {
	result := db.Exec("UPDATE trustmesh_entries SET transaction_status = 'COMMITTED', tendermint_block_id = ?, tendermint_transaction_timestamp = ? WHERE tendermint_transaction_id = ?",
		txResult.TxInfo.TxHeight,
		txResult.TxInfo.TxTimestamp,
		txResult.Job.TxHash)
	if result.RowsAffected == 1 {
		fmt.Printf("Tx %v committed \n", txResult.Job.TxHash)
	} else {
		fmt.Printf("Error setting tx status to committed %v\n", result.Error)
	}
}
