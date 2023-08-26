run-grpc-server:
	cd service-core && go run cmd/grpc-server/main.go

run-service-core-tests:
	cd service-core && go test ./...

generate-service-core-protoc:
	protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto