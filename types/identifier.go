package types

var _ Value = IdentifierValue{}

type IdentifierValue struct {
	Value string
}

func (a IdentifierValue) ToString(ctx *Context) (string, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return "", err
	}
	return aa.ToString(ctx)
}

func (a IdentifierValue) ToActualValue(ctx *Context) (Value, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return nil, err
	}
	return aa, err
}

func (a IdentifierValue) ToStringValue(ctx *Context) (StringValue, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return StringValue{}, err
	}
	return aa.ToStringValue(ctx)
}

func (a IdentifierValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return NumberValue{}, err
	}
	return aa.ToNumberValue(ctx)
}

func (a IdentifierValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return BooleanValue{}, err
	}
	return aa.ToBooleanValue(ctx)
}

func (a IdentifierValue) Add(ctx *Context, b Value) (Value, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return nil, err
	}
	return aa.Add(ctx, b)
}

func (a IdentifierValue) Subtract(ctx *Context, b Value) (Value, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return nil, err
	}
	return aa.Subtract(ctx, b)
}

func (a IdentifierValue) Assign(ctx *Context, b Value) (Value, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	ctx.Set(a.Value, ab)
	return ab, nil
}

func (a IdentifierValue) Compare(ctx *Context, b Value, strict bool) (int, bool, error) {
	aa, err := a.ToActualValue(ctx)
	if err != nil {
		return 0, false, err
	}
	return aa.Compare(ctx, b, strict)
}

func (a IdentifierValue) Call(ctx *Context) (Value, error) {
	aa, err := ctx.Get(a.Value)
	if err != nil {
		return nil, err
	}
	return aa.Call(ctx)
}
