## up: starts all containers in the background without forcing build
up:
	@echo "Starting docker images..."
	docker-compose -f ./docker-compose.dev.yml --env-file ./.env.dev up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker images..."
	docker-compose -f ./docker-compose.dev.yml --env-file ./.env.dev down
	@echo "Docker stopped!"

## post: stops post-service, removes docker image, builds service, and starts it
post: build
	@echo "Building post-service docker image..."
	docker-compose -f ./docker-compose.dev.yml --env-file ./.env.dev stop post-service
	docker-compose -f ./docker-compose.dev.yml --env-file ./.env.dev rm -f post-service
	docker-compose -f ./docker-compose.dev.yml --env-file ./.env.dev up --build -d post-service
	docker-compose -f ./docker-compose.dev.yml --env-file ./.env.dev start post-service
	@echo "post-service built and started!"

## build: builds the post-service binary as a linux executable
build:
	@echo "Building post-service binary.."
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o post-service.bin ./cmd/app
	@echo "Authentication post-service built!"

## start: start application
start:
	go run ./cmd/app/*
	
## doc: generating Swagger Docs
doc:
	@echo "Stopping generating Swagger Docs..."
	swag init -g ./cmd/app/* --output ./docs
	@echo "Swagger Docs prepared, look at /docs"

## init: run this command once for prepare MySQL database
init:
	@echo "Creating Database schema..."
	docker exec -i mysql mysql -uroot -proot < migrations/init-001.sql
	docker-compose -f ./docker-compose.dev.yml --env-file ./.env.dev restart post-service



## help: displays help
help: Makefile
	@echo " Choose a command:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'