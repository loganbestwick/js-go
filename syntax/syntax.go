package syntax

import (
	"fmt"
	"strconv"

	"github.com/loganbestwick/js-go/types"
)

const (
	ADD_OP      = "+"
	SUBTRACT_OP = "-"
	ASSIGN_OP   = "="
)

type Node interface {
	Eval(*types.Context) (types.Value, error)
}

type StatementsNode struct {
	Statements []Node
}

func (n StatementsNode) Eval(c *types.Context) (types.Value, error) {
	var ret types.Value
	var err error
	for _, statement := range n.Statements {
		ret, err = statement.Eval(c)
		if err != nil {
			return nil, err
		}
	}
	return ret.ToActualValue(c)
}

type BinaryOpNode struct {
	Left     Node
	Right    Node
	Operator string
}

func (n BinaryOpNode) Eval(c *types.Context) (types.Value, error) {
	lv, err := n.Left.Eval(c)
	if err != nil {
		return nil, err
	}
	rv, err := n.Right.Eval(c)
	if err != nil {
		return nil, err
	}
	switch n.Operator {
	case ADD_OP:
		return lv.Add(c, rv)
	case SUBTRACT_OP:
		return lv.Subtract(c, rv)
	case ASSIGN_OP:
		return lv.Assign(c, rv)
	default:
		return nil, fmt.Errorf("operator %s not recognized", n.Operator)
	}
}

type NumberNode struct {
	Value string
}

func (t NumberNode) Eval(c *types.Context) (types.Value, error) {
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

func (t StringNode) Eval(c *types.Context) (types.Value, error) {
	return types.StringValue{Value: t.Value[1 : len(t.Value)-1]}, nil
}

type VariableNode struct {
	Value string
}

func (t VariableNode) Eval(c *types.Context) (types.Value, error) {
	return types.VariableValue{Variable: t.Value}, nil
}
