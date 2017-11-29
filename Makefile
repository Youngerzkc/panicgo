GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")


all:
	go build
	./panicgo

install: deps
	govendor sync

.PHONY: test
test:
	go test -v -covermode=count -coverprofile=coverage.out

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)
