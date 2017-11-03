package test

import (
	"testing"

	"github.com/loganbestwick/js-go/types"
	. "github.com/smartystreets/goconvey/convey"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestInterop(t *testing.T) {
	Convey("interop", t, func() {
		ctx := &types.Context{}
		assertEvalWithContext(ctx, "max(1, 2);", intVal(2))
	})
}
