.PHONY: build
build:
	go build -v -o ./bin/prog ./cmd/server ; 
	swag init -g ./cmd/server/server.go

.PHONY: run
run:
	./bin/prog

.DEFAULT_GOAL := build