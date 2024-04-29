include .env

DB_DRIVER := $(DB_DRIVER)
DB_USER := $(DB_USER)
DB_PASSWORD := $(DB_PASSWORD)
DB_HOST := $(DB_HOST)
DB_PORT := $(DB_PORT)
DB_NAME := $(DB_NAME)
DB_SOURCE := $(DB_SOURCE)
DB_DOCKER_HOST := $(DB_DOCKER_HOST)

db/createmigration:
	migrate create -ext sql -dir database/migrations -seq $(name)

db/migrateup:
	migrate -path database/migrations -database $(DB_SOURCE) -verbose up

db/migratedown:
	migrate -path database/migrations -database $(DB_SOURCE) -verbose down

dc_up:
	docker compose up -d

dc_down:
	docker compose down -v

sqlc:
	sqlc generate

start:
	CompileDaemon -command="./goRepositoryPattern"

build:
	go build -o goRepositoryPattern

test:
	go test -v -cover ./...

.PHONY:	db/createmigration db/migrateup db/migratedown pg_up pg_down db_up db_down sqlc start test