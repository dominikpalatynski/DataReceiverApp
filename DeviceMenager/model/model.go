package model

type DeviceInfo struct {
	Id       int    `json:"id" binding:"required" db:"id"`
	OrgId    int    `json:"org_id" binding:"required" db:"org_id"`
	Interval int    `json:"interval" binding:"required" db:"interval"`
	Name     string `json:"name" binding:"required" db:"name"`
	Bucket   string `json:"bucket" binding:"required" db:"bucket"`
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
	Sensor       []SensorData `json:"sensor"`
}

type OrganizationData struct {
	Name   string `json:"name" binding:"required" db:"name"`
	Bucket string `json:"bucket" binding:"required" db:"bucket"`
}