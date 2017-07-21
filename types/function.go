package types

import (
	"errors"
)

var _ Value = FunctionValue{}

type Evalable interface {
	Eval(*Context) (Value, error)
}

type ErrReturn struct {
	ReturnValue Value
}

func (e *ErrReturn) Error() string {
	return "[ReturnValue]"
}

type FunctionValue struct {
	Statements Evalable
}

func (a FunctionValue) ToString(ctx *Context) (string, error) {
	return "function() {...}", nil
}

func (a FunctionValue) ToActualValue(ctx *Context) (Value, error) {
	return a, nil
}

func (a FunctionValue) ToStringValue(ctx *Context) (StringValue, error) {
	return StringValue{Value: "function() {...}"}, nil
}

func (a FunctionValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	return NaN, nil
}

func (a FunctionValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	return BooleanValue{Value: true}, nil
}

func (a FunctionValue) Add(ctx *Context, b Value) (Value, error) {
	sv, err := a.ToStringValue(ctx)
	if err != nil {
		return nil, err
	}
	return sv.Add(ctx, b)
}

func (a FunctionValue) Subtract(ctx *Context, b Value) (Value, error) {
	sa, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return sa.Subtract(ctx, b)
}

func (a FunctionValue) Assign(ctx *Context, b Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

func (a FunctionValue) Compare(ctx *Context, b Value, strict bool) (int, bool, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return 0, false, err
	}
	if a == ab {
		return 0, false, nil
	} else {
		return 0, true, nil
	}
}

func (a FunctionValue) Call(ctx *Context) (Value, error) {
	functionCtx := &Context{}
	return a.Statements.Eval(functionCtx)
}
