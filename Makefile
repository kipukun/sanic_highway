BINARY = highway
BUILD_DIR = cmd/highway

all: build run
build:
	go build ./$(BUILD_DIR)
run:
	./$(BINARY)

.PHONY: build run
