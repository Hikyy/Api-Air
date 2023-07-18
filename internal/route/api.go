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
		// r.Use(middlewares.CheckMJWTValidity)

		// Routes protégées

		// Route Condition
		r.Get("/conditions", handlerService.IndexCondition)
		r.Post("/conditions", handlerService.StoreCondition)

		// Route User
		r.Get("/profils", handlerService.IndexProfils)
		r.Patch("/profil/user/{id}", handlerService.Update)
		r.Delete("/profil/user/{id}", handlerService.Delete)

		// Route Actuator
		r.Get("/actuators", handlerService.IndexActuators)

		// Route Room
		r.Get("/rooms", handlerService.IndexRooms)

		// Route Sensors
		r.Get("/sensors", handlerService.IndexSensors)


		// Route Sensor Event
		r.Get("/sensor-events", handlerService.IndexSensorEvents)
		r.Get("/room/{id}/sensor-events", handlerService.IndexRoomSensorEvents)
		r.Get("/room/{id}/sensor-events/{date}", handlerService.IndexRoomSensorEventsByDate)
		r.Get("/room/{id}/sensor-events/{date-debut}/{date-fin}", handlerService.IndexRoomSensorEventsBetweenTwoDates)
		r.Get("/room/sensor-events", handlerService.IndexSensorEventsByIdByRoomByDate)
	})

	router.NotFound(notfound)
}
