.DEFAULT_GOAL := build
CH?=ch01
EXAMPLE?=ch01-hello-go

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
	# shadow ./...
.PHONY:vet

build: vet
	go build examples/$(CH)/$(EXAMPLE)/main.go
.PHONY:build
