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

func SetMQTT(broker string, username string, password string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messageHandler)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Abonnement au topic de réception
	topic := "groupe9/packet/1ac45e2c-2bc2-4027-a7f6-0dbcafcad53b"

	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())

	}

	log.Printf("Connecté au broker MQTT %s et abonné au topic %s\n ", broker, topic)

	// Envoi de messages
	data := models.SensorData{
		Current:          13,
		Voltage:          336,
		ActivePower:      0,
		FundamentalPower: 0,
		ReactivePower:    -2,
		ApparentPower:    4,
		Phase:            16402,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON:", err)
		return
	}
	publishMessage(client, topic, string(jsonData))

	<-c // Attente de l'interruption du signal (CTRL+C)
}
func messageHandler(client MQTT.Client, msg MQTT.Message) {

	var data map[string]interface{}
	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		fmt.Println("Erreur lors du décodage du payload JSON:", err)
		return
	}
	// fmt.Println("data:", json.Unmarshal([]byte(msg.Payload())))
	for key, value := range data {
		fmt.Println("value:", value, "key:", key)
	}
	fmt.Printf("Message reçu sur le topic: %s\n", msg.Topic())

	fmt.Printf("Vos daronne qui font l'emeutes")

}
func publishMessage(client MQTT.Client, topic, payload string) {
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
	time.Sleep(1 * time.Second) // Attente pour laisser le temps au message d'être publié
}
