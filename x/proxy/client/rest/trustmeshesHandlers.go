package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	"github.com/spf13/viper"
	"github.com/unibrightio/baseledger/x/proxy/types"
)

func listTrustmeshEntriesHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbHost, _ := viper.Get("DB_HOST").(string)
		dbPwd, _ := viper.Get("DB_UB_PWD").(string)
		sslMode, _ := viper.Get("DB_SSLMODE").(string)
		dbUser, _ := viper.Get("DB_BASELEDGER_USER").(string)
		dbName, _ := viper.Get("DB_BASELEDGER_NAME").(string)

		args := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=%s",
			dbHost,
			dbUser,
			dbPwd,
			dbName,
			sslMode,
		)

		db, err := gorm.Open("postgres", args)

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

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
