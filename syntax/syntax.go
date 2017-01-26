package syntax

import (
	"strconv"
)

type Node interface {
	Eval() (string, error)
}

type AddNode struct {
	Left  Node
	Right Node
}

func (n AddNode) Eval() (string, error) {
	lval, err := n.Left.Eval()
	if err != nil {
		return "", err
	}
	rval, err := n.Right.Eval()
	if err != nil {
		return "", err
	}
	a, err := strconv.Atoi(lval)
	if err != nil {
		return "", err
	}
	b, err := strconv.Atoi(rval)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(a + b), nil
}

type ValueNode struct {
	Value string
}

func (n ValueNode) Eval() (string, error) {
	return n.Value, nil
}
