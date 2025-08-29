# Features Documentation

This document outlines the core features implemented in this Go modular monolith template.  
These features serve as foundational APIs for all projects based on this template and provide essential monitoring capabilities for Kubernetes deployments.

---

## 1. Ping API ✅ **IMPLEMENTED**

### Purpose
Provides a lightweight endpoint to verify the service is alive and responding.

### Specification
- **Endpoint:** `GET /ping`
- **Response:** 
  - **Status Code:** `200 OK`
  - **Content-Type:** `application/json`
  - **Body:** `{"message": "PONG"}`

### Implementation Details
- **Module:** `internal/ping`
- **Handler:** `internal/ping/infrastructure/http/ping_handler.go`
- **Service:** `internal/ping/application/ping_service.go`
- **Domain:** `internal/ping/domain/ping.go`
- **Port:** `internal/ping/ports/ping_port.go`

### Usage
- Used as a simple health check and readiness probe in Kubernetes.
- Available at: `http://localhost:8080/ping`

### Notes
- Always returns success unless the service is fully down.
- Follows clean architecture principles with proper separation of concerns.

---

## 2. Health API ✅ **IMPLEMENTED**

### Purpose
A comprehensive health check that validates all critical external dependencies are accessible and functioning.

### Specification
- **Endpoint:** `GET /health`
- **Response:**
  - **Status Code:** `200 OK` (if all checks pass) or `503 Service Unavailable` (if any check fails)
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "status": "healthy|unhealthy",
      "checks": {
        "database": {
          "status": "up|down",
          "response_time_ms": 15
        },
        "redis": {
          "status": "up|down",
          "response_time_ms": 8
        }
      },
      "timestamp": "2024-01-15T10:30:00Z"
    }
    ```

### Implementation Details
- **Module:** `internal/probes` (health sub-module)
- **Handler:** `internal/probes/presentation/http/health_handler.go`
- **Service:** `internal/probes/application/query/health_query.go`
- **Domain:** `internal/probes/domain/health.go`
- **Ports:** `internal/probes/ports/health_port.go`
- **Infrastructure:** `internal/probes/infrastructure/database_checker.go`, `internal/probes/infrastructure/redis_checker.go`

### Usage
- Used as a readiness probe in Kubernetes to ensure the service is fully operational before receiving traffic.
- Available at: `http://localhost:8080/health`

### Notes
- Checks connectivity with database and Redis.  
- Measures response times and provides detailed status for debugging.  
- Includes timeout mechanisms to prevent hanging.
- Follows clean architecture principles with proper separation of concerns.  

---

## 3. Liveness API

### Purpose
Detects if the service is in a healthy state and capable of processing requests. Helps Kubernetes determine if a pod should be restarted.

### Specification
- **Endpoint:** `GET /liveness`
- **Response:**
  - **Status Code:** `200 OK` (service is alive) or `503 Service Unavailable` (service should be restarted)
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "status": "alive|dead",
      "uptime_seconds": 3600,
      "timestamp": "2024-01-15T10:30:00Z"
    }
    ```

### Usage
- Used as a Kubernetes liveness probe to determine when a pod should be restarted.  

### Notes
- Focuses only on the service’s internal state (not external dependencies).  
- Should be lightweight and fast.  

---

## 4. API Documentation (Swagger) ✅ **IMPLEMENTED**

### Purpose
Provides interactive API documentation using Swagger/OpenAPI specification for all endpoints in the service.

### Specification
- **Swagger UI Endpoint:** `GET /swagger/`
- **OpenAPI Spec Endpoint:** `GET /openapi.yaml`
- **Response:** Interactive Swagger UI interface and OpenAPI specification files

### Implementation Details
- **Module:** `internal/docs`
- **Handler:** `internal/docs/infrastructure/http/docs_handler.go`
- **Service:** `internal/docs/application/docs_service.go`
- **Domain:** `internal/docs/domain/service.go`
- **Port:** `internal/docs/ports/docs_port.go`
- **Adapter:** `internal/docs/infrastructure/docs_adapter.go`
- **OpenAPI Spec:** `api/openapi.yaml`
- **Generated Docs:** `docs/` directory (auto-generated)

### Features
- Interactive Swagger UI for testing API endpoints
- Auto-generated documentation from code annotations
- OpenAPI 3.0 specification format
- Comprehensive API metadata including contact info and license
- Support for request/response examples

### Usage
- Access Swagger UI at: `http://localhost:8080/swagger/`
- Download OpenAPI spec from: `http://localhost:8080/openapi.yaml`
- API documentation is automatically updated when code annotations change

### Notes
- Uses `swaggo/swag` for generating documentation from Go annotations
- Follows clean architecture with proper separation of concerns
- Documentation is served both as interactive UI and downloadable specifications

---

## 5. Implementation Guidelines for Features

### Error Handling
- Graceful degradation when external services are unavailable.  
- Proper HTTP status codes for different failure scenarios.  
- Detailed error messages in development, sanitized in production.  

### Performance
- Health checks should complete within 5 seconds.  
- Connection pooling should be used to avoid overhead.  
- Results may be cached briefly to prevent overloading dependencies.  

### Security
- No sensitive information should be exposed in responses.  
- Consider rate limiting for health endpoints.  
- Ensure endpoints do not become attack vectors.  

### Testing
- Unit tests for feature logic.  
- Integration tests with real dependencies where applicable.  
- Mocks should be used for isolated testing.  

---

## 6. Future Enhancements

### Potential Extensions
- Metrics collection and exposure (Prometheus format).  
- Custom health checks for business-specific dependencies.  
- Configurable health check intervals and thresholds.  
- Integration with distributed tracing systems.  
- Support for graceful shutdown procedures.  

### Monitoring Integration
- These APIs provide the foundation for comprehensive monitoring.  
- Can be extended to integrate with systems like Prometheus, Grafana, or DataDog.  
- Support for custom metrics and alerting based on health check results.  

---

**Note:** This document is focused on features only.  
For implementation details, please refer to `architecture.md` and `tech-stack.md`.
