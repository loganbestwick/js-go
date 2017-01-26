package parser

import (
	"github.com/loganbestwick/js-go/syntax"
	"io"
)

func createValueNode(o yySymType) yySymType {
	node := syntax.ValueNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createAddNode(left yySymType, right yySymType) yySymType {
	node := syntax.AddNode{
		Left:  left.node,
		Right: right.node,
	}
	return yySymType{node: node}
}

func Parse(reader io.Reader) syntax.Node {
	lexer := NewLexer(reader)
	yyParse(lexer)
	return lexer.parseResult.(yySymType).node
}

func setParseResult(lexer interface{}, o yySymType) {
	lexer.(*Lexer).parseResult = o
}
