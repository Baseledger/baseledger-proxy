package types

import (
	proxytypes "github.com/unibrightio/baseledger/x/proxy/types"
)

// these structs are related to tendermint jsonrpc
type TxResult struct {
	Hash   string `json:"hash"`
	Height string `json:"height"`
}

type Header struct {
	Time string `json:"time"`
}

type Block struct {
	Header Header `json:"header"`
}

type BlockResult struct {
	Block Block `json:"block"`
}

type TxResp struct {
	TxResult TxResult `json:"result"`
}

type BlockResp struct {
	BlockResult BlockResult `json:"result"`
}

// these structs are related to worker pool
type Job struct {
	TrustmeshEntry proxytypes.TrustmeshEntry
}
type Result struct {
	Job    Job
	TxInfo TxInfo
}

type TxInfo struct {
	TxHeight    string
	TxTimestamp string
	TxCommitted bool
}
