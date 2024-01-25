.PHONY: build run test clean

build: fmt
	go build -o bin/hey main.go

run: build
	./bin/hey

test:
	go test ./...

clean:
	go clean
	rm -rf bin/

fmt:
	go fmt ./...
