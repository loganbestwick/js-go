package types

import (
	"errors"
	"strconv"
)

type Context struct {
	Variables map[string]Value
}

func (c *Context) Set(s string, v Value) Value {
	if c.Variables == nil {
		c.Variables = make(map[string]Value)
	}
	av, err := v.ToActualValue(c)
	if err != nil {
		return nil
	}
	c.Variables[s] = av
	return v
}

func (c Context) Get(s string) (Value, error) {
	v, ok := c.Variables[s]
	if !ok {
		return nil, errors.New("ReferenceError: " + s + " is not defined")
	}
	return v, nil
}

type Value interface {
	ToString(*Context) (string, error)

	ToActualValue(*Context) (Value, error)
	ToStringValue(*Context) (StringValue, error)
	ToNumberValue(*Context) (NumberValue, error)
	ToBoolValue(*Context) (BooleanValue, error)

	// Rules for addition:
	// If either operand is a string, do string concatenation
	// If both operands are numbers, do addition
	Add(*Context, Value) (Value, error)

	// Rules for subtraction:
	// Convert both operands to numbers, do subtraction
	Subtract(*Context, Value) (Value, error)

	// Rules for assignment:
	Assign(*Context, Value) (Value, error)
}

// Interface assertions
var _ Value = StringValue{}
var _ Value = NumberValue{}
var _ Value = BooleanValue{}

// Constants definition
var NaN NumberValue = NumberValue{NaN: true}

type IdentifierValue struct {
	Value string
}

func (i IdentifierValue) ToString(ctx *Context) (string, error) {
	av, err := ctx.Get(i.Value)
	if err != nil {
		return "", err
	}
	return av.ToString(ctx)
}

func (i IdentifierValue) ToActualValue(ctx *Context) (Value, error) {
	av, err := ctx.Get(i.Value)
	if err != nil {
		return nil, err
	}
	return av, err
}

func (i IdentifierValue) ToStringValue(ctx *Context) (StringValue, error) {
	av, err := ctx.Get(i.Value)
	if err != nil {
		return StringValue{}, err
	}
	return av.ToStringValue(ctx)
}

func (i IdentifierValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	av, err := ctx.Get(i.Value)
	if err != nil {
		return NumberValue{}, err
	}
	return av.ToNumberValue(ctx)
}

func (i IdentifierValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	av, err := ctx.Get(i.Value)
	if err != nil {
		return BooleanValue{}, err
	}
	return av.ToBoolValue(ctx)
}

func (i IdentifierValue) Add(ctx *Context, v Value) (Value, error) {
	av, err := ctx.Get(i.Value)
	if err != nil {
		return nil, err
	}
	return av.Add(ctx, v)
}

func (i IdentifierValue) Subtract(ctx *Context, v Value) (Value, error) {
	av, err := ctx.Get(i.Value)
	if err != nil {
		return nil, err
	}
	return av.Subtract(ctx, v)
}

func (i IdentifierValue) Assign(ctx *Context, value Value) (Value, error) {
	av, err := value.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	ctx.Set(i.Value, av)
	return av, nil
}

type StringValue struct {
	Value string
}

func (t StringValue) ToString(ctx *Context) (string, error) {
	return "'" + t.Value + "'", nil
}

func (t StringValue) ToActualValue(ctx *Context) (Value, error) {
	return t, nil
}

func (t StringValue) ToStringValue(ctx *Context) (StringValue, error) {
	return t, nil
}

func (t StringValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return NaN, nil
	}
	return NumberValue{Value: i}, nil
}

func (t StringValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	if len(t.Value) > 0 {
		return BooleanValue{Value: true}, nil
	}
	return BooleanValue{Value: false}, nil
}

func (t StringValue) Add(ctx *Context, v Value) (Value, error) {
	sv, err := v.ToStringValue(ctx)
	if err != nil {
		return nil, err
	}
	return StringValue{Value: t.Value + sv.Value}, nil
}

func (t StringValue) Subtract(ctx *Context, v Value) (Value, error) {
	nv, err := t.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return nv.Subtract(ctx, v)
}

func (t StringValue) Assign(ctx *Context, value Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

type NumberValue struct {
	NaN   bool
	Value int64
}

type ObjectValue struct {
	Keys map[string]Value
}

func (t NumberValue) ToString(ctx *Context) (string, error) {
	if t.NaN {
		return "NaN", nil
	}
	return strconv.FormatInt(t.Value, 10), nil
}

func (t NumberValue) ToActualValue(ctx *Context) (Value, error) {
	return t, nil
}

func (t NumberValue) ToStringValue(ctx *Context) (StringValue, error) {
	if t.NaN {
		return StringValue{Value: "NaN"}, nil
	}
	s := strconv.FormatInt(t.Value, 10)
	return StringValue{Value: s}, nil
}

func (t NumberValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	return t, nil
}

func (t NumberValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	if t.Value != 0 && t.NaN != true {
		return BooleanValue{Value: true}, nil
	}
	return BooleanValue{Value: false}, nil
}

func (t NumberValue) Add(ctx *Context, v Value) (Value, error) {
	av, err := v.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	if iv, ok := av.(NumberValue); ok {
		if t.NaN || iv.NaN {
			return NaN, nil
		}
		return NumberValue{Value: t.Value + iv.Value}, nil
	}
	sv, err := t.ToStringValue(ctx)
	if err != nil {
		return nil, err
	}
	return sv.Add(ctx, v)
}

func (t NumberValue) Subtract(ctx *Context, v Value) (Value, error) {
	iv, err := v.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	if t.NaN || iv.NaN {
		return NaN, nil
	}
	return NumberValue{Value: t.Value - iv.Value}, nil
}

func (t NumberValue) Assign(ctx *Context, value Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

type BooleanValue struct {
	Value bool
}

func (b BooleanValue) ToString(ctx *Context) (string, error) {
	return strconv.FormatBool(b.Value), nil
}

func (b BooleanValue) ToActualValue(ctx *Context) (Value, error) {
	return b, nil
}

func (b BooleanValue) ToStringValue(ctx *Context) (StringValue, error) {
	return strconv.FormatBool(b.Value), nil
}

func (b BooleanValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	return b, nil
}

func (b BooleanValue) Add(ctx *Context, v Value) (Value, error) {

}

func (b BooleanValue) Subtract(ctx *Context, v Value) (Value, error) {

}

func (b BooleanValue) Assign(ctx *Context, v Value) (Value, error) {

}
