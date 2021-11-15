package handler

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/spf13/viper"
	"github.com/unibrightio/proxy-api/common"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/proxyutil"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/synctree"
	"github.com/unibrightio/proxy-api/types"
	"github.com/unibrightio/proxy-api/workgroups"
)

type sendSuggestionDto struct {
	WorkgroupId                string   `json:"workgroup_id"`
	Recipient                  string   `json:"recipient"`
	WorkstepType               string   `json:"workstep_type"`
	WorkflowId                 string   `json:"workflow_id"`
	BaseledgerBusinessObjectId string   `json:"baseledger_business_object_id"`
	BusinessObjectType         string   `json:"business_object_type"`
	BusinessObjectId           string   `json:"business_object_id"`
	BusinessObjectJson         string   `json:"business_object_json"`
	KnowledgeLimiters          []string `json:"knowledge_limiters"`
}

type sendSuggestionResponseDto struct {
	WorkflowId                 string `json:"workflow_id"`
	WorkstepId                 string `json:"workstep_id"`
	BaseledgerBusinessObjectId string `json:"baseledger_business_object_id"`
	TransactionHash            string `json:"transaction_hash"`
	Error                      string `json:"error"`
}

// @Security BasicAuth
// Create Suggestion ... Create Suggestion
// @Summary Create new suggestion based on parameters
// @Description Create new suggestion
// @Tags Suggestions
// @Accept json
// @Param suggestion body sendSuggestionDto true "Suggestion Request"
// @Success 200 {string} txHash
// @Failure 400,422,500 {string} errorMessage
// @Router /suggestion [post]
func CreateSuggestionRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseDto := &sendSuggestionResponseDto{}

		buf, err := c.GetRawData()
		if err != nil {
			responseDto.Error = err.Error()
			restutil.Render(responseDto, 400, c)
			return
		}

		dto := &sendSuggestionDto{}
		err = json.Unmarshal(buf, &dto)
		if err != nil {
			responseDto.Error = err.Error()
			restutil.Render(responseDto, 422, c)
			return
		}

		if dto.WorkgroupId == "" {
			// valid only until we go with the assumtion 1 recipient == 1 workgroup
			workgroupClient := &workgroups.PostgresWorkgroupClient{}
			workgroupMembership := workgroupClient.GetRecipientWorkgroupMember(dto.Recipient)

			if workgroupMembership == nil {
				responseDto.Error = "failed to find a workgroup membership for recipient"
				restutil.Render(responseDto, 400, c)
				return
			}

			dto.WorkgroupId = workgroupMembership.WorkgroupId
		}

		newSuggestionRequest := &types.NewSuggestionRequest{}

		// there is no bboid and no trustmesh id that the suggestion references - we treat it as INITIAL
		if dto.BaseledgerBusinessObjectId == "" && dto.WorkflowId == "" {
			newSuggestionRequest = createNewInitialSuggestionRequest(*dto)
		} else {
			// either go with bboid or workflow id
			latestTrustmeshEntry := &types.TrustmeshEntry{}
			if dto.BaseledgerBusinessObjectId != "" {
				latestTrustmeshEntry, err = types.GetLatestTrustmeshEntryBasedOnBboid(dto.BaseledgerBusinessObjectId)

				if err != nil {
					responseDto.Error = err.Error()
					restutil.Render(responseDto, 400, c)
					return
				}
			} else if dto.WorkflowId != "" {
				latestTrustmeshEntry, err = types.GetLatestTrustmeshEntryBasedOnTrustmeshId(dto.WorkflowId)

				if err != nil {
					responseDto.Error = err.Error()
					restutil.Render(responseDto, 400, c)
					return
				}
			}

			// consider opening up for other previous worksteps.
			// currently we are rigid that feedback has to be the step that came before this one
			// in the scenario where workflow id is provided
			if latestTrustmeshEntry.EntryType != common.FeedbackReceivedTrustmeshEntryType && latestTrustmeshEntry.EntryType != common.FeedbackSentTrustmeshEntryType {
				responseDto.Error = "Previous workstep is not feedback sent/received"
				restutil.Render(responseDto, 400, c)
				return
			}

			if dto.WorkstepType == common.WorkstepTypeNewVersion {
				newSuggestionRequest = createNewVersionSuggestionRequestFromLatestTrustmeshEntry(*dto, *latestTrustmeshEntry)
			} else if dto.WorkstepType == common.WorkstepTypeNextWorkstep {
				newSuggestionRequest = createNextWorkstepOrFinalSuggestionRequestFromLatestTrustmeshEntry(*dto, *latestTrustmeshEntry, common.WorkstepTypeNextWorkstep)
			} else if dto.WorkstepType == common.WorkstepTypeFinal {
				newSuggestionRequest = createNextWorkstepOrFinalSuggestionRequestFromLatestTrustmeshEntry(*dto, *latestTrustmeshEntry, common.WorkstepTypeFinal)
			} else {
				responseDto.Error = "Workstep type is invalid for suggestion"
				restutil.Render(responseDto, 400, c)
				return
			}
		}

		syncTree := synctree.CreateFromBusinessObjectJson(newSuggestionRequest.BusinessObjectJson, newSuggestionRequest.KnowledgeLimiters)
		logger.Infof("Sync tree %v", syncTree)

		syncTreeJson, err := json.Marshal(syncTree)
		if err != nil {
			responseDto.Error = "error marshaling sync tree"
			logger.Errorf(responseDto.Error, err.Error())
			restutil.Render(responseDto, 500, c)
			return
		}

		logger.Infof("Sync tree json %v", string(syncTreeJson))

		transactionId := uuid.NewV4()

		offchainMsg := createNewSuggestionOffchainMessage(*newSuggestionRequest, transactionId, string(syncTreeJson), syncTree.RootProof)

		if !offchainMsg.Create() {
			responseDto.Error = "error when creating new offchain msg entry"
			logger.Errorf(responseDto.Error)
			restutil.Render(responseDto, 500, c)
			return
		}

		payload := proxyutil.CreateNewSuggestionBaseledgerTransactionPayload(newSuggestionRequest, &offchainMsg)

		signAndBroadcastPayload := restutil.SignAndBroadcastPayload{
			TransactionId: transactionId.String(),
			Payload:       payload,
			OpCode:        uint32(getRandomSuggestionOpCode()),
		}

		transactionHash := restutil.SignAndBroadcast(signAndBroadcastPayload)

		if transactionHash == nil {
			responseDto.Error = "sign and broadcast transaction error"
			logger.Errorf(responseDto.Error)
			restutil.Render(responseDto, 500, c)
			return
		}

		trustmeshEntry := createSuggestionSentTrustmeshEntry(*newSuggestionRequest, transactionId, offchainMsg, *transactionHash)

		if !trustmeshEntry.Create() {
			responseDto.Error = "error when creating new trustmesh entry"
			logger.Errorf(responseDto.Error)
			restutil.Render(responseDto, 500, c)
			return
		}

		// quick fix to get trustmesh id, consider other options
		createdTrustmesh, _ := types.GetTrustmeshEntryById(trustmeshEntry.Id)

		responseDto.WorkflowId = createdTrustmesh.TrustmeshId.String()
		responseDto.WorkstepId = createdTrustmesh.Id.String()
		responseDto.BaseledgerBusinessObjectId = createdTrustmesh.BaseledgerBusinessObjectId
		responseDto.TransactionHash = createdTrustmesh.TransactionHash

		restutil.Render(responseDto, 200, c)
	}
}

