package main

// il est necessaire de relancer le code à chaque fois
// chi.router
// gofiber en server http2
// GORM à vérifier

// MOCHI MQTT broker
// PAHO MQTT client MQTT pour se co au Broker

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"App/internal/controllers"
	"App/internal/models"
)

var psqlInfo string
var broker = "mqtt://mqtt.arcplex.fr:2295" // Remplacez par l'adresse et le port de votre broker MQTT
var clientID = "go-mqtt-client"

func main() {

	var (
		host     = "127.0.0.1"
		port     = "3307"
		dbuser   = "root"
		password = "root"
		dbname   = "postgres"
	)

	// test SetMQTT
	//handlers.SetMQTT(broker, clientID)

	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbuser, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	fmt.Println(us, err)

	listen := "localhost:8097"

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	us.Ping()

	usersC := controllers.NewUsers(us) // Handling User Controller
	r := mux.NewRouter()

	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("w => ", &w)
		usersC.Create(w, r)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		usersC.Login(w, r)
	}).Methods("POST")

	fmt.Println("listening on port ", listen)

	http.ListenAndServe(listen, r)
}
