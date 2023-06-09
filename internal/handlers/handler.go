package handlers

import (
	"App/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*chi.Mux
}

func Handlers(mux *chi.Mux) {
	handler := &Handler{
		mux,
	}

	handler.Use(middlewares.FormRequestCall)
	handler.Get("/login", Login)
	handler.Route("/api/login", func(r chi.Router) {
	})
}
