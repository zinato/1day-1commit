.DEFAULT_GOAL := build
EXAMPLE_DIR?=ch01-hello-go

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build examples/$(EXAMPLE_DIR)/main.go
.PHONY:build
