package main

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    opts := mqtt.NewClientOptions().AddBroker("tcp://mosquitto:1883").SetClientID("go_subscriber")
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("Błąd połączenia: %s", token.Error())
    }

    if token := client.Subscribe("test/topic", 0, func(client mqtt.Client, msg mqtt.Message) {
        log.Printf("Odebrano wiadomość: %s", msg.Payload())
    }); token.Wait() && token.Error() != nil {
        log.Fatalf("Błąd subskrypcji: %s", token.Error())
    }

    select {}
}