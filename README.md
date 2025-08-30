# Go Clean Architecture Template

A production-ready **Modular Monolith** template for Go applications, designed to provide the development velocity of a monolith with a clear path to microservices extraction.

## 🎯 Purpose

This template serves as a foundation for all Go projects, implementing clean architecture principles with strict domain boundaries. It's designed to:

- **Scale Development Teams**: Multiple teams can work on different modules without conflicts
- **Maintain Code Quality**: Enforced architectural boundaries prevent technical debt
- **Enable Future Migration**: Clean module separation allows easy extraction to microservices
- **Provide Production Readiness**: Includes monitoring, health checks, and Kubernetes integration

## 🏗️ Architecture

This project follows a **Modular Monolith** architecture with clean separation of concerns:

```
├── cmd/                # Application entry points
├── internal/           # Private application modules
│   └── module-name/    # Self-contained business domains
│       ├── application/   # Use cases and business flows
│       ├── domain/        # Core business logic (pure)
│       ├── ports/         # Interface definitions
│       └── infrastructure/# External integrations
├── platform/           # Shared foundational code
├── docs/               # Project documentation
└── test/               # Integration and E2E tests
```

### Key Principles

- **No Direct Inter-Module Imports**: Modules communicate only through defined ports
- **Clean Architecture Layers**: Domain → Application → Infrastructure separation
- **Dependency Injection**: All dependencies injected via constructors
- **Pure Domain Layer**: Business logic with zero external dependencies

For detailed architecture guidelines, see [`docs/architecture.md`](docs/architecture.md).

## 🛠️ Technology Stack

