RUN_DIR=./cmd/server
DOCKER_COMPOSE_PATH=deployments/docker-compose.yml
DEV_ENV_NAME=ping-battle

run:
	go run $(RUN_DIR)

start-dev-env:
	docker-compose -f $(DOCKER_COMPOSE_PATH) -p $(DEV_ENV_NAME) up -d

stop-dev-env:
	docker-compose -f $(DOCKER_COMPOSE_PATH) -p $(DEV_ENV_NAME) down