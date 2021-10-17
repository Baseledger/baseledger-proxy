package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unibrightio/proxy-api/logger"
)

var ethClient *ethclient.Client

func GetClient() *ethclient.Client {
	if ethClient == nil {
		client, err := ethclient.Dial("https://ropsten.infura.io/v3/0f43a95e908e4114a72f4f9e7e3913a7")

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
