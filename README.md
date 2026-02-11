
# order-fulfillment

## Architecture

This project follows a **hexagonal (ports and adapters) architecture**:

- **Domain**: Core business logic and models, independent of frameworks and infrastructure.
- **Ports**: Interfaces for application and infrastructure boundaries (e.g., repositories, services).
- **Adapters**: Implementations of ports for HTTP (handlers) and database (repositories).
- **Application/Service**: Orchestrates use cases and business rules.
- **Infrastructure**: Database, server, and external system integrations.
- **API**: Entry point wiring up dependencies and HTTP routes.

This structure ensures separation of concerns, testability, and flexibility for future changes.

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
