package handlers

import (
	"App/internal/middlewares"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func route(router *chi.Mux, userHandler *Users) {

	router.Post("/signup", userHandler.Create)

	router.Post("/login", userHandler.Login)

	router.Group(func(r chi.Router) {
		// Middleware d'authentification
		r.Use(middlewares.CheckMJWTValidity)

		// Routes protégées
		r.Post("/conditions", userHandler.StoreCondition)
		r.Patch("/profil/user/{id}", userHandler.Update)
		r.Get("/conditions", userHandler.IndexCondition)
		r.Get("/profil", userHandler.GetAll)
		r.Get("/sensor-events", userHandler.IndexSensorEvents)
		r.Get("/rooms", userHandler.IndexRooms)
		r.Get("/room/{id}/sensor-events", userHandler.IndexRoomSensorEvents)
		r.Get("/room/{id}/sensor-events/{date}", userHandler.IndexRoomSensorEventsByDate)
		r.Get("/room/{id}/sensor-events/{date-debut}/{date-fin}", userHandler.IndexRoomSensorEventsBetweenTwoDates)
	})

	router.NotFound(notfound)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page not found")
}
