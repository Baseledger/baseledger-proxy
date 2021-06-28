package dbutil

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type dbInstance struct {
	db *gorm.DB
}

var Db dbInstance

func InitConnection() {
	fmt.Println("init app db connection")

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
		panic(err)
	}

	fmt.Println("app db connection successful")

	Db = dbInstance{db: db}
}

func (instance *dbInstance) GetConn() *gorm.DB {
	return instance.db
}
