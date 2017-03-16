package types

import (
	"fmt"
)

type Context struct {
	Variables map[string]Value
}

func (c *Context) Get(variable string) (Value, error) {
	if value, ok := c.Variables[variable]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("%s is not defined", variable)
}

func (c *Context) Set(variable string, value Value) {
	if c.Variables == nil {
		c.Variables = make(map[string]Value)
	}
	c.Variables[variable] = value
}