| Component | Technology | Purpose |
|-----------|------------|----------|
| **HTTP Framework** | [GoFiber](https://github.com/gofiber/fiber) | Fast HTTP server and routing |
| **Database** | [PostgreSQL](https://postgresql.org) + [pgx](https://github.com/jackc/pgx) | Primary data persistence |
| **Cache** | [Redis](https://redis.io) + [go-redis](https://github.com/redis/go-redis) | Caching and temporary storage |
| **Logging** | [zerolog](https://github.com/rs/zerolog) | Structured JSON logging |
| **Configuration** | [viper](https://github.com/spf13/viper) | Configuration management |
| **Dependency Injection** | [wire](https://github.com/google/wire) | Compile-time DI |
| **Testing** | Go testing + [testify](https://github.com/stretchr/testify) | Unit and integration tests |
| **Validation** | [validator](https://github.com/go-playground/validator) | Request validation |

For complete tech stack guidelines, see [`docs/tech-stack.md`](docs/tech-stack.md).

## ✨ Features

This template includes essential APIs for production deployment:

### Health & Monitoring APIs

- **`GET /ping`** - Simple alive check returning "PONG"
- **`GET /health`** - Comprehensive health check validating all dependencies
- **`GET /liveness`** - Internal service health for Kubernetes restart decisions

These endpoints are designed as Kubernetes probes for:
- **Readiness Probes**: `/ping` and `/health`
- **Liveness Probes**: `/liveness`

For detailed feature specifications, see [`docs/features.md`](docs/features.md).

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- PostgreSQL 14+
- Redis 6+
- Docker & Docker Compose (for local development)

### Local Development Setup

1. **Clone the template**:
   ```bash
   git clone <repository-url>
   cd go-clean
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Start infrastructure services**:
   ```bash
   docker-compose up -d postgres redis
   ```

4. **Set up configuration**:
   ```bash
   cp configs/config.yaml.example configs/config.yaml
   # Edit configs/config.yaml with your settings
   ```

5. **Run database migrations**:
   ```bash
   make migrate-up
   ```

6. **Start the application**:
   ```bash
   go run cmd/app/main.go
   ```

7. **Verify the setup**:
   ```bash
   curl http://localhost:8080/ping
   curl http://localhost:8080/health
   ```

### Using as a Template

1. **Use this repository as a template** when creating new projects
2. **Rename the module** in `go.mod` to your project name
3. **Update configuration** in `configs/` directory
4. **Add your business modules** in `internal/`
5. **Update documentation** to reflect your project specifics

## 📁 Project Structure

```
.
├── api/                # API contracts (OpenAPI, protobuf)
├── build/              # CI/CD and deployment configs
├── cmd/
│   └── app/            # Main application entry point
├── configs/            # Configuration templates
├── docs/               # Project documentation
│   ├── architecture.md # Architecture guidelines
│   ├── features.md     # Feature specifications
│   └── tech-stack.md   # Technology decisions
├── internal/           # Private application code
│   └── [modules]/      # Business domain modules
├── platform/           # Shared infrastructure code
├── pkg/                # Public library code
├── scripts/            # Development and deployment scripts
├── test/               # Integration and E2E tests
├── docker-compose.yml  # Local development environment
├── Dockerfile          # Container image definition
└── Makefile            # Development commands
```

## 🧪 Testing

```bash
# Run unit tests
make test

# Run integration tests
make test-integration

# Run all tests with coverage
make test-coverage

# Run linting
make lint
```

## 🐳 Docker Support

### Development Environment
```bash
# Start all services
docker-compose up

# Start only infrastructure
docker-compose up postgres redis
```

### Production Build
```bash
# Build application image
docker build -t go-clean-app .

# Run with docker-compose
docker-compose -f docker-compose.prod.yml up
```

## ☸️ Kubernetes Deployment

The application includes health check endpoints designed for Kubernetes:

```yaml
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: app
    image: go-clean-app
    ports:
    - containerPort: 8080
    livenessProbe:
      httpGet:
        path: /liveness
        port: 8080
      initialDelaySeconds: 30
      periodSeconds: 10
    readinessProbe:
      httpGet:
        path: /health
        port: 8080
      initialDelaySeconds: 5
      periodSeconds: 5
```

## 🔧 Development Commands

### Database Migrations

```bash
# Install migration tool
make migrate-install

# Apply all pending migrations
make migrate-up

# Rollback last migration
make migrate-down

# Check migration status
make migrate-status

# Create new migration
make migrate-create NAME=add_user_roles

# Reset database (⚠️ DESTRUCTIVE)
make migrate-reset

# Force migration version (emergency only)
make migrate-force
```

### Migration Workflow by Environment

**Development**:
- Migrations run automatically via `docker-compose up`
- Use `make migrate-*` commands for manual control
- Database is disposable, aggressive changes are acceptable

**CI/CD**:
- Migrations run as separate job before service deployment
- Must complete successfully before deploying application
- Use `make migrate-up` in pipeline

**Production**:
- ⚠️ **NEVER** run migrations automatically
- Always run manually or via approved pipeline
- Backup database before applying migrations
- Test on staging environment first

### Other Development Commands

```bash
# Install development tools
make install-tools

# Generate code (wire, mocks, etc.)
make generate

# Format code
make fmt

# Run security checks
make security

# Build application
make build
```

## 📖 Documentation

- **[Architecture Guide](docs/architecture.md)** - Detailed architectural principles and rules
- **[Technology Stack](docs/tech-stack.md)** - Official technology decisions and guidelines
- **[Features](docs/features.md)** - Implemented features and specifications
- **[ADRs](docs/adrs/)** - Architecture Decision Records (when applicable)

## 🤝 Contributing

1. **Follow the architecture guidelines** defined in `docs/architecture.md`
2. **Use only approved technologies** from `docs/tech-stack.md`
3. **Use Wire for dependency injection** - Manual DI is strictly prohibited
4. **Update documentation** when adding features
5. **Write tests** for all new functionality
6. **Follow Git flow** for branching and commits

### Code Review Checklist

- [ ] No direct inter-module imports
- [ ] Dependencies injected via constructors **using Wire**
- [ ] Each module has `wire.go` file with proper build tags
- [ ] No manual dependency injection in production code
- [ ] Domain layer remains pure
- [ ] Approved technologies used
- [ ] Tests included
- [ ] Documentation updated

## 📄 License

[Add your license here]

## 🆘 Support

For questions about this template:

1. Check the documentation in `docs/`
2. Review existing issues
3. Create a new issue with detailed description

---

**Note**: This is a template repository. When using it for your project, update this README to reflect your specific application details, remove template-specific content, and add your project's unique information.