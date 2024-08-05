
GOCMD ?= go
GOBUILD = $(GOCMD) build
GOFMT = gofmt
GOTEST = $(GOCMD) test


default: fmt

all:
	test fmt

test: 
	$(GOTEST) -v ./...


fmt:
	$(GOFMT) -w ./..
