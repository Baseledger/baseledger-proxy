package rest

import (
	"database/sql"
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
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func listTrustmeshesHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var trustmeshes []types.Trustmesh
		entries := dbutil.Db.GetConn().Order("trustmeshes.created_at ASC")
		// preload seems good enough for now, revisit if it turns out to be performance bottleneck
		dbutil.Paginate(entries, &types.Trustmesh{}, r, w).Preload("Entries").Find(&trustmeshes)

		for i := 0; i < len(trustmeshes); i++ {
			processTrustmesh(&trustmeshes[i])
		}

		res, err := json.Marshal(trustmeshes)
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

func processTrustmesh(trustmesh *types.Trustmesh) {
	if len(trustmesh.Entries) == 0 {
		return
	}
	startTime := trustmesh.Entries[0].TendermintTransactionTimestamp
	endTime := trustmesh.Entries[0].TendermintTransactionTimestamp
	senders := ""
	receivers := ""
	businessObjectTypes := ""
	finalized := false
	containsRejection := false

	for _, entry := range trustmesh.Entries {
		startTime = getBeforeTime(startTime, entry.TendermintTransactionTimestamp)
		endTime = getAfterTime(endTime, entry.TendermintTransactionTimestamp)

		senders = senders + getSeparator(senders) + entry.SenderOrgId.String()
		receivers = receivers + getSeparator(receivers) + entry.ReceiverOrgId.String()
		businessObjectTypes = businessObjectTypes + getSeparator(businessObjectTypes) + entry.BusinessObjectType

		if entry.WorkstepType == "Final" && !finalized {
			finalized = true
		}

		if entry.BaseledgerTransactionType == "Reject" && !containsRejection {
			containsRejection = true
		}
	}

	trustmesh.StartTime = startTime.Time
	trustmesh.EndTime = endTime.Time
	trustmesh.Participants = senders + ", " + receivers
	trustmesh.BusinessObjectTypes = businessObjectTypes
	trustmesh.Finalized = finalized
	trustmesh.ContainsRejections = containsRejection
}

func getSeparator(str string) string {
	if str == "" {
		return ""
	} else {
		return ", "
	}
}

func getBeforeTime(first sql.NullTime, second sql.NullTime) sql.NullTime {
	if !first.Valid {
		return second
	}

	if !second.Valid {
		return first
	}

	if first.Time.Before(second.Time) {
		return first
	}

	return second
}

func getAfterTime(first sql.NullTime, second sql.NullTime) sql.NullTime {
	if !first.Valid {
		return second
	}

	if !second.Valid {
		return first
	}

	if first.Time.Before(second.Time) {
		return second
	}

	return first
}
