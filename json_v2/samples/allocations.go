package main

import (
	jsonv2 "encoding/json/v2"
	"fmt"
	"runtime"
)

type LargeStruct struct {
	Users    []User            `json:"users"`
	Metadata map[string]string `json:"metadata"`
	Settings []Setting         `json:"settings"`
}

type User struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Tags    []string `json:"tags"`
	Profile Profile  `json:"profile"`
}

type Profile struct {
	Bio      string            `json:"bio"`
	Avatar   string            `json:"avatar"`
	Socials  map[string]string `json:"socials"`
	Metadata interface{}       `json:"metadata"`
}

type Setting struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func main() {
	data := []byte(`{
        "users": [
            {
                "id": 1,
                "name": "João Silva",
                "email": "joao@example.com",
                "tags": ["admin", "premium"],
                "profile": {
                    "bio": "Desenvolvedor Go há 5 anos",
                    "avatar": "https://example.com/avatar1.jpg",
                    "socials": {"github": "joaosilva", "twitter": "@joao"},
                    "metadata": {"level": 5, "badges": ["expert", "mentor"]}
                }
            }
        ],
        "metadata": {
            "version": "2.1",
            "timestamp": "2024-12-01T10:00:00Z"
        },
        "settings": [
            {"key": "theme", "value": "dark"},
            {"key": "notifications", "value": true}
        ]
    }`)

	var result LargeStruct

	// Measure allocations
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	jsonv2.Unmarshal(data, &result)

	runtime.GC()
	runtime.ReadMemStats(&m2)

	fmt.Printf("Alocações: %d bytes\n", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Printf("Objetos criados: %d\n", m2.Mallocs-m1.Mallocs)
}
