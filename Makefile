debug:
	dlv debug --headless --listen localhost:1234 .

build:
	go build .

run:
	go run .

init:
	go run . --init