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
│   │   ├── application/   # Contains use cases following CQRS pattern with command/ and query/ subdirectories.
│   │   │   ├── command/   # Command handlers for write operations.
│   │   │   └── query/     # Query handlers for read operations.
│   │   ├── domain/        # Contains only entities and their methods/validations (pure, no dependencies).
│   │   ├── ports/         # Defines the interfaces (contracts) this module depends on or exposes.
│   │   ├── presentation/  # Contains HTTP handlers and route registration (can access application layer directly).
│   │   │   └── http/      # HTTP handlers with RegisterRoutes methods.
│   │   └── infrastructure/# Contains implementations for external concerns (DB, external APIs, etc.).
│   └── module-two/     # Another self-contained business domain.
│       ├── application/
│       │   ├── command/
│       │   └── query/
│       ├── domain/
│       ├── ports/
│       ├── presentation/
│       │   └── http/
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

### Rule 4: The Domain Layer is Pure and Entity-Focused
The domain directory within any module contains **only entities** of that module and their related methods and validations. It must be completely self-contained and have **zero external dependencies**.

- It **cannot** import from the application, infrastructure, presentation, or ports layers of its own module.
- It **cannot** contain any code related to databases, web frameworks, or any other external concern.
- It should only contain domain entities, value objects, and their business logic methods.

### Rule 5: The Presentation and Infrastructure Layer Separation
The **presentation** layer contains HTTP handlers and route registration, and can have direct access to the application layer.
The **infrastructure** layer contains implementations for external concerns like database repositories and third-party API clients.

- **Presentation Layer:** Contains HTTP handlers, route registration, and can directly access the application layer.
- **Infrastructure Layer:** Contains database repositories, external API clients, and other driven adapters.
- HTTP API routes should be registered in `RegisterRoutes` methods within each module's `presentation/http/` directory, not in main.go.

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

### Rule 8: Application Layer CQRS Pattern
The application layer must follow the CQRS (Command Query Responsibility Segregation) pattern with separate command and query folders:

- **Command handlers** (in `application/command/`) handle write operations and business logic that modifies state.
- **Query handlers** (in `application/query/`) handle read operations and data retrieval.
- This separation ensures clear responsibility boundaries and supports scalability.

### Rule 9: Documentation Folder Restrictions
The `/docs/` folder is exclusively for project blueprint documentation (architecture, features, guidelines, tech stack, etc.).

- **Prohibited:** Swagger files, OpenAPI specs, or any files needed for project runtime or building.
- **Allowed:** Architecture docs, feature specifications, technical guidelines, ADRs (Architecture Decision Records).
- Runtime API documentation should be placed in `/api/` directory.

### Rule 10: Database Migration Management
- All database schema changes must be managed through versioned migration files in `/scripts/migrations/`.
- Use `golang-migrate` tool with sequential numbering (e.g., `000001_initial_schema.up.sql`).
- Every migration must have both `.up.sql` and `.down.sql` files for rollback capability.
- Migration execution varies by environment:
  - **Development**: Auto-run via docker-compose for convenience
  - **CI/CD**: Run as separate job before service deployment
  - **Production**: Manual execution through controlled deployment process
- Use Makefile commands (`make migrate-up`, `make migrate-down`) for consistency.

### Rule 11: Presentation Layer Direct Access Pattern
The presentation layer of each module **MUST** access command and query handlers of the same module directly, not through interfaces defined in ports.

- **Rationale:** The presentation layer is the entry point for external requests and should have direct access to the application layer within the same module for simplicity and performance.
- **Implementation:** HTTP handlers in `presentation/http/` should directly instantiate and call command/query handlers from `application/command/` and `application/query/`.
- **Prohibited:** Creating port interfaces for intra-module communication between presentation and application layers.
- **Example:** `UserHandler` should directly use `CreateUserCommandHandler` and `GetUserQueryHandler`, not through a `UserServicePort` interface.

### Rule 12: Logger Injection Requirements
Logger must be injected to all layers except the domain layer which contains entities and value objects.

- **Domain Layer Exception:** The domain layer (entities and VOs) must remain pure and cannot have logger dependencies to maintain business logic isolation.
- **Required Injection:** All other layers (application, infrastructure, presentation) must receive logger instances through dependency injection.
- **Implementation:** Use constructor injection to pass logger instances to handlers, repositories, and other components.
- **Logger Interface:** All modules must use the `logger.Logger` interface from `platform/logger` package, never import specific logging libraries directly (e.g., zerolog, logrus).
- **Abstraction Benefit:** This allows switching underlying logger implementations without modifying module code.
- **Consistency:** Use the same logger interface across all modules to ensure uniform logging behavior and maintain loose coupling.
- **Example:** `CreateUserCommandHandler(logger logger.Logger, userRepo UserRepository)` but `User` entity has no logger dependency.

---

## 4. Migration Path to Microservices
When extracting a module into its own service:
1. Promote its `/application + domain + ports + infrastructure` into a new repo.  
2. Keep its existing ports contracts as the new API surface.  
3. Replace local wiring with remote communication (HTTP, gRPC, messaging).  
4. Keep other modules consuming the same ports interfaces to minimize impact.  

---

