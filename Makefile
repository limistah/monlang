
GOCMD ?= go
GOBUILD = $(GOCMD) build
GOFMT = gofmt
GOTEST = $(GOCMD) test
BINARYNAME = monlang

default: fmt

all:
	test fmt

test: 
	$(GOTEST) -v ./...

build: 
	$(GOBUILD) -o $(BINARYNAME) ./cmd/monlang/main.go  


fmt:
	$(GOFMT) -w ./..
