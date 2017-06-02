package types

import (
	"errors"
)

type Context struct {
	Variables map[string]Value
}

func (c *Context) Set(s string, v Value) Value {
	if c.Variables == nil {
		c.Variables = make(map[string]Value)
	}
	av, err := v.ToActualValue(c)
	if err != nil {
		return nil
	}
	c.Variables[s] = av
	return v
}

func (c Context) Get(s string) (Value, error) {
	v, ok := c.Variables[s]
	if !ok {
		return nil, errors.New("ReferenceError: " + s + " is not defined")
	}
	return v, nil
}

type Value interface {
	ToString(*Context) (string, error)

	ToActualValue(*Context) (Value, error)
	ToStringValue(*Context) (StringValue, error)
	ToNumberValue(*Context) (NumberValue, error)
	ToBooleanValue(*Context) (BooleanValue, error)

	// Rules for addition:
	// If either operand is a string, do string concatenation
	// If both operands are numbers, do addition
	Add(*Context, Value) (Value, error)

	// Rules for subtraction:
	// Convert both operands to numbers, do subtraction
	Subtract(*Context, Value) (Value, error)

	// Rules for assignment:
	Assign(*Context, Value) (Value, error)

	// Rules for equality:
	// Strict equality rules only (===)
	// Left value must equal type and value of right value
	Equal(*Context, Value) (Value, error)

	// Rules for inequality:
	// Strict equality rules only (!==)
	// Left value's type or value is not equal to right's
	NotEqual(*Context, Value) (Value, error)
}
