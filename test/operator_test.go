package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOperators(t *testing.T) {
	Convey("Addition", t, func() {
		// Int tests
		assertEval("1 + 1;", intVal(2))
		assertEval("1 + 2;", intVal(3))
		assertEval("1 + 2 + 3;", intVal(6))
		assertEval("1 + '2';", strVal("12"))
		assertEval("1 + NaN;", nanVal())
		assertEval("1 + true;", intVal(2))
		assertEval("1 + false;", intVal(1))

		// NaN tests
		assertEval("NaN + 1;", nanVal())
		assertEval("NaN + '2';", strVal("NaN2"))
		assertEval("NaN + true;", nanVal())

		// String tests
		assertEval("'1' + 1;", strVal("11"))
		assertEval("'1' + 2 + 3;", strVal("123"))
		assertEval("'1' + NaN;", strVal("1NaN"))
		assertEval("'1' + true;", strVal("1true"))

		// Boolean tests
		assertEval("true + true;", intVal(2))
		assertEval("true + false;", intVal(1))
		assertEval("true + NaN;", nanVal())
		assertEval("true + '1';", strVal("true1"))
	})

	Convey("Subtraction", t, func() {
		// Int tests
		assertEval("2 - 1;", intVal(1))
		assertEval("5 - 2;", intVal(3))
		assertEval("5 - 2 - 1;", intVal(2))
		assertEval("1 - 2;", intVal(-1))
		assertEval("1 - 2 - 3;", intVal(-4))
		assertEval("3 - '1';", intVal(2))
		assertEval("3 - 'a';", nanVal())
		assertEval("2 - NaN;", nanVal())
		assertEval("3 - true;", intVal(2))
		assertEval("3 - false;", intVal(3))

		// String tests
		assertEval("'1' - '1';", intVal(0))
		assertEval("'5' - '2' - '1';", intVal(2))
		assertEval("'3' - 1;", intVal(2))
		assertEval("'3' - 'a';", nanVal())
		assertEval("'3' - NaN;", nanVal())
		assertEval("'3' - true;", intVal(2))
		assertEval("'3' - false;", intVal(3))
		assertEval("'a' - 1;", nanVal())
		assertEval("'a' - '1';", nanVal())
		assertEval("'a' - 'a';", nanVal())
		assertEval("'a' - NaN;", nanVal())
		assertEval("'a' - true;", nanVal())
		assertEval("'a' - false;", nanVal())

		// NaN tests
		assertEval("NaN - 1;", nanVal())
		assertEval("NaN - '1';", nanVal())
		assertEval("NaN - true;", nanVal())
		assertEval("NaN - false;", nanVal())

		// Boolean tests
		assertEval("true - true;", intVal(0))
		assertEval("true - false;", intVal(1))
		assertEval("false - 1;", intVal(-1))
		assertEval("false - '1';", intVal(-1))
		assertEval("false - 'a';", nanVal())
		assertEval("false - NaN;", nanVal())
	})

	Convey("Assignment", t, func() {
		assertEval("x = 1;", intVal(1))
		assertEval("x = y = 1;", intVal(1))
		assertEval("x = y = z = 1;", intVal(1))
		assertEval("x = true;", boolVal(true))
	})

	Convey("Equality", t, func() {
		// Int tests
		assertEval("1 === 1;", boolVal(true))
		assertEval("1 === 2;", boolVal(false))
		assertEval("1 + 1 === 3 - 2;", boolVal(false))

		// String tests
		assertEval("'dog' === 'dog';", boolVal(true))
		assertEval("'dog' === 'cat';", boolVal(false))

		// Bool tests
		assertEval("true === true;", boolVal(true))
		assertEval("true === false;", boolVal(false))

		// NaN tests
		assertEval("NaN === NaN;", boolVal(false))
		assertEval("NaN === 'hi';", boolVal(false))
	})

	Convey("Inequality", t, func() {
		// Int tests
		assertEval("1 !== 1;", boolVal(false))
		assertEval("1 !== 2;", boolVal(true))
		assertEval("1 + 1 !== 3 - 2;", boolVal(false))

		// String tests
		assertEval("'dog' !== 'dog';", boolVal(false))
		assertEval("'dog' !== 'cat';", boolVal(true))

		// Bool tests
		assertEval("true !== true;", boolVal(false))
		assertEval("true !== false;", boolVal(true))

		// NaN tests
		assertEval("NaN !== NaN;", boolVal(true))
		assertEval("NaN !== 'hi';", boolVal(true))
	})
}
