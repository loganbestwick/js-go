package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstants(t *testing.T) {
	Convey("Numbers", t, func() {
		assertEval("1", intVal(1))
	})

	Convey("Strings", t, func() {
		assertEval("'1'", strVal("1"))
	})
}
