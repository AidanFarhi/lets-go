.DEFAULT_GOAL := build

.PHONY:fmt vet build
clean:
	go clean ./...

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet clean
	go build
