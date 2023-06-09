package main

// il est necessaire de relancer le code à chaque fois 
// chi.router 
// gofiber en server http2
// GORM à vérifier 

// MOCHI MQTT broker 
// PAHO MQTT client MQTT pour se co au Broker

import (
	// "App/internal/handlers"
	//"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
	r.Use(middleware.RequestID)
    r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

	//fmt.Println("working")

    http.ListenAndServe(":8099", r)
}
