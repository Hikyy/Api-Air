package middlewares

import (
	"App/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func FormRequestCall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(rw, r)
	})
}

func CheckMJWTValidity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Cookie("TokenBearer"))

		cookie, err := r.Cookie("TokenBearer")

		successAndCookie := models.TokenValidityToken{Jwt: true, Cookie: cookie}
		successStatus, _ := json.Marshal(successAndCookie)

		if err != nil {
			success := models.TokenValidity{Jwt: false}
			successStatus, _ = json.Marshal(success)
			w.Write(successStatus)
			return
		}

		w.Write(successStatus)
		next.ServeHTTP(w, r)
	})
}
