package storage

import (
	"context"
	"data_viewer/model"
	"fmt"

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

func (ic *InfluxClient) FetchData(queryParams *model.QueryParams) ([]map[string]interface{}, error) {
	queryAPI := ic.client.QueryAPI(ic.org)

	query := fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -1h)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["_field"] == "%s")
	`, queryParams.Bucket, queryParams.Measurement, queryParams.VariableName)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	for result.Next() {
		data = append(data, map[string]interface{}{
			"time":  result.Record().Time(),
			"value": result.Record().Value(),
		})
	}

	return data, nil
}