# Traefik with File Provider

This configuration uses Traefik's **File Provider** instead of the Docker Provider, making it compatible with Docker Desktop's **Enhanced Container Isolation (ECI)**.

## Why File Provider?

ECI blocks the Docker socket mount (`/var/run/docker.sock`), preventing Traefik from auto-discovering containers. With the File Provider, we define routes manually in YAML files.

## Structure

```
traefik/
├── traefik.yml        # Static configuration
└── dynamic/           # Dynamic configurations (routes)
    └── app.yml        # Routes for the application
```

## How to Use

1. **Start Traefik:**
   ```bash
   make traefik-up
   ```

2. **Build and start your application:**
   ```bash
   make build
   make preview-up
   ```

3. **Access:**
   - App via Traefik: http://demo.127.0.0.1.sslip.io
   - App direct: http://localhost:8090
   - Traefik Dashboard: http://localhost:8081

4. **Check status:**
   ```bash
   make status
   ```

## Adding New Routes

To add a new service, create a file in `dynamic/`:

```yaml
# dynamic/my-service.yml
http:
  routers:
    my-service:
      rule: "Host(`my-service.127.0.0.1.sslip.io`)"
      entryPoints: ["web"]
      service: my-service

  services:
    my-service:
      loadBalancer:
        servers:
          - url: "http://my-service:8080"
```

Traefik automatically detects new files and updates the routes!
