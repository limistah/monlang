GOCMD ?= go
GOBUILD = $(GOCMD) build
GOFMT = gofmt
INTERPRETER_BINARY_NAME = monlang
COVEROUT=cover.out


UNAME := $(shell uname -m)
ifeq ($(UNAME), s390x)
# go test does not support -race flag on s390x architecture
	RACE=
else
	RACE=-race
endif

GOTEST_QUIET=$(GOCMD) test $(RACE)
GOTEST=$(GOTEST_QUIET) -v

default: fmt \
	test

all:
	test fmt

.PHONY: test
test:
	$(GOTEST_QUIET)  ./...

test-ci: test cover

.PHONY: cover
cover:
	bash -c "set -e; set -o pipefail; STORAGE=memory $(GOTEST) -timeout 5m -coverprofile $(COVEROUT) ./... | tee test-results.json"
	go tool cover -html=$(COVEROUT) -o cover.html

build-all: build-interpreter

.PHONY: build-interpreter
build-interpreter:
	$(GOBUILD) -o $(INTERPRETER_BINARY_NAME) ./cmd/monlang/main.go

fmt:
	$(GOFMT) -w ./..

start-repl:
	$(GOCMD) run ./cmd/monlang/main.go

