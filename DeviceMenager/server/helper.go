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

func getOrganizationDataFromAPI(c *gin.Context, post *model.OrganizationData) error{
	if err := c.ShouldBindJSON(post); err != nil {
		return err
	}
	return nil
}