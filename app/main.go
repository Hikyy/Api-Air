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
	// "gorm.io/driver/postgres"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
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

	// !!!!!!!!!!!!! REMPLACER LE HOST PAR LA VALEUR DONNEE PAR CETTE COMMANDE !!!!!!!!!!!!!!
	// docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' {{ID_CONTAINER}}
	var (
		host     = "127.0.0.1"
		port     = 	"3307"
		dbuser   = "root"
		password = "root"
		dbname   = "db"
	)

    // str := cfg.FormatDSN()
	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, port, dbuser, password, dbname)


	us, err := models.NewUserService(psqlInfo)
	fmt.Println(us, err)
	listenPort := "localhost:8097"
	
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// defer us.Close()
	// us.DBDestructiveReset()
	us.Ping()

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

	fmt.Println("après new router", r)
	// r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println("listening on port ", listenPort)
    http.ListenAndServe(listenPort, r)
	fmt.Println("EOF")
}
