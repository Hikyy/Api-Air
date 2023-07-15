package handlers

import (
	"App/internal/config"
	"App/internal/database"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	host     = "localhost"
	port     = "3307"
	dbuser   = "root"
	password = "root"
	dbname   = "postgres"
)

func SetWebSocket() {

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbuser, password, dbname)

	configPostgres := config.Postgres()
	configJSON, err := json.Marshal(configPostgres)
	if err != nil {
		return
	}
	configString := string(configJSON)
	fmt.Println(configString)
	db, err := gorm.Open(postgres.Open(database.BuildConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer db.Close()

	test := db.Exec("LISTEN sensor_event_inserted")
	fmt.Printf("DB EXEC => %s", *test)

	listener := pq.NewListener(connString, time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	err = listener.Listen("sensor_event_inserted")
	if err != nil {
		log.Fatal(err)
		return
	}
	notificationChannel := make(chan *pq.Notification)
	go func() {
		for {
			select {
			case notification := <-notificationChannel:
				if notification.Channel == "sensor_event_inserted" {
					//fmt.Println("Sensor event inserted:", notification.Extra)
					// Faites ce que vous voulez lorsque vous recevez une notification pour un événement de capteur inséré
				}
			}
		}
	}()

	for {
		select {}
	}

}
