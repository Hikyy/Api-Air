package handlers

import (
	"App/internal/config"
	"App/internal/helpers"
	"App/internal/models"
	"encoding/json"
	"fmt"

	// "log"

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

	sourceAddress, err := getSourceAddress(msg)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}

	datas := &Datas{
		dts: models.Db,
	}

	var sendDataToDB = func(dt *Datas, data *models.SensorDataToDb, sourceAddress string) {
		dt.dts.AddDataToDb(data, sourceAddress)
	}

	converted := helpers.TimeStampConverter(sensorData.EventTimestamp)
	sensorDatatoDb.SensorID = sensorData.SensorID
	sensorDatatoDb.EventTimestamp = converted
	sensorDatatoDb.EventData = sensorData.Data

	for key, value := range sensorData.Data {
		sendDataToDB(datas, &sensorDatatoDb, sourceAddress)
		fmt.Printf("%s : %v\n", key, value)
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

	//time.Sleep(5 * time.Minute)

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

func SendRequest(jsonData []byte, roomId int) {
	client := setMQTT()
	if !client.IsConnected() {
		fmt.Println("MQTT client is not connected.")
		return
	}

	var gatewayId string

	models.InitGorm.Db.Table("rooms").Where("room_id= ?", roomId).Select("room_key").Find(&gatewayId)

	topic := "groupe9/request/" + gatewayId

	fmt.Printf("%s", jsonData)
	token := client.Publish(topic, 0, false, jsonData)
	token.Wait()

	if token.Error() != nil {
		fmt.Printf("Error sending request to topic %s: %v\n", topic, token.Error())
	} else {
		fmt.Printf("Requete envoyée au topic: %s\n", topic)
	}
}

func getSourceAddress(message MQTT.Message) (string, error) {
	var payload map[string]interface{}

	err := json.Unmarshal([]byte(message.Payload()), &payload)

	if err != nil {
		return "", err
	}

	if sourceAddress, ok := payload["source_address"].(string); ok {
		return sourceAddress, nil
	}

	return "", fmt.Errorf("Clé 'source_address' introuvable ou de type incorrect")
}
