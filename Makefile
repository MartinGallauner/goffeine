.PHONY: test

test:
	go test -v ./...
	staticcheck ./...
	gosec ./...

build: test
	templ generate && go build -o goffeine .

run:
	go run .