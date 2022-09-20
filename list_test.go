package containers

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedList_Size(t *testing.T) {
	Convey("Init data", t, func() {
		var ls = NewLinkedList[int]()
		ls.Add(1, 2, 3, 4, 5)
		fmt.Println(ls)
	})
}
