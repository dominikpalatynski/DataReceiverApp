package server

import (
	"ConfigApp/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getDeviceInfoFromAPI(c *gin.Context, post *model.AddDeviceInfo) error{
	if err := c.ShouldBindJSON(post); err != nil {
		return err
	}
	return nil
}

func getSensorInfoFromApi(c *gin.Context, post *model.SensorRequest) error{
	if err := c.ShouldBindJSON(post); err != nil {
		return err
	}
	return nil
}

func getOrganizationDataFromAPI(c *gin.Context, post *model.OrganizationDataRequest) error{
	if err := c.ShouldBindJSON(post); err != nil {
		return err
	}
	return nil
}

func (s APIServer) getDeviceData(deviceId string) (*model.DeviceData, error) {

	deviceKey := fmt.Sprintf("device:%s", deviceId)

	deviceDataFromCache, ok := s.cache.GetDeviceDataFromCache(deviceKey)
	if ok == nil {
		return deviceDataFromCache, nil
	}
	fmt.Print(ok)

	deviceIdToInt, err := strconv.Atoi(deviceId)
	if err != nil {
		fmt.Print("Can not convert deviceId to int")
		return nil, err
	}

	deviceData, err := s.storage.GetDeviceDataByDeviceId(deviceIdToInt)

	if err != nil {
		return nil, err
	}

	s.cache.SetDeviceDataToCache(*deviceData, deviceKey)

	return deviceData, nil
}

func (s APIServer) getDeviceStateCredentials(deviceId string) (*model.DeviceStateCredentials, error) {

	deviceKey := fmt.Sprintf("deviceStateCreds:%s", deviceId)

	deviceStateCredentialsFromCache, ok := s.cache.GetDeviceStateCredentialsToCache(deviceKey)
	if ok == nil {
		return deviceStateCredentialsFromCache, nil
	}
	fmt.Print(ok)

	deviceIdToInt, err := strconv.Atoi(deviceId)
	if err != nil {
		fmt.Print("Can not convert deviceId to int")
		return nil, err
	}

	deviceStateCredentials, err := s.storage.GetDeviceStateCredentials(deviceIdToInt)

	if err != nil {
		return nil, err
	}

	s.cache.SetDeviceStateCredentialsToCache(*deviceStateCredentials, deviceKey)

	return deviceStateCredentials, nil
}