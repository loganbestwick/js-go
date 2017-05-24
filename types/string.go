package types

import (
	"errors"
	"strconv"
)

var _ Value = StringValue{}

type StringValue struct {
	Value string
}

func (a StringValue) ToString(ctx *Context) (string, error) {
	return "'" + a.Value + "'", nil
}

func (a StringValue) ToActualValue(ctx *Context) (Value, error) {
	return a, nil
}

func (a StringValue) ToStringValue(ctx *Context) (StringValue, error) {
	return a, nil
}

func (a StringValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	i, err := strconv.ParseInt(a.Value, 10, 64)
	if err != nil {
		return NaN, nil
	}
	return NumberValue{Value: i}, nil
}

func (a StringValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	if len(a.Value) > 0 {
		return BooleanValue{Value: true}, nil
	}
	return BooleanValue{Value: false}, nil
}

func (a StringValue) Add(ctx *Context, b Value) (Value, error) {
	sb, err := b.ToStringValue(ctx)
	if err != nil {
		return nil, err
	}
	return StringValue{Value: a.Value + sb.Value}, nil
}

func (a StringValue) Subtract(ctx *Context, b Value) (Value, error) {
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return na.Subtract(ctx, b)
}

func (a StringValue) Assign(ctx *Context, value Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}
