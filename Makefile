BIN = bin/go-pinger
REGISTRY = abhisheknsit/go-pinger
GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
GIT_HEAD_REF=$(shell git show-ref --hash --verify refs/heads/$(GIT_BRANCH))
IMG_TAG?=$(GIT_HEAD_REF)

.PHONY: build docker


build:
	go build -o $(BIN) ./main.go

docker:
	docker build -t "$(REGISTRY):$(IMG_TAG)" .
	docker tag "$(REGISTRY):$(IMG_TAG)" "$(REGISTRY)"
	
clean:
	rm -rf bin/

push: docker
	docker push "$(REGISTRY):$(IMG_TAG)"
	docker push "$(REGISTRY)"
