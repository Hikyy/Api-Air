package main

// il est necessaire de relancer le code à chaque fois 
// chi.router 
// gofiber en server http2
// GORM à vérifier 

// MOCHI MQTT broker 
// PAHO MQTT client MQTT pour se co au Broker

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

	"App/internal/controllers"
	"App/internal/models"
)

// Error Handling
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// HTTP 404 NotFound
func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page not found")
}

var psqlInfo string


func main() {

	var (
		host     = "127.0.0.1"
		port     = 	"3307"
		dbuser   = "root"
		password = "root"
		dbname   = "db"
	)

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

	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json") // Définit le header en "Application/Json"
		fmt.Println("w => ", &w)
		// Appel de la fonction Create du contrôleur Users
		usersC.Create(w, r)
	}).Methods("POST")

	fmt.Println("listening on port ", listen)

    http.ListenAndServe(listen, r)
}
