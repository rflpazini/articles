package main

import (
	jsonv2 "encoding/json/v2"
	"fmt"
)

func main() {
	complexJSON := `{
        "data": {
            "items": [
                {"id": 1, "values": [1.1, 2.2, 3.3, 4.4, 5.5]},
                {"id": 2, "values": [6.6, 7.7, 8.8, 9.9, 10.0]},
                {"id": 3, "values": [11.1, 12.2, 13.3, 14.4, 15.5]}
            ],
            "metadata": {
                "count": 3,
                "sum": 120.0,
                "avg": 40.0
            }
        },
        "config": {
            "precision": 2,
            "format": "decimal"
        }
    }`

	// v2 otimiza especialmente este tipo de estrutura:
	// - Arrays de números (fast path para []float64)
	// - Objetos aninhados com tipos primitivos
	// - Maps com chaves string conhecidas
	var result struct {
		Data struct {
			Items []struct {
				ID     int       `json:"id"`
				Values []float64 `json:"values"` // fast path otimizado
			} `json:"items"`
			Metadata map[string]float64 `json:"metadata"` // fast path otimizado
		} `json:"data"`
		Config map[string]interface{} `json:"config"`
	}

	err := jsonv2.Unmarshal([]byte(complexJSON), &result)
	if err != nil {
		fmt.Printf("Erro: %v\\n", err)
		return
	}

	fmt.Printf("Processados %d items\\n", len(result.Data.Items))
	for _, item := range result.Data.Items {
		fmt.Printf("Item %d tem %d valores\\n", item.ID, len(item.Values))
	}
}
