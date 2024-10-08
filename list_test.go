package lisq

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var add = func(a, b int) int { return a + b }

func TestFold(t *testing.T) {
	Convey("empty case", t, func() {
		So(Fold(add, 0, Zero[int]()), ShouldEqual, 0)
	})
	Convey("finite case", t, func() {

	})
}
