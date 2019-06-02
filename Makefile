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
container: build-linux containerize
containerize:
	docker build -t magento-hackathon/$(APP_NAME):$(VERSION_NUMBER) .
docker-run:
	docker run \
		--rm \
		-p 3322:3322 \
		agento-hackathon/$(APP_NAME):latest
docker-push:
	echo "$(DOCKER_PASS)" | docker login -u "$(DOCKER_USER)" --password-stdin
	docker push agento-hackathon/$(APP_NAME)
