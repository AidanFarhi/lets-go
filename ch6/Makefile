.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	staticcheck ./...
.PHONY:lint

vet: fmt
	go vet ./...
	# uncomment the below to enable shadowing detection
	# shadow ./...
.PHONY:vet

build: vet
	go build main.go
.PHONY:build
