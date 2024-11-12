package storage

import "data_viewer/model"

type Storage interface {
	FetchData(*model.QueryParams) ([]map[string]interface{}, error)
}