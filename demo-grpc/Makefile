serv1:
	go run main.go -port 50051

serv2:
	go run main.go -port 50052

serv:
	go run main.go

client:
	go run client/client.go

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=demo_gprc_doc \
    proto/*.proto
	statik -f -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 8088 -r repl

tools:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: serv proto evans tools serv1 serv2 client