package model

type DeviceState struct {
	DeviceState string `json:"state"`
}

type DeviceStateCredentials struct {
	Organization Organization `json:"organization"`
	Name         string       `json:"name"`
}

type Organization struct {
	BucketName string `json:"bucket"`
}

type QueryParams struct {
	Bucket     string `json:"bucket"`
	DeviceName string `json:"device_name"`
}

type DataPoint struct {
	Time  interface{} `json:"time"`
	Value interface{} `json:"value"`
}