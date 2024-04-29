server:
	@go build -o grpc_server server/main.go
	@./grpc_server

client:
	@go build -o grpc_client client/main.go
	@./grpc_client

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protofiles/*.proto

.PHONY: server client proto