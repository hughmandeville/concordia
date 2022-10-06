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
## DOCKER:

.PHONY: docker-build-db
## docker-build-db: Build Docker Concordia DB image (concordia-db).
docker-build-db:
	cd db; docker build -t concordia-db -f concordia-db.dockerfile .

.PHONY: docker-open-db
## docker-open-db: Connect to Concordia DB running in Docker.
docker-open-db:
	mysql -h 127.0.0.1 -u root concordia

.PHONY: docker-run-db
## docker-run-db: Run Docker container with Concordia DB image (concordia-db).
docker-run-db:
	docker run -d --name concordia-db -p 3306:3306 -t concordia-db

.PHONY: docker-stop-db
## docker-stop-db: Stop and delete Docker Concordia DB container (concordia-db).
docker-stop-db:
	docker stop concordia-db && docker rm concordia-db


## :
## RUN:

.PHONY: run-server
## run-server: Run Go server locally (on port 8080).
run-server:
	cd cmd/server; go run ./...

## :
