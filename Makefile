install:
	go mod download
	go build 
dev: install
	./engo-db