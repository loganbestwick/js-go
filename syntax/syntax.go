package syntax

import (
	"fmt"
	"strconv"

	"github.com/loganbestwick/js-go/types"
)

const (
	ADD_OP                   = "+"
	SUBTRACT_OP              = "-"
	ASSIGNMENT_OP            = "="
	GREATER_THAN_OP          = ">"
	LESS_THAN_OP             = "<"
	GREATER_THAN_OR_EQUAL_OP = ">="
	LESS_THAN_OR_EQUAL_OP    = "<="
	EQUALITY_OP              = "=="
	INEQUALITY_OP            = "!="
	EQUALITY_OP_STRICT       = "==="
	INEQUALITY_OP_STRICT     = "!=="
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

type IdentifiersNode struct {
	Identifiers []string
}

func (n *IdentifiersNode) Append(identifier string) {
	n.Identifiers = append(n.Identifiers, identifier)
}

func (n IdentifiersNode) Eval(ctx *types.Context) (types.Value, error) {
	return nil, nil
}

type FunctionNode struct {
	Statements    *StatementsNode
	ArgumentNames []string
}

func (n FunctionNode) Eval(ctx *types.Context) (types.Value, error) {
	return types.FunctionValue{Statements: n.Statements, Variables: n.ArgumentNames}, nil
}

type ReturnNode struct {
	Expression Node
}

func (n ReturnNode) Eval(ctx *types.Context) (types.Value, error) {
	expr, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	val, err := expr.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	return nil, types.ErrReturn{ReturnValue: val}

}

type CallNode struct {
	Expression Node
	Arguments  []Node
}

func (n CallNode) Eval(ctx *types.Context) (types.Value, error) {
	expr, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	values := []types.Value{}
	for _, node := range n.Arguments {
		expr, err := node.Eval(ctx)
		if err != nil {
			return nil, err
		}
		val, err := expr.ToActualValue(ctx)
		if err != nil {
			return nil, err
		}
		values = append(values, val)
	}
	return expr.Call(ctx, values)
}

type ArgumentsNode struct {
	Arguments []Node
}

func (n *ArgumentsNode) Append(argument Node) {
	n.Arguments = append(n.Arguments, argument)
}

func (n ArgumentsNode) Eval(ctx *types.Context) (types.Value, error) {
	return nil, nil
}

type IfNode struct {
	Expression Node
	Statements *StatementsNode
}

func (n IfNode) Eval(ctx *types.Context) (types.Value, error) {
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

type ForNode struct {
	InitExpression Node
	Condition      Node
	LoopExpression Node
	Statements     *StatementsNode
}

func (n ForNode) Eval(ctx *types.Context) (types.Value, error) {
	var lastVal types.Value
	_, err := n.InitExpression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	loop := true
	for loop {
		val, err := n.Condition.Eval(ctx)
		if err != nil {
			return nil, err
		}
		boolVal, err := val.ToBooleanValue(ctx)
		if err != nil {
			return nil, err
		}
		loop = boolVal.Value
		if loop {
			lastVal, err = n.Statements.Eval(ctx)
			if err != nil {
				return nil, err
			}
			_, err := n.LoopExpression.Eval(ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return lastVal, nil
}

type WhileNode struct {
	Expression Node
	Statements *StatementsNode
}

func (n WhileNode) Eval(ctx *types.Context) (types.Value, error) {
	var lastVal types.Value
	loop := true
	for loop {
		val, err := n.Expression.Eval(ctx)
		if err != nil {
			return nil, err
		}
		boolVal, err := val.ToBooleanValue(ctx)
		if err != nil {
			return nil, err
		}
		loop = boolVal.Value
		if loop {
			lastVal, err = n.Statements.Eval(ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return lastVal, nil
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
	case EQUALITY_OP:
		return comp(ctx, lv, rv, false, false, 0)
	case INEQUALITY_OP:
		return comp(ctx, lv, rv, false, true, 0)
	case EQUALITY_OP_STRICT:
		return comp(ctx, lv, rv, true, false, 0)
	case INEQUALITY_OP_STRICT:
		return comp(ctx, lv, rv, true, true, 0)
	case GREATER_THAN_OP:
		return comp(ctx, lv, rv, false, false, 1)
	case GREATER_THAN_OR_EQUAL_OP:
		return comp(ctx, lv, rv, false, false, 0, 1)
	case LESS_THAN_OP:
		return comp(ctx, lv, rv, false, false, -1)
	case LESS_THAN_OR_EQUAL_OP:
		return comp(ctx, lv, rv, false, false, -1, 0)
	default:
		return nil, fmt.Errorf("operator %s not recognized", n.Operator)
	}
}

func comp(ctx *types.Context, lv types.Value, rv types.Value, strict bool, invert bool, expect ...int) (types.Value, error) {
	cmp, forceFalse, err := lv.Compare(ctx, rv, strict)
	if err != nil {
		return nil, err
	}
	if forceFalse {
		if invert {
			return types.BooleanValue{Value: true}, nil
		} else {
			return types.BooleanValue{Value: false}, nil
		}
	}
	successValue := !invert
	for _, val := range expect {
		if val == cmp {
			return types.BooleanValue{Value: successValue}, nil
		}
	}
	return types.BooleanValue{Value: !successValue}, nil
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
