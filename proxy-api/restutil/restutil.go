package restutil

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const defaultResponseContentType = "application/json; charset=UTF-8"

type BaseReq struct {
	From    string `json:"from"`
	ChainID string `json:"chain_id"`
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
