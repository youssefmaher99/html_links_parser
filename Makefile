build:
	go build -o main main.go

run: build
	./main

test:
	go test -v ./...