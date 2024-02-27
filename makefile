.PHONY: default run build test docs clean

APP_NAME=application-openings

default: run

install:
	@go mod tidy
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
run:
	@go run cmd/application-openings/main.go
docs:
	@swag init -g ./cmd/application-openings/main.go -o ./api/
clean:
	@rm -f @(APP_NAME)
	@rm -rf ./api
