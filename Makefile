# Makefile

## HELP:
.PHONY: help
## help: Show this help message.
help:
	@echo "Usage: make [target]\n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## :
## BUILD:

.PHONY: build
## build: Build Go server code.
build:
	go build -o /dev/null ./...

## :
## DATABASE:

.PHONE: db-start
## db-start: Start database in Docker.
db-start:
	docker run --name 14.1-alpine -e POSTGRES_PASSWORD=dbpw -d postgres

## :
## DEPENDENCIES:

.PHONY: dep-clean
## dep-clean: Clean up dependency files.
dep-clean:
	@rm go.mod go.sum

.PHONY: dep-get
## dep-get: Get Go modules.
dep-get:
	go mod tidy

.PHONY: dep-init
## dep-init: Initialize Go modules.
dep-init:
	go mod init

.PHONY: dep-update
## dep-update: Update Go modules.
dep-update:
	go get -u ./...
	go mod tidy

## :
## RUN:

.PHONY: run-server
## run-server: Run Go server locally (on port 8080).
run-server:
	cd cmd/server; go run ./...

## :
