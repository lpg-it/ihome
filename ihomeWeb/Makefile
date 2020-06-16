
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:

	go build -o ihomeWeb-web main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t ihomeWeb-web:latest
