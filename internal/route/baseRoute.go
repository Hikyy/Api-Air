package route

import (
	"App/internal/handlers"
	"App/internal/models"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page not found")
}

func SetupRouter() http.Handler {
	router := chi.NewRouter()

	// router.Use(middlewares.TransactionMiddleware(models.InitGorm))

	// router.Use(middlewares.TransactionMiddleware)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	router.Use(cors.Handler)

	handlerService := handlers.NewHandler(models.Db)

	route(router, handlerService)

	return router
}
