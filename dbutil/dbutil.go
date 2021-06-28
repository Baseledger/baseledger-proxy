package dbutil

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
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

func InitDbIfNotExists() {
	dbHost, _ := viper.Get("DB_HOST").(string)
	dbSuperUser, _ := viper.Get("DB_UB_USER").(string)
	dbPwd, _ := viper.Get("DB_UB_PWD").(string)
	dbDefaultName, _ := viper.Get("DB_UB_NAME").(string)
	sslMode, _ := viper.Get("DB_SSLMODE").(string)
	dbUser, _ := viper.Get("DB_BASELEDGER_USER").(string)
	dbName, _ := viper.Get("DB_BASELEDGER_NAME").(string)

	args := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbSuperUser,
		dbPwd,
		dbDefaultName,
		sslMode,
	)
	db, err := gorm.Open("postgres", args)

	if err != nil {
		fmt.Printf("db connection failed %v\n", err.Error())
		panic(err)
	}

	fmt.Printf("admin db connection successful")

	exists := db.Exec(fmt.Sprintf("select 1 from pg_roles where rolname='%s'", dbUser))

	if exists.RowsAffected == 1 {
		fmt.Printf("row already exits")
		return
	}

	result := db.Exec(fmt.Sprintf("create user %s with superuser password '%s'", dbUser, dbPwd))

	if result.Error != nil {
		fmt.Printf("failed to create user %v\n", result.Error)
		panic(result.Error)
	}

	result = db.Exec(fmt.Sprintf("create database %s owner %s", dbName, dbUser))

	if result.Error != nil {
		fmt.Printf("failed to create baseledger db %v\n", result.Error)
		panic(result.Error)
	}

	db.DB().Close() // Close admin connection
}

func PerformMigrations() {
	dbHost, _ := viper.Get("DB_HOST").(string)
	dbPwd, _ := viper.Get("DB_UB_PWD").(string)
	sslMode, _ := viper.Get("DB_SSLMODE").(string)
	dbUser, _ := viper.Get("DB_BASELEDGER_USER").(string)
	dbName, _ := viper.Get("DB_BASELEDGER_NAME").(string)

	dsn := fmt.Sprintf("postgres://%s/%s?user=%s&password=%s&sslmode=%s",
		dbHost,
		dbName,
		dbUser,
		dbPwd,
		sslMode,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("migrations failed 1: %s", err.Error())
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Printf("migrations failed 2: %s", err.Error())
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./ops/migrations", dbName, driver)
	if err != nil {
		fmt.Printf("migrations failed 3: %s", err.Error())
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Printf("migrations failed 4: %s", err.Error())
	}
}
