package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = "3307"
	dbuser   = "root"
	password = "root"
	dbname   = "postgres"
)

// func SetWebSocket() {
//
//	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbuser, password, dbname)
//
//	configPostgres := config.Postgres()
//	configJSON, err := json.Marshal(configPostgres)
//	if err != nil {
//		return
//	}
//	configString := string(configJSON)
//	fmt.Println(configString)
//	db, err := gorm.Open(postgres.Open(database.BuildConnectionString()), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	//defer db.Close()
//
//	result := db.Exec("LISTEN sensor_event_inserted")
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	fmt.Printf("DB EXEC => %+v\n", *result)
//
//	listener := pq.NewListener(connString, time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
//		if err != nil {
//			log.Fatal(err)
//			return
//		}
//	})
//
//	err = listener.Listen("sensor_event_inserted")
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	notificationChannel := make(chan *pq.Notification)
//	go func() {
//		for {
//			select {
//			case notification := <-notificationChannel:
//				if notification.Channel == "sensor_event_inserted" {
//					//fmt.Println("Sensor event inserted:", notification.Extra)
//					// Faites ce que vous voulez lorsque vous recevez une notification pour un événement de capteur inséré
//				}
//			}
//		}
//	}()
//
//	for {
//		select {}
//	}
//
// }
func StartSQL(c chan os.Signal) {
	//var conninfo string = "dbname=exampledb user=webapp password=webapp"
	conninfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbuser, password, dbname)

	_, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, reportProblem)
	err = listener.Listen("events")
	if err != nil {
		panic(err)
	}

	fmt.Println("Start monitoring PostgreSQL...")
	for {
		WaitForNotification(listener)
	}
}
func WaitForNotification(l *pq.Listener) {
	for {
		select {
		case n := <-l.Notify:
			fmt.Println("Received data from channel [", n.Channel, "] :")
			// Prepare notification payload for pretty print
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
			if err != nil {
				fmt.Println("Error processing JSON: ", err)
				return
			}
			fmt.Println(string(prettyJSON.Bytes()))
			return
		case <-time.After(90 * time.Second):
			fmt.Println("Received no events for 90 seconds, checking connection")
			go func() {
				l.Ping()
			}()
			return
		}
	}
}
