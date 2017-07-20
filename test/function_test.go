package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFunctions(t *testing.T) {
	Convey("functions", t, func() {
		assertEval("x = function() { return 1; };", intVal(1))
		assertEval("function x() { return 1; } x();", intVal(1))
	})
}
