package middlewares

import (
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
		fmt.Println(cookie)

		if err != nil {
			return
		}

		next.ServeHTTP(w, r)
	})
}
