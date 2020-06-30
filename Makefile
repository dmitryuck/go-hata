.PHONY: build

build:
	go build -v ./

run:
	go run ./

.DEFAULT_GOAL := build