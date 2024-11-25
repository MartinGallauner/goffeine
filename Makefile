.PHONY: test

test:
	go test -v ./...

build:
	go build -o goffeine .

run:
	go run .