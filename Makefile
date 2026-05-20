APP_NAME=book-scraper

run:
	go run ./cmd/scraper

test:
	go test ./...

lint:
	golangci-lint run

build:
	go build -o bin/$(APP_NAME) ./cmd/scraper

docker-build:
	docker build -t $(APP_NAME):latest .