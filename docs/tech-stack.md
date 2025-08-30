# Tech Stack & Implementation Guidelines

This document defines the official technology stack and implementation guidelines for this project.  
All contributors **must follow these rules** to ensure consistency, maintainability, and performance.

---

## 0. Go Version Requirements

### Go Language
- **Version:** Go 1.24.4 or higher  
- **Rationale:** Required for GoFiber compatibility and latest language features  
- **Guidelines:**  
  - Always use the latest stable Go version when possible  
  - Update go.mod file to reflect minimum required version  
  - CI/CD pipelines must use the specified minimum version or higher  

---

## 1. Databases & Persistence

### PostgreSQL
- **Library:** [`pgx`](https://github.com/jackc/pgx) (with connection pooling).  
- **Migration Tool:** [`golang-migrate`](https://github.com/golang-migrate/migrate) for database schema management.  
- **Usage:**  
  - Default database for all persistent data.  
  - Repository implementations in `/internal/module-x/infrastructure` must use `pgx`.  
  - SQL queries should be written explicitly (avoid ORMs).  
  - All schema changes must go through versioned migration files.  
- **Guidelines:**  
  - Use connection pooling (`pgxpool`) for high concurrency.  
  - Transactions must be managed explicitly at the application/service layer.  
  - Migration scripts must be stored in `/scripts/migrations/` following golang-migrate naming convention.  
  - Never run migrations automatically in production; always use controlled deployment process.  
  - Development environments can auto-run migrations via docker-compose for convenience.  

---

## 2. Caching & In-Memory Storage

### Redis
- **Library:** [`go-redis`](https://github.com/redis/go-redis).  
- **Usage:**  
  - Temporary and non-critical data (e.g., OTP codes, short-lived sessions, rate limiting).  
  - Never store critical system-of-record data.  
- **Guidelines:**  
  - Redis logic should reside in `/internal/module-x/infrastructure`.  
  - Define clear TTLs for all keys.  
  - Avoid storing large payloads (keep Redis usage lightweight).  

---

## 3. API Layer

### HTTP Framework
- **Library:** [`GoFiber`](https://github.com/gofiber/fiber).  
- **Usage:**  
  - All HTTP request handling must be implemented using GoFiber.  
  - Controllers/handlers live in `/internal/module-x/infrastructure/http/`.  
- **Guidelines:**  
  - JSON is the default serialization format.  
  - Middlewares (auth, logging, tracing, etc.) should be configured in `/platform/http`.  
  - Avoid mixing business logic in handlers; delegate to application layer.  

---

## 4. Logging

- **Library:** [`zerolog`](https://github.com/rs/zerolog).  
- **Usage:**  
  - Global logger initialized in `/platform/logger`.  
  - Structured JSON logging must be used.  
  - Log levels: `debug`, `info`, `warn`, `error`, `fatal`.  
- **Guidelines:**  
  - No `fmt.Println` or raw `log` usage.  
  - Every module logs via injected logger dependency.  

---

## 5. Configuration

- **Library:** [`viper`](https://github.com/spf13/viper).  
- **Usage:**  
  - Centralized config in `/platform/config`.  
  - Supports `.yaml` + environment variable overrides.  
- **Guidelines:**  
  - Never hardcode credentials or secrets.  
  - Use environment variables for sensitive values.  

---

## 6. Testing

- **Unit Testing:** Go’s built-in `testing` package.  
- **Mocks:** [`testify/mock`](https://github.com/stretchr/testify).  
- **Integration/E2E Tests:** Located in `/test/`, may spin up Postgres + Redis using Docker.  

---

## 7. Dependency Injection

- **Library:** [`wire`](https://github.com/google/wire) (**MANDATORY**).  
- **Usage:**  
  - All dependencies **MUST** be injected via constructors using Wire.  
  - Manual dependency injection is **PROHIBITED**.  
  - Final wiring done only in `/cmd/app/main.go` using Wire's generated code.  
  - Each module **MUST** define its providers in a `wire.go` file.  
- **Guidelines:**  
  - Use `//go:build wireinject` build tags only for Wire injector functions.  
  - Provider functions should follow the pattern: `NewModuleName(deps...) *ModuleName`.  
  - Wire provider sets should be organized by module for clarity.  
  - Never instantiate dependencies manually in production code.  

---

## 8. Other Libraries

- **Validation:** [`go-playground/validator`](https://github.com/go-playground/validator).  
- **UUIDs:** [`google/uuid`](https://pkg.go.dev/github.com/google/uuid).  
- **Time Handling:** [`github.com/jinzhu/now`](https://github.com/jinzhu/now) for parsing helpers, standard `time` for core logic.  

---

## 9. Rules & Enforcement

- Do not introduce alternative libraries for the same concern (e.g., no `gorm` or `database/sql` if `pgx` is the standard).  
- Any proposed tech stack change requires an **Architecture Decision Record (ADR)** in `/docs/adrs/`.  
- Violations will be flagged in code review.  

---

👉 This doc becomes the **single source of truth** for technology decisions in the project.
