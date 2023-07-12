package handlers

import (
	"App/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

func SetMQTT(broker string, username string, password string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(MessagePubHandler)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Abonnement au topic de réception

	//packets := [3]string{"1ac45e2c-2bc2-4027-a7f6-0dbcafcad53b", "1ac45e2c-2bc2-4027-a7f6-0dbcafcad53b", "1ac45e2c-2bc2-4027-a7f6-0dbcafcad53b"}
	// dict := map[interface{}]interface{}{
	// 	1:     "hello",
	// 	"hey": 2,
	// }
	// fmt.Println(dict) // map[1:hello hey:2]
	topic := []string{"groupe9/packet/5e178fd2-5321-4cf5-b04c-4c6a8a827d88/db0b2380-acf0-4688-b219-04ad29c369f3/102", "groupe9/packet/a95cec4a-8aaf-4204-9fa2-b6c4aa8779e7/5072c2d9-cd1b-4e3d-aedb-9ddf19b25abc/131", "groupe9/packet/1ac45e2c-2bc2-4027-a7f6-0dbcafcad53b/42d166ae-94b8-4445-b122-91fbc77bb3c1/118"}
	for _, topic := range topic {
		if token := client.Subscribe(topic, 0, MessagePubHandler); token.Wait() && token.Error() != nil {
			fmt.Printf("Error subscribing to topic %s: %v\n", topic, token.Error())
		} else {
			fmt.Printf("Subscribed to topic: %s\n", topic)
		}
	}

	//Envoi de messages
	data := models.SensorData{
		Current:          13,
		Voltage:          336,
		ActivePower:      0,
		FundamentalPower: 0,
		ReactivePower:    -2,
		ApparentPower:    4,
		Phase:            16402,
	}

	//envoyer message
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON:", err)
		return
	}
	for _, topic := range topic {
		publishMessage(client, topic, string(jsonData))
	}

	<-c // Attente de l'interruption du signal (CTRL+C)
}
func publishMessage(client MQTT.Client, topic, payload string) {
	token := client.Publish(topic, 0, false, payload)
	fmt.Println(payload)
	token.Wait()
	time.Sleep(1 * time.Second) // Attente pour laisser le temps au message d'être publié
}

//type MessageHandler func(clien MQTT.Client, message MQTT.Message)

// func messageHandler(client MQTT.Client, msg MQTT.Message) {
// 	var data map[string]interface{}
// 	err := json.Unmarshal(msg.Payload(), &data)

// 	if err != nil {
// 		fmt.Println("Erreur lors du décodage du payload JSON:", err)
// 		return
// 	}

// }
