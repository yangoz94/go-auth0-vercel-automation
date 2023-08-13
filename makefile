.PHONY: run build test

run:
	go run cmd/api/main.go

build:
	go build -o build/script cmd/api/main.go

test:
	go test ./...