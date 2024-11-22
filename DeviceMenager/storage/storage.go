package storage

import "ConfigApp/model"

type Storage interface {
	AssignDeviceToOrganization(model.AddDeviceInfo) (*model.DeviceInfo, error)
	CreateDeviceInfo(model.AddDeviceInfo) (model.DeviceInfo, error)
	CreateSensor(model.SensorRequest) (model.SensorResponse, error)
	CreateOrganization(model.OrganizationDataRequest) (model.OrganizationDataReponse, error)
	CreateUserOrganizationConnection(int, string, string) (model.UserOrganizationConnection, error)
	CreateInitialSensorsForDevice(int, int) error

	UpdateSensor(model.SensorUpdate) (*model.SensorUpdate, error)

	GetOrganizationsConnectedToUser(string) ([]model.UserOrganization, error)
	GetDeviceInfoByOrgId(int) ([]model.DeviceInfo, error)
	GetDeviceInfoByDeviceId(int) (*model.DeviceInfo, error)
	GetDeviceInfoByMAC(string) (*model.DeviceInfo, error)
	GetDeviceDataByDeviceId(int) (*model.DeviceData, error)
}