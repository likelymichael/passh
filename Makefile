.PHONY: test
test:
	go test ./... -cover

bin:
	mkdir bin

build:
	go build -o passh .
