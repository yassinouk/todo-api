build:
	@go build -o bin/todo-api
run: build
	@./bin/todo-api
test:
	@go test -v ./...
.PHONY: build run test