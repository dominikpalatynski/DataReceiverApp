package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func getEnvOrDefault(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

type Point struct {
    Name string `json:"name"`
    Meta map[string]string `json:"meta"`
    Data map[string]interface{} `json:"data"`
    TimeStamp string `json:"timestamp"`
}

type Snapshot struct {
    TimeStamp string `json:"timestamp"`
    DeviceId string      `json:"deviceID"`
    Sensors    []SensorData `json:"sensors"`
}

type SensorData struct {
    Data       int    `json:"data"`
    SensorID   string `json:"sensorID"`
}


func main() {
    brokerHost := getEnvOrDefault("MQTT_HOST", "localhost")
    brokerPort := getEnvOrDefault("MQTT_PORT", "1883")
    clientID := getEnvOrDefault("CLIENT_ID", "go_publisher")
	interval := getEnvOrDefault("INTERVAL", "5")
    topic := getEnvOrDefault("MQTT_TOPIC", "devices/mybucket/measurements")

    brokerURL := fmt.Sprintf("tcp://%s:%s", brokerHost, brokerPort)
    
    opts := mqtt.NewClientOptions().
        AddBroker(brokerURL).
        SetClientID(clientID)

    opts.SetWill("device/37/status", `{"state":"disconnected"}`, 1, true)

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Połączono z brokerem MQTT")
		token := c.Publish("device/37/status", 1, true, `{"state":"connected"}`)
		token.Wait()
		fmt.Println("Wysłano status: connected")
	}

    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("Błąd połączenia: %s", token.Error())
    }

    log.Printf("Połączono z brokerem MQTT: %s", brokerURL)
    log.Printf("Publikowanie na topic: %s", topic)
	counter := 0
    temperature := 20
    for {
        point := Snapshot{
            DeviceId: "37",
            Sensors: []SensorData{
            {Data: temperature, SensorID: "19"},},
            TimeStamp: time.Now().Format(time.RFC3339),
        }
        temperature += 10
        jsonData, err := json.Marshal(point)
    
        fmt.Println(string(jsonData))

        if err != nil {
            log.Fatal(err)
        }
        token := client.Publish(topic, 0, false, jsonData)
        token.Wait()
		counter++
        log.Println("Wysłano wiadomość")
		num, err := strconv.Atoi(interval)
		if err != nil {
			fmt.Println("Błąd konwersji:", err)
		} else {
			fmt.Println("Wartość int:", num)
		}
        time.Sleep(time.Duration(num) * time.Second)
    }
}