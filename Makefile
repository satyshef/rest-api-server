APP_PATH="apiserver"

.PHONY: build
build:
	go build -v -o ${APP_PATH} ./cmd/apiserver

.DEFAULT_GOAL := build