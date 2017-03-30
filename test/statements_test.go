package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStatements(t *testing.T) {
	Convey("Statements", t, func() {
		assertEval("1", intVal(1))
		assertEval("1; 2", intVal(2))
		assertEval("1; 2;", intVal(2))
	})
}
