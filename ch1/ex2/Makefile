.DEFAULT_GOAL := build

.PHONY:fmt vet build clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: clean vet
	go build

clean:
	go clean