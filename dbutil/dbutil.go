package dbutil

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func InitBaseledgerDBConnection() (db *gorm.DB, err error) {
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

	return gorm.Open("postgres", args)
}
