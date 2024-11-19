package model

type QueryParams struct {
	Bucket       string `json:"bucket"`
	Measurement  string `json:"measurement"`
	VariableName string `json:"variable_name"`
}

type DataPoint struct {
	Time  interface{} `json:"time"`
	Value interface{} `json:"value"`
}