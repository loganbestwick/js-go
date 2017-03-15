package parser

import (
	"io"

	"github.com/loganbestwick/js-go/syntax"
)

func createNumberNode(o yySymType) yySymType {
	node := syntax.NumberNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createStringNode(o yySymType) yySymType {
	node := syntax.StringNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createBinaryOpNode(operator yySymType, left yySymType, right yySymType) yySymType {
	node := syntax.BinaryOpNode{
		Left:     left.node,
		Right:    right.node,
		Operator: operator.s,
	}
	return yySymType{node: node}
}

func createAssignmentNode(left yySymType, right yySymType) yySymType {
	node := syntax.AssignmentNode{
		Left: left.s,
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
