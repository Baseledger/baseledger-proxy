package cron

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
)

// these structs are related to tendermint jsonrpc
type txResult struct {
	Hash   string `json:"hash"`
	Height string `json:"height"`
}

type header struct {
	Time string `json:"time"`
}

type block struct {
	Header header `json:"header"`
}

type blockResult struct {
	Block block `json:"block"`
}

type txResp struct {
	TxResult txResult `json:"result"`
}

type blockResp struct {
	BlockResult blockResult `json:"result"`
}

// these structs are related to worker pool
type Job struct {
	txHash string
}
type Result struct {
	job    Job
	txInfo TxInfo
}

type TxInfo struct {
	txHeight    string
	txTimestamp string
	txCommitted bool
}

func queryTrustmeshes() {
	fmt.Println("query trustmeshes start")
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
	}

	var trustmeshEntries []proxytypes.TrustmeshEntry
	db.Where("transaction_status='UNCOMMITTED'").Find(&trustmeshEntries)

	fmt.Printf("found %v trustmesh entries\n", len(trustmeshEntries))
	var jobs = make(chan Job, len(trustmeshEntries))
	var results = make(chan Result, len(trustmeshEntries))
	createWorkerPool(1, jobs, results)
	for _, trustmeshEntry := range trustmeshEntries {
		fmt.Printf("creating job for %v\n", trustmeshEntry.TendermintTransactionId)
		job := Job{txHash: trustmeshEntry.TendermintTransactionId}
		jobs <- job
	}

	close(jobs)

	for result := range results {
		fmt.Printf("Tx hash %v, height %v, timestamp %v\n", result.job.txHash, result.txInfo.txHeight, result.txInfo.txTimestamp)
	}
	fmt.Println("query trustmeshes end")
}

func getTxInfo(txHash string) (txInfo *TxInfo, err error) {
	// fetching tx details, if it's not found it will return 500, otherwise 200
	str := "http://localhost:26657/tx?hash=0x" + txHash
	httpRes, err := http.Get(str)
	if err != nil {
		fmt.Printf("error during http tx req %v\n", err)
		return &TxInfo{}, errors.New("get tx info error")
	}

	if httpRes.StatusCode == 500 {
		fmt.Println("tx not committed yet")
		return &TxInfo{txCommitted: false}, errors.New("get tx info error")
	}

	// if it's found should be committed at this point, decode
	if httpRes.StatusCode == 200 {
		var committedTx txResp
		err = json.NewDecoder(httpRes.Body).Decode(&committedTx)
		if err != nil {
			fmt.Printf("error decoding tx %v\n", err)
			return &TxInfo{}, errors.New("error decoding tx")
		}
		// query for block at specific height to find timestamp
		str = "http://localhost:26657/block?height" + committedTx.TxResult.Height
		httpRes, err = http.Get(str)
		if err != nil {
			fmt.Printf("error during http block req %v\n", err)
			return &TxInfo{}, errors.New("get blcok info error")
		}
		var commitedBlock blockResp
		err = json.NewDecoder(httpRes.Body).Decode(&commitedBlock)
		if err != nil {
			fmt.Printf("error decoding block %v\n", err)
			return &TxInfo{}, errors.New("error decoding block")
		}
		// EXAMPLE OF RESULT: DECODED COMMITTED TX HEIGHT 10 AND TIMESTAMP 2021-06-08T16:15:44.8835491Z
		fmt.Printf("DECODED COMMITTED TX HEIGHT %v AND TIMESTAMP %v\n", committedTx.TxResult.Height, commitedBlock.BlockResult.Block.Header.Time)
		return &TxInfo{
			txHeight:    committedTx.TxResult.Height,
			txTimestamp: commitedBlock.BlockResult.Block.Header.Time,
			txCommitted: true,
		}, nil
	}

	// TODO: refactor, this should not happen
	return &TxInfo{}, nil
}

func worker(jobs chan Job, results chan Result) {
	for job := range jobs {
		txInfo, err := getTxInfo(job.txHash)
		if err != nil {
			// TODO: what to do here? it would be http error
			// it seems that we can just let it go through result channel
			fmt.Printf("result error %v\n", err)
		}
		fmt.Printf("result tx %v\n", txInfo)
		output := Result{job: job, txInfo: *txInfo}
		results <- output
	}
	close(results)
}

func createWorkerPool(noOfWorkers int, jobs chan Job, results chan Result) {
	for i := 0; i < noOfWorkers; i++ {
		fmt.Printf("worker no %v\n", i)
		go worker(jobs, results)
	}
}

func StartCron() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().SingletonMode().Do(queryTrustmeshes)

	s.StartAsync()
}
