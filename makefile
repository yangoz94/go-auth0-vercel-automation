.PHONY: run build test

run:
	go run cmd/main.go

build:
	go build -o build/script cmd/main.go

test:
	go test ./...