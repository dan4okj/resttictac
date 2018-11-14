GOCMD=go

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=resttictac

all: test build
build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./src/main.go
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_NAME)
run: build
	./bin/$(BINARY_NAME)
deps:
	$(GOGET) github.com/gorilla/mux

