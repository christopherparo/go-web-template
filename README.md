# go-web-template

A lightweight Go web service template that serves both a JSON API and a static frontend, with Docker and GitHub Actions CI/CD baked in.

## Project Structure

```
.
├── cmd/server/          # Application entrypoint
│   └── main.go
├── internal/server/     # HTTP server, router, and handlers
│   ├── server.go        # Server with graceful shutdown
│   ├── router.go        # chi router + middleware + static file serving
│   ├── handlers.go      # API endpoint handlers
│   └── handlers_test.go # Handler tests
├── web/                 # Embedded frontend assets
│   ├── web.go           # go:embed directive
│   └── static/          # HTML, CSS, JS served at /
├── .github/workflows/
│   └── ci.yml           # Test + build + push to GHCR
├── Dockerfile           # Multi-stage build
├── Makefile             # Dev commands
└── go.mod
```

## Quick Start

### Run locally

```bash
# Default port 8080
make run

# Custom port
PORT=3000 make run
```

Then open [http://localhost:8080](http://localhost:8080). The frontend will fetch `/api/health` and display the result.

### Run with Docker

```bash
# Build and run
make docker-run

# Or manually
docker build -t go-web-template:local .
docker run --rm -p 8080:8080 go-web-template:local
```

### Run tests

```bash
make test
```

## API

| Method | Path          | Description          |
|--------|---------------|----------------------|
| GET    | `/api/health` | Returns `{"status":"ok"}` |
| GET    | `/*`          | Serves static frontend |

## CI/CD Pipeline

The GitHub Actions workflow (`.github/workflows/ci.yml`) runs on every push and pull request:

1. **Test** — Runs `go test -race ./...` on all branches and PRs.
2. **Build & Push** — On pushes to `main` or version tags (`v*`), builds the Docker image and pushes to GitHub Container Registry.

### Image Tags

| Trigger | Tag |
|---------|-----|
| Push to `main` | `ghcr.io/christopherparo/go-web-template:latest` |
| Any push | `ghcr.io/christopherparo/go-web-template:sha-<short>` |
| Tag `v1.2.3` | `ghcr.io/christopherparo/go-web-template:1.2.3`, `:1.2` |

### Pull the image

```bash
docker pull ghcr.io/christopherparo/go-web-template:latest
```

## Makefile Targets

| Target | Description |
|--------|-------------|
| `make build` | Compile Go binary to `bin/server` |
| `make run` | Run locally via `go run` |
| `make test` | Run tests with race detection |
| `make lint` | Run golangci-lint |
| `make docker-build` | Build Docker image |
| `make docker-run` | Build and run Docker container |
| `make clean` | Remove build artifacts |

## Customizing the Frontend

The frontend lives in `web/static/` and is embedded into the Go binary at compile time via `//go:embed`. To swap it:

1. Replace the contents of `web/static/` with your frontend build output.
2. Rebuild the Go binary — the new assets are automatically embedded.

No changes to Go code required.

## Tech Stack

- **Go 1.24** with [chi](https://github.com/go-chi/chi) router
- **Docker** multi-stage build (Alpine runtime)
- **GitHub Actions** for CI/CD
- **GHCR** for container images
