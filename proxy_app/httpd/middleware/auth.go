package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/unibrightio/proxy-api/logger"
	"github.com/unibrightio/proxy-api/token"
)

func AuthorizeJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
