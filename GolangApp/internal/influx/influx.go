package influx

import (
	"context"
	"data_receiver/internal/models"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

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

func (ic *InfluxClient) WriteData(ctx context.Context, point *models.Point) error {
    writeAPI := ic.client.WriteAPIBlocking(ic.org, point.Bucket)

	p := influxdb2.NewPoint(
		point.Name,
		point.Meta,
		point.Data,
		time.Now(),
	)
	if err := writeAPI.WritePoint(ctx, p); err != nil {
		return err
	}

	log.Printf("Zapisano dane z urządzenia %s", point.Name)
    return nil
}