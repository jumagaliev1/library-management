MIGRATIONS_PATH = migrations/
DATABASE_URI = postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
DOCKER_COMPOSE_PATH = build/docker-compose.yml
DOCKER_FILE_PATH = build/Dockerfile

.PHONY: run
run:
	go run cmd/main.go

.PHONY: migrate.up
migrate.up:
	migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URI) up

.PHONY: migrate.down
migrate.down:
	migrate -path $(MIGRATIONS_PATH) -database $(DATABASE_URI) down

.PHONY: docker.compose
docker.compose:
	docker compose -f $(DOCKER_COMPOSE_PATH) up --build app

.PHONY: docker.build
docker.build:
	docker build -f $(DOCKER_FILE_PATH) .