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
		assertEval("x = function() { return 1; }; a = x(); a;", intVal(1))
		assertEval("x = function(age) { return age; }; x(15);", intVal(15))
		assertEval("x = function(age) { return age; }; x(10 + 5);", intVal(15))

		maxFunc := `
		max = function(a, b) {
			if (a > b) {
				return a;
			}
			return b;
		};
		`
		assertEval(maxFunc+"max(5, 10);", intVal(10))
		assertEval(maxFunc+"max(10, 5);", intVal(10))
		assertEval(maxFunc+"max(max(5, 10), 3);", intVal(10))
		assertEval(maxFunc+"x = 10; max(x, x + 1);", intVal(11))
	})
}
