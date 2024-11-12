package server

import (
	"data_viewer/model"

	"github.com/gin-gonic/gin"
)

func getQueryParamsFromApi(c *gin.Context, post *model.QueryParams) error{
	if err := c.ShouldBindJSON(post); err != nil {
		return err
	}
	return nil
}