package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type DeviceManager struct {
    influxClient influxdb2.Client
    org          string
    writeAPIs    map[string]api.WriteAPIBlocking
    mu           sync.RWMutex
}

type Point struct {
    Name string `json:"name"`
    Meta map[string]string `json:"meta"`
    Data map[string]interface{} `json:"data"`
    TimeStamp time.Time `json:"timestamp"`
}

func NewDeviceManager(influxURL, token, org string) *DeviceManager {
    client := influxdb2.NewClient(influxURL, token)
    return &DeviceManager{
        influxClient: client,
        org:          org,
        writeAPIs:    make(map[string]api.WriteAPIBlocking),
    }
}

func (dm *DeviceManager) getWriteAPI(deviceID string) api.WriteAPIBlocking {
    dm.mu.RLock()
    writeAPI, exists := dm.writeAPIs[deviceID]
    dm.mu.RUnlock()

    if !exists {
        dm.mu.Lock()
        writeAPI, exists = dm.writeAPIs[deviceID]
        if !exists {
            writeAPI = dm.influxClient.WriteAPIBlocking(dm.org, deviceID)
            dm.writeAPIs[deviceID] = writeAPI
        }
        dm.mu.Unlock()
    }

    return writeAPI
}

func (dm *DeviceManager) ensureBucketExists(deviceID string) error {
    queryAPI := dm.influxClient.QueryAPI(dm.org)
    
    query := fmt.Sprintf(`buckets() |> filter(fn: (r) => r.name == "%s")`, deviceID)
    result, err := queryAPI.Query(context.Background(), query)
    if err != nil {
        return fmt.Errorf("błąd sprawdzania bucketu: %v", err)
    }

    bucketExists := false
    for result.Next() {
        bucketExists = true
        break
    }

    if !bucketExists {
        return fmt.Errorf("nie można utworzyć bucketu dla urządzenia %s: %v", deviceID, err)
    }

    return nil
}

func (dm *DeviceManager) processMQTTMessage(topic string, data []byte) error {
    parts := strings.Split(topic, "/")
    if len(parts) != 3 {
        return fmt.Errorf("nieprawidłowy format topicu: %s", topic)
    }
    
    deviceID := parts[1]

    if err := dm.ensureBucketExists(deviceID); err != nil {
        return err
    }

    writeAPI := dm.getWriteAPI(deviceID)

    var point Point
    err := json.Unmarshal(data, &point)

    if(err != nil){
        log.Println("Błąd dekodowania json", err)
        return err
    }

    mappedPoint := influxdb2.NewPoint(
        point.Name,
        point.Meta,
        point.Data,
        point.TimeStamp,
    )

    return writeAPI.WritePoint(context.Background(), mappedPoint)
}

func main() {
    deviceManager := NewDeviceManager(
        "http://localhost:8086",
        "mytoken",
        "myorg",
    )

    opts := mqtt.NewClientOptions().
        AddBroker("tcp://localhost:1883")

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    topic := "devices/+/measurements"
    if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
        err := deviceManager.processMQTTMessage(msg.Topic(), msg.Payload())
        if err != nil {
            log.Printf("Błąd przetwarzania wiadomości: %v", err)
        } else{
            log.Printf("Przetworzono wiadomość: %s", msg.Payload())
        }
    }); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    select {}
}