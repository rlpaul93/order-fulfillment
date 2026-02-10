MIGRATIONS_DIR=internal/infrastructure/db/migrations
GOBIN:=$(shell go env GOBIN)
MIGRATE_BIN=$(GOBIN)/migrate
DB_URL=postgres://user:password@localhost:5432/order_fulfillment?sslmode=disable

.PHONY: migrate-create migrate-install migrate-up migrate-down docker-build docker-run compose-up compose-down api-logs
# View logs from the API container
api-logs:
	docker-compose logs -f api

# Start all services with docker-compose
compose-up:
	docker-compose up -d --build

# Stop all services with docker-compose
compose-down:
	docker-compose down
# Build Docker image
docker-build:
	docker build -t order-fulfillment .

# Run Docker container
docker-run:
	docker run --rm -it -p 8080:8080 -e DATABASE_URL=$${DATABASE_URL} order-fulfillment

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATIONS_DIR) -format 20260210174532 $(name)

# Run all up migrations
migrate-up:
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database $(DB_URL) up

# Run all down migrations
migrate-down:
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database $(DB_URL) down
