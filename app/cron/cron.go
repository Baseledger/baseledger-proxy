package cron

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"

	businessprocess "github.com/unibrightio/baseledger/app/business_process"
	"github.com/unibrightio/baseledger/app/types"
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"

	"github.com/unibrightio/baseledger/dbutil"
)

func queryTrustmeshes() {
	fmt.Println("query trustmeshes start")
	db, err := dbutil.InitBaseledgerDBConnection()

	if err != nil {
		fmt.Printf("error when connecting to db %v\n", err)
	}

	var trustmeshEntries []proxytypes.TrustmeshEntry
	db.Where("transaction_status='UNCOMMITTED'").Find(&trustmeshEntries)

	fmt.Printf("found %v trustmesh entries\n", len(trustmeshEntries))
	var jobs = make(chan types.Job, len(trustmeshEntries))
	var results = make(chan types.Result, len(trustmeshEntries))
	createWorkerPool(1, jobs, results)

	for _, trustmeshEntry := range trustmeshEntries {
		fmt.Printf("creating job for %v\n", trustmeshEntry.TendermintTransactionId)
		job := types.Job{TxHash: trustmeshEntry.TendermintTransactionId}
		jobs <- job
	}
	close(jobs)

	for result := range results {
		fmt.Printf("Tx hash %v, height %v, timestamp %v\n", result.Job.TxHash, result.TxInfo.TxHeight, result.TxInfo.TxTimestamp)
		if result.TxInfo.TxHeight != "" && result.TxInfo.TxTimestamp != "" {
			businessprocess.SetTxStatusToCommitted(result, db)
		}
	}
	fmt.Println("query trustmeshes end")
}

func getTxInfo(txHash string) (txInfo *types.TxInfo, err error) {
	// fetching tx details
	str := "http://localhost:26657/tx?hash=0x" + txHash
	httpRes, err := http.Get(str)
	if err != nil {
		fmt.Printf("error during http tx req %v\n", err)
		return &types.TxInfo{}, errors.New("get tx info error")
	}

	// if transaction is not found it is not yet committed
	if httpRes.StatusCode != 200 {
		fmt.Println("tx not committed yet")
		return &types.TxInfo{TxCommitted: false}, errors.New("get tx info error")
	}

	// if it's found should be committed at this point, decode
	var committedTx types.TxResp
	err = json.NewDecoder(httpRes.Body).Decode(&committedTx)
	if err != nil {
		fmt.Printf("error decoding tx %v\n", err)
		return &types.TxInfo{}, errors.New("error decoding tx")
	}
	// query for block at specific height to find timestamp
	str = "http://localhost:26657/block?height" + committedTx.TxResult.Height
	httpRes, err = http.Get(str)
	if err != nil {
		fmt.Printf("error during http block req %v\n", err)
		return &types.TxInfo{}, errors.New("get blcok info error")
	}
	var commitedBlock types.BlockResp
	err = json.NewDecoder(httpRes.Body).Decode(&commitedBlock)
	if err != nil {
		fmt.Printf("error decoding block %v\n", err)
		return &types.TxInfo{}, errors.New("error decoding block")
	}
	fmt.Printf("DECODED COMMITTED TX HEIGHT %v AND TIMESTAMP %v\n", committedTx.TxResult.Height, commitedBlock.BlockResult.Block.Header.Time)
	return &types.TxInfo{
		TxHeight:    committedTx.TxResult.Height,
		TxTimestamp: commitedBlock.BlockResult.Block.Header.Time,
		TxCommitted: true,
	}, nil
}

func worker(jobs chan types.Job, results chan types.Result) {
	defer close(results)
	for job := range jobs {
		txInfo, err := getTxInfo(job.TxHash)
		if err != nil {
			// here it would be http error
			// it seems that we can just let it go through result channel
			fmt.Printf("result error %v\n", err)
		}
		fmt.Printf("result tx %v\n", txInfo)
		output := types.Result{Job: job, TxInfo: *txInfo}
		results <- output
	}
}

func createWorkerPool(noOfWorkers int, jobs chan types.Job, results chan types.Result) {
	for i := 0; i < noOfWorkers; i++ {
		go worker(jobs, results)
	}
}

func StartCron() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().SingletonMode().Do(queryTrustmeshes)

	s.StartAsync()
}
