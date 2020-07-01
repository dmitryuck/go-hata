.PHONY: build

build:
	go build -v ./

run:
	go run ./cmd/apiserver

.DEFAULT_GOAL := build