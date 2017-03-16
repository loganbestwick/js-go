package types

import (
	"errors"
	"strconv"
)

var (
	// Errors
	ErrInvalidAssignment = errors.New("Invalid left-hand side in assignment")

	// Constants
	NaN = NumberValue{NaN: true}
)

type Value interface {
	ToString() string

	ToActualValue(*Context) (Value, error)
	ToStringValue(*Context) (StringValue, error)
	ToNumberValue(*Context) (NumberValue, error)

	// Rules for addition:
	// If either operand is a string, do string concatenation
	// If both operands are numbers, do addition
	Add(*Context, Value) (Value, error)

	// Rules for subtraction:
	// Convert both operands to numbers, do subtraction
	Subtract(*Context, Value) (Value, error)

	// Rules for assign:
	// Left side must be a variable, returns assigned value
	Assign(*Context, Value) (Value, error)
}

// Interface assertions
var _ Value = StringValue{}
var _ Value = NumberValue{}
var _ Value = VariableValue{}

type StringValue struct {
	Value string
}

func (t StringValue) ToString() string {
	return "'" + t.Value + "'"
}

func (t StringValue) ToActualValue(c *Context) (Value, error) {
	return t, nil
}

func (t StringValue) ToStringValue(c *Context) (StringValue, error) {
	return t, nil
}

func (t StringValue) ToNumberValue(c *Context) (NumberValue, error) {
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return NaN, nil
	}
	return NumberValue{Value: i}, nil
}

func (t StringValue) Add(c *Context, v Value) (Value, error) {
	sv, err := v.ToStringValue(c)
	if err != nil {
		return nil, err
	}
	return StringValue{Value: t.Value + sv.Value}, nil
}

func (t StringValue) Subtract(c *Context, v Value) (Value, error) {
	nv, err := t.ToNumberValue(c)
	if err != nil {
		return nil, err
	}
	return nv.Subtract(c, v)
}

func (t StringValue) Assign(c *Context, v Value) (Value, error) {
	return nil, ErrInvalidAssignment
}

type NumberValue struct {
	NaN   bool
	Value int64
}

func (t NumberValue) ToString() string {
	if t.NaN {
		return "NaN"
	}
	return strconv.FormatInt(t.Value, 10)
}

func (t NumberValue) ToActualValue(c *Context) (Value, error) {
	return t, nil
}

func (t NumberValue) ToStringValue(c *Context) (StringValue, error) {
	if t.NaN {
		return StringValue{Value: "NaN"}, nil
	}
	s := strconv.FormatInt(t.Value, 10)
	return StringValue{Value: s}, nil
}

func (t NumberValue) ToNumberValue(c *Context) (NumberValue, error) {
	return t, nil
}

func (t NumberValue) Add(c *Context, v Value) (Value, error) {
	av, err := v.ToActualValue(c)
	if err != nil {
		return nil, err
	}
	if nv, ok := av.(NumberValue); ok {
		if t.NaN || nv.NaN {
			return NaN, nil
		}
		return NumberValue{Value: t.Value + nv.Value}, nil
	}
	sv, err := t.ToStringValue(c)
	if err != nil {
		return nil, err
	}
	return sv.Add(c, v)
}

func (t NumberValue) Subtract(c *Context, v Value) (Value, error) {
	nv, err := v.ToNumberValue(c)
	if err != nil {
		return nil, err
	}
	if t.NaN || nv.NaN {
		return NaN, nil
	}
	return NumberValue{Value: t.Value - nv.Value}, nil
}

func (t NumberValue) Assign(c *Context, v Value) (Value, error) {
	return nil, ErrInvalidAssignment
}

type VariableValue struct {
	Variable string
}

func (t VariableValue) ToString() string {
	return t.Variable
}

func (t VariableValue) ToActualValue(c *Context) (Value, error) {
	val, err := c.Get(t.Variable)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (t VariableValue) ToStringValue(c *Context) (StringValue, error) {
	val, err := c.Get(t.Variable)
	if err != nil {
		return StringValue{}, err
	}
	return val.ToStringValue(c)
}

func (t VariableValue) ToNumberValue(c *Context) (NumberValue, error) {
	val, err := c.Get(t.Variable)
	if err != nil {
		return NumberValue{}, err
	}
	return val.ToNumberValue(c)
}

func (t VariableValue) Add(c *Context, v Value) (Value, error) {
	val, err := c.Get(t.Variable)
	if err != nil {
		return nil, err
	}
	return val.Add(c, v)
}

func (t VariableValue) Subtract(c *Context, v Value) (Value, error) {
	val, err := c.Get(t.Variable)
	if err != nil {
		return nil, err
	}
	return val.Subtract(c, v)
}

func (t VariableValue) Assign(c *Context, v Value) (Value, error) {
	val, err := v.ToActualValue(c)
	if err != nil {
		return nil, err
	}
	c.Set(t.Variable, val)
	return val, nil
}
