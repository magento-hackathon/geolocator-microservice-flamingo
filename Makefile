GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_FOLDER=bin/
APP_NAME=geolocatorservice
BINARY_UNIX=$(APP_NAME)_unix
CONTEXT?=dev
VERSION_NUMBER?=latest

.PHONY: serve

all: test build build-linux
build:
	$(GOBUILD) -o $(BIN_FOLDER)$(APP_NAME) -v
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN_FOLDER)$(BINARY_UNIX) -v
test:
	$(GOTEST) ./... -v
clean:
	$(GOCLEAN)
	rm -f $(BIN_FOLDER)$(APP_NAME)
	rm -f $(BIN_FOLDER)$(BINARY_UNIX)
serve:
	CONTEXT=$(CONTEXT) go run main.go serve
update-flamingo:
	go get flamingo.me/flamingo/v3

