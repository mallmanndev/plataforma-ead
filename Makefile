.PHONY: up
up:
	docker compose up -d

.PHONY: build
build:
	docker compose up -d --build

.PHONY: restart
restart: 
	docker compose restart

.PHONY: logs
logs:
	docker logs -f service-core

.PHONY: down
down:
	docker compose down

.PHONY: psql
psql:
	docker exec -it service-core-db psql -U postgres service-core

.PHONY: bash
bash:
	docker exec -it service-core /bin/bash


servcore-tests:
	docker exec -it service-core go test ./...

generate-service-core-protoc:
	protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto