debug:
	dlv debug --headless --listen localhost:1234 .

build:
	go build .

static:
	CGO_ENABLED=0 go build . -a

run:
	go run .

init:
	go run . --init