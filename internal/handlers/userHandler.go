package handlers

import (
	"App/internal/requests"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	Safe(requests.StoreUserRequest(), writer)
	//fmt.Printf("%v", validated)
}
