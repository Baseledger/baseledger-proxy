package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

type createSorWebhookRequest struct {
	Url             string               `json:"url"`               // url representing the sor endpoint to trigger. params in the url must be defined with the following syntax {{param_name}}
	UrlParams       []types.RequestParam `json:"url_params"`        // list of key values represting url parameter name (param_name) -> url parameter value field (representing trustmesh entry field to populate the value from i.e BaseledgerTransactionId)
	HttpMethod      string               `json:"http_method"`       // PUT or POST
	WebhookType     types.WebhookType    `json:"webhook_type"`      // 0 - Create object, 1 - Update object
	AuthType        types.AuthType       `json:"auth_type"`         // 0 - None, 1 - Basic auth
	AuthUsername    string               `json:"auth_username"`     // Mandatory if auth_type > 0
	AuthPassword    string               `json:"auth_password"`     // Mandatory if auth_type > 0
	XcsrfUrl        string               `json:"xcsrf_url"`         // If provided, used to fetch the token and place it in header of every request. Provided auth_type will also be applied.
	BodyContentType string               `json:"body_content_type"` // JSON, XML or other format. JSON supported for now
	Body            string               `json:"body"`              // body of the request. params in the body must be defined with the following syntax {{param_name}}. Special param {{business_object_json_payload}} can be used to inject sync tree information when types.WebhookType create object. Special param {{new_object_status}} can be used when types.WebhookType update object
	BodyParams      []types.RequestParam `json:"body_params"`       // list of key values represting body parameter name -> body parameter value field (representing trustmesh entry field to populate the value from i.e BaseledgerTransactionId). Special params do not have to be listed here.
}

type sorWebhookDetailsDto struct {
	Id              uuid.UUID            `json:"id"`
	Url             string               `json:"url"`
	UrlParams       []types.RequestParam `json:"url_params"`
	HttpMethod      string               `json:"http_method"`
	WebhookType     types.WebhookType    `json:"webhook_type"`
	AuthType        types.AuthType       `json:"auth_type"`
	AuthUsername    string               `json:"auth_username"`
	AuthPassword    string               `json:"auth_password"`
	BodyContentType string               `json:"body_content_type"`
	Body            string               `json:"body"`
	BodyParams      []types.RequestParam `json:"body_params"`
}

// @Security BasicAuth
// GetSorWebhook ... Get all sor  webhooks
// @Summary Get sor webhooks
// @Description get sor webhooks
// @Tags SOR Webhooks
// @Produce json
// @Accept json
// @Param id path string format "uuid" "id"
// @Success 200 {array} workgroupMemberDetailsDto
// @Router /sorwebhook [get]
func GetSorWebhooksHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sorWebhooks []types.SorWebhook

		dbutil.Db.GetConn().Find(&sorWebhooks)

		var sorWebhookDtos []sorWebhookDetailsDto

		for i := 0; i < len(sorWebhooks); i++ {
			sorWebhookDto := &sorWebhookDetailsDto{}
			sorWebhookDto.Id = sorWebhooks[i].Id
			sorWebhookDto.Url = sorWebhooks[i].Url
			sorWebhookDto.UrlParams = types.ParseStringIntoRequestParams(sorWebhooks[i].UrlParams)
			sorWebhookDto.HttpMethod = sorWebhooks[i].HttpMethod
			sorWebhookDto.WebhookType = sorWebhooks[i].WebhookType
			sorWebhookDto.AuthType = sorWebhooks[i].AuthType
			sorWebhookDto.AuthUsername = sorWebhooks[i].AuthUsername
			sorWebhookDto.AuthPassword = sorWebhooks[i].AuthPassword
			sorWebhookDto.BodyContentType = sorWebhooks[i].BodyContentType
			sorWebhookDto.Body = sorWebhooks[i].Body
			sorWebhookDto.BodyParams = types.ParseStringIntoRequestParams(sorWebhooks[i].BodyParams)
			sorWebhookDtos = append(sorWebhookDtos, *sorWebhookDto)
		}

		restutil.Render(sorWebhookDtos, 200, c)
	}
}

// @Security BasicAuth
// Create SOR Webhook ... Create SOR Webhook
// @Summary Create new SOR webhook based on parameters
// @Description Create new SOR webhook
// @Tags SOR Webhook
// @Accept json
// @Param sorWebhook body object true "Create SOR webhook"
// @Success 200 {string} types.SorWebhook
// @Failure 400,422,500 {string} errorMessage
// @Router /sorwebhook [post]
func CreateSorWebhookHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		req := &createSorWebhookRequest{}
		err = json.Unmarshal(buf, &req)
		if err != nil {
			restutil.RenderError(err.Error(), 422, c)
			return
		}

		newSorWebhook := newSorWebhook(*req)

		if !newSorWebhook.Create() {
			logger.Errorf("error when creating new sor webhook")
			restutil.RenderError("error when creating new sor webhook", 500, c)
			return
		}

		restutil.Render(newSorWebhook.Id, 200, c)
	}
}

// @Security BasicAuth
// Delete SorWebhook Member... Delete SorWebhook Member
// @Summary Delete sorWebhook member
// @Description Delete sorWebhook member
// @Tags SorWebhook Members
// @Param id path string format "uuid" "id"
// @Success 204
// @Failure 404,500 {string} errorMessage
// @Router /sorwebhook/{id} [delete]
func DeleteSorWebhookHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		sorWebhookId := c.Param("id")

		var existingSorWebhook types.SorWebhook
		dbError := dbutil.Db.GetConn().First(&existingSorWebhook, "id = ?", sorWebhookId).Error

		if dbError != nil {
			logger.Errorf("error trying to fetch sor webhook with id %s\n", sorWebhookId)
			restutil.RenderError("sor webhook not found", 404, c)
			return
		}

		if !existingSorWebhook.Delete() {
			logger.Errorf("error when deleting sor webhook")
			restutil.RenderError("error when deleting sor webhook", 500, c)
			return
		}

		restutil.Render(nil, 204, c)
	}
}

func newSorWebhook(req createSorWebhookRequest) *types.SorWebhook {
	return &types.SorWebhook{
		Url:             req.Url,
		UrlParams:       types.ParseRequestParamsIntoString(req.UrlParams),
		HttpMethod:      req.HttpMethod,
		WebhookType:     req.WebhookType,
		AuthType:        req.AuthType,
		AuthUsername:    req.AuthUsername,
		AuthPassword:    req.AuthPassword,
		BodyContentType: req.BodyContentType,
		Body:            req.Body,
		BodyParams:      types.ParseRequestParamsIntoString(req.BodyParams),
	}
}
