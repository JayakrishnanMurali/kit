BINARY_NAME=kit

GO=go
GO_BUILD=$(GO) build
GO_MOD=$(GO) mod
GOFMT=$(GO) fmt

all: build

build:
	$(GO_BUILD) -o $(BINARY_NAME)

fmt:
	$(GOFMT) ./...

test:
	$(GO) test ./...

clean:
	rm -f $(BINARY_NAME)

run:
	./$(BINARY_NAME) $(args)

.PHONY: all build fmt test clean run
