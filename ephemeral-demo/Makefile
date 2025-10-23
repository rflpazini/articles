REGISTRY ?= local
IMAGE ?= preview-app
TAG ?= demo
IMAGE_REF := $(REGISTRY)/$(IMAGE):$(TAG)

SUBDOMAIN ?= demo.127.0.0.1.sslip.io
PROJECT_NAME ?= demo

# Capture git commit if available
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
VERSION ?= $(TAG)

.PHONY: traefik-up build preview-up open teardown rebake

traefik-up: ## Note: Traefik requires Docker socket access (blocked by org policy). Using direct port mapping instead.
	@echo "⚠️  Traefik disabled due to Enhanced Container Isolation policy"
	@echo "Using direct port mapping on localhost:8080 instead"

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
	@echo "URL: http://localhost:8090"
	@curl -sS "http://localhost:8090/health" || true

teardown:
	IMAGE_REF=$(IMAGE_REF) SUBDOMAIN=$(SUBDOMAIN) PROJECT_NAME=$(PROJECT_NAME) \
	docker compose -f deploy/compose.preview.yml down -v || true
