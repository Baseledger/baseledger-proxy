package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func GetToken(email string) (string, error) {
	jwtSecret := []byte(viper.GetString("JWT_SECRET"))

	// Create the Claims
	proxyClaims := jwt.MapClaims{}
	proxyClaims["authorized"] = true
	proxyClaims["email"] = email
	proxyClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, proxyClaims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %v", token.Header["alg"])
		}
		jwtSecret := []byte(viper.GetString("JWT_SECRET"))
		return []byte(jwtSecret), nil
	})
}
