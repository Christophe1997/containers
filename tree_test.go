package containers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRBTree_debug(t *testing.T) {
	Convey("Init data", t, func() {
		tree1 := &rbTree[int]{}
		ShouldBeTrue(tree1.IsEmpty())

		node1 := &treeNode[int]{
			left:   nil,
			right:  nil,
			parent: nil,
			isRed:  false,
			val:    0,
		}
		node2 := &treeNode[int]{
			left:   nil,
			right:  nil,
			parent: nil,
			isRed:  true,
			val:    1,
		}
		node1.left = node2
		node2.parent = node1
		tree2 := &rbTree[int]{root: node1}
		ShouldBeTrue(tree2.debug())
	})
}
