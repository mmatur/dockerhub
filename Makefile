BINARY_NAME = dockerhub

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse --short HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

export GO111MODULE=on

GOFILES := $(shell git ls-files '*.go' | grep -v '^vendor/')

default: clean check test build

$(DIST_DIR):
	mkdir -p $(DIST_DIR)

dependencies:
	go mod download

clean:
	rm -rf dockerhub

test: clean
	go test -v -cover ./...

build:
	CGO_ENABLED=0 go build -ldflags="-s -w \
	-X github.com/mmatur/$(BINARY_NAME)/cmd/version.version=$(VERSION) \
	-X github.com/mmatur/$(BINARY_NAME)/cmd/version.commit=$(SHA) \
	-X github.com/mmatur/$(BINARY_NAME)/cmd/version.date=$(BUILD_DATE)" \
	$(CURDIR)/

check:
	golangci-lint run

fmt:
	@gofmt -s -l -w $(GOFILES)

imports:
	@goimports -w $(GOFILES)

.PHONY: clean check test build dependencies fmt imports
