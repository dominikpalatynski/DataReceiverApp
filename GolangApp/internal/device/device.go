package device

import (
	"context"
	"data_receiver/internal/influx"
	"data_receiver/internal/models"
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DeviceManager struct {
    influxClient *influx.InfluxClient
}

func NewDeviceManager(influxURL, token, org string) *DeviceManager {
    influxClient := influx.NewClient(influxURL, token, org)
    return &DeviceManager{influxClient: influxClient}
}

func (dm *DeviceManager) ProcessMQTTMessage(client mqtt.Client, msg mqtt.Message) {
    var snapshot models.Snapshot
    if err := json.Unmarshal(msg.Payload(), &snapshot); err != nil {
        log.Printf("Błąd parsowania JSON: %v", err)
    }

	deviceData, ok := fetchDeviceData(snapshot.DeviceId)
	if ok != nil {
		log.Printf("Błąd pobierania danych urządzenia: %v", ok)
		return
	}

	point := preparePoint(*deviceData, snapshot)

	log.Printf("Pobrano dane urządzenia: %v", deviceData)

    err :=  dm.influxClient.WriteData(context.Background(), point)
	if err != nil {
		log.Printf("Błąd zapisu do InfluxDB: %v", err)
	}
}