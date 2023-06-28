BIN := "./bin/cbrwsdltojson"
DOCKER_IMG="cbrwsdltojson:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/cbrwsdltojson

run: build
	$(BIN) -config ./configs/config.env > cbrwsdltojsonCLog.log

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

stop-img: 
	docker stop $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/... 

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

up:
	docker-compose -f ./deployments/docker-compose.yaml up --build > deployLog.log

down:
	docker-compose -f ./deployments/docker-compose.yaml down

integration-tests:
	docker-compose -f ./deployments/docker-compose.yaml -f ./deployments/docker-compose.test.yaml up --build --exit-code-from integration_tests && \
	docker-compose -f ./deployments/docker-compose.yaml -f ./deployments/docker-compose.test.yaml down > deployIntegrationTestsLog.log

.PHONY:  build run build-img run-img stop-img version test lint up down integration-tests 