func getRandomSuggestionOpCode() int {
	rand.Seed(time.Now().UnixNano())
	min := 7
	max := 8

	return rand.Intn(max-min+1) + min
}

func createNewInitialSuggestionRequest(req sendSuggestionDto) *types.NewSuggestionRequest {
	return &types.NewSuggestionRequest{
		WorkgroupId:                uuid.FromStringOrNil(req.WorkgroupId),
		Recipient:                  req.Recipient,
		WorkstepType:               common.WorkstepTypeInitial,
		BusinessObjectType:         req.BusinessObjectType,
		BusinessObjectId:           req.BusinessObjectId,
		BaseledgerBusinessObjectId: uuid.NewV4().String(),
		BusinessObjectJson:         req.BusinessObjectJson,
		KnowledgeLimiters:          req.KnowledgeLimiters,
	}
}

func createNewVersionSuggestionRequestFromDto(req sendSuggestionDto) *types.NewSuggestionRequest {
	return &types.NewSuggestionRequest{
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		Recipient:                            req.Recipient,
		WorkstepType:                         common.WorkstepTypeNewVersion,
		BusinessObjectType:                   req.BusinessObjectType,
		BusinessObjectId:                     req.BusinessObjectId,
		BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: req.BaseledgerBusinessObjectId,
		BusinessObjectJson:                   req.BusinessObjectJson,
		KnowledgeLimiters:                    req.KnowledgeLimiters,
	}
}

func createNextWorkstepOrFinalSuggestionRequestFromDto(req sendSuggestionDto, workstepType string) *types.NewSuggestionRequest {
	return &types.NewSuggestionRequest{
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		Recipient:                            req.Recipient,
		WorkstepType:                         workstepType,
		BusinessObjectType:                   req.BusinessObjectType,
		BusinessObjectId:                     req.BusinessObjectId,
		BaseledgerBusinessObjectId:           uuid.NewV4().String(),
		ReferencedBaseledgerBusinessObjectId: req.BaseledgerBusinessObjectId,
		BusinessObjectJson:                   req.BusinessObjectJson,
		KnowledgeLimiters:                    req.KnowledgeLimiters,
	}
}

