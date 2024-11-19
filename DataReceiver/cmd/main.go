// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"strings"
// 	"sync"
// 	"time"

// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// 	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
// 	"github.com/influxdata/influxdb-client-go/v2/api"
// )

// type DeviceManager struct {
//     influxClient influxdb2.Client
//     org          string
//     writeAPIs    map[string]api.WriteAPIBlocking
//     mu           sync.RWMutex
// }

// type Point struct {
//     Name string `json:"name"`
//     Meta map[string]string `json:"meta"`
//     Data map[string]interface{} `json:"data"`
//     TimeStamp time.Time `json:"timestamp"`
// }

// type SensorData struct {
// 	Data       int    `json:"data"`
// 	Variable   string `json:"variable"`
// 	SensorName string `json:"sensorName"`
// }

// type Snapshot struct {
// 	TimeStamp time.Time     `json:"timeStamp"`
// 	BucketName      string        `json:"bucketName"`
// 	DeviceName      string        `json:"deviceName"`
// 	Sensors   []SensorData  `json:"sensors"`
// }

// func NewDeviceManager(influxURL, token, org string) *DeviceManager {
//     client := influxdb2.NewClient(influxURL, token)
//     return &DeviceManager{
//         influxClient: client,
//         org:          org,
//         writeAPIs:    make(map[string]api.WriteAPIBlocking),
//     }
// }

// func (dm *DeviceManager) getWriteAPI(deviceID string) api.WriteAPIBlocking {
//     dm.mu.RLock()
//     writeAPI, exists := dm.writeAPIs[deviceID]
//     dm.mu.RUnlock()

//     if !exists {
//         dm.mu.Lock()
//         writeAPI, exists = dm.writeAPIs[deviceID]
//         if !exists {
//             writeAPI = dm.influxClient.WriteAPIBlocking(dm.org, deviceID)
//             dm.writeAPIs[deviceID] = writeAPI
//         }
//         dm.mu.Unlock()
//     }

//     return writeAPI
// }

// func (dm *DeviceManager) ensureBucketExists(deviceID string) error {
//     queryAPI := dm.influxClient.QueryAPI(dm.org)

//     query := fmt.Sprintf(`buckets() |> filter(fn: (r) => r.name == "%s")`, deviceID)
//     result, err := queryAPI.Query(context.Background(), query)
//     if err != nil {
//         return fmt.Errorf("błąd sprawdzania bucketu: %v", err)
//     }

//     bucketExists := false
//     for result.Next() {
//         bucketExists = true
//         break
//     }

//     if !bucketExists {
//         return fmt.Errorf("nie można utworzyć bucketu dla urządzenia %s: %v", deviceID, err)
//     }

//     return nil
// }

// func (dm *DeviceManager) processMQTTMessage(topic string, data []byte) error {
//     parts := strings.Split(topic, "/")
//     if len(parts) != 3 {
//         return fmt.Errorf("nieprawidłowy format topicu: %s", topic)
//     }

//     var snapshot Snapshot

//     if err := json.Unmarshal(data, &snapshot); err != nil{
//         log.Println("Błąd dekodowania json", err)
//         return err
//     }

//     deviceID := parts[1]

//     if err := dm.ensureBucketExists(snapshot.BucketName); err != nil {
//         return err
//     }

//     writeAPI := dm.getWriteAPI(deviceID)

//     for _, sensor := range snapshot.Sensors {
//         mappedPoint := influxdb2.NewPoint(
//             snapshot.DeviceName,
//             map[string]string{
//                 sensor.SensorName: sensor.SensorName,
//             },
//             map[string]interface{}{
//                 sensor.Variable: sensor.Data,
//             },
//             snapshot.TimeStamp,
//         )

//         if err := writeAPI.WritePoint(context.Background(), mappedPoint); err != nil {
//             return err
//         }
//     }

//     return nil
// }

package main

import (
	"data_receiver/internal/broker"
	"data_receiver/internal/config"
	"data_receiver/internal/device"
	"log"
)

func main() {

    config, err := config.LoadConfig()

    if err != nil {
        log.Fatalf("Fatal error during config load %v", err)
    }

    deviceManager := device.NewDeviceManager(config.Database.Url, config.Database.Token, config.Database.Org)

    mqttClient, err := broker.NewClient(config.Broker.BrokerUrl)
    if err != nil {
        log.Fatalf("Błąd konfiguracji MQTT: %v", err)
    } else {
        log.Println("Połączono z brokerem MQTT")
    }

    if token := mqttClient.Subscribe(config.Broker.TopicPattern, 0, deviceManager.ProcessMQTTMessage); token.Wait() && token.Error() != nil {
        log.Fatalf("Błąd subskrybowania: %v", token.Error())
    } else {
        log.Println("Subskrybowano temat devices/+/measurements")
    }

    select {}
}