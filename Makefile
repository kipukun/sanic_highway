BINARY = highway
BUILD_DIR = cmd/highway

all: gen build run
build:
	go build ./$(BUILD_DIR)
gen:
	go generate ./$(BUILD_DIR)
run:
	./$(BINARY)

.PHONY: gen build run
