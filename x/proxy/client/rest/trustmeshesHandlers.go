package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/types/rest"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	"github.com/unibrightio/baseledger/dbutil"
	"github.com/unibrightio/baseledger/logger"
	"github.com/unibrightio/baseledger/x/proxy/types"
)

func listTrustmeshEntriesHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var trustmeshEntries []types.TrustmeshEntry
		entries := dbutil.Db.GetConn().Find(&trustmeshEntries)

		res, err := json.Marshal(entries)
		if err != nil {
			logger.Errorf("error when getting results from db %v\n", err)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
