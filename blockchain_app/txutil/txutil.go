package txutil

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unibrightio/baseledger/logger"
)

const (
	errCodeMismatch = 32
)

var (
	// errors are of the form:
	// "account sequence mismatch, expected 25, got 27: incorrect account sequence"
	recoverRegexp = regexp.MustCompile(`^account sequence mismatch, expected (\d+), got (\d+):`)
)

func BuildClientCtx(clientCtx client.Context) (*client.Context, error) {
	keyring, err := NewKeyringInstance()

	keysList, err := keyring.List()
	if err != nil {
		logger.Errorf("error getting key list %v\n", err.Error())
		return nil, errors.New("")
	}

	if len(keysList) == 0 {
		logger.Errorf("key list empty %v\n", err.Error())
		return nil, errors.New("")
	}

	if err != nil {
		logger.Errorf("error getting key %v\n", err.Error())
		return nil, errors.New("")
	}

	// every node should configure key for this purpose, and it should be first in key list
	logger.Infof("key found %v %v\n", keysList[0], keysList[0].GetName())

	clientCtx = clientCtx.
		WithKeyring(keyring).
		WithFromAddress(keysList[0].GetAddress()).
		WithSkipConfirmation(true).
		WithFromName(keysList[0].GetName()).
		WithBroadcastMode("sync")

	return &clientCtx, nil
}

// TODO: change test keyring with other (file?) - new ticket for this
func NewKeyringInstance() (keyring.Keyring, error) {
	kr, err := keyring.New("baseledger", "test", "~/.baseledger", nil)

	if err != nil {
		logger.Errorf("error fetching test keyring %v\n", err.Error())
		return nil, errors.New("error fetching key ring")
	}

	return kr, nil
}

func SignTxAndGetTxBytes(clientCtx client.Context, msg sdk.Msg, accNum uint64, accSeq uint64) ([]byte, error) {
	logger.Infof("retrieved account %v %v\n", accNum, accSeq)
	txFactory := tx.Factory{}.
		WithChainID("baseledger").
		WithGas(1000000). // hardcoding gasWanted to high number since fees will allways be 1 token
		WithTxConfig(clientCtx.TxConfig).
		WithAccountNumber(accNum).
		WithSequence(accSeq).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring)

	txFactory, err := tx.PrepareFactory(clientCtx, txFactory)
	if err != nil {
		logger.Errorf("prepare factory error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	transaction, err := tx.BuildUnsignedTx(txFactory, msg)
	if err != nil {
		logger.Errorf("build unsigned tx error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	err = tx.Sign(txFactory, clientCtx.GetFromName(), transaction, true)
	if err != nil {
		logger.Errorf("sign tx error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(transaction.GetTx())
	if err != nil {
		logger.Errorf("tx encoder %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	return txBytes, nil
}

func BroadcastAndGetTxHash(clientCtx client.Context, msg sdk.Msg, accNum uint64, accSeq uint64, retried bool) (*string, error) {
	txBytes, err := SignTxAndGetTxBytes(clientCtx, msg, accNum, accSeq)
	if err != nil {
		return nil, err
	}
	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		logger.Errorf("error while broadcasting tx %v\n", err.Error())
		return nil, err
	}

	// if broadcast was successful, return txHash
	if res.Code == 0 {
		logMsg := "BROADCASTED"
		if retried {
			logMsg = "REBROADCASTED"
		}
		logger.Infof("TRANSACTION %v WITH RESULT %v\n", logMsg, res)
		return &res.TxHash, nil
	}

	// if code is not 0 and is not missmatch, don't handle it and return
	if res.Code != errCodeMismatch {
		return nil, err
	}

	if res.Code != 0 && res.Code != errCodeMismatch {
		logger.Errorf("broadcast failed with code different than missmatch %v %v\n", res.Code, res)
		return nil, err
	}

	// if code is missmatch and it is retrying, don't handle it and return
	if retried {
		return nil, errors.New("broadcast failed after retrying")
	}

	// if code is missmatch first time, parse log and try again
	logger.Infof("ACCOUNT SEQUENCE MISSMATCH %v\n", res.RawLog)

	nextSequence, ok := parseNextSequence(accSeq, res.RawLog)

	if !ok {
		return nil, errors.New("broadcast failed when parsing sequence")
	}

	logger.Infof("RETRYING WITH SEQUENCE %v\n", nextSequence)

	return BroadcastAndGetTxHash(clientCtx, msg, accNum, nextSequence, true)
}

func parseNextSequence(current uint64, message string) (uint64, bool) {
	// "account sequence mismatch, expected 25, got 27: incorrect account sequence"
	matches := recoverRegexp.FindStringSubmatch(message)

	if len(matches) != 3 {
		return 0, false
	}

	if len(matches[1]) == 0 || len(matches[2]) == 0 {
		return 0, false
	}

	expected, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil || expected == 0 {
		return 0, false
	}

	received, err := strconv.ParseUint(matches[2], 10, 64)
	if err != nil || received == 0 {
		return 0, false
	}

	if received != current {
		return expected, true
	}

	return expected, true
}
