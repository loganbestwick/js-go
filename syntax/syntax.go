package syntax

import (
	"strconv"

	"github.com/loganbestwick/js-go/types"
	"fmt"
	"code.justin.tv/web/audrey/_vendor/github.com/davecgh/go-spew/spew"
)


const (
	ADD_OP = "+"
	SUBTRACT_OP = "-"
)

type Node interface {
	Eval() (types.Value, error)
}

type BinaryOpNode struct {
	Left  Node
	Right Node
	Operator string
}

func (n BinaryOpNode) Eval() (types.Value, error) {
	lv, err := n.Left.Eval()
	if err != nil {
		return nil, err
	}
	rv, err := n.Right.Eval()
	if err != nil {
		return nil, err
	}
	spew.Dump(lv)
	switch n.Operator {
	case ADD_OP:
		return lv.Add(rv)
	case SUBTRACT_OP:
		return lv.Subtract(rv)
	default:
		return nil, fmt.Errorf("operator %s not recognized", n.Operator)
	}
}

type NumberNode struct {
	Value string
}

func (t NumberNode) Eval() (types.Value, error) {
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return nil, err
	}
	return types.IntegerValue{Value: i}, nil
}

type StringNode struct {
	Value string
}

func (t StringNode) Eval() (types.Value, error) {
	return types.StringValue{Value: t.Value[1 : len(t.Value)-1]}, nil
}
