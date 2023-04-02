MIGRATIONS_PATH = migrations/
DATABASE_URI = postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: run
run:
	go run cmd/main.go

.PHONY: migrate.up
migrate.up:
	migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URI) up

.PHONY: migrate.down
migrate.down:
	migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URI) down