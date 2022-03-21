CURRENT_DIR = $(shell pwd)

GREEN 		= \033[0;32m
YELLOW 		= \033[0;33m
NC 			= \033[0m

APP_VERSION = `git describe --tag --abbrev=0`
BUILD_DATE 	= `date -u +%Y%m%d.%H%M%S`

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-19s\033[0m %s\n", $$1, $$2}'

# ============= General use-cases =============

start-up: install dependencies-up migrate run ## Start service from zero
start: dependencies-up migrate run ## UP docker dependensies adn run service
check: linting test ## Linting Golang code and run tests in one command
validate: fmt linting test ## Formatting Linting Golang code and run tests in one command

# ============= General commands =============
migrate: ## Migrate db
	@echo "\n${GREEN}Applying migrations: migrations run on startup${NC}"

install: ## Install / Rebuild the binary
	@echo "\n${GREEN}Install / Rebuild the binary${NC}"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0

dependencies-up: ## Pull and start the Docker containers with dependencies in the background
	@echo "\n${GREEN}Pull and start the Docker containers with dependencies in the background${NC}"
	# docker-compose up -d database
	cmd/accrual/accrual_darwin_amd64 -a :8091

dependencies-down: ## Down the Docker containers with dependencies
	@echo "\n${YELLOW}Down the Docker containers with dependencies${NC}"
	docker-compose down
	docker-compose ps

run: ## Run the application
	@echo "\n${GREEN}Run the application${NC}"
	go run cmd/gophermart/main.go

clean: ## Clear temporary information, stop Docker containers
	@echo "\n${YELLOW}Clear cache directories${NC}"
	go clean
	go mod tidy
	golangci-lint cache clean

test: ## Run unit-tests
	@echo "\n${GREEN}Running unit-tests${NC}"
	go test -race -v -covermode=atomic -coverprofile=coverage.out $$(go list ./... | grep -v cmd)
	# go tool cover -func coverage.out | grep total | awk '{print "" $$3}'

fmt: ## Auto formatting Golang code
	@echo "\n${GREEN}Auto formatting golang code with gofmt${NC}"
	gofmt -w -l $$(go list -f "{{ .Dir }}" ./...); if [ "$${errors}" != "" ]; then echo "$${errors}"; fi
	@echo "\n${GREEN}Auto formatting golang code with golangci-lint${NC}"
	golangci-lint run --fix

linting: golangci-lint ## Linting Golang code

# ============= Other project specific commands =============

golangci-lint: ## Linting Golang code with golangci
	@echo "\n${GREEN}Linting Golang code with golangci${NC}"
	golangci-lint --version
	golangci-lint run ./... -v --timeout 240s
