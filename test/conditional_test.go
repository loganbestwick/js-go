package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConditionals(t *testing.T) {
	Convey("if", t, func() {
		// Bool Tests
		assertEval("if (true) { 1; }", intVal(1))
		assertEval("a = 1; if (true) { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; if (false) { a = a + 1; } a;", intVal(1))

		// Num Tests
		assertEval("a = 1; if (1) { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; if (1 - 1) { a = a + 1; } a;", intVal(1))
		assertEval("a = 1; if (1 + 1) { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; if (NaN) { a = a + 1; } a;", intVal(1))

		// Identifier Tests
		assertEval("a = 1; b = true; if (b) { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; b = true; if (b = false) { a = a + 1; } a;", intVal(1))
		assertEval("a = 1; b = false; if (b) { a = a + 1; } a;", intVal(1))

		// String Tests
		assertEval("a = 1; if ('hello') { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; if ('') { a = a + 1; } a;", intVal(1))
		assertEval("a = 1; if ('hello' + NaN) { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; if ('hello' + 1) { a = a + 1; } a;", intVal(2))
	})
}
