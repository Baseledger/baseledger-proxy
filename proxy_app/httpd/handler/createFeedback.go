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
	"github.com/unibrightio/proxy-api/types"
)

type sendFeedbackDto struct {
	WorkflowId                 string `json:"workflow_id"`
	BaseledgerBusinessObjectId string `json:"baseledger_business_object_id"`
	Approved                   bool   `json:"approved"`
	FeedbackMessage            string `json:"feedback_message"`
}

type sendFeedbackResponseDto struct {
	WorkflowId                 string `json:"workflow_id"`
	WorkstepId                 string `json:"workstep_id"`
	BaseledgerBusinessObjectId string `json:"baseledger_business_object_id"`
	TransactionHash            string `json:"transaction_hash"`
	Error                      string `json:"error"`
}

// @Security BasicAuth
// Create Feedback ... Create Feedback
// @Summary Create new feedback based on parameters
// @Description Create new feedback
// @Tags Feedbacks
// @Accept json
// @Param feedback body sendFeedbackDto true "Feedback Request"
// @Success 200 {string} txHash
// @Failure 400,422,500 {string} errorMessage
// @Router /feedback [post]
func CreateSynchronizationFeedbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseDto := &sendFeedbackResponseDto{}

		buf, err := c.GetRawData()
		if err != nil {
			responseDto.Error = err.Error()
			restutil.Render(responseDto, 400, c)
			return
		}

		dto := &sendFeedbackDto{}
		err = json.Unmarshal(buf, &dto)
		if err != nil {
			responseDto.Error = err.Error()
			restutil.Render(responseDto, 422, c)
			return
		}

		latestTrustmeshEntry := &types.TrustmeshEntry{}

		if dto.BaseledgerBusinessObjectId != "" {
			latestTrustmeshEntry, err = types.GetLatestTrustmeshEntryBasedOnBboid(dto.BaseledgerBusinessObjectId)

			if err != nil {
				responseDto.Error = err.Error()
				restutil.Render(responseDto, 400, c)
				return
			}

			if latestTrustmeshEntry.EntryType != common.SuggestionReceivedTrustmeshEntryType {
				responseDto.Error = "Previous trustmesh entry is not of type Suggest"
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

			if latestTrustmeshEntry.EntryType != common.SuggestionReceivedTrustmeshEntryType {
				responseDto.Error = "Previous trustmesh entry is not of type Suggest"
				restutil.Render(responseDto, 400, c)
				return
			}
		} else {
			responseDto.Error = "Both bboid and workflow id are missing. At least one must be provided"
			restutil.Render(responseDto, 400, c)
			return
		}

		suggestionReceivedOffchainMessage, err := types.GetOffchainMsgById(latestTrustmeshEntry.OffchainProcessMessageId)

		if err != nil {
			responseDto.Error = err.Error()
			restutil.Render(responseDto, 400, c)
			return
		}

		newFeedbackRequest := newFeedbackRequest(*dto, *latestTrustmeshEntry, suggestionReceivedOffchainMessage)

		transactionId := uuid.NewV4()

		feedbackOffchainMessage := createFeedbackOffchainMessage(*newFeedbackRequest, suggestionReceivedOffchainMessage, transactionId)

		if !feedbackOffchainMessage.Create() {
			responseDto.Error = "error when creating new offchain msg entry"
			logger.Errorf(responseDto.Error)
			restutil.Render(responseDto, 500, c)
			return
		}

		payload := proxyutil.CreateNewFeedbackBaseledgerTransactionPayload(newFeedbackRequest, &feedbackOffchainMessage)

		signAndBroadcastPayload := restutil.SignAndBroadcastPayload{
			TransactionId: transactionId.String(),
			Payload:       payload,
			OpCode:        uint32(getRandomFeedbackOpCode()),
		}

		transactionHash := restutil.SignAndBroadcast(signAndBroadcastPayload)

		if transactionHash == nil {
			responseDto.Error = "sign and broadcast transaction error"
			restutil.Render(responseDto, 500, c)
			return
		}

		feedbackSentTrustmeshEntry := createFeedbackSentTrustmeshEntry(*newFeedbackRequest, feedbackOffchainMessage, *transactionHash)

		if !feedbackSentTrustmeshEntry.Create() {
			responseDto.Error = "error when creating new trustmesh entry"
			logger.Errorf(responseDto.Error)
			restutil.Render(responseDto, 500, c)
			return
		}

		responseDto.WorkflowId = feedbackSentTrustmeshEntry.TrustmeshId.String()
		responseDto.WorkstepId = feedbackSentTrustmeshEntry.Id.String()
		responseDto.BaseledgerBusinessObjectId = feedbackSentTrustmeshEntry.ReferencedBaseledgerBusinessObjectId
		responseDto.TransactionHash = feedbackSentTrustmeshEntry.TransactionHash

		restutil.Render(responseDto, 200, c)
	}
}

func getRandomFeedbackOpCode() int {
	rand.Seed(time.Now().UnixNano())
	min := 7
	max := 8

	return rand.Intn(max-min+1) + min
}

