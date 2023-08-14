.PHONY: run build test

run:
	go run cmd/api/main.go

build:
	go build -o build/application cmd/api/main.go

test:
	go test ./...