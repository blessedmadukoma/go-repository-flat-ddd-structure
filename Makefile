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
	migrate create -ext sql -dir db/migrations -seq $(name)

db/migrateup:
	migrate -path database/migrations -database $(DB_SOURCE) -verbose up
	# migrate -path db/migrations -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

db/migratedown:
	# migrate -path database/migrations -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down
	migrate -path database/migrations -database $(DB_SOURCE) -verbose down

dc_up:
	docker compose up -d

dc_down:
	docker compose down -v

# db_up:
# 	# create database
# 	docker exec -it fintrax_db createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)
# 	docker exec -it fintrax_db_live createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

# db_down:
# 	# drop database
# 	docker exec -it fintrax_db dropdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)
# 	docker exec -it fintrax_db_live dropdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

sqlc:
	sqlc generate

start:
	CompileDaemon -command="./fintrax"

test:
	go test -v -cover ./...

.PHONY:	db/createmigration db/migrateup db/migratedown pg_up pg_down db_up db_down sqlc start test