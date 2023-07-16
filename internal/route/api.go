package route

import (
	"App/internal/handlers"
	"App/internal/middlewares"

	"github.com/go-chi/chi/v5"
)

func route(router *chi.Mux, handlerService *handlers.HandlerService) {

	router.Use(middlewares.SetJSONHeaders)
	router.Post("/signup", handlerService.StoreUser)

	router.Post("/login", handlerService.Login)

	// Groupe de routes protégées par le middleware d'authentification

	router.Group(func(r chi.Router) {
		// Middleware d'authentification
		r.Use(middlewares.CheckMJWTValidity)

		// Routes protégées
		r.Get("/profil", handlerService.GetAll)
		r.Patch("/profil/user/{id}", handlerService.Update)
		r.Get("/getDatas", handlerService.GetDatasFromDates)
		r.Get("/getRooms", handlerService.GetAllRooms)
		r.Get("/GetAllDatasByRoom", handlerService.GetAllDatasbyRooms)
		r.Get("/getAllDatasByRoomByDates", handlerService.GetAllDatasbyRoomsByDate)
	})

	router.NotFound(notfound)
}
