
GOCMD ?= go
GOBUILD = $(GOCMD) build
GOFMT = gofmt
GOTEST = $(GOCMD) test


default: fmt

all:
	test fmt

test: 
	$(GOTEST) -v ./...

build-cmd: 
	$(GOBUILD) -o monlang ./cmd/monlang/main.go  


fmt:
	$(GOFMT) -w ./..
