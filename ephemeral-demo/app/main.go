package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	Version    = "dev"
	CommitHash = "unknown"
	BuildTime  = "unknown"
)

type HealthResponse struct {
	Status      string `json:"status"`
	ImageRef    string `json:"image_ref,omitempty"`
	ProjectName string `json:"project_name,omitempty"`
	Commit      string `json:"commit,omitempty"`
	Version     string `json:"version,omitempty"`
	BuildTime   string `json:"build_time,omitempty"`
}

func main() {
	sub := os.Getenv("SUBDOMAIN")
	if sub == "" {
		sub = "demo.127.0.0.1.sslip.io"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Preview online at %s\n", sub)
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		health := HealthResponse{
			Status:      "healthy",
			ImageRef:    os.Getenv("IMAGE_REF"),
			ProjectName: os.Getenv("PROJECT_NAME"),
			Commit:      CommitHash,
			Version:     Version,
			BuildTime:   BuildTime,
		}

		if err := json.NewEncoder(w).Encode(health); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
