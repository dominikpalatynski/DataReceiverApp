package server

import (
	"ConfigApp/model"

	"github.com/gin-gonic/gin"
)

func getDeviceInfoFromAPI(c *gin.Context, post *model.DeviceInfo) error{
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