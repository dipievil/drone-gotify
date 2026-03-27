.PHONY: build test tidy run

build:
	CGO_ENABLED=0 GOOS=linux go build -v -a -o release/linux/amd64/drone-gotify .

test:
	go test -v -cover ./...

tidy:
	go mod tidy

run: tidy build
	./release/linux/amd64/drone-gotify
