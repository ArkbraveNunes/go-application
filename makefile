.PHONY: default run build test docs clean

APP_NAME=application-openings

default: run

run:
	@go run cmd/application-openings/main.go
install:
	@go mod tidy
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init ./cmd/application-openings/main.go -o ./api
clean:
	@rm -f @(APP_NAME)
	@rm -rf ./api
