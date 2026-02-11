
# order-fulfillment

## Architecture

This project follows a **hexagonal (ports and adapters) architecture**:

- **Domain**: Core business logic and models, independent of frameworks and infrastructure.
- **Ports**: Interfaces for application and infrastructure boundaries (e.g., repositories, services).
- **Adapters**: Implementations of ports for HTTP (handlers) and database (repositories).
- **Application/Service**: Orchestrates use cases and business rules.
- **Infrastructure**: Database, server, and external system integrations.
- **API**: Entry point wiring up dependencies and HTTP routes.

### Project Structure

```
.
├── cmd/
│   └── api/
│       ├── main.go              # Application entry point
│       ├── config/              # Configuration loading
│       └── factory/             # Dependency injection / wiring
├── docs/                        # Swagger generated documentation
├── internal/
│   ├── adapters/
│   │   ├── in/                  # Inbound adapters (HTTP handlers)
│   │   └── out/                 # Outbound adapters (DB repositories)
│   ├── domain/
│   │   ├── model/               # Domain entities (Product, Pack)
│   │   ├── port/                # Interfaces (repository contracts)
│   │   └── service/             # Business logic services
│   └── infrastructure/
│       ├── db/                  # Database connection
│       └── server/              # HTTP server and routing
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── go.mod
```

This structure ensures separation of concerns, testability, and flexibility for future changes.

## API Documentation

Swagger UI is available at: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

To regenerate swagger docs after code changes, run:
```bash
make swagger
```

## Makefile Commands

- **make migrate-install**: Install the golang-migrate tool with PostgreSQL support.
- **make migrate-create name=your_migration_name**: Create a new migration file.
- **make migrate-up**: Apply all up migrations to the database.
- **make migrate-down**: Revert all migrations (danger: this will drop tables!).
- **make docker-build**: Build the API Docker image.
- **make docker-run**: Run the API Docker image locally.
- **make compose-up**: Start all services (API and DB) with Docker Compose.
- **make compose-down**: Stop all services started by Docker Compose.
- **make api-logs**: View logs from the running API container.

## Running the Full Solution Locally

The solution includes three services: **Frontend** (React), **API** (Go), and **Database** (PostgreSQL). All can be run together using Docker Compose.

### Prerequisites

- Docker and Docker Compose installed
- The frontend repository `order-form-ui` cloned as a sibling directory:
  ```
  /your-path/
  ├── order-fulfillment/     # This repository
  └── order-form-ui/         # Frontend repository
  ```

### Quick Start

1. **Start all services:**
   ```bash
   make compose-up
   ```

2. **Run database migrations** (first time only):
   ```bash
   make migrate-up
   ```

3. **Access the applications:**
   - **Frontend:** http://localhost:5173
   - **API:** http://localhost:8080
   - **Swagger UI:** http://localhost:8080/swagger/index.html

4. **View logs:**
   ```bash
   make api-logs              # API logs only
   docker-compose logs -f     # All services
   ```

5. **Stop all services:**
   ```bash
   make compose-down
   ```

### Services Overview

| Service   | Port | Description                          |
|-----------|------|--------------------------------------|
| frontend  | 5173 | React application (Vite dev server)  |
| api       | 8080 | Go REST API                          |
| db        | 5432 | PostgreSQL database                  |

## Next steps
- Create middleware for authentication and authorization
- Improve logging for better traceability
- Add versioning to the API