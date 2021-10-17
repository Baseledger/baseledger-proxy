package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/models"
	"github.com/unibrightio/proxy-api/restutil"
)

func CreateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		user := &models.User{}
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

func LoginUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, err := c.GetRawData()
		if err != nil {
			restutil.RenderError(err.Error(), 400, c)
			return
		}

		user := &models.User{}
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

func CreateTransactionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
