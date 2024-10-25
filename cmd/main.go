package main

import (
	"context"
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

func (dm *DeviceManager) processMQTTMessage(topic string, value string) error {
    parts := strings.Split(topic, "/")
    if len(parts) != 3 {
        return fmt.Errorf("nieprawidłowy format topicu: %s", topic)
    }
    
    deviceID := parts[1]
    measurement := parts[2]

    if err := dm.ensureBucketExists(deviceID); err != nil {
        return err
    }

    writeAPI := dm.getWriteAPI(deviceID)

    point := influxdb2.NewPoint(
        measurement,
        map[string]string{
            "deviceId": deviceID,
        },
        map[string]interface{}{
            "value": value,
        },
        time.Now(),
    )

    return writeAPI.WritePoint(context.Background(), point)
}

func main() {
    deviceManager := NewDeviceManager(
        "http://influxdb2:8086",
        "mytoken",
        "myorg",
    )

    opts := mqtt.NewClientOptions().
        AddBroker("tcp://mosquitto:1883")

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    topic := "devices/+/measurements"
    if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
        err := deviceManager.processMQTTMessage(msg.Topic(), string(msg.Payload()))
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