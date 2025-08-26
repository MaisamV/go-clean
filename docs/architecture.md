# Go Project Structure: A Modular Monolith Template (Improved - Ports in Modules, Platform at Root)

## 1. Philosophy
This document outlines the standard project structure for Go applications, designed as a **Modular Monolith**. 
The primary goal is to support a large project within a single repository and binary, while enforcing strict boundaries between business domains.

This approach provides the development velocity of a monolith with a clear and low-effort path to extracting domains into separate microservices in the future.

The core principles are:
- **High Cohesion, Low Coupling:** Code related to a single business domain is kept together. Dependencies between domains are minimized and strictly controlled.
- **Clean Architecture:** The structure separates concerns into distinct layers: Domain (business rules), Application (use cases), and Infrastructure (external concerns like databases, APIs, etc.).
- **Future-Proofing:** The architecture is designed to make the future transition to a microservices architecture a surgical and low-risk process.

## 2. Directory Structure
```structure
.
├── api/                # Contains API contract files (e.g., OpenAPI specs, .proto files).
├── build/              # Contains CI configurations and build-related artifacts (e.g., Helm, k8s manifests).
├── bin/                # The location of compiled binaries.
├── cmd/                # Contains the application's entry points (main packages).
│   └── app/            # An executable for the main application.
│       └── main.go     # The main function: initializes and wires up all components.
├── configs/            # Contains configuration file templates (e.g., config.yaml.example).
├── docs/               # Contains project documentation (e.g., ADRs, architecture, features, tech stack docs).
├── internal/           # Contains all private application code, not importable by other projects, each subfolder must be a module.
│   ├── module-one/     # A self-contained business domain (e.g., "users", "billing").
│   │   ├── application/   # Contains the application use cases and business flows.
│   │   ├── domain/        # Contains the core business models and logic (pure, no dependencies).
│   │   ├── ports/         # Defines the interfaces (contracts) this module depends on or exposes.
│   │   └── infrastructure/# Contains implementations for external concerns (DB, APIs, handlers).
│   └── module-two/     # Another self-contained business domain.
│       ├── application/
│       ├── domain/
│       ├── ports/
│       └── infrastructure/
├── platform/           # Shared, non-business foundational code (e.g., DB conn, HTTP server, logger).
├── pkg/                # Shared library code intended to be imported by external projects (true external reusables).
├── scripts/            # Contains helper scripts for development (e.g., migrations, code generation).
│   └── migrations/     # Database migration files using golang-migrate (versioned SQL files).
├── test/               # Contains end-to-end and integration tests. Unit tests stay in-module as *_test.go files.
├── .gitignore          
├── .golangci.yml       
├── docker-compose.yml  
├── Dockerfile          
├── go.mod              
├── go.sum              
├── Makefile            
└── README.md           
```

## 3. Architectural Guidelines for All Contributors
Adherence to these rules is mandatory to maintain the integrity and scalability of the architecture.

### Rule 1: The Golden Rule - No Direct Inter-Module Imports
Code within one module (e.g., /internal/module-one) **MUST NEVER** directly import code from another module (e.g., /internal/module-two). 
All inter-module communication must go through `ports` defined inside the provider or consumer module.

> ⚠️ This rule is not enforced by Go alone. It should be checked by code reviews and optionally custom linters.

### Rule 2: Communication Through Module Ports
Each module defines the interfaces it depends on (ports) in its own `/ports` directory.  
Consumer modules depend on these interfaces, and provider modules implement them.

- **Example:** If the billing module needs to check if a user is active, it will depend on a `UserServicePort` interface inside `billing/ports`.  
  The users module will provide the concrete implementation.

### Rule 3: Dependency Injection is Mandatory
Modules must not instantiate their own dependencies. 
All dependencies (database connections, repository implementations, other module services via their ports) must be passed into a module's components via constructors.

The final dependency graph for the entire application is assembled **only once**, in the `/cmd/app/main.go` file. This is the "wiring" layer.

### Rule 4: The Domain Layer is Pure
The domain directory within any module represents the core business logic. It must be completely self-contained and have **zero external dependencies**.

- It **cannot** import from the application, infrastructure, or ports layers of its own module.
- It **cannot** contain any code related to databases, web frameworks, or any other external concern.

### Rule 5: The Infrastructure Layer is the Boundary
The infrastructure directory is the single, unified home for all code that interacts with the outside world. 
This includes both:

- **Driving Adapters (Input):** Code that drives the application, like HTTP handlers, gRPC services, and message queue consumers.
- **Driven Adapters (Output):** Code that is driven by the application, like database repository implementations and clients for third-party APIs.

This creates a clear boundary: your core logic (application + domain) is on one side, and all external details are on the other.

### Rule 6: The Platform Folder is for Non-Business Code
The `/platform` directory is for shared, foundational code that is **not specific to any business domain**. 
This includes setting up the database connection pool, configuring the shared HTTP server, or initializing a common logger.

**Rule of thumb:**
- If it is business-related → put it in the module’s `infrastructure`.
- If it is cross-cutting and shared → put it in `platform`.

### Rule 7: Testing Conventions
- **Unit tests**: co-located with the code they test (`*_test.go` inside each module).  
- **Integration/E2E tests**: live in `/test/`.  
- Shared mocks or test utilities can live in `/platform/testing/`.

### Rule 8: Database Migration Management
- All database schema changes must be managed through versioned migration files in `/scripts/migrations/`.
- Use `golang-migrate` tool with sequential numbering (e.g., `000001_initial_schema.up.sql`).
- Every migration must have both `.up.sql` and `.down.sql` files for rollback capability.
- Migration execution varies by environment:
  - **Development**: Auto-run via docker-compose for convenience
  - **CI/CD**: Run as separate job before service deployment
  - **Production**: Manual execution through controlled deployment process
- Use Makefile commands (`make migrate-up`, `make migrate-down`) for consistency.

---

## 4. Migration Path to Microservices
When extracting a module into its own service:
1. Promote its `/application + domain + ports + infrastructure` into a new repo.  
2. Keep its existing ports contracts as the new API surface.  
3. Replace local wiring with remote communication (HTTP, gRPC, messaging).  
4. Keep other modules consuming the same ports interfaces to minimize impact.  

---

