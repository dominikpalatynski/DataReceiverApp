package broker

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewClient(brokerURL string) (mqtt.Client, error) {
    opts := mqtt.NewClientOptions().AddBroker(brokerURL)
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }
    
    return client, nil
}