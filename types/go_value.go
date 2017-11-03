package types

import (
	"errors"
	"fmt"
)

var _ Value = GoValue{}

type GoValue struct {
	Value interface{}
}

func (a GoValue) ToString(ctx *Context) (string, error) {
	return fmt.Sprintf("GoValue(%v)", a.Value), nil
}

func (a GoValue) ToActualValue(ctx *Context) (Value, error) {
	return nil, nil
}

func (a GoValue) ToStringValue(ctx *Context) (StringValue, error) {
	av, err := a.ToActualValue(ctx)
	if err != nil {
		return StringValue{}, err
	}
	return av.ToStringValue(ctx)
}

func (a GoValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	av, err := a.ToActualValue(ctx)
	if err != nil {
		return NumberValue{}, err
	}
	return av.ToNumberValue(ctx)
}

func (a GoValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	av, err := a.ToActualValue(ctx)
	if err != nil {
		return BooleanValue{}, err
	}
	return av.ToBooleanValue(ctx)
}

func (a GoValue) Add(ctx *Context, b Value) (Value, error) {
	av, err := a.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	return av.Add(ctx, b)
}

func (a GoValue) Subtract(ctx *Context, b Value) (Value, error) {
	av, err := a.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	return av.Subtract(ctx, b)
}

func (a GoValue) Assign(ctx *Context, b Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

func (a GoValue) Compare(ctx *Context, b Value, strict bool) (int, bool, error) {
	av, err := a.ToActualValue(ctx)
	if err != nil {
		return 0, false, err
	}
	return av.Compare(ctx, b, strict)
}

func (a GoValue) Call(ctx *Context, arguments []Value) (Value, error) {
	return nil, errors.New("not implemented")
}
