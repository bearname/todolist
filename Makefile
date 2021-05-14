all: test build run

modules:
	go mod tidy

go-build: modules
	go build -o bin/todo ./cmd/todo/.

build:
	docker-compose build

build-docker:
	docker build -t mikhailmi/todolist  .

run:
	docker-compose up

run-docker:
	docker run -p 8000:8000 mikhailmi/todolist

go-test:
	go test ./...

test:
	newman run test/go.postman_collection.json