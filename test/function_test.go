package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFunctions(t *testing.T) {
	Convey("functions", t, func() {
		assertEval("x = function() { return 1; }; x();", intVal(1))
		assertEval("x = function() { return 2; return 1; }; x();", intVal(2))
		assertEval("x = function() { a = 1; return a; }; x();", intVal(1))
		assertEval("function x() { return 1; } x();", intVal(1))
	})
}
