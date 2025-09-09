RUN_DIR=./cmd/server
DOCKER_COMPOSE_PATH=deployments/docker-compose.yml
APP_NAME=ping-battle
DEV_ENV_NAME=ping-battle-dev

.PHONY: run
run:
	go run $(RUN_DIR)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(APP_NAME) $(RUN_DIR)

.PHONY: test
test:
	go test ./... -v

.PHONY: docker-build
docker-build:
	docker build -t $(APP_NAME):latest .

.PHONY: docker-run
docker-run: docker-build
	docker run --rm --env-file .env -p 8080:8080 $(APP_NAME):latest

.PHONY: dev-up
dev-up:
	docker-compose -f $(DOCKER_COMPOSE_PATH) -p $(DEV_ENV_NAME) up -d

.PHONY: dev-down
dev-down:
	docker-compose -f $(DOCKER_COMPOSE_PATH) -p $(DEV_ENV_NAME) down

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: format
format:
	golangci-lint fmt