func createNewVersionSuggestionRequestFromLatestTrustmeshEntry(
	req sendSuggestionDto,
	latestFeedbackTrustmeshEntry types.TrustmeshEntry) *types.NewSuggestionRequest {
	return &types.NewSuggestionRequest{
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		Recipient:                            latestFeedbackTrustmeshEntry.SenderOrgId.String(),
		WorkstepType:                         common.WorkstepTypeNewVersion,
		BusinessObjectType:                   latestFeedbackTrustmeshEntry.BusinessObjectType,
		BusinessObjectId:                     latestFeedbackTrustmeshEntry.SorBusinessObjectId,
		BaseledgerBusinessObjectId:           latestFeedbackTrustmeshEntry.ReferencedBaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: latestFeedbackTrustmeshEntry.ReferencedBaseledgerBusinessObjectId,
		ReferencedBaseledgerTransactionId:    latestFeedbackTrustmeshEntry.BaseledgerTransactionId.String(),
		BusinessObjectJson:                   req.BusinessObjectJson,
		KnowledgeLimiters:                    req.KnowledgeLimiters,
	}
}

func createNextWorkstepOrFinalSuggestionRequestFromLatestTrustmeshEntry(
	req sendSuggestionDto,
	latestFeedbackTrustmeshEntry types.TrustmeshEntry,
	workstepType string) *types.NewSuggestionRequest {
	return &types.NewSuggestionRequest{
		WorkgroupId:                          uuid.FromStringOrNil(req.WorkgroupId),
		Recipient:                            latestFeedbackTrustmeshEntry.SenderOrgId.String(),
		WorkstepType:                         workstepType,
		BusinessObjectType:                   req.BusinessObjectType,
		BusinessObjectId:                     req.BusinessObjectId,
		BaseledgerBusinessObjectId:           uuid.NewV4().String(),
		ReferencedBaseledgerBusinessObjectId: latestFeedbackTrustmeshEntry.ReferencedBaseledgerBusinessObjectId,
		ReferencedBaseledgerTransactionId:    latestFeedbackTrustmeshEntry.BaseledgerTransactionId.String(),
		BusinessObjectJson:                   req.BusinessObjectJson,
		KnowledgeLimiters:                    req.KnowledgeLimiters,
	}
}

func createNewSuggestionOffchainMessage(
	req types.NewSuggestionRequest, transactionId uuid.UUID, syncTreeJson string, rootProof string) types.OffchainProcessMessage {
	offchainMessage := types.OffchainProcessMessage{
		SenderId:                             uuid.FromStringOrNil(viper.Get("ORGANIZATION_ID").(string)),
		ReceiverId:                           uuid.FromStringOrNil(req.Recipient),
		Topic:                                req.WorkgroupId.String(), // TODO: BAS-79 why is this called topic? rename to workgroup id
		WorkstepType:                         req.WorkstepType,
		BaseledgerSyncTreeJson:               syncTreeJson,
		BusinessObjectProof:                  rootProof,
		BusinessObjectType:                   req.BusinessObjectType,
		SorBusinessObjectId:                  req.BusinessObjectId,
		BaseledgerBusinessObjectId:           req.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: req.ReferencedBaseledgerBusinessObjectId,
		BaseledgerTransactionType:            common.BaseledgerTransactionTypeSuggest,
		EntryType:                            common.SuggestionSentTrustmeshEntryType, // can we ditch this as we have one above that says that this is a suggestion
		StatusTextMessage:                    req.WorkstepType + " " + common.BaseledgerTransactionTypeSuggest,
		BaseledgerTransactionIdOfStoredProof: transactionId,
		ReferencedBaseledgerTransactionId:    uuid.FromStringOrNil(req.ReferencedBaseledgerTransactionId),
		TendermintTransactionIdOfStoredProof: transactionId,
	}

	return offchainMessage
}

func createSuggestionSentTrustmeshEntry(req types.NewSuggestionRequest, transactionId uuid.UUID, offchainMsg types.OffchainProcessMessage, txHash string) *types.TrustmeshEntry {
	return &types.TrustmeshEntry{
		EntryType:                         common.SuggestionSentTrustmeshEntryType,
		SenderOrgId:                       offchainMsg.SenderId,
		ReceiverOrgId:                     uuid.FromStringOrNil(req.Recipient),
		WorkgroupId:                       req.WorkgroupId,
		WorkstepType:                      offchainMsg.WorkstepType,
		BaseledgerTransactionType:         offchainMsg.BaseledgerTransactionType,
		BusinessObjectType:                req.BusinessObjectType,
		SorBusinessObjectId:               req.BusinessObjectId,
		BaseledgerBusinessObjectId:        offchainMsg.BaseledgerBusinessObjectId,
		OffchainProcessMessageId:          offchainMsg.Id,
		TendermintTransactionId:           transactionId,
		TransactionHash:                   txHash,
		BaseledgerTransactionId:           transactionId,
		ReferencedBaseledgerTransactionId: uuid.FromStringOrNil(req.ReferencedBaseledgerTransactionId),
	}
}
