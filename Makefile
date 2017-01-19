setup:
	go get github.com/blynn/nex

gen:
	go tool yacc -o=parser/parser.gen.go parser/parser.y
	nex -o parser/lexer.gen.go parser/lexer.nex

build: gen
	go build

clean:
	-rm y.output
