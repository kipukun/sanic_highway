all: run
build:
	go build cmd/highway/main.go
run:
	go build cmd/highway/main.go
	./cmd/highway/highway
