package hash 

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// HMAC est un wrapper englobant package crypto/hmac
type HMAC struct {
	hmac hash.Hash
}

// NewHMAC créé HMAC object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// Hash l'input en utilisant HMAC
// Avec la clé secrète ajouté quand HMAC est créé
func (h HMAC) Hashing(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}