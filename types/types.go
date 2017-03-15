package types

import (
	"strconv"
)

type Context struct {
	Variables map[string]Value
}

func (c Context) Set(s string, v Value) Value {
	c.Variables[s] = v
	return v
}

func (c Context) Get(s string) (Value, error) {
	v := c.Variables[s]
	if v == nil {
		// Return ReferenceError
	}
	return v, nil
}

type Value interface {
	ToString() string

	ToStringValue() StringValue
	ToNumberValue() NumberValue

	// Rules for addition:
	// If either operand is a string, do string concatenation
	// If both operands are numbers, do addition
	Add(Value) (Value, error)

	// Rules for subtraction:
	// Convert both operands to numbers, do subtraction
	Subtract(Value) (Value, error)
}

// Interface assertions
var _ Value = StringValue{}
var _ Value = NumberValue{}

// Constants definition
var NaN NumberValue = NumberValue{NaN: true}

type StringValue struct {
	Value string
}

func (t StringValue) ToString() string {
	return "'" + t.Value + "'"
}

func (t StringValue) ToStringValue() StringValue {
	return t
}

func (t StringValue) ToNumberValue() NumberValue {
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return NaN
	}
	return NumberValue{Value: i}
}

func (t StringValue) Add(v Value) (Value, error) {
	sv := v.ToStringValue()
	return StringValue{Value: t.Value + sv.Value}, nil
}

func (t StringValue) Subtract(v Value) (Value, error) {
	return t.ToNumberValue().Subtract(v)
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

func (t NumberValue) ToStringValue() StringValue {
	if t.NaN {
		return StringValue{Value: "NaN"}
	}
	s := strconv.FormatInt(t.Value, 10)
	return StringValue{Value: s}
}

func (t NumberValue) ToNumberValue() NumberValue {
	return t
}

func (t NumberValue) Add(v Value) (Value, error) {
	if iv, ok := v.(NumberValue); ok {
		if t.NaN || iv.NaN {
			return NaN, nil
		}
		return NumberValue{Value: t.Value + iv.Value}, nil
	}
	return t.ToStringValue().Add(v)
}

func (t NumberValue) Subtract(v Value) (Value, error) {
	iv := v.ToNumberValue()
	if t.NaN || iv.NaN {
		return NaN, nil
	}
	return NumberValue{Value: t.Value - iv.Value}, nil
}
