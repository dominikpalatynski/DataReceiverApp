package storage

import "ConfigApp/model"

type Storage interface {
	CreateDeviceInfo(model.DeviceInfo) (model.DeviceInfo, error)
}