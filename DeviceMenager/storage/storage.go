package storage

import "ConfigApp/model"

type Storage interface {
	CreateDeviceInfo(model.DeviceInfoRequest) (model.DeviceInfo, error)
	CreateSensor(model.SensorRequest) (model.SensorResponse, error)
	CreateOrganization(model.OrganizationDataRequest) (model.OrganizationDataReponse, error)
	CreateUserOrganizationConnection(int, string, string) (model.UserOrganizationConnection, error)
	CreateSlotsForDevice(int, int) error

	UpdateSlot(int, int, int) error

	GetOrganizationsConnectedToUser(string) ([]model.UserOrganization, error)
	GetDeviceInfoByOrgId(int) ([]model.DeviceInfo, error)
	GetDeviceInfoByDeviceId(int) (*model.DeviceInfo, error)
	GetDeviceDataByDeviceId(int) (*model.DeviceData, error)
	GetSlotsByDeviceId(int) ([]model.Slot, error)
}