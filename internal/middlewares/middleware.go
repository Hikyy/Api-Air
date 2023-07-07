package middlewares

import (
	"net/http"
)

func FormRequestCall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(rw, r)
	})
}

// ici j'implementerais le middleWare présent sur mon windows
