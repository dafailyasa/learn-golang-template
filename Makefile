DB_URL = mysql://root:dafa@tcp(localhost:3306)/golang-template?charset=utf8mb4&parseTime=True&loc=Local
APP_NAME = golang-template
BUILD_DIR = $(PWD)/build

server-dev: 
	go run cmd/http/main.go
lint: 
	golangci-lint run
test: 
	go test -v -cover ./...
test-coverage: 
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out && rm coverage.out
migrateup:
	migrate -path database/migrations -database "$(DB_URL)" -verbose up
migrateup1:
	migrate -path database/migrations -database "$(DB_URL)" -verbose up 1
migratedown: 
	migrate -path database/migrations -database "$(DB_URL)" -verbose down
migratedown1: 
	migrate -path database/migrations -database "$(DB_URL)" -verbose down 1
build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) cmd/http/main.go
clean-build:
	rm -rf ./build

.PHONY: server-dev lint test test-coverage migrateup migrateup1 migratedown migratedown1 build clean-build