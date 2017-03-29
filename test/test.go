package test

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/loganbestwick/js-go/parser"
	"github.com/loganbestwick/js-go/types"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	DEBUG = false
)

func eval(ctx *types.Context, code string) (types.Value, error) {
	if !strings.HasSuffix(code, ";") {
		code += ";"
	}
	node := parser.Parse(strings.NewReader(code))
	if DEBUG {
		fmt.Println("")
		fmt.Println("-- CODE --")
		fmt.Println(code)
		fmt.Println("-- AST --")
		spew.Dump(node)
	}
	return node.Eval(ctx)
}

func intVal(i int64) types.NumberValue {
	return types.NumberValue{Value: i}
}

func nanVal() types.NumberValue {
	return types.NaN
}

func strVal(s string) types.StringValue {
	return types.StringValue{Value: s}
}

func identVal(s string) types.IdentifierValue {
	return types.IdentifierValue{Value: s}
}

func assertEval(code string, value types.Value) {
	ctx := &types.Context{}
	testName, _ := value.ToString(ctx)
	Convey(code+" = "+testName, func() {
		result, err := eval(ctx, code)
		if DEBUG {
			fmt.Println("-- VAL --")
			spew.Dump(result)
			fmt.Println("-- ERR --")
			spew.Dump(err)
			fmt.Println("-- END --")
		}
		So(err, ShouldBeNil)
		So(result, ShouldResemble, value)
	})
}
