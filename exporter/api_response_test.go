package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func apiResponse() ApiResponse {
	var apiResponse ApiResponse

	apiResponseBytes, _ := ioutil.ReadFile("../tests/api_response.json")
	json.Unmarshal(apiResponseBytes, &apiResponse)

	return apiResponse
}

func TestApiResponse(t *testing.T) {
	apiResponse := apiResponse()

	if len(apiResponse.Body.PerformanceKeys.PerformanceKey) != 19 {
		t.Errorf("Api response parse fails.\n")
	}
}

func findInMetrics(key string, metrics []Metrics) bool {
	for _, v := range metrics {
		if key == v.Key {
			return true
		}
	}

	return false
}

func TestApiResponseMetrics(t *testing.T) {
	apiResponse := apiResponse()

	metrics := apiResponse.Metrics()
	if !findInMetrics("com_delete", metrics) {
		t.Errorf("Parse metrics fails.\n")
	}
}
