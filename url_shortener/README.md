# url_shortener

`url_shortener` is a Go-based application that shortens URLs using a Base62 encoding algorithm and stores the shortened URLs in a Redis database. The application is designed to be simple, fast, and easy to deploy with Docker Compose.



## Project Structure

```bash
url_shortener
├── cmd
│ └── main.go                    # Entry point of the application
├── internal
│ ├── monitoring
│ │ └── prometheus.go            # Prometheus setup and monitoring
│ ├── server
│ │ └── server.go                # Server setup and initialization
│ └── shortener
│     ├── model.go               # Data models for URL information
│     ├── repository.go          # Redis repository for storing URLs
│     ├── service.go             # Business logic for URL shortening
│     └── service_test.go        # Unit tests for the shortener service
├── pkg
│ ├── handler
│ │ └── shortener
│ │     ├── handler.go                  # HTTP handlers for API endpoints
│ │     └── handler_integration_test.go # Integration tests for handlers
│ └── utils
│     └── base62
│         ├── hash.go              # Base62 encoding implementation
│         └── hash_test.go         # Unit tests for Base62 encoding
├── Dockerfile                     # Dockerfile for building the application
├── Dockerfile.alpine              # Dockerfile for building a smaller image with Alpine
├── Dockerfile.golang              # Dockerfile for building with Go SDK
├── README.md                      # Project documentation
├── compose.yml                    # Docker Compose configuration
├── go.mod                         # Go module file
├── go.sum                         # Go dependencies checksum file
└── prometheus.yml                 # Prometheus configuration file
```

## Features
- URL Shortening: Accepts long URLs and shortens them using a Base62 encoding algorithm.
- Redis Storage: Stores the shortened URLs and their corresponding original URLs in a Redis database.
- RESTful API: Provides HTTP endpoints for URL shortening and retrieval.


## Getting Started
### Prerequisites
- Go 1.18 or higher
- Docker and Docker Compose

### Installation

1. Clone the repository:
```bash 
git clone https://github.com/rflpazini/url_shortener.git
cd url_shortener
```

2. Install Go dependencies:
```bash 
go mod tidy
```
3. Start the application using Docker Compose:
```bash 
docker-compose up --build 
```

### Configuration
The application is configured using environment variables. You can set these variables in the `compose.yaml` file or directly in your shell environment.

### Monitoring with Prometheus and Grafana
The `url_shortener` application includes monitoring capabilities using Prometheus and Grafana. These tools allow you to track the performance and health of your application in real-time.

##### Prometheus
Prometheus is configured to scrape metrics from the application at the `/metrics` endpoint. The metrics include:

Total HTTP Requests: Number of requests received by the application, categorized by path and method.
Request Duration: Duration of HTTP requests in seconds.

##### Grafana
Grafana is used to visualize the metrics collected by Prometheus. It provides a powerful and flexible dashboard for monitoring the application.
##### Setting Up Monitoring
Prometheus and Grafana are included in the `compose.yml` configuration. When you start the application with Docker Compose, these services will be automatically started and configured.

To access Grafana:
1. Open your browser and go to http://localhost:3000.
2. The default username and password are both admin.
3. You can then add Prometheus as a data source and create custom dashboards to monitor the application.

### Running tests

- To run unit tests only:
```bash
go test ./...
```

- To run integration tests, you should pass the tag `integration` during the build
```bash
go test ./... -tags=integration
```

### API Endpoints

- POST /v1/shortener: Shorten a new URL.
 
_Request_
```json
{
     "url": "https://hub.docker.com/_/golang"
}

```
_Response_
```json
{
  "url": "https://hub.docker.com/_/golang",
  "short": "9h1sH1",
  "created_at": "2024-08-09T17:42:23.942767Z"
}
```
- GET /v1/shortener?url={shortened_url}: Retrieve the original URL using a shortened URL.

_Response (Redirects to the original URL)_
```
HTTP/1.1 302 Found
Location: https://hub.docker.com/_/golang
```

- GET /v1/shortener: Retrieves all URLs that was stored

_Response_
```json
[
     {
          "url": "https://dev.to/rflpazini",
          "short": "SuWSPc",
          "created_at": "2024-08-09T17:49:32.013842Z",
          "updated_at": "2024-08-09T17:49:33.375284Z"
     },
     {
          "url": "https://hub.docker.com/_/golang",
          "short": "9h1sH1",
          "created_at": "2024-08-09T17:42:23.942767Z",
          "updated_at": "2024-08-09T17:48:54.041113Z"
     }
]
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request for review.

## License
This project is licensed under the MIT License - see the [LICENSE](https://rflpazini.mit-license.org/) file for details.