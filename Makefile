REPO_NAME ?= docker-golang
REPO_URI ?= github.com/whyistilley/${REPO_NAME}
BUILD_DATE ?= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT ?= $(shell git rev-parse HEAD)
GIT_VERSION ?= $(shell git describe --tags --abbrev=0 | tr -d '\n')
IMAGE_NAME ?= ${REPO_NAME}
IMAGE_TAG ?= ${REPO_URI}:${GIT_VERSION}-${GIT_COMMIT}
DOCKERFILE ?= ./_build/Dockerfile
COMPOSE_FILE ?= ./_deploy/docker-compose.yml
LDFLAGS ?= "-X '${REPO_URI}/internal/service.VersionNumber=${GIT_VERSION}' -X '${REPO_URI}/internal/service.CommitHash=${GIT_COMMIT}' -X '${REPO_URI}/internal/service.BuildDateTime=${BUILD_DATE}'"

include ./.env

.PHONY: build run clean \
 	docker-build docker-run docker-logs docker-remove docker-remove-image docker-clean docker-export docker-remove-export \
 	compose-build compose-run compose-up compose-down compose-remove compose-clean

# To find symbols to overwrite, run the following command: go tool nm ./app | grep app
build-flags: clean
	go build -v -p 1 -race -o ./bin/ -ldflags=${LDFLAGS} ./...

build: clean
	go build -v -p 1 -race -o ./bin/ ./...

run:
	POSTGRES_URI=${POSTGRES_URI} ./bin/app

clean:
	rm -rf ./bin ./coverage.out

test:
	go test -v -p 1 -race -bench=. -benchmem -cover -coverprofile=coverage.out -failfast ./...

test-flags:
	go test -v -p 1 -race -bench=. -benchmem -cover -coverprofile=coverage.out -failfast -ldflags="${LDFLAGS}" ./...

docker-build:
	docker build --build-arg LDFLAGS="${LDFLAGS}" --build-arg TEST_LDFLAGS="${TEST_LDFLAGS}" --tag ${IMAGE_TAG} --file ${DOCKER_FILE} --no-cache .

docker-run:
	docker container run --name ${REPO_NAME} ${IMAGE_TAG}

docker-logs:
	docker container logs --follow --details --timestamps ${REPO_NAME}

docker-remove:
	docker container rm --volumes --force ${REPO_NAME}

docker-remove-image:
	docker image rm --force ${IMAGE_TAG}

docker-clean: docker-remove docker-remove-image

docker-export:
	docker container export --output="./${REPO_NAME}.tar" ${REPO_NAME}

docker-remove-export:
	rm -f ${REPO_NAME}.tar

compose-build:
	docker compose --file ${COMPOSE_FILE} build --no-cache

compose-run:
	docker compose --file ${COMPOSE_FILE} run --rm db

compose-up: compose-build
	docker compose --file ${COMPOSE_FILE} up --build --force-recreate --detach

compose-logs:
	docker compose --file ${COMPOSE_FILE} logs --follow --timestamps

compose-down:
	docker compose --file ${COMPOSE_FILE} down --volumes --rmi="local"

compose-remove:
	docker compose --file ${COMPOSE_FILE} rm --force --stop --volumes

compose-clean: compose-remove compose-down
