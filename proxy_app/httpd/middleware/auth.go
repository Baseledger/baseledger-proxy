package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/token"
)

func BasicAuth(fallbackToJwt bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		basicAuthUser, _ := viper.Get("API_UB_USER").(string)
		basicAuthPwd, _ := viper.Get("API_UB_PWD").(string)
		// Get the Basic Authentication credentials
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == basicAuthUser && password == basicAuthPwd {
			logger.Info("Basic auth successful")
			// Setting this flag inside this context so next middleware knows it is already auth
			if fallbackToJwt {
				c.Set("auth", true)
			}
		} else {
			if fallbackToJwt {
				logger.Error("Basic auth failed, trying with jwt")
				c.Next()
			} else {
				logger.Error("Basic auth failed")
				c.Abort()
				c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
				c.JSON(http.StatusForbidden, map[string]interface{}{"error": "auth failed"})
			}
		}
	}
}

func AuthorizeJWTMiddleware(isFallback bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		basicAuthSucceded, basicAuthParamExists := c.Get("auth")
		if isFallback && basicAuthParamExists == true && basicAuthSucceded == true {
			logger.Infof("Already auth with basic auth")
			c.Next()
			return
		}
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BEARER_SCHEMA)+1 {
			logger.Errorf("Auth error, no token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		token, err := token.ValidateToken(tokenString)
		if err != nil {
			logger.Errorf("Auth error %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			logger.Infof("Token valid, claims %v", claims)
		} else {
			logger.Errorf("Auth error %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
