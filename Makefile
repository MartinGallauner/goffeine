
.PHONY: setup
setup:
	go install github.com/a-h/templ/cmd/templ@v0.3.924
	npm install tailwindcss @tailwindcss/cli

.PHONY: test
test:
	test -z $(go fmt ./...) && go test -v ./... && staticcheck ./... && gosec ./...

.PHONY: build
build:
	templ generate && npx @tailwindcss/cli -i ./assets/src/style.css -o ./assets/dist/style.css && go build -o bin/goffeine ./cmd

.PHONY: run
run: build
	./bin/goffeine

.PHONY: hotreload
hotreload:
	templ generate --watch --proxy="http://localhost:8081" --cmd="go run ./cmd"
