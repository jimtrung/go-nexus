run:
	templ generate
	go build -o ./tmp/main ./cmd/server

include .env
export $(shell sed 's/=.*//' .env)

MIGRATION_DIR=migrations

MIGRATE_CMD=migrate -database "$(DB_URL)" -path "$(MIGRATION_DIR)"

.PHONY: migrate-up migrate-down migrate-force

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down 1

migrate-force:
	$(MIGRATE_CMD) force $(VERSION)

lines:
	 find . -type f \( -name "*.go" -o -name "*.templ" \) -print0 | xargs -0 wc -l
