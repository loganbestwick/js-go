setup:
	go get github.com/blynn/nex

gen:
	go tool yacc -o=parser/gen_parser.go parser/parser.y
	nex -o parser/gen_lexer.go parser/lexer.nex

build: gen
	go build

clean:
	-rm y.output
