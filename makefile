# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

# Name of the binary to be built
BINARY_NAME=ChiApplication


build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v .
	./bin/$(BINARY_NAME)
