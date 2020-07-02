.PHONY: build

build:
	go build -v ./cmd/apiserver

run:
	go build -o ./bin/apiserver.exe -v ./cmd/apiserver
	./bin/apiserver.exe

.DEFAULT_GOAL := build