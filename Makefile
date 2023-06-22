all: deps

deps:
	go mod tidy
	go mod vendor

run:
	go run cmd/main.go