build:
	go build -o dist/main

unittest:
	go test -v ./...
