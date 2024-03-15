migrate:
	@echo "Running migration script..."
	@bash scripts/bash/migrate.sh

build:
	@echo "Building docker containers..."
	@docker-compose build

up:
	@echo "Starting docker containers..."
	@docker-compose up -d

down:
	@echo "Stopping docker containers..."
	@docker-compose down

remove:
	@echo "Removing docker containers..."
	@docker-compose down --rmi all

logs:
	@echo "Fetching docker containers logs..."
	@docker-compose logs
