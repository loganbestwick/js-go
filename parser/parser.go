package parser

import (
	"io"
	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	s string
}

func node(o yySymType) yySymType {
	return yySymType{
		node: Node{s: o.s},
	}
}

func Parse(reader io.Reader) {
	lexer := NewLexer(reader)
	yyParse(lexer)
	spew.Dump(lexer)
}

func setParseResult(lexer interface{}, o yySymType) {
	lexer.(*Lexer).parseResult = o
}
