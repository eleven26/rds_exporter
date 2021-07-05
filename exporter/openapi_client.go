package exporter

import (
	"encoding/json"
	"fmt"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	rds20140815 "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func createClient(accessKeyId *string, accessKeySecret *string) (result *rds20140815.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}

	// 访问的域名
	config.Endpoint = tea.String("rds.aliyuncs.com")
	result = &rds20140815.Client{}
	result, _err = rds20140815.NewClient(config)

	return result, _err
}

func formatTime(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02dZ",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(),
	)
}

func FetchLastMinuteDBInstancePerformance(config Config) (response *MetricsResponse, err error) {
	var apiResponse ApiResponse

	client, err := createClient(tea.String(config.AccessKeyId), tea.String(config.AccessKeySecret))
	if err != nil {
		return nil, err
	}

	// 获取过去一分钟的性能数据，接口需要使用 UTC 时间
	currentTime := time.Now().UTC()
	startTime := formatTime(currentTime.Add(-time.Minute))
	endTime := formatTime(currentTime)

	describeDBInstancePerformanceRequest := &rds20140815.DescribeDBInstancePerformanceRequest{
		DBInstanceId: tea.String(config.InstanceId),
		Key:          tea.String(PerformanceKeysString()),
		StartTime:    tea.String(startTime),
		EndTime:      tea.String(endTime),
	}

	result, err := client.DescribeDBInstancePerformance(describeDBInstancePerformanceRequest)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(result.String()), &apiResponse)

	metricsResponse := MetricsResponse{
		apiResponse: apiResponse,
	}

	return &metricsResponse, nil
}
