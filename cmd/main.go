package main

import (
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    opts := mqtt.NewClientOptions().AddBroker("tcp://mosquitto:1883").SetClientID("go_subscriber")
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("Błąd połączenia: %s", token.Error())
    }

    if token := client.Subscribe("devices/+/measurements", 0, func(client mqtt.Client, msg mqtt.Message) {
        topicParts := strings.Split(msg.Topic(), "/")
        if len(topicParts) >= 3 {
            deviceId := topicParts[1]
            log.Printf("Odebrano wiadomość z urządzenia %s: %s", deviceId, msg.Payload())
        } else {
            log.Printf("Nieprawidłowy format tematu: %s", msg.Topic())
        }
    }); token.Wait() && token.Error() != nil {
        log.Fatalf("Błąd subskrypcji: %s", token.Error())
    }

    select {}
}