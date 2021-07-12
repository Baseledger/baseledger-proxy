package dbutil

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/unibrightio/baseledger/logger"
)

const defaultResultsPerPage = 5

type dbInstance struct {
	db *gorm.DB
}

var Db dbInstance

func InitConnection() {
	logger.Info("init app db connection")

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
		logger.Errorf("error when connecting to db %v\n", err)
		panic(err)
	}

	logger.Info("app db connection successful")

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
		logger.Errorf("db connection failed %v\n", err.Error())
		panic(err)
	}

	logger.Info("admin db connection successful")

	exists := db.Exec(fmt.Sprintf("select 1 from pg_roles where rolname='%s'", dbUser))

	if exists.RowsAffected == 1 {
		logger.Errorf("row already exits")
		return
	}

	result := db.Exec(fmt.Sprintf("create user %s with superuser password '%s'", dbUser, dbPwd))

	if result.Error != nil {
		logger.Errorf("failed to create user %v\n", result.Error)
		panic(result.Error)
	}

	result = db.Exec(fmt.Sprintf("create database %s owner %s", dbName, dbUser))

	if result.Error != nil {
		logger.Errorf("failed to create baseledger db %v\n", result.Error)
		panic(result.Error)
	}

	db.DB().Close() // Close admin connection
}

// TODO: check if we can reuse db connection from method above
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
		logger.Errorf("migrations failed 1: %s", err.Error())
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Errorf("migrations failed 2: %s", err.Error())
		panic(err)
	}

	// TODO: These have to be some kind of embeded resources
	m, err := migrate.NewWithDatabaseInstance("file://./ops/migrations", dbName, driver)
	if err != nil {
		logger.Errorf("migrations failed 3: %s", err.Error())
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Errorf("migrations failed 4: %s", err.Error())
	}
}

func Paginate(db *gorm.DB, model interface{}, request *http.Request, w http.ResponseWriter) (query *gorm.DB) {
	pageNum := int64(1)
	pageSize := int64(defaultResultsPerPage)
	queryParams := request.URL.Query()

	if queryParams.Get("pageNum") != "" {
		if _page, err := strconv.ParseInt(queryParams.Get("pageNum"), 10, 64); err == nil {
			pageNum = _page
		}
	}

	if queryParams.Get("pageSize") != "" {
		if _rpp, err := strconv.ParseInt(queryParams.Get("pageSize"), 10, 64); err == nil {
			pageSize = _rpp
		}
	}

	totalResults := 0
	db.Model(model).Count(&totalResults)
	w.Header().Set("x-total-results-count", fmt.Sprintf("%d", totalResults))

	return db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
}
