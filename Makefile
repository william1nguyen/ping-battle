RUN_DIR=./cmd/server

run:
	go run $(RUN_DIR)

start-dev-env:
	docker-compose -f deployments/docker-compose.yml -p ping-battle up -d --build

stop-dev-env:
	docker-compose -f deployments/docker-compose.yml -p ping-battle down