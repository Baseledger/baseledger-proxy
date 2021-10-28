package eth

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"github.com/unibrightio/proxy-api/logger"

	proxyCommon "github.com/unibrightio/proxy-api/common"
	contracts "github.com/unibrightio/proxy-api/contracts"
)

var ethClient *ethclient.Client

func GetClient() *ethclient.Client {
	if ethClient == nil {
		client, err := ethclient.Dial(viper.GetString("ETHEREUM_API_URL"))

		if err != nil {
			logger.Errorf("Error connecting to infure %v", err.Error())
			return nil
		}
		logger.Infof("New eth client initialized")
		ethClient = client
	}

	logger.Infof("Using existing eth client")
	return ethClient
}

func AddNewProof(txId string, proof string) {
	instance, auth := getContractInstance()
	if instance == nil || auth == nil {
		logger.Error("Error getting contract instance")
		return
	}

	tx, err := instance.Add(auth, txId, proof)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Infof("eth tx sent: %s", tx.Hash().Hex())
}

func GetProof(txId string) {
	instance, auth := getContractInstance()
	if instance == nil || auth == nil {
		logger.Error("Error getting contract instance")
		return
	}

	result, err := instance.Get(nil, txId)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Infof("Proof for tx id %v is %v", txId, string(result[:]))
}

func getContractInstance() (*contracts.Contracts, *bind.TransactOpts) {
	client := GetClient()

	if client == nil {
		return nil, nil
	}

	privateKey, err := crypto.HexToECDSA(viper.GetString("ETHEREUM_PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, nil
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(proxyCommon.EthProofsSmartContractAddress)

	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	return instance, auth
}
