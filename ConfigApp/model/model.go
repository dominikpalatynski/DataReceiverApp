package model

type DeviceInfo struct {
	Id       int    `json:"id" binding:"required" db:"id"`
	OrgId    int    `json:"org_id" binding:"required" db:"org_id"`
	Interval int    `json:"interval" binding:"required" db:"interval"`
	Name     string `json:"name" binding:"required" db:"name"`
	Bucket   string `json:"bucket" binding:"required" db:"bucket"`
}