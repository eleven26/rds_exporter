package exporter

import "strings"

type PerformanceKeys struct {
	PerformanceKey []PerformanceKeyItem
}

type PerformanceValue struct {
	PerformanceValue []PerformanceValueItem
}

type PerformanceValueItem struct {
	Value string
	Date  string
}

type Body struct {
	PerformanceKeys PerformanceKeys
	EndTime         string
	StartTime       string
	DBInstanceId    string
	Engine          string
}

type PerformanceKeyItem struct {
	Key         string
	Unit        string
	Values      PerformanceValue
	ValueFormat string
}

type ApiResponse struct {
	Body Body `json:"body"`
}

func (k PerformanceKeyItem) performanceMetrics() map[string]string {
	metrics := strings.Split(k.ValueFormat, "&")
	values := strings.Split(k.Values.PerformanceValue[0].Value, "&")

	result := make(map[string]string, len(metrics))

	for k, metric := range metrics {
		result[metric] = values[k]
	}

	return result
}

type Metrics struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (api ApiResponse) Metrics() []Metrics {
	var result []Metrics

	for _, item := range api.Body.PerformanceKeys.PerformanceKey {
		// 这两个在我的测试实例上返回空的值，先跳过
		if item.Key == "MySQL_ReplicationThread" || item.Key == "MySQL_ReplicationDelay" {
			continue
		}

		metrics := item.performanceMetrics()

		for k, v := range metrics {
			result = append(result, Metrics{
				Key:   k,
				Value: v,
			})
		}
	}

	return result
}

func (api ApiResponse) Mappings() map[string]string {
	metrics := api.Metrics()
	result := make(map[string]string, len(metrics))

	for _, metric := range metrics {
		result[metric.Key] = metric.Value
	}

	return result
}
