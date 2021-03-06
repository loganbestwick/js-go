.PHONY: all precommit setup test build run fix gen

setup:
	go get github.com/blynn/nex

clean:
	-rm y.output

gen: clean
	goyacc -o=parser/parser.gen.go parser/parser.y
	nex -o parser/lexer.gen.go parser/lexer.nex

fix:
	gofmt -w -l .
	goimports -w -l .

test:
	go test ./...

precommit: gen fix test

build: clean gen test
	go build

all: clean setup build
