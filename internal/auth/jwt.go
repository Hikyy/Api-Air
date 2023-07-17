package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var Secretkey = "zkdopqkdQZaLDZMdoqkoSMPDQZPdl8QSdmq"

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["id"] = "JWT"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
