package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetToken(email string) (string, error) {
	// move to env
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	proxyClaims := jwt.MapClaims{}
	proxyClaims["authorized"] = true
	proxyClaims["email"] = email
	proxyClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, proxyClaims)
	return token.SignedString(mySigningKey)
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %v", token.Header["alg"])
		}
		// move to env
		mySigningKey := []byte("AllYourBase")
		return []byte(mySigningKey), nil
	})
}
