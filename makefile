.DEFAULT_GOAL := build

.PHONY: build_front
build_front:
	cd front-app 	&&
	npm run build	&&
	cd ..  


.PHONY: build
build:
	GOOS=linux go build -v -o ./bin/prog ./cmd/server
	swag init -g ./cmd/server/server.go
	docker build -t goapp -f ./docker/app/Dockerfile .
	docker build -t nginxapp -f ./docker/nginx/Dockerfile .

# make run ARGS="-config config2.yaml"
# make run ARGS="-config config.yaml"
.PHONY: run
run:
	docker compose up
# sleep 1						;
# ./bin/prog $(ARGS)

.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database "mysql://root:2280@tcp(localhost:3306)/piraterDB" up

.PHONY: migrate_down
migrate_down:
	migrate -path migrations -database "mysql://root:2280@tcp(localhost:3306)/piraterDB" down


.PHONY: clean
clean: rm_im
	docker compose down					;
	rm --force logs/server/access.log									;
	rm --force logs/server/error.log									;

Trash_Images := $(shell docker images -q --filter dangling=true)
empty:=
ifneq ($(Trash_Images),$(empty))
.PHONY: rm_im
rm_im: rm_conts
	docker rmi $(Trash_Images)
else
.PHONY: rm_im
rm_im: rm_conts
endif

Stopped_conts := $(shell docker ps -aq)
ifneq ($(Stopped_conts),$(empty))
.PHONY: rm_conts
rm_conts:
	docker rm $(Stopped_conts)
else
.PHONY: rm_conts
rm_conts:
endif