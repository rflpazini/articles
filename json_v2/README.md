# Go 1.25 – JSON v2 e GreenteaGC

Chegou o Go 1.25 e, finalmente, temos duas mudanças que fazem diferença no dia a dia:

- **JSON v2 experimental (`encoding/json/v2`)** – muito mais rápido, com menos alocações e suporte nativo a streaming.
- **GreenteaGC** – coletor de lixo alternativo que promete reduzir overhead de CPU em workloads intensivos.

Este repositório contém **exemplos práticos e benchmarks** para você testar as novidades.

---

## Pré-requisitos

- Go **1.25+** instalado
- Dados de teste (ex.: issues do GitHub ou [GitHub Archive](https://www.gharchive.org/)) para rodar benchmarks realistas

---

## Como rodar o JSON v2

O pacote ainda é experimental. Você **precisa habilitar via GOEXPERIMENT**:

```bash
GOEXPERIMENT=jsonv2 go run main.go      # para ativar o jsonv2
GOEXPERIMENT=greenteagc go run main.go  # para habilitar o novo gc
```