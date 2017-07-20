package parser

import (
	"fmt"
	"io"

	"github.com/loganbestwick/js-go/syntax"
)

func createNumberNode(o yySymType) yySymType {
	node := syntax.NumberNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createBooleanNode(o yySymType) yySymType {
	node := syntax.BooleanNode{
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

func createIdentifierNode(o yySymType) yySymType {
	node := syntax.IdentifierNode{
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

func createIfNode(expr yySymType, statements yySymType) yySymType {
	statementsNode := statements.node.(*syntax.StatementsNode)
	node := syntax.IfNode{
		Expression: expr.node,
		Statements: statementsNode,
	}
	return yySymType{node: node}
}

func createForNode(initExpr yySymType, condition yySymType, loopExpr yySymType, statements yySymType) yySymType {
	statementsNode := statements.node.(*syntax.StatementsNode)
	node := syntax.ForNode{
		InitExpression: initExpr.node,
		Condition:      condition.node,
		LoopExpression: loopExpr.node,
		Statements:     statementsNode,
	}
	return yySymType{node: node}
}

func createWhileNode(expr yySymType, statements yySymType) yySymType {
	statementsNode := statements.node.(*syntax.StatementsNode)
	node := syntax.WhileNode{
		Expression: expr.node,
		Statements: statementsNode,
	}
	return yySymType{node: node}
}

func createFunctionNode(statements yySymType) yySymType {
	statementsNode := statements.node.(*syntax.StatementsNode)
	node := syntax.FunctionNode{
		Statements: statementsNode,
	}
	return yySymType{node: node}
}

func createReturnNode(expr yySymType) yySymType {
	node := syntax.ReturnNode{
		Expression: expr.node,
	}
	return yySymType{node: node}
}

func createCallNode(expr yySymType) yySymType {
	node := syntax.CallNode{
		Expression: expr.node,
	}
	return yySymType{node: node}
}

func appendStatement(statements *yySymType, statement yySymType) yySymType {
	var node *syntax.StatementsNode
	if statements != nil {
		if statementsNode, ok := statements.node.(*syntax.StatementsNode); ok {
			node = statementsNode
		} else {
			panic(fmt.Sprintf("not a statements node: %+v", statements.node))
		}
	} else {
		node = &syntax.StatementsNode{}
	}
	node.Append(statement.node)
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
