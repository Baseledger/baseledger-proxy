package txutil

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BuildClientCtx(clientCtx client.Context, from string) (*client.Context, error) {
	fromAddress, err := sdk.AccAddressFromBech32(from)

	keyring, err := NewKeyringInstance()
	key, err := keyring.KeyByAddress(fromAddress)

	if err != nil {
		fmt.Printf("error getting key %v\n", err.Error())
		return nil, errors.New("")
	}

	fmt.Printf("key found %v %v\n", key, key.GetName())

	clientCtx = clientCtx.
		WithKeyring(keyring).
		WithFromAddress(fromAddress).
		WithSkipConfirmation(true).
		WithFromName(key.GetName()).
		WithBroadcastMode("sync")

	return &clientCtx, nil
}

// TODO: change test keyring with other (file?) - new ticket for this
func NewKeyringInstance() (keyring.Keyring, error) {
	kr, err := keyring.New("baseledger", "test", "~/.baseledger", nil)

	if err != nil {
		fmt.Printf("error fetching test keyring %v\n", err.Error())
		return nil, errors.New("error fetching key ring")
	}

	return kr, nil
}

func SignTxAndGetTxBytes(clientCtx client.Context, msg sdk.Msg) ([]byte, error) {
	accNum, accSeq, err := clientCtx.AccountRetriever.GetAccountNumberSequence(clientCtx, clientCtx.FromAddress)

	if err != nil {
		fmt.Printf("error while retrieving acc %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}
	fmt.Printf("retrieved account %v %v\n", accNum, accSeq)
	txFactory := tx.Factory{}.
		WithChainID(clientCtx.ChainID).
		WithGas(100000).
		WithTxConfig(clientCtx.TxConfig).
		WithAccountNumber(accNum).
		WithSequence(accSeq).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring)

	txFactory, err = tx.PrepareFactory(clientCtx, txFactory)
	if err != nil {
		fmt.Printf("prepare factory error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	transaction, err := tx.BuildUnsignedTx(txFactory, msg)
	if err != nil {
		fmt.Printf("build unsigned tx error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	err = tx.Sign(txFactory, clientCtx.GetFromName(), transaction, true)
	if err != nil {
		fmt.Printf("sign tx error %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(transaction.GetTx())
	if err != nil {
		fmt.Printf("tx encoder %v\n", err.Error())
		return nil, errors.New("sign tx error")
	}

	return txBytes, nil
}
