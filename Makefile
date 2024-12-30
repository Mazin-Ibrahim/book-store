include .$(PWD)/.env
MIGRATIONS_PATH = ./internal/adapter/storage/postgres/migrations

.PHONY: test
test:
	@go test -v ./...


.PHONY: migration
migration:
	@if [ "$(filter-out $@,$(MAKECMDGOALS))" = "" ]; then \
		echo "Error: Migration name is required. Usage: make migration <name>"; \
		exit 1; \
	fi
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))



.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database="postgres://admin:adminpassword@localhost:5432/book?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY:	db-seed
db-seed:
	@go run /cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag

# Prevent Make from interpreting extra arguments as Makefile targets
%:
	@:
