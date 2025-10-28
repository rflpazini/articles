# [DOCKER] - Ship Faster and Test Smarter with Ephemeral Environments

Use this flow on your laptop first. In the event you can switch the IP to a VPS later.

## Prerequisites

- Docker Engine or Docker Desktop with Buildx enabled
- Internet to pull images on first run
- A terminal in the repo root

## Quickstart

1. Warm up images

```bash
docker pull traefik:v3.1 golang:1.22 gcr.io/distroless/base-debian12:nonroot
```

2. Start the router

```bash
make traefik-up
```

3. Build with Buildx and Bake

```bash
make build TAG=live-001
```

4. Bring up the preview

```bash
make preview-up TAG=live-001 SUBDOMAIN=demo.127.0.0.1.sslip.io PROJECT_NAME=demo
```

5. Smoke test

```bash
make open
```

6. Edit the app and rebuild fast

Open `app/main.go` and change the string to show a visible diff. Then:

```bash
make rebake
make preview-up TAG=$(date +%H%M%S)
make open
```

7. Teardown

```bash
make teardown
```

## How it works

Ephemeral environments are created on developer demand for specific tasks such as running tests, previewing features, or staging deployments. Once the task is completed, they are disposed of. Key characteristics:

- **Temporary nature** - They are designed to exist only for a short duration
- **On-demand creation** - Built when needed, not pre-provisioned
- **Clean disposal** - Removed completely after use

Technical flow:

- Buildx and Bake produce a tagged image with inline cache and SBOM
- Compose deploys the stack with Traefik labels
- Traefik watches Docker and routes by host rule
- sslip.io resolves to 127.0.0.1 (zero DNS configuration)
- One tag per iteration allows safe roll forward and rollback

## Files included

- `app/main.go` - simple Go server with health check
- `app/Dockerfile` - distroless runtime
- `docker-bake.hcl` - Buildx Bake target with inline cache and SBOM
- `deploy/compose.traefik.yml` - router stack
- `deploy/compose.preview.yml` - preview stack with Traefik labels
- `Makefile` - all commands you need on stage
- `README.md` - quickstart and runbook
- `slides.md` - 4 concise slides to paste into the shared deck

## Rehearsal checklist

- Terminal in large font
- Pre pull images and test every command once
- Keep a 60 second screen recording as plan B
- If the venue network is flaky use the 127.0.0.1 sslip.io URL in your own browser
- If you use a VPS replace 127.0.0.1 with the public IP inside `SUBDOMAIN`
