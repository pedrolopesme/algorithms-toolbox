GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/stack
BINARY_UNIX=$(BINARY_NAME)_unix

test:
	@echo "Running ... tests"
	$(GOTEST) -v ./...

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w *.go