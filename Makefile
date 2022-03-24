build:
	go build -o dist/main

unittest:
	go test -v ./...

buildimage: 
	docker build --platform=linux/amd64 -t zhaoyi0113/es-index-scheduler .

publishimage:
	docker push zhaoyi0113/es-index-scheduler
