package auth

import (
	"App/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var Secretkey = "zkdopqkdQZaLDZMdoqkoSMPDQZPdl8QSdmq"

func GenerateJWT(email string, id int) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func DecodeJWT(token string) (*models.TokenClaim, error) {
	var mySigningKey = []byte(Secretkey)

	claims := &models.TokenClaim{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !parsedToken.Valid {
		//return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}
