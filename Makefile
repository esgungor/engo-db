install:
	go mod download
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/comm.proto
server: install
	go build -o server ./cmd/server

client: install
	go build -o cli ./cmd/client

