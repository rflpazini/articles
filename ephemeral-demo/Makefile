REGISTRY ?= local
IMAGE ?= preview-app

TAG ?= live-001
IMAGE_REF := $(REGISTRY)/$(IMAGE):$(TAG)

# Default subdomain for demo (can be overridden per deployment)
SUBDOMAIN ?= demo.127.0.0.1.sslip.io
ENV_NAME ?= demo

# Capture git commit if available
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
VERSION ?= $(TAG)

.PHONY: traefik-up traefik-down traefik-logs build generate-traefik-config cleanup-preview preview-up open teardown rebake status clean

clean: ## Kill all containers and clean up
	@docker ps -aq --filter "network=edge" | xargs -r docker rm -f 2>/dev/null || true
	@docker compose -f deploy/compose.traefik.yml down 2>/dev/null || true
	@rm -f deploy/traefik/dynamic/*.yml
	@echo "All containers cleaned up"

traefik-up: ## Start Traefik with File provider (ECI compatible)
	docker compose -f deploy/compose.traefik.yml up -d

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

rebake: ## Quick rebuild using cache
	$(MAKE) build TAG=$(shell date +%H%M%S)

generate-traefik-config: ## Generate Traefik routing configuration for the environment
	@mkdir -p deploy/traefik/dynamic
	@echo "# Auto-generated route for $(ENV_NAME)" > deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "http:" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "  routers:" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "    $(ENV_NAME):" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "      rule: \"Host(\`$(SUBDOMAIN)\`)\"" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "      entryPoints: [\"web\"]" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "      service: $(ENV_NAME)" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "  services:" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "    $(ENV_NAME):" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "      loadBalancer:" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "        servers:" >> deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "          - url: \"http://$(ENV_NAME)-app:8080\"" >> deploy/traefik/dynamic/$(ENV_NAME).yml

cleanup-preview: ## Stop and remove existing preview containers
	@docker stop $(ENV_NAME)-app > /dev/null 2>&1 || true
	@docker rm $(ENV_NAME)-app > /dev/null 2>&1 || true

preview-up: generate-traefik-config cleanup-preview ## Deploy preview environment
	@echo "Deploying $(ENV_NAME)..."
	@echo "  Image: $(IMAGE_REF)"
	@echo "  URL: http://$(SUBDOMAIN)"
	IMAGE_REF=$(IMAGE_REF) SUBDOMAIN=$(SUBDOMAIN) ENV_NAME=$(ENV_NAME) \
	docker compose -f deploy/compose.preview.yml up -d
	@sleep 2
	@echo "Deployment complete ✔"

open: ## curls api endpoints
	@curl -sS "http://$(SUBDOMAIN)/health" | python3 -m json.tool 2>/dev/null || curl -sS "http://$(SUBDOMAIN)/health"
	@echo ""
	@curl -sS "http://$(SUBDOMAIN)/"

status: ## Show container status
	@echo "🔍 Container status:"
	@docker ps --filter "name=traefik" --filter "name=$(ENV_NAME)" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

teardown: ## Remove a specific preview environment
	IMAGE_REF=$(IMAGE_REF) SUBDOMAIN=$(SUBDOMAIN) ENV_NAME=$(ENV_NAME) \
	docker compose -f deploy/compose.preview.yml down -v || true
	@rm -f deploy/traefik/dynamic/$(ENV_NAME).yml
	@echo "Environment $(ENV_NAME) removed"
