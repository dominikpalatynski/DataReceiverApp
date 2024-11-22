package device

import (
	"data_receiver/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func fetchDeviceData(deviceId string) (*models.DeviceData, error) {
	url := fmt.Sprintf("http://device_menager:5000/deviceData/%s", deviceId)

	resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("błąd podczas wysyłania zapytania: %v", err)
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("błąd w odpowiedzi: %s", resp.Status)
    }

    var deviceData models.DeviceData
    err = json.NewDecoder(resp.Body).Decode(&deviceData)
    if err != nil {
        return nil, fmt.Errorf("błąd podczas dekodowania odpowiedzi: %v", err)
    }

    return &deviceData, nil

}

func preparePoint(deviceData models.DeviceData, snapshot models.Snapshot) *models.Point {
	var point models.Point
	point.Data = make(map[string]interface{})
	sensorMap := make(map[int]int)

	for _, sensor := range snapshot.Sensors {
		sensorId, _ := strconv.Atoi(sensor.SensorID)
		sensorMap[sensorId] = sensor.Data
	}
	for _, sensorData := range deviceData.Sensors {
		point.Data[sensorData.VariableName] = sensorMap[sensorData.Id]
	}

	point.Bucket = deviceData.Organization.BucketName
	point.Name = deviceData.Name

	parsedTime, err := time.Parse(time.RFC3339, snapshot.TimeStamp)
	if err != nil {
		fmt.Printf("Błąd parsowania czasu: %v", err)
		return nil
	}
	point.TimeStamp = parsedTime

	return &point
}