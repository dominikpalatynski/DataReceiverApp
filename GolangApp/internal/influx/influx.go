package influx

import (
	"context"
	"data_receiver/internal/models"
	"log"

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

func (ic *InfluxClient) WriteData(ctx context.Context, snapshot models.Snapshot) error {
    writeAPI := ic.client.WriteAPIBlocking(ic.org, snapshot.BucketName)
    for _, sensor := range snapshot.Sensors {
        p := influxdb2.NewPoint(
            snapshot.DeviceName,
            map[string]string{"sensor": sensor.SensorName},
            map[string]interface{}{sensor.Variable: sensor.Data},
            snapshot.TimeStamp,
        )
        if err := writeAPI.WritePoint(ctx, p); err != nil {
            return err
        }
    }
	log.Printf("Zapisano dane z urządzenia %s", snapshot.DeviceName)
    return nil
}