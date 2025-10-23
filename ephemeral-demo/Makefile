REGISTRY ?= local
IMAGE ?= preview-app
TAG ?= demo
IMAGE_REF := $(REGISTRY)/$(IMAGE):$(TAG)

SUBDOMAIN ?= demo.127.0.0.1.sslip.io
PROJECT_NAME ?= demo

# Capture git commit if available
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
VERSION ?= $(TAG)

.PHONY: traefik-up traefik-down traefik-logs build preview-up open teardown rebake status

traefik-up: ## Start Traefik with File provider (ECI compatible)
	@echo "🚀 Starting Traefik..."
	docker compose -f deploy/compose.traefik.yml up -d
	@echo "Traefik running at http://localhost:80"
	@echo "Dashboard available at http://localhost:8080"

traefik-down: ## Stop Traefik
	docker compose -f deploy/compose.traefik.yml down

traefik-logs: ## Show Traefik logs
	docker compose -f deploy/compose.traefik.yml logs -f

build:
	docker buildx bake --file docker-bake.hcl \
	  --set app.tags=$(IMAGE_REF) \
	  --set app.cache-to=type=inline \
	  --set app.args.VERSION=$(VERSION) \
	  --set app.args.COMMIT=$(COMMIT)

rebake: ## Build a new tag quickly to show cached rebuild
	$(MAKE) build TAG=$(shell date +%H%M%S)

preview-up:
	IMAGE_REF=$(IMAGE_REF) SUBDOMAIN=$(SUBDOMAIN) PROJECT_NAME=$(PROJECT_NAME) \
	docker compose -f deploy/compose.preview.yml up -d

open:
	@echo "📍 Direct access (port): http://localhost:8090"
	@echo "🌐 Access via Traefik: http://demo.127.0.0.1.sslip.io"
	@curl -sS "http://localhost:8090/health" || true

status: ## Show container status
	@echo "🔍 Container status:"
	@docker ps --filter "name=traefik" --filter "name=$(PROJECT_NAME)" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

teardown:
	IMAGE_REF=$(IMAGE_REF) SUBDOMAIN=$(SUBDOMAIN) PROJECT_NAME=$(PROJECT_NAME) \
	docker compose -f deploy/compose.preview.yml down -v || true
