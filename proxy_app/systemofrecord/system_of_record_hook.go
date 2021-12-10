package systemofrecord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"reflect"
	"strings"

	"github.com/oleiade/reflections"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/types"
)

var client http.Client

func InitClient() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	client = http.Client{
		Jar: jar,
	}
}

func TriggerSorWebhook(
	webhookType types.WebhookType,
	trustmeshEntry *types.TrustmeshEntry,
	payload string,
	approved string, // TODO This is 1 or 0 for feedback Approved or Rejected when origin is the counterparty, and is 1 or 0 for proxy status update Success or Failuer when origin is empty.
	message string,
	origin string,
) bool {

	webhook := types.FetchWebhookByType(webhookType)

	if webhook == nil {
		return false
	}

	targetUrl := buildWebhookRequestUrl(webhook.Url, webhook.UrlParams, trustmeshEntry)

	if targetUrl == "" {
		return false
	}

	if webhook.AuthType == types.BasicAuth {
		targetUrl = handleBasicAuth(targetUrl, webhook.AuthUsername, webhook.AuthPassword)
	}

	requestBody := buildWebhookRequestBody(
		webhook.Body,
		webhook.BodyParams,
		webhook.WebhookType,
		trustmeshEntry,
		payload,
		approved,
		message,
		origin)

	if requestBody == "" {
		return false
	}

	request := prepareWebhookRequest(requestBody, targetUrl, webhook.HttpMethod)

	if request == nil {
		return false
	}

	if webhook.XCSRFUrl != "" {
		addXcsrfTokenToRequest(webhook, request)
	}

	triggerWebhookRequest(request)
	return true
}

func buildWebhookRequestUrl(targetUrl string, urlParams string, trustmeshEntry *types.TrustmeshEntry) string {
	logger.Infof("Building hook url..")

	requestParams := types.ParseStringIntoRequestParams(urlParams)

	for _, param := range requestParams {
		field := reflect.Indirect(reflect.ValueOf(trustmeshEntry)).FieldByName(param.ParamValueField)
		if !field.IsValid() {
			logger.Errorf("Error fetching request parameter field %v from trustmesh entry \n", param.ParamValueField)
			return ""
		}
		targetUrl = strings.Replace(targetUrl, "{{"+param.ParamName+"}}", fmt.Sprintf("%v", field), 1)
	}

	return targetUrl
}

func handleBasicAuth(targetUrl string, username string, password string) string {
	targetUrlParts := strings.Split(targetUrl, "//")
	targetUrl = targetUrlParts[0] + "//" + username + ":" + password + "@" + targetUrlParts[1]
	logger.Infof("Included basic auth credentials into url")
	return targetUrl
}

func buildWebhookRequestBody(
	bodyTemplate string,
	bodyParams string,
	webhookType types.WebhookType,
	trustmeshEntry *types.TrustmeshEntry,
	payload string, // TODO: Move special params templating to a dedicated method
	approved string,
	message string,
	origin string) string {
	logger.Infof("Building body..")

	requestBody := bodyTemplate
	bodyRequestParams := types.ParseStringIntoRequestParams(bodyParams)

	for _, param := range bodyRequestParams {
		paramValue, err := reflections.GetField(trustmeshEntry, param.ParamValueField)
		if err != nil {
			logger.Errorf("Error %v fetching request parameter field %v from trustmesh entry \n", err.Error(), param.ParamValueField)
			return ""
		}
		requestBody = strings.Replace(requestBody, "{{"+param.ParamName+"}}", fmt.Sprintf("%v", paramValue), 1)
	}

	if webhookType == types.CreateObject {
		requestBody = strings.Replace(requestBody, "{{business_object_json_payload}}", jsonEscape(payload), 1)
	} else {
		requestBody = strings.Replace(requestBody, "{{approved}}", approved, 1)
		requestBody = strings.Replace(requestBody, "{{message}}", message, 1)
	}

	requestBody = strings.Replace(requestBody, "{{origin}}", origin, 1)

	logger.Infof("Body built %v\n", requestBody)

	return requestBody
}

func prepareWebhookRequest(requestBody string, targetUrl string, httpMethod string) *http.Request {
	logger.Infof("Preparing request with body.. %v", requestBody)

	var jsonStr = []byte(requestBody)

	req, err := http.NewRequest(httpMethod, targetUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		logger.Errorf("Error preparing request %v\n", err.Error())
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	return req
}

func addXcsrfTokenToRequest(webhook *types.SorWebhook, request *http.Request) {
	xcsrfUrl := webhook.XCSRFUrl

	if webhook.AuthType == types.BasicAuth {
		xcsrfUrl = handleBasicAuth(xcsrfUrl, webhook.AuthUsername, webhook.AuthPassword)
	}

	xcsrfRequest := prepareXcsrfTokenRequest(xcsrfUrl, "GET")
	token, cookies := triggerXcsrfTokenRequest(xcsrfRequest)

	if token == "" {
		return
	}

	request.Header.Add("X-CSRF-Token", token)

	for _, c := range cookies {
		request.AddCookie(c)
	}
}

func prepareXcsrfTokenRequest(targetUrl string, httpMethod string) *http.Request {
	logger.Infof("Preparing X-CSRF Token request..")

	req, err := http.NewRequest(httpMethod, targetUrl, nil)
	if err != nil {
		logger.Errorf("Error preparing request %v\n", err.Error())
		return nil
	}

	req.Header.Add("X-CSRF-Token", "Fetch")

	return req
}

func triggerXcsrfTokenRequest(request *http.Request) (string, []*http.Cookie) {
	logger.Infof("Firing away X-CSRF Token request..")

	resp, err := client.Do(request)
	if err != nil {
		logger.Errorf("Error firing away request %v\n", err.Error())
		return "", nil
	}

	if resp.StatusCode != http.StatusNoContent {
		logger.Errorf("X-CSRF Token request error, wrong status code %v\n", resp.StatusCode)
		return "", nil
	}

	if resp.Header.Get("X-CSRF-Token") == "" {
		logger.Errorf("X-CSRF Token response header missing")
		return "", nil
	}

	return resp.Header.Get("X-CSRF-Token"), resp.Cookies()
}

func triggerWebhookRequest(request *http.Request) {
	logger.Infof("Firing away request..")

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		logger.Errorf("DumpRequest error %s\n", err.Error())
	}

	logger.Infof("DumpRequest before result " + string(requestDump))

	resp, err := client.Do(request)
	if err != nil {
		logger.Errorf("Error firing away request %v\n", err.Error())
		return
	}

	// Save a copy of this request for debugging.
	requestDump, err = httputil.DumpRequest(request, true)
	if err != nil {
		logger.Errorf("DumpRequest error %s\n", err.Error())
	}

	logger.Infof("DumpRequest after result " + string(requestDump))

	logger.Infof("Sor Webhook request succesfull")
	logger.Infof("Sor Webhook request response %v\n", resp)
	logger.Infof("Sor Webhook request response body %v\n", resp.Body)
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
}
