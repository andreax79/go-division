BASENAME=`basename $$(pwd)`
IMAGE_NAME=${BASENAME}
VERSION=`cat VERSION`
OWNER=andreax79

.PHONY: help build test image push all

help:
	@echo "- make build        Build"
	@echo "- make test         Run tests"
	@echo "- make image        Build docker image"
	@echo "- make push         Push docker image"

.DEFAULT_GOAL := help

build:
	@go build -o division -ldflags "-X main.version=${VERSION}"

test:
	@go test

image:
	@DOCKER_BUILDKIT=1 docker build \
		 --tag ${IMAGE_NAME}:latest \
		 --tag ${IMAGE_NAME}:${VERSION} \
		 .
push:
	docker tag ${IMAGE_NAME}:${VERSION} ghcr.io/${OWNER}/${IMAGE_NAME}:${VERSION}
	docker tag ${IMAGE_NAME}:${VERSION} ghcr.io/${OWNER}/${IMAGE_NAME}:latest
	docker push ghcr.io/${OWNER}/${IMAGE_NAME}:${VERSION}
	docker push ghcr.io/${OWNER}/${IMAGE_NAME}:latest

all: build
