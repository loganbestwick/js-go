package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOperators(t *testing.T) {
	Convey("Addition", t, func() {
		assertEval("1 + 1", intVal(2))
		assertEval("1 + 2", intVal(3))
		assertEval("1 + 2 + 3", intVal(6))

		assertEval("'2' + 1", strVal("21"))
		assertEval("1 + '2'", strVal("12"))
		assertEval("1 + '2' + 3", strVal("123"))

		assertEval("NaN + 1", nanVal())
		assertEval("1 + NaN", nanVal())
		assertEval("NaN + '1'", strVal("NaN1"))
		assertEval("'1' + NaN", strVal("1NaN"))
	})

	Convey("Subtraction", t, func() {
		assertEval("2 - 1", intVal(1))
		assertEval("5 - 2", intVal(3))
		assertEval("5 - 2 - 1", intVal(2))
		assertEval("1 - 2", intVal(-1))
		assertEval("1 - 2 - 3", intVal(-4))

		assertEval("'1' - 1", intVal(0))
		assertEval("1 - '1'", intVal(0))
		assertEval("'1' - '1'", intVal(0))
		assertEval("'5' - '2' - '1'", intVal(2))

		assertEval("'a' - 1", nanVal())
		assertEval("1 - 'a'", nanVal())
		assertEval("'a' - '1'", nanVal())
		assertEval("'1' - 'a'", nanVal())

		assertEval("NaN - 1", nanVal())
		assertEval("1 - NaN", nanVal())
		assertEval("NaN - '1'", nanVal())
		assertEval("'1' - NaN", nanVal())
	})

	Convey("Assignment", t, func() {
		assertEval("x = 1", intVal(1))
		assertEval("x = y = 1", intVal(1))
		assertEval("x = y = z = 1", intVal(1))
	})
}
