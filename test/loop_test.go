package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoops(t *testing.T) {
	Convey("while", t, func() {
		assertEval("x = 0; while(false) { x = x + 1 }; x;", intVal(0))
		assertEval("x = 0; while(x < 3) { x = x + 1 }; x;", intVal(3))
	})
}
