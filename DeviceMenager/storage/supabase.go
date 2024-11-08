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

func (s *SupabaseStorage) CreateDeviceInfo(deviceInfo model.DeviceInfo) (model.DeviceInfo, error) {

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

func (s *SupabaseStorage) CreateOrganization(organization model.OrganizationData) (model.OrganizationData, error) {

	ctx := context.Background()
	var results []model.OrganizationData
	if err := s.client.DB.From("Organization").Insert(organization).ExecuteWithContext(ctx, &results); err != nil {
		fmt.Println(err)
		return model.OrganizationData{}, err
	}

	return results[0], nil
}
