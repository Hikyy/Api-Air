package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func route(router *chi.Mux, userHandler *Users) {

	router.Post("/signup", userHandler.Create)

	router.Post("/login", userHandler.Login)

	router.Get("/profil", userHandler.GetAll)

	router.Patch("/profil/user/{id}", userHandler.Update)

	// router.Get("/{salle}", userHandler.GetRoom)

	// router.Get("/{salle}/{captor}", userHandler.GetCaptorByRoom)

	router.NotFound(notfound)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page not found")
}