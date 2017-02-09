package syntax

import (
	"strconv"

	"github.com/loganbestwick/js-go/types"
)

type Node interface {
	Eval() (types.Value, error)
}

type AddNode struct {
	Left  Node
	Right Node
}

func (n AddNode) Eval() (types.Value, error) {
	lv, err := n.Left.Eval()
	if err != nil {
		return nil, err
	}
	rv, err := n.Right.Eval()
	if err != nil {
		return nil, err
	}
	return lv.Add(rv)
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
