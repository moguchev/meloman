export GO111MODULE=on

.PHONY: up
up:
	docker-compose build
	docker-compose up -d pg bouncer server nginx

.PHONY: down
down:
	docker-compose down

.PHONY: stop
stop:
	docker-compose stop pg bouncer server nginx

.PHONY: start
start:
	docker-compose start pg bouncer server nginx

.PHONY: up_db
up_db:
	docker-compose up -d pg bouncer

.PHONY: build
build:
	go build -o ./bin/meloman -v ./cmd/meloman

.PHONY: migration
migration:
	migrate create -ext sql -dir ./db/migrations -seq $(SEQ)

.PHONY: generate
generate:
	buf beta mod update 
	buf generate
	cp ./api/meloman/meloman.swagger.json ./swaggerui/

.PHONY: deps
deps:
	# go get -v -d ./...
	go mod tidy

.PHONY: run
run:
	DATABASE_URL=postgres://user:password@localhost:6432/meloman?sslmode=disable ./bin/meloman
