# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
SERVER_BINARY=./build/server
SERVER_BINARY_UNIX=./build/$(SERVER_BINARY)_unix
TELEGRAM_BINARY=./build/telegrambot
TELEGRAM_BINARY_UNIX=./build/$(TELEGRAM_BINARY)_unix

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	$(GOCMD) mod tidy -v
	$(GOCMD) fmt ./...

.PHONY: build-server
build-server: tidy
	$(GOBUILD) -o $(SERVER_BINARY) ./cmd/server

.PHONY: build-telegram
build-telegram:
	$(GOBUILD) -o $(TELEGRAM_BINARY) ./cmd/telegrambot

.PHONY: run-server
run-server: build-server
	./$(SERVER_BINARY)

.PHONY: run-telegram
run-telegram: build-telegram
	./$(TELEGRAM_BINARY)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(SERVER_BINARY) $(SERVER_BINARY_UNIX) $(TELEGRAM_BINARY) $(TELEGRAM_BINARY_UNIX)

.PHONY: build-linux-server
build-linux-server:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(SERVER_BINARY_UNIX) ./cmd/server

.PHONY: build-linux-telegram
build-linux-telegram:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(TELEGRAM_BINARY_UNIX) ./cmd/telegrambot