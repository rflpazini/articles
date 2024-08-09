# url_shortener

`url_shortener` is a Go-based application that shortens URLs using a Base62 encoding algorithm and stores the shortened URLs in a Redis database. The application is designed to be simple, fast, and easy to deploy with Docker Compose.



## Project Structure

```bash
url_shortener
├── cmd
│   └── main.go                   # Entry point of the application
├── compose.yaml                  # Docker Compose configuration
├── go.mod                        # Go module file
├── go.sum                        # Go dependencies checksum file
├── internal
│   ├── server
│   │   └── server.go             # Server setup and initialization
│   └── shortener
│       ├── model.go              # Data models for URL information
│       ├── repository.go         # Redis repository for storing URLs
│       └── service.go            # Business logic for URL shortening
└── pkg
    ├── api
    │   └── shortener
    │       └── handler.go        # HTTP handlers for API endpoints
    └── utils
        └── base62
            └── hash.go           # Base62 encoding implementation
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