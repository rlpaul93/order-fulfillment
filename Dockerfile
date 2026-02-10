# Start from the official Golang image for build
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o order-fulfillment ./cmd/api

# Use a minimal image for running
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-fulfillment ./order-fulfillment
COPY internal/infrastructure/db/migrations ./migrations

# Set environment variable for database URL (override in production)
ENV DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable

EXPOSE 8080
CMD ["./order-fulfillment"]
