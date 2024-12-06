package storage

import (
	"DeviceStateHandler/internal/model"
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Storage interface {
    SetState(model.DeviceState, model.DeviceStateCredentials) error
    GetDeviceStates(queryParams *model.QueryParams, since string) ([]model.DataPoint, error)
}

type InfluxClient struct {
	client influxdb2.Client
	org string
}

func NewClient(url, token, org string) *InfluxClient {
	client := influxdb2.NewClient(url, token)
	return &InfluxClient{
		client: client,
		org: org,
	}
}


func (ic *InfluxClient) SetState(deviceState model.DeviceState, deviceCreds model.DeviceStateCredentials) error {
    writeAPI := ic.client.WriteAPIBlocking(ic.org, deviceCreds.Organization.BucketName)

	p := influxdb2.NewPoint(
		deviceCreds.Name,
		map[string]string{},
		map[string]interface{}{
			"state": deviceState.DeviceState,
		},
		time.Now(),
	)

	ctx := context.Background()
	if err := writeAPI.WritePoint(ctx, p); err != nil {
		return err
	}

	log.Printf("Zapisano dane z urządzenia %s", deviceCreds.Name)
    return nil
}

func (ic *InfluxClient) GetDeviceStates(queryParams *model.QueryParams, since string) ([]model.DataPoint, error) {
	queryAPI := ic.client.QueryAPI(ic.org)

	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: %s)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["_field"] == "state")
	`, queryParams.Bucket, since, queryParams.DeviceName)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var data []model.DataPoint
	for result.Next() {
		data = append(data, model.DataPoint{
			Time:  result.Record().Time(),
			Value: result.Record().Value(),
		})
	}

	return data, nil
}