# Resume API (`resume-api`)

~~Production-grade~~ Go microservice for querying and retrieving resume profiles, built with feature-sliced domain architecture and strict dependency boundaries on top of the [`github.com/nawaphonOHM/whatever`](https://github.com/nawaphonOHM/whatever) framework.

---

## Architecture Overview

`resume-api` follows the **Feature-Slice Domain Pattern** for high cohesion, low coupling, and rapid developer velocity:

```
.
├── .env.example
├── .gitignore
├── .golangci.yml
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── cmd/
│   └── resume-api/
│       └── main.go          # Application lifecycle, dependency wiring, server bootstrap
└── internal/
    └── resume/
        ├── model.go         # Domain models (Resume, ContactInfo) and fallback records
        ├── sections.go      # Experience and Education section entities
        ├── repository.go    # Data persistence interface and memory fallback
        ├── mongo.go         # MongoDB persistence driver integration
        ├── service.go       # Domain business logic and error handling
        ├── handler.go       # Declarative REST handlers and RFC 9457 error mappings
        └── registration.go  # Route registration blueprint for /v1/resumes
```

### Core Technologies & Principles
- **Go 1.27+ Standard Toolchain**
- **Whatever Framework (`pkg/rest`, `pkg/mongodb`)**: Declarative HTTP routing, RFC 9457 Problem Details envelopes, managed MongoDB connections, and Kubernetes-ready lifecycle probes.
- **Strict Dependency Guard**: Zero unvetted third-party dependencies. Static analysis allows exclusively Go standard library and `github.com/nawaphonOHM/whatever`.
- **Dual-Mode Persistence**: Seamless MongoDB persistence with automatic fallback to built-in standalone records when database host is unconfigured for offline local development.

---

## Configuration

The application is configured through environment variables prefixed with `OHM9996_`. A baseline template is provided in `.env.example`.

### Server Settings (`OHM9996_SERVER_*`)

| Variable | Default | Description |
|---|---|---|
| `OHM9996_SERVER_HOST` | `0.0.0.0` | Server bind host address |
| `OHM9996_SERVER_PORT` | `8080` | Server HTTP port |
| `OHM9996_GIN_MODE` | `release` | Web engine mode (`debug` or `release`) |
| `OHM9996_SERVER_SHUTDOWN_TIMEOUT` | `10s` | Graceful shutdown draining period |
| `OHM9996_SERVER_ENABLE_ACCESS_LOG` | `true` | Enable standard request access logging |

### MongoDB Settings (`OHM9996_MONGODB_*`)

| Variable | Default | Description |
|---|---|---|
| `OHM9996_MONGODB_HOST` | *(empty)* | MongoDB server host (standalone fallback used if unset) |
| `OHM9996_MONGODB_PORT` | `27017` | MongoDB connection port |
| `OHM9996_MONGODB_DATABASE` | `resume_db` | Target database name |
| `OHM9996_MONGODB_USERNAME` | *(empty)* | Authentication username |
| `OHM9996_MONGODB_PASSWORD` | *(empty)* | Authentication password |
| `OHM9996_MONGODB_AUTH_SOURCE` | `admin` | Authentication database source |

### OpenTelemetry Settings (`OHM9996_OTEL_*`)

| Variable | Default | Description |
|---|---|---|
| `OHM9996_OTEL_ENABLED` | `true` | Enable or disable distributed tracing (`true` or `false`) |
| `OHM9996_OTEL_SERVICE_NAME` | `resume-api` | Logical service name attached to exported traces |
| `OHM9996_OTEL_EXPORTER_OTLP_ENDPOINT` | `localhost:4317` | Collector receiver endpoint (e.g. `localhost:4317` for gRPC or `localhost:4318` for HTTP) |
| `OHM9996_OTEL_EXPORTER_OTLP_PROTOCOL` | `grpc` | Exporter transport protocol (`grpc`, `http`, `http/protobuf`, `http/json`) |
| `OHM9996_OTEL_INSECURE` | `true` | Disable TLS for collector connection (`true` or `false`) |
| `OHM9996_OTEL_SAMPLE_RATE` | `1.0` | Trace sampling probability ratio between `0.0` (0%) and `1.0` (100%) |
| `OHM9996_OTEL_SKIP_PATHS` | `/health,/ready` | Comma-separated list of route paths excluded from trace generation |

---

## API Endpoints

### 1. Health & Readiness Probes
- `GET /health`: Liveness probe managed by framework runtime.
- `GET /ready`: Readiness probe indicating service traffic readiness.

### 2. Resume Resource Endpoints
- `GET /v1/resumes`: Retrieve a collection of all resume profiles.
    - **Success Response (200 OK)**:
      ```json
      {
        "success": true,
        "data": [
          {
            "id": "1",
            "name": "Nawaphon",
            "title": "Senior Software Engineer",
            "summary": "Experienced backend engineer...",
            "contact": {
              "email": "nawaphon@example.com",
              "phone": "+66 81 234 5678",
              "location": "Bangkok, Thailand",
              "website": "https://github.com/nawaphonOHM"
            },
            "skills": ["Go", "MongoDB", "Distributed Systems"],
            "experience": [...],
            "education": [...]
          }
        ],
        "timestamp": "2026-09-28T22:00:00Z"
      }
      ```
- `GET /v1/resumes/:id`: Retrieve a specific resume profile by ID.
    - **Success Response (200 OK)**: Wrapped resume object inside standard JSON envelope.
    - **Error Response (404 Not Found)**: Standard RFC 9457 Problem Details (`application/problem+json`):
      ```json
      {
        "type": "about:blank",
        "title": "Not Found",
        "status": 404,
        "detail": "resume profile not found",
        "code": "RESUME_NOT_FOUND"
      }
      ```

---

## Static Analysis & Linter Policy

Static analysis is enforced via `.golangci.yml` incorporating the complete ruleset:
- **`depguard`**: Strictly restricts imported packages to `$gostd`, `github.com/nawaphonOHM/whatever`, and `github.com/nawaphonOHM/resume-api`.
- **`revive`**: Enforces strict complexity limits (max cyclomatic complexity 3, function length <= 10 statements, file length <= 100 lines).
- **`govet`**, **`errcheck`**, **`ineffassign`**, **`staticcheck`**, **`unused`**, **`misspell`**.

---

## Developer Workflows (`Makefile`)

| Command | Action |
|---|---|
| `make build` | Compiles binary to `bin/resume-api` |
| `make run` | Starts the service locally |
| `make test` | Executes unit tests with race detection (`go test -race ./...`) |
| `make test-coverage` | Runs tests and generates `coverage.html` |
| `make vet` | Runs `go vet ./...` |
| `make lint` | Executes `golangci-lint run` across the codebase |
| `make tidy` | Tidies and synchronizes `go.mod` and `go.sum` |
| `make clean` | Cleans up binaries and coverage profiles |
