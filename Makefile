# Load .env configurations
ifneq (,$(wildcard ./.env))
    include .env
    export
    MIGRATIONS_SOURCE = file://db/migrations
	DATABASE_URL = postgres://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=$(DATABASE_SSLMODE)
endif

# Tools as dependencies installer bash
install-tools:
	bash scripts/install_tools

# Run REST API server
rest:
	go run cmd/rest-server/main.go

# Generate OpenAPI json & yaml
openapi-gen:
	go run cmd/openapi-gen/main.go

# Run golang-migrate up
migrate-up:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" up

# Run golang-migrate down
migrate-down:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" down -all

# Run golang-migrate create
migrate-create:
	migrate create -ext sql -dir ./db/migrations $(filter-out $@,$(MAKECMDGOALS))

# Run SQLC generator
sqlc-gen:
	sqlc generate

# Run gofmt code formatter
gofmt:
	gofmt -s -w .

# Generate protocol buffer files.
proto:
	protoc --go_out=. --go-grpc_out=. pkg/proto/*