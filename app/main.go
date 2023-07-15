package main

import (
	"App/internal/handlers"
	"App/internal/models"
	"fmt"
	"net/http"
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

	handlers.StartSQL()
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
