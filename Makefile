include .env
GO ?= go
GOBUILD ?= $(GO) build

MAIN_FILE ?= ./cmd/server/main.go


DB_DSN := "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) mod download
	$(GO) install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: build
build:
	@echo "Building the application..."
	$(GOBUILD) -o bin/main $(MAIN_FILE)
