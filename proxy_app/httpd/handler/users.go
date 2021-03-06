package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	uuid "github.com/kthomas/go.uuid"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/restutil"
	"github.com/unibrightio/proxy-api/types"
)

type createTxDto struct {
	Payload string `json:"payload"`
	OpCode  uint   `json:"op_code"`
}

type userDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Register user ... Register user
// @Summary Register user
// @Description register user
// @Tags Dev
// @Accept json
// @Param user body userDto true "User data"
// @Success 200 {string} email
// @Failure 400,422,500 {string} errorMessage
// @Router /dev/users [post]
func CreateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		user := &types.User{}
		err = json.Unmarshal(buf, &user)
		if err != nil {
			restutil.RenderError(err.Error(), 422, c)
			return
		}

		if !user.Create() {
			logger.Error("error when creating new user")
			restutil.RenderError("error when creating new user", 500, c)
			return
		}

		restutil.Render(user.Email, 200, c)
	}
}

// Login user ... Login user
// @Summary Login user
// @Description login user
// @Tags Dev
// @Accept json
// @Param user body userDto true "User data"
// @Success 200 {string} acessToken
// @Failure 400,422,500 {string} errorMessage
// @Router /dev/auth [post]
func LoginUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		user := &types.User{}
		err = json.Unmarshal(buf, &user)
		if err != nil {
			restutil.RenderError(err.Error(), 422, c)
			return
		}

		token, err := user.Login()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		restutil.Render(token, 200, c)
	}
}

// @Security BearerAuth
// Generate transaction with custom payload ... Generate transaction with custom payload
// @Summary Generate transaction with custom payload
// @Description generate transaction with custom payload
// @Tags Dev
// @Accept json
// @Param user body createTxDto true "Transaction payload"
// @Success 200 {string} txHash
// @Failure 400,422,500 {string} errorMessage
// @Router /dev/tx [post]
func CreateTransactionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		req := &createTxDto{}
		err = json.Unmarshal(buf, &req)
		if err != nil {
			restutil.RenderError(err.Error(), 422, c)
			return
		}

		if req.OpCode != 9 {
			restutil.RenderError("currently only op code 9 is supported", 400, c)
			return
		}

		maximumPayloadSize := 128
		if len(req.Payload) > maximumPayloadSize {
			restutil.RenderError("payload maximum size exceeded", 400, c)
			return
		}

		transactionId := uuid.NewV4()
		signAndBroadcastPayload := restutil.SignAndBroadcastPayload{
			TransactionId: transactionId.String(),
			Payload:       req.Payload,
			OpCode:        uint32(req.OpCode),
		}

		txHash := restutil.SignAndBroadcast(signAndBroadcastPayload)

		logger.Infof("Transaction hash of custom payload %v\n", txHash)
		restutil.Render(txHash, 200, c)
	}
}
