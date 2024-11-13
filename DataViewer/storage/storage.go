package storage

import (
	"data_viewer/model"
	"time"
)

type Storage interface {
	FetchData(*model.QueryParams, time.Time) ([]model.DataPoint, error)
}