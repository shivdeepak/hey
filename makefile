.PHONY: build run test clean

build: fmt tidy
	go build -o bin/hey main.go

tidy:
	go mod tidy

run: build
	./bin/hey

test:
	go test ./...

clean:
	go clean
	rm -rf bin/

fmt:
	go fmt ./...
