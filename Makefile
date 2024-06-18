include .env

PROJECTNAME=$(shell basename "$(PWD)")

## help: Display this help screen
help: Makefile
	@echo " Choose a command to run in $(PROJECTNAME):"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## install: Install missing dependencies. Runs `go mod tidy` and `go mod download`.
install:
	go mod tidy
	go mod download

## start: Start the application using Docker Compose.
start:
	docker-compose up --build

## stop: Stop the application using Docker Compose.
stop:
	docker-compose down

## test: Run tests inside the Go container.
test:
	docker-compose run web go test -v ./...

## lint: Run linter to check for issues.
lint:
	docker-compose run web golangci-lint run ./...

## build: Build the application using Docker.
build:
	docker-compose build

## clean: Remove built files and Docker volumes.
clean:
	docker-compose down -v
	rm -rf bin/

.PHONY: help install start stop test lint build clean
