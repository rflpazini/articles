# Ephemeral Environments Lightning Demo

Production-like preview for every PR with Docker Buildx, Docker Bake, Compose and a tiny router.

## Quickstart

Prereqs
- Docker Engine or Docker Desktop with Buildx enabled
- Internet to pull images on first run
- A terminal in the repo root

Warm up images before the talk
```
docker pull traefik:v3.1 golang:1.22 gcr.io/distroless/base-debian12:nonroot
```

Run the 6 minute sequence

```
make traefik-up
make build TAG=live-001
make preview-up TAG=live-001 SUBDOMAIN=demo.127.0.0.1.sslip.io PROJECT_NAME=demo
make open
# edit app/main.go and change the message line
make rebake
make preview-up TAG=$(date +%H%M%S)
make open
make teardown
```

## How it works

- Buildx and Bake create a tagged image with inline cache
- Compose deploys a small stack that Traefik routes by hostname
- sslip.io resolves demo.127.0.0.1.sslip.io to 127.0.0.1 so no DNS change is needed

## Notes

Use a VPS with a public IP if you want remote attendees to open the URL on their devices. Replace 127.0.0.1 with your public IP inside the SUBDOMAIN value.
