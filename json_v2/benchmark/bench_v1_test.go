package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json" // v1
	"os"
	"testing"

	"github.com/rflpazini/jsonv2/internal"
)

type Event = map[string]any

func benchmarkNDJSON(b *testing.B, path string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f, err := os.Open(path)
		if err != nil {
			b.Fatal(err)
		}
		gz, err := gzip.NewReader(f)
		if err != nil {
			b.Fatal(err)
		}
		s := internal.NewNDJSONScanner(bufio.NewReader(gz))

		var count int
		for s.Scan() {
			line := s.Bytes()
			var ev Event
			if err := json.Unmarshal(line, &ev); err != nil {
				b.Fatal(err)
			}
			if _, ok := ev["type"]; ok {
				count++
			}
		}
		if err := s.Err(); err != nil {
			b.Fatal(err)
		}
		gz.Close()
		f.Close()
		b.SetBytes(int64(count)) // rough: eventos processados
	}
}

func Benchmark_V1_GHArchive_UnmarshalMap(b *testing.B) {
	benchmarkNDJSON(b, "../data/2025-07-01-12.json.gz")
}
