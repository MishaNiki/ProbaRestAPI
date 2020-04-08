.PHONY: build

build: 
	go build -o apiserver.exe -v ./apiserver
	go build -o client.exe -v ./client


.PHONY: test

test:
	go test -v -race -timeout 30s ./...
	
.DEFAULT_GOAL := build