package types

import (
	"errors"
	"strconv"
)

var _ Value = NumberValue{}

// Constants definition
var NaN NumberValue = NumberValue{NaN: true}

type NumberValue struct {
	NaN   bool
	Value int64
}

func (a NumberValue) ToString(ctx *Context) (string, error) {
	if a.NaN {
		return "NaN", nil
	}
	return strconv.FormatInt(a.Value, 10), nil
}

func (a NumberValue) ToActualValue(ctx *Context) (Value, error) {
	return a, nil
}

func (a NumberValue) ToStringValue(ctx *Context) (StringValue, error) {
	if a.NaN {
		return StringValue{Value: "NaN"}, nil
	}
	s := strconv.FormatInt(a.Value, 10)
	return StringValue{Value: s}, nil
}

func (a NumberValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	return a, nil
}

func (a NumberValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	if a.Value != 0 && a.NaN != true {
		return BooleanValue{Value: true}, nil
	}
	return BooleanValue{Value: false}, nil
}

func (a NumberValue) Add(ctx *Context, b Value) (Value, error) {
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
	nab, err := ab.ToNumberValue(ctx)
	if a.NaN || nab.NaN {
		return NaN, nil
	}
	return NumberValue{Value: a.Value + nab.Value}, nil
}

func (a NumberValue) Subtract(ctx *Context, b Value) (Value, error) {
	nb, err := b.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	if a.NaN || nb.NaN {
		return NaN, nil
	}
	return NumberValue{Value: a.Value - nb.Value}, nil
}

func (a NumberValue) Assign(ctx *Context, value Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

func (a NumberValue) Equal(ctx *Context, b Value) (Value, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	if nb, ok := ab.(NumberValue); ok {
		if !a.NaN && !nb.NaN && a.Value == nb.Value {
			return BooleanValue{Value:true}, nil
		}
	}
	return BooleanValue{Value:false}, nil
}

func (a NumberValue) NotEqual(ctx *Context, b Value) (Value, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	if nb, ok := ab.(NumberValue); ok {
		if !a.NaN && !nb.NaN && a.Value == nb.Value {
			return BooleanValue{Value:false}, nil
		}
	}
	return BooleanValue{Value:true}, nil
}

func (a NumberValue) Compare(ctx *Context, b Value) (*int, error) {
	nb, err := b.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	cmp := int(a.Value - nb.Value)
	return &cmp, nil
}
