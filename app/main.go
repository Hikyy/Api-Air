package main

// il est necessaire de relancer le code à chaque fois 
// chi.router 
// gofiber en server http2
// GORM à vérifier 

// MOCHI MQTT broker 
// PAHO MQTT client MQTT pour se co au Broker

import (
	//"App/internal/handlers"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
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
		host     = "localhost"
		port     = 	"1333"
		dbuser   = "root"
		password = "root"
		dbname   = "db"
	)

	// psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, dbuser, password, dbname)

	// dsn := "host=localhost user=root password=root dbname=db port=1333 sslmode=disable TimeZone=Europe/Paris"
	dsn := "host=" + host + " user=" + dbuser + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"


	fmt.Println(psqlInfo)
	us, err := models.NewUserService(dsn)
	fmt.Println(us, err)
	listenPort := "8080"
	
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	defer us.Close()
	us.DBDestructiveReset()

	usersC := controllers.NewUsers(us) // Handling User Controller
	
    // r := chi.NewRouter()
	// r.Use(middleware.RequestID)
    // r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

    // r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type","application/json")
    //     w.Write([]byte("Hello World!"))
    // })

	r := mux.NewRouter()

	// r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println("listening on port ", listenPort)
    http.ListenAndServe(listenPort, r)
}
