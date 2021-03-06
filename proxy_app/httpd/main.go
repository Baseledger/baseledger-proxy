package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/cron"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/httpd/handler"
	proxyMiddleware "github.com/unibrightio/proxy-api/httpd/middleware"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/messaging"

	"github.com/unibrightio/proxy-api/types"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	docs "github.com/unibrightio/proxy-api/httpd/docs"

	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

// @title Baseledger Proxy API documentation
// @version 1.0.0
// @host localhost:8081
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	setupViper()
	docs.SwaggerInfo.Host = viper.GetString("SWAGGER_HOST")
	logger.SetupLogger()
	setupDb()
	cron.StartCron()
	subscribeToWorkgroupMessages()

	rate := limiter.Rate{
		Period: 1 * time.Hour * 24,
		Limit:  10,
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)
	rateMiddleware := mgin.NewMiddleware(instance)

	r := gin.Default()
	r.Use(proxyMiddleware.CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/trustmeshes", proxyMiddleware.BasicAuth(false), handler.GetTrustmeshesHandler())
	r.GET("/trustmeshes/:id", proxyMiddleware.BasicAuth(false), handler.GetTrustmeshHandler())
	r.POST("/suggestion", proxyMiddleware.BasicAuth(true), proxyMiddleware.AuthorizeJWTMiddleware(true), handler.CreateSuggestionRequestHandler())
	r.POST("/feedback", proxyMiddleware.BasicAuth(true), proxyMiddleware.AuthorizeJWTMiddleware(true), handler.CreateSynchronizationFeedbackHandler())
	r.GET("/sunburst/:txId", proxyMiddleware.BasicAuth(false), handler.GetSunburstHandler())
	r.GET("/organization", proxyMiddleware.BasicAuth(false), handler.GetOrganizationsHandler())
	r.POST("/organization", proxyMiddleware.BasicAuth(false), handler.CreateOrganizationHandler())
	r.DELETE("/organization/:id", proxyMiddleware.BasicAuth(false), handler.DeleteOrganizationHandler())
	r.GET("/workgroup", proxyMiddleware.BasicAuth(true), proxyMiddleware.AuthorizeJWTMiddleware(true), handler.GetWorkgroupsHandler())
	r.POST("/workgroup", proxyMiddleware.BasicAuth(true), proxyMiddleware.AuthorizeJWTMiddleware(true), handler.CreateWorkgroupHandler())
	r.DELETE("/workgroup/:id", proxyMiddleware.BasicAuth(true), proxyMiddleware.AuthorizeJWTMiddleware(true), handler.DeleteWorkgroupHandler())
	r.GET("/workgroup/:id/participation", proxyMiddleware.BasicAuth(false), handler.GetWorkgroupMembersHandler())
	r.POST("/workgroup/:id/participation", proxyMiddleware.BasicAuth(true), proxyMiddleware.AuthorizeJWTMiddleware(true), handler.CreateWorkgroupMemberHandler())
	r.DELETE("/workgroup/:id/participation/:participationId", proxyMiddleware.BasicAuth(false), handler.DeleteWorkgroupMemberHandler())
	r.GET("/sorwebhook", proxyMiddleware.BasicAuth(false), handler.GetSorWebhooksHandler())
	r.POST("/sorwebhook", proxyMiddleware.BasicAuth(false), handler.CreateSorWebhookHandler())
	r.DELETE("/sorwebhook/:id", proxyMiddleware.BasicAuth(false), handler.DeleteSorWebhookHandler())
	// TODO: BAS-29 r.POST("/workgroup/invite", handler.InviteToWorkgroupHandler())
	// full details of workgroup, including organization
	r.GET("/workflow/new/:workgroup_id", proxyMiddleware.AuthorizeJWTMiddleware(false), handler.GetNewWorkflowHandler())
	r.GET("/workflow/latestState/:bo_id", proxyMiddleware.AuthorizeJWTMiddleware(false), handler.GetLatestWorkflowStateHandler())
	r.POST("/dev/users", handler.CreateUserHandler())
	r.POST("/dev/auth", handler.LoginUserHandler())
	r.POST("/dev/tx", proxyMiddleware.AuthorizeJWTMiddleware(false), rateMiddleware, handler.CreateTransactionHandler())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// discuss if we should use config struct or this is enough
func setupViper() {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // Overwrite config with env variables if exist, important for debugging session
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Printf("viper read config error %v\n", err))
	}
}

// migrate should be separate package, and we should have .sh script for running, see provide services
// leaving this for first version but we should separate definetely
func setupDb() {
	dbutil.InitDbIfNotExists()
	dbutil.PerformMigrations()
	// TODO: BAS-29 Add own org id to database with some dummy name
	dbutil.InitConnection()
}

func subscribeToWorkgroupMessages() {
	natsServerUrl, _ := viper.Get("NATS_URL").(string)
	natsToken := "testToken1" // TODO: Read from configuration
	logger.Infof("subscribeToWorkgroupMessages natsServerUrl %v", natsServerUrl)
	messagingClient := &messaging.NatsMessagingClient{}
	messagingClient.Subscribe(natsServerUrl, natsToken, common.BaseledgerNatsSubject, receiveOffchainProcessMessage)
	messagingClient.Subscribe(natsServerUrl, natsToken, common.EthTxHashNatsSubject, receiveTxEthHashUpdateMessage)
}

func receiveOffchainProcessMessage(sender string, natsMsg *nats.Msg) {
	// TODO: should we move this parsing to nats client and just get struct in this callback?
	var natsMessage types.NatsMessage
	err := json.Unmarshal(natsMsg.Data, &natsMessage)
	if err != nil {
		logger.Errorf("Error parsing nats message %v\n", err)
		return
	}

	logger.Infof("message received %v\n", natsMessage)

	natsMessage.ProcessMessage.Id = uuid.Nil // set to nil so that it can be created in the DB
	if !natsMessage.ProcessMessage.Create() {
		logger.Errorf("error when creating new offchain msg entry")
		return
	}

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
		SorBusinessObjectId:                  natsMessage.ProcessMessage.SorBusinessObjectId,
		TransactionHash:                      natsMessage.TxHash,
		EntryType:                            entryType,
	}

	if !trustmeshEntry.Create() {
		logger.Errorf("error when creating new trustmesh entry")
	}

}

func receiveTxEthHashUpdateMessage(sender string, natsMsg *nats.Msg) {
	var natsTrustmeshUpdateMessage types.NatsTrustmeshUpdateMessage
	err := json.Unmarshal(natsMsg.Data, &natsTrustmeshUpdateMessage)
	if err != nil {
		logger.Errorf("Error parsing nats message %v\n", err)
		return
	}

	logger.Infof("message received %v", natsTrustmeshUpdateMessage)

	trustmeshEntry, err := types.GetLatestTrustmeshEntryBasedOnBboid(natsTrustmeshUpdateMessage.BaseledgerBusinessObjectId)
	if err != nil {
		logger.Errorf("Error getting latest trustmesh entry by bbod %v", err.Error())
	}
	types.UpdateTrustmeshEthTxHash(trustmeshEntry.TrustmeshId, natsTrustmeshUpdateMessage.EthExitTxHash)
	return
}
