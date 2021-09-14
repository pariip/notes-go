# Basic go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=notes
BINARY_LINUX=notes-go_linux

# Path
MAIN_PATH=./cmd/notes/main.go

go-all: go-test go-build

go-build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
go-test:
	$(GOTEST) -v ./...
go-clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)
go-run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
	./$(BINARY_NAME)