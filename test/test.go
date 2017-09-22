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

func boolVal(b bool) types.BooleanValue {
	return types.BooleanValue{Value: b}
}

func testEval(ctx *types.Context, code string) (types.Value, error) {
	result, err := eval(ctx, code)
	if DEBUG {
		fmt.Println("-- VAL --")
		spew.Dump(result)
		fmt.Println("-- ERR --")
		spew.Dump(err)
		fmt.Println("-- END --")
	}
	return result, err
}

func assertEval(code string, value types.Value) {
	ctx := &types.Context{}
	val, _ := value.ToString(ctx)
	Convey(collapseCode(code)+" = "+val, func() {
		result, err := testEval(ctx, code)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, value)
	})
}

func collapseCode(code string) string {
	code = strings.Replace(code, "\t", "", -1)
	code = strings.Replace(code, "\n", " ", -1)
	code = strings.TrimSpace(code)
	return code
}

func assertError(code string, errString string) {
	ctx := &types.Context{}
	Convey(code+" = "+errString, func() {
		result, err := eval(ctx, code)
		if DEBUG {
			fmt.Println("-- VAL --")
			spew.Dump(result)
			fmt.Println("-- ERR --")
			spew.Dump(err)
			fmt.Println("-- END --")
		}
		So(err.Error(), ShouldContainSubstring, errString)
		So(result, ShouldBeNil)
	})
}
