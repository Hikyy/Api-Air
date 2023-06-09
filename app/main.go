package main

import (
	"App/internal/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	handlers.Handlers(router)

	//config.Sql()
	err := http.ListenAndServe(":8097", router)
	if err != nil {
		_ = fmt.Errorf("impossible de lancer le serveur : %w", err)
		return
	}
}
