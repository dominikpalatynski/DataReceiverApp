package influx

import (
	"context"
	"data_receiver/internal/models"
	"fmt"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxClient struct {
	client influxdb2.Client
	org string
}

func NewClient(url, token, org string) (*InfluxClient, error) {
	client := influxdb2.NewClient(url, token)
    ctx := context.Background()
    health, err := client.Health(ctx)
    if err != nil {
        return nil, err
    }

    if health.Status != "pass" {
        return nil, fmt.Errorf("InfluxDB health check failed: %s", health.Message)
    }

	return &InfluxClient{
		client: client,
		org: org,
	}, nil
}

func (ic *InfluxClient) WriteData(ctx context.Context, point *models.Point) error {
    writeAPI := ic.client.WriteAPIBlocking(ic.org, point.Bucket)

	p := influxdb2.NewPoint(
		point.Name,
		point.Meta,
		point.Data,
		point.TimeStamp,
	)
	if err := writeAPI.WritePoint(ctx, p); err != nil {
		return err
	}

	log.Printf("Zapisano dane z urządzenia %s", point.Name)
    return nil
}