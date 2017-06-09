package types

import (
	"errors"
	"strconv"
)

var _ Value = BooleanValue{}

type BooleanValue struct {
	Value bool
}

func (a BooleanValue) ToString(ctx *Context) (string, error) {
	str := strconv.FormatBool(a.Value)
	return str, nil
}

func (a BooleanValue) ToActualValue(ctx *Context) (Value, error) {
	return a, nil
}

func (a BooleanValue) ToStringValue(ctx *Context) (StringValue, error) {
	str := strconv.FormatBool(a.Value)
	return StringValue{Value: str}, nil
}

func (a BooleanValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	var num int64
	if a.Value {
		num = 1
	}
	return NumberValue{Value: num}, nil
}

func (a BooleanValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	return a, nil
}

func (a BooleanValue) Add(ctx *Context, b Value) (Value, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	if _, ok := ab.(StringValue); ok {
		sa, err := a.ToStringValue(ctx)
		if err != nil {
			return nil, err
		}
		return sa.Add(ctx, b)
	}
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return na.Add(ctx, b)
}

func (a BooleanValue) Subtract(ctx *Context, b Value) (Value, error) {
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return na.Subtract(ctx, b)
}

func (a BooleanValue) Assign(ctx *Context, b Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

func (a BooleanValue) Equal(ctx *Context, b Value) (Value, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	if bb, ok := ab.(BooleanValue); ok {
		if a.Value == bb.Value {
			return BooleanValue{Value: true}, nil
		}
	}
	return BooleanValue{Value: false}, nil
}

func (a BooleanValue) NotEqual(ctx *Context, b Value) (Value, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	if bb, ok := ab.(BooleanValue); ok {
		if a.Value == bb.Value {
			return BooleanValue{Value: false}, nil
		}
	}
	return BooleanValue{Value: true}, nil
}

func (a BooleanValue) Compare(ctx *Context, b Value, strict bool) (int, bool, error) {
	if strict {
		ab, err := b.ToActualValue(ctx)
		if err != nil {
			return 0, false, err
		}
		if _, ok := ab.(BooleanValue); !ok {
			return 0, true, nil
		}
	}
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return 0, false, err
	}
	return na.Compare(ctx, b, false)
}
