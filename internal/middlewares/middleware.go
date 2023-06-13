package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

func FormRequestCall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//var user User

		//err := json.NewDecoder(r.Body).Decode(&user)
		/*		if err != nil {
					fmt.Println(err)
					return
				}

				ctx := context.WithValue(r.Context(), "user", user)*/

		next.ServeHTTP(rw, r)
	})
}

func middleware(next http.Handler) http.Handler {
	return http.Handler(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		}
		next.ServeHTTP(w, r)
	})
}
