package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoops(t *testing.T) {
	Convey("while", t, func() {
		assertEval("x = 0; while(false) { x = x + 1; } x;", intVal(0))
		assertEval("x = 0; while(x < 3) { x = x + 1; } x;", intVal(3))
		assertEval("x = 0; while(x < 2) { x = x + 1; }", intVal(2))
	})

	Convey("for", t, func() {
		assertEval("z = 0; for(x = 0; x < 1; x = x + 1) { z = 10; } z;", intVal(10))
		assertEval("z = 0; for(x = 0; x < 1; x = x + 1) { z = 10; }", intVal(10))
		assertEval("z = 0; for(x = 0; x < 20; x = x + 1) { z = z + 1; } z;", intVal(20))

		assertEval("for(x = 0; x < 5; x = x + 1) { x; } x;", intVal(5))
		assertEval("for(x = 0; x < 0; x = x + 1) { x; } x;", intVal(0))
	})
}
