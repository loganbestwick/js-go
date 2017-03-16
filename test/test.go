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

func eval(code string) (types.Value, error) {
	node := parser.Parse(strings.NewReader(code))
	if DEBUG {
		fmt.Println("")
		fmt.Println("-- CODE --")
		fmt.Println(code)
		fmt.Println("-- AST --")
		spew.Dump(node)
	}
	ctx := &types.Context{}
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

func assertEval(code string, value types.Value) {
	Convey(code+" = "+value.ToString(), func() {
		if !strings.HasSuffix(code, ";") {
			code = code + ";"
		}
		result, err := eval(code)
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

func assertError(code string, e string) {
	Convey(code+" errors '"+e+"'", func() {
		if !strings.HasSuffix(code, ";") {
			code = code + ";"
		}
		result, err := eval(code)
		if DEBUG {
			fmt.Println("-- VAL --")
			spew.Dump(result)
			fmt.Println("-- ERR --")
			spew.Dump(err)
			fmt.Println("-- END --")
		}
		So(err, ShouldNotBeNil)
		So(result, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, e)
	})
}
