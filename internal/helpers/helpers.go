package helpers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"golang.org/x/crypto/bcrypt"
)

// ParseForm pour décoder l'http request de Gorillan schema
func ParseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println(bytes)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}