package storage

import "ConfigApp/model"

type Storage interface {
	CreateDeviceInfo(model.DeviceInfo) (model.DeviceInfo, error)
	GetDeviceInfoByOrgId(int) ([]model.DeviceInfo, error)
	GetDeviceInfoByDeviceId(int) (*model.DeviceInfo, error)
	GetDeviceDataByDeviceId(int) (*model.DeviceData, error)
}