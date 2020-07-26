.PHONY: build test start

build:
	go build -v ./cmd/makeba

test: 
	go test -v -race -timeout 30s ./...

start: 
	go build -v ./cmd/makeba
	./makeba

.DEFAULT_GOAL := build