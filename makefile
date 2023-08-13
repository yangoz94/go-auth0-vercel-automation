.PHONY: run build test

run:
	go run pkg/main.go

build:
	go build -o build/script pkg/main.go

test:
	go test ./...