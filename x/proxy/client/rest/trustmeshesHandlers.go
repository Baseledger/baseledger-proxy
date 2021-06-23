package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/types/rest"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	"github.com/unibrightio/baseledger/dbutil"
	"github.com/unibrightio/baseledger/x/proxy/types"
)

func listTrustmeshEntriesHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := dbutil.InitBaseledgerDBConnection()

		if err != nil {
			fmt.Printf("error when connecting to db %v\n", err)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		var trustmeshEntries []types.TrustmeshEntry
		entries := db.Find(&trustmeshEntries)

		res, err := json.Marshal(entries)
		if err != nil {
			fmt.Printf("error when getting results from db %v\n", err)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		db.Close()

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
