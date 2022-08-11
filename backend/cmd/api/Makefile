SERVICE_NAME=api
PROJECT_ROOT=$(shell git rev-parse --show-toplevel)
API_ROOT=$(PROJECT_ROOT)/backend/cmd/api

build:
	env GOOS=linux GOARCH=amd64 go build -o $(API_ROOT)