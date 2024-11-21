package storage

import (
	"ConfigApp/model"
	"context"
	"errors"
	"fmt"
	"strconv"

	supa "github.com/nedpals/supabase-go"
)

type SupabaseStorage struct {
    client *supa.Client
}

func NewSupabaseStorage(client *supa.Client) *SupabaseStorage {
    return &SupabaseStorage{client: client}
}

func (s *SupabaseStorage) CreateDeviceInfo(deviceInfo model.AddDeviceInfo) (model.DeviceInfo, error) {

	ctx := context.Background()
	var results []model.DeviceInfo
	if err := s.client.DB.From("DeviceInfo").Insert(deviceInfo).ExecuteWithContext(ctx, &results); err != nil {
		return model.DeviceInfo{}, err
	}

	return results[0], nil
}

func (s *SupabaseStorage) GetDeviceInfoByOrgId(orgId int) ([]model.DeviceInfo, error) {
	ctx := context.Background()
	var results []model.DeviceInfo
	if err := s.client.DB.From("DeviceInfo").Select("*").Eq("org_id", strconv.Itoa(orgId)).ExecuteWithContext(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *SupabaseStorage) GetDeviceInfoByDeviceId(deviceId int) (*model.DeviceInfo, error) {
	ctx := context.Background()
	var results []model.DeviceInfo
	if err := s.client.DB.From("DeviceInfo").Select("*").Eq("id", strconv.Itoa(deviceId)).ExecuteWithContext(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("deviceInfo not found for the given orgId")
	}

	return &results[0], nil
}

func (s *SupabaseStorage) GetDeviceInfoByMAC(mac string) (*model.DeviceInfo, error) {
    ctx := context.Background()

    var results []model.DeviceInfo
    err := s.client.DB.From("DeviceInfo").
        Select("*").
		Eq("mac_adress", mac).
        ExecuteWithContext(ctx, &results)
    if err != nil {
        return nil, err
    }

    if len(results) == 0 {
        return nil, errors.New("device not found in database")
    }

    return &results[0], nil
}

func (s *SupabaseStorage) GetDeviceDataByDeviceId(deviceId int) (*model.DeviceData, error) {
	ctx := context.Background()
	var results []model.DeviceData
	if err := s.client.DB.From("DeviceInfo").Select("name, Organization(bucket), Sensor(id, variable_name)").Eq("id", strconv.Itoa(deviceId)).ExecuteWithContext(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("deviceInfo not found for the given orgId")
	}

	return &results[0], nil
}

func (s *SupabaseStorage) CreateOrganization(organization model.OrganizationDataRequest) (model.OrganizationDataReponse, error) {

	ctx := context.Background()
	var results []model.OrganizationDataReponse
	if err := s.client.DB.From("Organization").Insert(organization).ExecuteWithContext(ctx, &results); err != nil {
		fmt.Println(err)
		return model.OrganizationDataReponse{}, err
	}

	return results[0], nil
}

func (s *SupabaseStorage) CreateUserOrganizationConnection(orgId int, userId string, role string) (model.UserOrganizationConnection, error) {

	var insertData model.UserOrganizationConnection

	insertData.OrgId = orgId
	insertData.UserId = userId
	insertData.Role = role

	ctx := context.Background()
	var results []model.UserOrganizationConnection
	if err := s.client.DB.From("UserOrganization").Insert(insertData).ExecuteWithContext(ctx, &results); err != nil {
		fmt.Println(err)
		return model.UserOrganizationConnection{}, err
	}

	return results[0], nil
}

func (s *SupabaseStorage) GetOrganizationsConnectedToUser(userId string) ([]model.UserOrganization, error) {
	ctx := context.Background()
	var results []model.UserOrganization
	if err := s.client.DB.From("UserOrganization").Select("role, Organization(id, name)").Eq("user_id", userId).ExecuteWithContext(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *SupabaseStorage) CreateSensor(sensorRequest model.SensorRequest) (model.SensorResponse, error) {
	
	ctx := context.Background()

	insertData := struct {
		DeviceId      int    `json:"device_id" binding:"required" db:"device_id"`
		Variable_name string `json:"variable_name" binding:"required" db:"variable_name"`
		Name 		string `json:"name" binding:"required" db:"name"`
	}{
        Name:        sensorRequest.Name,
        Variable_name: sensorRequest.Variable_name,
        DeviceId:    sensorRequest.DeviceId,
    }


	var results []model.SensorResponse
	if err := s.client.DB.From("Sensor").Insert(insertData).ExecuteWithContext(ctx, &results); err != nil {
		fmt.Println(err)
		return model.SensorResponse{}, err
	}

	

	return results[0], nil
}

func (s *SupabaseStorage) CreateInitialSensorsForDevice(deviceId int, initialSensors int) error {
	sensors := make([]model.SensorInsert, initialSensors)

	for i := 0; i < initialSensors; i++ {
		sensors[i] = model.SensorInsert{
			DeviceId:   deviceId,
			SlotNumber: i + 1,
		}
	}

	if err := s.client.DB.From("Sensor").Insert(sensors).Execute(nil); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *SupabaseStorage) UpdateSensor(sensor model.SensorUpdate) (*model.SensorUpdate, error) {

	updateData := map[string]interface{}{
		"variable_name": sensor.VariableName,
		"name": sensor.Name,
	}

	ctx := context.Background()
	var results []model.SensorUpdate
	if err := s.client.DB.From("Sensor").Update(updateData).Eq("id", strconv.Itoa(sensor.Id)).ExecuteWithContext(ctx, &results); err != nil {
		return nil, err
	}

	if condition := len(results) == 0; condition {
		return nil, errors.New("sensor not found")
		
	}

	return &results[0], nil
}

func (s *SupabaseStorage) AssignDeviceToOrganization(deviceInfo model.AddDeviceInfo) (*model.DeviceInfo, error) {

	updateData := map[string]interface{}{
		"org_id": deviceInfo.OrgId,
		"name": deviceInfo.Name,
		"interval": deviceInfo.Interval,
	}

	ctx := context.Background()
	var results []model.DeviceInfo
	if err := s.client.DB.From("DeviceInfo").Update(updateData).Eq("mac_adress", deviceInfo.MAC).ExecuteWithContext(ctx, &results); err != nil {
		return nil, err
	}

	if condition := len(results) == 0; condition {
		return nil, errors.New("device not found")
		
	}

	return &results[0], nil
}
