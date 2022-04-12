all: up lint run

build:
	go build github.com/snapp-incubator/errandboi

docker-build:
	go docker build -t errandboi

run:
	go run github.com/snapp-incubator/errandboi serve

up:
	docker-compose up -d

down:
	docker-compose down

lint:
	golangci-lint run
.PHONY: lint
