package syntax

import (
	"fmt"
	"strconv"

	"github.com/loganbestwick/js-go/types"
)

const (
	ADD_OP        = "+"
	SUBTRACT_OP   = "-"
	ASSIGNMENT_OP = "="
)

type Node interface {
	Eval(*types.Context) (types.Value, error)
}

type StatementsNode struct {
	Statements []Node
}

func (n *StatementsNode) Append(statement Node) {
	n.Statements = append(n.Statements, statement)
}

func (n StatementsNode) Eval(ctx *types.Context) (types.Value, error) {
	var ret types.Value
	for _, statement := range n.Statements {
		var err error
		ret, err = statement.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ret.ToActualValue(ctx)
}

type ConditionalNode struct {
	Expression Node
	Statements *StatementsNode
}

func (n ConditionalNode) Eval(ctx *types.Context) (types.Value, error) {
	expr, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	exprBool, err := expr.ToBooleanValue(ctx)
	if err != nil {
		return nil, err
	}
	if exprBool.Value == true {
		return n.Statements.Eval(ctx)
	}
	return nil, nil
}

type BinaryOpNode struct {
	Left     Node
	Right    Node
	Operator string
}

func (n BinaryOpNode) Eval(ctx *types.Context) (types.Value, error) {
	lv, err := n.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	rv, err := n.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	switch n.Operator {
	case ADD_OP:
		return lv.Add(ctx, rv)
	case SUBTRACT_OP:
		return lv.Subtract(ctx, rv)
	case ASSIGNMENT_OP:
		return lv.Assign(ctx, rv)
	default:
		return nil, fmt.Errorf("operator %s not recognized", n.Operator)
	}
}

type IdentifierNode struct {
	Value string
}

func (i IdentifierNode) Eval(ctx *types.Context) (types.Value, error) {
	return types.IdentifierValue{Value: i.Value}, nil
}

type BooleanNode struct {
	Value string
}

func (t BooleanNode) Eval(ctx *types.Context) (types.Value, error) {
	if t.Value == "true" {
		return types.BooleanValue{Value: true}, nil
	}
	return types.BooleanValue{Value: false}, nil
}

type NumberNode struct {
	Value string
}

func (t NumberNode) Eval(ctx *types.Context) (types.Value, error) {
	if t.Value == "NaN" {
		return types.NaN, nil
	}
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return nil, err
	}
	return types.NumberValue{Value: i}, nil
}

type StringNode struct {
	Value string
}

func (t StringNode) Eval(ctx *types.Context) (types.Value, error) {
	return types.StringValue{Value: t.Value[1 : len(t.Value)-1]}, nil
}
