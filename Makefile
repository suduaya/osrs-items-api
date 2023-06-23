BINARY_DIR=bin
BINARY_NAME=app
GOCMD=go
SWAGCMD=swag init

all: deps test build swagger

deps:
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

run:
	$(GOCMD) run cmd/app/main.go

build:
	$(GOCMD) build -mod vendor -o $(BINARY_DIR)/$(BINARY_NAME) ./cmd/app

swagger: 
	$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest
	$(SWAGCMD) --parseInternal --parseDependency --parseVendor --parseDepth 3 -o ./api -g ./cmd/app/main.go

test:
	$(GOCMD) test -v ./...