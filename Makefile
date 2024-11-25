.PHONY: test

test:
	go test -v ./...

build:
	templ generate && go build -o goffeine .

run:
	go run .