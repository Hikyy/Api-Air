package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	host     = "localhost"
	port     = "5432"
	dbuser   = "root"
	password = "root"
	dbname   = "postgres"
)
var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade la connexion HTTP en une connexion WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur lors de la mise à niveau du WebSocket :", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connecté.")
	clients[conn] = true // Ajoutez le client WebSocket à la liste

	// Commencez à écouter les messages WebSocket du client (optionnel)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erreur de lecture du WebSocket :", err)
			delete(clients, conn)
			break
		}
	}
}

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

	http.Handle("/ws", http.HandlerFunc(handleWebSocketConnection))

	fmt.Println("Start monitoring PostgreSQL...")
	for {
		WaitForNotification(listener)
	}
}

func sendToClients(data []byte) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Erreur lors de l'envoi du message WebSocket :", err)
			client.Close()
			delete(clients, client)
		}
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
			sendToClients(prettyJSON.Bytes())
			return
		case <-time.After(10 * time.Second):
			fmt.Println("Received no events for 90 seconds, checking connection")
			go func() {
				l.Ping()
			}()
			return
		}
	}
}
