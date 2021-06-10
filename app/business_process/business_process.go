package businessprocess

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func SetTxStatusToCommitted(txHash string, db *gorm.DB) {
	result := db.Exec("UPDATE trustmesh_entries SET transaction_status = 'COMMITTED' WHERE tendermint_transaction_id = ?", txHash)
	if result.RowsAffected == 1 {
		fmt.Printf("Tx %v committed \n", txHash)
	} else {
		fmt.Printf("Error setting tx status to committed %v\n", result.Error)
	}
}
