package types

import (
	"strconv"
)

type Value interface {
	ToString() string

	ToStringValue() StringValue
	ToIntegerValue() (IntegerValue, error)

	Add(Value) (Value, error)
}

var _ Value = StringValue{}
var _ Value = IntegerValue{}

type StringValue struct {
	Value string
}

func (t StringValue) ToString() string {
	return "'" + t.Value + "'"
}

func (t StringValue) ToStringValue() StringValue {
	return t
}

func (t StringValue) ToIntegerValue() (IntegerValue, error) {
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return IntegerValue{}, nil
	}
	return IntegerValue{Value: i}, nil
}

func (t StringValue) Add(v Value) (Value, error) {
	sv := v.ToStringValue()
	return StringValue{Value: t.Value + sv.Value}, nil
}

type IntegerValue struct {
	Value int64
}

func (t IntegerValue) ToString() string {
	return strconv.FormatInt(t.Value, 10)
}

func (t IntegerValue) ToStringValue() StringValue {
	s := strconv.FormatInt(t.Value, 10)
	return StringValue{Value: s}
}

func (t IntegerValue) ToIntegerValue() (IntegerValue, error) {
	return t, nil
}

func (t IntegerValue) Add(v Value) (Value, error) {
	if iv, ok := v.(IntegerValue); ok {
		return IntegerValue{Value: t.Value + iv.Value}, nil
	}
	return t.ToStringValue().Add(v)
}
