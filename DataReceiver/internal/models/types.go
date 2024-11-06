package models

type Snapshot struct {
	TimeStamp string   `json:"timeStamp"`
	DeviceId  string   `json:"deviceID"`
	Sensors   []Sensor `json:"sensors"`
}

type Sensor struct {
	Data     int    `json:"data"`
	SensorID string `json:"sensorID"`
}

type Organization struct {
	BucketName string `json:"bucket"`
}

type SensorData struct {
	Id           int    `json:"id"`
	VariableName string `json:"variable_name"`
}

type DeviceData struct {
	Organization Organization `json:"organization"`
	Name         string       `json:"name"`
	Sensors      []SensorData `json:"sensor"`
}

type Point struct {
	Bucket string
	Name   string
	Meta   map[string]string
	Data   map[string]interface{}
}