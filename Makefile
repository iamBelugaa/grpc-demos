tidy:
	@go mod tidy

gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/hello.proto

build-server: gen-proto
	@go build -o bin/server/main cmd/server/main.go

run-server: build-server
	@./bin/server/main

build-client:
	@go build -o bin/client/main cmd/client/main.go

run-client: build-client
	@./bin/client/main