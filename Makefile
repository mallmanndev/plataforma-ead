.PHONY: up
up:
	docker compose up -d

.PHONY: build
build:
	docker compose up -d --build

.PHONY: restart
restart: 
	docker compose restart

.PHONY: servcore-logs
servcore-logs:
	docker logs -f service-core

.PHONY: servcourse-logs
servcourse-logs:
	docker logs -f service-course

.PHONY: down
down:
	docker compose down

.PHONY: psql
psql:
	docker exec -it service-core-db psql -U postgres service-core

.PHONY: servcore-bash
servcore-bash:
	docker exec -it service-core /bin/bash

.PHONY: servcourse-bash
servcourse-bash:
	docker exec -it service-course /bin/bash

.PHONY: servcore-test
servcore-test:
	docker exec -it service-core go test ./...

.PHONY: servcourse-test
servcourse-test:
	docker exec -it service-course go test ./...

.PHONY: servcourse-coverage
servcourse-coverage:
	docker exec -it service-course go test -coverprofile=coverage.out ./...
	docker exec -it service-course go tool cover -html=coverage.out -o coverage.html

.PHONY: servcore-protoc
servcore-protoc:
	protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto

.PHONY: servcourse-protoc
servcourse-protoc:
	protoc --go_out=./service-course --go-grpc_out=./service-course ./service-course/protos/*.proto
