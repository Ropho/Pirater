.DEFAULT_GOAL := build

.PHONY: build
build:
	cd front-app && npm run build && cd ..  ;
	docker build -t nginxapp .				;
	go build -v -o ./bin/prog ./cmd/server 	; 
	swag init -g ./cmd/server/server.go

# make run ARGS="-config config2.yaml"
# make run ARGS="-config config.yaml"
.PHONY: run
run:
	docker compose up -d		;
	sleep 1						;
	./bin/prog $(ARGS)

.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database "mysql://root:2280@tcp(localhost:3306)/piraterDB" up

.PHONY: migrate_down
migrate_down:
	migrate -path migrations -database "mysql://root:2280@tcp(localhost:3306)/piraterDB" down


.PHONY: clean
clean: rm_im
	docker compose down													;
	rm --force logs/server/access.log									;
	rm --force logs/server/error.log									;


Trash_Images := $(shell docker images -q --filter dangling=true)
empty:=
ifneq ($(Trash_Images),$(empty))
.PHONY: rm_im
rm_im:
	docker rmi $(Trash_Images)
else
.PHONY: rm_im
rm_im:
endif