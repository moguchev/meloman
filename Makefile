export GO111MODULE=on

.PHONY: up
up:
	docker-compose build
	docker-compose up -d pg bouncer server

.PHONY: down
down:
	docker-compose down

.PHONY: build
build:
	go build -o ./bin/meloman -v ./cmd/meloman

.PHONY: migration
migration:
	migrate create -ext sql -dir ./db/migrations -seq $(SEQ)  