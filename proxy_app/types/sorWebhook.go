package types

import (
	"strings"

	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/dbutil"
	"github.com/unibrightio/proxy-api/logger"
)

type AuthType int8

const (
	None      AuthType = iota // Webhook does not use auth
	BasicAuth                 // Authenticate with basic auth
	XCSRF                     // Authenticate with X-CSRF token
)

type WebhookType int8

const (
	CreateObject WebhookType = iota // Create new business object
	UpdateObject                    // Update business object
)

type RequestParam struct {
	ParamName       string
	ParamValueField string
}

type SorWebhook struct {
	Id              uuid.UUID
	Url             string
	UrlParams       string // semicolon delimited list of param_name:param_value_field
	HttpMethod      string
	WebhookType     WebhookType
	AuthType        AuthType
	AuthUsername    string
	AuthPassword    string // hashed
	XCSRFUrl        string
	BodyContentType string
	Body            string // JSON, XML or other format defined in BodyContentType
	BodyParams      string // semicolon delimited list of param_name:param_value_field
}

func (t *SorWebhook) Create() bool {
	if dbutil.Db.GetConn().NewRecord(t) {
		result := dbutil.Db.GetConn().Create(&t)
		rowsAffected := result.RowsAffected
		errors := result.GetErrors()
		if len(errors) > 0 {
			logger.Errorf("errors while creating new entry %v\n", errors)
			return false
		}
		return rowsAffected > 0
	}

	return false
}

func (t *SorWebhook) Delete() bool {
	result := dbutil.Db.GetConn().Delete(&t)
	rowsAffected := result.RowsAffected
	errors := result.GetErrors()

	if len(errors) > 0 {
		logger.Errorf("errors while deleting entry %v\n", errors)
		return false
	}

	return rowsAffected > 0
}

func FetchWebhookByType(webhookType WebhookType) *SorWebhook {
	logger.Infof("Fetching SOR webhook of type %v..", webhookType)

	var webhook *SorWebhook
	dbError := dbutil.Db.GetConn().First(&webhook, "webhook_type = ?", webhookType).Error

	if dbError != nil {
		logger.Warnf("SOR webhook does not exist")
		return nil
	}

	return webhook
}

func ParseRequestParamsIntoString(requestParams []RequestParam) string {
	requestParamsDbFormat := ""

	for _, param := range requestParams {
		requestParamsDbFormat += param.ParamName + ":" + param.ParamValueField + ";"
	}

	return requestParamsDbFormat
}

func ParseStringIntoRequestParams(requestParamsDbFormat string) []RequestParam {
	var requestParams []RequestParam

	paramsStringRepresentation := strings.Split(requestParamsDbFormat, ";")

	for _, param := range paramsStringRepresentation {
		paramNameAndFieldValue := strings.Split(param, ":")
		if len(paramNameAndFieldValue) == 2 {
			requestParams = append(requestParams, RequestParam{paramNameAndFieldValue[0], paramNameAndFieldValue[1]})
		}
	}

	return requestParams
}