func createFeedbackOffchainMessage(
	newFeedbackRequest types.NewFeedbackRequest,
	suggestionReceivedOffchainMessage *types.OffchainProcessMessage,
	transactionId uuid.UUID,
) types.OffchainProcessMessage {
	baseledgerTransactionType := common.BaseledgerTransactionTypeApprove

	if !newFeedbackRequest.Approved {
		baseledgerTransactionType = common.BaseledgerTransactionTypeReject
	}

	offchainMessage := types.OffchainProcessMessage{
		SenderId:                             uuid.FromStringOrNil(viper.Get("ORGANIZATION_ID").(string)),
		ReceiverId:                           suggestionReceivedOffchainMessage.SenderId,
		Topic:                                suggestionReceivedOffchainMessage.Topic,
		WorkstepType:                         common.WorkstepTypeFeedback,
		BaseledgerSyncTreeJson:               suggestionReceivedOffchainMessage.BaseledgerSyncTreeJson,
		BusinessObjectProof:                  suggestionReceivedOffchainMessage.BusinessObjectProof,
		BaseledgerBusinessObjectId:           "", // empty because we are giving feedback
		ReferencedBaseledgerBusinessObjectId: suggestionReceivedOffchainMessage.BaseledgerBusinessObjectId,
		StatusTextMessage:                    newFeedbackRequest.FeedbackMessage,
		BaseledgerTransactionIdOfStoredProof: transactionId,
		TendermintTransactionIdOfStoredProof: transactionId,
		BusinessObjectType:                   suggestionReceivedOffchainMessage.BusinessObjectType,
		BaseledgerTransactionType:            baseledgerTransactionType,
		ReferencedBaseledgerTransactionId:    suggestionReceivedOffchainMessage.BaseledgerTransactionIdOfStoredProof,
		EntryType:                            common.FeedbackSentTrustmeshEntryType,
		SorBusinessObjectId:                  suggestionReceivedOffchainMessage.SorBusinessObjectId,
		ReferencedWorkstepType:               suggestionReceivedOffchainMessage.WorkstepType,
	}

	return offchainMessage
}

func createFeedbackSentTrustmeshEntry(newFeedbackRequest types.NewFeedbackRequest, offchainMsg types.OffchainProcessMessage, txHash string) *types.TrustmeshEntry {
	trustmeshEntry := &types.TrustmeshEntry{
		TendermintTransactionId:              offchainMsg.BaseledgerTransactionIdOfStoredProof,
		OffchainProcessMessageId:             offchainMsg.Id,
		SenderOrgId:                          uuid.FromStringOrNil(viper.Get("ORGANIZATION_ID").(string)),
		ReceiverOrgId:                        uuid.FromStringOrNil(offchainMsg.ReceiverId.String()),
		WorkgroupId:                          uuid.FromStringOrNil(offchainMsg.Topic),
		WorkstepType:                         offchainMsg.WorkstepType,
		BaseledgerTransactionType:            offchainMsg.BaseledgerTransactionType,
		BaseledgerTransactionId:              offchainMsg.BaseledgerTransactionIdOfStoredProof,
		ReferencedBaseledgerTransactionId:    uuid.FromStringOrNil(newFeedbackRequest.OriginalBaseledgerTransactionId),
		BusinessObjectType:                   offchainMsg.BusinessObjectType,
		BaseledgerBusinessObjectId:           offchainMsg.BaseledgerBusinessObjectId,
		ReferencedBaseledgerBusinessObjectId: offchainMsg.ReferencedBaseledgerBusinessObjectId,
		TransactionHash:                      txHash,
		EntryType:                            common.FeedbackSentTrustmeshEntryType,
		SorBusinessObjectId:                  offchainMsg.SorBusinessObjectId,
	}

	return trustmeshEntry
}

func newFeedbackRequest(dto sendFeedbackDto, suggestionReceivedTrustmeshEntry types.TrustmeshEntry, suggestionReceivedOffchainMessage *types.OffchainProcessMessage) *types.NewFeedbackRequest {
	return &types.NewFeedbackRequest{
		WorkgroupId:        suggestionReceivedTrustmeshEntry.WorkgroupId,
		Recipient:          suggestionReceivedTrustmeshEntry.SenderOrgId.String(),
		BusinessObjectType: suggestionReceivedTrustmeshEntry.BusinessObjectType,
		BaseledgerBusinessObjectIdOfApprovedObject: suggestionReceivedTrustmeshEntry.BaseledgerBusinessObjectId,
		OriginalBaseledgerTransactionId:            suggestionReceivedTrustmeshEntry.BaseledgerTransactionId.String(), // TODO: BAS-79 is this correct?
		Approved:                                   dto.Approved,
		FeedbackMessage:                            dto.FeedbackMessage,
		BaseledgerProvenBusinessObjectJson:         suggestionReceivedOffchainMessage.BaseledgerSyncTreeJson,
		HashOfObjectToApprove:                      suggestionReceivedOffchainMessage.BusinessObjectProof,
		OriginalOffchainProcessMessageId:           suggestionReceivedOffchainMessage.Id.String(), // TODO: BAS-79 is this correct?
	}
}
