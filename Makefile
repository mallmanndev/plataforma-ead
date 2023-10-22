# Define colors based on tput
BOLD := $(shell tput bold)
RED := $(shell tput -Txterm setaf 1)
GREEN := $(shell tput -Txterm setaf 2)
BLUE := $(shell tput -Txterm setaf 6)
RESET := $(shell tput -Txterm sgr0)

.PHONY: help up restart logs down psql run-service-core-tests generate-service-core-protoc

help: ## Shows help for Makefile commands.
	@echo ""
	@echo "$(BOLD)$(BLUE)Available commands:$(RESET)"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9 -]+:.*?## / {printf "$(BOLD)$(GREEN)%-30s$(RESET)%s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

up: # Start Docker Compose services in the background
	@echo "$(BOLD)$(BLUE)Starting Docker Compose services...$(RESET)"
	@docker compose up -d && echo "$(GREEN)Docker Compose started successfully$(RESET)" || echo "$(RED)Docker Compose failed to start$(RESET)"

restart: ## Restart docker compose services
	@echo "$(BOLD)$(BLUE)Restarting Docker Compose services...$(RESET)"
	@docker compose restart && echo "$(GREEN)Docker Compose restarted successfully$(RESET)" || echo "$(RED)Docker Compose failed to restart$(RESET)"

log: ## Display logs from the service-core container
	@echo "$(BOLD)$(BLUE)Displaying logs from service-core container...$(RESET)"
	@docker logs -f service-core || echo "$(RED)Failed to display logs from service-core container$(RESET)"

down: ## Stop and remove docker compose services
	@echo "$(BOLD)$(BLUE)Stopping Docker Compose services...$(RESET)"
	@docker compose down && echo "$(GREEN)Docker Compose stopped successfully$(RESET)" || echo "$(RED)Docker Compose failed to stop$(RESET)"

psql: ## Open a psql session in the service-core-db container
	@echo "$(BOLD)$(BLUE)Opening psql session in service-core-db container...$(RESET)"
	@docker exec -it service-core-db psql -U postgres service-core || echo "$(RED)Failed to open psql session in service-core-db container$(RESET)"

run-service-core-tests: ## Run tests in the service-core directory
	@echo "$(BOLD)$(BLUE)Running tests in service-core directory...$(RESET)"
	@cd service-core && go test ./... || echo "$(RED)Failed to run tests in service-core directory$(RESET)"

generate-service-core-protoc: ## Generate Go code from Protobuf files
	@echo "$(BOLD)$(BLUE)Generating Go code from Protobuf files...$(RESET)"
	@protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto && echo "$(GREEN)Generated Protobuf files successfully$(RESET)" || echo "$(RED)Failed to generate Protobuf files$(RESET)"