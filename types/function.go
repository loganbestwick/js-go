package types

import (
	"errors"
)

var _ Value = FunctionValue{}

type FunctionValue struct {
	Value
}

func (a FunctionValue) ToString(ctx *Context) (string, error) {
	//str := strconv.FormatBool(a.Value)
	return "hi", nil
}

func (a FunctionValue) ToActualValue(ctx *Context) (Value, error) {
	return a, nil
}

func (a FunctionValue) ToStringValue(ctx *Context) (StringValue, error) {
	return StringValue{ Value: "TBD" }, nil
}

func (a FunctionValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	return NaN, nil
}

func (a FunctionValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	return BooleanValue{Value:true}, nil
}

func (a FunctionValue) Add(ctx *Context, b Value) (Value, error) {
	//ab, err := b.ToActualValue(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//if _, ok := ab.(StringValue); ok {
	//	sa, err := a.ToStringValue(ctx)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return sa.Add(ctx, b)
	//}
	//na, err := a.ToNumberValue(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//return na.Add(ctx, b)
	return StringValue{Value: "hi"}, nil
}

func (a FunctionValue) Subtract(ctx *Context, b Value) (Value, error) {
	return NaN, nil
}

func (a FunctionValue) Assign(ctx *Context, b Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

func (a FunctionValue) Compare(ctx *Context, b Value, strict bool) (int, bool, error) {
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

