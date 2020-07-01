.PHONY: build

build:
	go build -v ./cmd/apiserver

run:
	go build -v ./cmd/apiserver
	./apiserver.exe

.DEFAULT_GOAL := build