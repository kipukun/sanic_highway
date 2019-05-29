BINARY = highway
BUILD_DIR = cmd/highway

all: gen build run
build:
	go build ./$(BUILD_DIR)
gen:
	qtc -dir=templates
run:
	./$(BINARY)

.PHONY: gen build run
