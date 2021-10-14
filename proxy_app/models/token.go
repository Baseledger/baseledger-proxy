package models

import (
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
