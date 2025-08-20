package main

import (
	jsonv2 "encoding/json/v2"
	"fmt"
	"strings"
	"time"
)

func main() {
	timeSeriesData := `{
		"metric": "cpu_usage",
		"timestamps": [1609459200, 1609459260, 1609459320, 1609459380, 1609459440],
		"values": [45.2, 67.8, 23.1, 89.5, 12.7, 56.3, 78.9, 34.6, 91.2, 18.4],
		"labels": ["server1", "server2", "server3", "server4", "server5"]
	}`

	// Repete o JSON para simular um dataset maior
	// Cada item separado por vírgula
	repeticoes := 1000
	bigData := strings.Repeat(timeSeriesData+",", repeticoes-1) + timeSeriesData
	bigJSON := "[" + bigData + "]"

	type TimeSeriesPoint struct {
		Metric     string    `json:"metric"`
		Timestamps []int64   `json:"timestamps"`
		Values     []float64 `json:"values"`
		Labels     []string  `json:"labels"`
	}

	start := time.Now()

	var results []TimeSeriesPoint
	err := jsonv2.Unmarshal([]byte(bigJSON), &results)

	parseTime := time.Since(start)

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	totalValues := 0
	for _, result := range results {
		totalValues += len(result.Values)
	}

	fmt.Printf("Parsed %d time series points com %d valores em %v\n",
		len(results), totalValues, parseTime)
	fmt.Printf("Throughput: %.0f valores/segundo\n",
		float64(totalValues)/parseTime.Seconds())
}
