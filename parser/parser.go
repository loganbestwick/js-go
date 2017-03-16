package parser

import (
	"io"

	"github.com/loganbestwick/js-go/syntax"
)

func createStatementsNode(a yySymType) yySymType {
	node := syntax.StatementsNode{
		Statements: []syntax.Node{a.node},
	}
	return yySymType{node: node}
}

func appendStatementsNode(a yySymType, b yySymType) yySymType {
	sn := a.node.(syntax.StatementsNode)
	node := syntax.StatementsNode{
		Statements: sn.Statements,
	}
	node.Statements = append(node.Statements, b.node)
	return yySymType{node: node}
}

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

func createVariableNode(o yySymType) yySymType {
	node := syntax.VariableNode{
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

func Parse(reader io.Reader) syntax.Node {
	lexer := NewLexer(reader)
	yyParse(lexer)
	return lexer.parseResult.(yySymType).node
}

func setParseResult(lexer interface{}, o yySymType) {
	lexer.(*Lexer).parseResult = o
}
