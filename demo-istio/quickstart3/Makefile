gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb
server:
	go run main.go -port 8091

.PHONY: gen server