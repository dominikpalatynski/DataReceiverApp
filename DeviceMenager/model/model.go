package model

type DeviceInfo struct {
	Id       int    `json:"id" db:"id"`
	OrgId    int    `json:"org_id" binding:"required" db:"org_id"`
	Interval int    `json:"interval" binding:"required" db:"interval"`
	Name     string `json:"name" db:"name"`
	MAC      string `json:"mac_adress" binding:"required" db:"mac_adress"`
}

type AddDeviceInfo struct {
	OrgId    int    `json:"org_id" binding:"required" db:"org_id"`
	Interval int    `json:"interval" binding:"required" db:"interval"`
	Name     string `json:"name" db:"name"`
	MAC      string `json:"mac_adress" binding:"required" db:"mac_adress"`
}

type DeviceInitDataRequest struct {
	MAC   string `json:"mac"`
	Token string `json:"token"`
}
type DeviceInitDataResponse struct {
	Id string `json:"id"`
}

type Organization struct {
	BucketName string `json:"bucket"`
}

type SensorData struct {
	Id           int    `json:"id"`
	VariableName string `json:"variable_name"`
}

type SensorUpdate struct {
	Id           int    `json:"id"`
	VariableName string `json:"variable_name"`
	Name         string `json:"name"`
}

type DeviceData struct {
	Organization Organization `json:"organization"`
	Name         string       `json:"name"`
	Sensor       []SensorData `json:"sensor"`
}

type OrganizationDataRequest struct {
	Name   string `json:"name" binding:"required" db:"name"`
	Bucket string `json:"bucket" binding:"required" db:"bucket"`
}

type OrganizationDataReponse struct {
	ID     int    `json:"id" binding:"required" db:"id"`
	Name   string `json:"name" binding:"required" db:"name"`
	Bucket string `json:"bucket" binding:"required" db:"bucket"`
}

type UserOrganizationConnection struct {
	OrgId  int    `json:"org_id" binding:"required" db:"org_id"`
	UserId string `json:"user_id" binding:"required" db:"user_id"`
	Role   string `json:"role" binding:"required" db:"role"`
}

type OrganizationName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserOrganization struct {
	Organization OrganizationName `json:"organization"`
	Role         string           `json:"role"`
}

type SensorRequest struct {
	DeviceId      int    `json:"device_id" binding:"required" db:"device_id"`
	Variable_name string `json:"variable_name" binding:"required" db:"variable_name"`
	Name          string `json:"name" binding:"required" db:"name"`
	Slot          int    `json:"slot" binding:"required" db:"slot"`
}

type SensorResponse struct {
	ID            int    `json:"id" binding:"required" db:"id"`
	DeviceId      int    `json:"device_id" binding:"required" db:"device_id"`
	Variable_name string `json:"variable_name" binding:"required" db:"variable_name"`
	Name          string `json:"name" binding:"required" db:"name"`
}

type SensorInsert struct {
	SlotNumber int `json:"slot_number"`
	DeviceId   int `json:"device_id"`
}
