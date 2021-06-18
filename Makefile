SHELL := /bin/bash
current_dir = $(shell pwd)

all: test build itest

test:
	go fmt
	go test -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

build:
	go mod download
	go build

itest: build
	INPUT_FILE=itest/input.yml OUTPUT_FILE=itest/app-config-nginx.conf ./nginx-config-generator
	docker run -t -v ${current_dir}:/app --entrypoint '/app/itest/itest.sh' nginx:stable-alpine

container:
	docker build . -t nginx-config-generator:$(shell git rev-parse --short HEAD)
