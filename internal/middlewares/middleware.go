package middlewares

import (
	"net/http"
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
