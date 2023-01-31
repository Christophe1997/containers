package containers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRBTree_debug(t *testing.T) {
	Convey("Init data", t, func() {

		tree0 := &rbTree[int]{}

		node11 := &treeNode[int]{
			isRed: false,
			val:   0,
		}
		node12 := &treeNode[int]{
			isRed: true,
			val:   1,
		}
		node11.left = node12
		node12.parent = node11
		tree1 := &rbTree[int]{root: node11}

		node21 := &treeNode[int]{isRed: false, val: 7}
		node22 := &treeNode[int]{isRed: false, val: 3}
		node23 := &treeNode[int]{isRed: true, val: 18}
		node24 := &treeNode[int]{isRed: false, val: 10}
		node25 := &treeNode[int]{isRed: false, val: 22}
		node26 := &treeNode[int]{isRed: true, val: 8}
		node27 := &treeNode[int]{isRed: true, val: 11}
		node28 := &treeNode[int]{isRed: true, val: 26}

		node21.left = node22
		node22.parent = node21
		node21.right = node23
		node23.parent = node21
		node23.left = node24
		node24.parent = node23
		node23.right = node25
		node25.parent = node23
		node24.left = node26
		node26.parent = node24
		node24.right = node27
		node27.parent = node24
		node25.right = node28
		node28.parent = node25

		tree2 := &rbTree[int]{root: node21}

		Convey("rbtree test", func() {
			So(tree0.IsEmpty(), ShouldBeTrue)
			So(tree1.debug(), ShouldBeTrue)
			So(tree2.debug(), ShouldBeTrue)
		})

	})
}

func TestRbTree_InsertWith(t *testing.T) {
	Convey("Init data", t, func() {
		cmp := func(l, r int) int {
			return l - r
		}

		Convey("rbtree insert", func() {
			tree0 := &rbTree[int]{}
			tree0.InsertWith(cmp, 7)
			So(tree0.debug(), ShouldBeTrue)
			tree0.InsertWith(cmp, 18)
			So(tree0.debug(), ShouldBeTrue)

			tree0.InsertWith(cmp, 3, 10, 22, 8, 11, 26)
			So(tree0.debug(), ShouldBeTrue)
		})
	})
}
