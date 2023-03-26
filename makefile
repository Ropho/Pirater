.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -v -o ./bin/prog ./cmd/server ; 
	swag init -g ./cmd/server/server.go


.PHONY: run
run:
	./bin/prog


.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database "mysql://root:2280@tcp(localhost:3306)/piraterDB" up

.PHONY: migrate_down
migrate_down:
	migrate -path migrations -database "mysql://root:2280@localhost:3306/piraterDB" down