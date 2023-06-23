all: deps swagger

deps:
	go mod tidy
	go mod vendor

run:
	go run cmd/app/main.go

swagger: 
	swag init --parseInternal --parseDependency --parseVendor --parseDepth 3 -o ./api -g ./cmd/app/main.go