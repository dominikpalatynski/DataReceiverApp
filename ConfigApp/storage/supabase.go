package storage

import (
	"ConfigApp/model"
	"context"
	"strconv"

	supa "github.com/nedpals/supabase-go"
)

type SupabaseStorage struct {
    client *supa.Client
}

func NewSupabaseStorage(url, apiKey string) *SupabaseStorage {
    client := supa.CreateClient(url, apiKey)
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