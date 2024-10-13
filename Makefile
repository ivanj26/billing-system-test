ENV_FILE ?= .env
BINARY_NAME="loan-app"
DB_NAME ?= $(shell sed -n 's/^DB_DATABASE=[[:space:]]*\(.*\)/\1/p' $(ENV_FILE))
DB_PORT ?= $(shell sed -n 's/^DB_PORT=[[:space:]]*\(.*\)/\1/p' $(ENV_FILE))
DB_PASS ?= $(shell sed -n 's/^DB_PASSWORD=[[:space:]]*\(.*\)/\1/p' $(ENV_FILE))
DB_HOST ?= $(shell sed -n 's/^DB_HOST=[[:space:]]*\(.*\)/\1/p' $(ENV_FILE))
DB_USER ?= $(shell sed -n 's/^DB_USERNAME=[[:space:]]*\(.*\)/\1/p' $(ENV_FILE))

build-local:
	go build -v -o ${BINARY_NAME} app.go

run:
	make build-local && ENV_LOC=.env ./$(BINARY_NAME)

db-create:
	mysql -h $(DB_HOST) -u $(DB_USER) -p$(DB_PASS) -P $(DB_PORT) -e "CREATE DATABASE $(DB_NAME);"

db-down:
	mysql -h $(DB_HOST) -u $(DB_USER) -p$(DB_PASS) -P $(DB_PORT) -e "DROP DATABASE $(DB_NAME);"

db-migrate:
	for file in $$(ls ddl/*.sql); do \
		mysql -h $(DB_HOST) -u $(DB_USER) -p$(DB_PASS) -P $(DB_PORT) "$(DB_NAME)" < $${file}; \
	done

db-seed:
	mysql -h $(DB_HOST) -u $(DB_USER) -p$(DB_PASS) -P $(DB_PORT) "$(DB_NAME)" < dummy.sql