
.PHONY: setup
setup:
	go install github.com/a-h/templ/cmd/templ@v0.3.865

.PHONY: test
test:
	test -z $(go fmt ./...) && go test -v ./... && staticcheck ./... && gosec ./...

.PHONY: build
build: test
	templ generate && go build -o goffeine .

.PHONY: run
run:
	go run .
