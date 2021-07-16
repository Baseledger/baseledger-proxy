package restutil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unibrightio/proxy-api/logger"
)

const defaultResponseContentType = "application/json; charset=UTF-8"

type BaseReq struct {
	From    string `json:"from"`
	ChainID string `json:"chain_id"`
}

type SignAndBroadcastPayload struct {
	BaseReq       BaseReq `json:"base_req"`
	TransactionId string  `json:"transaction_id"`
	Payload       string  `json:"payload"`
}

func SignAndBroadcast(payload SignAndBroadcastPayload, c *gin.Context) *string {
	jsonValue, err := json.Marshal(payload)

	if err != nil {
		logger.Error("Error marshaling sign and broadcast json")
		return nil
	}

	// All of these must be read from ENV. target should be localhost from host and host.docker.internal if dockerized
	resp, err := http.Post("http://host.docker.internal:1317/signAndBroadcast", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		logger.Errorf("error while sending feedback request %v\n", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("error while reading sign and broadcast transaction response %v\n", err.Error())
		return nil
	}

	txHash := string(body)
	return &txHash
}

// Render an object and status using the given gin context
func Render(obj interface{}, status int, c *gin.Context) {
	c.Header("content-type", defaultResponseContentType)
	c.Writer.WriteHeader(status)
	if &obj != nil && status != http.StatusNoContent {
		encoder := json.NewEncoder(c.Writer)
		encoder.SetIndent("", "    ")
		if err := encoder.Encode(obj); err != nil {
			panic(err)
		}
	} else {
		c.Header("content-length", "0")
	}
}

func RenderError(message string, status int, c *gin.Context) {
	err := map[string]*string{}
	err["message"] = &message
	Render(err, status, c)
}
