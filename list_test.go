package containers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkedList(t *testing.T) {
	Convey("LinkedList test", t, func() {
		var ls = NewLinkedList[int]()
		ShouldBeTrue(ls.IsEmpty())
		ls.Add(1, 2, 3, 4, 5)
		So(ls.Size(), ShouldEqual, 5)
		ls.RemoveBack()
		So(ls.Size(), ShouldEqual, 4)
		ls.PushFront(6)
		So(ls.ToArray(), ShouldResemble, []int{6, 1, 2, 3, 4})
		ls.PushFront(7).RemoveFront().RemoveFront()
		So(ls.ToArray(), ShouldResemble, []int{1, 2, 3, 4})
		So(ls.Front(), ShouldEqual, 1)
		So(ls.Back(), ShouldEqual, 4)
	})
}

func TestEmpty(t *testing.T) {
	Convey("empty test", t, func() {
		ls := NewLinkedList[int]()
		ShouldPanic(func() {
			ls.Remove(2)
		})
		ShouldPanic(func() {
			ls.Get(0)
		})
		ShouldPanic(func() {
			ls.Front()
		})
		ShouldPanic(func() {
			ls.Back()
		})
		ShouldPanic(func() {
			ls.RemoveBack()
		})
		ShouldPanic(func() {
			ls.RemoveFront()
		})
		ShouldPanic(func() {
			ls.Set(1, 1)
		})
	})
}

func TestList(t *testing.T) {
	Convey("List test", t, func() {
		var ls List[int] = NewLinkedList[int]()
		// Add and Remove
		ls.Add(1, 2)
		ls.Remove(2)
		So(ls.ToArray(), ShouldResemble, []int{1})
		ls.Add(3, 4, 5, 6)
		ls.Remove(4)
		So(ls.ToArray(), ShouldResemble, []int{1, 3, 5, 6})
		ls.RemoveWith(func(lhs, rhs int) bool { return lhs == rhs }, 5)
		So(ls.ToArray(), ShouldResemble, []int{1, 3, 6})

		// Contains
		ShouldBeTrue(ls.Contains(3))
		ShouldBeTrue(ls.ContainsWith(func(lhs, rhs int) bool { return lhs == rhs }, 6))

		// Index(implement by IndexWith)
		idx, ok := ls.Index(3)
		So(idx, ShouldEqual, 1)
		ShouldBeTrue(ok)
		idx, ok = ls.Index(7)
		So(idx, ShouldEqual, -1)
		ShouldBeFalse(ok)

		// Get and Set
		So(ls.Get(1), ShouldEqual, 3)
		So(ls.Get(0), ShouldEqual, 1)
		ls.Set(1, 7)
		So(ls.Get(1), ShouldEqual, 7)

		// String
		So(ls.String(), ShouldEqual, "[1, 7, 6]")
		So(NewLinkedList[int]().String(), ShouldEqual, "[]")
	})
}
