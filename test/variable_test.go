package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVariables(t *testing.T) {
	Convey("Assignment", t, func() {
		assertEval("a = 1; a + a", intVal(2))

		// Transitive assignment
		assertEval("a = 1; b = a; b", intVal(1))
		assertEval("a = 1; b = a; a = 2; b", intVal(1))
		assertEval("a = 1; b = a; a = 2; a", intVal(2))

		// Nested assignment
		assertEval("a = b = 2; a", intVal(2))
		assertEval("a = b = 2; b", intVal(2))

		// Not defined
		assertError("a", "a is not defined")

		// Invalid assignment
		assertError("1 = 1", "Invalid left-hand side in assignment")
		assertError("'a' = 'a'", "Invalid left-hand side in assignment")
	})
}
