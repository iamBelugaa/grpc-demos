tidy:
	@go mod tidy

gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/hello.proto

build:
	@go build -o bin/main main.go

run: build
	@./bin/main