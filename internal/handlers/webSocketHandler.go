package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"os"
	"time"
)

var (
	host     = "localhost"
	port     = "3307"
	dbuser   = "root"
	password = "root"
	dbname   = "postgres"
)

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
