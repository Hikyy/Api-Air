package handlers

import (
	"App/internal/config"
	"App/internal/helpers"
	"App/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var Client MQTT.Client

var MessagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("Messageeeeeeee %s received on topic %s\n", msg.Payload(), msg.Topic())
	//Test()

	var jsonString = msg.Payload()
	var sensorData models.SensorData
	var sensorDatatoDb models.SensorDataToDb
	err := json.Unmarshal([]byte(jsonString), &sensorData)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}
	datas := &Datas{
		dts: models.Db, // Initialisez le champ dts avec l'objet approprié
	}
	var prout = func(dt *Datas, data *models.SensorDataToDb) (error, *http.Request) {
		return dt.dts.AddDataToDb(data), nil
	}

	//fmt.Println("SensorAddress :", sensorData.SensorAddress)
	fmt.Println("SensorID:", sensorData.SensorID)
	fmt.Println("TimeEpoch  :", helpers.TimeStampConverter(sensorData.EventTimestamp))
	converted := helpers.TimeStampConverter(sensorData.EventTimestamp)
	sensorDatatoDb.SensorID = sensorData.SensorID
	sensorDatatoDb.EventTimestamp = converted
	sensorDatatoDb.EventData = sensorData.Data

	for key, value := range sensorData.Data {
		prout(datas, &sensorDatatoDb)
		fmt.Printf("%s: %v\n", key, value)
	}
}

func SetMQTT(broker string, username string, password string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(MessagePubHandler)
	opts.SetOrderMatters(true)
	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	time.Sleep(time.Second)
	topics := config.Salles

	for key, value := range topics {
		for keySensor, valueSensor := range value {
			request := key + keySensor + valueSensor
			if token := client.Subscribe(request, 0, MessagePubHandler); token.Wait() && token.Error() != nil {
				fmt.Printf("Error subscribing to topic %s: %v\n", request, token.Error())
			} else {
				fmt.Printf("Subscribed to topic: %s\n", request)
			}
		}
	}

	<-c // Attente de l'interruption du signal (CTRL+C)
}
func SendRequest(client MQTT.Client) {
	// set le type Message
	groupe := "groupe9"
	geteway_id := "a95cec4a-8aaf-4204-9fa2-b6c4aa8779e7"
	message := models.Message{
		CmdID:              "1",
		DestinationAddress: "3933a900-4f73-4716-9a97-42f5353c9254",
		Ack_Flags:          "0",
		Cmd_Type:           "207",
	}

	// Convertissez l'objet Message en JSON
	request, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON :", err)
		return
	}
	nodeRoute := fmt.Sprintf("%s/%s/%s", groupe, request, geteway_id)
	// boucler sur les topics et envoyer le message
	topics := config.Salles
	for key, value := range topics {
		for keySensor, valueSensor := range value {
			topic := key + keySensor + valueSensor
			token := client.Publish(topic, 0, false, nodeRoute)
			token.Wait()
			if token.Error() != nil {
				fmt.Printf("Error sending request to topic %s: %v\n", request, token.Error())
			} else {
				fmt.Printf("Requete envoyer au topic: %s\n", topic)
			}
		}
	}
}
