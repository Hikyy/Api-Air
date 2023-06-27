package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// RememberTokenBytes utilisé comme valeur par défault de taille de byte de token
const RememberTokenBytes = 32

// Bytes pour génerer des bytes random
// return une erreur si existe
// utilise crypto/rand pour safety remember token
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// CryptoString genere un slice de byte de la taille nByte
// puis retourn une string en base64 URL encoded de ce slice de byte
func CryptoString(nByte int) (string, error) {
	b, err := Bytes(nByte)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken est une func helper pour génerer remember token 
// avec le prédéfinis byte size "32"
func RememberToken() (string, error) {
	return CryptoString(RememberTokenBytes)
}