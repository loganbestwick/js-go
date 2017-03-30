package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstants(t *testing.T) {
	Convey("Numbers", t, func() {
		assertEval("1", intVal(1))
	})

	Convey("NaN", t, func() {
		assertEval("NaN", nanVal())
	})

	Convey("Strings", t, func() {
		assertEval("'1'", strVal("1"))
	})

	Convey("Identifier", t, func() {
		assertEval("x = 1; x", intVal(1))
		assertError("x", "x is not defined")
	})
}
