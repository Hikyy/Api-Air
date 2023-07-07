package handlers

import (
	"App/internal/models"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"time"
)

func SetMQTT(broker string, clientID string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Abonnement au topic de réception
	topic := "mytopic"

	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Printf("Connecté au broker MQTT %s et abonné au topic %s\n", broker, topic)

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
	fmt.Printf("Message reçu sur le topic: %s\n", msg.Topic())
	fmt.Printf("Payloadss: %s\n", msg.Payload())
	// var data map[string]string

	var data map[string]interface{}
	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		fmt.Println("Erreur lors du décodage du payload JSON:", err)
		return
	}

	current, ok := data["current"].(float64)
	if !ok {
		fmt.Println("Erreur lors de l'accès à l'indice 'current'")
		return
	}

	fmt.Println("Valeur de 'current':", current)

}

func publishMessage(client MQTT.Client, topic, payload string) {
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
	time.Sleep(1 * time.Second) // Attente pour laisser le temps au message d'être publié
}
