package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConditionals(t *testing.T) {
	Convey("if", t, func() {
		assertEval("if (true) { 1; }", intVal(1))
		assertEval("a = 1; if (true) { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; if (false) { a = a + 1; } a;", intVal(1))
	})
}
