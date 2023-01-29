## up: starts all containers in the background without forcing build
up:
	@echo "Starting docker images..."
	docker-compose -f ./docker-compose.dev.yml up -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker images..."
	docker-compose -f ./docker-compose.dev.yml down
	@echo "Docker stopped!"

## start: start application
start:
	@echo "Stopping docker images..."
	go run ./cmd/app/*
	@echo "Stopping docker images..."
	
## doc: genereting Swagger Docs
doc:
	@echo "Stopping genereting Swagger Docs..."
	swag init -g ./cmd/app/* --output ./docs
	@echo "Swagger Docs prepared, look at /docs"

## help: displays help
help: Makefile
	@echo " Choose a command:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'