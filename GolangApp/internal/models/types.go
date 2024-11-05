package models

import "time"

type Snapshot struct {
    TimeStamp  time.Time   `json:"timeStamp"`
    BucketName string      `json:"bucketName"`
    DeviceName string      `json:"deviceName"`
    Sensors    []SensorData `json:"sensors"`
}

type SensorData struct {
    Data       int    `json:"data"`
    Variable   string `json:"variable"`
    SensorName string `json:"sensorName"`
}
