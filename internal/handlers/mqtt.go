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
	var message models.LightData

	err := json.Unmarshal(msg.Payload(), &message)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON :", err)
		return
	}
	fmt.Println(err)
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

	time.Sleep(time.Second)
	topic := []string{"groupe9/packet/5e178fd2-5321-4cf5-b04c-4c6a8a827d88/db0b2380-acf0-4688-b219-04ad29c369f3/102", "groupe9/packet/a95cec4a-8aaf-4204-9fa2-b6c4aa8779e7/5072c2d9-cd1b-4e3d-aedb-9ddf19b25abc/131", "groupe9/packet/1ac45e2c-2bc2-4027-a7f6-0dbcafcad53b/42d166ae-94b8-4445-b122-91fbc77bb3c1/118"}
	for _, topic := range topic {
		if token := client.Subscribe(topic, 0, MessagePubHandler); token.Wait() && token.Error() != nil {
			fmt.Printf("Error subscribing to topic %s: %v\n", topic, token.Error())
		} else {
			fmt.Printf("Subscribed to topic: %s\n", topic)
		}
	}

	log.Printf("Connecté au broker MQTT %s et abonné au topic %s\n ", broker, topic)

	<-c // Attente de l'interruption du signal (CTRL+C)
}

