package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProgram(t *testing.T) {
	Convey("Multiple Statements", t, func() {
		assertEval("1; 2;", intVal(2))
	})
}
