package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/cron"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/httpd/handler"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/messaging"
	"github.com/unibrightio/proxy-api/types"
)

func main() {
	setupViper()
	logger.SetupLogger()
	setupDb()
	cron.StartCron()
	subscribeToWorkgroupMessages()

	r := gin.Default()
	r.GET("/trustmeshes", handler.GetTrustmeshesHandler())
	r.POST("/suggestion", handler.CreateInitialSuggestionRequestHandler())
	r.POST("/feedback", handler.CreateSynchronizationFeedbackHandler())
	r.POST("send_offchain_message", handler.SendOffchainMessageHandler())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// discuss if we should use config struct or this is enough
func setupViper() {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper read config error")
	}
}

// migrate should be separate package, and we should have .sh script for running, see provide services
// leaving this for first version but we should separate definetely
func setupDb() {
	dbutil.InitDbIfNotExists()
	dbutil.PerformMigrations()
	dbutil.InitConnection()
}

func subscribeToWorkgroupMessages() {
	natsServerUrl, _ := viper.Get("NATS_URL").(string)
	natsToken := "testToken1" // TODO: Read from configuration
	logger.Infof("subscribeToWorkgroupMessages natsServerUrl %v", natsServerUrl)
	messagingClient := &messaging.NatsMessagingClient{}
	messagingClient.Subscribe(natsServerUrl, natsToken, "baseledger", receiveOffchainProcessMessage)
}

func receiveOffchainProcessMessage(sender string, natsMsg *nats.Msg) {
	// TODO: should we move this parsing to nats client and just get struct in this callback?
	var natsMessage types.NatsMessage
	err := json.Unmarshal(natsMsg.Data, &natsMessage)
	if err != nil {
		logger.Errorf("Error parsing nats message %v\n", err)
	}

	logger.Infof("message received %v\n", natsMessage)
	entryType := common.SuggestionReceivedTrustmeshEntryType
	if natsMessage.ProcessMessage.EntryType == common.FeedbackSentTrustmeshEntryType {
		entryType = common.FeedbackReceivedTrustmeshEntryType
	}

	trustmeshEntry := &types.TrustmeshEntry{
		TendermintTransactionId:              natsMessage.ProcessMessage.BaseledgerTransactionIdOfStoredProof,
		OffchainProcessMessageId:             natsMessage.ProcessMessage.Id,
		SenderOrgId:                          natsMessage.ProcessMessage.SenderId,
		ReceiverOrgId:                        natsMessage.ProcessMessage.ReceiverId,
		WorkgroupId:                          uuid.FromStringOrNil(natsMessage.ProcessMessage.Topic),
		WorkstepType:                         natsMessage.ProcessMessage.WorkstepType,
		BaseledgerTransactionType:            natsMessage.ProcessMessage.BaseledgerTransactionType,
		BaseledgerTransactionId:              natsMessage.ProcessMessage.BaseledgerTransactionIdOfStoredProof,
		ReferencedBaseledgerTransactionId:    natsMessage.ProcessMessage.ReferencedBaseledgerTransactionId,
		BusinessObjectType:                   natsMessage.ProcessMessage.BusinessObjectType,
		BaseledgerBusinessObjectId:           natsMessage.ProcessMessage.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: natsMessage.ProcessMessage.ReferencedBaseledgerBusinessObjectId,
		ReferencedProcessMessageId:           natsMessage.ProcessMessage.ReferencedOffchainProcessMessageId,
		TransactionHash:                      natsMessage.TxHash,
		EntryType:                            entryType,
	}

	if !trustmeshEntry.Create() {
		logger.Errorf("error when creating new trustmesh entry")
	}

}
