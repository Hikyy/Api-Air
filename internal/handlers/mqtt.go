package handlers

import (
	"App/internal/config"
	"App/internal/helpers"
	"App/internal/models"
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	connSuccess = make(chan bool)
	connError   = make(chan error)
)
var MessagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	var jsonString = msg.Payload()
	var sensorData models.SensorData
	var sensorDatatoDb models.SensorDataToDb

	err := json.Unmarshal([]byte(jsonString), &sensorData)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}

	datas := &Datas{
		dts: models.Db,
	}

	var sendDataToDB = func(dt *Datas, data *models.SensorDataToDb) (error, *http.Request) {
		return dt.dts.AddDataToDb(data), nil
	}
	fmt.Println("SensorID:", sensorData.SensorID)
	fmt.Println("TimeEpoch  :", helpers.TimeStampConverter(sensorData.EventTimestamp))

	converted := helpers.TimeStampConverter(sensorData.EventTimestamp)
	sensorDatatoDb.SensorID = sensorData.SensorID
	sensorDatatoDb.EventTimestamp = converted
	sensorDatatoDb.EventData = sensorData.Data

	for key, value := range sensorData.Data {
		sendDataToDB(datas, &sensorDatatoDb)
		fmt.Printf("%s: %v\n", key, value)
	}
}

func setMQTT() MQTT.Client {
	config := config.ClientConfig()
	opts := MQTT.NewClientOptions()

	opts.AddBroker(config["broker"])
	opts.SetUsername(config["username"])
	opts.SetPassword(config["password"])
	opts.SetDefaultPublishHandler(MessagePubHandler)
	opts.SetOrderMatters(true)

	client := MQTT.NewClient(opts)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		fmt.Println("MQTT client is not connected. Error:", token.Error())
		return nil
	} else {
		fmt.Println("MQTT client is connected.")
	}
	return client
}

func SubscribeTopic(c chan os.Signal) {
	client := setMQTT()
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		connError <- token.Error()
		return
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
	connSuccess <- true
	<-c
}

func SendRequest(c chan os.Signal) {
	client := setMQTT()
	if !client.IsConnected() {
		fmt.Println("MQTT client is not connected.")
		return
	}

	select {
	case err := <-connError: // Receive error from connError channel
		fmt.Println("Erreur de connexion MQTT: ", err)
		return
	case <-connSuccess: // Success received from connSuccess channel

		command := map[string]interface{}{
			"cmd_id":              102,
			"destination_address": "db0b2380-acf0-4688-b219-04ad29c369f3",
			"ack_flags":           0,
			"cmd_type":            208,
		}

		jsonData, err := json.Marshal(command)
		if err != nil {
			fmt.Println("Erreur lors de la conversion en JSON :", err)
			return
		}

		topic := "groupe9/request/5e178fd2-5321-4cf5-b04c-4c6a8a827d88"
		token := client.Publish(topic, 0, false, jsonData)
		token.Wait()

		if token.Error() != nil {
			fmt.Printf("Error sending request to topic %s: %v\n", topic, token.Error())
		} else {
			fmt.Printf("Requete envoyée au topic: %s\n", topic)
		}
	}
	<-c
}
