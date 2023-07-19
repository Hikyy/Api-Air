package handlers

import (
	"App/internal/models"
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
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[*websocket.Conn]bool)

type WebSocketContext struct {
	C        chan os.Signal
	Listener *pq.Listener
}

func HandleWebSocketConnection(w http.ResponseWriter, r *http.Request, ctx *WebSocketContext) {
	// upgrade en ws
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur lors de la mise à niveau du WebSocket :", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connecté.")
	clients[conn] = true // Ajoutez le client WebSocket à la liste

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

	ctx := &WebSocketContext{
		C:        c,
		Listener: listener,
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		HandleWebSocketConnection(w, r, ctx)

	})

	go http.ListenAndServe(":9098", nil)

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
			if err != nil {
				return
			}
			if err != nil {
				fmt.Println("Erreur lors du décodage JSON:", err)
				return
			}
			decodedMessage := BreakWs(&prettyJSON)
			CompareConditions([]models.Conditions{}, decodedMessage)
			sendToClients(prettyJSON.Bytes())
			return
		case <-time.After(10 * time.Second):
			fmt.Println("Received no events for 10 seconds, checking connection")
			go func() {
				l.Ping()
			}()
			return
		}
	}
}

func CompareConditions(Conditions []models.Conditions, websocket []models.SensorEventJson) {
	datas, err := models.GetConditions()
	if err != nil {
		return
	}
	Conditions = datas
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	fmt.Println(websocket)
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")

	fmt.Println(Conditions)
	for _, condition := range Conditions {
		fmt.Println(condition.DataKey, condition.Operator, condition.Value)
		for _, event := range websocket {
			fmt.Println("##################################################################")
			fmt.Println("eventtttt => data.data", event.Data.Data)
			fmt.Println("eventtttt => data.sensorId", event.Data.SensorID)
			for _, test := range event.Data.Data {
				fmt.Println("C LA LUTTE FINALE", test)
			}
		}
		switch condition.Operator {
		case ">":
			// Avant cela, il faut vérifier si condition.DataKey est égal à JWTWS.DataKey
			// En d'autres termes, ici, si JWTWS.Value > condition.Value =>
			fmt.Println("case : condition.DataKey ", condition.DataKey, " EST ", condition.Operator, " A ", condition.Value)
			break
		case "<":
			fmt.Println("case : ", condition.Operator)
			break
		case "=":
			fmt.Println("case : ", condition.Operator)
			break
		}
	}
}

func BreakWs(webSocket *bytes.Buffer) []models.SensorEventJson {
	webSocketData := models.SensorEventJson{}

	if err := json.Unmarshal(webSocket.Bytes(), &webSocketData); err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
	}

	fmt.Println("webSocketData =>", webSocketData)

	decodedMessages := []models.SensorEventJson{webSocketData}
	return decodedMessages
}
