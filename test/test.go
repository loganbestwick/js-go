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
	return node.Eval()
}

func intVal(i int64) types.IntegerValue {
	return types.IntegerValue{Value: i}
}

func strVal(s string) types.StringValue {
	return types.StringValue{Value: s}
}

func assertEval(code string, value types.Value) {
	Convey(code+" = "+value.ToString(), func() {
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
