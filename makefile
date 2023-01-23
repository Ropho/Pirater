.PHONY: build
build:
	go build -v -o ./bin/prog ./cmd/server

.PHONY: run
run:
	./bin/prog

.DEFAULT_GOAL := build