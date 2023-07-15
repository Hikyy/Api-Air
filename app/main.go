package main

import (
	"App/internal/handlers"
	"App/internal/models"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

var (
	broker   = "mqtt://mqtt.arcplex.fr:2295"
	username = "groupe9"
	password = "Pu3a76ZS0pgT"
)

func main() {

	err := models.DatabaseServiceProvider()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer models.InitGorm.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go handlers.StartSQL(c)
	go handlers.SetMQTT(broker, username, password, c)

	//handlers.StartSQL()
	//handlers.SetMQTT(broker, username, password)

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	router := handlers.SetupRouter()

	fmt.Println("Server listening on port 8097")

	// err = godotenv.Load()

	http.ListenAndServe(":8097", router)

}
