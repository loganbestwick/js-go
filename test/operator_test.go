package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOperators(t *testing.T) {
	Convey("Addition", t, func() {
		assertEval("1 + 1", intVal(2))
		assertEval("1 + 2", intVal(3))

		assertEval("'2' + 1", strVal("21"))
		assertEval("1 + '2'", strVal("12"))
	})
}
