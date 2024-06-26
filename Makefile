.PHONY: up
up:
	docker compose up -d

.PHONY: build
build:
	docker compose up -d --build

.PHONY: restart
restart: 
	docker compose restart

.PHONY: backend-logs
backend-logs:
	docker logs -f backend

.PHONY: front-logs
front-logs:
	docker logs -f front

.PHONY: down
down:
	docker compose down

.PHONY: up-prod
up-prod:
	docker compose -f docker-compose.prod.yml up -d

.PHONY: build-prod
build-prod:
	docker compose -f docker-compose.prod.yml up -d --build

.PHONY: restart-prod
restart-prod:
	docker compose -f docker-compose.prod.yml restart

.PHONY: backend-bash
backend-bash:
	docker exec -it service-course /bin/bash

.PHONY: front-bash
front-bash:
	docker exec -it front /bin/bash

.PHONY: servcourse-test
servcourse-test:
	docker exec -it service-course go test ./...

.PHONY: backend-swag
backend-swag:
	docker exec -it service-course swag init -g ./cmd/rest/main.go

.PHONY: servcourse-coverage
servcourse-coverage:
	docker exec -it service-course go test -coverprofile=coverage.out ./...
	docker exec -it service-course go tool cover -html=coverage.out -o coverage.html

.PHONY: servcourse-protoc
servcourse-protoc:
	protoc --go_out=./service-course --go-grpc_out=./service-course ./service-course/protos/*.proto
