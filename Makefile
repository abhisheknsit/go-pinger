BIN = bin/go-pinger
REGISTRY = abhisheknsit/go-pinger
GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
IMG_TAG?=$(GIT_BRANCH)

.PHONY: build docker


build:
	go build -o $(BIN) ./main.go

docker:
	docker build -t "$(REGISTRY):$(IMG_TAG)" .
	
clean:
	rm -rf bin/
