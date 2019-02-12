BINARY = highway
BUILD_DIR = cmd/highway

all: build gen run
build:
	go build ./$(BUILD_DIR)
gen:
	qtc -dir=templates
run:
	./$(BINARY)

.PHONY: build gen run
