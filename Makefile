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

.PHONY: generate
generate:
	buf beta mod update 
	buf generate
	cp ./api/meloman/meloman.swagger.json ./swaggerui/

.PHONY: deps
deps:
	# go get -v -d ./...
	go mod tidy