package rest

import (
	"encoding/json"
	"fmt"
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
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func listTrustmeshesHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		var trustmeshes []types.Trustmesh
		entries := dbutil.Db.GetConn().Order("trustmeshes.created_at ASC")
		db, totalResults := dbutil.Paginate(entries, &types.Trustmesh{}, queryParams.Get("pageNum"), queryParams.Get("pageSize"))
		db.Find(&trustmeshes)

		db = dbutil.Db.GetConn()
		var trustmeshEntries []types.TrustmeshEntry
		for i := 0; i < len(trustmeshes); i++ {
			db.Where("trustmesh_id = ?", trustmeshes[i].Id).Find(&trustmeshEntries)
			trustmeshes[i].Entries = trustmeshEntries[:]
		}

		res, err := json.Marshal(trustmeshes)
		if err != nil {
			logger.Errorf("error when getting results from db %v\n", err)
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-total-results-count", fmt.Sprintf("%d", *totalResults))
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
