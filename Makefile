.PHONY: up
up:
	docker compose up -d

.PHONY: restart
restart: 
	docker compose restart

.PHONY: log
log:
	docker logs -f service-core

.PHONY: down
down:
	docker compose down

run-service-core-tests:
	cd service-core && go test ./...

generate-service-core-protoc:
	protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto