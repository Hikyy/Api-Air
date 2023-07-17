package main

import (
	"App/internal/models"
	"App/internal/route"
	"fmt"
	"net/http"
	// "os"
	// "os/signal"
)

func main() {

	err := models.DatabaseServiceProvider()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	defer models.InitGorm.Close()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)

	//go handlers.SubscribeTopic(c)
	//go handlers.SendRequest(c)
	//go handlers.StartSQL(c)

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	router := route.SetupRouter()

	// customRouter := &handlers.CustomRouter{
	// 	Router: router,
	// 	C:      c,
	// }

	fmt.Println("Server listening on port 8097")

	http.ListenAndServe(":8097", router)
}
