package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

func GetTrustmeshesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var trustmeshes []types.Trustmesh
		db := dbutil.Db.GetConn().Order("trustmeshes.created_at ASC")
		// preload seems good enough for now, revisit if it turns out to be performance bottleneck
		dbutil.Paginate(c, db, &types.Trustmesh{}).Preload("Entries").Find(&trustmeshes)

		for i := 0; i < len(trustmeshes); i++ {
			processTrustmesh(&trustmeshes[i])
		}

		restutil.Render(trustmeshes, 200, c)
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

		if entry.WorkstepType == "FinalWorkstep" && !finalized {
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
