.PHONY: db-create db-reset db-migrate db-rollback db-migrate-create dev dev-setup

setup: install-dependencies db-reset

install-dependencies:
	go install github.com/cosmtrek/air@latest
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go get github.com/mattn/go-sqlite3

db-create:
	@if [ ! -f database.db ]; then \
		sqlite3 database.db < schema.sql; \
		echo "Created database.db"; \
	else \
		echo "Database already exists"; \
	fi

db-migrate:
	@if [ -z "$$(find migrations -name '*.up.sql' -print -quit)" ]; then \
		echo "No migrations found."; \
	else \
		migrate -path migrations -database "sqlite3://database.db" up; \
	fi

db-reset: db-create
	@if [ -f database.db ]; then \
		sqlite3 database.db ".exit"; \
		rm -f database.db; \
		make db-create; \
		make db-seed; \
		make db-migrate; \
	else \
		echo "Database does not exist"; \
	fi

db-rollback:
	@if [ -f database.db ]; then \
		migrate -path migrations -database "sqlite3://database.db" down 1; \
	else \
		echo "Database does not exist"; \
	fi

db-migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

dev-setup:
	@go mod download

db-shell:
	@sqlite3 database.db

db-seed:
	@sqlite3 database.db < seed.sql
	@echo "Seeded database.db"

dev:
	air -c .air.toml & pnpm css:watch