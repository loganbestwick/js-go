package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAssignment(t *testing.T) {
	Convey("Variable assignment", t, func() {
		assertEval("x = 1", intVal(1))
		assertEval("x = y = 1", intVal(1))
		assertEval("x = y = z = 1", intVal(1))

		assertEval("x = 1; y = x + 1", intVal(2))

		assertEval("x = 1; y = x; x = 2; y", intVal(1))
		assertEval("x = 1; y = x; x = 2; y + 0", intVal(1))
	})
}
