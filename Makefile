install:
	go mod download
	go build -o engo-db ./cmd/server
dev: install
	./engo-db