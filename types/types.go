package types

import (
	"strconv"
)

type Value interface {
	ToString() string

	ToStringValue() StringValue
	ToIntegerValue() IntegerValue

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
var _ Value = IntegerValue{}

// Constants definition
var NaN IntegerValue = IntegerValue{NaN: true}

type StringValue struct {
	Value string
}

func (t StringValue) ToString() string {
	return "'" + t.Value + "'"
}

func (t StringValue) ToStringValue() StringValue {
	return t
}

func (t StringValue) ToIntegerValue() IntegerValue {
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return NaN
	}
	return IntegerValue{Value: i}
}

func (t StringValue) Add(v Value) (Value, error) {
	sv := v.ToStringValue()
	return StringValue{Value: t.Value + sv.Value}, nil
}

func (t StringValue) Subtract(v Value) (Value, error) {
	return t.ToIntegerValue().Subtract(v)
}

type IntegerValue struct {
	NaN   bool
	Value int64
}

func (t IntegerValue) ToString() string {
	if t.NaN {
		return "NaN"
	}
	return strconv.FormatInt(t.Value, 10)
}

func (t IntegerValue) ToStringValue() StringValue {
	if t.NaN {
		return StringValue{Value: "NaN"}
	}
	s := strconv.FormatInt(t.Value, 10)
	return StringValue{Value: s}
}

func (t IntegerValue) ToIntegerValue() IntegerValue {
	return t
}

func (t IntegerValue) Add(v Value) (Value, error) {
	if iv, ok := v.(IntegerValue); ok {
		if t.NaN || iv.NaN {
			return NaN, nil
		}
		return IntegerValue{Value: t.Value + iv.Value}, nil
	}
	return t.ToStringValue().Add(v)
}

func (t IntegerValue) Subtract(v Value) (Value, error) {
	iv := v.ToIntegerValue()
	if t.NaN || iv.NaN {
		return NaN, nil
	}
	return IntegerValue{Value: t.Value - iv.Value}, nil
}
