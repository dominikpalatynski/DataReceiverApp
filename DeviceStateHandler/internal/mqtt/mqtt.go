package mqtt

import (
	"DeviceStateHandler/internal/model"
	"DeviceStateHandler/internal/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)
type MQTTClient struct {
	client mqtt.Client
	cacheClient storage.Cache
	storageClient storage.Storage
	deviceManagerURL string
}

func NewMqttClient(brokerURL string, cacheClient storage.Cache, storageClient storage.Storage, deviceManagerURL string) (*MQTTClient, error) {
	client, err := newClient(brokerURL)

	if err != nil {
		return nil, err
	}

	mqttClient := &MQTTClient{
		client: client,
		cacheClient: cacheClient,
		storageClient: storageClient,
		deviceManagerURL: deviceManagerURL,
	}

	return mqttClient, nil
}

func (mq *MQTTClient) Subcribe(topicPattern string) error{
	if token := mq.client.Subscribe("device/+/status", 0, mq.processIncomingState); token.Wait() && token.Error() != nil {
        return token.Error()
    }

	log.Println("Subscribed on topic: device/+/status")

	return nil
}

func (mq *MQTTClient) processIncomingState(client mqtt.Client, msg mqtt.Message) {
    state :=  new(model.DeviceState)
    if err := json.Unmarshal(msg.Payload(), state); err != nil {
        log.Printf("Błąd parsowania JSON: %v", err)
		return
    }

	parts := strings.Split(msg.Topic(), "/")
	if len(parts) != 3 {
		log.Printf("Unsuported Topic pattern: %s", msg.Topic())
		return
	}

	deviceId := parts[1]
	deviceKey := fmt.Sprintf("deviceState:%s", deviceId)

	if err := mq.cacheClient.SetDeviceState(*state, deviceKey); err != nil {
		log.Printf("Can not save state in cache, %s", err)
	}

	url := fmt.Sprintf(mq.deviceManagerURL + deviceId)

	resp, err := http.Get(url)
    if err != nil {
        log.Printf("błąd podczas wysyłania zapytania: %v", err)
		return
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
        log.Printf("błąd w odpowiedzi: %s", resp.Status)
		return
    }

	var deviceData model.DeviceStateCredentials
    err = json.NewDecoder(resp.Body).Decode(&deviceData)
    if err != nil {
        log.Printf("błąd podczas dekodowania odpowiedzi: %v", err)
		return
    }

	if err := mq.storageClient.SetState(*state, deviceData); err != nil {
		log.Printf("błąd podczas zapiswyania statusu: %v", err)
	}

	log.Printf("Zapisano pomyślnie %v", deviceData.Name)
}

func newClient(brokerURL string) (mqtt.Client, error) {
    opts := mqtt.NewClientOptions().AddBroker(brokerURL)
    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }
    
    return client, nil
}